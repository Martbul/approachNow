//gRPC server near_users

package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/hashicorp/go-hclog"
	"github.com/joho/godotenv"
	"github.com/martbul/near_users/data"
	protosNearUsers "github.com/martbul/near_users/protos/near_users"
	"github.com/martbul/near_users/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)



func main() {
	logger := hclog.Default()

	// Find .env file
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Unable to load .env file", "error", err)
	}

	// Initialize the database connection
	data.InitDB()
	defer data.Close() // Ensure that the connection is closed on exit


	// Getting and using a value from .env
	nearUsersGrpcServerPort := os.Getenv("NEAR_USERS_GRPC_PORT")

	// gRPC server is listening on port 8000 for requests
	network, err := net.Listen("tcp", nearUsersGrpcServerPort)
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
		logger.Info("gRPC server listening:", "port", nearUsersGrpcServerPort)
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
