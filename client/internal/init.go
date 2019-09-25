package internal

import (
	"github.com/wailsapp/wails"
)

type FH struct {
	Runtime *wails.Runtime
	Log     *wails.CustomLogger
}

func NewFH() (*FH, error) {
	result := &FH{}
	return result, nil
}

// WailsInit .
func (w *FH) WailsInit(runtime *wails.Runtime) error {
	w.Log = runtime.Log.New("Init")
	runtime.Window.SetColour("#fff")

	w.Runtime = runtime

	// get a list of files stored on server
	//go api.ListFiles()
	w.ListFiles()

	return nil
}
