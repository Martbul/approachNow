// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.3
// source: near_users.proto

package near_users

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	NearUsers_StreamNearbyUsers_FullMethodName = "/near_users.NearUsers/StreamNearbyUsers"
)

// NearUsersClient is the client API for NearUsers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NearUsersClient interface {
	StreamNearbyUsers(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[UserTokenAndLocation, NearbyUsersResponse], error)
}

type nearUsersClient struct {
	cc grpc.ClientConnInterface
}

func NewNearUsersClient(cc grpc.ClientConnInterface) NearUsersClient {
	return &nearUsersClient{cc}
}

func (c *nearUsersClient) StreamNearbyUsers(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[UserTokenAndLocation, NearbyUsersResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &NearUsers_ServiceDesc.Streams[0], NearUsers_StreamNearbyUsers_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[UserTokenAndLocation, NearbyUsersResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type NearUsers_StreamNearbyUsersClient = grpc.BidiStreamingClient[UserTokenAndLocation, NearbyUsersResponse]

// NearUsersServer is the server API for NearUsers service.
// All implementations must embed UnimplementedNearUsersServer
// for forward compatibility.
type NearUsersServer interface {
	StreamNearbyUsers(grpc.BidiStreamingServer[UserTokenAndLocation, NearbyUsersResponse]) error
	mustEmbedUnimplementedNearUsersServer()
}

// UnimplementedNearUsersServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNearUsersServer struct{}

func (UnimplementedNearUsersServer) StreamNearbyUsers(grpc.BidiStreamingServer[UserTokenAndLocation, NearbyUsersResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamNearbyUsers not implemented")
}
func (UnimplementedNearUsersServer) mustEmbedUnimplementedNearUsersServer() {}
func (UnimplementedNearUsersServer) testEmbeddedByValue()                   {}

// UnsafeNearUsersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NearUsersServer will
// result in compilation errors.
type UnsafeNearUsersServer interface {
	mustEmbedUnimplementedNearUsersServer()
}

func RegisterNearUsersServer(s grpc.ServiceRegistrar, srv NearUsersServer) {
	// If the following call pancis, it indicates UnimplementedNearUsersServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NearUsers_ServiceDesc, srv)
}

func _NearUsers_StreamNearbyUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NearUsersServer).StreamNearbyUsers(&grpc.GenericServerStream[UserTokenAndLocation, NearbyUsersResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type NearUsers_StreamNearbyUsersServer = grpc.BidiStreamingServer[UserTokenAndLocation, NearbyUsersResponse]

// NearUsers_ServiceDesc is the grpc.ServiceDesc for NearUsers service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NearUsers_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "near_users.NearUsers",
	HandlerType: (*NearUsersServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamNearbyUsers",
			Handler:       _NearUsers_StreamNearbyUsers_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "near_users.proto",
}
