package logic

import (
	"github.com/distanceNing/testapp/common"
	"github.com/distanceNing/testapp/conf"
)

// Service 提供的所有服务功能的封装
type Service struct {
	loginSvc   *LoginService
	articleMgr *ArticleManager
	imageMgr   *ImageManager
}

func NewService(conf *conf.ServerConf) *Service {
	return &Service{loginSvc: NewLoginService(&conf.RedisConf), articleMgr: NewArticleManager(), imageMgr: NewImageManager()}
}

func (svc *Service) Login(req *LoginRequest, rsp *common.Rsp) common.Status {
	return svc.loginSvc.Login(req, rsp)
}

func (svc *Service) Register(req *LoginRequest, rsp *common.Rsp) common.Status {
	return svc.loginSvc.Register(req, rsp)
}

func (svc *Service) GetArticle(req *GetArticleReq, rsp *common.Rsp) common.Status {
	return svc.articleMgr.GetArticle(req, rsp)
}

func (svc *Service) CreateArticle(req *CreateArticleReq, rsp *common.Rsp) common.Status {
	return svc.articleMgr.CreateArticle(req, rsp)
}

func (svc *Service) GetChannels(req *CreateArticleReq, rsp *common.Rsp) common.Status {
	return svc.articleMgr.GetChannels(req, rsp)
}

func (svc *Service) SearchArticle(req *GetArticlePageReq, rsp *common.Rsp) common.Status {
	return svc.articleMgr.SearchArticle(req, rsp)
}

func (svc *Service) Upload(req *UploadImageReq, rsp *common.Rsp) common.Status {
	return svc.imageMgr.Upload(req, rsp)
}

func (svc *Service) DeleteArticle(req *DeleteArticleReq, rsp *common.Rsp) common.Status {
	return svc.articleMgr.DeleteArticle(req, rsp)
}

func (svc *Service) UpdateArticle(req *UpdateArticleReq, rsp *common.Rsp) common.Status {
	return svc.articleMgr.UpdateArticle(req, rsp)
}
func (svc *Service) GetImages(req *GetImageReq, rsp *common.Rsp) common.Status {
	return svc.imageMgr.GetImages(req, rsp)
}