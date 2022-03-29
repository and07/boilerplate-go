package main

import (
	"context"
	"os"
	"time"

	"github.com/and07/boilerplate-go/api/gen-boilerplate-go/api"
	"github.com/and07/boilerplate-go/internal/app/boilerplate"
	"github.com/and07/boilerplate-go/internal/app/serv"
	log "github.com/and07/boilerplate-go/internal/pkg/logger"
	"github.com/and07/boilerplate-go/internal/pkg/template"
	"github.com/and07/boilerplate-go/internal/pkg/tracing"
	"github.com/and07/boilerplate-go/pkg/data"
	"github.com/and07/boilerplate-go/pkg/utils"
	"github.com/caarlos0/env"
	"github.com/gorilla/sessions"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"

	"github.com/markbates/goth"
	//"github.com/markbates/goth/providers/apple"
	//"github.com/markbates/goth/providers/facebook"

	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
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

	googleKEY := os.Getenv("GOOGLE_KEY")
	if googleKEY == "" {
		googleKEY = "328290909614-3casaiclr5c4ftspb91kun5ckl1av5om.apps.googleusercontent.com"
	}
	googleSECRET := os.Getenv("GOOGLE_SECRET")
	if googleSECRET == "" {
		googleSECRET = "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"
	}

	goth.UseProviders(
		//facebook.New(os.Getenv("FACEBOOK_KEY"), os.Getenv("FACEBOOK_SECRET"), "http://localhost:3000/auth/facebook/callback"),
		google.New(googleKEY, googleSECRET, "http://localhost:"+cfg.Port+"/auth/google/callback"),
		//apple.New(os.Getenv("APPLE_KEY"), os.Getenv("APPLE_SECRET"), "http://localhost:3000/auth/apple/callback", nil, apple.ScopeName, apple.ScopeEmail),
	)

	key := ""            // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30 // 30 days
	isProd := false      // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

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

	srv := serv.New(ctx,
		serv.WithPublicPort(cfg.Port),
		serv.WithDebugPort(cfg.PortDebug),
		serv.WithGRPC(ctx, cfg.PortGRPC, hw, bc),
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

	// create a new connection to the postgres db store
	db, err := data.NewConnection(configs, logger)
	if err != nil {
		logger.Error("unable to connect to db", "error", err)
		panic(err)
	}
	defer db.Close()

	srv.Run(ctx, publicHandle(ctx, tpl, db, configs, logger), api.RegisterBlockchainServiceHandlerFromEndpoint, api.RegisterHttpBodyExampleServiceHandlerFromEndpoint)

	log.Info("STOP APP")
}
