package main

import (
	"context"
	"time"

	api "github.com/and07/boilerplate-go/api/gen-boilerplate-go"
	apiTrening "github.com/and07/boilerplate-go/api/trening"
	"github.com/and07/boilerplate-go/configs"
	"github.com/and07/boilerplate-go/internal/app/boilerplate"
	"github.com/and07/boilerplate-go/internal/app/serv"
	trening "github.com/and07/boilerplate-go/internal/app/trening/handlers"
	log "github.com/and07/boilerplate-go/internal/pkg/logger"
	"github.com/and07/boilerplate-go/internal/pkg/storage/s3"
	"github.com/and07/boilerplate-go/internal/pkg/template"
	"github.com/and07/boilerplate-go/internal/pkg/tracing"
	"github.com/and07/boilerplate-go/pkg/data"
	"github.com/and07/boilerplate-go/pkg/service"
	"github.com/and07/boilerplate-go/pkg/token"
	"github.com/and07/boilerplate-go/pkg/utils"
	"github.com/caarlos0/env/v6"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	//"github.com/markbates/goth/providers/apple"
	//"github.com/markbates/goth/providers/facebook"
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
		return
	}

	cfgS3 := &configs.S3Configs{}
	if err := env.Parse(cfgS3); err != nil {
		log.Error(err)
		return
	}

	tpl := template.NewTemplate("tpl/layouts/", "tpl/", `{{define "main" }} {{ template "base" . }} {{ end }}`)
	tpl.Init()

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

	// authService contains all methods that help in authorizing a user request
	authService := service.NewAuthService(logger, configs)

	// repository contains all the methods that interact with DB to perform CURD operations for user.
	repositoryAuth := data.NewAuthPostgresRepository(db, logger)
	//repositoryAuth := data.NewAuthMemoryRepository(logger)

	treningRepository := data.NewTreningPostgresRepository(db, logger)

	hw := func(grpcSrv *grpc.Server) {
		impl := boilerplate.New(ctx)
		api.RegisterHttpBodyExampleServiceServer(grpcSrv, impl)
	}

	eventBus := make(chan interface{})

	bc := func(grpcSrv *grpc.Server) {
		api.RegisterBlockchainServiceServer(grpcSrv, boilerplate.NewBlockchainServer(eventBus))
	}

	uploader := s3.NewFileUploader(s3.UploaderCfg{
		Region:    cfgS3.Region,
		Endpoint:  cfgS3.Host,
		AccessKey: cfgS3.AccessKey,
		SecretKey: cfgS3.SecretKey,
	})

	treningHandler := trening.NewTreningHandler(treningRepository, uploader, logger)
	facadeTrening := apiTrening.NewTeningFacade(token.NewExtractor(), authService, treningHandler, logger)
	ts := func(grpcSrv *grpc.Server) {
		apiTrening.RegisterTreningServiceServer(grpcSrv, facadeTrening)
	}

	srv := serv.New(ctx,
		serv.WithPublicPort(cfg.Port),
		serv.WithDebugPort(cfg.PortDebug),
		serv.WithGRPC(ctx, cfg, facadeTrening, hw, bc, ts),
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

	srv.Run(
		ctx,
		publicHandle(ctx, tpl, repositoryAuth, configs, authService, logger),
		logger,
		api.RegisterBlockchainServiceHandlerFromEndpoint,
		api.RegisterHttpBodyExampleServiceHandlerFromEndpoint,
		apiTrening.RegisterTreningServiceHandlerFromEndpoint,
	)

	log.Info("STOP APP")
}
