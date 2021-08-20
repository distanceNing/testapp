package logic

import (
	"github.com/distanceNing/testapp/common"
	"github.com/distanceNing/testapp/repo"
	"time"
)

const (
	commentReplyCollection   = "comment_reply"
	articleCommentCollection = "comment_info"
)

type CommentManger struct {
	store *repo.MongoDbInstance
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

func (mgr *CommentManger) CreateComment(req *CreateCommentReq, rsp *common.Rsp) common.ErrorCode {
	status := common.NewSuccCode()

	return status
}

func (mgr *CommentManger) CreateCommentReply(req *CommentReplyReq, rsp *common.Rsp) common.ErrorCode {
	status := common.NewSuccCode()
	return status
}

func (mgr *CommentManger) GetCommentReply(req *GetCommentReplyReq, rsp *common.Rsp) common.ErrorCode {
	status := common.NewSuccCode()
	return status
}
