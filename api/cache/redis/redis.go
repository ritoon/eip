package redis

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	client *redis.Client
}

func New(addr string, password string, db int) *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	return &Redis{client: rdb}
}

func (r *Redis) Set(c context.Context, key string, value interface{}, duration time.Duration) error {
	err := r.client.Set(c, key, value, duration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Get(c context.Context, key string) ([]byte, error) {
	val, err := r.client.Get(c, key).Bytes()
	if err != nil {
		log.Println("cache: Get key: not found:", key)
		return nil, err
	}
	log.Println("cache: Get key: found:", key)

	return val, nil
}
