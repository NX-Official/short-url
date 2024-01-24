package main

import (
	"embed"
	"io/fs"
	"short-url/cmd/server"
)

//go:embed web/build/*
var FrontFiles embed.FS

func init() {
	sub, err := fs.Sub(FrontFiles, "web/build")
	if err != nil {
		panic(err)
	}
	server.FrontFiles = sub
}

func main() {
	server.Init()
	server.Run()
}
