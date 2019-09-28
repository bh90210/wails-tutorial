package api

import (
	context "context"
	e "grpc-tutorial/errors"
)

// DeleteFile a file from server
func DeleteFile(filePath string) string {
	// connect to server
	conn, client := ConnectToServer()
	// and drop connection when done
	defer conn.Close()

	reply, err := client.Delete(context.Background(), &File{Path: filePath})
	e.Handle(err)

	if reply.Feedback == "succ" {
		return "succ"
	}

	return "err"
}
