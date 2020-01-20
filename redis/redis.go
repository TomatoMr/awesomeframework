package redis

import (
	"github.com/TomatoMr/awesomeframework/config"
	"github.com/TomatoMr/awesomeframework/logger"
	"github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var client *redis.Client

func InitRedis() error {
	conf := config.GetConfig()
	client = redis.NewClient(&redis.Options{
		Addr:     conf.RedisConfig.Addr,
		Password: conf.RedisConfig.Password,
		DB:       conf.RedisConfig.DB,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		err = errors.Wrap(err, "InitRedis")
		return err
	}
	logger.GetLogger().Info("Redis ping:", zap.String("ping", pong))
	return nil
}
