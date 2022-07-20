package usermodels

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    string `json:"id" gorm:"column:id;" validate:"required"`
	Name  string `json:"name" gorm:"column:name;" validate:"required"`
	Email string `json:"email" gorm:"column:email;" validate:"required,email"`
}

func (User) TableName() string { return "user" }
