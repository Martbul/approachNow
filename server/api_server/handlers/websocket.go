package handlers

import (
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
}

func NewWebsocketConnectionUserLocation(logger hclog.Logger) *WebSocketConnectioUserLocation {
	return &WebSocketConnectioUserLocation{logger}
}

func (wscul *WebSocketConnectioUserLocation) HandleWebsocketConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	wscul.logger.Info("GOT REQUEST")

	if err != nil {
		wscul.logger.Error("Cannot upgrade to websocket connection", "error", err)
		return
		//TODO: handle error
	}


	// defer conn.Close()

	// for {
	// 	var loc UserLocation

	// 	err := conn.ReadJSON(&loc)
	// 	if err != nil {
	// 		wscul.logger.Error("Unable to read json from client", "error", err)
	// 		break
	// 	}

	// 	// Send location data to near_users via gRPC
	// 	grpcLoc := &protosNearUsers.UserLocation{
	// 		Latitude:  loc.Latitude,
	// 		Longitude: loc.Longitude,
	// 	}

	// 	wscul.logger.Info("LOCATION", grpcLoc)
	// 	//! maybe NEED TO BE FIXED
	// 	// _, err = grpcClient.ReceiveLocation(context.Background(), grpcLoc)
	// 	_, err = grpcClient.GetNearbyUsers(context.Background(), grpcLoc)
	// 	if err != nil {
	// 		wscul.logger.Error("Unsucesful sending of the gRPC request", "error", err)
	// 	}

	// 	// Optionally, send a response back to the WebSocket client
	// 	conn.WriteJSON(map[string]string{"status": "location received"})

	// }

	
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

		wscul.logger.Info("MESSAGE", string(p[:]))

		// Echo the message back
		if err := conn.WriteMessage(messageType, p); err != nil {
			wscul.logger.Error("Unable to write message", "err", err)
			return
		}
	}
	

}
