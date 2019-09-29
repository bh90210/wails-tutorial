package pkg

import (
	"grpc-tutorial/api"
	e "grpc-tutorial/errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/wailsapp/wails"
)

// NewFH .
func NewFH() *FH {
	return &FH{}
}

// FH .
type FH struct {
	Runtime *wails.Runtime
	Log     *wails.CustomLogger
	helper  *api.GrpcHelper
}

// WailsInit .
func (w *FH) WailsInit(runtime *wails.Runtime) error {
	connClient := api.NewGrpcHelper()

	w.Log = runtime.Log.New("Init")
	w.Runtime = runtime
	w.helper = connClient

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

// DeleteFile .
func (w *FH) DeleteFile(filePath string) string {
	reply := w.helper.DeleteFile(filePath)
	log.Print(reply)
	w.ListFiles()

	return reply
}

// UploadFile .
func (w *FH) UploadFile(filePath string, data []byte) string {
	log.Print(filePath)
	reply := w.helper.UploadFile(filePath, data)
	log.Print(reply)
	w.ListFiles()

	return reply
}

// DownloadFile .
func (w *FH) DownloadFile(filePath string) string {
	data := w.helper.DownloadFile(filePath)
	dir, err := os.Getwd()
	e.Handle(err)

	os.Chdir(dir + "/downloads")

	err = ioutil.WriteFile(dir+"/"+filePath, data, 0644)
	e.Handle(err)

	return "succ"
}
