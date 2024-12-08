package server

import (
	"net"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func RunGRPCServer(serviceName string, registerServer func(*grpc.Server)) {
	addr := viper.Sub(serviceName).GetString("grpc_addr")
	if addr == "" {
		logrus.Panicf("please provide grpc address for %s service", serviceName)
	}
	RunGRPCServerOnAddr(addr, registerServer)
}

func RunGRPCServerOnAddr(addr string, registerServer func(*grpc.Server)) {
	s := grpc.NewServer()
	registerServer(s)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Panic(err)
	}

	logrus.Infof("grpc server is running on %s", addr)
	if err := s.Serve(lis); err != nil {
		logrus.Panic(err)
	}
}
