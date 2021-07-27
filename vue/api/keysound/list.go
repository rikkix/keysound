package keysound

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
)

type MusicItem struct {
	ID         string
	Label      string
	Product    string
	Title      string
	Name       string
	SectionNum int
}

type MusicList []*MusicItem

func ParseList(path string) (MusicList, error) {
	ml := MusicList{}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return MusicList{}, err
	}
	reader := csv.NewReader(bytes.NewReader(file))
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return MusicList{}, err
		}
		if len(record) != 6 {
			return MusicList{}, fmt.Errorf("expect 6 columns, get %d columns", len(record))
		}
		num, err := strconv.Atoi(record[5])
		if err != nil {
			return MusicList{}, err
		}
		ml = append(ml, &MusicItem{
			ID:         record[0],
			Label:      record[1],
			Product:    record[2],
			Title:      record[3],
			Name:       record[4],
			SectionNum: num,
		})
	}
	return ml, nil
}
