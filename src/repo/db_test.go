package repo

import (
	"github.com/distanceNing/testapp/src/conf"
	"reflect"
	"testing"
)

func TestNewDbInstance(t *testing.T) {
	dbconf := conf.DbConf{"127.0.0.1:3306", "root", "DLJn@123456!"}
	tests := []struct {
		name string
		want error
	}{
		{"base", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, got := NewDbInstance(&dbconf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDbInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}
