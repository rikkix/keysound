package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
)

type NewResponse struct {
	ID     string            `json:"id"`
	Base   string            `json:"base"`
	Pieces PieceListResponse `json:"pieces"`
}

func NewQuizHandler(w http.ResponseWriter, r *http.Request) {
	if (r.URL.Path != "/api/new" && r.URL.Path != "/new") ||
		r.Method != http.MethodPost {
		fmt.Printf("unknown request with path: %s, method: %s\n",
			r.URL.Path, r.Method)
		w.WriteHeader(400)
		return
	}
	plr := newList().RandomPieceList().PieceListResponse(allChoices)
	id, err := MongoNewQuiz(plr)
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

func newList() MusicList {
	ml := MusicList{}
	indexes := randomIndexes(10, len(musicList)-1)
	for i := 0; i < len(indexes); i++ {
		ml = append(ml, musicList[indexes[i]])
	}
	return ml
}

func randomIndexes(n, b int) []int {
	var il []int
	if n > b+1 {
		for i := 0; i <= b; i++ {
			il = append(il, i)
		}
		return il
	}
	for {
		in := randomIndex(b)
		if exist(in, il) {
			continue
		}
		il = append(il, in)
		if len(il) >= n {
			break
		}
	}
	return il
}

func exist(n int, l []int) bool {
	for i := 0; i < len(l); i++ {
		if l[i] == n {
			return true
		}
	}
	return false
}

// [0,b]
func randomIndex(b int) int {
	return rand.Intn(b + 1)
}
