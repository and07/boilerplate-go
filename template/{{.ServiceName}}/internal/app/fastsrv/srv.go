package fastsrv

import (
	"context"
	"net/http"

	log "github.com/{{.User}}/{{.ServiceName}}/internal/pkg/logger"
	"github.com/opentracing/opentracing-go"
	"github.com/valyala/fasthttp"
)

// Fastsrv ...
type Fastsrv struct {
	portPublicHTTP  string
	portPrivateHTTP string
}

// Option ...
type Option struct {
	PortPublicHTTP  string
	PortPrivateHTTP string
}

// New ...
func New(ctx context.Context, opt Option) *Fastsrv {
	span, _ := opentracing.StartSpanFromContext(ctx, "NewFastsrv")
	defer span.Finish()
	f := &Fastsrv{
		portPublicHTTP:  opt.PortPublicHTTP,
		portPrivateHTTP: opt.PortPrivateHTTP,
	}
	return f
}

// Run ...
func (s *Fastsrv) Run(ctx context.Context, handles func(ctx *fasthttp.RequestCtx)) {
	span, _ := opentracing.StartSpanFromContext(ctx, "Fastsrv.Run")
	defer span.Finish()

	srvPublic := s.runPublicHTTP(handles)
	srvPrivate := s.runPrivateHTTP()

	idleSrvConnsClosed := graceful(srvPrivate, srvPublic)
	go func() {
		log.Info("http.Private start")
		if errSrvPrivateListenAndServe := srvPrivate.ListenAndServe(); errSrvPrivateListenAndServe != http.ErrServerClosed {
			log.Errorf("HTTP server ListenAndServe: %v", errSrvPrivateListenAndServe)
		}
	}()
	go func() {
		log.Info("http.Public start")
		if errSrvPublicListenAndServe := srvPublic.ListenAndServe(":" + s.portPublicHTTP); errSrvPublicListenAndServe != http.ErrServerClosed {
			log.Errorf("HTTP server ListenAndServe: %v", errSrvPublicListenAndServe)
		}
	}()
	<-idleSrvConnsClosed
}
