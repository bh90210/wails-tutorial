package pkg

import (
	e "grpc-tutorial-server/errors"
	"os"
	"path/filepath"
)

// ListFiles .
func ListFiles() map[string]map[string]int64 {
	// get current working directory
	dir, err := os.Getwd()
	e.Handle(err)

	// cd in '/files' which is where files are stored server side
	os.Chdir(dir + "/files")
	list := make(map[string]map[string]int64)
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		e.Handle(err)

		// leave directories out
		if !info.IsDir() {
			list[info.Name()] = map[string]int64{path: info.Size()}
		}
		return nil
	})
	e.Handle(err)

	return list
}
