# gRPC Interceptor
gRPC stream and unary interceptor

## Setup
There are several ways to setup interceptor.

## Default Setup
This will setup interceptor with default configuration.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/grpc/interceptor"
)

func main() {
    i := interceptor.New([]interceptor.Option{})
}
```

### With Options
Setting up interceptor with options from `interceptor.With`. 

#### Debug
`interceptor.WithDebug` `boolean` parameter. It will set the interceptor `debug` value.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/grpc/interceptor"
)

func main() {
    i := interceptor.New([]interceptor.Option{interceptor.WithDebug(true)})
}
```

#### Logger
`interceptor.WithLogger` `pkg/logger` parameter. It will set the interceptor `logger` value.
```go
package main

import (
    Logger "github.com/agitdevcenter/gopkg/logger"
    "github.com/agitdevcenter/gopkg/transport/grpc/interceptor"
    "time"
)

func main() {
    logger := Logger.New(Logger.Options{
        FileLocation:    "tdr.log",
        FileTdrLocation: "log.log",
        FileMaxAge:      time.Hour,
        Stdout:          true,
    })
    i := interceptor.New([]interceptor.Option{interceptor.WithLogger(logger)})
}
```

#### Handle Crash
`interceptor.WithHandleCrash` `boolean` parameter. It will set the interceptor `handlerCrash` value.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/grpc/interceptor"
)

func main() {
    i := interceptor.New([]interceptor.Option{interceptor.WithHandleCrash(true)})
}
```

#### Session
`interceptor.WithSession` session `boolean`, name, version `string`, port `int` parameters. It will set the interceptor request session. It needs application name, version, and port to generate request session.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/grpc/interceptor"
)

func main() {
    i := interceptor.New([]interceptor.Option{interceptor.WithSession(true, "LinkAja", "1.0.0", 80)})
}
```

#### Internal Server Error Message
`interceptor.WithInternalServerErrorMessage` `string` parameter. It will set the interceptor `internalServerErrorMessage` value.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/grpc/interceptor"
)

func main() {
    i := interceptor.New([]interceptor.Option{interceptor.WithInternalServerErrorMessage("Server Gangguan")})
}
```

#### Skip RPC
`interceptor.WithSkip` `[]string` parameter. It will set the interceptor `skipRPCs` value, to skip logging.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/grpc/interceptor"
)

func main() {
    i := interceptor.New([]interceptor.Option{interceptor.WithSkip([]string{"/", "", "health"})})
}
```

