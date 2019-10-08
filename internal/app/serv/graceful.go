package serv

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	log "gitlab.com/and07/boilerplate-go/internal/pkg/logger"
)

func graceful(ctx context.Context, srvs ...*http.Server) chan struct{} {
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)

		// interrupt signal sent from terminal
		signal.Notify(sigint, os.Interrupt)
		// sigterm signal sent from kubernetes
		signal.Notify(sigint, syscall.SIGTERM)

		<-sigint

		// We received an interrupt signal, shut down.
		for i := range srvs {
			if err := srvs[i].Shutdown(ctx); err != nil {
				// Error from closing listeners, or context timeout:
				log.Errorf("HTTP server Shutdown %s: %v", srvs[i].Addr, err)
			}
		}

		close(idleConnsClosed)
		log.Info("HTTP server Shutdow")
	}()
	return idleConnsClosed
}
