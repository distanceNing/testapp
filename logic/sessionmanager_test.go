package logic

import (
	"reflect"
	"testing"

	"github.com/distanceNing/testapp/repo"
)

func TestSessionManager_CreateSession(t *testing.T) {
	type fields struct {
		redisCli *repo.RedisInstance
	}
	type args struct {
		userId string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"base", fields{repo.GetDefaultRedis()}, args{"test"}}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mgr := &SessionManager{
				redisCli: tt.fields.redisCli,
			}
			status, token := mgr.CreateSession(tt.args.userId)
			if !status.Ok() {
				t.Error("redis process failed")
			}

			if token != "xxxx" {
				t.Errorf("CreateSession() = %v, want %v", token, "xxxx")
			}

		})
	}
}

func TestSessionManager_QuerySessionToken(t *testing.T) {
	type fields struct {
		redisCli *repo.RedisInstance
	}
	type args struct {
		userId string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"base", fields{repo.GetDefaultRedis()}, args{"test"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mgr := &SessionManager{
				redisCli: tt.fields.redisCli,
			}
			got, _ := mgr.QuerySessionToken(tt.args.userId)
			if !reflect.DeepEqual(got.Code(), tt.want) {
				t.Errorf("SessionManager.QuerySessionToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
