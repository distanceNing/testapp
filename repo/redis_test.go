package repo

import (
	"reflect"
	"testing"

	"github.com/distanceNing/testapp/conf"
)



func TestNewRedisInstance(t *testing.T) {
	type args struct {
		conf *conf.RedisConf
	}
	tests := []struct {
		name string
		args args
		want *RedisInstance
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRedisInstance(tt.args.conf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRedisInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}
