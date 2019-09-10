package clientCommunication

import (
	"context"
	"net"

	pb "mockServer/api/protobuf"
	e "mockServer/errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	port = ":50051"
)

// StartServer start listening for clients
func StartServer() {
	lis, err := net.Listen("tcp", port)
	e.Handle(err)

	// Create the TLS credentials
	creds, err := credentials.NewServerTLSFromFile("cert/cert.pem", "cert/key.pem")
	e.Handle(err)

	// Create an array of gRPC options with the credentials
	opts := []grpc.ServerOption{grpc.Creds(creds)}
	// create a gRPC server object
	s := grpc.NewServer(opts...)
	// attach service to the server
	pb.RegisterIntercommServer(s, &intercommService{})
	// start the server
	err2 := s.Serve(lis)
	e.Handle(err2)
}

type intercommService struct{}

func (s *intercommService) Upload(stream pb.Intercomm_UploadServer) error {

	return nil
}

func (s *intercommService) Download(in *pb.DownloadRequest, stream pb.Intercomm_DownloadServer) error {

	return nil
}

func (s *intercommService) Delete(cont context.Context, in *pb.DeleteRequest) (*pb.DeleteReply, error) {
	reply := &pb.DeleteReply{}
	return reply, nil
}

func (s *intercommService) ListFiles(in *pb.ListFilesRequest, stream pb.Intercomm_ListFilesServer) error {

	return nil
}
