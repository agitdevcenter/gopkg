package http

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	Logger "github.com/agitdevcenter/gopkg/logger"
	Middleware "github.com/agitdevcenter/gopkg/transport/http/middleware"
	Router "github.com/agitdevcenter/gopkg/transport/http/router"
	Echo "github.com/labstack/echo/v4"
	"golang.org/x/net/http2"
	"io"
	"os"
	"strings"

	"github.com/labstack/echo-contrib/jaegertracing"

	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"time"
)

type Server struct {
	inherit                 bool
	debug                   bool
	logger                  Logger.Logger
	host                    string
	serverName              string
	port                    int
	tls                     bool
	h2c                     bool
	h2cMaxConcurrentStreams uint32
	h2cMaxReadFrameSize     uint32
	insecureSkipVerify      bool
	certificateFile         string
	keyFile                 string
	rootCAFile              string
	readTimeout             time.Duration
	writeTimeout            time.Duration
	idleTimeout             time.Duration
	keepAliveEnabled        bool
	gracefulShutdownTime    time.Duration
	router                  Router.Router
	middleware              *Middleware.Middleware
	tracing                 bool
	tracingName             string
	skipTracingURLs         []string
}

func New(opts []Option) *Server {
	s := &Server{
		port:                 2202,
		gracefulShutdownTime: 5 * time.Second,
		inherit:              true,
	}

	for _, opt := range opts {
		opt(s)
	}

	if s.logger == nil {
		s.logger = Logger.Noop()
	}

	return s
}

func (s *Server) SetDebug(enabled bool) {
	s.debug = enabled
}

func (s *Server) SetLogger(logger Logger.Logger) {
	s.logger = logger
}

func (s *Server) SetupMiddleware(opts []Middleware.Option) {
	s.middleware = Middleware.New(opts)
	s.inheritMiddleware()
}

func (s *Server) inheritMiddleware() {
	if s.inherit && s.middleware != nil {
		s.middleware.SetDebug(s.debug)
		s.middleware.SetLogger(s.logger)
		s.middleware.SetPort(s.port)
	}
}

func (s *Server) SetMiddleware(middleware *Middleware.Middleware) {
	s.middleware = middleware
	s.inheritMiddleware()
}

func (s *Server) Address() string {
	return fmt.Sprintf("%s:%d", s.host, s.port)
}

func (s *Server) Start(ctx context.Context, wg *sync.WaitGroup) func() error {
	s.inheritMiddleware()
	return func() error {

		if checkPort, err := net.Listen("tcp", s.Address()); err != nil {
			return fmt.Errorf("%s already in use, error : %+v", s.Address(), err)
		} else {
			if err := checkPort.Close(); err != nil {
				return fmt.Errorf("closing http connection on %s, error : %+v", s.Address(), err)
			}
		}

		echo := Echo.New()

		var tracing io.Closer
		if s.tracing {
			if len(strings.TrimSpace(s.tracingName)) > 0 {
				os.Setenv("JAEGER_SERVICE_NAME", s.tracingName)
			}
			tracing = jaegertracing.New(echo, func(e Echo.Context) bool {
				for _, url := range s.skipTracingURLs {
					if strings.HasPrefix(e.Path(), url) {
						return true
					}
				}
				return false
			})
		}

		echo.Server.Addr = s.Address()

		if s.readTimeout > 0 {
			echo.Server.ReadTimeout = s.readTimeout
		}
		if s.writeTimeout > 0 {
			echo.Server.WriteTimeout = s.writeTimeout
		}
		if s.idleTimeout > 0 {
			echo.Server.IdleTimeout = s.idleTimeout
		}

		echo.Server.SetKeepAlivesEnabled(s.keepAliveEnabled)

		if s.middleware != nil {
			s.middleware.Setup(echo)
		}

		if s.router != nil {
			s.router.Route(echo)
		}

		errorServer := make(chan error, 1)

		go func() {
			<-ctx.Done()
			shutdownContext, cancel := context.WithTimeout(context.Background(), s.gracefulShutdownTime)
			defer cancel()
			if err := echo.Server.Shutdown(shutdownContext); err != nil {
				errorServer <- fmt.Errorf("error shutting down http server on %s, error : %+v", s.Address(), err)
			}
			if s.tracing {
				if err := tracing.Close(); err != nil {
					errorServer <- fmt.Errorf("error closing down tracing : %+v", err)
				}
			}
			close(errorServer)
			wg.Done()
		}()

		if s.debug {
			s.logger.Info(fmt.Sprintf("starting http server on %s", s.Address()))
		}

		var errorStartingServer error
		if s.tls {
			if len(s.rootCAFile) > 0 {
				certPool := x509.NewCertPool()
				ca, err := ioutil.ReadFile(s.rootCAFile)
				if err != nil {
					return fmt.Errorf("starting http server error : %+v", err)
				}
				if ok := certPool.AppendCertsFromPEM(ca); !ok {
					return fmt.Errorf("starting http server error : append certs from pem")
				}
				echo.Server.TLSConfig = &tls.Config{
					ServerName:         s.serverName,
					RootCAs:            certPool,
					InsecureSkipVerify: s.insecureSkipVerify,
				}
			}
			errorStartingServer = echo.Server.ListenAndServeTLS(s.certificateFile, s.keyFile)
		} else if s.h2c {
			h2c := &http2.Server{
				MaxConcurrentStreams: s.h2cMaxConcurrentStreams,
				MaxReadFrameSize:     s.h2cMaxReadFrameSize,
			}

			if s.idleTimeout > 0 {
				h2c.IdleTimeout = s.idleTimeout
			}

			echo.HideBanner = true
			echo.HidePort = true

			errorStartingServer = echo.StartH2CServer(s.Address(), h2c)
		} else {
			errorStartingServer = echo.Server.ListenAndServe()
		}

		if errorStartingServer != http.ErrServerClosed {
			return fmt.Errorf("starting http server error : %+v", errorStartingServer)
		}

		if s.debug {
			s.logger.Info(fmt.Sprintf("http server stopped on : %s", s.Address()))
		}

		err := <-errorServer
		wg.Wait()
		return err
	}
}
