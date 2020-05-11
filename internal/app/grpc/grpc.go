package grpc

import (
	"context"
	"net"
	"net/http"
	"strconv"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/tmc/grpc-websocket-proxy/wsproxy"

	//"github.com/and07/boilerplate-go/service"
	_ "github.com/jnewmano/grpc-json-proxy/codec" // GRPC Proxy
	errch "github.com/proxeter/errors-channel"
	log "gitlab.com/and07/boilerplate-go/internal/pkg/logger"
	"go.elastic.co/apm/module/apmgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

const defaultGRPCPort = 8842

// GRPC ...
type GRPC struct {
	grpcSrv *grpc.Server
	address string
}

// NewServer ...
func NewServer(ctx context.Context, GRPCPort string) *GRPC {
	var address string
	if GRPCPort != "" {
		address = ":" + GRPCPort
	} else {
		address = ":" + strconv.Itoa(defaultGRPCPort)
	}

	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_prometheus.UnaryServerInterceptor,
				apmgrpc.NewUnaryServerInterceptor(apmgrpc.WithRecovery()),
			),
		),
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
	)

	return &GRPC{
		grpcSrv: server,
		address: address,
	}
}

// RegisterEndpoints ...
func (g *GRPC) RegisterEndpoints(ctx context.Context, RegisterEndpointFns ...func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error) error {
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(50000000)),
	}

	headers := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

			h.ServeHTTP(w, r)
		})
	}

	for _, fn := range RegisterEndpointFns {
		if err := fn(ctx, mux, ":8842", opts); err != nil {
			log.Error(err)
		}
	}

	select {
	case err := <-errch.Register(func() error { return http.ListenAndServe(":8843", wsproxy.WebsocketProxy(headers(mux))) }):
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RunGRPC ...
func (g *GRPC) RunGRPC(ctx context.Context, fns ...func(*grpc.Server)) error {

	for _, fn := range fns {
		// Call the option giving the instantiated
		// *Serv as the argument
		fn(g.grpcSrv)
	}

	conn, err := net.Listen("tcp4", g.address)
	if err != nil {
		log.Fatal("error while listen socket for grpc service ", zap.Error(err))
	}

	healthpb.RegisterHealthServer(g.grpcSrv, health.NewServer())

	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(g.grpcSrv)

	reflection.Register(g.grpcSrv)

	log.Info("Start grpc server", zap.String("address", g.address))
	select {
	case err := <-errch.Register(func() error { return g.grpcSrv.Serve(conn) }):
		return err
	case <-ctx.Done():
		log.Infof("Shutdown grpc server %s", g.address)
		g.grpcSrv.GracefulStop()

		return ctx.Err()
	}
}
