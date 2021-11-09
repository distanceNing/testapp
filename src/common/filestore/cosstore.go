package filestore

import (
	"context"
	"github.com/distanceNing/testapp/src/common/parser"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type CosConf struct {
	AppId        string `yaml:"AppId"`
	CosSecretId  string `yaml:"CosSecretId"`
	CosSecretKey string `yaml:"CosSecretKey"`
	CosUrl       string `yaml:"CosUrl"`
}

func GetDefaultCosConf() *CosConf {
	DefaultCosConf := &CosConf{}
	err := parser.ParseYamlFile("default.conf.yaml", &DefaultCosConf)
	if err != nil {
		return nil
	}
	return DefaultCosConf
}

type CosFileStore struct {
	url string
	c   *cos.Client
}

func NewCosFileStore(conf *CosConf) *CosFileStore {
	u, _ := url.Parse(conf.CosUrl)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		//设置超时时间
		Timeout: 100 * time.Second,
		Transport: &cos.AuthorizationTransport{
			//如实填写账号和密钥，也可以设置为环境变量
			SecretID:  conf.CosSecretId,
			SecretKey: conf.CosSecretKey,
		},
	})
	return &CosFileStore{url: conf.CosUrl, c: c}
}

func (cfs *CosFileStore) Get(path string) error {
	return nil
}

func (cfs *CosFileStore) Put(path string, r io.Reader) (string, error) {
	_, err := cfs.c.Object.Put(context.Background(), path, r, nil)
	return cfs.url + "/" + path, err
}
