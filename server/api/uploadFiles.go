package api

import (
	fh "grpc-tutorial-server/pkg"
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

		go fh.UploadFile(file.GetPath(), file.GetData())

	}
	//return nil
}
