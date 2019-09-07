package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

// Stats .
type Stats struct {
	log *wails.CustomLogger
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

	go func() {
		for {
			runtime.Events.Emit("cpu_usage", rand.Intn(100))
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

	js := mewn.String("./frontend/build/static/js/main.js")
	css := mewn.String("./frontend/build/static/css/main.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "gRPC Tutorial",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})

	biindd, _ := NewTodos()

	app.Bind(biindd)
	app.Bind(basic)
	app.Run()
}
