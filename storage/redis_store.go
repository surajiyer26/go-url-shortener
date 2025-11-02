package storage

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisStore(addr string) *RedisStore {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &RedisStore{
		client: rdb,
		ctx:    context.Background(),
	}
}

func (s *RedisStore) Save(id, url string) {
	s.client.Set(s.ctx, id, url, 0)
}

func (s *RedisStore) Get(id string) (string, bool) {
	val, err := s.client.Get(s.ctx, id).Result()
	if err == redis.Nil {
		return "", false
	} else if err != nil {
		panic(err)
	}
	return val, true
}
