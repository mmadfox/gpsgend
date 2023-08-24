// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.1
// source: proto/gpsgend/v1/tracker_service.proto

package gpsgendproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TrackerService_Subscribe_FullMethodName      = "/proto.gpsgend.v1.TrackerService/Subscribe"
	TrackerService_Unsubscribe_FullMethodName    = "/proto.gpsgend.v1.TrackerService/Unsubscribe"
	TrackerService_GetClientsInfo_FullMethodName = "/proto.gpsgend.v1.TrackerService/GetClientsInfo"
)

// TrackerServiceClient is the client API for TrackerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrackerServiceClient interface {
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (TrackerService_SubscribeClient, error)
	Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*UnsubscribeResponse, error)
	GetClientsInfo(ctx context.Context, in *GetClientsInfoRequest, opts ...grpc.CallOption) (*GetClientsInfoResponse, error)
}

type trackerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTrackerServiceClient(cc grpc.ClientConnInterface) TrackerServiceClient {
	return &trackerServiceClient{cc}
}

func (c *trackerServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (TrackerService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &TrackerService_ServiceDesc.Streams[0], TrackerService_Subscribe_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &trackerServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TrackerService_SubscribeClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type trackerServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *trackerServiceSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *trackerServiceClient) Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*UnsubscribeResponse, error) {
	out := new(UnsubscribeResponse)
	err := c.cc.Invoke(ctx, TrackerService_Unsubscribe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackerServiceClient) GetClientsInfo(ctx context.Context, in *GetClientsInfoRequest, opts ...grpc.CallOption) (*GetClientsInfoResponse, error) {
	out := new(GetClientsInfoResponse)
	err := c.cc.Invoke(ctx, TrackerService_GetClientsInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrackerServiceServer is the server API for TrackerService service.
// All implementations must embed UnimplementedTrackerServiceServer
// for forward compatibility
type TrackerServiceServer interface {
	Subscribe(*SubscribeRequest, TrackerService_SubscribeServer) error
	Unsubscribe(context.Context, *UnsubscribeRequest) (*UnsubscribeResponse, error)
	GetClientsInfo(context.Context, *GetClientsInfoRequest) (*GetClientsInfoResponse, error)
	mustEmbedUnimplementedTrackerServiceServer()
}

// UnimplementedTrackerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTrackerServiceServer struct {
}

func (UnimplementedTrackerServiceServer) Subscribe(*SubscribeRequest, TrackerService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedTrackerServiceServer) Unsubscribe(context.Context, *UnsubscribeRequest) (*UnsubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unsubscribe not implemented")
}
func (UnimplementedTrackerServiceServer) GetClientsInfo(context.Context, *GetClientsInfoRequest) (*GetClientsInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientsInfo not implemented")
}
func (UnimplementedTrackerServiceServer) mustEmbedUnimplementedTrackerServiceServer() {}

// UnsafeTrackerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrackerServiceServer will
// result in compilation errors.
type UnsafeTrackerServiceServer interface {
	mustEmbedUnimplementedTrackerServiceServer()
}

func RegisterTrackerServiceServer(s grpc.ServiceRegistrar, srv TrackerServiceServer) {
	s.RegisterService(&TrackerService_ServiceDesc, srv)
}

func _TrackerService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TrackerServiceServer).Subscribe(m, &trackerServiceSubscribeServer{stream})
}

type TrackerService_SubscribeServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type trackerServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *trackerServiceSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TrackerService_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackerServiceServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrackerService_Unsubscribe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackerServiceServer).Unsubscribe(ctx, req.(*UnsubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrackerService_GetClientsInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientsInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackerServiceServer).GetClientsInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrackerService_GetClientsInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackerServiceServer).GetClientsInfo(ctx, req.(*GetClientsInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TrackerService_ServiceDesc is the grpc.ServiceDesc for TrackerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrackerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.gpsgend.v1.TrackerService",
	HandlerType: (*TrackerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Unsubscribe",
			Handler:    _TrackerService_Unsubscribe_Handler,
		},
		{
			MethodName: "GetClientsInfo",
			Handler:    _TrackerService_GetClientsInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _TrackerService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/gpsgend/v1/tracker_service.proto",
}