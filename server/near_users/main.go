package main

import (

	protosNearUsers "github.com/martbul/near_users/protos/near_users"

	"google.golang.org/grpc"
)

func main() {

	grpcServer := grpc.NewServer()

	protosNearUsers.RegisterNearUsersServer(grpcServer, &someImplementation{})

}
