package userbiz

import (
	"context"
	"fmt"

	usermodels "github.com/lamhai1401/gin-gorm-ex/user/model"
	"github.com/lamhai1401/gin-gorm-ex/utils"
)

type UserStorage interface {
	FindUserStorage
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

	user, _ := biz.store.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if user != nil && user.Email == data.Email {
		return fmt.Errorf("%s user already exist", data.Name)
	}

	// add snowflake ID
	data.ID = utils.GetSnowflakeID()

	// add data
	err = biz.store.CreateUser(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
