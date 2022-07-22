package main

import "ftgo-order/pkg/adapter/inbound/rest"

func main() {
	ginServer := rest.NewGinServer()
	rest.StartHTTPServer(ginServer)
}
