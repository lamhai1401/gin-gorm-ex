package userbiz

import (
	"context"
	"fmt"

	"github.com/lamhai1401/gin-gorm-ex/utils"
)

const unAuthErr = "An Authenticate"

type LoginUserStorage interface {
	FindUserStorage
}

type loginBiz struct {
	store LoginUserStorage
}

func NewLoginUserBiz(store LoginUserStorage) *loginBiz {
	return &loginBiz{store: store}
}

func (biz *loginBiz) Login(
	ctx context.Context,
	credential *utils.Credentials,
) (string, error) {
	// find client in table
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"email": credential.Email})
	if err != nil {
		return "", err
	}

	if !utils.CheckPasswordHash(credential.Password, user.Password) {
		return "", fmt.Errorf("%v of user %v", unAuthErr, credential.Email)
	}

	// create token
	token, err := utils.GenerateToken(credential.Email, credential.Password)
	if err != nil {
		return "", fmt.Errorf("%v of user %v err: %v", unAuthErr, credential.Email, err.Error())
	}

	return token, nil
}
