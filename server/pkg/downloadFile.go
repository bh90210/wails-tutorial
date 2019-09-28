package pkg

import (
	"fmt"
	"io/ioutil"
	"os"

	e "grpc-tutorial-server/errors"
)

// DownloadFile .
func DownloadFile(filePath string) []byte {
	dir, err := os.Getwd()
	e.Handle(err)
	fmt.Println(filePath)
	// here we don't need to include '/files' directory
	// because in 'listFiles.go' we have this line which is ran when the app inits
	// os.Chdir(dir + "/files") thus the working directory is already '/files'
	//err = os.Remove(dir + "/" + filePath)
	//e.Handle(err)
	//fmt.Println("==> done deleting file")
	dat, err := ioutil.ReadFile(dir + "/" + filePath)
	e.Handle(err)

	return dat
}
