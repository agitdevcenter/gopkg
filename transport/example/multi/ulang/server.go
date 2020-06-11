package ulang

import (
	"context"
	"fmt"
	"github.com/agitdevcenter/gopkg/backoff"
	Logger "github.com/agitdevcenter/gopkg/logger"
	"sync"
	"time"
)

type Server struct {
	logger Logger.Logger
	debug  bool
}

func New() *Server {

	return &Server{}
}

func (k *Server) SetLogger(logger Logger.Logger) {
	k.logger = logger
}

func (k *Server) SetDebug(enabled bool) {
	k.debug = enabled
}

func (k *Server) Start(ctx context.Context, wg *sync.WaitGroup) func() error {
	return func() error {

		loop := true

		errorServer := make(chan error, 1)

		go func() {
			<-ctx.Done()
			loop = false
			close(errorServer)
			wg.Done()
		}()

		if k.debug {
			k.logger.Info("starting ulang")
		}
		i := 0
		for loop {
			duration := backoff.Default.Duration(i)
			k.logger.Info(fmt.Sprintf("sleep for : %+v", duration))
			time.Sleep(duration)
			i++
		}

		if k.debug {
			k.logger.Info("ulang stopped")
		}

		err := <-errorServer
		wg.Wait()
		return err
	}
}
