package main

import (
	"github.com/distanceNing/testapp/conf"
	"github.com/distanceNing/testapp/logic"
	"github.com/distanceNing/testapp/repo"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// user_info := make(map[string]string)

func SayHello(w http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		body_fd := request.Body
		if body_fd == nil {
			return
		}
		bodyBytes, err := ioutil.ReadAll(body_fd)
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Print(string(bodyBytes))
	}

	if request.Method == http.MethodGet {
		log.Println(request.URL.RawQuery)
	}
	var code = request.URL.Query().Get("code")
	// GET https://api.weixin.qq.com/sns/jscode2session?
	// appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
	user_login := logic.UserLogin{}
	user_login.Login(code)

	file := "return_txt.log"
	fd, err := os.Open(file)
	if err != nil {
		log.Fatal(file, err.Error())
	}

	defer fd.Close()
	fileInfo, _ := fd.Stat()

	buffer := make([]byte, fileInfo.Size())

	_, err = fd.Read(buffer)
	if err != nil {
		log.Fatal("read file", err.Error())
	}
	//str := "{\"FrontSeqNo\":\"2005060250124115\",\"CnsmrSeqNo\":\"F088722005061581951001\",\"ret_msg\":\"ok\",\"ret_code\":\"0\"}"
	_, err = w.Write([]byte(buffer))
	if err != nil {
		log.Fatal("http resp write fail", err)
	}
}

func loginOp() {
	ul := logic.UserLogin{}
	ul.Init()
	gin.ForceConsoleColor()

	r := gin.Default()
	r.GET("/logic", func(c *gin.Context) {
		code := c.Request.URL.Query().Get("code")
		status := ul.Login(code)
		c.JSON(200, gin.H{
			"ret": status.Code(),
			"msg": status.Msg(),
		})
	})

	r.GET("/auth", func(c *gin.Context) {
		auth_req := logic.AuthRequest{Signature: c.Query("signature"),
			Code:          c.Query("code"),
			EncryptedData: c.Query("encryptedData"),
			Iv:            c.Query("iv"),
			RawData:       c.Query("rawData"),
		}
		status := ul.Auth(&auth_req)
		c.JSON(200, gin.H{
			"ret": status.Code(),
			"msg": status.Msg(),
		})
	})

	err := r.Run("127.0.0.1:18888")
	if err != nil {
		log.Printf(err.Error())
		return
	}
}

func main() {
	status, svrConf := conf.ReadConf("conf.yaml")
	if !status.Ok() {
		return
	}

	repo.Testgorm(&svrConf.DbConf)

}
