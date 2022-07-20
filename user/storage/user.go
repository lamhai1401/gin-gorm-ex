package userstorage

import (
	"context"

	"gorm.io/gorm"

	"github.com/go-playground/validator/v10"
	usermodels "github.com/lamhai1401/gin-gorm-ex/user/model"
)

type createUserStorage struct {
	db       *gorm.DB
	validate *validator.Validate
}

func NewMySQLStorage(db *gorm.DB, validate *validator.Validate) *createUserStorage {
	return &createUserStorage{db: db, validate: validate}
}

func (s *createUserStorage) CreateUser(ctx context.Context, data *usermodels.User) error {
	err := s.db.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *createUserStorage) Validate(data interface{}) error {
	// returns nil or ValidationErrors ( []FieldError )
	return s.validate.Struct(data)
}
