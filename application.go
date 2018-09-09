package main

import (
	"log"
	"net/http"

	"github.com/nabinkh/go-backend/handler"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	overrideLogFile("backendlog.log")
	log.Println("Startin backend ...")
	mux := http.NewServeMux()

	mux.HandleFunc("/search", handler.RequestHandler)
	log.Println(http.ListenAndServe(":9000", mux))
}

func overrideLogFile(file string) {
	l := &lumberjack.Logger{
		Filename:   file,
		MaxSize:    500, //megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
	}
	log.SetOutput(l)
}
