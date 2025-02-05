package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisClient interface {
	Get(key string) (res string, err error)
	Set(key, value string) (err error)
}

type redisClient struct {
	c   *redis.Client
	ctx context.Context
}

func (rc redisClient) Get(key string) (res string, err error) {
	res, err = rc.c.Get(rc.ctx, key).Result()
	if err == redis.Nil {
		err = nil
	}
	return
}

func (rc redisClient) Set(key, value string) (err error) {
	err = rc.c.Set(rc.ctx, key, value, 0).Err()
	return
}

func NewClient(ctx context.Context) RedisClient {
	return &redisClient{
		c: redis.NewClient(&redis.Options{
			Addr:       "localhost:6379",
			DB:         0,
			MaxRetries: 5,
		}),
		ctx: ctx,
	}
}
