package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	count := map[string]int{}
	err := filepath.WalkDir("data", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		parts := strings.Split(path, "/")
		if len(parts) < 3 {
			return nil
		}

		count[parts[1]]++

		fmt.Println(path)

		return nil
	})
	if err != nil {
		log.Println(err)
		return
	}

	file, err := os.Open("list.csv")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	nfile, err := os.OpenFile("list.withnum.csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	defer nfile.Close()

	writer := csv.NewWriter(nfile)

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Println(err)
				return
			}
		}
		err = writer.Write(append(record, strconv.Itoa(count[record[0]]-2)))
		if err != nil {
			log.Println(err)
			return
		}
	}
	writer.Flush()
}
