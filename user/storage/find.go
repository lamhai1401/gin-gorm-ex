package userstorage

import (
	"context"

	usermodels "github.com/lamhai1401/gin-gorm-ex/user/model"
	"gorm.io/gorm"
	"gorm.io/hints"
)

func (s *userStorage) FindUser(
	ctx context.Context,
	condition map[string]interface{},
) (*usermodels.User, error) {
	var userData usermodels.User

	// if err := s.db.Where(condition).First(&itemData).Error; err != nil {
	// 	if err == gorm.ErrRecordNotFound { // data not found
	// 		return nil, usermodels.ErrItemNotFound
	// 	}

	// 	return nil, err // other errors
	// }

	err := s.db.Clauses(hints.UseIndex(usermodels.IdxEmail)).Find(&userData, condition).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { // data not found
			return nil, usermodels.ErrItemNotFound
		}
		return nil, err // other errors
	}

	return &userData, nil
}
