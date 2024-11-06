package billing_repository

import (
	"time"

	"github.com/herryg91/billing/rest-api/app/entity"
	"gorm.io/gorm"
)

type Repository interface {
	Create(req []entity.Billing) error
	GetOutstandings(loan_ids []int) (map[int]float64, error)
	GetExceedDueDate(loan_ids []int) (map[int]int, error)

	GetOverDueByUserId(user_id int, add_days int) ([]entity.Billing, error)
	GetByInstallmentNumber(loan_id int, installment_number int) (*entity.Billing, error)

	UpdatePayment(req entity.Billing) error

	GetByLoanId(loan_id int) ([]entity.Billing, error)
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(req []entity.Billing) error {
	models := []LoanBillingModel{}
	for _, r := range req {
		b := LoanBillingModel{}.New(r)
		b.CreatedAt = time.Now()
		b.UpdatedAt = time.Now()
		models = append(models, *b)
	}
	err := r.db.Table("loan_billing").Create(&models).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetByLoanId(loan_id int) ([]entity.Billing, error) {
	models := []LoanBillingModel{}
	err := r.db.Table("loan_billing").Where("loan_id = ?", loan_id).Order("installment_number asc").Find(&models).Error
	if err != nil {
		return []entity.Billing{}, err
	}
	resp := []entity.Billing{}
	for _, m := range models {
		resp = append(resp, *m.Parse())
	}
	return resp, nil
}

type OutstandingsModel struct {
	LoanId      int     `gorm:"column:loan_id"`
	Outstanding float64 `gorm:"column:total_outstanding"`
}

func (r *repository) GetOutstandings(loan_ids []int) (map[int]float64, error) {
	resp := map[int]float64{}
	if len(loan_ids) == 0 {
		return resp, nil
	}
	models := []OutstandingsModel{}
	err := r.db.Raw(`SELECT loan_id, sum(principal + interest_amount) as total_outstanding 
	FROM loan_billing  
	WHERE loan_id in (?)
	AND payment_status != 'PAID'
	GROUP BY loan_id
	`, loan_ids).Find(&models).Error
	if err != nil {
		return resp, err
	}

	for _, m := range models {
		resp[m.LoanId] = m.Outstanding
	}
	return resp, nil
}

type ExceedDueDateModel struct {
	LoanId             int `gorm:"column:loan_id"`
	TotalExceedDueDate int `gorm:"column:total_exceed_due_date"`
}

func (r *repository) GetExceedDueDate(loan_ids []int) (map[int]int, error) {
	resp := map[int]int{}
	if len(loan_ids) == 0 {
		return resp, nil
	}
	models := []ExceedDueDateModel{}
	err := r.db.Raw(`SELECT loan_id, count(id) as total_exceed_due_date 
	FROM loan_billing  
	WHERE loan_id in (?)
	AND due_date < now()::date
	and payment_status != 'PAID'
	GROUP BY loan_id
	`, loan_ids).Find(&models).Error
	if err != nil {
		return resp, err
	}

	for _, m := range models {
		resp[m.LoanId] = m.TotalExceedDueDate
	}
	return resp, nil
}

func (r *repository) GetOverDueByUserId(user_id int, add_days int) ([]entity.Billing, error) {
	models := []LoanBillingModel{}
	q := `SELECT lb.*, l.code as loan_code FROM loan l 
		inner join loan_billing lb on l.id = lb.loan_id 
		where l.user_id = ?
		and l.status = 'ACTIVE'
		and lb.payment_status != 'PAID'
		and lb.due_date < NOW()`
	if add_days > 0 {
		q += ` + interval '7' day`
	}
	err := r.db.Raw(q, user_id).Find(&models).Error
	if err != nil {
		return []entity.Billing{}, err
	}
	resp := []entity.Billing{}
	for _, m := range models {
		resp = append(resp, *m.Parse())
	}
	return resp, nil
}

func (r *repository) GetByInstallmentNumber(loan_id int, installment_number int) (*entity.Billing, error) {
	model := LoanBillingModel{}
	err := r.db.Table("loan_billing").Where("loan_id = ? AND installment_number = ?", loan_id, installment_number).First(&model).Error
	if err != nil {
		return nil, err
	}

	return model.Parse(), nil
}

func (r *repository) UpdatePayment(req entity.Billing) error {
	err := r.db.Table("loan_billing").Where("id = ?", req.Id).Updates(map[string]interface{}{
		"payment_bank":       req.PaymentBank,
		"payment_va":         req.PaymentVA,
		"payment_status":     req.PaymentStatus,
		"payment_expired_at": req.PaymentExpiredAt,
		"payment_ref":        req.PaymentRef,
		"updated_at":         time.Now(),
	}).Error
	if err != nil {
		return err
	}

	return nil
}
