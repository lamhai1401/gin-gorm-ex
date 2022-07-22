package userstorage

import (
	"context"

	usermodels "github.com/lamhai1401/gin-gorm-ex/user/model"
)

func (s *userStorage) UpdateUser(
	ctx context.Context,
	condition map[string]interface{},
	dataUpdate *usermodels.User,
) error {
	tx := s.db.Begin()
	err := tx.Where(condition).Updates(dataUpdate).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
