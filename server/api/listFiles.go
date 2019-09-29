package api

import (
	fh "grpc-tutorial-server/pkg"
)

func (s *intercommService) ListFiles(in *Request, stream Intercomm_ListFilesServer) error {
	// use a different package function
	// to collect all files
	files := fh.ListFiles()
	// iterate over the result to get the info
	for name, value := range files {
		// then iterate again to get the nested map values
		for path, size := range value {
			info := &File{Path: path, Name: name, Size: int32(size)}
			// stream each file's info to client
			stream.Send(info)
		}
	}

	return nil
}
