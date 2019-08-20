package serv

import (
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

func (s *Serv) runPrivateHTTP() *http.Server {
	srvPrivate := &http.Server{Addr: ":" + s.portPrivateHTTP, Handler: privateHandle()}
	return srvPrivate
}

func (s *Serv) runPublicHTTP(h *http.ServeMux) *http.Server {
	srvPublic := &http.Server{Addr: ":" + s.portPublicHTTP, Handler: h}
	return srvPublic
}
