package scratch

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func graceful(srvs ...*http.Server) chan struct{} {
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
			if err := srvs[i].Shutdown(context.Background()); err != nil {
				// Error from closing listeners, or context timeout:
				log.Printf("HTTP server Shutdown %s: %v", srvs[i].Addr, err)
			}
		}

		close(idleConnsClosed)
		log.Println("HTTP server Shutdow")
	}()
	return idleConnsClosed
}
