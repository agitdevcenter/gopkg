package interceptor

import (
	"context"
	"fmt"
	Logger "github.com/agitdevcenter/gopkg/logger"
	Session "github.com/agitdevcenter/gopkg/session"
	"github.com/agitdevcenter/gopkg/utils"
	ValueObject "github.com/agitdevcenter/gopkg/vo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"strings"
)

const (
	InternalServerErrorMessage = "Internal Server Error"
	Name                       = "LinkAja"
	Version                    = "1.0.0"
)

var additionalHandlers []func(interface{})

type WrappedServerStream struct {
	grpc.ServerStream
	WrappedContext context.Context
}

func (w *WrappedServerStream) Context() context.Context {
	return w.WrappedContext
}

func WrapServerStream(stream grpc.ServerStream) *WrappedServerStream {
	if existing, ok := stream.(*WrappedServerStream); ok {
		return existing
	}
	return &WrappedServerStream{ServerStream: stream, WrappedContext: stream.Context()}
}

type Interceptor struct {
	logger                     Logger.Logger
	debug                      bool
	session                    bool
	name                       string
	version                    string
	port                       int
	handleCrash                bool
	skipRPCs                   []string
	internalServerErrorMessage string
}

func New(opts []Option) *Interceptor {
	i := &Interceptor{
		name:                       Name,
		version:                    Version,
		handleCrash:                true,
		skipRPCs:                   []string{},
		session:                    true,
		internalServerErrorMessage: InternalServerErrorMessage,
	}

	for _, opt := range opts {
		opt(i)
	}

	if i.logger == nil {
		i.logger = Logger.Noop()
	}

	return i
}

func (i *Interceptor) SetLogger(logger Logger.Logger) {
	i.logger = logger
}

func (i *Interceptor) SetDebug(enabled bool) {
	i.debug = enabled
}

func (i *Interceptor) SetPort(port int) {
	i.port = port
}

func (i *Interceptor) Stream() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		var session *Session.Session
		if i.session {
			session = Session.New(i.logger).
				SetAppName(i.name).
				SetAppVersion(i.version).
				SetPort(i.port).
				//SetRequest(). @TODO how to get request?
				SetURL(info.FullMethod).
				SetThreadID(getXID(stream.Context())).
				SetSrcIP(getRealIP(stream.Context())).
				SetIP(getIP(stream.Context())).
				SetMethod("gRPC")
		}

		defer handleCrash(func(r interface{}) {
			err = i.panicError(r, session)
		})

		var ctx context.Context

		if !i.skip(info.FullMethod) {
			ctx = context.WithValue(ctx, ValueObject.AppSession, session)
		}

		wrapped := WrapServerStream(stream)
		wrapped.WrappedContext = ctx
		return handler(srv, wrapped)
	}
}

func (i *Interceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (response interface{}, err error) {
		var session *Session.Session
		if i.session {
			session = Session.New(i.logger).
				SetAppName(i.name).
				SetAppVersion(i.version).
				SetPort(i.port).
				SetRequest(req).
				SetURL(info.FullMethod).
				SetThreadID(getXID(ctx)).
				SetSrcIP(getRealIP(ctx)).
				SetIP(getIP(ctx)).
				SetMethod("gRPC")
		}

		if !i.skip(info.FullMethod) {
			if i.session && session != nil {
				session.T1("Incoming Request")
			}
			ctx = context.WithValue(ctx, ValueObject.AppSession, session)
		}

		if i.handleCrash {
			defer handleCrash(func(r interface{}) {
				err = i.panicError(r, session)
			})
		}

		response, err = handler(ctx, req)

		if i.session && session != nil {
			session.T4(response)
		}

		return
	}
}

func (i *Interceptor) skip(method string) (skip bool) {
	for _, url := range i.skipRPCs {
		if strings.HasPrefix(strings.ToLower(method), url) {
			skip = true
			return
		}
	}
	return
}

func getXID(ctx context.Context) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if xid, ok := md["X-Request-ID"]; ok {
			if len(xid) > 0 {
				return xid[0]
			}
		}
		if xid, ok := md["xid"]; ok {
			if len(xid) > 0 {
				return xid[0]
			}
		}
	}
	return utils.GenerateThreadId()
}

func getRealIP(ctx context.Context) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if xRealIP, ok := md["X-Real-IP"]; ok {
			if len(xRealIP) > 0 {
				return xRealIP[0]
			}
		}
		if xRealIP, ok := md["x-real-ip"]; ok {
			if len(xRealIP) > 0 {
				return xRealIP[0]
			}
		}
		if xForwardedFor, ok := md["X-Forwarded-For"]; ok {
			if len(xForwardedFor) > 0 {
				return xForwardedFor[0]
			}
		}
		if xForwardedFor, ok := md["x-forwarded-for"]; ok {
			if len(xForwardedFor) > 0 {
				return xForwardedFor[0]
			}
		}
	}
	return getIP(ctx)
}

func getIP(ctx context.Context) string {
	if p, ok := peer.FromContext(ctx); ok {
		return p.Addr.String()
	}
	return ""
}

func handleCrash(handler func(interface{})) {
	if r := recover(); r != nil {
		handler(r)

		if additionalHandlers != nil {
			for _, fn := range additionalHandlers {
				fn(r)
			}
		}
	}
}

func (i *Interceptor) panicError(r interface{}, session *Session.Session) error {
	message := fmt.Sprintf("gRPC error : %+v", r)
	if i.session && session != nil {
		session.T4(message)
	}
	i.logger.Error(message)
	return status.Errorf(codes.Internal, i.internalServerErrorMessage)
}
