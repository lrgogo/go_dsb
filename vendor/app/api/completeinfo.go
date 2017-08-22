package api

import (
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	"app/db"
	"app/util"
)

func init() {
	http.HandleFunc("/completeinfo", func(writer http.ResponseWriter, request *http.Request) {
		//是否为POST
		if request.Method != "POST" {
			log.Println("request is not POST")
			return
		}
		//解析token
		token := request.Header.Get("x-access-token")
		if token == "" {
			log.Println("token is invalid")
			return
		}
		uid := util.VerifyJWT(token)
		if uid == -1 {
			log.Println("token is invalid")
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
		profession := kv["profession"]
		business := kv["business"]
		corp := kv["corp"]

		if profession == "" || business == "" || corp == "" {
			return
		}

		if db.CompleteInfo(uid, profession, business, corp) == -1 {
			log.Println("mysql error")
			return
		}
		sendSuccessNoData(writer)
	})
}
