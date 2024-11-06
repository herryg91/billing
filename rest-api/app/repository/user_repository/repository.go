package user_repository

import (
	"errors"
	"time"

	"github.com/herryg91/billing/rest-api/app/entity"
	"github.com/jackc/pgconn"

	"gorm.io/gorm"
)

type Repository interface {
	GetByEmail(email string) (*entity.User, error)
	Create(param entity.User) (int, error)
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) GetByEmail(email string) (*entity.User, error) {
	data := &UserModel{}
	err := r.db.Table("users").Where("email = ?", email).First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return data.ToUser(), nil
}

func (r *repository) Create(param entity.User) (int, error) {
	model := UserModel{}.New(param)
	model.CreatedAt = time.Now()
	model.UpdatedAt = time.Now()
	err := r.db.Table("users").Create(&model).Error
	if err != nil {
		pqerr := &pgconn.PgError{}
		if errors.As(err, &pqerr) && pqerr.Code == "23505" {
			return 0, ErrDuplicateEmail
		}
		return 0, err
	}
	return model.Id, nil
}
