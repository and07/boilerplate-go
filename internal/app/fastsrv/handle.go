package fastsrv

import (
	"net/http"
	"net/http/pprof"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp"
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

func privateHandle() *http.ServeMux {
	rPrivate := http.NewServeMux()
	rPrivate.Handle("/metrics", prometheusHandler())

	// pprof
	rPrivate.HandleFunc("/debug/pprof/", pprof.Index)
	rPrivate.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	rPrivate.HandleFunc("/debug/pprof/profile", pprof.Profile)
	rPrivate.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	rPrivate.HandleFunc("/debug/pprof/trace", pprof.Trace)

	rPrivate.HandleFunc("/healthz", healthz)
	rPrivate.HandleFunc("/readyz", readyz)
	return rPrivate
}

func (s *Fastsrv) runPrivateHTTP() *http.Server {
	srvPrivate := &http.Server{Addr: ":" + s.portPrivateHTTP, Handler: privateHandle()}
	return srvPrivate
}

func (s *Fastsrv) runPublicHTTP(h func(ctx *fasthttp.RequestCtx)) *fasthttp.Server {
	srvPublic := &fasthttp.Server{
		Handler: h,
	}
	return srvPublic
}
