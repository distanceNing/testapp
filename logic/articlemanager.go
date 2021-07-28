package logic

import (
	"github.com/distanceNing/testapp/common"
	"github.com/distanceNing/testapp/repo"
)

const (
	DraftStatus = 0
	CheckingStatus
	CheckedStatus
	CheckedFailedStatus
	DeletedStatus
)

type CreateArticleReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  int    `json:"status"`
}
type ArticleManager struct {
}

func NewArticleManager() *ArticleManager {
	return &ArticleManager{}
}

func (mgr *ArticleManager) CreateArticle(req *CreateArticleReq, rsp *common.Rsp) common.Status {
	obj := repo.ArticleInfo{Title: req.Title, Content: req.Content, Status: req.Status}
	status := repo.CreateObject(&obj)
	return status
}
