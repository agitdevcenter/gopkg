# Custom Services
This is useful if you wants to add others thread blocking services like Kafka consumer for example. It can be anything actually, like Cron etc.

## Service
Must implement `custom.Service` interfaces.
```
package main

import (
    Logger "github.com/agitdevcenter/gopkg/logger"
)

type Example struct{
    logger Logger.Logger
    debug bool 
}

func New() *Example {
    return &Example{}
}

func (e *Example) SetLogger(logger Logger.Logger) {
    e.logger = logger
}

func (e *Example) SetDebug(enabled bool) {
    e.debug = enabled
}

func (e *Example) Start(ctx context.Context, wg *sync.WaitGroup) func() error {
    // must at least implement this
    return func() error {
        errorServer := make(chan error, 1)
        
        go func() {
            <-ctx.Done()
            // closing blocking thread
            // errorServer <- fmt.Errorf("error closing")
            close(errorServer)
            wg.Done()
        }()

        // you put your thread blocking process here
        
        err := <-errorServer
        wg.Wait()
        return err
    }
}
```

## Holder
Holds service and wait time.
```
package main

import (
    Logger "github.com/agitdevcenter/gopkg/logger"
    "github.com/agitdevcenter/gopkg/transport/custom"
    "github.com/agitdevcenter/gopkg/transport"
    "time"
    "fmt"
)

type Example struct{
    logger Logger.Logger
    debug bool 
}

func New() *Example {
    return &Example{}
}

func (e *Example) SetLogger(logger Logger.Logger) {
    e.logger = logger
}

func (e *Example) SetDebug(enabled bool) {
    e.debug = enabled
}

func (e *Example) Start(ctx context.Context, wg *sync.WaitGroup) func() error {
    // must at least implement this
    return func() error {
        errorServer := make(chan error, 1)
        
        go func() {
            <-ctx.Done()
            // closing blocking thread
            // errorServer <- fmt.Errorf("error closing")
            close(errorServer)
            wg.Done()
        }()

        // you put your thread blocking process here
        
        err := <-errorServer
        wg.Wait()
        return err
    }
}

func main() {
    service := New()
    holder := custom.New(custom.OptionService(service))
    
    t := transport.New(transport.OptionCustom(holder))
    
    if err := t.Run(); err != nil {
        fmt.ErrorF("error transport : %+v", err)
    }
}

func mainMultiService() {
    service1 := New()
    holder1 := custom.New(custom.OptionService(service1))
    
    service2 := New()
    // hold for 5 seconds before starting second service
    holder2 := custom.New(
        custom.OptionService(service2),
        custom.OptionHold(5 * time.Second),
    )

    t := transport.New(
        transport.OptionCustom(holder1),
        transport.OptionCustom(holder2),
    )
    
    if err := t.Run(); err != nil {
        fmt.ErrorF("error transport : %+v", err)
    }
}
```

## Working Example
There is working example using Kafka consumer that you can see [here](../example/kafka/main.go).