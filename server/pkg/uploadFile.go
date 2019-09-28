package pkg

import (
	"fmt"
	"io/ioutil"
	"os"

	e "grpc-tutorial-server/errors"
)

// UploadFile .
func UploadFile(filePath string, data []byte) {
	dir, err := os.Getwd()
	e.Handle(err)
	fmt.Println(filePath)
	// here we don't need to include '/files' directory
	// because in 'listFiles.go' we have this line which is ran when the app inits
	// os.Chdir(dir + "/files") thus the working directory is already '/files'
	err = ioutil.WriteFile(dir+"/"+filePath, data, 0644)
	e.Handle(err)
}
