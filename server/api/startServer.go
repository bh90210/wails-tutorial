package api

import (
	"flag"
	fmt "fmt"
	"net"

	e "mockServer/errors"

	"google.golang.org/grpc"
)

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

type intercommService struct{}
