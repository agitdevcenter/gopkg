package tracer

import (
	"github.com/opentracing/opentracing-go"
	"io"
	"time"

	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics/prometheus"
)

func New(serviceName string) (opentracing.Tracer, io.Closer, error) {
	defcfg := config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}
	cfg, _ := defcfg.FromEnv()
	cfg.ServiceName = serviceName
	return cfg.NewTracer(
		config.Metrics(prometheus.New()),
	)
}
