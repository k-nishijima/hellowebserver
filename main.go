package hellowebserver

import (
	"fmt"
	"net/http"
)

type HelloWebServer struct{}

type HelloWebServerConfig struct {
	Port string
}

func (srv HelloWebServer) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func (srv HelloWebServer) Startup(config HelloWebServerConfig) {
	fmt.Println("HelloWebServer startup port:" + config.Port)
	http.HandleFunc("/", srv.handler)
	http.ListenAndServe(":"+config.Port, nil)
}
