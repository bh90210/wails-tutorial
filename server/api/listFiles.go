package api

import (
	fmt "fmt"

	fh "mockServer/pkg"
)

func (s *intercommService) ListFiles(in *Request, stream Intercomm_ListFilesServer) error {
	files := fh.ListFiles()
	fmt.Printf("%+v \n", files)

	// list := &File{Path: "/test", Name: "test"}
	// stream.Send(list)
	// return nil
	for _, file := range files {
		info := &File{Path: "/test", Name: string(file)}
		stream.Send(info)
	}

	return nil
}

// func (fts *fileTransferService) ListFiles(_ *proto.ListRequestType, stream proto.FileTransferService_ListFilesServer) error {
// 	err := filepath.Walk(fts.root, func(p string, info os.FileInfo, err error) error {
// 		name, err := filepath.Rel(fts.root, p)
// 		if err != nil {
// 			return err
// 		}
// 		name = filepath.ToSlash(name)
// 		modTime := new(google_protobuf.Timestamp)
// 		modTime.Seconds = int64(info.ModTime().Unix())
// 		modTime.Nanos = int32(info.ModTime().UnixNano())
// 		f := &proto.ListResponseType{Name: name, Size: info.Size(), Mode: uint32(info.Mode()), ModTime: modTime}
// 		return stream.Send(f)
// 	})
// 	return err
// }
