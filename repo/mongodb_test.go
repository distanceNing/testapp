package repo

import (
	"context"
	"testing"

	"github.com/distanceNing/testapp/conf"
)

func TestNewMongoDbInstance(t *testing.T) {
	type args struct {
		conf *conf.MongoDbConf
	}

	mongoDbConf := conf.MongoDbConf{Addr: "127.0.0.1:27017", Database: "test"}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"base", args{&mongoDbConf}, false},
		{"base_failed", args{&mongoDbConf}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewMongoDbInstance(tt.args.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMongoDbInstance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMongoDbInstance_InsertOne(t *testing.T) {
	mongoDbConf := conf.MongoDbConf{Addr: "127.0.0.1:27017", Database: "test"}
	type TestObj struct {
		Filed1 string
		Filed2 string
	}

	type TestArrayObj struct {
		Filed1 string
		Filed2 []string
	}

	type args struct {
		ctx         context.Context
		collectName string
		v           interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"base", args{context.TODO(), "test", TestObj{"x", "x"}}, false},
		{"base_array", args{context.TODO(), "test", TestArrayObj{Filed1: "x", Filed2: []string{"x", "y"}}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mongo, _ := NewMongoDbInstance(&mongoDbConf)
			if err := mongo.InsertOne(tt.args.ctx, tt.args.collectName, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("MongoDbInstance.InsertOne() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Log("mongo insert succ")
		})
	}
}
