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

type RedisConf struct {
	Addr     string `yaml:"addr"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type AppConf struct {
	Name string `yaml:"name"`
	Addr string `yaml:"addr"`
}

type ServerConf struct {
	AppConf AppConf `yaml:"appconf"`
	DbConf  DbConf  `yaml:"dbconf"`
}

func ReadConf(confPath string) (common.Status, *ServerConf) {
	status := common.NewStatus()
	f, err := os.Open(confPath)
	if err != nil {
		status.Set(-1, "open file : "+err.Error())
		return status, nil
	}

	decoder := yaml.NewDecoder(f)
	SvrConf := new(ServerConf)
	err = decoder.Decode(SvrConf)
	if err != nil {
		status.Set(-1, "decode failed ,"+err.Error())
		return status, nil
	}

	log.Println(SvrConf)
	return status, SvrConf
}
