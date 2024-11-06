package loan_usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/herryg91/billing/rest-api/app/entity"
	"github.com/herryg91/billing/rest-api/app/repository/billing_repository"
	"github.com/herryg91/billing/rest-api/app/repository/loan_repository"
	"gorm.io/gorm"
)

type usecase struct {
	loan_repo        loan_repository.Repository
	billing_repo     billing_repository.Repository
	flatInterestRate float64
}

func New(loan_repo loan_repository.Repository, billing_repo billing_repository.Repository, flatInterestRate float64) UseCase {
	return &usecase{
		loan_repo:        loan_repo,
		billing_repo:     billing_repo,
		flatInterestRate: flatInterestRate,
	}
}

func (uc *usecase) SimulateLoan(l entity.Loan) (entity.Loan, []entity.Billing) {
	l.InterestPercent = uc.flatInterestRate
	l.InstallmentCycle = entity.InstallmentCycle_Weekly
	billings := l.SimulateBilling(time.Now())
	l.InterestAmount = l.CalculateTotalInterestFlat()
	l.TotalAmount = l.Principal + l.InterestAmount
	return l, billings
}

func (uc *usecase) CreateLoanRequest(l entity.Loan) error {
	var err error
	l.Code, err = uc.generateLoanCode()
	if err != nil {
		return err
	}
	l.InterestPercent = uc.flatInterestRate
	l.InstallmentCycle = entity.InstallmentCycle_Weekly
	l.InterestType = entity.InterestType_Flat
	l.InterestAmount = l.CalculateTotalInterestFlat()
	l.TotalAmount = l.Principal + l.InterestAmount
	l.Status = entity.LoanStatus_Approved // entity.LoanStatus_Pending
	// 	Approval Process was out of scope, therefore we assume that it is auto approved and generate the billing schedule

	loanId, err := uc.loan_repo.Create(l)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrUnexpected, err.Error())
	}
	l.Id = loanId

	return uc.DisburseAndCreateSchedule(l)
}

func (uc *usecase) DisburseAndCreateSchedule(l entity.Loan) error {
	// 1. Generate Billing Schedule
	billings := l.SimulateBilling(time.Now())
	for idx := range billings {
		billings[idx].LoanId = l.Id
	}
	err := uc.billing_repo.Create(billings)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrUnexpected, err.Error())
	}

	// 2. Update Disbursed At
	time_now := time.Now()
	err = uc.loan_repo.UpdateDisburse(l.Id, &time_now)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrUnexpected, err.Error())
	}

	// 3. Update Loan Status to ACTIVE
	err = uc.loan_repo.UpdateStatus(l.Id, entity.LoanStatus_Active)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrUnexpected, err.Error())
	}

	return nil
}

func (uc *usecase) GetLoans(user_id int) ([]entity.LoanWithOutstanding, error) {
	loans, err := uc.loan_repo.GetByUserId(user_id)
	if err != nil {
		return []entity.LoanWithOutstanding{}, fmt.Errorf("%w: %s", ErrUnexpected, err.Error())
	}

	for idx := range loans {
		if loans[idx].Status != entity.LoanStatus_Active {
			loans[idx].Outstanding = 0
		}
	}

	return loans, nil
}

func (uc *usecase) GetUserSummary(user_id int) (loan_count int, total_outstanding float64, is_delinquent bool, err error) {
	loan_count, total_outstanding, is_delinquent, err = 0, 0, false, nil
	// 1. Find active loans
	active_loans, err := uc.loan_repo.GetByStatus(user_id, entity.LoanStatus_Active)
	if err != nil {
		return 0, 0, false, fmt.Errorf("%w: %s", ErrUnexpected, err.Error())
	}
	loan_count = len(active_loans)

	active_loan_ids := []int{}
	for _, l := range active_loans {
		active_loan_ids = append(active_loan_ids, l.Id)
	}
	// 2. Find & calculate total outstandings
	outstandings, err := uc.billing_repo.GetOutstandings(active_loan_ids)
	if err != nil {
		return 0, 0, false, fmt.Errorf("%w: %s", ErrUnexpected, err.Error())
	}
	for _, amount := range outstandings {
		total_outstanding += amount
	}

	// 3. Check delinquent
	exceed_due_date_info, err := uc.billing_repo.GetExceedDueDate(active_loan_ids)
	if err != nil {
		return 0, 0, false, fmt.Errorf("%w: %s", ErrUnexpected, err.Error())
	}
	for _, count := range exceed_due_date_info {
		if count > 1 {
			is_delinquent = true
			break
		}
	}

	return
}

func (uc *usecase) GetLoanByCode(user_id int, code string) (*entity.Loan, error) {
	current_loan, err := uc.loan_repo.GetByCode(code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("%w: %s", ErrUnexpected, err.Error())
	}

	if current_loan.UserId != user_id {
		return nil, ErrUnauthorized
	}

	return current_loan, nil
}
