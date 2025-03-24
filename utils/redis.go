package utils

import (
	"gin-learn-notes/config"
	"time"
)

func RedisSet(key string, value interface{}, expiration time.Duration) error {
	return config.RedisClient.Set(config.Ctx, key, value, expiration).Err()
}

func RedisGet(key string) (string, error) {
	return config.RedisClient.Get(config.Ctx, key).Result()
}

func RedisDel(key string) error {
	return config.RedisClient.Del(config.Ctx, key).Err()
}
