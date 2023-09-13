package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisKV struct {
	client *redis.Client
}

func NewRedisKV(opts *redis.Options) *RedisKV {
	return &RedisKV{
		client: redis.NewClient(opts),
	}
}

func (rkv *RedisKV) Get(ctx context.Context, key string) (string, error) {
	return rkv.client.Get(ctx, key).Result()
}

func (rkv *RedisKV) Set(ctx context.Context, key string, val string) error {
	return rkv.client.Set(ctx, key, val, 0).Err()
}

func (rkv *RedisKV) Exists(ctx context.Context, key string) (bool, error) {
	res, err := rkv.client.Exists(ctx, key).Result()
	if res == 0 {
		return false, err
	}
	return true, err
}

var _ KeyValuer = (*RedisKV)(nil)
