package conf

import (
	"github.com/distanceNing/testapp/comm"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type DbConf struct {
	Addr     string `yaml:"addr"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type AppConf struct {
	Name string `yaml:"name"`
}

type ServerConf struct {
	AppConf AppConf `yaml:"appconf"`
	DbConf  DbConf  `yaml:"dbconf"`
}

func ReadConf(confPath string) (comm.Status, *ServerConf) {
	status := comm.NewStatus()
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
