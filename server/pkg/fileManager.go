package fileManager

import (
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
func DownloadFile(fileName string) []byte {
	// get current working directory
	dir, err := os.Getwd()
	e.Handle(err)

	dat, err := ioutil.ReadFile(dir + "/" + fileName)
	e.Handle(err)

	return dat
}

// UploadFile .
func UploadFile(fileName string, data []byte) {
	// get wirking directory
	dir, err := os.Getwd()
	e.Handle(err)

	err = ioutil.WriteFile(dir+"/"+fileName, data, 0644)
	e.Handle(err)
}

// DeleteFile .
func DeleteFile(fileName string) {
	// get current working directory
	dir, err := os.Getwd()
	e.Handle(err)

	err = os.Remove(dir + "/" + fileName)
	e.Handle(err)
}
