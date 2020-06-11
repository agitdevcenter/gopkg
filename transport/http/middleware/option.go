package middleware

import (
	Logger "github.com/agitdevcenter/gopkg/logger"
)

type Option func(*Middleware)

func WithLogger(logger Logger.Logger) Option {
	return func(m *Middleware) {
		m.logger = logger
	}
}

func WithDebug(enabled bool) Option {
	return func(m *Middleware) {
		m.debug = enabled
	}
}

func WithProfiling(enabled bool) Option {
	return func(m *Middleware) {
		m.profiling = enabled
	}
}

func WithRecover(enabled bool) Option {
	return func(m *Middleware) {
		m.recover = enabled
	}
}

func WithCORS(enabled bool) Option {
	return func(m *Middleware) {
		m.cors = enabled
	}
}

func WithGZip(enabled bool) Option {
	return func(m *Middleware) {
		m.gzip = enabled
	}
}

func WithValidator(enabled bool) Option {
	return func(m *Middleware) {
		m.validator = enabled
	}
}

func WithErrorHandler(enabled bool) Option {
	return func(m *Middleware) {
		m.errorHandler = enabled
	}
}

func WithAcceptJSON(enabled bool) Option {
	return func(m *Middleware) {
		m.acceptJSON = enabled
	}
}

func WithSession(enabled bool, name, version string, port int) Option {
	return func(m *Middleware) {
		m.session = enabled
		m.name = name
		m.version = version
		m.port = port
	}
}

func WithInternalServerErrorMessage(internalServerErrorMessage string) Option {
	return func(m *Middleware) {
		m.internalServerErrorMessage = internalServerErrorMessage
	}
}

func WithSkip(urls []string) Option {
	return func(m *Middleware) {
		m.skipURLs = urls
	}
}

func WithHealth(url string) Option {
	return func(m *Middleware) {
		m.health = true
		m.healthURL = url
		m.skipURLs = append(m.skipURLs, m.healthURL)
	}
}

func WithAvailability(url string) Option {
	return func(m *Middleware) {
		m.availabilityEnabled = true
		m.availabilityURLPrefix = url
	}
}

func WithEndpointAvailability(urls []string) Option {
	return func(m *Middleware) {
		m.endpointAvailabilityURLs = urls
	}
}
