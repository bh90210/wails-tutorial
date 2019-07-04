package serverCommunication

import (
	//"bufio"
	"context"
	"io"
	"log"
	//"os"
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
var Conn *grpc.ClientConn
var CpuPlay = make(chan struct{})
var CpuPause = make(chan struct{})
var DiskPlay = make(chan struct{})
var DiskPause = make(chan struct{})
var LoadPlay = make(chan struct{})
var LoadPause = make(chan struct{})
var MemPlay = make(chan struct{})
var MemPause = make(chan struct{})
var wg sync.WaitGroup
var LastCall string

// Monitoring .
func Monitoring() {
	connect()
	//defer Conn.Close()

	wg.Add(1)
	go cpu()
	CpuPause <- struct{}{}
	go disk()
	DiskPause <- struct{}{}
	go load()
	LoadPause <- struct{}{}
	go mem()
	MemPause <- struct{}{}
}

// PausePrevRout .
func PausePrevRout() {
	if LastCall == "" {
		//
	} else {
		if LastCall == "1" {
			CpuPause <- struct{}{}
		}
		if LastCall == "2" {
			DiskPause <- struct{}{}
		}
		if LastCall == "3" {
			LoadPause <- struct{}{}
		}
		if LastCall == "4" {
			MemPause <- struct{}{}
		}
	}
}

func connect() {
	// Create the client TLS credentials
	creds, err := credentials.NewClientTLSFromFile("cert/cert.pem", "")
	e.Handle(err)
	// Initiate a connection with the server
	Conn, err = grpc.Dial(address, grpc.WithTransportCredentials(creds))
	e.Handle(err)
	client := pb.NewIntercommClient(Conn)
	varClient = client

	// do logging stuff

}

func cpu() {
	for {
		select {
		case <-CpuPause:
			log.Println("cpu pause")
			select {
			case <-CpuPlay:
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
		case <-DiskPause:
			log.Println("disk pause")
			select {
			case <-DiskPlay:
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
		case <-LoadPause:
			log.Println("load pause")
			select {
			case <-LoadPlay:
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
		case <-MemPause:
			log.Println("mem pause")
			select {
			case <-MemPlay:
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
