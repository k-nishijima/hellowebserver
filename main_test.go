package hellowebserver

import (
	"testing"
)

func TestStartup(t *testing.T) {
	var srv HelloWebServer

	config := HelloWebServerConfig{
		Port: "1234",
	}
	srv.Startup(config)
}
