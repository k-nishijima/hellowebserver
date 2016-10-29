package hellowebserver

import (
	"log"

	"fmt"
	"net/http"

	"github.com/comail/colog"
)

type HelloWebServer struct{}

type HelloWebServerConfig struct {
	Port string
}

func init() {
	colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()
}

func (srv HelloWebServer) handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("debug: handler requset:%v", r.RequestURI)
	fmt.Fprintf(w, "Hello, World")
}

func (srv HelloWebServer) Startup(config HelloWebServerConfig) {
	log.Printf("info: HelloWebServer startup... port:%v", config.Port)
	http.HandleFunc("/", srv.handler)
	http.ListenAndServe(":"+config.Port, nil)
}
