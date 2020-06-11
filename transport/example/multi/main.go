package main

import (
	"fmt"
	Logger "github.com/agitdevcenter/gopkg/logger"
	"github.com/agitdevcenter/gopkg/transport"
	"github.com/agitdevcenter/gopkg/transport/custom"
	Handler "github.com/agitdevcenter/gopkg/transport/example/grpc/handler"
	HTTPHandler "github.com/agitdevcenter/gopkg/transport/example/http/handler"
	HTTPRouter "github.com/agitdevcenter/gopkg/transport/example/http/router"
	"github.com/agitdevcenter/gopkg/transport/example/kafka/service"
	"github.com/agitdevcenter/gopkg/transport/example/multi/ulang"
	"github.com/agitdevcenter/gopkg/transport/grpc"
	Interceptor "github.com/agitdevcenter/gopkg/transport/grpc/interceptor"
	"github.com/agitdevcenter/gopkg/transport/http"
	HTTPMiddleware "github.com/agitdevcenter/gopkg/transport/http/middleware"
	"time"
)

func main() {
	logger := Logger.New(Logger.Options{
		FileLocation:    "tdr.log",
		FileTdrLocation: "log.log",
		FileMaxAge:      time.Hour,
		Stdout:          true,
	})

	// http service
	httpHandler := HTTPHandler.New()

	httpRouter := HTTPRouter.New(httpHandler)

	httpMiddleware := HTTPMiddleware.New([]HTTPMiddleware.Option{})

	httpServer := http.New([]http.Option{
		http.WithInherit(true),
		http.WithPort(8765),
		http.WithMiddleware(httpMiddleware),
		http.WithRouter(httpRouter),
	})

	// gRPC service
	interceptor := Interceptor.New([]Interceptor.Option{})

	handler := Handler.New()

	g := grpc.New([]grpc.Option{
		grpc.WithInherit(true),
		grpc.WithPort(4321),
		grpc.WithUnary(true),
		grpc.WithInterceptor(interceptor),
		grpc.WithHandler(handler),
	})

	// kafka service
	k := service.New()
	kafkaHolder := custom.New(custom.OptionService(k))

	// ulang service
	u := ulang.New()
	ulangHolder := custom.New(custom.OptionService(u))

	// transport
	t := transport.New([]transport.Option{
		transport.WithInherit(true),
		transport.WithDebug(true),
		transport.WithLogger(logger),
		transport.WithHTTPServer(httpServer),
		transport.WithGRPCServer(g),
		transport.WithCustom(kafkaHolder),
		transport.WithCustom(ulangHolder),
	})

	if err := t.Run(); err != nil {
		logger.Error(fmt.Sprintf("error transport : %+v", err))
	}
}
