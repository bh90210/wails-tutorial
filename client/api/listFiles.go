package api

import (
	context "context"
	e "grpc-tutorial/errors"
	"io"
	"log"
	"strconv"
)

// ListFiles get all files from server
func ListFiles() [][]string {
	// connect to server
	conn, client := ConnectToServer()
	// and drop connection when done
	defer conn.Close()

	stream, err := client.ListFiles(context.Background(), &Request{List: true})
	e.Handle(err)

	//list := make(map[string]map[string]int32)
	var filesList [][]string

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
			log.Printf("++++++++ Got message %s, %s, %v", in.Path, in.Name, in.Size)
			//list[in.Name] = map[string]int32{in.Path: in.Size}
			// create an array to place the files
			intToString := strconv.FormatInt(int64(in.Size), 10)
			var entry []string
			entry = append(entry, in.Path)
			entry = append(entry, in.Name)
			entry = append(entry, intToString)
			filesList = append(filesList, entry)
		}
	}()
	stream.CloseSend()
	<-waitc

	return filesList
}
