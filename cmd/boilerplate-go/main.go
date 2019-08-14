package main

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"gitlab.com/and07/boilerplate-go/internal/app/scratch"
	log "gitlab.com/and07/boilerplate-go/internal/pkg/logger"
	"gitlab.com/and07/boilerplate-go/internal/pkg/tracing"
)

const (
	publicPort  = 8080
	privatePort = 8888
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
	srv := scratch.New(ctx, scratch.Option{PortPublicHTTP: publicPort, PortPrivateHTTP: privatePort})
	srv.Run(ctx, publicHandle(ctx))

}
