// Only used for grpc talk with athServerGrpc

package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	protosAuth "github.com/martbul/auth/protos/auth"

	"github.com/hashicorp/go-hclog"
)

type AuthHandler struct {
	logger     hclog.Logger
	authClient protosAuth.AuthClient
}

func NewAuthHandler(logger hclog.Logger, authClient protosAuth.AuthClient) *AuthHandler {
	return &AuthHandler{
		logger:     logger,
		authClient: authClient,
	}
}

func (ah *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req protosAuth.RegisterUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
log.Println("here1")
	//Call the gRPC RegisterUser services
	resp, err := ah.authClient.RegisterUser(context.Background(), &req)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	//return the result to the client
	json.NewEncoder(w).Encode(resp)
}


func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req protosAuth.LoginUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	//Call the gRPC LoginUser service
	resp, err := ah.authClient.LoginUser(context.Background(), &req)
	if err != nil {
	http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Return the result to the client
	json.NewEncoder(w).Encode(resp)
}
