package server

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/martbul/near_users/data" // Import the database package
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

		// След като тук имам локацията на 1 потребител, трябва да я запазя в базата данни, след това да взема всички хора които са в х радиус от локацията

		//insert/update the current user location in the user_locations table
		_, err = data.Query(stream.Context(), `
			INSERT INTO user_locations (latitude, longitude)
			VALUES ($1, $2)
		`, loc.Latitude, loc.Longitude)

		if err != nil {
			nus.logger.Error("Error inserting location into database", "error", err)
			return status.Errorf(codes.Internal, "error inserting location: %v", err)
		}

		// Perform a query to the database to get nearby users
		rows, err := data.Query(stream.Context(), `
  		   SELECT user_id, latitude, longitude, location
			FROM user_locations
			WHERE ST_DWithin(
    		ST_SetSRID(ST_MakePoint($1, $2), 4326),
    		location,
   		$3
			)
			`, loc.Longitude, loc.Latitude, 5000) // Adjust radius as needed

		if err != nil {
			nus.logger.Error("Error querying database", "error", err)
			return status.Errorf(codes.Internal, "error querying database: %v", err)
		}
		defer rows.Close()

		// Define variables to hold the column values
		var userID int
		var latitude float64
		var longitude float64
		var location string // Use string for GEOGRAPHY(POINT) or adjust as needed
		var nearbyUsers []*protosNearUsers.User
		for rows.Next() {
			err := rows.Scan(&userID, &latitude, &longitude, &location)
			if err != nil {
				nus.logger.Error("Error scanning row", "error", err)
				return status.Errorf(codes.Internal, "error scanning row: %v", err)
			}

			// Create a user object and add to the slice
			user := &protosNearUsers.User{
				Id:        fmt.Sprintf("%d", userID),
				Latitude:  latitude,
				Longitude: longitude,
				// Location field handling depends on how you need it
			}

			nearbyUsers = append(nearbyUsers, user)

			// if err := rows.Scan(&user.Id, &user.Latitude, &user.Longitude, &user.Name); err != nil {
			// 	nus.logger.Error("Error scanning row", "error", err)
			// 	return status.Errorf(codes.Internal, "error scanning row: %v", err)
			// }
			// nearbyUsers = append(nearbyUsers, &user)
		}

		if err := rows.Err(); err != nil {
			nus.logger.Error("Error iterating rows", "error", err)
			return status.Errorf(codes.Internal, "error iterating rows: %v", err)
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
