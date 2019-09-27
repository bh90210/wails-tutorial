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
	err = os.Remove(dir + "/files/" + filePath)
	e.Handle(err)
	fmt.Println("==> done deleting file")
}
