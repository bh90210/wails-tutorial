package api

import (
	context "context"
	"grpc-tutorial/errors"
	"io"
	"log"
)

// ListFiles get all files from server
func ListFiles() {
	// connect to server
	conn, client := ConnectToServer()
	// and drop connection when done
	defer conn.Close()

	stream, err := client.ListFiles(context.Background(), &Request{List: true})
	errors.Handle(err)

	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note : %v", err)
			}
			log.Printf("Got message %s, %s, %v)", in.Path, in.Name, in.Size)
		}
	}()
	stream.CloseSend()
	<-waitc
}
