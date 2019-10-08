package serv

import (
	"context"
	"net/http"

	"github.com/opentracing/opentracing-go"
	log "gitlab.com/and07/boilerplate-go/internal/pkg/logger"
)

// Serv ...
type Serv struct {
	portPublicHTTP  string
	portPrivateHTTP string
}

// Option ...
type Option struct {
	PortPublicHTTP  string
	PortPrivateHTTP string
}

// New ...
func New(ctx context.Context, opt Option) *Serv {
	span, _ := opentracing.StartSpanFromContext(ctx, "NewServ")
	defer span.Finish()
	s := &Serv{
		portPublicHTTP:  opt.PortPublicHTTP,
		portPrivateHTTP: opt.PortPrivateHTTP,
	}
	return s
}

// Run ...
func (s *Serv) Run(ctx context.Context, handles *http.ServeMux) {
	span, _ := opentracing.StartSpanFromContext(ctx, "Serv.Run")
	defer span.Finish()

	srvPublic := s.runPublicHTTP(handles)
	srvPrivate := s.runPrivateHTTP()

	idleSrvConnsClosed := graceful(ctx, srvPrivate, srvPublic)
	go func() {
		log.Info("http.Private start")
		if errSrvPrivateListenAndServe := srvPrivate.ListenAndServe(); errSrvPrivateListenAndServe != http.ErrServerClosed {
			log.Errorf("HTTP server ListenAndServe: %v", errSrvPrivateListenAndServe)
		}
	}()
	go func() {
		log.Info("http.Public start")
		if errSrvPublicListenAndServe := srvPublic.ListenAndServe(); errSrvPublicListenAndServe != http.ErrServerClosed {
			log.Errorf("HTTP server ListenAndServe: %v", errSrvPublicListenAndServe)
		}
	}()
	<-idleSrvConnsClosed
}
