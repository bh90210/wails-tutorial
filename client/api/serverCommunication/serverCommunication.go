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

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

var input string
var loop1 int
var loop2 int
var loop3 int
var loop4 int
var screen string

// Monitoring .
func Monitoring() {
	var conn *grpc.ClientConn
	// Create the client TLS credentials
	creds, err := credentials.NewClientTLSFromFile("cert/cert.pem", "")
	if err != nil {
		log.Fatalf("could not load tls cert: %s", err)
	}
	// Initiate a connection with the server
	conn, err = grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	client := pb.NewIntercommClient(conn)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if scanner.Scan() {
			if scanner.Text() == "1" {
				loop1 = 1
				loop2 = 0
				loop3 = 0
				loop4 = 0
			}
			if scanner.Text() == "2" {
				loop1 = 0
				loop2 = 1
				loop3 = 0
				loop4 = 0
			}
			if scanner.Text() == "3" {
				loop1 = 0
				loop2 = 0
				loop3 = 1
				loop4 = 0
			}
			if scanner.Text() == "4" {
				loop1 = 0
				loop2 = 0
				loop3 = 0
				loop4 = 1
			}

			switch scanner.Text() {
			case "1":
				//
				for {
					if scanner.Text() != "1" {
						break
					}

					done := make(chan bool)

					stream, err := client.GetCPUStats(context.Background(), &pb.CPUStatsRequest{Name: "cpu"})
					if err != nil {
						log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
					}

					response, err := stream.Recv()
					if err == io.EOF {
						close(done)
						return
					}
					if err != nil {
						log.Fatalf("can not receive %v", err)
					}
					log.Println(response.Percentage)
					log.Println(response.User)
					log.Println(response.System)
					log.Println(response.Idle)
					log.Println(response.Nice)

					time.Sleep(1 * time.Second)
				}
			case "2":
				//
				for {
					if loop2 == 0 {
						break
					}

					log.Println("a")
					time.Sleep(1 * time.Second)
				}
			case "3":
				//
				for {
					if loop3 == 0 {
						break
					}

					time.Sleep(1 * time.Second)
				}
			case "4":
				//
				for {
					if loop4 == 0 {
						break
					}

					time.Sleep(1 * time.Second)
				}
			}
		}
	}
}
