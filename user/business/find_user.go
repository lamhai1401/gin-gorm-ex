package userbiz

import (
	"context"

	"github.com/lamhai1401/gin-gorm-ex/caching"
	usermodels "github.com/lamhai1401/gin-gorm-ex/user/model"
	"github.com/lamhai1401/gin-gorm-ex/utils"
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
	// check cache here
	var err error
	var itemData *usermodels.User
	id, ok := condition["id"].(string)
	if ok {
		data, err := caching.GetLocalCache(id)
		if err == nil {
			// parsing data
			err = utils.Decode(data, &itemData)
			if err != nil {
				return nil, err
			}
			return itemData, nil
		}
	}

	itemData, err = biz.store.FindUser(ctx, condition)

	if err != nil {
		return nil, err
	}

	caching.AddLocalCache(itemData.ID, itemData)
	return itemData, nil
}
