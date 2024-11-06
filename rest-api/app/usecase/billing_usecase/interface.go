package billing_usecase

import (
	"errors"

	"github.com/herryg91/billing/rest-api/app/entity"
)

var ErrUnexpected = errors.New("Unexpected internal error")
var ErrNotFound = errors.New("Billing not found")
var ErrUnauthorized = errors.New("You are unauthorized to access this data")

type UseCase interface {
	GetBillingByLoanCode(user_id int, loan_code string) ([]entity.Billing, error)
	GetBillingOverDue(user_id int) ([]entity.Billing, error)

	GeneratePaymentInfo(user_id int, loan_code string, installment_number int) (*entity.Billing, error)
	SettlePayment(user_id int, loan_code string, installment_number int) error
}
