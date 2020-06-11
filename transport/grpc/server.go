package grpc

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/agitdevcenter/gopkg/grpc/health"
	Logger "github.com/agitdevcenter/gopkg/logger"
	Handler "github.com/agitdevcenter/gopkg/transport/grpc/handler"
	Interceptor "github.com/agitdevcenter/gopkg/transport/grpc/interceptor"
	"github.com/agitdevcenter/gopkg/transport/grpc/tracer"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	gRPCOpenTracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"io"
	"io/ioutil"
	"net"
	"strings"
	"sync"
)

type Server struct {
	logger                    Logger.Logger
	debug                     bool
	inherit                   bool
	port                      int
	host                      string
	tls                       bool
	certificateFile           string
	keyFile                   string
	rootCAFile                string
	insecureSkipVerify        bool
	serverName                string
	keepAlivePolicy           keepalive.EnforcementPolicy
	keepAliveServerParameters keepalive.ServerParameters
	stream                    bool
	unary                     bool
	handler                   Handler.Handler
	interceptor               *Interceptor.Interceptor
	tracing                   bool
	tracingName               string
}

func New(opts []Option) *Server {
	s := &Server{
		port:    2202,
		inherit: true,
	}

	for _, opt := range opts {
		opt(s)
	}

	if s.logger == nil {
		s.logger = Logger.Noop()
	}

	return s
}

func (s *Server) SetLogger(logger Logger.Logger) {
	s.logger = logger
}

func (s *Server) SetDebug(enabled bool) {
	s.debug = enabled
}

func (s *Server) Address() string {
	return fmt.Sprintf("%s:%d", s.host, s.port)
}

func (s *Server) inheritInterceptor() {
	if s.inherit && s.interceptor != nil {
		s.interceptor.SetDebug(s.debug)
		s.interceptor.SetLogger(s.logger)
		s.interceptor.SetPort(s.port)
	}
}

func (s *Server) Start(ctx context.Context, wg *sync.WaitGroup) func() error {
	s.inheritInterceptor()
	return func() error {

		var trc opentracing.Tracer
		var closer io.Closer
		if s.tracing {
			if len(strings.TrimSpace(s.tracingName)) == 0 {
				s.tracingName = "gRPC"
			}
			var err error
			if trc, closer, err = tracer.New(s.tracingName); err != nil {
				return fmt.Errorf("error : %+v", err)
			}
			opentracing.SetGlobalTracer(trc)
		}

		if !s.unary && !s.stream {
			return fmt.Errorf("please specify grpc service method (unary &/ stream)")
		}

		var listener net.Listener
		var err error
		if listener, err = net.Listen("tcp", s.Address()); err != nil {
			return fmt.Errorf("%s already in use, error : %+v", s.Address(), err)
		}

		// initialize grpc options
		var options []grpc.ServerOption

		if s.tls {
			certificate, err := tls.LoadX509KeyPair(s.certificateFile, s.keyFile)
			if err != nil {
				return fmt.Errorf("could not load grpc certificates : %+v", certificate)
			}

			tlsConfig := &tls.Config{
				Certificates: []tls.Certificate{certificate},
			}

			if len(s.rootCAFile) > 0 {
				certPool := x509.NewCertPool()

				ca, err := ioutil.ReadFile(s.rootCAFile)
				if err != nil {
					return fmt.Errorf("could not load grpc ca authority : %+v", err)
				}

				if ok := certPool.AppendCertsFromPEM(ca); !ok {
					return fmt.Errorf("could not append grpc certs from pem")
				}
				tlsConfig.ServerName = s.serverName
				tlsConfig.RootCAs = certPool
				tlsConfig.InsecureSkipVerify = s.insecureSkipVerify
			}

			credential := credentials.NewTLS(tlsConfig)
			options = append(options, grpc.Creds(credential))
		}

		options = append(options, grpc.KeepaliveEnforcementPolicy(s.keepAlivePolicy))
		options = append(options, grpc.KeepaliveParams(s.keepAliveServerParameters))

		if s.interceptor != nil {
			if s.unary {
				if s.tracing {
					options = append(options, grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
						s.interceptor.Unary(),
						gRPCOpenTracing.UnaryServerInterceptor(gRPCOpenTracing.WithTracer(trc)),
					)))
				} else {
					options = append(options, grpc.UnaryInterceptor(s.interceptor.Unary()))
				}
			}

			if s.stream {
				if s.tracing {
					options = append(options, grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
						s.interceptor.Stream(),
						gRPCOpenTracing.StreamServerInterceptor(gRPCOpenTracing.WithTracer(trc)),
					)))
				} else {
					options = append(options, grpc.StreamInterceptor(s.interceptor.Stream()))
				}
			}
		}

		server := grpc.NewServer(options...)

		health.RegisterHealthServer(server)

		if s.handler != nil {
			s.handler.Register(server)
		}

		errorServer := make(chan error, 1)

		go func() {
			<-ctx.Done()
			server.GracefulStop()
			if s.tracing {
				if err := closer.Close(); err != nil {
					errorServer <- fmt.Errorf("error closing down tracing : %+v", err)
				}
			}
			close(errorServer)
			wg.Done()
		}()

		if s.debug {
			s.logger.Info(fmt.Sprintf("starting grpc server on %s", s.Address()))
		}

		if err := server.Serve(listener); err != nil {
			return fmt.Errorf("error starting grpc server : %+v", err)
		}

		if s.debug {
			s.logger.Info(fmt.Sprintf("grpc server stopped on %s", s.Address()))
		}

		err = <-errorServer
		wg.Wait()
		return err
	}
}
