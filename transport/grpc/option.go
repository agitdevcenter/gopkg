package grpc

import (
	Logger "github.com/agitdevcenter/gopkg/logger"
	Handler "github.com/agitdevcenter/gopkg/transport/grpc/handler"
	"github.com/agitdevcenter/gopkg/transport/grpc/interceptor"
	"google.golang.org/grpc/keepalive"
)

type Option func(*Server)

func WithStream(enabled bool) Option {
	return func(s *Server) {
		s.stream = enabled
	}
}

func WithUnary(enabled bool) Option {
	return func(s *Server) {
		s.unary = enabled
	}
}

func WithKeepAliveEnforcementPolicy(policy keepalive.EnforcementPolicy) Option {
	return func(s *Server) {
		s.keepAlivePolicy = policy
	}
}

func WithKeepAliveServerParameters(parameters keepalive.ServerParameters) Option {
	return func(s *Server) {
		s.keepAliveServerParameters = parameters
	}
}

func WithInherit(inherit bool) Option {
	return func(s *Server) {
		s.inherit = inherit
	}
}

func WithInterceptor(interceptor *interceptor.Interceptor) Option {
	return func(s *Server) {
		s.interceptor = interceptor
	}
}

func WithHandler(handler Handler.Handler) Option {
	return func(s *Server) {
		s.handler = handler
	}
}

func WithHost(host string) Option {
	return func(s *Server) {
		s.host = host
	}
}

func WithPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

func WithAddress(host string, port int) Option {
	return func(s *Server) {
		s.host = host
		s.port = port
	}
}

func WithDebug(enabled bool) Option {
	return func(s *Server) {
		s.debug = enabled
	}
}

func WithLogger(logger Logger.Logger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}

func WithTLS(certificateFile, keyFile, rootCAFile, serverName string, insecureSkipVerify bool) Option {
	return func(s *Server) {
		s.tls = true
		s.certificateFile = certificateFile
		s.keyFile = keyFile
		s.rootCAFile = rootCAFile
		s.serverName = serverName
		s.insecureSkipVerify = insecureSkipVerify
	}
}

func WithTracing(tracing bool, tracingName string) Option {
	return func(s *Server) {
		s.tracing = tracing
		s.tracingName = tracingName
	}
}
