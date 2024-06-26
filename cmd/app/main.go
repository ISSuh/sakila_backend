package main

import (
	"os"

	"github.com/ISSuh/sakila_backend/internal/app"
	"github.com/ISSuh/sakila_backend/pkg/adapter/log"
)

func main() {
	l := log.NewLogger()

	args := os.Args[1:]
	if len(args) < 1 {
		l.Fatalln("need config file path")
	}

	path := args[0]
	var application *app.Application
	var err error

	if application, err = app.NewApplication(l, path); err != nil {
		l.Fatalln(err)
	}

	if err = application.Init(); err != nil {
		l.Fatalln(err)
	}

	if err = application.Start(); err != nil {
		l.Fatalln(err)
	}
}
