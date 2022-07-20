package userstorage

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type userStorage struct {
	db       *gorm.DB
	validate *validator.Validate
}

func NewMySQLStorage(db *gorm.DB) *userStorage {
	return &userStorage{db: db}
}

func (u *userStorage) SetValidate(validate *validator.Validate) {
	u.validate = validate
}

func (s *userStorage) Validate(data interface{}) error {
	// returns nil or ValidationErrors ( []FieldError )
	return s.validate.Struct(data)
}
