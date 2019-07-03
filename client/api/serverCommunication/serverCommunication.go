package serverCommunication

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"sync"
	"time"

	pb "simpleclient/api/protobuf"
	e "simpleclient/errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	address = "localhost:50051"
)

var varClient pb.IntercommClient
var conn *grpc.ClientConn
var cpuPlay = make(chan struct{})
var cpuPause = make(chan struct{})
var diskPlay = make(chan struct{})
var diskPause = make(chan struct{})
var loadPlay = make(chan struct{})
var loadPause = make(chan struct{})
var memPlay = make(chan struct{})
var memPause = make(chan struct{})
var wg sync.WaitGroup
var lastCall string

// Monitoring .
func Monitoring() {
	connect()
	defer conn.Close()

	wg.Add(1)
	go cpu()
	cpuPause <- struct{}{}
	go disk()
	diskPause <- struct{}{}
	go load()
	loadPause <- struct{}{}
	go mem()
	memPause <- struct{}{}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if scanner.Scan() {
			switch scanner.Text() {
			case "1":
				pausePrevRout()
				cpuPlay <- struct{}{}
				lastCall = "1"
			case "2":
				pausePrevRout()
				diskPlay <- struct{}{}
				lastCall = "2"
			case "3":
				pausePrevRout()
				loadPlay <- struct{}{}
				lastCall = "3"
			case "4":
				pausePrevRout()
				memPlay <- struct{}{}
				lastCall = "4"
			}
		}
	}
}

func pausePrevRout() {
	if lastCall == "" {
		//
	} else {
		if lastCall == "1" {
			cpuPause <- struct{}{}
		}
		if lastCall == "2" {
			diskPause <- struct{}{}
		}
		if lastCall == "3" {
			loadPause <- struct{}{}
		}
		if lastCall == "4" {
			memPause <- struct{}{}
		}
	}
}

func connect() {
	// Create the client TLS credentials
	creds, err := credentials.NewClientTLSFromFile("cert/cert.pem", "")
	e.Handle(err)
	// Initiate a connection with the server
	conn, err = grpc.Dial(address, grpc.WithTransportCredentials(creds))
	e.Handle(err)
	client := pb.NewIntercommClient(conn)
	varClient = client

	// do logging stuff

}

func cpu() {
	for {
		select {
		case <-cpuPause:
			log.Println("cpu pause")
			select {
			case <-cpuPlay:
				log.Println("cpu play")
				/*case <-quit:
				wg.Done()
				return*/
			}
		/*case <-quit:
		wg.Done()
		return*/
		default:
			cpuCallnReturn()
		}
	}
}

func cpuCallnReturn() {
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

func disk() {
	for {
		select {
		case <-diskPause:
			log.Println("disk pause")
			select {
			case <-diskPlay:
				log.Println("disk play")
			}
		default:
			diskCallnReturn()
		}
	}
}

func diskCallnReturn() {
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

func load() {
	for {
		select {
		case <-loadPause:
			log.Println("load pause")
			select {
			case <-loadPlay:
				log.Println("load play")
			}
		default:
			loadCallnReturn()
		}
	}
}

func loadCallnReturn() {
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

func mem() {
	for {
		select {
		case <-memPause:
			log.Println("mem pause")
			select {
			case <-memPlay:
				log.Println("mem play")
			}
		default:
			memCallnReturn()
		}
	}
}

func memCallnReturn() {
	done := make(chan bool)

	stream, err := varClient.GetMemStats(context.Background(), &pb.MemStatsRequest{Name: "mem"})
	e.Handle(err)

	response, err := stream.Recv()
	if err == io.EOF {
		close(done)
		return
	}
	e.Handle(err)

	log.Println(response.Total)
	log.Println(response.Used)
	log.Println(response.Free)
	log.Println(response.Sin)
	log.Println(response.Sout)

	time.Sleep(1 * time.Second)
}
