package pkg

import (
	"fmt"
	"os"

	e "grpc-tutorial-server/errors"
)

// DeleteFile .
func DeleteFile(filePath string) {
	dir, err := os.Getwd()
	e.Handle(err)
	fmt.Println(filePath)
	// here we don't need to include '/files' directory
	// because in 'listFiles.go' we have this line which is ran when the app inits
	// os.Chdir(dir + "/files") thus the working directory is already '/files'
	err = os.Remove(dir + "/" + filePath)
	e.Handle(err)
	//fmt.Println("==> done deleting file")
}
