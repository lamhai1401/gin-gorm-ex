package userstorage

import (
	"context"

	usermodels "github.com/lamhai1401/gin-gorm-ex/user/model"
)

func (s *userStorage) DeleteUser(
	ctx context.Context,
	condition map[string]interface{},
) error {
	tx := s.db.Begin()
	err := tx.Table(usermodels.User{}.TableName()).Where(condition).Delete(nil).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
