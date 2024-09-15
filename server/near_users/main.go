package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/hashicorp/go-hclog"
	protosNearUsers "github.com/martbul/near_users/protos/near_users"
	"github.com/martbul/near_users/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8000"
)

func main() {
	logger := hclog.Default()

	network, err := net.Listen("tcp", port)
	if err != nil {
		logger.Error("Unable to listen on port 8000", "error", err)
	}

	grpcServer := grpc.NewServer()

	nearUsersServer := server.NewNearUsersServer(logger)

	// Register the NearUsers service
	protosNearUsers.RegisterNearUsersServer(grpcServer, nearUsersServer)

	// Enable server reflection for easier debugging
	reflection.Register(grpcServer)

	go func() {
		logger.Info("gRPC server started", "port", port)
		err := grpcServer.Serve(network)
		if err != nil {

			logger.Error("Unable to serve", "error", err)

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
