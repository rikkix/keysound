package main

import (
	"fmt"
	"net/http"

	handler "github.com/iochen/keysound/api"
)

func main() {
	http.HandleFunc("/api/get", keysound.GetQuizHandler)
	http.HandleFunc("/api/new", handler.NewQuizHandler)
	fmt.Println(http.ListenAndServe(":8088", nil))
}
