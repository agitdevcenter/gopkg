package interceptor

import Logger "github.com/agitdevcenter/gopkg/logger"

type Option func(*Interceptor)

func WithDebug(enabled bool) Option {
	return func(i *Interceptor) {
		i.debug = enabled
	}
}

func WithLogger(logger Logger.Logger) Option {
	return func(i *Interceptor) {
		i.logger = logger
	}
}

func WithHandleCrash(enabled bool) Option {
	return func(i *Interceptor) {
		i.handleCrash = enabled
	}
}

func WithSession(enabled bool, name, version string, port int) Option {
	return func(i *Interceptor) {
		i.session = enabled
		i.name = name
		i.version = version
		i.port = port
	}
}

func WithInternalServerErrorMessage(internalServerErrorMessage string) Option {
	return func(i *Interceptor) {
		i.internalServerErrorMessage = internalServerErrorMessage
	}
}

func WithSkip(urls []string) Option {
	return func(i *Interceptor) {
		i.skipRPCs = urls
	}
}
