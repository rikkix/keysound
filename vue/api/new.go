package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/iochen/keysound/vue/keysound"
)

type NewResponse struct {
	ID     string                     `json:"id"`
	Base   string                     `json:"base"`
	Pieces keysound.PieceListResponse `json:"pieces"`
}

func NewQuizHandler(w http.ResponseWriter, r *http.Request) {
	if (r.URL.Path != "/api/new" && r.URL.Path != "/new") ||
		r.Method != http.MethodPost {
		fmt.Printf("unknown request with path: %s, method: %s\n",
			r.URL.Path, r.Method)
		w.WriteHeader(400)
		return
	}
	plr := keysound.NewList().RandomPieceList().PieceListResponse(keysound.DefaultChoices)
	id, err := keysound.MongoNewQuiz(plr)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	err = json.NewEncoder(w).Encode(NewResponse{
		ID:     id,
		Base:   os.Getenv("BASE_URL"),
		Pieces: plr,
	})
	if err != nil {
		if err == bson.ErrNilRegistry {
			w.WriteHeader(404)
			return
		}
		log.Println(err)
		w.WriteHeader(500)
		return
	}
}
