package pkg

import "grpc-tutorial/api"

// DeleteFile .
func (w *FH) DeleteFile(filePath string) {
	reply := api.DeleteFile(filePath)

	if reply == "succ" {
		w.ListFiles()
	}
}
