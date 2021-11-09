package logic

import (
	"github.com/distanceNing/testapp/src/common/errcode"
	"github.com/distanceNing/testapp/src/common/types"
	"github.com/distanceNing/testapp/src/conf"
	"github.com/distanceNing/testapp/src/repo"
	"time"
)

const TokenTimeOut = 60 * 60

var UserTypeMap = map[string]int{"teacher": 0, "manager": 1}

type LoginRequest struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
	UserType string `json:"userType"`
	Email    string `json:"email"`
	NickName string `json:"nickName"`
}

type LoginService struct {
	sessionMgr *SessionManager
}

func NewLoginService(conf *conf.RedisConf) *LoginService {
	return &LoginService{NewSessionManager(conf)}
}

// Login
func (loginSvc *LoginService) Register(req *LoginRequest, rsp *types.Rsp) error {
	err, _ := repo.QueryUserInfo(req.UserId)
	if err != nil {
		return errcode.New(errcode.ErrUserAlreadyExist, "user already exist")
	} else if errcode.Code(err) != errcode.ErrUserNotExist {
		return err
	}
	err = repo.CreateObject(&repo.UserInfo{UserId: req.UserId, NickName: req.NickName, UserType: UserTypeMap[req.UserType],
		UserPassword: req.Password, Email: req.Email, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	if err != nil {
		return err
	}
	return err
}

// Login
func (loginSvc *LoginService) Login(req *LoginRequest, rsp *types.Rsp) error {
	err, userInfo := repo.QueryUserInfo(req.UserId)
	if err != nil {
		return err
	}
	if userInfo.UserPassword != req.Password {
		return errcode.New(errcode.ErrPasswordNotMatch, "user password not match")
	}
	var token string
	err, token = loginSvc.sessionMgr.CreateSession(req.UserId)
	if err != nil {
		return err
	}
	rsp.Set("token", token)
	return err
}

func (loginSvc *LoginService) CheckSessionToken(userId string, token string) error {
	err, tokenInSvr := loginSvc.sessionMgr.QuerySessionToken(userId)
	if err != nil {
		return err
	}
	if token != tokenInSvr {
		return errcode.New(errcode.ErrTokenNotMatch, "req token not match in svr")
	}
	return err
}
