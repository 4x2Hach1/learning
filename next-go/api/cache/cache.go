package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	cache *redis.Client
}

func NewCache(cache *redis.Client) *Cache {
	return &Cache{cache: cache}
}

func (c *Cache) SetCache(key CacheKey, value CacheValue, exSeconds int64) error {
	_, err := c.cache.Set(context.Background(),
		key.Key(), value.Value(), time.Duration(exSeconds)*time.Second).Result()
	if err != nil && err != redis.Nil {
		return err
	}

	return nil
}

func (c *Cache) GetCache(key CacheKey, value CacheValue) error {
	res, err := c.cache.Get(context.Background(), key.Key()).Result()
	if err != nil && err != redis.Nil {
		return err
	}

	value.FromStr(res)
	return nil
}

func (c *Cache) DeleteCache(key CacheKey) error {
	_, err := c.cache.Del(context.Background(), key.Key()).Result()
	if err != nil && err != redis.Nil {
		return err
	}

	return nil
}
