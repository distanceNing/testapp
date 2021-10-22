package filestore

import (
	"io"
	"strings"
	"testing"
)

func Test_test_cos(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"base"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp := ""
			test_cos(&rsp)
			t.Error(rsp)
		})
	}
}

func TestCosFileStore_Put(t *testing.T) {
	type args struct {
		path string
		r    io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"base", args{"test.txt", strings.NewReader("test cos")}, false}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfs := NewCosFileStore()
			if err := cfs.Put(tt.args.path, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("CosFileStore.Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
