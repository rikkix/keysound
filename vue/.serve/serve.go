package main

import (
	"fmt"
	"net/http"

	handler "github.com/iochen/keysound/vue/api"
)

func main() {
	http.HandleFunc("/api/get", handler.GetQuizHandler)
	http.HandleFunc("/api/new", handler.NewQuizHandler)
	fmt.Println(http.ListenAndServe(":8088", nil))
}
