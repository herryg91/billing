package loan_repository

import (
	"time"

	"github.com/herryg91/billing/rest-api/app/entity"
)

type LoanModel struct {
	Id                int                     `gorm:"primary_key;column:id"`
	Code              string                  `gorm:"column:code"`
	UserId            int                     `gorm:"column:user_id"`
	Description       string                  `gorm:"column:description"`
	InstallmentCycle  entity.InstallmentCycle `gorm:"column:installment_cycle"`
	InstallmentLength int                     `gorm:"column:installment_length"`
	InterestType      entity.InterestType     `gorm:"column:interest_type"`
	InterestPercent   float64                 `gorm:"column:interest_percent"`
	Principal         float64                 `gorm:"column:principal"`
	InterestAmount    float64                 `gorm:"column:interest_amount"`
	TotalAmount       float64                 `gorm:"column:total_amount"`
	Status            entity.LoanStatus       `gorm:"column:status"`
	DisbursedAt       *time.Time              `gorm:"column:disbursed_at"`
	CreatedAt         time.Time               `gorm:"column:created_at"`
	UpdatedAt         time.Time               `gorm:"column:updated_at"`

	Outstanding float64 `gorm:"column:outstanding;->"`
}

func (LoanModel) New(l entity.Loan) *LoanModel {
	return &LoanModel{
		Id:                l.Id,
		Code:              l.Code,
		UserId:            l.UserId,
		Description:       l.Description,
		InstallmentCycle:  l.InstallmentCycle,
		InstallmentLength: l.InstallmentLength,
		InterestType:      l.InterestType,
		InterestPercent:   l.InterestPercent,
		Principal:         l.Principal,
		InterestAmount:    l.InterestAmount,
		TotalAmount:       l.TotalAmount,
		Status:            l.Status,
		CreatedAt:         l.CreatedAt,
	}
}

func (l *LoanModel) Parse() *entity.Loan {
	return &entity.Loan{
		Id:                l.Id,
		Code:              l.Code,
		UserId:            l.UserId,
		Description:       l.Description,
		InstallmentCycle:  l.InstallmentCycle,
		InstallmentLength: l.InstallmentLength,
		InterestType:      l.InterestType,
		InterestPercent:   l.InterestPercent,
		Principal:         l.Principal,
		InterestAmount:    l.InterestAmount,
		TotalAmount:       l.TotalAmount,
		Status:            l.Status,
		CreatedAt:         l.CreatedAt,
	}
}
