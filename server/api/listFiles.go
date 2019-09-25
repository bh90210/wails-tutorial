package api

import (
	fh "grpc-tutorial-server/pkg"
)

func (s *intercommService) ListFiles(in *Request, stream Intercomm_ListFilesServer) error {
	files := fh.ListFiles()
	for key, value := range files {
		for path, size := range value {
			info := &File{Path: path, Name: key, Size: int32(size)}
			stream.Send(info)
		}
	}

	return nil
}
