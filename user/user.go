package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/health"
	grpcHealth "google.golang.org/grpc/health/grpc_health_v1"

	scgRPC "github.com/crafted-systems/smartcollect-pro/common/go/grpc"
	"github.com/crafted-systems/smartcollect-pro/common/go/metrics"
	"github.com/crafted-systems/smartcollect-pro/user/config"
	"github.com/crafted-systems/smartcollect-pro/user/logging"
	"github.com/crafted-systems/smartcollect-pro/user/proto"
	"github.com/crafted-systems/smartcollect-pro/user/rpc"
)

const SleepCleanupTime = 5

func main() {
	configuration, err := config.Parse()

	if err != nil {
		panic(err)
	}

	err = logging.Configure(configuration.LogLevel)

	if err != nil {
		panic(err)
	}

	s := scgRPC.NewServer(scgRPC.ServerOpts{
		EnableReflection: configuration.GRPCReflection,
	})

	svc, err := rpc.New()

	if err != nil {
		zap.L().Fatal("Unable to create gRPC server", zap.Error(err))
	}

	proto.RegisterUserServer(s, &svc)

	healthsrv := health.NewServer()
	grpcHealth.RegisterHealthServer(s, healthsrv)

	metrics.StartServer("/metrics", ":"+configuration.MetricsPort)

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

		defer signal.Stop(sig)

		<-sig

		zap.L().Info("Starting graceful server shutdown")
		s.GracefulStop()
	}()

	err = scgRPC.StartServer(s, ":"+configuration.Port)

	if err != nil {
		zap.L().Fatal("Unable to start gRPC server", zap.Error(err))
	}

	zap.L().Info(fmt.Sprintf("Sleeping %d seconds to allow cleanup", SleepCleanupTime))
	time.Sleep(SleepCleanupTime * time.Second)
	zap.L().Info("Finished shutdown")
}
