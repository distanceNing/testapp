package logic

import (
	"github.com/distanceNing/testapp/common"
	"github.com/distanceNing/testapp/conf"
)

// Service 提供的所有服务功能的封装
type Service struct {
	loginSvc   *LoginService
	articleMgr *ArticleManager
}

func NewService(conf *conf.ServerConf) *Service {
	return &Service{loginSvc: NewLoginService(&conf.RedisConf), articleMgr: NewArticleManager()}
}

func (svc *Service) Login(req *LoginRequest, rsp *common.Rsp) common.Status {
	return svc.loginSvc.Login(req, rsp)
}

func (svc *Service) Register(req *LoginRequest, rsp *common.Rsp) common.Status {
	return svc.loginSvc.Register(req, rsp)
}

func (svc *Service) CreateArticle(req *CreateArticleReq, rsp *common.Rsp) common.Status {
	return svc.articleMgr.CreateArticle(req, rsp)
}
