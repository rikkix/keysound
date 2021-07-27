package handler

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
}

func TestNewHandler(t *testing.T) {
	list := newList()
	plr := list.RandomPieceList().PieceListResponse(allChoices)
	for i := range plr {
		fmt.Println(*plr[i])
	}
}

func TestMongoNewQuiz(t *testing.T) {
	list := newList()
	plr := list.RandomPieceList().PieceListResponse(allChoices)
	for i := range plr {
		fmt.Println(*plr[i])
	}
	id, err := MongoNewQuiz(plr)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(id)
}
