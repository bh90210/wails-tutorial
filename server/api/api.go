package api

import (
	context "context"
	"flag"
	fmt "fmt"
	"net"

	e "grpc-tutorial-server/errors"
	fh "grpc-tutorial-server/fileManager"

	"google.golang.org/grpc"
)

type intercommService struct{}

// StartServer start listening for clients
func StartServer() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	e.Handle(err)

	s := grpc.NewServer()
	// attach service to the server
	RegisterIntercommServer(s, &intercommService{})
	// start the server
	e.Handle(s.Serve(lis))
}

func (s *intercommService) ListFiles(in *Request, stream Intercomm_ListFilesServer) error {
	// use a different package function
	// to collect all files
	files := fh.ListFiles()
	// iterate over the result to get the info
	for name, value := range files {
		// then iterate again to get the nested map values
		for path, size := range value {
			info := &File{Path: path, Name: name, Size: size}
			// stream each file's info to client
			stream.Send(info)
		}
	}

	return nil
}

func (s *intercommService) Upload(ctx context.Context, file *File) (*Reply, error) {
	path := file.GetPath()
	data := file.GetData()
	fh.UploadFile(path, data)

	reply := &Reply{Feedback: "succ"}
	return reply, nil
}

func (s *intercommService) Download(ctx context.Context, in *Request) (*File, error) {
	f := fh.DownloadFile(in.GetPath())
	file := &File{Data: f}

	return file, nil
}

func (s *intercommService) Delete(ctx context.Context, file *File) (*Reply, error) {
	fh.DeleteFile(file.Path)
	reply := &Reply{Feedback: "succ"}

	return reply, nil
}
