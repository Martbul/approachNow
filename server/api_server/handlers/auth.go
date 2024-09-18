// // Only used for grpc talk with athServerGrpc

// package handlers

// import (
// 	"context"
// 	"encoding/json"
// 	"net/http"

// 	protosAuth "github.com/martbul/auth/protos/auth"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"

// 	"github.com/hashicorp/go-hclog"
// )

// type AuthHandler struct {
// 	logger     hclog.Logger
// 	authClient protosAuth.AuthClient
// }

// func NewAuthHandler(logger hclog.Logger, authClient protosAuth.AuthClient) *AuthHandler {
// 	return &AuthHandler{
// 		logger:     logger,
// 		authClient: authClient,
// 	}
// }

// func (ah *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
// 	var req protosAuth.RegisterUserRequest
// 	err := json.NewDecoder(r.Body).Decode(&req)
// 	if err != nil {
// 		http.Error(w, "Invalid request", http.StatusBadRequest)
// 		return
// 	}

// 	//Call the gRPC RegisterUser services
// 	resp, err := ah.authClient.RegisterUser(context.Background(), &req)
// 	if err != nil {
// 		if grpcStatus, ok := status.FromError(err); ok {
// 			switch grpcStatus.Code() {
// 			case codes.AlreadyExists:
// 				http.Error(w, grpcStatus.Message(), http.StatusConflict) // HTTP 409 Conflict
// 				return
// 			case codes.Internal:
// 				http.Error(w, "Internal server error", http.StatusInternalServerError)
// 				return
// 			}
// 		}
// 	}

// 	//return the result to the client
// 	json.NewEncoder(w).Encode(resp)
// }

// func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
// 	var req protosAuth.LoginUserRequest
// 	err := json.NewDecoder(r.Body).Decode(&req)
// 	if err != nil {
// 		http.Error(w, "Invalid request", http.StatusBadRequest)
// 		return
// 	}

// 	//Call the gRPC LoginUser service
// 	resp, err := ah.authClient.LoginUser(context.Background(), &req)
// 	if err != nil {
// 		if grpcStatus, ok := status.FromError(err); ok {
// 			switch grpcStatus.Code() {
// 			case codes.Unauthenticated:
// 				http.Error(w, grpcStatus.Message(), http.StatusForbidden)
// 				return
// 			case codes.Internal:
// 				http.Error(w, "Internal server error", http.StatusInternalServerError)
// 				return
// 			}
// 		}
// 	}

// 	// Return the result to the client
// 	json.NewEncoder(w).Encode(resp)
// }



















// package handlers

// import (
// 	"context"
// 	"encoding/json"
// 	"github/martbul/api_server/utils"
// 	"net/http"

// 	"github.com/hashicorp/go-hclog"
// 	protosAuth "github.com/martbul/auth/protos/auth"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// )



// type AuthHandler struct {
// 	logger     hclog.Logger
// 	authClient protosAuth.AuthClient
// }

// func NewAuthHandler(logger hclog.Logger, authClient protosAuth.AuthClient) *AuthHandler {
// 	return &AuthHandler{
// 		logger:     logger,
// 		authClient: authClient,
// 	}
// }



// // Register handles user registration
// func (ah *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
// 	var req protosAuth.RegisterUserRequest

// 	// Decode request
// 	 err := json.NewDecoder(r.Body).Decode(&req)
// 	if err != nil {
// 		ah.logger.Error("Error decoding request", "error", err)
// 		http.Error(w, "Invalid request format", http.StatusBadRequest)
// 		return
// 	}


// 	// Validate input
// 	if !utils.ValidateEmail(req.Email) {
// 		http.Error(w, "Invalid email format", http.StatusBadRequest)
// 		return
// 	}
// 	if !utils.ValidatePassword(req.Password) {
// 		http.Error(w, "Password must be at least 8 characters long", http.StatusBadRequest)
// 		return
// 	}

// 	// Call the gRPC RegisterUser service
// 	resp, err := ah.authClient.RegisterUser(context.Background(), &req)
// 	if err != nil {
// 		if grpcStatus, ok := status.FromError(err); ok {
// 			switch grpcStatus.Code() {
// 			case codes.AlreadyExists:
// 				http.Error(w, grpcStatus.Message(), http.StatusConflict) 
// 				return
// 			case codes.Internal:
// 				http.Error(w, "Internal server error", http.StatusInternalServerError)
// 				return
// 			default:
// 				http.Error(w, "Unexpected error occurred", http.StatusInternalServerError)
// 				return
// 			}
// 		}
// 		http.Error(w, "Unexpected error occurred", http.StatusInternalServerError)
// 		return
// 	}

