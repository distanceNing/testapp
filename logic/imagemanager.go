package logic

import (
	"github.com/distanceNing/testapp/common"
	"github.com/distanceNing/testapp/repo"
)

type ImageManager struct {
	idGenerator *common.Snowflake
}

func NewImageManager() *ImageManager {
	return &ImageManager{idGenerator: &common.Snowflake{}}
}

type UploadImageReq struct {
	Url      string
	BelongTo string
}
type GetImageReq struct {
	Collect   bool
	PageNum   int
	PageCount int
}

func (mgr *ImageManager) Upload(req *UploadImageReq, rsp *common.Rsp) common.Status {
	id := mgr.idGenerator.NextVal()
	s := repo.CreateObject(repo.ImageInfo{id, req.BelongTo, req.Url})
	if !s.Ok() {
		return s
	}
	rsp.Set("id", id)
	rsp.Set("url", req.Url)
	return s
}

func (mgr *ImageManager) GetImages(req *GetImageReq, rsp *common.Rsp) common.Status {
	cond := &repo.ImageInfo{}
	var totalCnt int64
	status := repo.QueryObjectCount(cond, &totalCnt)
	if !status.Ok() {
		return status
	}

	var objs []repo.ImageInfo
	repo.QueryObjectByPage(cond, &objs, req.PageCount, req.PageNum)
	if status.Code() == common.ErrRecordNotExist {
		status.Set(common.ErrRecordNotExist, "images not exist")
		return status
	} else if !status.Ok() {
		status.Set(common.ErrSystem, "query filed")
		return status
	}

	type ImageInfo struct {
		Id  int64  `json:"id"`
		Url string `json:"url"`
	}
	var rspObjs []ImageInfo
	for i := range objs {
		rspObjs = append(rspObjs, ImageInfo{objs[i].Id, objs[i].Url})
	}

	rsp.Set("total_count", totalCnt)
	rsp.Set("page", req.PageNum)
	rsp.Set("per_page_count", req.PageCount)
	rsp.Set("results", rspObjs)
	return status
}
