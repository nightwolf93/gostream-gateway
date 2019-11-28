package main

import (
	"github.com/nightwolf93/gostream-server/gateway"
)

func main() {
	gateway := gateway.NewGateway(gateway.GatewayConfig{
		Port:     1935,
		Password: "test",
	})
	gateway.Listen()
}
