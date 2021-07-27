package repo

import (
	"reflect"
	"testing"
)

func TestCreateUser(t *testing.T) {
	GetDefaultTestDb()
	type args struct {
		userId   string
		userType int
		password string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"base", args{"test", 0, "test"}, 0,},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateUser(tt.args.userId, tt.args.userType, tt.args.password); !reflect.DeepEqual(got.Code(), tt.want) {
				t.Errorf("CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
