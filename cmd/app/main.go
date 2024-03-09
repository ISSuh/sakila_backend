package main

import (
	"os"

	"github.com/ISSuh/msago-sample/internal/adapter/log"
	"github.com/ISSuh/msago-sample/internal/app"
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
		l.WithError(err)
		return
	}

	if err = application.Init(); err != nil {
		l.WithError(err)
		return
	}

	application.Start()
}
