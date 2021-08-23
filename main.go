package main

import (
	"github.com/distanceNing/testapp/common"
	"github.com/distanceNing/testapp/conf"
	"github.com/distanceNing/testapp/repo"
	"log"
)

const confFilePath = "conf.yaml"
const (
	s1 = 0
	s2 = 1
	s3 = 2
	s4 = 3
)

func DoS1(req interface{}, rsp interface{}, ext map[string]interface{}) (common.State, error) {
	log.Println("s1 next s2")
	return s2, common.NewSuccCode()
}

func DoS2(req interface{}, rsp interface{}, ext map[string]interface{}) (common.State, error) {
	log.Println("s2 next s3")
	ext["s3data"] = "xxxx"
	return s3, nil
}

func DoS3(req interface{}, rsp interface{}, ext map[string]interface{}) (common.State, error) {
	log.Printf("s3 next s4 , ext[%v]", ext)
	return s4, nil
}

func flowTest() {
	sm := common.NewStateMachine()
	sm.AddState([]*common.StateNode{common.NewStateNode(s1, DoS1), common.NewStateNode(s2, DoS2),
		common.NewStateNode(s3, DoS3)})
	sm.SetBeginState(s1)
	sm.SetEndState(s4)
	exec := common.NewExecutor()
	exec.RegisterFlow("test", sm)
	exec.ProcRequest("test", 1, 2)
}

func main() {
	// flowTest()
	err, svrConf := conf.ReadConf(confFilePath)
	if err != nil {
		return
	}
	err = repo.InitStorage(&svrConf.DbConf)
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
