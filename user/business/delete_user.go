package userbiz

import (
	"context"

	"github.com/lamhai1401/gin-gorm-ex/caching"
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

func (biz *deleteBiz) DeleteUser(
	ctx context.Context,
	condition map[string]interface{},
) error {

	// step 1: Find item by conditions
	user, err := biz.store.FindUser(ctx, condition)
	if err != nil {
		return err
	}

	// Step 2: call to storage to delete item
	if err := biz.store.DeleteUser(ctx, condition); err != nil {
		return err
	}

	// Step 3: remove local cache
	caching.RemoveLocalCache(user.Email)
	return nil
}
