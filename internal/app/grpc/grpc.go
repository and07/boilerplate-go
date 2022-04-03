package grpc

import (
	"bytes"
	"context"
	"net"
	"net/http"
	"os"
	"strconv"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hashicorp/go-hclog"
	"golang.org/x/sync/errgroup"

	//"github.com/and07/boilerplate-go/service"
	"github.com/and07/boilerplate-go/configs"
	log "github.com/and07/boilerplate-go/internal/pkg/logger"
	_ "github.com/jnewmano/grpc-json-proxy/codec" // GRPC Proxy
	"go.elastic.co/apm/module/apmgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

const defaultGRPCPort = 8842

// GRPC ...
type GRPC struct {
	grpcSrv *grpc.Server
	address string
	cfg     *configs.Configs
}

// NewServer ...
func NewServer(ctx context.Context, cfg *configs.Configs) *GRPC {
	simpleLogger, err := zap.NewDevelopment()
	if err != nil {
		os.Exit(1)
	}
	var address string
	if cfg.PortGrpc != "" {
		address = ":" + cfg.PortGrpc
	} else {
		address = ":" + strconv.Itoa(defaultGRPCPort)
	}

	grpc_prometheus.EnableClientHandlingTimeHistogram()
	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc_middleware.WithUnaryServerChain(
			grpc_prometheus.UnaryServerInterceptor,
			apmgrpc.NewUnaryServerInterceptor(apmgrpc.WithRecovery()),
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.UnaryServerInterceptor(simpleLogger),
			grpc_validator.UnaryServerInterceptor(),
		),
	)

	return &GRPC{
		grpcSrv: server,
		address: address,
		cfg:     cfg,
	}
}

// RegisterEndpoints ...
func (g *GRPC) RegisterEndpoints(ctx context.Context, logger hclog.Logger, RegisterEndpointFns ...func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error) error {
	mux := runtime.NewServeMux(
		runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
			m := make(map[string]string)
			m["authorization"] = req.Header.Get("Authorization")

			b := bytes.NewBufferString("")
			req.Header.Write(b)

			logger.Debug("metadata", req.Header.Get("Authorization"))

			return metadata.New(m)
		}),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{MarshalOptions: protojson.MarshalOptions{
			EmitUnpopulated: true,
		},
			UnmarshalOptions: protojson.UnmarshalOptions{},
		}),
	)
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
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
		if err := fn(ctx, mux, ":"+g.cfg.PortGrpc, opts); err != nil {
			log.Error(err)
		}
	}

	var group errgroup.Group
	group.Go(func() error {
		return http.ListenAndServe(":"+g.cfg.Port, headers(mux))
	})
	if err := group.Wait(); err != nil {
		log.Error(err)
		return err
	}

	<-ctx.Done()

	return nil
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

	var group errgroup.Group
	group.Go(func() error {
		return g.grpcSrv.Serve(conn)
	})
	if err := group.Wait(); err != nil {
		log.Error(err)
		return err
	}

	log.Infof("Shutdown grpc server %s", g.address)
	<-ctx.Done()
	g.grpcSrv.GracefulStop()

	return nil
}
