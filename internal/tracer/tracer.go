package tracer

import (
	"github.com/uber/jaeger-client-go/config"
	"time"
)

func GetBaseConfig(serviceName string) *config.Configuration {
	return &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            false,
			BufferFlushInterval: 1 * time.Second,
		},
		ServiceName: serviceName,
	}
}
