package transport

import (
	Logger "github.com/agitdevcenter/gopkg/logger"
	"github.com/agitdevcenter/gopkg/transport/custom"
	"github.com/agitdevcenter/gopkg/transport/grpc"
	"github.com/agitdevcenter/gopkg/transport/http"
)

type Option func(*Transport)

func WithInherit(inherit bool) Option {
	return func(t *Transport) {
		t.inherit = inherit
	}
}

func WithLogger(logger Logger.Logger) Option {
	return func(t *Transport) {
		t.logger = logger
	}
}

func WithDebug(enabled bool) Option {
	return func(t *Transport) {
		t.debug = enabled
	}
}

func WithHTTPServer(httpServer *http.Server) Option {
	return func(t *Transport) {
		t.httpServer = httpServer
	}
}

func WithGRPCServer(grpcServer *grpc.Server) Option {
	return func(t *Transport) {
		t.grpcServer = grpcServer
	}
}

func WithCustom(service *custom.Holder) Option {
	return func(t *Transport) {
		t.services = append(t.services, service)
	}
}
