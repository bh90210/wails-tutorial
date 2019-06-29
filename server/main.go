package main

import (
	"log"
	"net"
	"time"

	pb "mockServer/api/protobuf"
	"mockServer/pkg/monitoring"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	port = ":50051"
)

type server struct{}

// GetCPUStats implements helloworld.GreeterServer
func (s *server) GetCPUStats(in *pb.CPUStatsRequest, stream pb.Intercomm_GetCPUStatsServer) error {

	percent, user, system, idle, nice := monitoring.GetCPU()
	data := pb.CPUStatsReply{Percentage: percent, User: user, System: system, Idle: idle, Nice: nice}
	stream.Send(&data)
	//runtime.Events.Emit("cpu_usage", s.GetCPUUsage())

	log.Println("response.Percentage")
	time.Sleep(1 * time.Second)
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create the TLS credentials
	creds, err := credentials.NewServerTLSFromFile("cert/cert.pem", "cert/key.pem")
	if err != nil {
		log.Fatalf("could not load TLS keys: %s", err)
	}

	// Create an array of gRPC options with the credentials
	opts := []grpc.ServerOption{grpc.Creds(creds)}
	// create a gRPC server object
	s := grpc.NewServer(opts...)
	// attach the Ping service to the server
	pb.RegisterIntercommServer(s, &server{})
	// start the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
