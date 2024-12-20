package helpers

import (
	"fmt"
	"path/filepath"
)

func INclude(path string) []string {
	files, err := filepath.Glob("admin/views/templates/*.html")
	pathFiles, _ := filepath.Glob("admin/views/" + path + "/*.html")
	for _, file := range pathFiles {
		files = append(files, file)
	}
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return files
}
