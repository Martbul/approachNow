//gRPC server auth

package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/hashicorp/go-hclog"
	"github.com/joho/godotenv"
	protosAuth "github.com/martbul/auth/protos/auth"
	"github.com/martbul/auth/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	logger := hclog.Default()

	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Unable to load .env file", "error", err)
	}

	authGrpcServerPort := os.Getenv("DIAL_AUTH_PORT")

	network, err := net.Listen("tcp", authGrpcServerPort)
	if err != nil {
		logger.Error("Unable to listen on port 8000", "error", err)
	}

	grpcServer := grpc.NewServer()

	authServer := server.NewAuthServer(logger)

	protosAuth.RegisterAuthServer(grpcServer, authServer)

	reflection.Register(grpcServer)

	go func() {
		logger.Info("gRPC server listening:", "port", authGrpcServerPort)
		err := grpcServer.Serve(network)
		if err != nil {
			logger.Error("Unable to serve gRPC server", "error", err)
		}
	}()

	// Graceful shutdown on interrupt or termination signals
	waitForShutdown(grpcServer, logger)

}

// waitForShutdown gracefully handles server shutdown
func waitForShutdown(grpcServer *grpc.Server, logger hclog.Logger) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received
	<-signalChan

	logger.Info("\nReceived shutdown signal, gracefully stopping gRPC server...")
	grpcServer.GracefulStop()
	logger.Info("gRPC server stopped.")
}
