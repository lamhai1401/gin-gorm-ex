package userstorage

import (
	"context"

	usermodels "github.com/lamhai1401/gin-gorm-ex/user/model"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
	"gorm.io/hints"
)

func (s *userStorage) FindUser(
	ctx context.Context,
	condition map[string]interface{},
) (*usermodels.User, error) {
	var userData usermodels.User
	var userCondition usermodels.User

	err := mapstructure.Decode(condition, &userCondition)
	if err != nil {
		return nil, err
	}

	err = s.db.
		Clauses(hints.UseIndex(usermodels.Idx_Email)).
		Find(&userData, userCondition).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { // data not found
			return nil, usermodels.ErrItemNotFound
		}
		return nil, err // other errors
	}

	return &userData, nil
}
