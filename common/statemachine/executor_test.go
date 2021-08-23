package statemachine

import "testing"

func TestExecutor_ProcRequest(t *testing.T) {
	type fields struct {
		msgCodec RetryMessageCodec
		flowMap  map[string]*StateMachine
	}
	type args struct {
		flowName string
		req      interface{}
		rsp      interface{}
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
			e := &Executor{
				msgCodec: tt.fields.msgCodec,
				flowMap:  tt.fields.flowMap,
			}
			ProcRequest(tt.args.flowName, tt.args.req, tt.args.rsp)
		})
	}
}
