package serv

import (
	"context"
	"net/http"

	"github.com/opentracing/opentracing-go"
	g "gitlab.com/and07/boilerplate-go/internal/app/grpc"
	"google.golang.org/grpc"
)

// Serv ...
type Serv struct {
	portPublicHTTP  string
	portPrivateHTTP string
	portGRPC        string
}

// ServOption ...
type ServOption func(*Serv)

func WithDebugPort(portPrivateHTTP string) ServOption {
	return func(s *Serv) {
		s.portPrivateHTTP = portPrivateHTTP
	}
}

func WithPublicPort(portPublicHTTP string) ServOption {
	return func(s *Serv) {
		s.portPublicHTTP = portPublicHTTP
	}
}

func WithGRPCPort(portGRPC string) ServOption {
	return func(s *Serv) {
		s.portGRPC = portGRPC
	}
}

// New ...
func New(ctx context.Context, opts ...ServOption) *Serv {
	span, _ := opentracing.StartSpanFromContext(ctx, "NewServ")
	defer span.Finish()
	s := &Serv{}

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
		go g.NewServer(ctx, s.portGRPC).RunGRPC(ctx, fns...)
	}

	graceful(ctx, cancel, servs...)

	<-ctx.Done()
}
