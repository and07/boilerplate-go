package grpc

import (
	"context"
	"net"
	"strconv"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

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

type GRPC struct {
	grpcSrv *grpc.Server
	address string
}

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

/*
func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	lis, err := net.Listen("tcp", ":8842")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)
	eventBus := make(chan interface{})
	gw.RegisterBlockchainServiceServer(grpcServer, service.NewBlockchainServer(eventBus))
	grpc_prometheus.Register(grpcServer)

	var group errgroup.Group

	group.Go(func() error {
		return grpcServer.Serve(lis)
	})

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(50000000)),
	}

	group.Go(func() error {
		return gw.RegisterBlockchainServiceHandlerFromEndpoint(ctx, mux, ":8842", opts)
	})
	group.Go(func() error {
		return http.ListenAndServe(":8843", wsproxy.WebsocketProxy(mux))
	})
	group.Go(func() error {
		return http.ListenAndServe(":2662", promhttp.Handler())
	})
	group.Go(func() error {
		for i := 0; i < 100; i++ {
			eventBus <- struct {
				Type             byte
				Coin             string
				Value            int
				TransactionCount int
				Timestamp        time.Time
			}{
				Type:             1,
				Coin:             "BIP",
				TransactionCount: i,
				Timestamp:        time.Now(),
			}
		}
		return nil
	})

	return group.Wait()
}
*/
