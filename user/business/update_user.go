package userbiz

import (
	"context"

	usermodels "github.com/lamhai1401/gin-gorm-ex/user/model"
)

type UpdateUserStorage interface {
	FindUser(
		ctx context.Context,
		condition map[string]interface{},
	) (*usermodels.User, error)

	UpdateUser(
		ctx context.Context,
		condition map[string]interface{},
		dataUpdate *usermodels.User,
	) error
}

type updateBiz struct {
	store UpdateUserStorage
}

func NewUpdateUserBiz(store UpdateUserStorage) *updateBiz {
	return &updateBiz{store: store}
}

func (biz *updateBiz) UpdateUser(
	ctx context.Context,
	condition map[string]interface{},
	dataUpdate *usermodels.User,
) error {
	// step 1: Find item by conditions
	_, err := biz.store.FindUser(ctx, condition)

	if err != nil {
		return err
	}

	// just a demo in case we dont allow update Finished item
	// if oldItem.Status == "Finished" {
	// 	return usermodels.ErrCannotUpdateFinishedItem
	// }

	// Step 2: call to storage to update item
	if err := biz.store.UpdateUser(ctx, condition, dataUpdate); err != nil {
		return err
	}

	return nil
}
