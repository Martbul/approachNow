//Responsible for handling HTTP requests, serving as a REST API

package main

import (
	"context"
	"github/martbul/api_server/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	protosNearUsers "github.com/martbul/near_users/protos/near_users"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

var grpcClient protosNearUsers.NearUsersClient

func main() {

	logger := hclog.Default()

	conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	//create gRPC client for near_users
	grpcClient = protosNearUsers.NewNearUsersClient(conn)

	// usersAddres, err := grpcClient.GetNearbyUsers(currentUserLocaton)

	webSocketConnectionUserLocationHandler := handlers.NewWebsocketConnectionUserLocation(logger)

	serverRouter := mux.NewRouter()

	// Define WebSocket route
	serverRouter.HandleFunc("/ws", webSocketConnectionUserLocationHandler.HandleWebsocketConnection)

	//! work to be done
	// handlers for API
	// getRouter := serverRouter.Methods(http.MethodGet).Subrouter()
	// 	getRouter.HandleFunc("/connections", productsHandler.ListAll)

	//CORS
	corsConfig := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:5000"}))

	server := http.Server{
		Addr:         ":9000",
		Handler:      corsConfig(serverRouter),
		ErrorLog:     logger.StandardLogger(&hclog.StandardLoggerOptions{}),
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	go func() {
		logger.Info("Starting api_server on port 9000")

		err := server.ListenAndServe()
		if err != nil {
			logger.Error("Error starting server", "error", err)
			os.Exit(1)
		}
	}()

	signalChannel := make(chan os.Signal)

	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	sig := <-signalChannel
	logger.Info("Got", "signal", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)

	server.Shutdown(timeoutContext)

}
