package zap

import (
	"context"

	grpc_logging "github.com/grpc-ecosystem/go-grpc-middleware/logging"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// WithIgnoreMethods is an option that prevents certain methods from logging. It can be used in AllowGRPCLogging.
func WithIgnoreMethods(ignoredMethods []string) grpc_zap.Option {
	return grpc_zap.WithDecider(func(fullMethodName string, err error) bool {
		// will not log gRPC calls if it was a call to one of the ignoredMethods and no error was raised.
		if err == nil && in(fullMethodName, ignoredMethods) {
			return false
		}

		// by default everything will be logged
		return true
	})
}

func in(target string, elements []string) bool {
	for _, e := range elements {
		if e == target {
			return true
		}
	}

	return false
}

// AllowGRPCLogging allows a zap logger to be called from within grpc functions. It has to be added before you can add LogRequests. In the middleware chain, this means calling `grpc_middleware.ChainUnaryServer(AllowGRPCLogging, LogRequests)`.
func AllowGRPCLogging(logger *zap.Logger, opts ...grpc_zap.Option) grpc.UnaryServerInterceptor {
	return grpc_zap.UnaryServerInterceptor(logger, opts...)
}

// WithLogRequestIgnoreMethods can be used for LogRequest to ignore certain methods.
func WithLogRequestIgnoreMethods(ignoredMethods []string) grpc_logging.ServerPayloadLoggingDecider {
	return func(_ context.Context, fullMethodName string, _ interface{}) bool {
		// will not log gRPC calls if it was a call to one of the ignoredMethods.
		if in(fullMethodName, ignoredMethods) { //nolint:gosimple // in this case it's more clear this way.
			return false
		}

		// by default everything will be logged
		return true
	}
}

// LogRequests is just a wrapper for clarity of what it *actually* does. You need to add `AllowGRPCLogging` beforehand. In the middleware chain, this means calling `grpc_middleware.ChainUnaryServer(AllowGRPCLogging, LogRequests)`.
func LogRequests(logger *zap.Logger, decider grpc_logging.ServerPayloadLoggingDecider) grpc.UnaryServerInterceptor {
	return grpc_zap.PayloadUnaryServerInterceptor(logger, decider)
}
