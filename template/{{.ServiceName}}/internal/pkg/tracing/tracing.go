package tracing

import (
	"fmt"
	"io"
	"time"

	log "github.com/{{.User}}/{{.ServiceName}}/internal/pkg/logger"
	opentracing "github.com/opentracing/opentracing-go"
	config "github.com/uber/jaeger-client-go/config"
)

// Init returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func Init(serviceName string) (opentracing.Tracer, io.Closer) {
	cfg, err := config.FromEnv()
	if err != nil {
		log.Fatalf("cannot parse Jaeger env vars %s", err)
	}
	cfg.ServiceName = serviceName
	cfg.Sampler.Type = "ratelimiting"
	cfg.Sampler.Param = 1

	// TODO(ys) a quick hack to ensure random generators get different seeds, which are based on current time.
	time.Sleep(100 * time.Millisecond)

	tracer, closer, err := cfg.NewTracer(config.Logger(log.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}
