package api

import (
	context "context"
	fh "grpc-tutorial-server/pkg"
)

func (s *intercommService) Delete(ctx context.Context, file *File) (*Reply, error) {
	fh.DeleteFile(file.Path)
	reply := &Reply{Feedback: "succ"}

	return reply, nil
}
