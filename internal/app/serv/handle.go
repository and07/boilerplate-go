package serv

import (
	"context"
	"net/http"
	"net/http/pprof"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// healthz is a liveness probe.
func healthz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// readyz is a readiness probe.
func readyz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func prometheusHandler() http.Handler {
	return promhttp.Handler()
}

func (s *Serv) privateHandle() *http.ServeMux {
	rPrivate := http.NewServeMux()
	rPrivate.Handle("/metrics", prometheusHandler())

	// pprof
	rPrivate.HandleFunc("/debug/pprof/", pprof.Index)
	rPrivate.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	rPrivate.HandleFunc("/debug/pprof/profile", pprof.Profile)
	rPrivate.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	rPrivate.HandleFunc("/debug/pprof/trace", pprof.Trace)

	// swagger
	if s.swaggerui {
		staticServer := http.FileServer(http.Dir("./assets/swaggerui"))
		sh := http.StripPrefix("/swaggerui/", staticServer)
		rPrivate.Handle("/swaggerui/", sh)
	}

	rPrivate.HandleFunc("/healthz", healthz)
	rPrivate.HandleFunc("/readyz", readyz)
	return rPrivate
}

func (s *Serv) runPrivateHTTP(ctx context.Context) *http.Server {
	srvPrivate := &http.Server{Addr: ":" + s.portPrivateHTTP, Handler: s.privateHandle()}
	go func() {
		s.Logger.Infof("http.Private start : %s", s.portPrivateHTTP)
		if errSrvPrivateListenAndServe := srvPrivate.ListenAndServe(); errSrvPrivateListenAndServe != nil && errSrvPrivateListenAndServe != http.ErrServerClosed {
			s.Logger.Errorf("HTTP server ListenAndServe: %v", errSrvPrivateListenAndServe)
		}
		<-ctx.Done()
	}()
	return srvPrivate
}

func (s *Serv) runPublicHTTP(ctx context.Context, h http.Handler) *http.Server {
	srvPublic := &http.Server{Addr: ":" + s.portPublicHTTP, Handler: h}
	go func() {

		s.Logger.Infof("http.Public start : %s", s.portPublicHTTP)
		if errSrvPublicListenAndServe := srvPublic.ListenAndServe(); errSrvPublicListenAndServe != nil && errSrvPublicListenAndServe != http.ErrServerClosed {
			s.Logger.Errorf("HTTP server ListenAndServe: %v", errSrvPublicListenAndServe)
		}
		<-ctx.Done()
	}()
	return srvPublic
}