// 	// Return the result to the client
// 	w.Header().Set("Content-Type", "application/json")
// 	if err := json.NewEncoder(w).Encode(resp); err != nil {
// 		ah.logger.Error("Error encoding response", "error", err)
// 		http.Error(w, "Failed to process response", http.StatusInternalServerError)
// 	}
// }


// // Login handles user login
// func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
// 	var req protosAuth.LoginUserRequest

// 	// Decode request
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		ah.logger.Error("Error decoding request", "error", err)
// 		http.Error(w, "Invalid request format", http.StatusBadRequest)
// 		return
// 	}

// 	// Validate input
// 	if !utils.ValidateEmail(req.Email) {
// 		http.Error(w, "Invalid email format", http.StatusBadRequest)
// 		return
// 	}
// 	if !utils.ValidatePassword(req.Password) {
// 		http.Error(w, "Invalid password", http.StatusBadRequest)
// 		return
// 	}

// 	// Call the gRPC LoginUser service
// 	resp, err := ah.authClient.LoginUser(context.Background(), &req)
// 	if err != nil {
// 		if grpcStatus, ok := status.FromError(err); ok {
// 			switch grpcStatus.Code() {
// 			case codes.Unauthenticated:
// 				http.Error(w, grpcStatus.Message(), http.StatusUnauthorized) // HTTP 401 Unauthorized
// 				return
// 			case codes.Internal:
// 				http.Error(w, "Internal server error", http.StatusInternalServerError)
// 				return
// 			default:
// 				http.Error(w, "Unexpected error occurred", http.StatusInternalServerError)
// 				return
// 			}
// 		}

// 		http.Error(w, "Unexpected error occurred", http.StatusInternalServerError)
// 		return
// 	}

// 	// Return the result to the client
// 	w.Header().Set("Content-Type", "application/json")
// 	if err := json.NewEncoder(w).Encode(resp); err != nil {
// 		ah.logger.Error("Error encoding response", "error", err)
// 		http.Error(w, "Failed to process response", http.StatusInternalServerError)
// 	}
// }



package handlers

import (
	"context"
	"encoding/json"
	"github/martbul/api_server/utils"
	"net/http"

	"github.com/hashicorp/go-hclog"
	protosAuth "github.com/martbul/auth/protos/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

// Register handles user registration
func (ah *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req protosAuth.RegisterUserRequest

	// Decode request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		ah.logger.Error("Error decoding request", "error", err)
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	// Validate input
	if !utils.ValidateEmail(req.Email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}
	if !utils.ValidatePassword(req.Password) {
		http.Error(w, "Password must be at least 8 characters long", http.StatusBadRequest)
		return
	}

	// Call the gRPC RegisterUser service
	resp, err := ah.authClient.RegisterUser(context.Background(), &req)
	if err != nil {
		if grpcStatus, ok := status.FromError(err); ok {
			switch grpcStatus.Code() {
			case codes.AlreadyExists:
				http.Error(w, grpcStatus.Message(), http.StatusConflict)
				return
			case codes.Internal:
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			default:
				http.Error(w, "Unexpected error occurred", http.StatusInternalServerError)
				return
			}
		}
		http.Error(w, "Unexpected error occurred", http.StatusInternalServerError)
		return
	}

	// Return the JWT token to the client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token":   resp.Token,   // Return the JWT token
		"message": resp.Message, // Return a success message
	})
}

// Login handles user login
func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req protosAuth.LoginUserRequest

	// Decode request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		ah.logger.Error("Error decoding request", "error", err)
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	// Validate input
	if !utils.ValidateEmail(req.Email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}
	if !utils.ValidatePassword(req.Password) {
		http.Error(w, "Invalid password", http.StatusBadRequest)
		return
	}

	// Call the gRPC LoginUser service
	resp, err := ah.authClient.LoginUser(context.Background(), &req)
	if err != nil {
		if grpcStatus, ok := status.FromError(err); ok {
			switch grpcStatus.Code() {
			case codes.Unauthenticated:
				http.Error(w, grpcStatus.Message(), http.StatusUnauthorized)
				return
			case codes.Internal:
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			default:
				http.Error(w, "Unexpected error occurred", http.StatusInternalServerError)
				return
			}
		}
		http.Error(w, "Unexpected error occurred", http.StatusInternalServerError)
		return
	}

	// Return the JWT token to the client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token":   resp.Token,   // Return the JWT token
		"message": resp.Message, // Return a success message
	})
}
