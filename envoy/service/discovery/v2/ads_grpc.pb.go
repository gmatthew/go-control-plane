// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.29.3
// source: envoy/service/discovery/v2/ads.proto

package discoveryv2

import (
	context "context"
	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AggregatedDiscoveryService_StreamAggregatedResources_FullMethodName = "/envoy.service.discovery.v2.AggregatedDiscoveryService/StreamAggregatedResources"
	AggregatedDiscoveryService_DeltaAggregatedResources_FullMethodName  = "/envoy.service.discovery.v2.AggregatedDiscoveryService/DeltaAggregatedResources"
)

// AggregatedDiscoveryServiceClient is the client API for AggregatedDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AggregatedDiscoveryServiceClient interface {
	// This is a gRPC-only API.
	StreamAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (AggregatedDiscoveryService_StreamAggregatedResourcesClient, error)
	DeltaAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (AggregatedDiscoveryService_DeltaAggregatedResourcesClient, error)
}

type aggregatedDiscoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAggregatedDiscoveryServiceClient(cc grpc.ClientConnInterface) AggregatedDiscoveryServiceClient {
	return &aggregatedDiscoveryServiceClient{cc}
}

func (c *aggregatedDiscoveryServiceClient) StreamAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (AggregatedDiscoveryService_StreamAggregatedResourcesClient, error) {
	stream, err := c.cc.NewStream(ctx, &AggregatedDiscoveryService_ServiceDesc.Streams[0], AggregatedDiscoveryService_StreamAggregatedResources_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &aggregatedDiscoveryServiceStreamAggregatedResourcesClient{stream}
	return x, nil
}

type AggregatedDiscoveryService_StreamAggregatedResourcesClient interface {
	Send(*v2.DiscoveryRequest) error
	Recv() (*v2.DiscoveryResponse, error)
	grpc.ClientStream
}

type aggregatedDiscoveryServiceStreamAggregatedResourcesClient struct {
	grpc.ClientStream
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesClient) Send(m *v2.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesClient) Recv() (*v2.DiscoveryResponse, error) {
	m := new(v2.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aggregatedDiscoveryServiceClient) DeltaAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (AggregatedDiscoveryService_DeltaAggregatedResourcesClient, error) {
	stream, err := c.cc.NewStream(ctx, &AggregatedDiscoveryService_ServiceDesc.Streams[1], AggregatedDiscoveryService_DeltaAggregatedResources_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &aggregatedDiscoveryServiceDeltaAggregatedResourcesClient{stream}
	return x, nil
}

type AggregatedDiscoveryService_DeltaAggregatedResourcesClient interface {
	Send(*v2.DeltaDiscoveryRequest) error
	Recv() (*v2.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type aggregatedDiscoveryServiceDeltaAggregatedResourcesClient struct {
	grpc.ClientStream
}

func (x *aggregatedDiscoveryServiceDeltaAggregatedResourcesClient) Send(m *v2.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceDeltaAggregatedResourcesClient) Recv() (*v2.DeltaDiscoveryResponse, error) {
	m := new(v2.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AggregatedDiscoveryServiceServer is the server API for AggregatedDiscoveryService service.
// All implementations should embed UnimplementedAggregatedDiscoveryServiceServer
// for forward compatibility
type AggregatedDiscoveryServiceServer interface {
	// This is a gRPC-only API.
	StreamAggregatedResources(AggregatedDiscoveryService_StreamAggregatedResourcesServer) error
	DeltaAggregatedResources(AggregatedDiscoveryService_DeltaAggregatedResourcesServer) error
}

// UnimplementedAggregatedDiscoveryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAggregatedDiscoveryServiceServer struct {
}

func (UnimplementedAggregatedDiscoveryServiceServer) StreamAggregatedResources(AggregatedDiscoveryService_StreamAggregatedResourcesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamAggregatedResources not implemented")
}
func (UnimplementedAggregatedDiscoveryServiceServer) DeltaAggregatedResources(AggregatedDiscoveryService_DeltaAggregatedResourcesServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaAggregatedResources not implemented")
}

// UnsafeAggregatedDiscoveryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AggregatedDiscoveryServiceServer will
// result in compilation errors.
type UnsafeAggregatedDiscoveryServiceServer interface {
	mustEmbedUnimplementedAggregatedDiscoveryServiceServer()
}

func RegisterAggregatedDiscoveryServiceServer(s grpc.ServiceRegistrar, srv AggregatedDiscoveryServiceServer) {
	s.RegisterService(&AggregatedDiscoveryService_ServiceDesc, srv)
}

func _AggregatedDiscoveryService_StreamAggregatedResources_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AggregatedDiscoveryServiceServer).StreamAggregatedResources(&aggregatedDiscoveryServiceStreamAggregatedResourcesServer{stream})
}

type AggregatedDiscoveryService_StreamAggregatedResourcesServer interface {
	Send(*v2.DiscoveryResponse) error
	Recv() (*v2.DiscoveryRequest, error)
	grpc.ServerStream
}

type aggregatedDiscoveryServiceStreamAggregatedResourcesServer struct {
	grpc.ServerStream
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesServer) Send(m *v2.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesServer) Recv() (*v2.DiscoveryRequest, error) {
	m := new(v2.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AggregatedDiscoveryService_DeltaAggregatedResources_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AggregatedDiscoveryServiceServer).DeltaAggregatedResources(&aggregatedDiscoveryServiceDeltaAggregatedResourcesServer{stream})
}

type AggregatedDiscoveryService_DeltaAggregatedResourcesServer interface {
	Send(*v2.DeltaDiscoveryResponse) error
	Recv() (*v2.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type aggregatedDiscoveryServiceDeltaAggregatedResourcesServer struct {
	grpc.ServerStream
}

func (x *aggregatedDiscoveryServiceDeltaAggregatedResourcesServer) Send(m *v2.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceDeltaAggregatedResourcesServer) Recv() (*v2.DeltaDiscoveryRequest, error) {
	m := new(v2.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AggregatedDiscoveryService_ServiceDesc is the grpc.ServiceDesc for AggregatedDiscoveryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AggregatedDiscoveryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.discovery.v2.AggregatedDiscoveryService",
	HandlerType: (*AggregatedDiscoveryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamAggregatedResources",
			Handler:       _AggregatedDiscoveryService_StreamAggregatedResources_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaAggregatedResources",
			Handler:       _AggregatedDiscoveryService_DeltaAggregatedResources_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/discovery/v2/ads.proto",
}
