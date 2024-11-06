package entity

import (
	"math"
	"time"

	"github.com/herryg91/billing/rest-api/pkg/helpers"
)

type InstallmentCycle string
type LoanStatus string
type InterestType string

const (
	InstallmentCycle_Weekly InstallmentCycle = "WEEKLY"

	InterestType_Flat InterestType = "FLAT"

	LoanStatus_Pending  LoanStatus = "PENDING"
	LoanStatus_Approved LoanStatus = "APPROVED"
	LoanStatus_Active   LoanStatus = "ACTIVE"
	LoanStatus_Done     LoanStatus = "DONE"
)

type Loan struct {
	Id                int              `json:"id"`
	Code              string           `json:"code"`
	UserId            int              `json:"user_id"`
	Description       string           `json:"description"`
	InstallmentCycle  InstallmentCycle `json:"installment_cycle"`
	InstallmentLength int              `json:"installment_length"`
	InterestType      InterestType     `json:"interestType"`
	InterestPercent   float64          `json:"interest_percent"`
	Principal         float64          `json:"amount"`
	InterestAmount    float64          `json:"interest_amount"`
	TotalAmount       float64          `json:"total_amount"`
	Status            LoanStatus       `json:"status"`
	CreatedAt         time.Time        `json:"createdAt"`
}

type LoanWithOutstanding struct {
	Loan
	Outstanding float64
}

func (l Loan) CalculateTotalInterestFlat() float64 {
	// We are using 50 week structure instead of traditional 52 week year
	return math.Ceil(l.Principal * (l.InterestPercent / 100) * (float64(l.InstallmentLength) / 50))
}

func (l Loan) SimulateBilling(baseline_date time.Time) []Billing {
	l.InterestAmount = l.CalculateTotalInterestFlat()
	principal_installment_amount := math.Ceil(l.Principal / float64(l.InstallmentLength))
	interest_installment_amount := math.Ceil(l.InterestAmount / float64(l.InstallmentLength))

	resp := []Billing{}

	cumm_principal := float64(0)
	cumm_interest := float64(0)
	installment_date := baseline_date
	for i := 0; i < l.InstallmentLength; i++ {
		installment_date = installment_date.AddDate(0, 0, 7)
		if i == l.InstallmentLength-1 {
			// Installment n (last installment)
			resp = append(resp, Billing{
				LoanId:            l.Id,
				InstallmentNumber: i + 1,
				DueDate:           helpers.Date(installment_date),
				Principal:         l.Principal - cumm_principal,
				InterestAmount:    l.InterestAmount - cumm_interest,
				TotalAmount:       (l.Principal - cumm_principal) + (l.InterestAmount - cumm_interest),
				PaymentBank:       "BCA",
				PaymentVA:         "",
				PaymentStatus:     BillingPaymentStatus_Unpaid,
				PaymentExpiredAt:  time.Now(),
				PaymentRef:        "",
			})
		} else {
			// Installment 1 ... n-1
			cumm_principal += principal_installment_amount
			cumm_interest += interest_installment_amount
			resp = append(resp, Billing{
				LoanId:            l.Id,
				InstallmentNumber: i + 1,
				DueDate:           helpers.Date(installment_date),
				Principal:         principal_installment_amount,
				InterestAmount:    interest_installment_amount,
				TotalAmount:       principal_installment_amount + interest_installment_amount,
				PaymentBank:       "BCA",
				PaymentVA:         "",
				PaymentStatus:     BillingPaymentStatus_Unpaid,
				PaymentExpiredAt:  time.Now(),
				PaymentRef:        "",
			})
		}

	}
	return resp
}
