package logic

import (
	"github.com/distanceNing/testapp/common/errcode"
	"github.com/distanceNing/testapp/common/types"
	"github.com/distanceNing/testapp/repo"
	"strconv"
	"strings"
	"time"
)

const (
	DraftStatus = 1
	CheckingStatus
	CheckedStatus
	CheckedFailedStatus
	DeletedStatus
)

type CoverInfo struct {
	Type   int      `json:"type"`
	Images []string `json:"images"`
}

type CreateArticleReq struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	ChannelId int       `json:"channel_id"`
	Status    int       `json:"status"`
	Cover     CoverInfo `json:"cover"`
}

type GetArticlePageReq struct {
	Status    int       `json:"status"`
	ChannelId int       `json:"channel_id"`
	Begin     time.Time `json:"begin_publish_date"`
	End       time.Time `json:"end_publish_date"`
	PageNum   int       `json:"page_num"`
	PageCount int       `json:"per_page_count"`
}

type SearchArticleReq struct {
	KeyWord string `json:"title"`
}

type GetArticleReq struct {
	Id string
}

type DeleteArticleReq struct {
	Id string
}

type UpdateArticleReq struct {
	Id        string
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	ChannelId int       `json:"channel_id"`
	Status    int       `json:"status"`
	Cover     CoverInfo `json:"cover"`
}

type ChannelInfo struct {
	Id   int    `json:"idgenerator"`
	Name string `json:"name"`
}

type ArticleManager struct {
}

func NewArticleManager() *ArticleManager {
	return &ArticleManager{}
}

func (mgr *ArticleManager) SearchArticle(req *GetArticlePageReq, rsp *types.Rsp) error {
	cond := &repo.ArticleInfo{}
	if req.ChannelId != 0 {
		cond.ChannelId = req.ChannelId
	}
	if req.Status != 0 {
		cond.Status = req.Status
	}
	var totalCnt int64
	err := repo.QueryObjectCount(cond, &totalCnt)
	if err != nil {
		return err
	}

	var objs []repo.ArticleInfo
	err = repo.QueryObjectByPage(cond, &objs, req.PageCount, req.PageNum)
	if errcode.Code(err) == errcode.ErrRecordNotExist {
		return errcode.NewErrorCode(errcode.ErrRecordNotExist, "article not exist")
	} else if err != nil {
		return errcode.NewErrorCode(errcode.ErrSystem, "query filed")
	}

	rsp.Set("total_count", totalCnt)
	rsp.Set("page", req.PageNum)
	rsp.Set("per_page", req.PageCount)
	rsp.Set("results", objs)
	return nil
}

func (mgr *ArticleManager) GetArticle(req *GetArticleReq, rsp *types.Rsp) error {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return errcode.NewErrorCode(errcode.ErrRequest, "idgenerator to int failed")
	}
	obj := repo.ArticleInfo{}
	err = repo.QueryObject(&repo.ArticleInfo{Id: id}, &obj)
	if errcode.Code(err) == errcode.ErrRecordNotExist {
		return errcode.NewErrorCode(errcode.ErrRecordNotExist, "article not exist")
	} else if err != nil {
		return errcode.NewErrorCode(errcode.ErrSystem, "query filed")
	}

	type ArticleInfo struct {
		Id        int       `json:"idgenerator"`
		ChannelId int       `json:"channel_id"`
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		Images    []string  `json:"images"`
		Status    int       `json:"status"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	images := strings.Split(obj.Images, "|")
	rsp.Set("article_info", ArticleInfo{obj.Id, obj.ChannelId, obj.Title, obj.Content, images,
		obj.Status, obj.CreatedAt, obj.UpdatedAt})
	return nil
}

func (mgr *ArticleManager) CreateArticle(req *CreateArticleReq, rsp *types.Rsp) error {
	if req.Title == "" || req.Content == "" {
		return errcode.NewErrorCode(errcode.ErrRequest, "title or content is empty ")
	}

	images := ""
	for i := range req.Cover.Images {
		images = images + req.Cover.Images[i] + "|"
	}
	obj := repo.ArticleInfo{Title: req.Title, ChannelId: req.ChannelId, Content: req.Content, Images: images, Status: req.Status}
	err := repo.CreateObject(&obj)
	return err
}

func (mgr *ArticleManager) DeleteArticle(req *DeleteArticleReq, rsp *types.Rsp) error {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return errcode.NewErrorCode(errcode.ErrRequest, "idgenerator to int failed")
	}
	err = repo.DeleteObject(&repo.ArticleInfo{Id: id})
	return err
}

func (mgr *ArticleManager) UpdateArticle(req *UpdateArticleReq, rsp *types.Rsp) error {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return errcode.NewErrorCode(errcode.ErrRequest, "idgenerator to int failed")
	}
	updateField := &repo.ArticleInfo{}
	if req.Status != 0 {
		updateField.Status = req.Status
	}
	if req.Content != "" {
		updateField.Content = req.Content
	}
	if req.ChannelId != 0 {
		updateField.ChannelId = req.ChannelId
	}

	err = repo.UpdateObject(&repo.ArticleInfo{Id: id}, updateField)
	return err
}

func (mgr *ArticleManager) GetChannels(req *CreateArticleReq, rsp *types.Rsp) error {
	rsp.Set("channels", []ChannelInfo{{1, "团队活动"}, {2, "科研获奖"}, {3, "教学获奖"}})
	return nil
}
