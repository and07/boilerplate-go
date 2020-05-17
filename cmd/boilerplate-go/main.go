package main

import (
	"context"
	"time"

	"github.com/and07/boilerplate-go/api/gen-boilerplate-go/api"
	"github.com/and07/boilerplate-go/internal/app/boilerplate"
	"github.com/and07/boilerplate-go/internal/app/serv"
	log "github.com/and07/boilerplate-go/internal/pkg/logger"
	"github.com/and07/boilerplate-go/internal/pkg/template"
	"github.com/and07/boilerplate-go/internal/pkg/tracing"
	"github.com/caarlos0/env"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
)

const (
	pPublicPort  = "8080"
	pPrivatePort = "8888"
)

var counter prometheus.Counter

func init() {
	counter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "hi_handler",
		})
	prometheus.MustRegister(counter)
}

func main() {
	log.Info("Start APP")

	ctx := context.Background()

	tracer, closer := tracing.Init("boilerplate-go")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	span := tracer.StartSpan("Main")
	//span.SetTag("hello-to", helloTo)
	defer span.Finish()

	ctx = opentracing.ContextWithSpan(ctx, span)
	/*
		redisHost := os.Getenv("REDIS_HOST")
		if redisHost == "" {
			redisHost = "localhost:6379"
		}

		r := redis.New(ctx, redisHost)
		defer r.Close()

		rabbitmqConnectionString := os.Getenv("RABBIT_CONNECTION_STRING")
		if rabbitmqConnectionString == "" {
			rabbitmqConnectionString = "amqp://guest:guest@localhost:5672"
		}

		clRabbit, err := rabbitmq.New(ctx, rabbitmqConnectionString)
		if err != nil {
			log.Printf("rabbitmq.New: %v", err)
		}
		if clRabbit != nil {
			defer clRabbit.Close()
		}

		//clRabbit.Publish(ctx,[]byte("test msg"), *exchange, *exchangeType)

		if errSubscribe := clRabbit.Subscribe(ctx, *exchange, *exchangeType, *queue, func(d amqp.Delivery) {
			log.Printf("Delivery: %#v", d)
		}); errSubscribe != nil {
			log.Printf("errSubscribe: %v", errSubscribe)
		}

		clickhouseURL := os.Getenv("CLICKHOUSE_URL")
		if clickhouseURL == "" {
			clickhouseURL = "tcp://localhost:9000?debug=true"
		}
		clickhouse.NewClickhouseClient(ctx, clickhouseURL)

		elasticURL := os.Getenv("ELASTIC_URL")
		if elasticURL == "" {
			elasticURL = "http://127.0.0.1:9200"
		}
		elastic.NewElasticClient(ctx, elasticURL)
	*/

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		log.Error(err)
	}

	if cfg.Port == "" {
		cfg.Port = pPublicPort
	}

	if cfg.PortDebug == "" {
		cfg.PortDebug = pPrivatePort
	}

	if cfg.PortGRPC == "" {
		cfg.PortGRPC = "8842"
	}

	tpl := template.NewTemplate("tpl/layouts/", "tpl/", `{{define "main" }} {{ template "base" . }} {{ end }}`)
	tpl.Init()

	srv := serv.New(ctx,
		serv.WithPublicPort(cfg.Port),
		serv.WithDebugPort(cfg.PortDebug),
		serv.WithGRPCPort(cfg.PortGRPC),
		serv.WithSwaggerUI(true),
	)

	hw := func(grpcSrv *grpc.Server) {
		impl := boilerplate.New(ctx)
		api.RegisterHttpBodyExampleServiceServer(grpcSrv, impl)
	}

	eventBus := make(chan interface{})

	bc := func(grpcSrv *grpc.Server) {
		api.RegisterBlockchainServiceServer(grpcSrv, boilerplate.NewBlockchainServer(eventBus))
	}

	//test
	go func() {
		time.Sleep(5 * time.Second)

		for i := 0; 1 < 100; i++ {
			eventBus <- struct {
				Type             byte
				Coin             string
				Value            int
				TransactionCount int
				Timestamp        time.Time
			}{
				Type:             1,
				Coin:             "BIP",
				TransactionCount: i,
				Timestamp:        time.Now(),
			}
			time.Sleep(1 * time.Second)
		}
	}()

	srv.Run(ctx, publicHandle(ctx, tpl), hw, bc)

}
