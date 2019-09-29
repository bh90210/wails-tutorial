package api

import (
	context "context"
	"flag"
	e "grpc-tutorial/errors"
	"io"
	"log"
	"strconv"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func NewGrpcHelper() *GrpcHelper {
	return &GrpcHelper{Conn: connHelper, Client: clientHelper}
}

type GrpcHelper struct {
	Conn   *grpc.ClientConn
	Client IntercommClient
}

var connHelper *grpc.ClientConn
var clientHelper IntercommClient

// ConnectToServer function initialises a new connection
// with the server and pass it down to calling function
func ConnectToServer() *grpc.ClientConn {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	co, err := grpc.Dial(port, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	cl := NewIntercommClient(co)

	//_ = &GrpcHelper{Conn: co, Client: cl}

	clientHelper = cl
	connHelper = co

	return connHelper
}

// UploadFile a file from server
func (h *GrpcHelper) UploadFile(filePath string, dataList []byte) string {
	// connect to server
	//conn, client := ConnectToServer()
	// and drop connection when done
	//defer conn.Close()

	stream, err := h.Client.Upload(context.Background())
	e.Handle(err)

	file := &File{Path: filePath, Data: dataList}
	err = stream.Send(file)
	e.Handle(err)

	reply, err := stream.CloseAndRecv()
	e.Handle(err)

	return reply.Feedback
}

// ListFiles get all files from server
func (h *GrpcHelper) ListFiles() [][]string {
	// connect to server
	//conn, client := ConnectToServer()
	// and drop connection when done
	//defer conn.Close()

	stream, err := h.Client.ListFiles(context.Background(), &Request{List: true})
	e.Handle(err)

	var filesList [][]string
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			e.Handle(err)

			log.Printf("++++++++ Got message %s, %s, %v", in.Path, in.Name, in.Size)
			// create an array to place the files
			intToString := strconv.FormatInt(int64(in.Size), 10)
			var entry []string
			entry = append(entry, in.Path)
			entry = append(entry, in.Name)
			entry = append(entry, intToString)
			filesList = append(filesList, entry)
		}
	}()
	stream.CloseSend()
	<-waitc

	return filesList
}

// DownloadFile a file from server
func (h *GrpcHelper) DownloadFile(filePath string) []byte {
	// connect to server
	//conn, client := ConnectToServer()
	// and drop connection when done
	//defer conn.Close()

	file, err := h.Client.Download(context.Background(), &Request{Path: filePath})
	e.Handle(err)

	data := file.GetData()

	return data
}

// DeleteFile a file from server
func (h *GrpcHelper) DeleteFile(filePath string) string {
	// connect to server
	//conn, client := ConnectToServer()
	// and drop connection when done
	//defer conn.Close()

	reply, err := h.Client.Delete(context.Background(), &File{Path: filePath})
	e.Handle(err)

	if reply.Feedback == "succ" {
		return "succ"
	}

	return "err"
}
