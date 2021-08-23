package conf

import (
	"github.com/distanceNing/testapp/common"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type DbConf struct {
	Addr     string `yaml:"addr"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type MongoDbConf struct {
	Addr     string `yaml:"addr"`
	Database string `yaml:"database"`
}

type RedisConf struct {
	Addr     string `yaml:"addr"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type AppConf struct {
	Name      string `yaml:"name"`
	Addr      string `yaml:"addr"`
	ImagePath string `yaml:"image_path"`
	CdnPath   string `yaml:"cdn_path"`
}

type ServerConf struct {
	AppConf     AppConf     `yaml:"appconf"`
	DbConf      DbConf      `yaml:"dbconf"`
	MongoDbConf MongoDbConf `yaml:"mongodbconf"`
	RedisConf   RedisConf   `yaml:"redisconf"`
}

func ReadConf(confPath string) (error, *ServerConf) {
	f, err := os.Open(confPath)
	if err != nil {
		log.Println(err.Error())
		return common.NewErrorCode(-1, "open file : "+err.Error()), nil
	}

	decoder := yaml.NewDecoder(f)
	SvrConf := new(ServerConf)
	err = decoder.Decode(SvrConf)
	if err != nil {
		log.Println(err.Error())
		return common.NewErrorCode(-1, "decode failed ,"+err.Error()), nil
	}
	return nil, SvrConf
}
