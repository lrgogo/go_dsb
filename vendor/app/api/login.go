package api

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"app/db"
	"app/util"
)

type LoginUserInfo struct {
	Token      string `json:"token"`
	Uid        int64 `json:"uid"`
	Mobile     string `json:"mobile"`
	Profession string `json:"profession"`
	Corp       string `json:"corp"`
	Business   string `json:"business"`
}

func init() {
	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		//是否为POST
		if request.Method != "POST" {
			log.Println("request is not POST")
			return
		}
		//解析JSON
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			return
		}
		log.Println(string(body))
		var kv map[string]string
		err = json.Unmarshal(body, &kv)
		if err != nil {
			return
		}
		mobile := kv["mobile"]
		pwd := kv["pwd"]
		//信息是否为空
		if mobile == "" || pwd == "" {
			return
		}
		//验证是否已注册
		user, uid := db.Login(mobile, pwd)
		if uid == -1 {
			log.Println(err, "this mobile is not register")
			return
		}
		if uid == -2 {
			log.Println("password is wrong")
			return
		}
		loginUserInfo := &LoginUserInfo{
			Token:      util.CreateAccessJWT(uid),
			Uid:        user.Uid,
			Mobile:     user.Mobile,
			Profession: user.Profession,
			Business:   user.Business,
			Corp:       user.Corp,
		}
		sendSuccess(writer, loginUserInfo)
	})
}
