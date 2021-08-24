package logic

import (
	"context"
	"github.com/distanceNing/testapp/common/errcode"
	"github.com/distanceNing/testapp/common/test"
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
	return strings.ToUpper(test.RandString(16))
}

func (mgr *SessionManager) QuerySessionToken(userId string) (error, string) {
	ctx := context.Background()
	res := mgr.redisCli.RedisCli.Get(ctx, userId)
	if res.Err() != nil {
		return errcode.NewErrorCode(errcode.ErrSystem, "query session failed"), ""
	}
	t, err := res.Result()
	if err != nil {
		return errcode.NewErrorCode(errcode.ErrSystem, "get redis op return val failed"), ""
	}
	return nil, t
}

func (mgr *SessionManager) CreateSession(userId string) (error, string) {
	ctx := context.Background()
	token := mgr.genLoginToken()
	s := `
			local val = redis.call('GET', KEYS[1])
			if val == nil or val == false then                  
				redis.call('SETEX', KEYS[1], ARGV[2] , ARGV[1]) 
				return ARGV[1]                                
			else                                               
				return val                                      
			end                                                 
		`
	res := mgr.redisCli.RedisCli.Eval(ctx, s, []string{userId}, token, TokenTimeOut)
	if res.Err() != nil {
		return errcode.NewErrorCode(errcode.ErrSystem, res.Err().Error()), ""
	}
	t, err := res.Result()
	if err != nil {
		return errcode.NewErrorCode(errcode.ErrSystem, "get redis op return val failed"), ""
	}
	return nil, t.(string)
}
