package http

import (
	Logger "github.com/agitdevcenter/gopkg/logger"
	Middleware "github.com/agitdevcenter/gopkg/transport/http/middleware"
	"github.com/agitdevcenter/gopkg/transport/http/router"
	"time"
)

type Option func(*Server)

func WithInherit(inherit bool) Option {
	return func(s *Server) {
		s.inherit = inherit
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
		s.h2c = false
		s.tls = true
		s.certificateFile = certificateFile
		s.keyFile = keyFile
		s.rootCAFile = rootCAFile
		s.serverName = serverName
		s.insecureSkipVerify = insecureSkipVerify
	}
}

func WithH2C(enabled bool, maxConcurrentStreams, maxReadFrameSize uint32) Option {
	return func(s *Server) {
		if enabled {
			s.h2c = enabled
			s.h2cMaxConcurrentStreams = maxConcurrentStreams
			s.h2cMaxReadFrameSize = maxReadFrameSize
			s.tls = false
		}
	}
}

func WithTimeout(read, write, idle time.Duration) Option {
	return func(s *Server) {
		s.readTimeout = read
		s.writeTimeout = write
		s.idleTimeout = idle
	}
}

func WithKeepAlive(enabled bool) Option {
	return func(s *Server) {
		s.keepAliveEnabled = enabled
	}
}

func WithGracefulShutdownTime(gracefulShutdownTime time.Duration) Option {
	return func(s *Server) {
		s.gracefulShutdownTime = gracefulShutdownTime
	}
}

func WithRouter(router router.Router) Option {
	return func(s *Server) {
		s.router = router
	}
}

func WithMiddleware(middleware *Middleware.Middleware) Option {
	return func(s *Server) {
		s.middleware = middleware
	}
}

func WithTracing(tracing bool, skipTracingURLs []string, tracingName string) Option {
	return func(s *Server) {
		s.tracing = tracing
		s.skipTracingURLs = skipTracingURLs
		s.tracingName = tracingName
	}
}
