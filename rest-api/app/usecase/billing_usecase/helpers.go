package billing_usecase

import (
	"errors"
	"fmt"

	"github.com/herryg91/billing/rest-api/app/entity"
	"gorm.io/gorm"
)

func (uc *usecase) checkLoanAuthorized(loan_code string, user_id int) (*entity.Loan, error) {
	l, err := uc.loan_repo.GetByCode(loan_code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUnauthorized
		}
		return nil, fmt.Errorf("%w: %s", ErrUnexpected, err.Error())
	}
	if l.UserId != user_id {
		return nil, ErrUnauthorized
	}
	return l, nil
}
