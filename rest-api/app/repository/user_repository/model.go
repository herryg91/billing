package user_repository

import (
	"time"

	"github.com/herryg91/billing/rest-api/app/entity"
)

type UserModel struct {
	Id        int       `gorm:"primary_key;column:id"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (UserModel) New(u entity.User) *UserModel {
	return &UserModel{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
		Name:     u.Name,
	}
}

func (u *UserModel) ToUser() *entity.User {
	return &entity.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
		Name:     u.Name,
	}
}
