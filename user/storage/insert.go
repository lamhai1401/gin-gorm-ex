package userstorage

import (
	"context"

	usermodels "github.com/lamhai1401/gin-gorm-ex/user/model"
)

func (s *userStorage) CreateUser(ctx context.Context, data *usermodels.User) error {
	err := s.db.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}
