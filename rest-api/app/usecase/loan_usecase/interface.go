package loan_usecase

import (
	"errors"

	"github.com/herryg91/billing/rest-api/app/entity"
)

var ErrUnexpected = errors.New("Unexpected internal error")
var ErrUnauthorized = errors.New("You are unauthorized to access this data")
var ErrNotFound = errors.New("not found")

type UseCase interface {
	SimulateLoan(l entity.Loan) (entity.Loan, []entity.Billing)
	CreateLoanRequest(l entity.Loan) error
	DisburseAndCreateSchedule(l entity.Loan) error

	GetUserSummary(user_id int) (loan_count int, total_outstanding float64, is_delinquent bool, err error)
	GetLoans(user_id int) ([]entity.LoanWithOutstanding, error)
	GetLoanByCode(user_id int, code string) (*entity.Loan, error)
}
