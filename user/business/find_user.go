package userbiz

import (
	"context"

	usermodels "github.com/lamhai1401/gin-gorm-ex/user/model"
)

type FindUserStorage interface {
	FindUser(
		ctx context.Context,
		condition map[string]interface{},
	) (*usermodels.User, error)
}

type findBiz struct {
	store FindUserStorage
}

func NewFindUserBiz(store FindUserStorage) *findBiz {
	return &findBiz{store: store}
}

func (biz *findBiz) FindUser(ctx context.Context, condition map[string]interface{}) (*usermodels.User, error) {
	itemData, err := biz.store.FindUser(ctx, condition)

	if err != nil {
		return nil, err
	}

	return itemData, nil
}
