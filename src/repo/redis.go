package repo

import (
	"github.com/distanceNing/testapp/src/conf"
	"github.com/go-redis/redis/v8"
)

type RedisInstance struct {
	RedisCli *redis.Client
}

func NewRedisInstance(rConf *conf.RedisConf) *RedisInstance {
	return &RedisInstance{redis.NewClient(&redis.Options{Addr: rConf.Addr, Password: rConf.Password, DB: 0})}
}

func GetDefaultRedis() *RedisInstance {
	return NewRedisInstance(&conf.RedisConf{Addr: "127.0.0.1:6379", Password: "123"})
}
