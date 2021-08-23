package main

import (
	"encoding/json"
	"github.com/distanceNing/testapp/common"
	"github.com/distanceNing/testapp/conf"
	"github.com/distanceNing/testapp/logic"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"strconv"
)

type HttpSvr struct {
	engine *gin.Engine
	conf   *conf.ServerConf
	svc    *logic.Service
}

func NewHttpSvr(serverConf *conf.ServerConf) *HttpSvr {
	return &HttpSvr{gin.Default(), serverConf, logic.NewService(serverConf)}
}

func (svr *HttpSvr) Run(addr string) error {
	svr.initRouter()
	err := svr.engine.Run(addr)
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	return nil
}

func decodeBody(body io.Reader, v interface{}) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(v)
	if err != nil {
		log.Println(err.Error())
		return common.NewErrorCode(common.ErrJsonDecodeFail, "json decode failed")
	}
	return nil
}

func getParamToInt(c *gin.Context, key string, val *int) error {
	v := c.Request.URL.Query().Get(key)
	if v == "" {
		*val = 0
		return nil
	}
	var err error
	*val, err = strconv.Atoi(v)
	if err != nil {
		return err
	}
	return nil
}

func (svr *HttpSvr) constructResponse(c *gin.Context, rsp *common.Rsp, err error) {
	rsp.Set("ret", common.Code(err))
	rsp.Set("msg", common.Msg(err))
	c.JSON(200, rsp.GetV())
}

func (svr *HttpSvr) initRouter() {
	svr.engine.POST("/register", func(c *gin.Context) {
		rsp := common.NewRsp()
		loginReq := logic.LoginRequest{}
		err := decodeBody(c.Request.Body, &loginReq)
		if err != nil {
			err = svr.svc.Register(&loginReq, rsp)
		}
		svr.constructResponse(c, rsp, err)
	})

	svr.engine.POST("/login", func(c *gin.Context) {
		rsp := common.NewRsp()
		loginReq := logic.LoginRequest{}
		err := decodeBody(c.Request.Body, &loginReq)
		if err != nil {
			err = svr.svc.Login(&loginReq, rsp)
		}
		svr.constructResponse(c, rsp, err)
	})

	svr.engine.POST("/articles/create", func(c *gin.Context) {
		rsp := common.NewRsp()
		createArticleReq := logic.CreateArticleReq{}
		err := decodeBody(c.Request.Body, &createArticleReq)
		if err != nil {
			err = svr.svc.CreateArticle(&createArticleReq, rsp)
		}
		svr.constructResponse(c, rsp, err)
	})

	svr.engine.GET("/articles/get/:id", func(c *gin.Context) {
		rsp := common.NewRsp()
		req := logic.GetArticleReq{Id: c.Params.ByName("id")}
		err := svr.svc.GetArticle(&req, rsp)
		svr.constructResponse(c, rsp, err)
	})

	svr.engine.DELETE("/articles/:id", func(c *gin.Context) {
		rsp := common.NewRsp()
		req := logic.DeleteArticleReq{Id: c.Params.ByName("id")}
		err := svr.svc.DeleteArticle(&req, rsp)
		svr.constructResponse(c, rsp, err)
	})

	svr.engine.PUT("/articles/:id", func(c *gin.Context) {
		rsp := common.NewRsp()
		req := logic.UpdateArticleReq{Id: c.Params.ByName("id")}
		err := decodeBody(c.Request.Body, &req)
		if err != nil {
			err = svr.svc.UpdateArticle(&req, rsp)
		}
		svr.constructResponse(c, rsp, err)
	})

	svr.engine.GET("/channels", func(c *gin.Context) {
		rsp := common.NewRsp()
		createArticleReq := logic.CreateArticleReq{}
		err := svr.svc.GetChannels(&createArticleReq, rsp)
		svr.constructResponse(c, rsp, err)
	})

	svr.engine.GET("/articles/search", func(c *gin.Context) {
		rsp := common.NewRsp()
		req := logic.GetArticlePageReq{}
		err := getParamToInt(c, "channel_id", &req.ChannelId)
		if err != nil {
			return
		}
		err = getParamToInt(c, "err", &req.Status)
		if err != nil {
			return
		}
		err = getParamToInt(c, "page_num", &req.PageNum)
		if err != nil {
			return
		}
		err = getParamToInt(c, "per_page_count", &req.PageCount)
		if err != nil {
			return
		}
		err = svr.svc.SearchArticle(&req, rsp)
		svr.constructResponse(c, rsp, err)
	})

	svr.engine.POST("/upload", func(c *gin.Context) {
		err := common.NewSuccCode()
		rsp := common.NewRsp()
		file, err := c.FormFile("image")
		if err != nil {
			svr.constructResponse(c, rsp, err)
			return
		}
		if err := c.SaveUploadedFile(file, svr.conf.AppConf.ImagePath+"image/"+file.Filename); err == nil {
			req := &logic.UploadImageReq{Url: svr.conf.AppConf.CdnPath + "/image/" + file.Filename, BelongTo: "xxx"}
			err = svr.svc.Upload(req, rsp)
		} else {
			err = common.NewErrorCode(common.ErrRequest, "save file failed")
		}
		svr.constructResponse(c, rsp, err)
	})

	svr.engine.GET("/images", func(c *gin.Context) {
		rsp := common.NewRsp()
		req := logic.GetImageReq{}
		err := getParamToInt(c, "page_num", &req.PageNum)
		if err != nil {
			return
		}
		err = getParamToInt(c, "per_page_count", &req.PageCount)
		if err != nil {
			return
		}
		req.Collect = c.Request.URL.Query().Get("collect") == "true"
		err = svr.svc.GetImages(&req, rsp)
		svr.constructResponse(c, rsp, err)
	})

	svr.engine.GET("/image/:file", func(c *gin.Context) {
		fileName := svr.conf.AppConf.ImagePath + c.Request.URL.Path
		c.File(fileName)
	})

}
