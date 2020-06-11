package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"

	ConsulGRPC "github.com/agitdevcenter/gopkg/consul/grpc"
	Session "github.com/agitdevcenter/gopkg/session"
)

type RpcConnection struct {
	options Options
	Conn    *grpc.ClientConn
}

func (rpc *RpcConnection) CreateContext(parent context.Context, session *Session.Session) (ctx context.Context) {
	ctx, _ = context.WithTimeout(parent, rpc.options.Timeout*time.Second)
	ctx = context.WithValue(ctx, sessionKey, session)
	return
}

func NewGRpcConnection(options Options) *RpcConnection {
	var conn *grpc.ClientConn
	var err error

	if options.Resolver == ConsulResolver {
		consulConnect := ConsulGRPC.ConsulConnect{
			Timeout: options.Timeout * time.Second,
			Target:  options.Address,
		}
		conn, err = ConsulGRPC.ConnectWithRoundRobinWithOption(&consulConnect, withClientUnaryInterceptor())
	} else {
		conn, err = grpc.Dial(options.Address, grpc.WithInsecure(), withClientUnaryInterceptor())
	}

	if err != nil {
		panic(err)
	}

	return &RpcConnection{
		Conn:    conn,
		options: options,
	}
}

func NewGRpcConnectionE(options Options) (rpc *RpcConnection, err error) {
	// todo still always insecure
	var conn *grpc.ClientConn

	if options.Resolver == ConsulResolver {
		consulConnect := ConsulGRPC.ConsulConnect{
			Timeout: options.Timeout * time.Second,
			Target:  options.Address,
		}
		conn, err = ConsulGRPC.ConnectWithRoundRobinWithOption(&consulConnect, withClientUnaryInterceptor())
	} else {
		conn, err = grpc.Dial(options.Address, grpc.WithInsecure(), withClientUnaryInterceptor())
	}

	if err != nil {
		return
	}

	rpc = &RpcConnection{
		Conn:    conn,
		options: options,
	}
	return
}

func clientInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	session := ctx.Value(sessionKey).(*Session.Session)
	ctxWithMetadata := metadata.NewOutgoingContext(ctx, metadata.Pairs(XRequestID, session.ThreadID))
	processTime := session.T2("[request]", method, req)
	err := invoker(ctxWithMetadata, method, req, reply, cc, opts...)
	if err != nil {
		session.T3(processTime, "[response][error]", method, err)
		return err
	}
	session.T3(processTime, "[response]", method, reply)
	return err
}

func withClientUnaryInterceptor() grpc.DialOption {
	return grpc.WithUnaryInterceptor(clientInterceptor)
}
