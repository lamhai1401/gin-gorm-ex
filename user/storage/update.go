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
	if err := s.db.Where(condition).Updates(dataUpdate).Error; err != nil {
		return err
	}

	return nil
}
