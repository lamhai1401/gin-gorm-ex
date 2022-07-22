package userstorage

import (
	"context"

	usermodels "github.com/lamhai1401/gin-gorm-ex/user/model"
)

func (s *userStorage) ListUsers(
	ctx context.Context,
	condition map[string]interface{},
	paging *usermodels.DataPaging,
) ([]usermodels.User, error) {
	offset := (paging.Page - 1) * paging.Limit

	var result []usermodels.User

	err := s.db.
		Table(usermodels.User{}.TableName()).
		Where(condition).
		Count(&paging.Total).
		Offset(offset).
		Order("id desc").
		Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}
