package api

import (
	context "context"
)

func (s *intercommService) Download(ctx context.Context, in *Request) (*File, error) {
	file := &File{Name: "test", Path: "path", Size: 666}
	return file, nil
}
