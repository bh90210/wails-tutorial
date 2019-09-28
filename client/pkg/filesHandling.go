package pkg

import (
	"grpc-tutorial/api"
	e "grpc-tutorial/errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/wailsapp/wails"
)

// FH .
type FH struct {
	Runtime *wails.Runtime
	Log     *wails.CustomLogger
}

// NewFH .
func NewFH() *FH {
	return &FH{}
}

// WailsInit .
func (w *FH) WailsInit(runtime *wails.Runtime) error {
	w.Log = runtime.Log.New("Init")
	w.Runtime = runtime

	// get a list of files stored on server
	w.ListFiles()

	runtime.Window.SetColour("#fff")

	// w.Runtime.Events.On("filesDropped", func(data ...interface{}) {
	// 	// You should probably do better error checking
	// 	fmt.Printf("I received the 'filesDropped' event with the message '%s'!\n", data[0])
	// })

	return nil
}

// ListFiles .
func (w *FH) ListFiles() {
	// get the files from server
	list := api.ListFiles()

	// "emit" it to front
	w.Runtime.Events.Emit("filesList", list)
	log.Print(list)
}

// DeleteFile .
//func (w *FH) DeleteFile(filePath string) {
func DeleteFile(filePath string) string {
	reply := api.DeleteFile(filePath)
	log.Print(reply)

	return reply
}

// UploadFile .
//func (w *FH) UploadFile() {
func UploadFile(list []string, data [][]byte) string {
	reply := api.UploadFile(list, data)
	log.Print("test")

	return reply
}

// DownloadFile .
//func (w *FH) DownloadFile() {
func DownloadFile(filePath string) string {
	data := api.DownloadFile(filePath)
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	os.Chdir(dir + "/downloads")

	err = ioutil.WriteFile(dir+"/"+filePath, data, 0644)
	e.Handle(err)

	return "succ"
}
