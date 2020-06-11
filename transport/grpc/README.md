# gRPC Server
gRPC Server using [grpc-go](https://github.com/grpc/grpc-go).

## Setup
There are several ways to setup gRPC server.

## Default Setup
This will setup gRPC server at port 2202
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/grpc"
)

func main() {
    g := grpc.New([]grpc.Option{})
}
```

### With Options
Setting up gRPC server with options from `grpc.Option`.

#### Port
`grpc.WithPort` `int` parameter. It will set the gRPC `port` value.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/grpc"
)

func main() {
    g := grpc.New([]grpc.Option{grpc.WithPort(80)})
}
```

#### Host
`grpc.WithHost` `string` parameter. It will set the gRPC `host` value.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/grpc"
)

func main() {
    g := grpc.New([]grpc.Option{grpc.WithHost("linkaja.id")})
}
```

#### Unary
`grpc.WithUnary` `boolean` parameter. It will set the gRPC `unary` value.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/grpc"
)

func main() {
    g := grpc.New([]grpc.Option{grpc.WithUnary(true)})
}
```

#### Stream
`grpc.WithStream` `boolean` parameter. It will set the gRPC `stream` value.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/grpc"
)

func main() {
    g := grpc.New([]grpc.Option{grpc.WithStream(true)})
}
```

#### Address
`grpc.WithAddress` host `string`, port `int` parameters. It will set the gRPC `host` and `port` value.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/grpc"
)

func main() {
    g := grpc.New([]grpc.Option{grpc.WithAddress("linkaja.id", 80)})
}
```

#### Debug
`grpc.WithDebug` `boolean` parameter. It will set the gRPC `debug` value.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/grpc"
)

func main() {
    g := grpc.New([]grpc.Option{grpc.WithDebug(true)})
}
```

#### Logger
`grpc.WithLogger` `pkg/logger` parameter. It will set the gRPC `logger` value.
```go
package main

import (
    Logger "github.com/agitdevcenter/gopkg/logger"
    "github.com/agitdevcenter/gopkg/transport/grpc"
    "time"
)

func main() {
    logger := Logger.New(Logger.Options{
        FileLocation:    "tdr.log",
        FileTdrLocation: "log.log",
        FileMaxAge:      time.Hour,
        Stdout:          true,
    })
    g := grpc.New([]grpc.Option{grpc.WithLogger(logger)})
}
```

#### TLS
`grpc.WithTLS` certificateFile, keyFile, rootCAFile, serverName `string`, insecureSkipVerify `boolean` parameters. `required` certificateFile, keyFile. `allow empty` rootCAFile, serverName. If rootCAFile is provided, serverName should be also provided and tied with the rootCAFile. This will setting up gRPC server to run on TLS.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/grpc"
)

func main() {
    g := grpc.New([]grpc.Option{grpc.WithTLS("server.crt", "server.key", "server.pem", "linkaja.id", false)})
}
```

#### Keep Alive Enforcement Policy
`grpc.WithKeepAliveEnforcementPolicy` `keepalive.EnforcementPolicy` parameter. It will set the gRPC `keepAlivePolicy` value.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/grpc"
    "google.golang.org/grpc/keepalive"
)

func main() {
    policy := keepalive.EnforcementPolicy{}
    g := grpc.New([]grpc.Option{grpc.WithKeepAliveEnforcementPolicy(policy)})
}
```

#### Keep Alive Server Parameters
`grpc.WithKeepAliveServerParameters` `keepalive.ServerParameters` parameter. It will set the gRPC `keepAliveServerParameters` value.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/grpc"
    "google.golang.org/grpc/keepalive"
)

func main() {
    serverParameters := keepalive.ServerParameters{}
    g := grpc.New([]grpc.Option{grpc.WithKeepAliveServerParameters(serverParameters)})
}
```

#### Tracing
`grpc.WithTracing` tracing `boolean`, tracingName `string` parameters. It will setup gRPC `tracing` and `tracingName` value. Tracing using [jaeger](https://www.jaegertracing.io/).
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/grpc"
)

func main() {
    g := grpc.New([]grpc.Option{grpc.WithTracing(true, "tegaralaga")})
}
```

#### Handler
`grpc.WithHandler` handler `grpc.Handler` parameter. It will setup gRPC request handler. Handler parameter must implement `grpc.Handler` interfaces.
```proto
syntax = "proto2";

package handler;
option go_package = "handler";

import "github.com/gogo/protobuf/gogoproto/gogo.proto";

option (gogoproto.marshaler_all) = true;
option (gogoproto.sizer_all) = true;
option (gogoproto.unmarshaler_all) = true;
option (gogoproto.goproto_getters_all) = false;

message HelloRequest {
    optional string name = 1 [(gogoproto.nullable) = false];
}

message HelloData {
    optional string hello = 1 [(gogoproto.nullable) = false];
}

message HelloResponse {
    optional string status = 1 [(gogoproto.nullable) = false];
    optional string message = 2 [(gogoproto.nullable) = false];
    optional HelloData data = 3 [(gogoproto.nullable) = false];
}

message DefaultResponse {
    optional string status = 1 [(gogoproto.nullable) = false];
    optional string message = 2 [(gogoproto.nullable) = false];
}

