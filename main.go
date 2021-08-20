package main

import (
	"github.com/distanceNing/testapp/common"
	"log"
)

const confFilePath = "conf.yaml"
const (
	s1 = 0
	s2 = 1
	s3 = 2
	s4 = 3
)

type Node1 struct {
}

func (n *Node1) Do(req interface{}, rsp interface{}, ext map[string]interface{}) (common.State, common.ErrorCode) {
	log.Println("s1 next s2")
	return s2, common.NewSuccCode()
}

func (n *Node1) State() common.State {
	return s1
}

type Node2 struct {
}

func (n *Node2) Do(req interface{}, rsp interface{}, ext map[string]interface{}) (common.State, common.ErrorCode) {
	log.Println("s2 next s3")
	ext["s3data"] = "xxxx"
	return s3, common.NewSuccCode()
}

func (n *Node2) State() common.State {
	return s2
}

type Node3 struct {
}

func (n *Node3) Do(req interface{}, rsp interface{}, ext map[string]interface{}) (common.State, common.ErrorCode) {
	log.Printf("s3 next s4 , ext[%v]", ext)
	return s4, common.NewSuccCode()
}

func (n *Node3) State() common.State {
	return s3
}

func flowTest() {
	sm := common.NewStateMachine()
	sm.AddState(&Node1{})
	sm.AddState(&Node2{})
	sm.AddState(&Node3{})
	sm.SetBeginState(s1)
	sm.SetEndState(s4)

	exec := common.NewExecutor()
	exec.RegisterFlow("test", sm)
	exec.ProcRequest("test", 1, 2)
}

func main() {
	flowTest()
	//status, svrConf := conf.ReadConf(confFilePath)
	//if !status.Ok() {
	//	return
	//}
	//err := repo.InitStorage(&svrConf.DbConf)
	//if err != nil {
	//	return
	//}
	//svr := NewHttpSvr(svrConf)
	//err = svr.Run(svrConf.AppConf.Addr)
	//if err != nil {
	//	log.Println(err.Error())
	//	return
	//}
}
