package engine

import (
	"fmt"
	"os"
	"path/filepath"
)

// GetFilesFromDir walks through a directory and returns
// file paths contained inside.
func GetFilesFromDir(dir string) (files []string) {
	// is dir?
	// is readable/exists?
	filepath.Walk(dir, func(fp string, fi os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		fi.Mode()
		if fi.IsDir() {
			return nil
		}

		files = append(files, fp)

		return nil
	})

	return
}
