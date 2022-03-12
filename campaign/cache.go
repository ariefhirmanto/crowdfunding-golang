package campaign

import (
	"github.com/go-redis/redis/v8"
)

type cacheRepository struct {
	redisClient *redis.Client
}

func NewCacheRepository(redisClient *redis.Client) *cacheRepository {
	return &cacheRepository{redisClient: redisClient}
}
