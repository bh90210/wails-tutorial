package api

import (
	context "context"
	"flag"
	fmt "fmt"
	"net"

	e "mockServer/errors"
	fh "mockServer/pkg"

	"google.golang.org/grpc"
)

// StartServer start listening for clients
func StartServer() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	e.Handle(err)

	// Create the TLS credentials
	//creds, err := credentials.NewServerTLSFromFile("cert/cert.pem", "cert/key.pem")
	//e.Handle(err)

	// Create an array of gRPC options with the credentials
	//opts := []grpc.ServerOption{grpc.Creds(creds)}
	// create a gRPC server object
	// s := grpc.NewServer(opts...)
	s := grpc.NewServer()
	// attach service to the server
	RegisterIntercommServer(s, &intercommService{})
	// start the server
	err2 := s.Serve(lis)
	e.Handle(err2)

	// grpcServer := grpc.NewServer(opts...)
	// pb.RegisterRouteGuideServer(grpcServer, newServer())
	// grpcServer.Serve(lis)
}

type intercommService struct{}

func (s *intercommService) Upload(stream Intercomm_UploadServer) error {

	return nil
}

func (s *intercommService) Download(ctx context.Context, in *DownloadRequest) (*File, error) {
	file := &File{Name: "test", Path: "path", Size: 666}
	return file, nil
}

func (s *intercommService) Delete(ctx context.Context, file *File) (*Feedback, error) {
	reply := &Feedback{Feedback: "succ"}
	return reply, nil
}

func (s *intercommService) ListFiles(in *ListFilesRequest, stream Intercomm_ListFilesServer) error {
	files := fh.ListFiles()
	fmt.Printf("%+v \n", files)

	list := &File{Path: "/test", Name: "test"}
	stream.Send(list)
	return nil
}

// func (fts *fileTransferService) ListFiles(_ *proto.ListRequestType, stream proto.FileTransferService_ListFilesServer) error {
// 	err := filepath.Walk(fts.root, func(p string, info os.FileInfo, err error) error {
// 		name, err := filepath.Rel(fts.root, p)
// 		if err != nil {
// 			return err
// 		}
// 		name = filepath.ToSlash(name)
// 		modTime := new(google_protobuf.Timestamp)
// 		modTime.Seconds = int64(info.ModTime().Unix())
// 		modTime.Nanos = int32(info.ModTime().UnixNano())
// 		f := &proto.ListResponseType{Name: name, Size: info.Size(), Mode: uint32(info.Mode()), ModTime: modTime}
// 		return stream.Send(f)
// 	})
// 	return err
// }
