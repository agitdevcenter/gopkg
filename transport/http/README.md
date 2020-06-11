# HTTP Server
HTTP Server using [echo](https://github.com/labstack/echo) version [4](https://github.com/labstack/echo/releases).

## Setup
There are several ways to setup HTTP server.

## Default Setup
This will setup HTTP server at port 2202
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/http"
)

func main() {
    h := http.New([]http.Option{})
}
```

### With Options
Setting up HTTP server with options from `http.Option`.

#### Port
`http.WithPort` `int` parameter. It will set the HTTP `port` value.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/http"
)

func main() {
    h := http.New([]http.Option{http.WithPort(80)})
}
```

#### Graceful Shutdown Time
`http.WithGracefulShutdownTime` `time.Duration` parameter. It will set the HTTP `gracefulShutdownTime` value.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/http"
    "time"
)

func main() {
    h := http.New([]http.Option{http.WithGracefulShutdownTime(10 * time.Second)})
}
```

#### Host
`http.WithHost` `string` parameter. It will set the HTTP `host` value.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/http"
)

func main() {
    h := http.New([]http.Option{http.WithHost("linkaja.id")})
}
```

#### Address
`http.WithAddress` host `string`, port `int` parameters. It will set the HTTP `host` and `port` value.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/http"
)

func main() {
    h := http.New([]http.Option{http.WithAddress("linkaja.id", 80)})
}
```

#### Debug
`http.WithDebug` `boolean` parameter. It will set the HTTP `debug` value.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/http"
)

func main() {
    h := http.New([]http.Option{http.WithDebug(true)})
}
```

#### Logger
`http.WithLogger` `pkg/logger` parameter. It will set the HTTP `logger` value.
```go
package main

import (
    Logger "github.com/agitdevcenter/gopkg/logger"
    "github.com/agitdevcenter/gopkg/transport/http"
    "time"
)

func main() {
    logger := Logger.New(Logger.Options{
        FileLocation:    "tdr.log",
        FileTdrLocation: "log.log",
        FileMaxAge:      time.Hour,
        Stdout:          true,
    })

    h := http.New([]http.Option{http.WithLogger(logger)})
}
```

#### TLS
`http.WithTLS` certificateFile, keyFile, rootCAFile, serverName `string`, insecureSkipVerify `boolean` parameters. `required` certificateFile, keyFile. `allow empty` rootCAFile, serverName. If rootCAFile is provided, serverName should be also provided and tied with the rootCAFile. This will setting up HTTP server to run on TLS.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/http"
)

func main() {
    h := http.New([]http.Option{http.WithTLS("server.crt", "server.key", "server.pem", "linkaja.id", false)})
}
```

#### H2C
`http.WithH2C` enabled `boolean`, maxConcurrentStreams, maxReadFrameSize `uint32` parameters, . Set HTTP `h2c` value, if it is `true`, it will run HTTP server on HTTP/2 Cleartext (insecure). H2C is useful for service to service (private usage), offers HTTP/2 features without the TLS handshake.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/http"
)

func main() {
    h := http.New([]http.Option{http.WithH2C(true, 250, 1048576)})
}
```

#### Timeout
`http.WithTimeout` read, write, idle `time.Duration` parameters. It will setup HTTP `ReadTimeout`, `WriteTimeout`, and `IdleTimeout`.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/http"
    "time"
)

func main() {
    h := http.New([]http.Option{http.WithTimeout(5 * time.Second, 5 * time.Second, 5 * time.Second)})
}
```

#### KeepAlive
`http.WithKeepAlive` `boolean` parameter. It will setup HTTP `KeepAlivesEnabled`.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/http"
)

func main() {
    h := http.New([]http.Option{http.WithKeepAlive(true)})
}
```

#### Tracing
`http.WithTracing` tracing `boolean`, skipTracingURLs `[]string` parameters, tracingName `string`. It will setup HTTP `tracing`, `skipTracingURLS`, and `tracingName` values. Tracing using [jaeger](https://www.jaegertracing.io/), set skipTracingURLS to skip tracing those urls.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/http"
)

