package main

import (
	"fmt"
	Logger "github.com/agitdevcenter/gopkg/logger"
	"github.com/agitdevcenter/gopkg/transport"
	HTTPHandler "github.com/agitdevcenter/gopkg/transport/example/http/handler"
	HTTPRouter "github.com/agitdevcenter/gopkg/transport/example/http/router"
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

	httpHandler := HTTPHandler.New()

	httpRouter := HTTPRouter.New(httpHandler)

	httpMiddleware := HTTPMiddleware.New([]HTTPMiddleware.Option{
		HTTPMiddleware.WithHealth("/ping"),
		HTTPMiddleware.WithRecover(true),
		HTTPMiddleware.WithErrorHandler(true),
		HTTPMiddleware.WithAvailability("/hello"),
		//HTTPMiddleware.WithProfiling(true),
	})

	httpServer := http.New([]http.Option{
		http.WithInherit(true),
		http.WithPort(8765),
		http.WithMiddleware(httpMiddleware),
		http.WithRouter(httpRouter),
		http.WithH2C(true, 750, 1048576),
		http.WithTimeout(5*time.Second, 10*time.Second, 0*time.Second),
	})

	t := transport.New([]transport.Option{
		transport.WithInherit(true),
		transport.WithDebug(true),
		transport.WithLogger(logger),
		transport.WithHTTPServer(httpServer),
	})

	if err := t.Run(); err != nil {
		logger.Error(fmt.Sprintf("error transport : %+v", err))
	}
}
