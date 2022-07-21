package usermodels

import (
	"errors"

	"gorm.io/gorm"
)

const (
	IdxEmail = "idx_email"
)

var (
	ErrTitleCannotBeBlank       = errors.New("title can not be blank")
	ErrItemNotFound             = errors.New("item not found")
	ErrCannotUpdateFinishedItem = errors.New("can not update finished item")
)

type User struct {
	gorm.Model
	ID    string `json:"id" gorm:"column:id;" validate:"required"`
	Name  string `json:"name" gorm:"column:name;" validate:"required"`
	Email string `json:"email" gorm:"column:email,uniqueIndex:idx_email;" validate:"required,email"`
}

func (User) TableName() string { return "user" }

type DataPaging struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"-"`
}

func (p *DataPaging) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 {
		p.Limit = 10
	}
}
