// Contains the actual implementation of the Auth gRPC service (handles user registration and login).

package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/martbul/auth/data"
	protosAuth "github.com/martbul/auth/protos/auth"
	"github.com/martbul/auth/utils"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthServer implements the Auth gRPC service
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

// TODO: Add data validation (check is the email is an indeed email, min length on email and password, ...)
// RegisterUser handles user registration and returns a JWT token for automatic login
func (as *AuthServer) RegisterUser(ctx context.Context, req *protosAuth.RegisterUserRequest) (*protosAuth.RegisterUserResponse, error) {


	// check if email already exists
	existingUser, err := data.GetUserByEmail(ctx, req.Email)
	//! BEST ERROR HANDLING (FOR NOW)
	if err != nil {
		as.logger.Error("Unable to check user's email", "error", err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	if existingUser != nil {
		as.logger.Error("Email already registered", "error", err)

		return nil, status.Error(codes.AlreadyExists, "Email already registered")
	}


	//hashing the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		as.logger.Error("Unable to hash the password", "error", err)
		return &protosAuth.RegisterUserResponse{
			Success: false,
			Message: "Internal server error",
		}, err
	}


	//store the user in DB
	err = data.CreateUser(ctx, req.Username, req.Email, string(hashedPassword))
	if err != nil {
		as.logger.Error("Failed to create user", "error", err)
		return &protosAuth.RegisterUserResponse{
			Success: false,
			Message: "Internal server error",
		}, err
	}


	//Generate JWT token
	tokenString, err := utils.GenerateJWT(req.Email)
	if err != nil {
		as.logger.Error("Unalbe to generate token", "error", err)
		return &protosAuth.RegisterUserResponse{
			Success: false,
			Message: "Failed to generate token",
		}, err
	}

	as.logger.Info("Register user and returned JWT token")


	return &protosAuth.RegisterUserResponse{
		Success: true,
		Message: "User registered successfully",
		Token:   tokenString,
	}, nil

}




// TODO: Add data validation (check is the email is an indeed email, min length on email and password, ...)
func (as *AuthServer) LoginUser(ctx context.Context, req *protosAuth.LoginUserRequest) (*protosAuth.LoginUserResponse, error) {
	// Get user by email
	user, err := data.GetUserByEmail(ctx, req.Email)
	if err != nil {
		as.logger.Error("Failed to get user by email", "error", err)
		return nil, status.Error(codes.Internal, "Internal server error")

	}

	if user == nil {
		as.logger.Error("Failed to get user by email", "error", err)
		return nil, status.Error(codes.Unauthenticated, "Invalid email or password")

	}

	//compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Invalid email or password")
	}

	// Generate JWT token
	tokenString, err := utils.GenerateJWT(req.Email)
	if err != nil {
		as.logger.Error("Failed to generate token", "error", err)
		return &protosAuth.LoginUserResponse{
			Success: false,
			Message: "Failed to generate token",
		}, err
	}

	as.logger.Info("Log user and returned JWT token")

	return &protosAuth.LoginUserResponse{
		Success: true,
		Token:   tokenString,
		Message: "Login successful",
	}, nil
}
