package redis

import (
	"context"
	"strconv"
	"time"

	"github.com/JackFazackerley/hackernews-microservice/internal/config"

	"github.com/JackFazackerley/hackernews-microservice/internal/cache"
	"github.com/JackFazackerley/hackernews-microservice/internal/client"
	redisCache "github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	ring  *redis.Ring
	cache *redisCache.Cache
}

func New(cfg config.Redis) (cache.Cache, error) {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server": cfg.RedisAddress(),
		},
	})

	c := redisCache.New(&redisCache.Options{
		Redis: ring,
	})

	return &Redis{
		ring:  ring,
		cache: c,
	}, nil
}

func (r Redis) Get(key string) (*client.Item, error) {
	var item *client.Item

	if err := r.cache.Get(context.Background(), key, &item); err != nil {
		return nil, err
	}

	return item, nil
}

func (r Redis) Put(item *client.Item) error {
	return r.cache.Set(&redisCache.Item{
		Ctx:   context.Background(),
		Key:   strconv.Itoa(item.ID),
		Value: item,
		TTL:   time.Minute * 2,
	})
}

func (r Redis) Close() error {
	return r.ring.Close()
}
