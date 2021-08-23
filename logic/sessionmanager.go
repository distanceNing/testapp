package logic

import (
	"context"
	"github.com/distanceNing/testapp/common"
	"github.com/distanceNing/testapp/conf"
	"github.com/distanceNing/testapp/repo"
	"strings"
)

type SessionManager struct {
	redisCli *repo.RedisInstance
}

func NewSessionManager(conf *conf.RedisConf) *SessionManager {
	return &SessionManager{redisCli: repo.NewRedisInstance(conf)}
}

func (mgr *SessionManager) genLoginToken() string {
	return strings.ToUpper(common.RandString(16))
}

func (mgr *SessionManager) QuerySessionToken(userId string) (error, string) {
	status := common.NewSuccCode()
	ctx := context.Background()
	res := mgr.redisCli.RedisCli.Get(ctx, userId)
	if res.Err() != nil {
		return common.NewErrorCode(common.ErrSystem, "query session failed"), ""
	}
	t, err := res.Result()
	if err != nil {
		return common.NewErrorCode(common.ErrSystem, "get redis op return val failed"), ""
	}
	return status, t
}

func (mgr *SessionManager) CreateSession(userId string) (error, string) {
	ctx := context.Background()
	token := mgr.genLoginToken()
	status := common.NewSuccCode()
	s := "local val = redis.call('GET', KEYS[1])             \n" +
		"if val == nil or val == false then                  \n" +
		"    redis.call('SETEX', KEYS[1], ARGV[2] , ARGV[1]) \n" +
		"    return ARGV[1]                                  \n" +
		"else                                                \n" +
		"    return val                                      \n" +
		"end                                                 \n"
	res := mgr.redisCli.RedisCli.Eval(ctx, s, []string{userId}, token, TokenTimeOut)
	if res.Err() != nil {
		return common.NewErrorCode(common.ErrSystem, res.Err().Error()), ""
	}
	t, err := res.Result()
	if err != nil {
		return common.NewErrorCode(common.ErrSystem, "get redis op return val failed"), ""
	}
	return status, t.(string)
}
