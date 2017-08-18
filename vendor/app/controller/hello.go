package controller

import (
	"net/http"
	"fmt"
)

func init() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "hello world")
	})
}
