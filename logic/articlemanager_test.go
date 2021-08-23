package logic

import (
	"github.com/distanceNing/testapp/common/types"
	"reflect"
	"testing"

	"github.com/distanceNing/testapp/repo"
)

func TestArticleManager_GetArticle(t *testing.T) {
	repo.GetDefaultTestDb()
	type args struct {
		req *GetArticleReq
		rsp *types.Rsp
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"base", args{&GetArticleReq{Id: 1}, types.NewRsp()}, 0}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mgr := &ArticleManager{}
			if got := mgr.GetArticle(tt.args.req, tt.args.rsp); !reflect.DeepEqual(got.Code(), tt.want) {
				t.Errorf("ArticleManager.GetArticle() = %v, want %v", got.Code(), tt.want)
			}
		})
	}
}

func TestArticleManager_CreateArticle(t *testing.T) {
	repo.GetDefaultTestDb()
	type args struct {
		req *CreateArticleReq
		rsp *types.Rsp
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"base", args{&CreateArticleReq{Title: "test", Content: "test content", Status: 0}, types.NewRsp()}, 0}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mgr := &ArticleManager{}
			if got := mgr.CreateArticle(tt.args.req, tt.args.rsp); !reflect.DeepEqual(got.Code(), tt.want) {
				t.Errorf("ArticleManager.CreateArticle() = %v, want %v", got, tt.want)
			}
		})
	}
}
