package main

import (
	"github.com/gin-gonic/gin"
	"github.com/peileiscott/gorder/common/config"
	"github.com/peileiscott/gorder/common/server"
	"github.com/peileiscott/gorder/order/ports"
	"github.com/spf13/viper"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		panic(err)
	}
}

func main() {
	serviceName := viper.GetString("order.service_name")
	server.RunHTTPServer(serviceName, func(router *gin.Engine) {
		ports.RegisterHandlersWithOptions(router, ports.NewHTTPServer(), ports.GinServerOptions{
			BaseURL:      "/api",
			Middlewares:  nil,
			ErrorHandler: nil,
		})
	})
}
