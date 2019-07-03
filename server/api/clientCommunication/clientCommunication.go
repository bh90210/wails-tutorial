package clientCommunication

import (
	"log"
	"net"
	"time"

	pb "mockServer/api/protobuf"
	e "mockServer/errors"
	"mockServer/pkg/monitoring"

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
	// attach the Ping service to the server
	pb.RegisterIntercommServer(s, &server{})
	// start the server
	err2 := s.Serve(lis)
	e.Handle(err2)
}

type server struct{}

// GetCPUStats implements helloworld.GreeterServer
func (s *server) GetCPUStats(in *pb.CPUStatsRequest, stream pb.Intercomm_GetCPUStatsServer) error {

	percent, user, system, idle, nice := monitoring.GetCPU()
	data := pb.CPUStatsReply{Percentage: percent, User: user, System: system, Idle: idle, Nice: nice}
	stream.Send(&data)

	log.Println("new request")
	time.Sleep(1 * time.Second)
	return nil
}

// GetDiskStats .
func (s *server) GetDiskStats(in *pb.DiskStatsRequest, stream pb.Intercomm_GetDiskStatsServer) error {

	usedPercent, inodesUsedPercent := monitoring.GetDISKCounters()
	data := pb.DiskStatsReply{UsedPercent: usedPercent, InodesUsedPercent: inodesUsedPercent}
	stream.Send(&data)

	return nil
}

// GetLoadStats .
func (s *server) GetLoadStats(in *pb.LoadStatsRequest, stream pb.Intercomm_GetLoadStatsServer) error {

	load1, load5, load15, procsRunning, procsBlocked, ctxt := monitoring.GetLoad()
	data := pb.LoadStatsReply{Load1: load1, Load5: load5, Load15: load15, ProcsRunning: procsRunning, ProcsBlocked: procsBlocked, Ctxt: ctxt}
	stream.Send(&data)

	return nil
}

// GetMemStats .
func (s *server) GetMemStats(in *pb.MemStatsRequest, stream pb.Intercomm_GetMemStatsServer) error {

	total, used, free, sin, sout := monitoring.GetMem()
	data := pb.MemStatsReply{Total: total, Used: used, Free: free, Sin: sin, Sout: sout}
	stream.Send(&data)

	return nil
}
