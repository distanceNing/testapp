package logic

import (
	"github.com/distanceNing/testapp/src/common/errcode"
	"github.com/distanceNing/testapp/src/common/filestore"
	"github.com/distanceNing/testapp/src/common/idgenerator"
	"github.com/distanceNing/testapp/src/common/types"
	"github.com/distanceNing/testapp/src/repo"
	"io"
)

type ImageManager struct {
	idGenerator *idgenerator.Snowflake
	store       filestore.FileStore
}

func NewImageManager() *ImageManager {
	return &ImageManager{idGenerator: &idgenerator.Snowflake{}, store: filestore.MakeFileStore(filestore.CosStore)}
}

type UploadImageReq struct {
	Filename string
	R        io.Reader
	BelongTo string
}
type GetImageReq struct {
	Collect   bool
	PageNum   int
	PageCount int
}

func (mgr *ImageManager) Upload(req *UploadImageReq, rsp *types.Rsp) error {
	url, err := mgr.store.Put(req.Filename, req.R)
	if err != nil {
		return err
	}
	id := mgr.idGenerator.NextVal()
	//err = repo.CreateObject(repo.ImageInfo{id, req.BelongTo, url})
	//if err != nil {
	//	return err
	//}
	rsp.Set("id", id)
	rsp.Set("url", url)
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
		return errcode.New(errcode.ErrRecordNotExist, "imageerrnot exist")
	} else if err != nil {
		return errcode.New(errcode.ErrSystem, "query filed")
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
	return nil
}
