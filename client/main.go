package main

import (
	"fmt"
	"grpc-tutorial/api"
	"time"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

// Stats .
type Stats struct {
	log     *wails.CustomLogger
	runtime *wails.Runtime
}

// NewTodos attempts to create a new Todo list
func NewTodos() (*Stats, error) {
	// Create new Todos instance
	result := &Stats{}
	// Return it
	return result, nil
}

// WailsInit .
func (s *Stats) WailsInit(runtime *wails.Runtime) error {
	s.log = runtime.Log.New("Stats")

	runtime.Window.SetColour("#fff")

	//server := api.NewServer()
	//go server.LS()

	go func() {
		for {
			runtime.Events.Emit("cpu_usage", [4]int{1, 2, 3, 4})
			time.Sleep(1 * time.Second)
		}
	}()

	runtime.Events.On("filesDropped", func(data ...interface{}) {
		// You should probably do better error checking
		fmt.Printf("I received the 'filesDropped' event with the message '%s'!\n", data[0])
	})

	return nil
}

func basic(name string, path string, size float64, file []byte) {
	//fmt.Printf(fileName)
	fmt.Println(name)
	fmt.Println(path)
	fmt.Println(size)
	fmt.Printf("basic bind with the message %s", file)
	//fmt.Println(fileSize)
	//return "works! :+1:"
}

func main() {
	// initiate a server connection
	// connect to server
	go api.ConnectToServer()
	// bring in a new server struct
	server := api.NewServer()
	// and defer it so when app closes
	// server connection closes too
	defer server.Connection.Close()

	//go server.ListFiles()

	// wails generated code
	// do not alter
	js := mewn.String("./frontend/build/static/js/main.js")
	css := mewn.String("./frontend/build/static/css/main.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  800,
		Height: 650,
		Title:  "gRPC Tutorial",
		JS:     js,
		CSS:    css,
		Colour: "#FFF",
	})

	// frontend binding section
	biindd, _ := NewTodos()
	app.Bind(biindd)
	app.Bind(basic)
	//app.Bind(server)
	app.Run()
}
