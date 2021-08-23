package logic

import (
	"github.com/distanceNing/testapp/common/types"
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

func (svc *Service) Login(req *LoginRequest, rsp *types.Rsp) error {
	return svc.loginSvc.Login(req, rsp)
}

func (svc *Service) Register(req *LoginRequest, rsp *types.Rsp) error {
	return svc.loginSvc.Register(req, rsp)
}

func (svc *Service) GetArticle(req *GetArticleReq, rsp *types.Rsp) error {
	return svc.articleMgr.GetArticle(req, rsp)
}

func (svc *Service) CreateArticle(req *CreateArticleReq, rsp *types.Rsp) error {
	return svc.articleMgr.CreateArticle(req, rsp)
}

func (svc *Service) GetChannels(req *CreateArticleReq, rsp *types.Rsp) error {
	return svc.articleMgr.GetChannels(req, rsp)
}

func (svc *Service) SearchArticle(req *GetArticlePageReq, rsp *types.Rsp) error {
	return svc.articleMgr.SearchArticle(req, rsp)
}

func (svc *Service) Upload(req *UploadImageReq, rsp *types.Rsp) error {
	return svc.imageMgr.Upload(req, rsp)
}

func (svc *Service) DeleteArticle(req *DeleteArticleReq, rsp *types.Rsp) error {
	return svc.articleMgr.DeleteArticle(req, rsp)
}

func (svc *Service) UpdateArticle(req *UpdateArticleReq, rsp *types.Rsp) error {
	return svc.articleMgr.UpdateArticle(req, rsp)
}
func (svc *Service) GetImages(req *GetImageReq, rsp *types.Rsp) error {
	return svc.imageMgr.GetImages(req, rsp)
}
