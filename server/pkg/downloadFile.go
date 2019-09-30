package pkg

import (
	"fmt"
	"io/ioutil"
	"os"

	e "grpc-tutorial-server/errors"
)

// DownloadFile .
func DownloadFile(filePath string) []byte {
	// get current working directory
	dir, err := os.Getwd()
	e.Handle(err)

	fmt.Println(filePath)
	// here we don't need to include '/files' directory
	// because in 'listFiles.go' we have this line which is ran when the app inits
	// os.Chdir(dir + "/files") thus the working directory is already '/files'
	dat, err := ioutil.ReadFile(dir + "/" + filePath)
	e.Handle(err)

	return dat
}
