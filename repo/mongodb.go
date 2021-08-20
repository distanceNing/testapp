package repo

import (
	"context"
	"github.com/distanceNing/testapp/conf"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoDbInstance struct {
	client *mongo.Client
	conf   *conf.MongoDbConf
}

func NewMongoDbInstance(conf *conf.MongoDbConf) (*MongoDbInstance, error) {
	uri := "mongodb://" + conf.Addr
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	return &MongoDbInstance{client: client, conf: conf}, nil
}

func (mongo *MongoDbInstance) InsertOne(ctx context.Context, collectName string, v interface{}) error {
	db := mongo.client.Database(mongo.conf.Database)
	insertResult, err := db.Collection(collectName).InsertOne(ctx, v)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(insertResult.InsertedID)
	return nil
}

func (mongo *MongoDbInstance) Update(ctx context.Context, collectName string, id interface{}, updateField interface{}) error {
	db := mongo.client.Database(mongo.conf.Database)
	res, err := db.Collection(collectName).UpdateByID(ctx, id, updateField)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.MatchedCount)
	return nil
}
