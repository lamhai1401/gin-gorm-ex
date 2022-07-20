package userbiz

import (
	"context"

	usermodels "github.com/lamhai1401/gin-gorm-ex/user/model"
)

type DeleteUserStorage interface {
	FindUser(
		ctx context.Context,
		condition map[string]interface{},
	) (*usermodels.User, error)

	DeleteUser(
		ctx context.Context,
		condition map[string]interface{},
	) error
}

type deleteBiz struct {
	store DeleteUserStorage
}

func NewDeleteUserBiz(store DeleteUserStorage) *deleteBiz {
	return &deleteBiz{store: store}
}

func (biz *deleteBiz) DeleteItem(
	ctx context.Context,
	condition map[string]interface{},
) error {

	// step 1: Find item by conditions
	_, err := biz.store.FindUser(ctx, condition)
	if err != nil {
		return err
	}

	// Step 2: call to storage to delete item
	if err := biz.store.DeleteUser(ctx, condition); err != nil {
		return err
	}

	return nil
}
