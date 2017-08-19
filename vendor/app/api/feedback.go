package api

import (
	"net/http"
	"fmt"
)

func init() {
	http.HandleFunc("/feedback", func(writer http.ResponseWriter, request *http.Request) {
		arr := []int{1, 2, 3}
		fmt.Fprint(writer, arr[5])
	})
}
