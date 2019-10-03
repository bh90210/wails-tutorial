package pkg

import (
	"grpc-tutorial/api"
	e "grpc-tutorial/errors"
	"io/ioutil"
	"os"

	"github.com/wailsapp/wails"
)

// Binding section

type FileManager interface {
	ListFiles() [][]string
	UploadFile(fileName string, data []byte) string
	DownloadFile(fileName string) []byte
	DeleteFile(fileName string) string
	Close()
}

// NewFH this function exists to help defer Close() connection
// from `main` when app exits
// and binds methods (including WailsInit) to front
func NewFH() *FilesHandling {
	return &FilesHandling{}
}

// FilesHandling frontend binding struct
type FilesHandling struct {
	runtime     *wails.Runtime
	log         *wails.CustomLogger
	fileManager FileManager
}

// Close drops connections with server
func (w *FilesHandling) Close() {
	w.fileManager.Close()
}

// WailsInit frontend binding method
func (w *FilesHandling) WailsInit(runtime *wails.Runtime) error {
	w.log = runtime.Log.New("Init")
	w.runtime = runtime
	runtime.Window.SetColour("#fff")
	// when we call NewGrpcHelper() a new connection with the server is established
	w.fileManager = api.NewGrpcHelper()

	// w.Runtime.Events.On("filesDropped", func(data ...interface{}) {
	// 	// You should probably do better error checking
	// 	fmt.Printf("I received the 'filesDropped' event with the message '%s'!\n", data[0])
	// })

	return nil
}

// ListFiles frontend binding method
func (w *FilesHandling) ListFiles() {
	serverList := w.fileManager.ListFiles()
	// "emit" it to front
	w.runtime.Events.Emit("filesList", serverList)
}

// UploadFile frontend binding method to be called when user
// presses the upload button
// window.backend.FH.UploadFile(string, []byte)
func (w *FilesHandling) UploadFile(fileName string, data []byte) string {
	reply := w.fileManager.UploadFile(fileName, data)
	w.ListFiles()

	return reply
}

// DownloadFile frontend binding method
func (w *FilesHandling) DownloadFile(fileName string) string {
	data := w.fileManager.DownloadFile(fileName)
	dir, err := os.Getwd()
	e.Handle(err)

	os.Chdir(dir + "/downloads")

	err = ioutil.WriteFile(dir+"/"+fileName, data, 0644)
	e.Handle(err)

	return "succ"
}

// DeleteFile frontend binding method
func (w *FilesHandling) DeleteFile(fileName string) string {
	reply := w.fileManager.DeleteFile(fileName)
	w.ListFiles()

	return reply
}
