package common

import "log"

type Executor struct {
	msgCodec RetryMessageCodec
	flowMap  map[string]*StateMachine
}

func NewExecutor() *Executor {
	return &Executor{flowMap: make(map[string]*StateMachine)}
}

func (e *Executor) RegisterFlow(name string, sm *StateMachine) {
	e.flowMap[name] = sm
}

func (e *Executor) ProcRetryMessage(msg *RetryMessage) {

}

func (e *Executor) ProcRequest(flowName string, req interface{}, rsp interface{}) {
	// find proc flow
	sm, ok := e.flowMap[flowName]
	if !ok {
		log.Printf("[%s] not register", flowName)
		return
	}
	log.Printf("do [%s] flow", flowName)
	e.do(sm.GetBeginState(), sm, req, rsp)
}

func (e *Executor) constructRetryMessage(cur State, sm *StateMachine, req interface{}) {
	msg := RetryMessage{}
	data, err := e.msgCodec.Serialize(&msg)
	if err != nil {
		return
	}
	// TODO push msg to mq
	log.Println(data)
}

func (e *Executor) do(cur State, sm *StateMachine, req interface{}, rsp interface{}) {
	err := sm.Run(&cur, req, rsp)
	if !err.Ok() && err.Code() == ErrNeedRetry {
		e.constructRetryMessage(cur, sm, e)
	}
}
