package filestore

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type CosConf struct {
	AppId        string
	CosSecretId  string
	CosSecretKey string
	CosUrl       string
}

var DefaultCosConf = CosConf{
	AppId:        "1258428352",
	CosSecretId:  "AKIDghU0GFxOXP0KNAgV3MO9dpmSr1ypWyA5",
	CosSecretKey: "kFemeFZdQwCBuMXR4RYOxFGzSeZDafT0",
	CosUrl:       "https://test-1258428352.cos.ap-guangzhou.myqcloud.com",
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

func test_cos(rsp *string) {
	// AppId        := "1258428352"
	CosSecretId := "AKIDghU0GFxOXP0KNAgV3MO9dpmSr1ypWyA5"
	CosSecretKey := "kFemeFZdQwCBuMXR4RYOxFGzSeZDafT0"
	CosUrl := "https://test-1258428352.cos.ap-guangzhou.myqcloud.com"
	//将<bucket>和<region>修改为真实的信息
	//bucket的命名规则为{name}-{appid} ，此处填写的存储桶名称必须为此格式
	u, _ := url.Parse(CosUrl)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		//设置超时时间
		Timeout: 100 * time.Second,
		Transport: &cos.AuthorizationTransport{
			//如实填写账号和密钥，也可以设置为环境变量
			SecretID:  CosSecretId,
			SecretKey: CosSecretKey,
		},
	})

	name := "test.png"
	resp, err := c.Object.Get(context.Background(), name, nil)
	if err != nil {
		panic(err)
	}
	bs, _ := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	*rsp = string(bs)
}
