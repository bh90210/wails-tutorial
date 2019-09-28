package api

import (
	fh "grpc-tutorial-server/internal"
	"io"
)

func (s *intercommService) Upload(stream Intercomm_UploadServer) error {

	for {
		file, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&Reply{Feedback: "succ"})
		}
		if err != nil {
			return err
		}

		path := file.GetPath()
		data := file.GetData()
		go fh.UploadFile(path, data)

	}
}
