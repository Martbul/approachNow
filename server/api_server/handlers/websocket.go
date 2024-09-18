// package handlers

// import (
// 	// "context"
// 	"context"
// 	"encoding/json"
// 	// "fmt"
// 	"net/http"

// 	protosNearUsers "github.com/martbul/near_users/protos/near_users"

// 	"github.com/gorilla/websocket"
// 	"github.com/hashicorp/go-hclog"
// )

// var grpcClient protosNearUsers.NearUsersClient

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

// type UserLocation struct {
// 	Latitude  float64 `json:"latitude"`
// 	Longitude float64 `json:"longitude"`
// }

// type WebSocketConnectioUserLocation struct {
// 	logger hclog.Logger
// 		grpcClient protosNearUsers.NearUsersClient

// }

// func NewWebsocketConnectionUserLocation(logger hclog.Logger, grpcClient protosNearUsers.NearUsersClient) *WebSocketConnectioUserLocation {
// 	return &WebSocketConnectioUserLocation{logger: logger, grpcClient: grpcClient}
// }

// func (wscul *WebSocketConnectioUserLocation) HandleWebsocketConnection(w http.ResponseWriter, r *http.Request) {
// 	conn, err := upgrader.Upgrade(w, r, nil) // Upgrade HTTP to WebSocket

// 	wscul.logger.Info("GOT REQUEST")

// 	if err != nil {
// 		wscul.logger.Error("Cannot upgrade to websocket connection", "error", err)
// 		return
// 		//TODO: handle error better
// 	}

// 	// Create a gRPC stream to the server
// 	stream, err := wscul.grpcClient.StreamNearbyUsers(context.Background())
// 	if err != nil {
// 		wscul.logger.Error("Failed to create gRPC stream", "error", err)
// 		return
// 	}
// 	defer stream.CloseSend()



// 	// Listen for incoming WebSocket messages (locations)
// 	go func() {
// 		for {
// 			messageType, p, err := conn.ReadMessage()
// 		if err != nil {
// 			wscul.logger.Error("Unable to read message", "err", err)
// 			return
// 		}
// 		if err := conn.WriteMessage(messageType, p); err != nil {
// 			wscul.logger.Error("Unable to read message", "err", err)
// 			return
// 		}

// 		var result map[string]float64
// 		err = json.Unmarshal(p, &result)
// 		if err != nil {
// 			wscul.logger.Error("Unable to unmarshaling JSON:", "error", err)
// 			return
// 		}

// 		grpcLoc := &protosNearUsers.UserLocation{
// 			Latitude:  result["latitude"],
// 			Longitude: result["longitude"],
// 		}
// 			wscul.logger.Info("Received location", "latitude", grpcLoc.Latitude, "longitude", grpcLoc.Longitude)

// 			// Send location to gRPC stream
// 			err = stream.Send(grpcLoc)
// 			if err != nil {
// 				wscul.logger.Error("Error sending to gRPC stream", "error", err)
// 				return
// 			}
// 		}
// 	}()

// 	// Receive nearby users streamed by the server
// 	for {
// 		resp, err := stream.Recv()
// 		if err != nil {
// 			wscul.logger.Error("Error receiving from gRPC stream", "error", err)
// 			return
// 		}

// 		// Send received gRPC response back to WebSocket client
// 		nearbyUsersJSON, err := json.Marshal(resp)
// 		if err != nil {
// 			wscul.logger.Error("Error marshaling gRPC response", "error", err)
// 			return
// 		}

// 		if err := conn.WriteMessage(websocket.TextMessage, nearbyUsersJSON); err != nil {
// 			wscul.logger.Error("Unable to send WebSocket message", "error", err)
// 			return
// 		}
// 	}

	

// }






// package handlers

// import (
// 	"context"
// 	"encoding/json"
// 	"net/http"

// 	"github.com/gorilla/websocket"
// 	"github.com/hashicorp/go-hclog"
// 	protosNearUsers "github.com/martbul/near_users/protos/near_users"
// 	"github.com/martbul/auth/utils"
// )

// var (
// 	grpcClient protosNearUsers.NearUsersClient
// 	upgrader   = websocket.Upgrader{
// 		ReadBufferSize:  1024,
// 		WriteBufferSize: 1024,
// 		CheckOrigin: func(r *http.Request) bool {
// 			return true // Adjust origin checks based on your security requirements
// 		},
// 	}
// 	sessionMap = make(map[string]*websocket.Conn) // Map of JWT to WebSocket connections
// )

// type UserLocation struct {
// 	Latitude  float64 `json:"latitude"`
// 	Longitude float64 `json:"longitude"`
// }

// type WebSocketConnectionUserLocation struct {
// 	logger      hclog.Logger
// 	grpcClient  protosNearUsers.NearUsersClient
// }

// func NewWebSocketConnectionUserLocation(logger hclog.Logger, grpcClient protosNearUsers.NearUsersClient) *WebSocketConnectionUserLocation {
// 	return &WebSocketConnectionUserLocation{logger: logger, grpcClient: grpcClient}
// }

