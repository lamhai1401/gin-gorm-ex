package userbiz

import (
	"context"

	usermodels "github.com/lamhai1401/gin-gorm-ex/user/model"
)

type UserStorage interface {
	CreateUser(ctx context.Context, data *usermodels.User) error
	Validate(data interface{}) error
}

type createUserBiz struct {
	store UserStorage
}

func NewCreateUserBiz(store UserStorage) *createUserBiz {
	return &createUserBiz{store: store}
}

func (biz *createUserBiz) CreateNewUser(ctx context.Context, data *usermodels.User) error {
	// maybe validate check error here
	err := biz.store.Validate(data)
	if err != nil {
		return err
	}

	// add data
	err = biz.store.CreateUser(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
