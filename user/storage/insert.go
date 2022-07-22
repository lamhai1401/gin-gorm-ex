package userstorage

import (
	"context"

	usermodels "github.com/lamhai1401/gin-gorm-ex/user/model"
)

func (s *userStorage) CreateUser(ctx context.Context, data *usermodels.User) error {
	tx := s.db.Begin()

	err := tx.Create(data).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
