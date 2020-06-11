package main

import (
	"fmt"
	Logger "github.com/agitdevcenter/gopkg/logger"
	"github.com/agitdevcenter/gopkg/transport"
	"github.com/agitdevcenter/gopkg/transport/custom"
	"github.com/agitdevcenter/gopkg/transport/example/kafka/service"
	"time"
)

func main() {
	logger := Logger.New(Logger.Options{
		FileLocation:    "tdr.log",
		FileTdrLocation: "log.log",
		FileMaxAge:      time.Hour,
		Stdout:          true,
	})

	k := service.New()
	holder := custom.New(custom.OptionService(k))

	t := transport.New([]transport.Option{
		transport.WithDebug(true),
		transport.WithLogger(logger),
		transport.WithInherit(true),
		transport.WithCustom(holder),
	})

	if err := t.Run(); err != nil {
		logger.Error(fmt.Sprintf("error transport : %+v", err))
	}
}
