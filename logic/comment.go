package logic

import "github.com/distanceNing/testapp/common"

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
