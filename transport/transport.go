package transport

import (
	"context"
	"fmt"
	Logger "github.com/agitdevcenter/gopkg/logger"
	"github.com/agitdevcenter/gopkg/transport/custom"
	"github.com/agitdevcenter/gopkg/transport/grpc"
	"github.com/agitdevcenter/gopkg/transport/http"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Transport struct {
	inherit    bool
	httpServer *http.Server
	grpcServer *grpc.Server
	logger     Logger.Logger
	debug      bool
	services   []*custom.Holder
}

func New(opts []Option) *Transport {
	t := &Transport{
		inherit: true,
	}

	for _, opt := range opts {
		opt(t)
	}

	if t.logger == nil {
		t.logger = Logger.Noop()
	}

	t.inheritServer()

	return t
}

func (t *Transport) SetHTTP(httpServer *http.Server) {
	t.httpServer = httpServer
	t.inheritHTTPServer()
}

func (t *Transport) SetGRPC(grpcServer *grpc.Server) {
	t.grpcServer = grpcServer
	t.inheritGRPCServer()
}

func (t *Transport) inheritServer() {
	t.inheritHTTPServer()
	t.inheritGRPCServer()
}

func (t *Transport) inheritHTTPServer() {
	if t.inherit && t.httpServer != nil {
		t.httpServer.SetDebug(t.debug)
		t.httpServer.SetLogger(t.logger)
	}
}

func (t *Transport) inheritGRPCServer() {
	if t.inherit && t.grpcServer != nil {
		t.grpcServer.SetDebug(t.debug)
		t.grpcServer.SetLogger(t.logger)
	}
}

func (t *Transport) SetupHTTP(opts []http.Option) {
	t.httpServer = http.New(opts)
	t.inheritHTTPServer()
}

func (t *Transport) SetupGRPC(opts []grpc.Option) {
	t.grpcServer = grpc.New(opts)
	t.inheritGRPCServer()
}

func (t *Transport) Run() (err error) {

	serverAvailable := 0
	if t.httpServer != nil {
		serverAvailable++
	}
	if t.grpcServer != nil {
		serverAvailable++
	}

	if len(t.services) > 0 {
		serverAvailable = serverAvailable + len(t.services)
	}

	var wg sync.WaitGroup
	wg.Add(serverAvailable)

	ctx, cancel := context.WithCancel(context.Background())

	eg, egc := errgroup.WithContext(context.Background())

	// http services
	if t.httpServer != nil {
		eg.Go(t.httpServer.Start(ctx, &wg))
	}

	// grpc service
	if t.grpcServer != nil {
		eg.Go(t.grpcServer.Start(ctx, &wg))
	}

	// custom services
	for _, h := range t.services {
		if t.inherit {
			h.Service().SetLogger(t.logger)
			h.Service().SetDebug(t.debug)
		}
		h.Hold()
		eg.Go(h.Service().Start(ctx, &wg))
	}

	go func() {
		<-egc.Done()
		cancel()
	}()

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Kill, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
		<-signals
		cancel()
	}()

	if err = eg.Wait(); err != nil {
		err = fmt.Errorf("server error : %+v", err)
		return
	}

	if t.debug {
		t.logger.Info("all server closed successfully")
	}

	return
}
