package hellowebserver

import "testing"

func TestStartup(t *testing.T) {
	var srv HelloWebServer

	srv.Startup()
}
