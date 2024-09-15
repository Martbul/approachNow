package server

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-hclog"
	protosNearUsers "github.com/martbul/near_users/protos/near_users"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NearUsersServer is a struct that implements the gRPC methods for NearUsers.
type NearUsersServer struct {
	protosNearUsers.UnimplementedNearUsersServer
	logger hclog.Logger
}

// NewNearUsersServer creates a new instance of the NearUsersServer with a logger.
func NewNearUsersServer(logger hclog.Logger) *NearUsersServer {
	return &NearUsersServer{
		logger: logger,
	}
}

func (nus *NearUsersServer) GetNearbyUsers(ctx context.Context, loc *protosNearUsers.UserLocation) (*protosNearUsers.NearbyUsersResponse, error) {
	nus.logger.Info("Received GetNearbyUsers request", "latitude", loc.Latitude, "longitude", loc.Longitude)

	// Validate input
	if err := validateLocation(loc); err != nil {
		nus.logger.Error("Invalid location", "error", err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid location: %v", err)
	}

	//TODO: Fetch real users
	// Simulate fetching nearby users
	time.Sleep(500 * time.Millisecond) // Simulate processing delay

	nearbyUsers := []*protosNearUsers.User{
		{
			Id:        "1",
			Latitude:  loc.Latitude + 0.01,
			Longitude: loc.Longitude + 0.01,
			Name:      "John Doe",
		},
		{
			Id:        "2",
			Latitude:  loc.Latitude + 0.02,
			Longitude: loc.Longitude + 0.02,
			Name:      "Jane Smith",
		},
	}

	resp := &protosNearUsers.NearbyUsersResponse{
		NearbyUsers: nearbyUsers,
	}

	nus.logger.Info("Returning nearby users", "count", len(nearbyUsers))
	return resp, nil
}
// StreamNearbyUsers handles streaming of UserLocation messages and sends back NearbyUsersResponse messages.
func (nus *NearUsersServer) StreamNearbyUsers(stream protosNearUsers.NearUsers_StreamNearbyUsersServer) error {
	nus.logger.Info("StreamNearbyUsers started")

	for {
		// Receive a UserLocation message from the stream
		loc, err := stream.Recv()
		if err != nil {
			if err == context.Canceled {
				nus.logger.Info("Stream cancelled by client")
			} else {
				nus.logger.Error("Error receiving from stream", "error", err)
			}
			return err
		}

		// Validate input
		if err := validateLocation(loc); err != nil {
			nus.logger.Error("Invalid location", "error", err)
			return status.Errorf(codes.InvalidArgument, "invalid location: %v", err)
		}

		// Simulate fetching nearby users
		time.Sleep(500 * time.Millisecond) // Simulate processing delay

		nearbyUsers := []*protosNearUsers.User{
			{
				Id:        "1",
				Latitude:  loc.Latitude + 0.01,
				Longitude: loc.Longitude + 0.01,
				Name:      "John Doe",
			},
			{
				Id:        "2",
				Latitude:  loc.Latitude + 0.02,
				Longitude: loc.Longitude + 0.02,
				Name:      "Jane Smith",
			},
		}

		resp := &protosNearUsers.NearbyUsersResponse{
			NearbyUsers: nearbyUsers,
		}

		nus.logger.Info("Sending nearby users", "count", len(nearbyUsers))

		// Send the response to the client
		if err := stream.Send(resp); err != nil {
			nus.logger.Error("Error sending to stream", "error", err)
			return err
		}
	}
}


// validateLocation ensures that the coordinates are within valid ranges.
func validateLocation(loc *protosNearUsers.UserLocation) error {
	if loc.Latitude < -90 || loc.Latitude > 90 {
		return fmt.Errorf("latitude must be between -90 and 90")
	}
	if loc.Longitude < -180 || loc.Longitude > 180 {
		return fmt.Errorf("longitude must be between -180 and 180")
	}
	return nil
}
