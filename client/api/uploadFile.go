package api

import (
	context "context"
	e "grpc-tutorial/errors"
)

// UploadFile a file from server
func UploadFile(filesList []string, dataList [][]byte) string {
	// connect to server
	conn, client := ConnectToServer()
	// and drop connection when done
	defer conn.Close()

	stream, err := client.Upload(context.Background())
	e.Handle(err)

	for i, info := range filesList {
		file := &File{Name: info, Data: dataList[i]}
		err := stream.Send(file)
		e.Handle(err)
	}
	reply, err := stream.CloseAndRecv()
	e.Handle(err)

	return reply.Feedback
}
