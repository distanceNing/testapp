package login

import (
	"encoding/json"
	"github.com/distanceNing/testapp/comm"
	"github.com/distanceNing/testapp/utils"
	"log"
	"net/http"
	"sync"
)

type UserLogin struct {
	mutex        sync.Mutex
	userSessions map[string]SessionInfo
}

func (ul *UserLogin) Init() {
	ul.userSessions = make(map[string]SessionInfo)
}

type SessionInfo struct {
	sessionKey string
	openId     string
	unionId    string
}

type AuthReply struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

var wxAuthUrl = "https://api.weixin.qq.com/sns/jscode2session"
var wxAppid = "xx"
var grantType = "authorization_code"
var secretId = "xx"

func (ul *UserLogin) Login(code string) comm.Status {
	status := comm.NewStatus()
	req, err := http.NewRequest("GET", wxAuthUrl, nil)
	if err != nil {
		log.Println(err.Error())
		return status
	}
	args := req.URL.Query()
	args.Add("appid", wxAppid);
	args.Add("secret", secretId);
	args.Add("js_code", code);
	args.Add("grant_type", grantType);
	req.URL.RawQuery = args.Encode()
	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)

	if err != nil {
		log.Println(err.Error())
		return status
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var reply AuthReply
	err = decoder.Decode(&reply)
	if err != nil {
		log.Println(err.Error())
		return status
	}
	if reply.ErrCode != 0 {
		log.Printf("error code [%d],error msg [%s]", reply.ErrCode, reply.ErrMsg)
		return status
	}
	ul.mutex.Lock()
	defer ul.mutex.Unlock()
	ul.userSessions[code] = SessionInfo{reply.SessionKey, reply.OpenId, reply.UnionId}
	return status
}

type AuthRequest struct {
	Signature     string `json:"signature"`
	Code          string `json:"code"`
	EncryptedData string `json:"encryptedData"`
	Iv            string `json:"iv"`
	RawData       string `json:"rawData"`
}
type WaterMark struct {
	timestamp int
	appid     string
}
type UserInfo struct {
	OpenId    string `json:"openid"`
	NickName  string
	gender    int
	language  string
	city      string
	province  string
	country   string
	AvatarUrl string
	UnionId   string
	Watermark WaterMark
}

func (ul *UserLogin) Auth(req *AuthRequest) comm.Status {
	status := comm.NewStatus()
	ul.mutex.Lock()
	info, ok := ul.userSessions [ req.Code ]
	ul.mutex.Unlock()
	if !ok {
		log.Printf("req.code [%s] not exist", req.Code)
		status.Set(-1, "req.code not exist");
		return status
	}
	orig_data := utils.AesDecrypt(info.sessionKey, req.EncryptedData, req.Iv)
	var user_info UserInfo
	err := json.Unmarshal(orig_data, &user_info)
	if err != nil {
		log.Println(err.Error())
		return status
	}
	// verify sign
	if !utils.VerifySignBySha1(req.RawData+info.sessionKey, req.Signature) {
		log.Println("verify sign fail")
		return status
	}
	return status
}
