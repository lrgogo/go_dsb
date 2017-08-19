package controller

import "net/http"

func init() {
	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {

	})
}
