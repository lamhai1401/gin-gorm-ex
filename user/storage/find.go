package userstorage

import (
	"context"

	usermodels "github.com/lamhai1401/gin-gorm-ex/user/model"
	"gorm.io/gorm"
)

func (s *userStorage) FindUser(
	ctx context.Context,
	condition map[string]interface{},
) (*usermodels.User, error) {
	var itemData usermodels.User

	if err := s.db.Where(condition).First(&itemData).Error; err != nil {
		if err == gorm.ErrRecordNotFound { // data not found
			return nil, usermodels.ErrItemNotFound
		}

		return nil, err // other errors
	}

	return &itemData, nil
}
