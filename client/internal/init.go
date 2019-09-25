package internal

import (
	"github.com/wailsapp/wails"
)

type FH struct {
	Runtime *wails.Runtime
	Log     *wails.CustomLogger
}

func NewFH() *FH {
	result := &FH{}
	return result
}

// WailsInit .
func (w *FH) WailsInit(runtime *wails.Runtime) error {
	w.Log = runtime.Log.New("Init")
	w.Runtime = runtime

	// get a list of files stored on server
	//go api.ListFiles()
	w.ListFiles()

	runtime.Window.SetColour("#fff")
	
	return nil
}
