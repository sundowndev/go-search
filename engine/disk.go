package engine

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// GetFilesFromDir walks through a directory and returns
// file paths contained inside.
func GetFilesFromDir(dir string) (files []string) {
	err := filepath.Walk(dir, func(fp string, fi os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if fi.IsDir() {
			return nil
		}

		files = append(files, fp)

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return files
}
