package api

import (
	"net/http"
	"log"
	"encoding/json"
)

const (
	SUCCESS      = 1
	SERVER_ERROR = 2
)

type Result struct {
	Code int `json:"code"`
	Msg  string `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func sendError(w http.ResponseWriter, err error) {
	log.Println(err)
	var r = &Result{
		Code: SERVER_ERROR,
		Msg:  "服务器错误",
	}
	j, err := json.Marshal(r)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	log.Println(string(j))
	w.Write(j)
}

func sendSuccess(w http.ResponseWriter, data interface{}) {
	var r = &Result{
		Code: SUCCESS,
		Msg:  "",
		Data: data,
	}
	j, err := json.Marshal(r)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	log.Println(string(j))
	w.Write(j)
}

