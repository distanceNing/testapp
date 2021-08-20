package logic

import (
	"github.com/distanceNing/testapp/common"
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
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ArticleManager struct {
}

func NewArticleManager() *ArticleManager {
	return &ArticleManager{}
}

func (mgr *ArticleManager) SearchArticle(req *GetArticlePageReq, rsp *common.Rsp) common.ErrorCode {
	cond := &repo.ArticleInfo{}
	if req.ChannelId != 0 {
		cond.ChannelId = req.ChannelId
	}
	if req.Status != 0 {
		cond.Status = req.Status
	}
	var totalCnt int64
	status := repo.QueryObjectCount(cond, &totalCnt)
	if !status.Ok() {
		return status
	}

	var objs []repo.ArticleInfo
	repo.QueryObjectByPage(cond, &objs, req.PageCount, req.PageNum)
	if status.Code() == common.ErrRecordNotExist {
		status.Set(common.ErrRecordNotExist, "article not exist")
		return status
	} else if !status.Ok() {
		status.Set(common.ErrSystem, "query filed")
		return status
	}

	rsp.Set("total_count", totalCnt)
	rsp.Set("page", req.PageNum)
	rsp.Set("per_page", req.PageCount)
	rsp.Set("results", objs)
	return status
}

func (mgr *ArticleManager) GetArticle(req *GetArticleReq, rsp *common.Rsp) common.ErrorCode {
	status := common.NewSuccCode()
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		status.Set(common.ErrRequest, "id to int failed")
		return status
	}
	obj := repo.ArticleInfo{}
	status = repo.QueryObject(&repo.ArticleInfo{Id: id}, &obj)
	if status.Code() == common.ErrRecordNotExist {
		status.Set(common.ErrRecordNotExist, "article not exist")
		return status
	} else if !status.Ok() {
		status.Set(common.ErrSystem, "query filed")
		return status
	}

	type ArticleInfo struct {
		Id        int       `json:"id"`
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
	return status
}

func (mgr *ArticleManager) CreateArticle(req *CreateArticleReq, rsp *common.Rsp) common.ErrorCode {
	status := common.NewSuccCode()
	if req.Title == "" || req.Content == "" {
		status.Set(common.ErrRequest, "title or content is empty ")
		return status
	}

	images := ""
	for i := range req.Cover.Images {
		images = images + req.Cover.Images[i] + "|"
	}
	obj := repo.ArticleInfo{Title: req.Title, ChannelId: req.ChannelId, Content: req.Content, Images: images, Status: req.Status}
	status = repo.CreateObject(&obj)
	return status
}

func (mgr *ArticleManager) DeleteArticle(req *DeleteArticleReq, rsp *common.Rsp) common.ErrorCode {
	status := common.NewSuccCode()
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		status.Set(common.ErrRequest, "id to int failed")
		return status
	}
	repo.DeleteObject(&repo.ArticleInfo{Id: id})
	return status
}

func (mgr *ArticleManager) UpdateArticle(req *UpdateArticleReq, rsp *common.Rsp) common.ErrorCode {
	status := common.NewSuccCode()
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		status.Set(common.ErrRequest, "id to int failed")
		return status
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

	repo.UpdateObject(&repo.ArticleInfo{Id: id}, updateField)
	return status
}

func (mgr *ArticleManager) GetChannels(req *CreateArticleReq, rsp *common.Rsp) common.ErrorCode {
	rsp.Set("channels", []ChannelInfo{{1, "团队活动"}, {2, "科研获奖"}, {3, "教学获奖"}})
	return common.NewSuccCode()
}
