package serv

import (
	"context"
	"net/http"

	"github.com/and07/boilerplate-go/api/gen-boilerplate-go/api"
	g "github.com/and07/boilerplate-go/internal/app/grpc"
	log "github.com/and07/boilerplate-go/internal/pkg/logger"
	"github.com/opentracing/opentracing-go"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

// Logger ...
type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

// Serv ...
type Serv struct {
	portPublicHTTP  string
	portPrivateHTTP string
	portGRPC        string
	swaggerui       bool
	Logger          Logger
}

// Option ...
type Option func(*Serv)

// WithSwaggerUI ...
func WithSwaggerUI(on bool) Option {
	return func(s *Serv) {
		s.swaggerui = on
	}
}

// WithDebugPort ...
func WithDebugPort(portPrivateHTTP string) Option {
	return func(s *Serv) {
		s.portPrivateHTTP = portPrivateHTTP
	}
}

// WithPublicPort ...
func WithPublicPort(portPublicHTTP string) Option {
	return func(s *Serv) {
		s.portPublicHTTP = portPublicHTTP
	}
}

// WithGRPCPort ...
func WithGRPCPort(portGRPC string) Option {
	return func(s *Serv) {
		s.portGRPC = portGRPC
	}
}

// WithLogger ...
func WithLogger(l Logger) Option {
	return func(s *Serv) {
		s.Logger = l
	}
}

// New ...
func New(ctx context.Context, opts ...Option) *Serv {
	span, _ := opentracing.StartSpanFromContext(ctx, "NewServ")
	defer span.Finish()
	s := &Serv{
		Logger: log.New(),
	}

	for _, opt := range opts {
		// Call the option giving the instantiated
		// *Serv as the argument
		opt(s)
	}
	return s
}

// Run ...
func (s *Serv) Run(ctx context.Context, handles *http.ServeMux, fns ...func(*grpc.Server)) {

	ctx, cancel := context.WithCancel(ctx)

	span, _ := opentracing.StartSpanFromContext(ctx, "Serv.Run")
	defer span.Finish()

	servs := make([]*http.Server, 2)
	if s.portPublicHTTP != "" {
		servs[0] = s.runPublicHTTP(ctx, handles)
	}
	if s.portPrivateHTTP != "" {
		servs[1] = s.runPrivateHTTP(ctx)
	}

	if s.portGRPC != "" {
		var group errgroup.Group
		grpcSrv := g.NewServer(ctx, s.portGRPC)
		group.Go(func() error {
			return grpcSrv.RunGRPC(ctx, fns...)
		})
		group.Go(func() error {
			return grpcSrv.RegisterEndpoints(ctx, api.RegisterBlockchainServiceHandlerFromEndpoint, api.RegisterHttpBodyExampleServiceHandlerFromEndpoint)
		})

		if err := group.Wait(); err != nil {
			log.Error(err)
		}
	}

	s.graceful(ctx, cancel, servs...)

	<-ctx.Done()
}
