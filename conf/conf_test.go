package conf

import (
	"reflect"
	"testing"
)

func TestReadConf(t *testing.T) {
	type args struct {
		confPath string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"base", args{"../conf.yaml"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ReadConf(tt.args.confPath)
			if !reflect.DeepEqual(got.Code(), tt.want) {
				t.Errorf("ReadConf() got = %v, want %v", got, tt.want)
			}

		})
	}
}
