package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
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

	for k, v := range count {
		path := filepath.Join("data", k,
			fmt.Sprintf("%d.mp3", v-2))
		err := os.Remove(path)
		if err != nil {
			panic(err)
			return
		}
	}

	fmt.Println(len(count))

}
