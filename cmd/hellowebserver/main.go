package main

import (
	"flag"
	"strconv"

	"github.com/k-nishijima/hellowebserver"
)

const defaultPort = 3000

func main() {
	var port = flag.Int("port", defaultPort, "use")
	flag.Parse()

	var srv hellowebserver.HelloWebServer
	config := hellowebserver.HelloWebServerConfig{
		Port: strconv.Itoa(*port),
	}
	srv.Startup(config)
}
