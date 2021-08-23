package logic

import (
	"github.com/distanceNing/testapp/common/errcode"
	"github.com/distanceNing/testapp/common/idgenerator"
	"github.com/distanceNing/testapp/common/types"
	"github.com/distanceNing/testapp/repo"
)

type ImageManager struct {
	idGenerator *idgenerator.Snowflake
}

func NewImageManager() *ImageManager {
	return &ImageManager{idGenerator: &idgenerator.Snowflake{}}
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

func (mgr *ImageManager) Upload(req *UploadImageReq, rsp *types.Rsp) error {
	id := mgr.idGenerator.NextVal()
	err := repo.CreateObject(repo.ImageInfo{id, req.BelongTo, req.Url})
	if err != nil {
		return err
	}
	rsp.Set("idgenerator", id)
	rsp.Set("url", req.Url)
	return nil
}

func (mgr *ImageManager) GetImages(req *GetImageReq, rsp *types.Rsp) error {
	cond := &repo.ImageInfo{}
	var totalCnt int64
	err := repo.QueryObjectCount(cond, &totalCnt)
	if err != nil {
		return err
	}

	var objs []repo.ImageInfo
	err = repo.QueryObjectByPage(cond, &objs, req.PageCount, req.PageNum)
	if errcode.Code(err) == errcode.ErrRecordNotExist {
		return errcode.NewErrorCode(errcode.ErrRecordNotExist, "imageerrnot exist")
	} else if err != nil {
		return errcode.NewErrorCode(errcode.ErrSystem, "query filed")
	}

	type ImageInfo struct {
		Id  int64  `json:"idgenerator"`
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
	return nil
}
