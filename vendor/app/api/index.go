package api

import (
	"net/http"
	"fmt"
)

func init() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "welcome")
	})
}
