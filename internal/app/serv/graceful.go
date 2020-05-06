package serv

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (s *Serv) graceful(ctx context.Context, fn context.CancelFunc, srvs ...*http.Server) {

	go func() {
		sigint := make(chan os.Signal, 1)

		// interrupt signal sent from terminal
		// sigterm signal sent from kubernetes
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)

		<-sigint

		// We received an interrupt signal, shut down.
		timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		for i := range srvs {
			if err := srvs[i].Shutdown(timeoutCtx); err != nil {
				// Error from closing listeners, or context timeout:
				s.Logger.Errorf("HTTP server Shutdown %s: %v", srvs[i].Addr, err)
			}
			s.Logger.Infof("HTTP server Shutdow %s", srvs[i].Addr)
		}

		fn()

	}()
	return
}
