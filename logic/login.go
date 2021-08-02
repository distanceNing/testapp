package logic

import (
	"github.com/distanceNing/testapp/common"
	"github.com/distanceNing/testapp/conf"
	"github.com/distanceNing/testapp/repo"
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
func (loginSvc *LoginService) Register(req *LoginRequest, rsp *common.Rsp) common.Status {
	status, _ := repo.QueryUserInfo(req.UserId)
	if status.Ok() {
		status.Set(common.ErrUserAlreadyExist, "user already exist")
		return status
	} else if status.Code() != common.ErrUserNotExist {
		return status
	}
	status = repo.CreateObject(&repo.UserInfo{UserId: req.UserId, NickName: req.NickName, UserType: UserTypeMap[req.UserType],
		UserPassword: req.Password, Email: req.Email, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	if !status.Ok() {
		return status
	}
	return status
}

// Login
func (loginSvc *LoginService) Login(req *LoginRequest, rsp *common.Rsp) common.Status {
	status, userInfo := repo.QueryUserInfo(req.UserId)
	if !status.Ok() {
		return status
	}
	if userInfo.UserPassword != req.Password {
		status.Set(common.ErrPasswordNotMatch, "user password not match")
		return status
	}
	var token string
	status, token = loginSvc.sessionMgr.CreateSession(req.UserId)
	if !status.Ok() {
		return status
	}
	rsp.Set("token", token)
	return status
}

func (loginSvc *LoginService) CheckSessionToken(userId string, token string) common.Status {
	status, tokenInSvr := loginSvc.sessionMgr.QuerySessionToken(userId)
	if !status.Ok() {
		return status
	}
	if token != tokenInSvr {
		status.Set(common.ErrTokenNotMatch, "req token not match in svr")
	}
	return status
}
