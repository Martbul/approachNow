package server

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-hclog"
	"github.com/martbul/near_users/data" 
	protosNearUsers "github.com/martbul/near_users/protos/near_users"
	"github.com/martbul/auth/utils"
	"github.com/martbul/near_users/validate"
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


// StreamNearbyUsers handles streaming of UserLocation messages and sends back NearbyUsersResponse messages.
func (nus *NearUsersServer) StreamNearbyUsers(stream protosNearUsers.NearUsers_StreamNearbyUsersServer) error {
	nus.logger.Info("StreamNearbyUsers started")

	for {
		// Receive a UserTokenAndLocation message from the stream
		tokAndloc, err := stream.Recv()
		if err != nil {
			if err == context.Canceled {
				nus.logger.Info("Stream cancelled by client")
			} else {
				nus.logger.Error("Error receiving from stream", "error", err)
			}
			return err
		}

		
		// Validate location
		if err := validate.ValidateLocation(tokAndloc); err != nil {
			nus.logger.Error("Invalid location", "error", err)
			return status.Errorf(codes.InvalidArgument, "invalid location: %v", err)
		}
		//! here you are just calling a func from other microservice
		//! maybe here must be used another gRPC connection
		userId,err := utils.GetUserIdFromJWT(tokAndloc.JwtToken)
		
		
		//Todo: finish error handling
		if err != nil{

		}

		if userId == -1 {

		}

		//! След като тук имам локацията на 1 потребител, трябва да я запазя в базата данни, след това да взема всички хора които са в х радиус от локацията
		//TODO: user userId to update user loc in db
		//insert/update the current user location in the user_locations table
		_, err = data.Query(stream.Context(), `
			INSERT INTO user_locations (latitude, longitude)
			VALUES ($1, $2)
		`, tokAndloc.Latitude, tokAndloc.Longitude)

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
			`, tokAndloc.Longitude, tokAndloc.Latitude, 5000)

		if err != nil {
			nus.logger.Error("Error querying database", "error", err)
			return status.Errorf(codes.Internal, "error querying database: %v", err)
		}
		defer rows.Close()


		var userID int
		var latitude float64
		var longitude float64
		var location string 
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

