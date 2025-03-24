package config

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     Conf.Redis.Addr,
		Password: Conf.Redis.Password,
		DB:       Conf.Redis.DB,
	})

	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		panic("Redis连接失败: " + err.Error())
	}
}
