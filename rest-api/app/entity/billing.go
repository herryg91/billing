package entity

import (
	"time"

	"github.com/herryg91/billing/rest-api/pkg/helpers"
)

type BillingPaymentStatus string

const (
	BillingPaymentStatus_Unpaid         BillingPaymentStatus = "UNPAID"
	BillingPaymentStatus_WaitForPayment BillingPaymentStatus = "WAIT_FOR_PAYMENT"
	BillingPaymentStatus_Paid           BillingPaymentStatus = "PAID"
)

type Billing struct {
	Id                int          `json:"id"`
	LoanId            int          `json:"loan_id"`
	LoanCode          string       `json:"loan_code"`
	InstallmentNumber int          `json:"installment_number"`
	DueDate           helpers.Date `json:"due_date"`
	Principal         float64      `json:"principal"`
	InterestAmount    float64      `json:"interest_amount"`
	TotalAmount       float64      `json:"total_amount"`

	// Payment
	PaymentBank      string               `json:"payment_bank"`
	PaymentVA        string               `json:"payment_va"`
	PaymentStatus    BillingPaymentStatus `json:"payment_status"`
	PaymentExpiredAt time.Time            `json:"payment_expired_at"`
	PaymentRef       string               `json:"payment_ref"`
}

func (b *Billing) GeneratePaymentInfo() (update_payment bool) {
	update_payment = false
	if b == nil {
		return
	} else if b.PaymentStatus == BillingPaymentStatus_Paid {
		return
	} else if b.PaymentStatus == BillingPaymentStatus_WaitForPayment && b.PaymentExpiredAt.Unix() >= time.Now().Unix() {
		return
	}

	b.PaymentBank = "BCA"
	b.PaymentVA = helpers.RandomStringIntOnly(10)
	b.PaymentStatus = BillingPaymentStatus_WaitForPayment
	b.PaymentExpiredAt = time.Now().Add(time.Hour * 6)
	update_payment = true
	return
}

func (b *Billing) SetPaymentStatus() {
	if b == nil {
		return
	}

	if b.PaymentStatus == BillingPaymentStatus_WaitForPayment && b.PaymentExpiredAt.Unix() < time.Now().Unix() {
		b.PaymentStatus = BillingPaymentStatus_Unpaid
	}
}
