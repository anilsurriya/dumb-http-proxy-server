package proxyserver

import (
	"log"
	"net/http"
)

type ProxyServer struct {
	dialer *http.Client
	addr   string
	logger *log.Logger
}

func New(listeraddr string, dialer *http.Client) *ProxyServer {
	if dialer == nil {
		dialer = http.DefaultClient
	}
	return &ProxyServer{
		dialer: dialer,
		addr:   listeraddr,
		logger: log.Default(),
	}
}
