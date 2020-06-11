package custom

import (
	"context"
	Logger "github.com/agitdevcenter/gopkg/logger"
	"sync"
)

type Service interface {
	SetDebug(enabled bool)
	SetLogger(logger Logger.Logger)
	Start(ctx context.Context, wg *sync.WaitGroup) func() error
}
