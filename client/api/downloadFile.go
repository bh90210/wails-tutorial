package api

import (
	context "context"
	e "grpc-tutorial/errors"
)

// DownloadFile a file from server
func DownloadFile(filePath string) []byte {
	// connect to server
	conn, client := ConnectToServer()
	// and drop connection when done
	defer conn.Close()

	file, err := client.Download(context.Background(), &Request{Path: filePath})
	e.Handle(err)

	data := file.GetData()

	return data
}
