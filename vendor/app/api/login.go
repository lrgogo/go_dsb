package api

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"app/db"
)

type LoginInfo struct {
	Mobile string `json:"mobile"`
	Pwd    string `json:"pwd"`
}

type LoginedInfo struct {
	Token string `json:"token"`
	User  *db.User `json:"user"`
}

func init() {
	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
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
		info := &LoginInfo{}
		err = json.Unmarshal(arr, info)
		if err != nil {
			return
		}
		//信息是否为空
		if info.Mobile == "" || info.Pwd == "" {
			return
		}
		//验证是否已注册
		user := db.User{}
		row, err := db.DB.Query("SELECT * FROM user WHERE user.mobile=?", info.Mobile)
		if !row.Next() {
			log.Println(err, "this mobile is not register")
			return
		}
		row.Scan(&user.Mobile, &user.Pwd)
		if info.Pwd != user.Pwd {
			log.Println("password is wrong")
			return
		}
		loginedInfo := &LoginedInfo{
			Token: "abc",
			User:  &user,
		}
		sendSuccess(writer, loginedInfo)
	})
}
