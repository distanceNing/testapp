package repo

import (
	"github.com/distanceNing/testapp/common"
	"reflect"
	"testing"
)

func TestCreateObject(t *testing.T) {
	GetDefaultTestDb()
	type args struct {
		obj interface{}
	}
	obj := ArticleInfo{Title: "test", Content: "test", Status: 0}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"base", args{obj: &obj}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateObject(tt.args.obj); !reflect.DeepEqual(got.Code(), tt.want) {
				t.Errorf("CreateObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateObject(t *testing.T) {
	type args struct {
		cond        interface{}
		updateField interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"base", args{cond: &ArticleInfo{Id: 1}, updateField: ArticleInfo{Status: 1}}, common.ErrNoAffected},
		{"base", args{cond: &ArticleInfo{Id: 2}, updateField: ArticleInfo{Status: 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateObject(tt.args.cond, tt.args.updateField); !reflect.DeepEqual(got.Code(), tt.want) {
				t.Errorf("UpdateObject() = %v, want %v", got, tt.want)
			}
		})
	}
}
