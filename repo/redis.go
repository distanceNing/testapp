package repo

import (
	"github.com/distanceNing/testapp/conf"
	"github.com/go-redis/redis/v8"
)

type RedisInstance struct {
	redisCli *redis.Client
}

func NewRedisInstance(conf *conf.RedisConf) *RedisInstance {
	return &RedisInstance{redis.NewClient(&redis.Options{Addr: conf.Addr, Password: "", DB: 0,})}
}

