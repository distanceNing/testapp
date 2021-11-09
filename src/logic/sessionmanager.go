package logic

import (
	"context"
	"github.com/distanceNing/testapp/src/common/errcode"
	"github.com/distanceNing/testapp/src/common/test"
	"github.com/distanceNing/testapp/src/conf"
	"github.com/distanceNing/testapp/src/repo"
	"strings"
)

type SessionManager struct {
	redisCli *repo.RedisInstance
}

const (
	SetSessionKeyScript = `
			local val = redis.call('GET', KEYS[1])
			if val == nil or val == false then                  
				redis.call('SETEX', KEYS[1], ARGV[2] , ARGV[1]) 
				return ARGV[1]                                
			else                                               
				return val                                      
			end                                                 
		`
)

func NewSessionManager(conf *conf.RedisConf) *SessionManager {
	return &SessionManager{redisCli: repo.NewRedisInstance(conf)}
}

func (mgr *SessionManager) genLoginToken() string {
	return strings.ToUpper(test.RandString(16))
}

func (mgr *SessionManager) QuerySessionToken(userId string) (error, string) {
	ctx := context.Background()
	res := mgr.redisCli.RedisCli.Get(ctx, userId)
	if res.Err() != nil {
		return errcode.New(errcode.ErrSystem, "query session failed"), ""
	}
	t, err := res.Result()
	if err != nil {
		return errcode.New(errcode.ErrSystem, "get redis op return val failed"), ""
	}
	return nil, t
}

func (mgr *SessionManager) CreateSession(userId string) (error, string) {
	ctx := context.Background()
	token := mgr.genLoginToken()

	res := mgr.redisCli.RedisCli.Eval(ctx, SetSessionKeyScript, []string{userId}, token, TokenTimeOut)
	if res.Err() != nil {
		return errcode.New(errcode.ErrSystem, res.Err().Error()), ""
	}
	t, err := res.Result()
	if err != nil {
		return errcode.New(errcode.ErrSystem, "get redis op return val failed"), ""
	}
	return nil, t.(string)
}
