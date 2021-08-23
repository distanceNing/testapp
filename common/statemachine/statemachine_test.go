package statemachine

import (
	"testing"
)

func TestNewStateMachine(t *testing.T) {
	s := NewStateMachine()

}

func TestStateMachine_AddState(t *testing.T) {
	type fields struct {
		states map[State]StateNode
	}
	type args struct {
		node StateNode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StateMachine{
				states: tt.fields.states,
			}
			AddState(tt.args.node)
		})
	}
}
