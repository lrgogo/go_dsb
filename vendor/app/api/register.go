package api

import (
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	"app/db"
)

func init() {
	http.HandleFunc("/register", func(writer http.ResponseWriter, request *http.Request) {
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
		vcode := kv["vcode"]
		//信息是否为空
		if mobile == "" || pwd == "" || vcode == "" {
			return
		}
		//验证手机验证码
		if vcode != "123456" {
			log.Println("vcode is wrong")
			return
		}
		//验证是否已注册
		uid := db.Register(mobile, pwd)
		if uid < 0 {
			log.Println("this mobile is registered")
			return
		}
		log.Println("register success uid = ", uid)
	})
}
