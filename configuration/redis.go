package configuration

import (
	"context"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"strconv"
)

func NewRedis(config Config) *cache.Cache {
	host := config.Get("REDIS_HOST")
	port := config.Get("REDIS_PORT")
	maxPoolSize, err := strconv.Atoi(config.Get("REDIS_POOL_MAX_SIZE"))
	minIdlePoolSize, err := strconv.Atoi(config.Get("REDIS_POOL_MIN_IDLE_SIZE"))
	exception.PanicLogging(err)

	redisStore := redis.NewClient(&redis.Options{
		Addr:         host + ":" + port,
		PoolSize:     maxPoolSize,
		MinIdleConns: minIdlePoolSize,
	})

	cacheManager := cache.New(&cache.Options{
		Redis: redisStore,
	})

	return cacheManager
}

func SetCache(cacheManager *cache.Cache, ctx context.Context, key string, value interface{}) {
	if err := cacheManager.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: value,
	}); err != nil {
		exception.PanicLogging(err)
	}
}

func GetCache(cacheManager *cache.Cache, ctx context.Context, key string) interface{} {
	var object interface{}
	if err := cacheManager.Get(ctx, key, &object); err == nil {
		return object
	}
	return nil
}
