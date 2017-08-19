package controller

import (
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	"app/database"
	"app/api"
)

type RegisterInfo struct {
	Mobile string `json:"mobile"`
	Pwd    string `json:"pwd"`
	Vcode  string `json:"vcode"`
}

func init() {
	http.HandleFunc("/register", func(writer http.ResponseWriter, request *http.Request) {
		//是否为POST
		if request.Method != "POST" {
			log.Println("request is not POST")
			return
		}
		//解析JSON
		arr, err := ioutil.ReadAll(request.Body)
		if err != nil {
			return
		}
		log.Println(string(arr))
		info := &RegisterInfo{}
		err = json.Unmarshal(arr, info)
		if err != nil {
			return
		}
		//信息是否为空
		if info.Mobile == "" || info.Pwd == "" || info.Vcode == "" {
			return
		}
		//验证是否已注册
		rows, err := database.SQL.Query("SELECT user.uid FROM user WHERE user.mobile=?", info.Mobile)
		if rows.Next() {
			log.Println("this mobile is registered")
			return
		}
		//验证手机验证码
		if info.Vcode != "123456" {
			log.Println("vcode is wrong")
			return
		}
		//插入数据
		result, err := database.SQL.Exec("INSERT INTO user(user.mobile, user.pwd)VALUES (?,?)", info.Mobile, info.Pwd)
		if err != nil {
			api.E(writer, err)
			return
		}
		newUid, err := result.LastInsertId()
		log.Println("register success uid = ", newUid)
	})
}
