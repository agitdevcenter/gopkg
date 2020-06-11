package main

import (
	"fmt"
	Logger "github.com/agitdevcenter/gopkg/logger"
	"github.com/agitdevcenter/gopkg/transport"
	Handler "github.com/agitdevcenter/gopkg/transport/example/grpc/handler"
	"github.com/agitdevcenter/gopkg/transport/grpc"
	Interceptor "github.com/agitdevcenter/gopkg/transport/grpc/interceptor"
	"time"
)

func main() {
	logger := Logger.New(Logger.Options{
		FileLocation:    "tdr.log",
		FileTdrLocation: "log.log",
		FileMaxAge:      time.Hour,
		Stdout:          true,
	})

	interceptor := Interceptor.New([]Interceptor.Option{})

	handler := Handler.New()

	g := grpc.New([]grpc.Option{
		grpc.WithInherit(true),
		grpc.WithUnary(true),
		grpc.WithInterceptor(interceptor),
		grpc.WithHandler(handler),
	})

	t := transport.New([]transport.Option{
		transport.WithDebug(true),
		transport.WithLogger(logger),
		transport.WithInherit(true),
		transport.WithGRPCServer(g),
	})

	if err := t.Run(); err != nil {
		logger.Error(fmt.Sprintf("error transport : %+v", err))
	}
}
