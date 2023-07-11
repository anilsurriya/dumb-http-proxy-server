package main

import (
	"log"

	"github.com/anilsurriya/dumb-http-proxy-server/proxyserver"
)

func main() {
	pserver := proxyserver.New(":8080", nil)
	if err := pserver.StartServer(); err != nil {
		log.Fatal(err)
	}
}
