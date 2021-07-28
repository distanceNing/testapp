package main

import (
	"github.com/distanceNing/testapp/conf"
	"github.com/distanceNing/testapp/repo"
	"log"
)

func main() {
	status, svrConf := conf.ReadConf("conf.yaml")
	if !status.Ok() {
		return
	}
	err := repo.InitStorage(&svrConf.DbConf)
	if err != nil {
		return
	}
	svr := NewHttpSvr(svrConf)

	err = svr.Run(svrConf.AppConf.Addr)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
