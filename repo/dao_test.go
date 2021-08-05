package repo

import (
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
		{"base", args{cond: &ArticleInfo{Id: 1}, updateField: ArticleInfo{Status: 1}}, 0},
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

func TestQueryObjectByPage(t *testing.T) {
	GetDefaultTestDb()
	type args struct {
		cond      interface{}
		objs      interface{}
		pageCount int
		pageNum   int
	}
	var objs []ArticleInfo
	arg := args{&ArticleInfo{Status: 0}, &objs, 0, 2}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"base", arg, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryObjectByPage(tt.args.cond, tt.args.objs, tt.args.pageCount, tt.args.pageNum); !reflect.DeepEqual(got.Code(), tt.want) {
				t.Errorf("QueryObjectByPage() = %v, want %v", got.Code(), tt.want)
			}
			t.Log(tt.args.objs)
		})
	}
}

func TestQueryObjectCount(t *testing.T) {
	GetDefaultTestDb()
	type args struct {
		cond  interface{}
		count *int64
	}
	var count int64
	tests := []struct {
		name  string
		args  args
		count int64
		want  int
	}{
		{"base", args{cond: &ArticleInfo{Id: 2}, count: &count}, 3, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryObjectCount(tt.args.cond, tt.args.count); !reflect.DeepEqual(tt.want, got.Code()) || tt.count != *tt.args.count {
				t.Errorf("QueryObjectCount() = %v, want %v,  count = %v, count = %v", got, tt.want, tt.count, *tt.args.count)
			}
		})
	}
}
