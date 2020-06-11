package logger

import "go.uber.org/zap"

type noop struct{}

func Noop() Logger {
	return &noop{}
}

func (n *noop) Debug(message string, fields ...zap.Field) {}

func (n *noop) Info(message string, fields ...zap.Field) {}

func (n *noop) Warn(message string, fields ...zap.Field) {}

func (n *noop) Error(message string, fields ...zap.Field) {}

func (n *noop) Fatal(message string, fields ...zap.Field) {}

func (n *noop) Panic(message string, fields ...zap.Field) {}

func (n *noop) TDR(tdr LogTdrModel) {}
