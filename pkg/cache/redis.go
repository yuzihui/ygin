package cache

import (
	"ecloudsystem/configs"
	"github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
)

var Client *Repo

type Repo struct {
	Redis *redis.Client
}

func InitCache()  {
	var cacheRepo *redis.Client
	cacheRepo , err := redisConnect(RedisConfig())
	if err != nil {
		panic("redis fail")
	}

	Client = &Repo{
		Redis: cacheRepo,
	}
}


func  RedisConfig() *redis.Options {
	cfg := configs.Get().Redis
	redisConfig := redis.Options{
		Addr:         cfg.Cache.Addr,
		Password:     cfg.Cache.Pass,
		DB:           cfg.Cache.Db,
		MaxRetries:   cfg.Cache.MaxRetries,
		PoolSize:     cfg.Cache.PoolSize,
		MinIdleConns: cfg.Cache.MinIdleConns,
	}
	return &redisConfig
}

func (d *Repo) RedisClose() error {
	err := d.Redis.Close()
	if err != nil {
		return err
	}
	return nil
}

func redisConnect(options *redis.Options) (*redis.Client, error) {
	client := redis.NewClient(options)
	if err := client.Ping().Err(); err != nil {
		return nil, errors.Wrap(err, "ping redis err")
	}
	return client, nil
}


