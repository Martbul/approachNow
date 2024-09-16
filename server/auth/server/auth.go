package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	protosAuth "github.com/martbul/auth/protos/auth"
)

// NearUsersServer is a struct that implements the gRPC methods for NearUsers.
type AuthServer struct {
	protosAuth.UnimplementedAuthServer
	logger hclog.Logger
}

// NewNearUsersServer creates a new instance of the NearUsersServer with a logger.
func NewAuthServer(logger hclog.Logger) *AuthServer {
	return &AuthServer{
		logger: logger,
	}
}


func(as *AuthServer) RegisterUser(ctx context.Context, loc *protosAuth.RegisterUserAuthData) (*protosAuth.RegisterUserResponse, error) {
	return &protosAuth.RegisterUserResponse{}, nil
}


func(as *AuthServer) LoginUser(ctx context.Context, loc *protosAuth.LogiUserAuthData) (*protosAuth.LogiUserResponse, error){
	return &protosAuth.LogiUserResponse{}, nil

}