# HTTP Middleware
HTTP Middleware using [echo](https://github.com/labstack/echo) version [4](https://github.com/labstack/echo/releases) [middleware](https://echo.labstack.com/middleware).

## Setup
There are several ways to setup middleware.

## Default Setup
This will setup middleware with default configuration.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/http/middleware"
)

func main() {
    m := middleware.New([]middleware.Option{})
}
```

### With Options
Setting up middleware with options from `middleware.Option`. 

#### Debug
`middleware.WithDebug` `boolean` parameter. It will set the middleware `debug` value.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/http/middleware"
)

func main() {
    m := middleware.New([]middleware.Option{middleware.WithDebug(true)})
}
```

#### Logger
`middleware.WithLogger` `pkg/logger` parameter. It will set the middleware `logger` value.
```go
package main

import (
    Logger "github.com/agitdevcenter/gopkg/logger"
    "github.com/agitdevcenter/gopkg/transport/http/middleware"
    "time"
)

func main() {
    logger := Logger.New(Logger.Options{
        FileLocation:    "tdr.log",
        FileTdrLocation: "log.log",
        FileMaxAge:      time.Hour,
        Stdout:          true,
    })
    m := middleware.New([]middleware.Option{middleware.WithLogger(logger)})
}
```

#### Recover
`middleware.WithRecover` `boolean` parameter. It will set the middleware `recover` value.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/http/middleware"
)

func main() {
    m := middleware.New([]middleware.Option{middleware.WithRecover(true)})
}
```

#### CORS
`middleware.WithCORS` `boolean` parameter. It will set the middleware `cors` value.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/http/middleware"
)

func main() {
    m := middleware.New([]middleware.Option{middleware.WithCORS(true)})
}
```

#### GZip
`middleware.WithGZip` `boolean` parameter. It will set the middleware `gzip` value.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/http/middleware"
)

func main() {
    m := middleware.New([]middleware.Option{middleware.WithGZip(true)})
}
```

#### Validator
`middleware.WithValidator` `boolean` parameter. It will set the middleware `validator` value.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/http/middleware"
)

func main() {
    m := middleware.New([]middleware.Option{middleware.WithValidator(true)})
}
```

#### Error Handler
`middleware.WithErrorHandler` `boolean` parameter. It will set the middleware `errorHandler` value.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/http/middleware"
)

func main() {
    m := middleware.New([]middleware.Option{middleware.WithErrorHandler(true)})
}
```

#### Accept JSON
`middleware.WithAcceptJSON` `boolean` parameter. It will set the middleware `acceptJSON` value.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/http/middleware"
)

func main() {
    m := middleware.New([]middleware.Option{middleware.WithAcceptJSON(true)})
}
```

#### Session
`middleware.WithSession` session `boolean`, name, version `string`, port `int` parameters. It will set the middleware request session. It needs application name, version, and port to generate request session.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/http/middleware"
)

func main() {
    m := middleware.New([]middleware.Option{middleware.WithSession(true, "LinkAja", "1.0.0", 80)})
}
```

#### Internal Server Error Message
`middleware.WithInternalServerErrorMessage` `string` parameter. It will set the middleware `internalServerErrorMessage` value.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/http/middleware"
)

func main() {
    m := middleware.New([]middleware.Option{middleware.WithInternalServerErrorMessage("Server Gangguan")})
}
```

#### Skip URL
`middleware.WithSkip` `[]string` parameter. It will set the middleware `skipURLs` value, to skip logging.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/http/middleware"
)

func main() {
    m := middleware.New([]middleware.Option{middleware.WithSkip([]string{"/", "", "health"})})
}
```

#### Health
`middleware.WithHealth` `string` parameter. It will set the middleware `healthURL` value, to enable health check.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/http/middleware"
)

func main() {
    m := middleware.New([]middleware.Option{middleware.WithHealth("/sehat")})
}
```

#### Availability
`middleware.WithAvailability` `string` parameter. It will set the middleware `availabilityURLPrefix` value, to enable availability.

**Limitation** : only for single instance, let me think about it again later.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/http/middleware"
)

func main() {
    m := middleware.New([]middleware.Option{middleware.WithAvailability("/toggle")})
}
```

#### Endpoint Availability
`middleware.WithEndpointAvailability` `[]string` parameter. It will set the middleware `endpointAvailabilityURLs` value, to enable endpoint specific availability.

**Limitation** : only for single instance, let me think about it again later.
```go
package main

import (
    "github.com/agitdevcenter/gopkg/transport/http/middleware"
)

func main() {
    m := middleware.New([]middleware.Option{middleware.WithEndpointAvailability([]string{"/ping/error"})})
}
```

