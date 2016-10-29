package hellowebserver

import (
	"fmt"
	"net/http"
)

type HelloWebServer struct{}

func (srv HelloWebServer) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func (srv HelloWebServer) Startup() {
	http.HandleFunc("/", srv.handler) // ハンドラを登録してウェブページを表示させる
	http.ListenAndServe(":8080", nil)
}