func main() {
    h := http.New([]http.Option{http.WithTracing(true, []string{}, "tegaralaga")})
}
```

#### Router
`http.WithRouter` router `router.Router` parameter. It will setup HTTP request routing. Router parameter must implement `router.Router` interfaces.
```go
package main

import (
    Echo "github.com/labstack/echo/v4"
    HTTP "github.com/agitdevcenter/gopkg/transport/http"
    "net/http"
)

type HealthHandler struct {}

func NewHealthHandler() *HealthHandler {
    return &HealthHandler{}
}

func (h *HealthHandler) Ping(ctx Echo.Context) (err error) {
    return ctx.NoContent(http.StatusNoContent)
}

type Handler struct {
    Health *HealthHandler
}

func NewHandler() *Handler {
    return &Handler{Health: NewHealthHandler()}
}

type Router struct{
    handler *Handler
}

func NewRouter(handler *Handler) *Router {
    return &Router{handler: handler}
}

// implement from router.Router
func (r *Router) Route(echo *Echo.Echo) {
    health := echo.Group("/health")
    health.GET("", r.handler.Health.Ping)
}


func main() {
    handler := NewHandler()
    router := NewRouter(handler)

    h := HTTP.New([]HTTP.Option{HTTP.WithRouter(router)})
}
```

#### Middleware
`http.WithMiddleware` `middleware.Middleware` parameter. It will setup HTTP middleware with default configuration. For more middleware configuration read [here](middleware/README.md)
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/http"
    "github.com/agitdevcenter/gopkg/transport/http/middleware"
)

func main() {
    m := middleware.New([]middleware.Option{})
    h := http.New([]http.Option{http.WithMiddleware(m)})
}
```

### Inherit
`http.WithInherit` `boolean` parameter. Middleware will inherit `debug` and `logger` value from `http`
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/http"
)

func main() {
    h := http.New([]http.Option{http.WithInherit(true)})
}
```

#### Multiple Options
Running HTTP health check server at port 80.
```go
package main

import (
    "fmt"
    Echo "github.com/labstack/echo/v4"
    Logger "github.com/agitdevcenter/gopkg/logger"
    "github.com/agitdevcenter/gopkg/transport"
    HTTP "github.com/agitdevcenter/gopkg/transport/http"
    "github.com/agitdevcenter/gopkg/transport/http/middleware"
    "github.com/agitdevcenter/gopkg/vo"
    "log"
    "net/http"
    "time"
)

type HealthHandler struct {}

func NewHealthHandler() *HealthHandler {
    return &HealthHandler{}
}

func (h *HealthHandler) Ping(ctx Echo.Context) (err error) {
    return ctx.NoContent(http.StatusNoContent)
}

type Handler struct {
    Health *HealthHandler
}

func NewHandler() *Handler {
    return &Handler{Health: NewHealthHandler()}
}

type Router struct{
    handler *Handler
}

func NewRouter(handler *Handler) *Router {
    return &Router{handler: handler}
}

// implement from router.Router
func (r *Router) Route(echo *Echo.Echo) {
    health := echo.Group("/health")
    health.GET("", r.handler.Health.Ping)
}

func main() {
    logger := Logger.New(Logger.Options{
        FileLocation:    "tdr.log",
        FileTdrLocation: "log.log",
        FileMaxAge:      time.Hour,
        Stdout:          true,
    })

    m := middleware.New([]middleware.Option{})

    handler := NewHandler()

    router := NewRouter(handler)

    h := HTTP.New([]HTTP.Option{
        HTTP.WithPort(80),
        HTTP.WithMiddleware(m),
        HTTP.WithRouter(router),
        HTTP.WithInherit(true),
    })
    
    t := transport.New([]transport.Option{
        transport.WithDebug(true),
        transport.WithLogger(logger),
        transport.WithHTTPServer(h),
        transport.WithInherit(true),
    })
    
    if err := t.Run(); err != nil {
        log.Println(fmt.Errorf("error transport : %+v", err))
    }
}
```