package common

import (
	"log"
)

type State int

type NodeHandler func(req interface{}, rsp interface{}, ext map[string]interface{}) (State, ErrorCode)

type StateNode struct {
	Do    NodeHandler
	State State
}

func NewStateNode(state State, do NodeHandler) *StateNode {
	return &StateNode{Do: do, State: state}
}

type StateMachine struct {
	states map[State]*StateNode
	begin  State
	end    State
}

func NewStateMachine() *StateMachine {
	return &StateMachine{states: make(map[State]*StateNode)}
}

func (s *StateMachine) SetBeginState(state State) {
	s.begin = state
}
func (s *StateMachine) SetEndState(state State) {
	s.end = state
}

func (s *StateMachine) GetBeginState() State {
	return s.begin
}

func (s *StateMachine) AddState(node *StateNode) {
	s.states[node.State] = node
}

func (s *StateMachine) Run(cur *State, req interface{}, rsp interface{}) ErrorCode {
	// 节点中间处理数据
	ext := make(map[string]interface{})
	for ; *cur != s.end; {
		node, ok := s.states[*cur]
		if !ok {
			log.Printf("state :%d not register ", *cur)
			return NewErrorCode(ErrSystem, "state not register")
		}
		next, err := node.Do(req, rsp, ext)
		if !err.Ok() {
			return err
		}
		*cur = next
	}
	return NewSuccCode()
}
