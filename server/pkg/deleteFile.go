package pkg

import (
	"os"

	e "grpc-tutorial-server/errors"
)

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
