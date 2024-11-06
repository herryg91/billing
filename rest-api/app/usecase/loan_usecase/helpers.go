package loan_usecase

import (
	"fmt"
	"strconv"
	"time"
)

func (uc *usecase) generateLoanCode() (string, error) {
	prefix := "LN-" + strconv.Itoa(time.Now().Year()) + "-"
	_, number, err := uc.loan_repo.LastCode(prefix)
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrUnexpected, err.Error())
	}
	increment_number_str := "00000" + strconv.Itoa(number+1)
	increment_number_str = increment_number_str[len(increment_number_str)-5:]
	return prefix + increment_number_str, nil
}
