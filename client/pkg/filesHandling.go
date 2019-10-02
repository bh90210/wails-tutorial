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
	*FH
}
type upload struct {
	fileName string
	data     []byte
	*FH
}
type download struct {
	fileName string
	*FH
}
type delete struct {
	fileName string
	*FH
}

func (i *list) handle() string {
	serverList := i.FH.helper.ListFiles()
	// "emit" it to front
	i.FH.runtime.Events.Emit("filesList", serverList)

	return "succ"
}

func (i *upload) handle() string {
	reply := i.FH.helper.UploadFile(i.fileName, i.data)
	i.FH.ListFiles()
	return reply
}

func (i *download) handle() string {
	data := i.FH.helper.DownloadFile(i.fileName)
	dir, err := os.Getwd()
	e.Handle(err)

	os.Chdir(dir + "/downloads")

	err = ioutil.WriteFile(dir+"/"+i.fileName, data, 0644)
	e.Handle(err)

	return "succ"
}

func (i *delete) handle() string {
	reply := i.FH.helper.DeleteFile(i.fileName)
	i.FH.ListFiles()

	return reply
}

func handleFiles(f files) string {
	reply := f.handle()

	return reply
}

// NewFH this function exists to help defer Close() connection
// from `main` when app exits
// and binds methods (including WailsInit) to front
func NewFH() *FH {
	return &FH{}
}

// FH frontend binding struct
type FH struct {
	runtime *wails.Runtime
	log     *wails.CustomLogger
	// *api.GrpcHelper holds gRPC's connection and client for use in this package
	helper *api.GrpcHelper
}

// Close drops connections with server
func (w *FH) Close() {
	w.helper.Conn.Close()
}

// WailsInit frontend binding method
func (w *FH) WailsInit(runtime *wails.Runtime) error {
	w.log = runtime.Log.New("Init")
	w.runtime = runtime
	runtime.Window.SetColour("#fff")
	// when we call NewGrpcHelper() a new connection with the server is established
	connClient := api.NewGrpcHelper()
	w.helper = connClient
	// w.Runtime.Events.On("filesDropped", func(data ...interface{}) {
	// 	// You should probably do better error checking
	// 	fmt.Printf("I received the 'filesDropped' event with the message '%s'!\n", data[0])
	// })

	return nil
}

// ListFiles frontend binding method
func (w *FH) ListFiles() {
	// get the files from server
	_ = handleFiles(&list{FH: w})
}

// UploadFile frontend binding method to be called when user
// presses the upload button
// window.backend.FH.UploadFile(string, []byte)
func (w *FH) UploadFile(fileName string, data []byte) string {
	uploadFile := &upload{fileName: fileName, data: data, FH: w}
	reply := handleFiles(uploadFile)

	return reply
}

// DownloadFile frontend binding method
func (w *FH) DownloadFile(fileName string) string {
	downloadFile := &download{fileName: fileName, FH: w}
	reply := handleFiles(downloadFile)

	return reply
}

// DeleteFile frontend binding method
func (w *FH) DeleteFile(fileName string) string {
	deleteFile := &delete{fileName: fileName, FH: w}
	reply := handleFiles(deleteFile)

	return reply
}
