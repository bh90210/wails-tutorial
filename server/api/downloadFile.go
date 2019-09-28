package api

import (
	context "context"
	fh "grpc-tutorial-server/pkg"
)

func (s *intercommService) Download(ctx context.Context, in *Request) (*File, error) {
	f := fh.DownloadFile(in.GetPath())
	file := &File{Data: f}

	return file, nil
}
