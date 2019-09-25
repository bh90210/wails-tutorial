package api

import (
	"flag"
	"log"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// ConnectToServer function initialises a new connection
// with the server and pass it down to calling function
func ConnectToServer() (*grpc.ClientConn, IntercommClient) {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(port, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	client := NewIntercommClient(conn)

	return conn, client
}
