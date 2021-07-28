package main

import (
	"encoding/json"
	"github.com/distanceNing/testapp/common"
	"github.com/distanceNing/testapp/conf"
	"github.com/distanceNing/testapp/logic"
	"github.com/gin-gonic/gin"

	"io"
	"log"
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

func decodeBody(body io.Reader, v interface{}) common.Status {
	status := common.NewStatus()
	decoder := json.NewDecoder(body)
	err := decoder.Decode(v)
	if err != nil {
		log.Println(err.Error())
		status.Set(common.ErrJsonDecodeFail, "json decode failed")
	}
	return status
}

func (svr *HttpSvr) initRouter() {
	svr.engine.POST("/register", func(c *gin.Context) {
		rsp := common.NewRsp()
		loginReq := logic.LoginRequest{}
		status := decodeBody(c.Request.Body, &loginReq)
		if status.Ok() {
			status = svr.svc.Register(&loginReq, rsp)
		}
		rsp.SetStatus(&status)
		c.JSON(200, rsp.GetV())
	})

	svr.engine.POST("/login", func(c *gin.Context) {
		rsp := common.NewRsp()
		loginReq := logic.LoginRequest{}
		status := decodeBody(c.Request.Body, &loginReq)
		if status.Ok() {
			status = svr.svc.Login(&loginReq, rsp)
		}
		rsp.SetStatus(&status)
		c.JSON(200, rsp.GetV())
	})

	svr.engine.POST("/articles/create", func(c *gin.Context) {
		rsp := common.NewRsp()
		createArticleReq := logic.CreateArticleReq{}
		status := decodeBody(c.Request.Body, &createArticleReq)
		if status.Ok() {
			status = svr.svc.CreateArticle(&createArticleReq, rsp)
		}
		rsp.SetStatus(&status)
		c.JSON(200, rsp.GetV())
	})
}
