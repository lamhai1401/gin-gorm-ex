package usermodels

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

const (
	Idx_Email = "idx_email"
)

var (
	ErrTitleCannotBeBlank       = errors.New("title can not be blank")
	ErrItemNotFound             = errors.New("item not found")
	ErrCannotUpdateFinishedItem = errors.New("can not update finished item")
)

type User struct {
	ID        string `json:"id" gorm:"default:id,column:id,primarykey;" mapstructure:"id"`
	Name      string `json:"name" gorm:"column:name;" validate:"required" mapstructure:"name"`
	Email     string `gorm:"column:email,size:255;index:idx_email,unique"  validate:"required,email" mapstructure:"email"`
	Password  string `json:"password" gorm:"column:password;" validate:"required" mapstructure:"password"`
	Token     string `json:"token" gorm:"column:token;" validate:"jwt" mapstructure:"token"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
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
