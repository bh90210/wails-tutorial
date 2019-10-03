package api

import (
	context "context"
	e "grpc-tutorial/errors"
	"io"
	"strconv"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// NewGrpcHelper initiates a connection with server
func NewGrpcHelper() *GrpcHelper {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(port, opts...)
	e.Handle(err)

	client := NewIntercommClient(conn)

	return &GrpcHelper{Conn: conn, Client: client}
}

// GrpcHelper helper struct to hold grpc's connection and client
type GrpcHelper struct {
	Conn   *grpc.ClientConn
	Client IntercommClient
}

// ListFiles get all files from server
func (h *GrpcHelper) ListFiles() [][]string {
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

// UploadFile a file from server
func (h *GrpcHelper) UploadFile(fileName string, dataList []byte) string {
	reply, err := h.Client.Upload(context.Background(), &File{Path: fileName, Data: dataList})
	e.Handle(err)

	return reply.Feedback
}

// DownloadFile a file from server
func (h *GrpcHelper) DownloadFile(fileName string) []byte {
	file, err := h.Client.Download(context.Background(), &Request{Path: fileName})
	e.Handle(err)

	data := file.GetData()

	return data
}

// DeleteFile a file from server
func (h *GrpcHelper) DeleteFile(fileName string) string {
	reply, err := h.Client.Delete(context.Background(), &File{Path: fileName})
	e.Handle(err)

	if reply.Feedback == "succ" {
		return "succ"
	}

	return "err"
}
