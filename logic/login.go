package logic

import (
	"github.com/distanceNing/testapp/common"
	"github.com/distanceNing/testapp/repo"
	"strings"
	"time"
)

const TokenTimeOut = 2 * 60

type LoginRequest struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
	UserType string `json:"userType"`
}

type LoginService struct {
}

func NewLoginService() *LoginService {
	return &LoginService{}
}

func genLoginToken() string {
	return strings.ToUpper(common.RandString(16))
}

func (loginSvc *LoginService) checkSessionToken(userId string) common.Status {
	status := common.NewStatus()
	var session repo.UserSession
	status, session = repo.QuerySessionToken(userId)
	if !status.Ok() {
		return status
	}

	d := time.Since(session.CreatedAt)
	if d.Seconds() > TokenTimeOut {
		status.Set(common.ErrTokenTimeOut, "token time out")
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

	token := genLoginToken()
	status = repo.CreateSession(req.UserId, token)
	if !status.Ok() && status.Code() != common.ErrDbDupKey {
		status.Set(common.ErrSystem, "proc failed")
		return status
	}
	if status.Code() == common.ErrDbDupKey {
		status = repo.UpdateSessionToken(req.UserId, token)
		if !status.Ok() {
			return status
		}
	}
	rsp.Set("token", token)
	return status
}