service HelloHandler {
    rpc Hello (HelloRequest) returns (HelloResponse) {};
}
```
```go
package main

import (
    gRPC "github.com/agitdevcenter/gopkg/transport/grpc"
    "google.golang.org/grpc"
    "context"
    Response "github.com/agitdevcenter/gopkg/response"
)

type helloHandler struct{}

func SetupHelloHandler() *helloHandler {
    return &helloHandler{}
}

func (h *helloHandler) Validate() HelloHandlerServer {
    return h
}

// implement generated proto
func (h *helloHandler) Hello(ctx context.Context, request *HelloRequest) (response *HelloResponse, err error) {
    response = &HelloResponse{
        Status:  Response.SuccessCode,
        Message: "Success",
        Data: HelloData{
            Hello: request.Name,
        },
    }
    return
}

type Handler struct {
    Hello HelloHandlerServer
}

func NewHandler() *Handler {
    return &Handler{Hello: SetupHelloHandler().Validate()}
}

func (h *Handler) Register(server *grpc.Server) {
    RegisterHelloHandlerServer(server, h.Hello)
}

func main() {
    handler := NewHandler()
    g := gRPC.New([]gRPC.Option{gRPC.WithHandler(handler)})
}
```

#### Interceptor
`grpc.WithInterceptor` `interceptor.Interceptor` parameter. It will setup gRPC interceptor with default configuration. For more interceptor configuration read [here](interceptor/README.md)
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/grpc"
    "github.com/agitdevcenter/gopkg/transport/grpc/interceptor"
)

func main() {
    i := interceptor.New([]interceptor.Option{})
    g := grpc.New([]grpc.Option{grpc.WithInterceptor(i)})
}
```

### Inherit
`grpc.WithInherit` `boolean` parameter. Interceptor will inherit `debug` and `logger` value from `grpc`
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/grpc"
)

func main() {
    g := grpc.New([]grpc.Option{grpc.WithInherit(true)})
}
```

#### Multiple Options
Running gRPC hello server at port 8080.
```proto
syntax = "proto2";

package handler;
option go_package = "handler";

import "github.com/gogo/protobuf/gogoproto/gogo.proto";

option (gogoproto.marshaler_all) = true;
option (gogoproto.sizer_all) = true;
option (gogoproto.unmarshaler_all) = true;
option (gogoproto.goproto_getters_all) = false;

message HelloRequest {
    optional string name = 1 [(gogoproto.nullable) = false];
}

message HelloData {
    optional string hello = 1 [(gogoproto.nullable) = false];
}

message HelloResponse {
    optional string status = 1 [(gogoproto.nullable) = false];
    optional string message = 2 [(gogoproto.nullable) = false];
    optional HelloData data = 3 [(gogoproto.nullable) = false];
}

message DefaultResponse {
    optional string status = 1 [(gogoproto.nullable) = false];
    optional string message = 2 [(gogoproto.nullable) = false];
}

service HelloHandler {
    rpc Hello (HelloRequest) returns (HelloResponse) {};
}
```

```go
package main

import (
    gRPC "github.com/agitdevcenter/gopkg/transport/grpc"
    "google.golang.org/grpc"
    "context"
    Response "github.com/agitdevcenter/gopkg/response"
    "github.com/agitdevcenter/gopkg/transport/grpc/interceptor"
    "fmt"
    Logger "github.com/agitdevcenter/gopkg/logger"
    "github.com/agitdevcenter/gopkg/transport"
    "time"
)

type helloHandler struct{}

func SetupHelloHandler() *helloHandler {
    return &helloHandler{}
}

func (h *helloHandler) Validate() HelloHandlerServer {
    return h
}

// implement generated proto
func (h *helloHandler) Hello(ctx context.Context, request *HelloRequest) (response *HelloResponse, err error) {
    response = &HelloResponse{
        Status:  Response.SuccessCode,
        Message: "Success",
        Data: HelloData{
            Hello: request.Name,
        },
    }
    return
}

type Handler struct {
    Hello HelloHandlerServer
}

func NewHandler() *Handler {
    return &Handler{Hello: SetupHelloHandler().Validate()}
}

func (h *Handler) Register(server *grpc.Server) {
    RegisterHelloHandlerServer(server, h.Hello)
}

func main() {
    logger := Logger.New(Logger.Options{
        FileLocation:    "tdr.log",
        FileTdrLocation: "log.log",
        FileMaxAge:      time.Hour,
        Stdout:          true,
    })
    
    i := interceptor.New([]interceptor.Option{})
    
    handler := NewHandler()
    
    g := gRPC.New([]gRPC.Option{
        gRPC.WithPort(8080),
        gRPC.WithInherit(true),
        gRPC.WithUnary(true),
        gRPC.WithInterceptor(i),
        gRPC.WithHandler(handler),
    })
    
    t := transport.New([]transport.Option{
        transport.WithDebug(true),
        transport.WithLogger(logger),
        transport.WithInherit(true),
        transport.WithGRPCServer(g),
    })
    
    if err := t.Run(); err != nil {
        logger.Error(fmt.Sprintf("error transport : %+v", err))
    }
}
```

## Miscellaneous
Compile protocol buffer definition
```shell script
protoc -I="${GOPATH}/src/" -I=./ --gogoslick_out=plugins=grpc:. *.proto
```