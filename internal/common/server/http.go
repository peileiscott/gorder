package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func RunHTTPServer(serviceName string, registerHandlers func(*gin.Engine)) {
	addr := viper.Sub(serviceName).GetString("http_addr")
	if addr == "" {
		logrus.Panicf("please provide http address for %s service", serviceName)
	}
	RunHTTPServerOnAddr(addr, registerHandlers)
}

func RunHTTPServerOnAddr(addr string, registerHandlers func(*gin.Engine)) {
	r := gin.New()
	registerHandlers(r)
	if err := r.Run(addr); err != nil {
		logrus.Panic(err)
	}
}
