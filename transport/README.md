# Transport
This package helps you build simple HTTP or gRPC server, or both, or any other thread blocker services.

## Example
For you that didn't bother to read this, please go [here](example) to see working examples.

## HTTP Server
This package HTTP Server using [echo](https://github.com/labstack/echo) version [4](https://github.com/labstack/echo/releases) as HTTP framework. For more detailed information you can read [here](http/README.md), or you can see the working example [here](example/http)
.
### Setup HTTP Server
There are several ways to setup HTTP server.

#### Using Option
This will start HTTP server at port 2202.
```go
package main

import (
    "fmt"
    "github.com/agitdevcenter/gopkg/transport"
    "github.com/agitdevcenter/gopkg/transport/http"
    "log"
)

func main() {
    h := http.New([]http.Option{})

    t := transport.New([]transport.Option{
        transport.WithHTTPServer(h),  
    })
    
    if err := t.Run(); err != nil {
        log.Println(fmt.Errorf("error transport : %+v", err))
    }
}
```

#### Using SetHTTP method
This will start HTTP server at port 2202.
```go
package main

import (
    "fmt"
    "github.com/agitdevcenter/gopkg/transport"
    "github.com/agitdevcenter/gopkg/transport/http"
    "log"
)

func main() {
    h := http.New([]http.Option{})
    
    t := transport.New([]transport.Option{})
    
    t.SetHTTP(h)
    
    if err := t.Run(); err != nil {
        log.Println(fmt.Errorf("error transport : %+v", err))
    }   
}
```

#### Using SetupHTTP method
This will start HTTP server at port 2202. Method SetupHTTP accepts `http.Option` parameters. For more information about `http.Option` read [here](http/README.md).
```go
package main

import (
    "fmt"
    "github.com/agitdevcenter/gopkg/transport"
    "github.com/agitdevcenter/gopkg/transport/http"
    "log"
)

func main() {
    t := transport.New([]transport.Option{})
    
    t.SetupHTTP([]http.Option{})
    
    if err := t.Run(); err != nil {
        log.Println(fmt.Errorf("error transport : %+v", err))
    }
}
```

## gRPC Server
gRPC server. For more detailed information you can read [here](grpc/README.md), or you can see the working example [here](example/grpc)
.
### Setup gRPC Server
There are several ways to setup gRPC server.

#### Using Option
This will start gRPC server at port 2202.
```go
package main

import (
    "fmt"
    "github.com/agitdevcenter/gopkg/transport"
    "github.com/agitdevcenter/gopkg/transport/grpc"
    "log"
)

func main() {
    g := grpc.New([]grpc.Option{})
    
    t := transport.New([]transport.Option{transport.WithGRPCServer(g)})
    
    if err := t.Run(); err != nil {
        log.Println(fmt.Errorf("error transport : %+v", err))
    }
}
```

#### Using SetGRPC method
This will start gRPC server at port 2202.
```go
package main

import (
    "fmt"
    "github.com/agitdevcenter/gopkg/transport"
    "github.com/agitdevcenter/gopkg/transport/grpc"
    "log"
)

func main() {
    g := grpc.New([]grpc.Option{})
    
    t := transport.New([]transport.Option{})
    
    t.SetGRPC(g)
    
    if err := t.Run(); err != nil {
        log.Println(fmt.Errorf("error transport : %+v", err))
    }   
}
```

#### Using SetupGRPC method
This will start gRPC server at port 2202. Method SetupGRPC accepts `grpc.Option` parameters. For more information about `grpc.Option` read [here](grpc/README.md).
```go
package main

import (
    "fmt"
    "github.com/agitdevcenter/gopkg/transport"
    "github.com/agitdevcenter/gopkg/transport/grpc"
    "log"
)

func main() {
    t := transport.New([]transport.Option{})
    
    t.SetupGRPC([]grpc.Option{})
    
    if err := t.Run(); err != nil {
        log.Println(fmt.Errorf("error transport : %+v", err))
    }
}
```

## Options
Available options using `transport.Option`

### Logger
`transport.WithLogger` `pkg/logger` parameter. Will set logger value for `transport`.
```go
package main

import (
    Logger "github.com/agitdevcenter/gopkg/logger"
    "github.com/agitdevcenter/gopkg/transport"
    "time"
)

func main() {
    logger := Logger.New(Logger.Options{
        FileLocation:    "tdr.log",
        FileTdrLocation: "log.log",
        FileMaxAge:      time.Hour,
        Stdout:          true,
    })

    t := transport.New([]transport.Option{transport.WithLogger(logger)})
}
```

### Debug
`transport.WithDebug` `boolean` parameter. Will set `debug` value for `transport`.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport"
)

func main() {
    t := transport.New([]transport.Option{transport.WithDebug(true)})
}
```

### HTTPServer
`transport.WithHTTPServer` `http.Server` parameter. Will set `httpServer` value for `transport`.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport"
    "github.com/agitdevcenter/gopkg/transport/http"
)

func main() {
    h := http.New([]http.Option{})

    t := transport.New([]transport.Option{transport.WithHTTPServer(h)})
}
```

### GRPCServer
`transport.WithGRPCServer` `grpc.Server` parameter. Will set `grpcServer` value for `transport`. @Unimplemented
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport"
    "github.com/agitdevcenter/gopkg/transport/grpc"
)

func main() {
    g := grpc.New()

    t := transport.New([]transport.Option{transport.WithGRPCServer(g)})
}
```

### Inherit
`transport.WithInherit` `boolean` parameter. HTTP or gRPC server, or both and any other services will inherit `debug` and `logger` value from `transport`
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport"
)

func main() {
    t := transport.New([]transport.Option{transport.WithInherit(false)})
}
```

### Custom Service
Adding thread blocking process service, you can read more detailed information [here](custom/README.md).

### Multiple Options
This example will start HTTP Server at port 2022 with `debug` and `logger` set up for `transport` and inherited to `http.Server`
```go
package main

import (
    "fmt"
    Logger "github.com/agitdevcenter/gopkg/logger"
    "github.com/agitdevcenter/gopkg/transport"
    "github.com/agitdevcenter/gopkg/transport/http"
    "log"
    "time"
)

func main() {
    h := http.New([]http.Option{})

    logger := Logger.New(Logger.Options{
        FileLocation:    "tdr.log",
        FileTdrLocation: "log.log",
        FileMaxAge:      time.Hour,
        Stdout:          true,
    })

    t := transport.New([]transport.Option{
        transport.WithHTTPServer(h),
        transport.WithLogger(logger),
        transport.WithDebug(true),
        transport.WithInherit(true),
    })
    
    if err := t.Run(); err != nil {
        log.Println(fmt.Errorf("error transport : %+v", err))
    }

}
```
