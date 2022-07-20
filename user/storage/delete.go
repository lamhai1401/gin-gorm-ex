package userstorage

import (
	"context"

	usermodels "github.com/lamhai1401/gin-gorm-ex/user/model"
)

func (s *userStorage) DeleteUser(
	ctx context.Context,
	condition map[string]interface{},
) error {

	if err := s.db.
		Table(usermodels.User{}.TableName()).
		Where(condition).Delete(nil).Error; err != nil {
		return err
	}

	return nil
}
