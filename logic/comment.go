package logic

import (
	"github.com/distanceNing/testapp/common"
	"time"
)

type CommentManger struct {
}

type CreateCommentReq struct {
	PublisherId string `json:"publisherId"`
	Content     string `json:"content"`
}

type CommentReplyReq struct {
	CommentId   int    `json:"commentId"`
	PublisherId string `json:"publisherId"`
	Content     string `json:"content"`
}

type GetCommentReplyReq struct {
	CommentId int    `json:"commentId"`
	PageNum   int    `json:"pageNum"`
	Count     string `json:"count"`
}

type CommentInfo struct {
	Id          int64  `json:"id"`
	ArticleId   int64  `gorm:"index"` // 属于那个作品
	PublisherId string `gorm:"index"`
	Content     string
	Status      int
	CreatedAt   time.Time
}

type CommentReply struct {
	ParentId    int    `gorm:"primaryKey"` // 父评论id
	PublisherId string `gorm:"index"`
	Content     string
	Status      int
	CreatedAt   time.Time
}

type ArticleComment struct {
}

func (mgr *CommentManger) CreateComment(req *CreateCommentReq, rsp *common.Rsp) common.Status {
	status := common.NewStatus()
	return status
}

func (mgr *CommentManger) CreateCommentReply(req *CommentReplyReq, rsp *common.Rsp) common.Status {
	status := common.NewStatus()
	return status
}

func (mgr *CommentManger) GetCommentReply(req *GetCommentReplyReq, rsp *common.Rsp) common.Status {
	status := common.NewStatus()
	return status
}
