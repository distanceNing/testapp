package repo

import (
	"context"
	"testing"

	"github.com/distanceNing/testapp/src/conf"
)

func TestNewRedisInstance(t *testing.T) {
	type args struct {
		conf *conf.RedisConf
	}
	rdi := GetDefaultRedis()

	var ctx = context.Background()
	err := rdi.RedisCli.Set(ctx, "key1", "value1", 0).Err()
	if err != nil {
		panic(err)
	}

	luaScript := "return {redis.call('GET',KEYS[1])}"

	res := rdi.RedisCli.Eval(ctx, luaScript, []string{"key1"})
	if res.Err() != nil {
		panic(err)
	}

}
