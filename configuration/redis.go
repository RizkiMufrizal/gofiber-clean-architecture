package configuration

import (
	"context"
	"encoding/json"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"github.com/go-redis/redis/v9"
	"strconv"
)

func NewRedis(config Config) *redis.Client {
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
	return redisStore
}

func SetCache[T any](cacheManager *redis.Client, ctx context.Context, prefix string, key string, executeData func(context.Context, string) (T, error)) *T {
	var data []byte
	var object T
	if err := cacheManager.Get(ctx, prefix+"_"+key).Scan(&data); err == nil {
		err := json.Unmarshal(data, &object)
		exception.PanicLogging(err)

		return &object
	}
	value, err := executeData(ctx, key)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	cacheValue, err := json.Marshal(value)
	exception.PanicLogging(err)

	if err := cacheManager.Set(ctx, prefix+"_"+key, cacheValue, -1).Err(); err != nil {
		exception.PanicLogging(err)
	}
	return &value
}
