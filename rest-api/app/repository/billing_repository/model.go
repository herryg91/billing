package billing_repository

import (
	"time"

	"github.com/herryg91/billing/rest-api/app/entity"
	"github.com/herryg91/billing/rest-api/pkg/helpers"
)

type LoanBillingModel struct {
	Id                int          `gorm:"primary_key;column:id"`
	LoanId            int          `gorm:"column:loan_id"`
	LoanCode          string       `gorm:"column:loan_code;->"`
	InstallmentNumber int          `gorm:"column:installment_number"`
	DueDate           helpers.Date `gorm:"column:due_date"`
	Principal         float64      `gorm:"column:principal"`
	InterestAmount    float64      `gorm:"column:interest_amount"`

	PaymentBank      string                      `gorm:"column:payment_bank"`
	PaymentVA        string                      `gorm:"column:payment_va"`
	PaymentStatus    entity.BillingPaymentStatus `gorm:"column:payment_status"`
	PaymentExpiredAt time.Time                   `gorm:"column:payment_expired_at"`
	PaymentRef       string                      `gorm:"column:payment_ref"`

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (LoanBillingModel) New(b entity.Billing) *LoanBillingModel {
	return &LoanBillingModel{
		Id:                b.Id,
		LoanId:            b.LoanId,
		LoanCode:          b.LoanCode,
		InstallmentNumber: b.InstallmentNumber,
		DueDate:           b.DueDate,
		Principal:         b.Principal,
		InterestAmount:    b.InterestAmount,
		PaymentBank:       b.PaymentBank,
		PaymentVA:         b.PaymentVA,
		PaymentStatus:     b.PaymentStatus,
		PaymentExpiredAt:  b.PaymentExpiredAt,
		PaymentRef:        b.PaymentRef,
	}
}

func (b *LoanBillingModel) Parse() *entity.Billing {
	resp := &entity.Billing{
		Id:                b.Id,
		LoanId:            b.LoanId,
		LoanCode:          b.LoanCode,
		InstallmentNumber: b.InstallmentNumber,
		DueDate:           b.DueDate,
		Principal:         b.Principal,
		InterestAmount:    b.InterestAmount,
		TotalAmount:       b.Principal + b.InterestAmount,
		PaymentBank:       b.PaymentBank,
		PaymentVA:         b.PaymentVA,
		PaymentStatus:     b.PaymentStatus,
		PaymentExpiredAt:  b.PaymentExpiredAt,
		PaymentRef:        b.PaymentRef,
	}
	resp.SetPaymentStatus()
	return resp
}