// func (wscul *WebSocketConnectionUserLocation) HandleWebSocketConnection(w http.ResponseWriter, r *http.Request) {
// 	// Extract token from query parameters
// 	token := r.URL.Query().Get("token")
// 	if token == "" {
// 		http.Error(w, "Token is required", http.StatusUnauthorized)
// 		return
// 	}

// 	// Validate JWT token
// 	_, err := utils.ValidateJWT(token)
// 	if err != nil {
// 		http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
// 		return
// 	}

// 	// Upgrade HTTP connection to WebSocket
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		wscul.logger.Error("Cannot upgrade to WebSocket connection", "error", err)
// 		return
// 	}
// 	defer conn.Close()

// 	// Store connection in session map
// 	sessionMap[token] = conn
// 	defer delete(sessionMap, token)

// 	// Create a gRPC stream to the server
// 	stream, err := wscul.grpcClient.StreamNearbyUsers(context.Background())
// 	if err != nil {
// 		wscul.logger.Error("Failed to create gRPC stream", "error", err)
// 		return
// 	}
// 	defer stream.CloseSend()

// 	// Listen for incoming WebSocket messages (locations)
// 	go func() {
// 		for {
// 			messageType, p, err := conn.ReadMessage()
// 			if err != nil {
// 				wscul.logger.Error("Unable to read message", "error", err)
// 				return
// 			}
// 			if messageType == websocket.TextMessage { // Ensure message type is text
// 				var location UserLocation
// 				if err := json.Unmarshal(p, &location); err != nil {
// 					wscul.logger.Error("Unable to unmarshal JSON", "error", err)
// 					return
// 				}

// 				grpcLoc := &protosNearUsers.UserLocation{
// 					Latitude:  location.Latitude,
// 					Longitude: location.Longitude,
// 				}
// 				wscul.logger.Info("Received location", "latitude", grpcLoc.Latitude, "longitude", grpcLoc.Longitude)

// 				// Send location to gRPC stream
// 				if err := stream.Send(grpcLoc); err != nil {
// 					wscul.logger.Error("Error sending to gRPC stream", "error", err)
// 					return
// 				}
// 			}
// 		}
// 	}()

// 	// Receive nearby users streamed by the server
// 	for {
// 		resp, err := stream.Recv()
// 		if err != nil {
// 			wscul.logger.Error("Error receiving from gRPC stream", "error", err)
// 			return
// 		}

// 		nearbyUsersJSON, err := json.Marshal(resp)
// 		if err != nil {
// 			wscul.logger.Error("Error marshaling gRPC response", "error", err)
// 			return
// 		}

// 		if err := conn.WriteMessage(websocket.TextMessage, nearbyUsersJSON); err != nil {
// 			wscul.logger.Error("Unable to send WebSocket message", "error", err)
// 			return
// 		}
// 	}
// }



package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/hashicorp/go-hclog"
	protosNearUsers "github.com/martbul/near_users/protos/near_users"
	"github.com/martbul/auth/utils"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // Adjust origin checks based on your security requirements
		},
	}
	sessionMap = make(map[string]*websocket.Conn) // Map of JWT to WebSocket connections
)

type UserLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type WebSocketConnectionUserLocation struct {
	logger      hclog.Logger
	grpcClient  protosNearUsers.NearUsersClient
}

func NewWebSocketConnectionUserLocation(logger hclog.Logger, grpcClient protosNearUsers.NearUsersClient) *WebSocketConnectionUserLocation {
	return &WebSocketConnectionUserLocation{logger: logger, grpcClient: grpcClient}
}

func (wscul *WebSocketConnectionUserLocation) HandleWebSocketConnection(w http.ResponseWriter, r *http.Request) {
	// Extract token from query parameters
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Token is required", http.StatusUnauthorized)
		return
	}

	// Validate JWT token
	_, err := utils.ValidateJWT(token)
	if err != nil {
		http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
		return
	}

	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		wscul.logger.Error("Cannot upgrade to WebSocket connection", "error", err)
		return
	}
	defer conn.Close()

	// Store connection in session map
	sessionMap[token] = conn
	defer delete(sessionMap, token)

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
				wscul.logger.Error("Unable to read message", "error", err)
				return
			}
			if messageType == websocket.TextMessage { // Ensure message type is text
				var location UserLocation
				if err := json.Unmarshal(p, &location); err != nil {
					wscul.logger.Error("Unable to unmarshal JSON", "error", err)
					return
				}

				grpcLoc := &protosNearUsers.UserLocation{
					Latitude:  location.Latitude,
					Longitude: location.Longitude,
				}
				wscul.logger.Info("Received location", "latitude", grpcLoc.Latitude, "longitude", grpcLoc.Longitude)

				// Send location to gRPC stream
				if err := stream.Send(grpcLoc); err != nil {
					wscul.logger.Error("Error sending to gRPC stream", "error", err)
					return
				}
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
