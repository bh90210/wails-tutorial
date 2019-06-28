package main

import (
	"context"
	"io"
	"log"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"

	pb "client/api/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func basic() string {
	return "World!"
}

func main() {
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

	stream, err := client.GetCPUStats(context.Background(), &pb.CPUStatsRequest{Name: "cpu"})
	if err != nil {
		log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
	}

	var perc int32
	ctx := stream.Context()
	done := make(chan bool)

	go func() {
		for {
			response, err := *stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			perc = response.Percentage
			log.Printf("new max %d received", perc)
			log.Println("test")
			log.Println(response.Percentage)
		}
	}()

	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		close(done)
	}()
	<-done
	log.Printf("finished with max=%d", perc)

	// wails
	js := mewn.String("./frontend/build/static/js/main.js")
	css := mewn.String("./frontend/build/static/css/main.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "client",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(basic)
	app.Run()
}
