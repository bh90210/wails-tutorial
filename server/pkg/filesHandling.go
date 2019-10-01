package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	e "grpc-tutorial-server/errors"
)

// type FilesHandle interface {
// 	DownloadFile()
// }

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

// DownloadFile .
func DownloadFile(filePath string) []byte {
	// get current working directory
	dir, err := os.Getwd()
	e.Handle(err)

	fmt.Println(filePath)
	dat, err := ioutil.ReadFile(dir + "/" + filePath)
	e.Handle(err)

	return dat
}

// UploadFile .
func UploadFile(filePath string, data []byte) {
	// get wirking directory
	dir, err := os.Getwd()
	e.Handle(err)

	// here we don't need to include '/files' directory
	// because in 'listFiles.go' we have this line which is ran when the app inits
	// os.Chdir(dir + "/files") thus the working directory is already '/files'
	err = ioutil.WriteFile(dir+"/"+filePath, data, 0644)
	e.Handle(err)
}

// DeleteFile .
func DeleteFile(filePath string) {
	// get current working directory
	dir, err := os.Getwd()
	e.Handle(err)

	// here we don't need to include '/files' directory
	// because in 'listFiles.go' we have this line which is ran when the app inits
	// os.Chdir(dir + "/files") thus the working directory is already '/files'
	err = os.Remove(dir + "/" + filePath)
	e.Handle(err)
}
