package loan_repository

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/herryg91/billing/rest-api/app/entity"
	"gorm.io/gorm"
)

var ErrNotFound = errors.New("User not found")

type Repository interface {
	LastCode(prefix string) (string, int, error)
	Create(req entity.Loan) (int, error)
	UpdateStatus(id int, status entity.LoanStatus) error
	UpdateDisburse(id int, disbursed_at *time.Time) error
	GetByStatus(user_id int, status entity.LoanStatus) ([]entity.Loan, error)
	GetByCode(code string) (*entity.Loan, error)

	// Query With Outstanding
	GetByUserId(user_id int) ([]entity.LoanWithOutstanding, error)
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) LastCode(prefix string) (string, int, error) {
	model := &LoanModel{}
	err := r.db.Table("loan").Where("code like ?", prefix+"%").Order("code desc").First(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", 0, nil
		}
		return "", 0, err
	}
	increment_number_str := strings.ReplaceAll(model.Code, prefix, "")
	increment_number, _ := strconv.Atoi(increment_number_str)
	return model.Code, increment_number, nil
}

func (r *repository) Create(req entity.Loan) (int, error) {
	model := LoanModel{}.New(req)
	model.CreatedAt = time.Now()
	model.UpdatedAt = time.Now()
	err := r.db.Table("loan").Create(&model).Error
	if err != nil {
		return 0, err
	}
	return model.Id, nil
}
func (r *repository) UpdateStatus(id int, status entity.LoanStatus) error {
	err := r.db.Table("loan").Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":     status,
			"updated_at": time.Now(),
		}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateDisburse(id int, disbursed_at *time.Time) error {
	err := r.db.Table("loan").Where("id = ?", id).
		Updates(map[string]interface{}{
			"disbursed_at": disbursed_at,
			"updated_at":   time.Now(),
		}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetByStatus(user_id int, status entity.LoanStatus) ([]entity.Loan, error) {
	models := []LoanModel{}
	err := r.db.Table("loan").Where("user_id = ? AND status = ?", user_id, status).Find(&models).Error
	if err != nil {
		return []entity.Loan{}, err
	}
	resp := []entity.Loan{}
	for _, m := range models {
		resp = append(resp, *m.Parse())
	}
	return resp, nil
}

func (r *repository) GetByCode(code string) (*entity.Loan, error) {
	model := LoanModel{}
	err := r.db.Table("loan").Where("code = ?", code).First(&model).Error
	if err != nil {
		return nil, err
	}

	return model.Parse(), nil
}

func (r *repository) GetByUserId(user_id int) ([]entity.LoanWithOutstanding, error) {
	models := []LoanModel{}
	err := r.db.Raw(`SELECT 
		l.*, 
		(select sum(lb.principal+lb.interest_amount) from loan_billing lb where l.id = lb.loan_id and lb.payment_status != 'PAID') as outstanding
	FROM loan l 
	WHERE l.user_id = ?
	ORDER BY l.created_at desc`, user_id).Find(&models).Error
	if err != nil {
		return []entity.LoanWithOutstanding{}, err
	}
	resp := []entity.LoanWithOutstanding{}
	for _, m := range models {
		resp = append(resp, entity.LoanWithOutstanding{
			Loan:        *m.Parse(),
			Outstanding: m.Outstanding,
		})
	}
	return resp, nil
}
