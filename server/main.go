package main

import (
	cc "mockServer/api"
	fh "mockServer/pkg"
)

func main() {
	go cc.StartServer()

	fh.TestFunc()
}
