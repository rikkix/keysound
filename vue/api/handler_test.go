package handler

import (
	"net/http"
	"testing"
)

func TestHandlers(t *testing.T) {
	http.HandleFunc("/api/new", NewQuizHandler)
	http.HandleFunc("/api/get", GetQuizHandler)
	t.Error(http.ListenAndServe(":8088", nil))
}
