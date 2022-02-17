package grpc

import (
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type ServerOpts struct {
	EnableReflection  bool
	UnaryInterceptors []grpc.UnaryServerInterceptor
	ServerOptions     []grpc.ServerOption
}

// NewServer creates a basic health server.
func NewServer(opts ServerOpts) *grpc.Server {
	// create default interceptors
	var interceptors []grpc.UnaryServerInterceptor
	interceptors = append(interceptors,
		grpc_prometheus.UnaryServerInterceptor,
		grpc_recovery.UnaryServerInterceptor(),
	)
	interceptors = append(interceptors, opts.UnaryInterceptors...)

	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(interceptors...)),
	}

	serverOptions = append(serverOptions, opts.ServerOptions...)

	s := grpc.NewServer(serverOptions...)

	if opts.EnableReflection {
		reflection.Register(s)
	}

	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(s)

	return s
}

// StartServer starts a server.
func StartServer(server *grpc.Server, port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	zap.S().Infof("Successfully created gRPC server on %s. Attempting to serve.", port)

	return server.Serve(lis)
}
