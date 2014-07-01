package generator

import (
	"socialapi/workers/helper"
	"socialapi/workers/sitemap/common"

	"github.com/koding/redis"
)

type FileSelector interface {
	Select() (string, error)
}

type SimpleFileSelector struct{}

func (s SimpleFileSelector) Select() (string, error) {
	return "sitemap", nil
}

type CachedFileSelector struct{}

func (s CachedFileSelector) Select() (string, error) {
	redisConn := helper.MustGetRedisConn()

	item, err := redisConn.PopSetMember(common.PrepareCurrentFileNameCacheKey())

	if err != redis.ErrNil {
		return "", err
	}

	return item, err
}
