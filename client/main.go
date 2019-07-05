package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
	sc "simpleclient/api/serverCommunication"
	w "simpleclient/pkg/wailscomm"
)

func main() {
	logic := &sc.WailsStruct{}
	// Start connection with server and client side goroutines
	go logic.Monitoring()
	// Close connection with server when program ends
	defer sc.Conn.Close()

	//choose := w.NewserviceChooser()
	choose := &w.ServiceChooser{}

	js := mewn.String("./frontend/build/static/js/main.js")
	css := mewn.String("./frontend/build/static/css/main.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "simpleClient",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(choose)
	app.Bind(logic)
	app.Run()
}
