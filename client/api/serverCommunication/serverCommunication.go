package serverCommunication

import (
	"bufio"
	"context"
	_ "fmt"
	"io"
	"log"
	"os"
	"time"

	pb "simpleClient/api/protobuf"
	e "simpleClient/errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	address = "localhost:50051"
)

var varClient pb.IntercommClient
var conn *grpc.ClientConn

func connect() {
	// Create the client TLS credentials
	creds, err := credentials.NewClientTLSFromFile("cert/cert.pem", "")
	e.Handle(err)
	// Initiate a connection with the server
	conn, err = grpc.Dial(address, grpc.WithTransportCredentials(creds))
	e.Handle(err)
	client := pb.NewIntercommClient(conn)
	varClient = client

	// do logging

}

func cpu(cpuChannel chan bool) {
	for {
		select {
		case <-cpuChannel:
			return
		default:

			// Do other stuff
			done := make(chan bool)

			stream, err := varClient.GetCPUStats(context.Background(), &pb.CPUStatsRequest{Name: "cpu"})
			e.Handle(err)

			response, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			e.Handle(err)
			log.Println(response.Percentage)
			log.Println(response.User)
			log.Println(response.System)
			log.Println(response.Idle)
			log.Println(response.Nice)

			time.Sleep(1 * time.Second)
		}
	}
}

func disk(diskChannel chan bool) {
	for {
		select {
		case <-diskChannel:
			return
		default:

			// Do other stuff
			done := make(chan bool)

			stream, err := varClient.GetDiskStats(context.Background(), &pb.DiskStatsRequest{Name: "disk"})
			e.Handle(err)

			response, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			e.Handle(err)
			log.Println(response.UsedPercent)
			log.Println(response.InodesUsedPercent)

			time.Sleep(1 * time.Second)
		}
	}
}

func load(loadChannel chan bool) {
	for {
		select {
		case <-loadChannel:
			return
		default:

			// Do other stuff
			done := make(chan bool)

			stream, err := varClient.GetLoadStats(context.Background(), &pb.LoadStatsRequest{Name: "load"})
			e.Handle(err)

			response, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			e.Handle(err)

			log.Println(response.Load1)
			log.Println(response.Load5)
			log.Println(response.Load15)
			log.Println(response.ProcsRunning)
			log.Println(response.ProcsBlocked)
			log.Println(response.Ctxt)

			time.Sleep(1 * time.Second)
		}
	}
}

// Monitoring .
func Monitoring() {
	connect()
	defer conn.Close()

	cpuChannel := make(chan bool, 1)
	diskChannel := make(chan bool, 1)
	loadChannel := make(chan bool, 1)
	//memChannel := make(chan bool, 1)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if scanner.Scan() {
			switch scanner.Text() {
			case "1":
				go cpu(cpuChannel)
				diskChannel <- true
				loadChannel <- true
			case "2":
				go disk(diskChannel)
				cpuChannel <- true
				loadChannel <- true
			case "3":
				go load(loadChannel)
				cpuChannel <- true
				diskChannel <- true
			}
		}
	}
}
