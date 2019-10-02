package main

import (
	"grpc-tutorial/pkg"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func main() {
	// wails generated code
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

	// bind 'files handling' methods
	// including WailsInit
	filesHandling := pkg.NewFH()
	defer filesHandling.Close()
	app.Bind(filesHandling)
	app.Run()
}
