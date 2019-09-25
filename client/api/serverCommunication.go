package api

import (
	context "context"
	"flag"
	"grpc-tutorial/errors"
	"io"
	"log"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// Server struct to store all information relevant with the server
type Server struct {
	Connection *grpc.ClientConn
	Client     IntercommClient
}

// ConnectToServer function initialises a new connection
// with the server and stores it in a struct for further use
func ConnectToServer() {
	flag.Parse()
	var opts []grpc.DialOption
	//creds, err := credentials.NewClientTLSFromFile("cert/cert.pem", "cert/key.pem")
	//opts = append(opts, grpc.WithTransportCredentials(creds))
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(port, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	client := NewIntercommClient(conn)

	server := &Server{}
	server.Connection = conn
	server.Client = client
	log.Printf("Connected with Server")

	stream, err := client.ListFiles(context.Background(), &ListFilesRequest{List: true})
	errors.Handle(err)

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			// read done.

			return
		}
		if err != nil {
			log.Fatalf("Failed to receive a note : %v", err)
		}
		log.Printf("Got message %s, %s, %v)", in.Path, in.Name, in.Size)
	}
}

// NewServer struct to move connection to other packages
func NewServer() *Server {
	newServer := &Server{}
	return newServer
}

// ListFiles get all files from server
func (s *Server) LS() {
	stream, err := s.Client.ListFiles(context.Background(), &ListFilesRequest{List: true})
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
