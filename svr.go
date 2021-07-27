package main

import (
	"encoding/json"
	"github.com/distanceNing/testapp/common"
	"github.com/distanceNing/testapp/conf"
	"github.com/distanceNing/testapp/logic"
	"github.com/distanceNing/testapp/repo"

	"github.com/gin-gonic/gin"

	"log"
)

// user_info := make(map[string]string)

type HttpSvr struct {
	engine *gin.Engine
}

func NewHttpSvr() *HttpSvr {
	return &HttpSvr{gin.Default()}
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

func (svr *HttpSvr) initRouter() {
	svr.engine.POST("/login", func(c *gin.Context) {
		svc := logic.NewLoginService()
		decoder := json.NewDecoder(c.Request.Body)
		loginReq := logic.LoginRequest{}
		err := decoder.Decode(&loginReq)
		if err != nil {
			log.Println(err.Error())
			c.JSON(200, gin.H{
				"ret": common.ErrJsonDecodeFail,
				"msg": "json decode failed",
			})
			return
		}
		rsp := common.NewRsp()
		status := svc.Login(&loginReq, rsp)
		rsp.SetStatus(&status)
		c.JSON(200, rsp.GetV())
	})
}

func main() {
	status, svrConf := conf.ReadConf("conf.yaml")
	if !status.Ok() {
		return
	}
	err := repo.InitStorage(&svrConf.DbConf)
	if err != nil {
		return
	}
	svr := NewHttpSvr()
	err = svr.Run(svrConf.AppConf.Addr)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
