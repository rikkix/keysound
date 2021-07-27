package main

import (
	"bytes"
	"encoding/csv"
	"io/fs"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	buf := &bytes.Buffer{}

	writer := csv.NewWriter(buf)

	err := filepath.Walk("Split", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		parts := strings.Split(path, "/")
		if len(parts) != 3 {
			return nil
		}

		base := strings.TrimSuffix(parts[2], ".wav")
		fn := strings.Split(base, "-")
		if len(fn) != 2 {
			if len(fn) < 2 {
				log.Println(fn)

			} else {
				fn[1] = strings.Join(fn[1:], "-")
			}
		}

		rb := RandStringBytes(6)

		err = writer.Write([]string{rb, "Key", parts[1], fn[1]})
		if err != nil {
			return err
		}

		err = os.Mkdir(rb, 0755)
		if err != nil {
			return err
		}

		err = os.Rename(path, filepath.Join(rb, "full.wav"))

		return nil
	})

	if err != nil {
		log.Println(err)
		return
	}

	writer.Flush()

	err = ioutil.WriteFile("list.csv", buf.Bytes(), 0644)
	if err != nil {
		log.Println(err)
		return
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyz0123456789"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
