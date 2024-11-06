package billing_usecase

import (
	"errors"
	"fmt"

	"github.com/herryg91/billing/rest-api/app/entity"
	"github.com/herryg91/billing/rest-api/app/repository/billing_repository"
	"github.com/herryg91/billing/rest-api/app/repository/loan_repository"
	"gorm.io/gorm"
)

type usecase struct {
	billing_repo billing_repository.Repository
	loan_repo    loan_repository.Repository
}

func New(billing_repo billing_repository.Repository, loan_repo loan_repository.Repository) UseCase {
	return &usecase{billing_repo: billing_repo, loan_repo: loan_repo}
}

func (uc *usecase) GetBillingByLoanCode(user_id int, loan_code string) ([]entity.Billing, error) {
	current_loan, err := uc.checkLoanAuthorized(loan_code, user_id)
	if err != nil {
		return []entity.Billing{}, err
	}

	billings, err := uc.billing_repo.GetByLoanId(current_loan.Id)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrUnexpected, err.Error())
	}

	return billings, nil
}

func (uc *usecase) GetBillingOverDue(user_id int) ([]entity.Billing, error) {
	billings, err := uc.billing_repo.GetOverDueByUserId(user_id, 7)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrUnexpected, err.Error())
	}

	return billings, nil
}

func (uc *usecase) GeneratePaymentInfo(user_id int, loan_code string, installment_number int) (*entity.Billing, error) {
	l, err := uc.checkLoanAuthorized(loan_code, user_id)
	if err != nil {
		return nil, err
	}

	// Get respective billing
	b, err := uc.billing_repo.GetByInstallmentNumber(l.Id, installment_number)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("%w: %s", ErrUnexpected, err.Error())
	}
	b.LoanCode = loan_code

	// Check if it need to generate payment or not. If status UNPAID or WAIT_FOR_PAYMENT but xpired, then it need generated
	is_payment_update := b.GeneratePaymentInfo()
	if is_payment_update {
		err = uc.billing_repo.UpdatePayment(*b)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrUnexpected, err.Error())
		}
	}
	return b, nil
}

func (uc *usecase) SettlePayment(user_id int, loan_code string, installment_number int) error {
	l, err := uc.checkLoanAuthorized(loan_code, user_id)
	if err != nil {
		return err
	}

	// 1. Get Current Billing
	current_billing, err := uc.billing_repo.GetByInstallmentNumber(l.Id, installment_number)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return fmt.Errorf("%w: %s", ErrUnexpected, err.Error())
	}
	// 2. Skip if already paid
	if current_billing.PaymentStatus == entity.BillingPaymentStatus_Paid {
		return nil
	}

	// 3. Update the billing into paid
	current_billing.PaymentStatus = entity.BillingPaymentStatus_Paid
	current_billing.PaymentRef = "{...}"
	err = uc.billing_repo.UpdatePayment(*current_billing)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrUnexpected, err.Error())
	}

	// 4. Check if all billings has paid
	is_all_paid := true
	all_billings, err := uc.billing_repo.GetByLoanId(l.Id)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrUnexpected, err.Error())
	}
	for _, b := range all_billings {
		if b.PaymentStatus != entity.BillingPaymentStatus_Paid {
			is_all_paid = false
			break
		}
	}

	// 5. If all paid, change loan status into DONE
	if is_all_paid {
		uc.loan_repo.UpdateStatus(l.Id, entity.LoanStatus_Done)
	}

	return nil
}
