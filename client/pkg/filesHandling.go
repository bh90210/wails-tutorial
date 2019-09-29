package pkg

import (
	"grpc-tutorial/api"
	"log"

	"github.com/wailsapp/wails"
)

// FH .
type FH struct {
	Runtime *wails.Runtime
	Log     *wails.CustomLogger
	helper  *api.GrpcHelper
}

// NewFH .
func NewFH() *FH {
	return &FH{}
}

// WailsInit .
func (w *FH) WailsInit(runtime *wails.Runtime) error {
	api.ConnectToServer()

	w.Log = runtime.Log.New("Init")
	w.Runtime = runtime
	w.helper = api.NewGrpcHelper()

	runtime.Window.SetColour("#fff")

	// w.Runtime.Events.On("filesDropped", func(data ...interface{}) {
	// 	// You should probably do better error checking
	// 	fmt.Printf("I received the 'filesDropped' event with the message '%s'!\n", data[0])
	// })

	return nil
}

func (w *FH) Close() {
	w.helper.Conn.Close()
}

// ListFiles .
func (w *FH) ListFiles() {
	// get the files from server
	list := w.helper.ListFiles()

	// "emit" it to front
	w.Runtime.Events.Emit("filesList", list)
	log.Print(list)
}

// // DeleteFile .
// //func (w *FH) DeleteFile(filePath string) {
// func (w *FH) DeleteFile(filePath string) string {
// 	reply := h.DeleteFile(filePath)
// 	log.Print(reply)
//  w.ListFiles()

// 	return reply
// }

// // UploadFile .
// //func (w *FH) UploadFile() {
// func (w *FH) UploadFile(filePath string, data []byte) string {
// 	log.Print(filePath)
// 	reply := h.UploadFile(filePath, data)
// 	log.Print(reply)
//  w.ListFiles()

// 	return reply
// }

// // DownloadFile .
// //func (w *FH) DownloadFile() {
// func (w *FH) DownloadFile(filePath string) string {
// 	data := h.DownloadFile(filePath)
// 	dir, err := os.Getwd()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	os.Chdir(dir + "/downloads")

// 	err = ioutil.WriteFile(dir+"/"+filePath, data, 0644)
// 	e.Handle(err)

// 	return "succ"
// }
