package userbiz

import (
	"context"

	usermodels "github.com/lamhai1401/gin-gorm-ex/user/model"
)

type ListUserStorage interface {
	ListUsers(
		ctx context.Context,
		condition map[string]interface{},
		paging *usermodels.DataPaging,
	) ([]usermodels.User, error)
}

type listBiz struct {
	store ListUserStorage
}

func NewListUserBiz(store ListUserStorage) *listBiz {
	return &listBiz{store: store}
}

func (biz *listBiz) ListUsers(ctx context.Context,
	condition map[string]interface{},
	paging *usermodels.DataPaging,
) ([]usermodels.User, error) {
	result, err := biz.store.ListUsers(ctx, condition, paging)

	if err != nil {
		return nil, err
	}

	return result, err
}
