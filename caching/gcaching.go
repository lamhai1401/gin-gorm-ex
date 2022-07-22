package caching

import (
	"fmt"
	"time"

	"github.com/bluele/gcache"
)

const (
	ExpireTime = time.Duration(3600)
)

var localCaching gcache.Cache
var AddLocalCacheErr = func(reason string) error {
	return fmt.Errorf("AddLocalCacheErr err: %v", reason)
}

var GetLocalCacheErr = func(reason string) error {
	return fmt.Errorf("GetLocalCacheErr err: %v", reason)
}

func InitLocalCaching(size int) {
	localCaching = gcache.New(size).LRU().Build()
}

func AddLocalCache(key string, value interface{}) error {
	err := localCaching.SetWithExpire(key, value, ExpireTime*time.Second)
	if err != nil {
		return AddLocalCacheErr(err.Error())
	}
	return nil
}

func RemoveLocalCache(key string) bool {
	return localCaching.Remove(key)
}

func GetLocalCache(key string) (interface{}, error) {
	data, err := localCaching.Get(key)
	if err != nil {
		return nil, GetLocalCacheErr(err.Error())
	}
	return data, nil
}
