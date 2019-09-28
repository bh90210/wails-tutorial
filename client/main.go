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

	// frontend binding section

	// bind 'files handling' methods
	// including WailsInit
	fh := pkg.NewFH()
	app.Bind(fh)
	//app.Bind(fh.ListFiles)
	app.Bind(pkg.DeleteFile)
	//app.Bind(listFiles)
	app.Bind(pkg.UploadFile)
	app.Bind(pkg.DownloadFile)
	app.Run()
}
