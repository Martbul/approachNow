package handlers

import (
	// "context"
	"context"
	"encoding/json"
	// "fmt"
	"net/http"

	protosNearUsers "github.com/martbul/near_users/protos/near_users"

	"github.com/gorilla/websocket"
	"github.com/hashicorp/go-hclog"
)

var grpcClient protosNearUsers.NearUsersClient

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type UserLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type WebSocketConnectioUserLocation struct {
	logger hclog.Logger
		grpcClient protosNearUsers.NearUsersClient

}

func NewWebsocketConnectionUserLocation(logger hclog.Logger, grpcClient protosNearUsers.NearUsersClient) *WebSocketConnectioUserLocation {
	return &WebSocketConnectioUserLocation{logger: logger, grpcClient: grpcClient}
}

func (wscul *WebSocketConnectioUserLocation) HandleWebsocketConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) // Upgrade HTTP to WebSocket

	wscul.logger.Info("GOT REQUEST")

	if err != nil {
		wscul.logger.Error("Cannot upgrade to websocket connection", "error", err)
		return
		//TODO: handle error
	}

	// Create a gRPC stream to the server
	stream, err := wscul.grpcClient.StreamNearbyUsers(context.Background())
	if err != nil {
		wscul.logger.Error("Failed to create gRPC stream", "error", err)
		return
	}
	defer stream.CloseSend()



	// Listen for incoming WebSocket messages (locations)
	go func() {
		for {
			messageType, p, err := conn.ReadMessage()
		if err != nil {
			wscul.logger.Error("Unable to read message", "err", err)
			return
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			wscul.logger.Error("Unable to read message", "err", err)
			return
		}

		var result map[string]float64
		err = json.Unmarshal(p, &result)
		if err != nil {
			wscul.logger.Error("Unable to unmarshaling JSON:", "error", err)
			return
		}

		grpcLoc := &protosNearUsers.UserLocation{
			Latitude:  result["latitude"],
			Longitude: result["longitude"],
		}
			wscul.logger.Info("Received location", "latitude", grpcLoc.Latitude, "longitude", grpcLoc.Longitude)

			// Send location to gRPC stream
			err = stream.Send(grpcLoc)
			if err != nil {
				wscul.logger.Error("Error sending to gRPC stream", "error", err)
				return
			}
		}
	}()

	// Receive nearby users streamed by the server
	for {
		resp, err := stream.Recv()
		if err != nil {
			wscul.logger.Error("Error receiving from gRPC stream", "error", err)
			return
		}

		// Send received gRPC response back to WebSocket client
		nearbyUsersJSON, err := json.Marshal(resp)
		if err != nil {
			wscul.logger.Error("Error marshaling gRPC response", "error", err)
			return
		}

		if err := conn.WriteMessage(websocket.TextMessage, nearbyUsersJSON); err != nil {
			wscul.logger.Error("Unable to send WebSocket message", "error", err)
			return
		}
	}

	

}
