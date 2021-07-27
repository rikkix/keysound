package keysound

import (
	"fmt"
	"math/rand"
)

type MusicPiece struct {
	MusicItem  *MusicItem
	SectionNum int
}

type PieceList []*MusicPiece

type Choices []string

type MusicPieceResponse struct {
	Path    string  `json:"path"`
	Full    string  `json:"full"`
	Name    string  `json:"name"`
	Title   string  `json:"title"`
	Choices Choices `json:"choices"`
}

type PieceListResponse []*MusicPieceResponse

func (ml MusicList) Choices() Choices {
	var c []string
	for i := range ml {
		c = append(c, ml[i].Name)
	}
	return removeDuplicateStr(c)
}

func (mi *MusicItem) RandomMusicPiece() *MusicPiece {
	return &MusicPiece{
		MusicItem:  mi,
		SectionNum: randomIndex(mi.SectionNum),
	}
}

func (ml MusicList) RandomPieceList() PieceList {
	pl := make([]*MusicPiece, len(ml))
	for i := 0; i < len(pl); i++ {
		pl[i] = ml[i].RandomMusicPiece()
	}
	return pl
}

func (mp *MusicPiece) ToResponse(choices Choices) *MusicPieceResponse {
	mpr := &MusicPieceResponse{
		Path:  fmt.Sprintf("%s/%d.mp3", mp.MusicItem.ID, mp.SectionNum),
		Full:  fmt.Sprintf("%s/full.mp3", mp.MusicItem.ID),
		Name:  mp.MusicItem.Name,
		Title: mp.MusicItem.Title,
	}
	for {
		indexes := randomIndexes(3, len(choices)-1)
		for i := 0; i < len(indexes); i++ {
			if choices[indexes[i]] == mpr.Name {
				continue
			}
		}

		for i := 0; i < len(indexes); i++ {
			mpr.Choices = append(mpr.Choices, choices[indexes[i]])
		}
		break
	}
	return mpr
}

func (pl PieceList) PieceListResponse(choices Choices) PieceListResponse {
	plr := make([]*MusicPieceResponse, len(pl))
	for i := range plr {
		plr[i] = pl[i].ToResponse(choices)
	}
	return plr
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	var list []string
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func NewList() MusicList {
	ml := MusicList{}
	indexes := randomIndexes(10, len(DefaultMusicList)-1)
	for i := 0; i < len(indexes); i++ {
		ml = append(ml, DefaultMusicList[indexes[i]])
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
