package main

import (
	"context"
	"time"

	api "github.com/and07/boilerplate-go/api/gen-boilerplate-go"
	apiTrening "github.com/and07/boilerplate-go/api/trening"
	"github.com/and07/boilerplate-go/configs"
	"github.com/and07/boilerplate-go/internal/app/boilerplate"
	"github.com/and07/boilerplate-go/internal/app/serv"
	log "github.com/and07/boilerplate-go/internal/pkg/logger"
	"github.com/and07/boilerplate-go/internal/pkg/template"
	"github.com/and07/boilerplate-go/internal/pkg/tracing"
	"github.com/and07/boilerplate-go/pkg/data"
	"github.com/and07/boilerplate-go/pkg/token"
	"github.com/and07/boilerplate-go/pkg/utils"
	"github.com/caarlos0/env"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	//"github.com/markbates/goth/providers/apple"
	//"github.com/markbates/goth/providers/facebook"
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

	cfg := &configs.Configs{}
	if err := env.Parse(cfg); err != nil {
		log.Error(err)
	}

	tpl := template.NewTemplate("tpl/layouts/", "tpl/", `{{define "main" }} {{ template "base" . }} {{ end }}`)
	tpl.Init()

	hw := func(grpcSrv *grpc.Server) {
		impl := boilerplate.New(ctx)
		api.RegisterHttpBodyExampleServiceServer(grpcSrv, impl)
	}

	eventBus := make(chan interface{})

	bc := func(grpcSrv *grpc.Server) {
		api.RegisterBlockchainServiceServer(grpcSrv, boilerplate.NewBlockchainServer(eventBus))
	}

	ts := func(grpcSrv *grpc.Server) {
		apiTrening.RegisterTreningServiceServer(grpcSrv, apiTrening.NewtTeningFacade(token.NewExtractor()))
	}

	srv := serv.New(ctx,
		serv.WithPublicPort(cfg.Port),
		serv.WithDebugPort(cfg.PortDebug),
		serv.WithGRPC(ctx, cfg, hw, bc, ts),
		serv.WithSwaggerUI(true),
	)

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

	logger := utils.NewLogger()

	configs := utils.NewConfigurations(logger)

	var db *sqlx.DB
	var err error
	if configs.DBConn != "" || configs.DBHost != "" {
		// create a new connection to the postgres db store
		db, err = data.NewConnection(configs, logger)
		if err != nil {
			logger.Error("unable to connect to db", "error", err)
			panic(err)
		}
		defer db.Close()
	}

	srv.Run(ctx, publicHandle(ctx, tpl, db, configs, logger), logger, api.RegisterBlockchainServiceHandlerFromEndpoint, api.RegisterHttpBodyExampleServiceHandlerFromEndpoint, apiTrening.RegisterTreningServiceHandlerFromEndpoint)

	log.Info("STOP APP")
}
