package fastsrv

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	log "gitlab.com/seobutik-dsp/dsp-proxy-new/internal/pkg/logger"
	"github.com/valyala/fasthttp"
)

func graceful(srv *http.Server, fsrv *fasthttp.Server) chan struct{} {
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)

		// interrupt signal sent from terminal
		signal.Notify(sigint, os.Interrupt)
		// sigterm signal sent from kubernetes
		signal.Notify(sigint, syscall.SIGTERM)

		<-sigint

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Errorf("HTTP server Shutdown : %v", err)
		}
	
		if err := fsrv.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Errorf("FastHTTP server Shutdown : %v", err)
		}
		close(idleConnsClosed)
		log.Info("HTTP server Shutdow")
	}()
	return idleConnsClosed
}
