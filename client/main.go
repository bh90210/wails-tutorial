package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
	sc "simpleclient/api/serverCommunication"
)

func main() {
	go sc.Monitoring()

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
	//app.Bind(backend)
	app.Run()
}
