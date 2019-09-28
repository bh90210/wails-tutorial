package api

import (
	context "context"
	"grpc-tutorial-server/pkg"
)

func (s *intercommService) Delete(ctx context.Context, file *File) (*Reply, error) {
	pkg.DeleteFile(file.Path)
	reply := &Reply{Feedback: "succ"}

	return reply, nil
}
