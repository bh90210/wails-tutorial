package api

import (
	context "context"
)

func (s *intercommService) Delete(ctx context.Context, file *File) (*Reply, error) {
	reply := &Reply{Feedback: "succ"}
	return reply, nil
}
