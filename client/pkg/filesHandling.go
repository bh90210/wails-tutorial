package pkg

import (
	"grpc-tutorial/api"
	e "grpc-tutorial/errors"
	"io/ioutil"
	"os"

	"github.com/wailsapp/wails"
)

type files interface {
	handle() string
}

type list struct {
	list [][]string
	// holds gRPC's connection, client
	// and wails' runtime, logger
	*FilesHandling
}
type upload struct {
	fileName string
	data     []byte
	*FilesHandling
}
type download struct {
	fileName string
	*FilesHandling
}
type delete struct {
	fileName string
	*FilesHandling
}

func (i *list) handle() string {
	serverList := i.FilesHandling.apiHelper.ListFiles()
	// "emit" it to front
	i.FilesHandling.runtime.Events.Emit("filesList", serverList)

	return "succ"
}

func (i *upload) handle() string {
	reply := i.FilesHandling.apiHelper.UploadFile(i.fileName, i.data)
	i.FilesHandling.ListFiles()
	return reply
}

func (i *download) handle() string {
	data := i.FilesHandling.apiHelper.DownloadFile(i.fileName)
	dir, err := os.Getwd()
	e.Handle(err)

	os.Chdir(dir + "/downloads")

	err = ioutil.WriteFile(dir+"/"+i.fileName, data, 0644)
	e.Handle(err)

	return "succ"
}

func (i *delete) handle() string {
	reply := i.FilesHandling.apiHelper.DeleteFile(i.fileName)
	i.FilesHandling.ListFiles()

	return reply
}

func handleFiles(f files) string {
	reply := f.handle()

	return reply
}

// NewFH this function exists to help defer Close() connection
// from `main` when app exits
// and binds methods (including WailsInit) to front
func NewFH() *FilesHandling {
	return &FilesHandling{}
}

// FH frontend binding struct
type FilesHandling struct {
	runtime *wails.Runtime
	log     *wails.CustomLogger
	// *api.GrpcHelper holds gRPC's connection and client for use in this package
	apiHelper *api.GrpcHelper
}

// Close drops connections with server
func (w *FilesHandling) Close() {
	w.apiHelper.Conn.Close()
}

// WailsInit frontend binding method
func (w *FilesHandling) WailsInit(runtime *wails.Runtime) error {
	w.log = runtime.Log.New("Init")
	w.runtime = runtime
	runtime.Window.SetColour("#fff")
	// when we call NewGrpcHelper() a new connection with the server is established
	connClient := api.NewGrpcHelper()
	w.apiHelper = connClient
	// w.Runtime.Events.On("filesDropped", func(data ...interface{}) {
	// 	// You should probably do better error checking
	// 	fmt.Printf("I received the 'filesDropped' event with the message '%s'!\n", data[0])
	// })

	return nil
}

// ListFiles frontend binding method
func (w *FilesHandling) ListFiles() {
	// get the files from server
	_ = handleFiles(&list{FilesHandling: w})
}

// UploadFile frontend binding method to be called when user
// presses the upload button
// window.backend.FH.UploadFile(string, []byte)
func (w *FilesHandling) UploadFile(fileName string, data []byte) string {
	uploadFile := &upload{fileName: fileName, data: data, FilesHandling: w}
	reply := handleFiles(uploadFile)

	return reply
}

// DownloadFile frontend binding method
func (w *FilesHandling) DownloadFile(fileName string) string {
	downloadFile := &download{fileName: fileName, FilesHandling: w}
	reply := handleFiles(downloadFile)

	return reply
}

// DeleteFile frontend binding method
func (w *FilesHandling) DeleteFile(fileName string) string {
	deleteFile := &delete{fileName: fileName, FilesHandling: w}
	reply := handleFiles(deleteFile)

	return reply
}
