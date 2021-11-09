package parser

import (
	"github.com/distanceNing/testapp/src/common/errcode"
	"gopkg.in/yaml.v2"
	"os"
)

func ParseYamlFile(path string, v interface{}) error {
	f, err := os.Open(path)
	if err != nil {
		return errcode.New(-1, "open file : "+err.Error())
	}

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(v)
	if err != nil {
		return errcode.New(-1, "decode failed ,"+err.Error())
	}
	return nil
}
