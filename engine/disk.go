package engine

import (
	"fmt"
	"os"
	"path/filepath"
)

// ScanDir walks through a directory and returns
// file paths contained inside.
func ScanDir(dir string) (files []string, err error) {
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if info.IsDir() {
			return nil
		}

		files = append(files, path)

		return nil
	})

	return files, err
}
