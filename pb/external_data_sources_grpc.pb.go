// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.7
// source: external_data_sources.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ExternalDataSourceService_DeleteExternalDataSource_FullMethodName           = "/pomerium.dashboard.ExternalDataSourceService/DeleteExternalDataSource"
	ExternalDataSourceService_GetExternalDataSource_FullMethodName              = "/pomerium.dashboard.ExternalDataSourceService/GetExternalDataSource"
	ExternalDataSourceService_ListExternalDataSources_FullMethodName            = "/pomerium.dashboard.ExternalDataSourceService/ListExternalDataSources"
	ExternalDataSourceService_ListExternalDataSourceRecordTypes_FullMethodName  = "/pomerium.dashboard.ExternalDataSourceService/ListExternalDataSourceRecordTypes"
	ExternalDataSourceService_ListExternalDataSourceRecordFields_FullMethodName = "/pomerium.dashboard.ExternalDataSourceService/ListExternalDataSourceRecordFields"
	ExternalDataSourceService_SetExternalDataSource_FullMethodName              = "/pomerium.dashboard.ExternalDataSourceService/SetExternalDataSource"
)

// ExternalDataSourceServiceClient is the client API for ExternalDataSourceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExternalDataSourceServiceClient interface {
	DeleteExternalDataSource(ctx context.Context, in *DeleteExternalDataSourceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetExternalDataSource(ctx context.Context, in *GetExternalDataSourceRequest, opts ...grpc.CallOption) (*GetExternalDataSourceResponse, error)
	ListExternalDataSources(ctx context.Context, in *ListExternalDataSourcesRequest, opts ...grpc.CallOption) (*ListExternalDataSourcesResponse, error)
	ListExternalDataSourceRecordTypes(ctx context.Context, in *ListExternalDataSourceRecordTypesRequest, opts ...grpc.CallOption) (*ListExternalDataSourceRecordTypesResponse, error)
	ListExternalDataSourceRecordFields(ctx context.Context, in *ListExternalDataSourceRecordFieldsRequest, opts ...grpc.CallOption) (*ListExternalDataSourceRecordFieldsResponse, error)
	SetExternalDataSource(ctx context.Context, in *SetExternalDataSourceRequest, opts ...grpc.CallOption) (*SetExternalDataSourceResponse, error)
}

type externalDataSourceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExternalDataSourceServiceClient(cc grpc.ClientConnInterface) ExternalDataSourceServiceClient {
	return &externalDataSourceServiceClient{cc}
}

func (c *externalDataSourceServiceClient) DeleteExternalDataSource(ctx context.Context, in *DeleteExternalDataSourceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ExternalDataSourceService_DeleteExternalDataSource_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalDataSourceServiceClient) GetExternalDataSource(ctx context.Context, in *GetExternalDataSourceRequest, opts ...grpc.CallOption) (*GetExternalDataSourceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetExternalDataSourceResponse)
	err := c.cc.Invoke(ctx, ExternalDataSourceService_GetExternalDataSource_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalDataSourceServiceClient) ListExternalDataSources(ctx context.Context, in *ListExternalDataSourcesRequest, opts ...grpc.CallOption) (*ListExternalDataSourcesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListExternalDataSourcesResponse)
	err := c.cc.Invoke(ctx, ExternalDataSourceService_ListExternalDataSources_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalDataSourceServiceClient) ListExternalDataSourceRecordTypes(ctx context.Context, in *ListExternalDataSourceRecordTypesRequest, opts ...grpc.CallOption) (*ListExternalDataSourceRecordTypesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListExternalDataSourceRecordTypesResponse)
	err := c.cc.Invoke(ctx, ExternalDataSourceService_ListExternalDataSourceRecordTypes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalDataSourceServiceClient) ListExternalDataSourceRecordFields(ctx context.Context, in *ListExternalDataSourceRecordFieldsRequest, opts ...grpc.CallOption) (*ListExternalDataSourceRecordFieldsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListExternalDataSourceRecordFieldsResponse)
	err := c.cc.Invoke(ctx, ExternalDataSourceService_ListExternalDataSourceRecordFields_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalDataSourceServiceClient) SetExternalDataSource(ctx context.Context, in *SetExternalDataSourceRequest, opts ...grpc.CallOption) (*SetExternalDataSourceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetExternalDataSourceResponse)
	err := c.cc.Invoke(ctx, ExternalDataSourceService_SetExternalDataSource_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExternalDataSourceServiceServer is the server API for ExternalDataSourceService service.
// All implementations should embed UnimplementedExternalDataSourceServiceServer
// for forward compatibility.
type ExternalDataSourceServiceServer interface {
	DeleteExternalDataSource(context.Context, *DeleteExternalDataSourceRequest) (*emptypb.Empty, error)
	GetExternalDataSource(context.Context, *GetExternalDataSourceRequest) (*GetExternalDataSourceResponse, error)
	ListExternalDataSources(context.Context, *ListExternalDataSourcesRequest) (*ListExternalDataSourcesResponse, error)
	ListExternalDataSourceRecordTypes(context.Context, *ListExternalDataSourceRecordTypesRequest) (*ListExternalDataSourceRecordTypesResponse, error)
	ListExternalDataSourceRecordFields(context.Context, *ListExternalDataSourceRecordFieldsRequest) (*ListExternalDataSourceRecordFieldsResponse, error)
	SetExternalDataSource(context.Context, *SetExternalDataSourceRequest) (*SetExternalDataSourceResponse, error)
}

// UnimplementedExternalDataSourceServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedExternalDataSourceServiceServer struct{}

func (UnimplementedExternalDataSourceServiceServer) DeleteExternalDataSource(context.Context, *DeleteExternalDataSourceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteExternalDataSource not implemented")
}
func (UnimplementedExternalDataSourceServiceServer) GetExternalDataSource(context.Context, *GetExternalDataSourceRequest) (*GetExternalDataSourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExternalDataSource not implemented")
}
func (UnimplementedExternalDataSourceServiceServer) ListExternalDataSources(context.Context, *ListExternalDataSourcesRequest) (*ListExternalDataSourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListExternalDataSources not implemented")
}
func (UnimplementedExternalDataSourceServiceServer) ListExternalDataSourceRecordTypes(context.Context, *ListExternalDataSourceRecordTypesRequest) (*ListExternalDataSourceRecordTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListExternalDataSourceRecordTypes not implemented")
}
func (UnimplementedExternalDataSourceServiceServer) ListExternalDataSourceRecordFields(context.Context, *ListExternalDataSourceRecordFieldsRequest) (*ListExternalDataSourceRecordFieldsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListExternalDataSourceRecordFields not implemented")
}
func (UnimplementedExternalDataSourceServiceServer) SetExternalDataSource(context.Context, *SetExternalDataSourceRequest) (*SetExternalDataSourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetExternalDataSource not implemented")
}
func (UnimplementedExternalDataSourceServiceServer) testEmbeddedByValue() {}

// UnsafeExternalDataSourceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExternalDataSourceServiceServer will
// result in compilation errors.
type UnsafeExternalDataSourceServiceServer interface {
	mustEmbedUnimplementedExternalDataSourceServiceServer()
}

func RegisterExternalDataSourceServiceServer(s grpc.ServiceRegistrar, srv ExternalDataSourceServiceServer) {
	// If the following call pancis, it indicates UnimplementedExternalDataSourceServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ExternalDataSourceService_ServiceDesc, srv)
}

func _ExternalDataSourceService_DeleteExternalDataSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteExternalDataSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalDataSourceServiceServer).DeleteExternalDataSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalDataSourceService_DeleteExternalDataSource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalDataSourceServiceServer).DeleteExternalDataSource(ctx, req.(*DeleteExternalDataSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalDataSourceService_GetExternalDataSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExternalDataSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalDataSourceServiceServer).GetExternalDataSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalDataSourceService_GetExternalDataSource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalDataSourceServiceServer).GetExternalDataSource(ctx, req.(*GetExternalDataSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalDataSourceService_ListExternalDataSources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListExternalDataSourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalDataSourceServiceServer).ListExternalDataSources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalDataSourceService_ListExternalDataSources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalDataSourceServiceServer).ListExternalDataSources(ctx, req.(*ListExternalDataSourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalDataSourceService_ListExternalDataSourceRecordTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListExternalDataSourceRecordTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalDataSourceServiceServer).ListExternalDataSourceRecordTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalDataSourceService_ListExternalDataSourceRecordTypes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalDataSourceServiceServer).ListExternalDataSourceRecordTypes(ctx, req.(*ListExternalDataSourceRecordTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalDataSourceService_ListExternalDataSourceRecordFields_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListExternalDataSourceRecordFieldsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalDataSourceServiceServer).ListExternalDataSourceRecordFields(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalDataSourceService_ListExternalDataSourceRecordFields_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalDataSourceServiceServer).ListExternalDataSourceRecordFields(ctx, req.(*ListExternalDataSourceRecordFieldsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalDataSourceService_SetExternalDataSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetExternalDataSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalDataSourceServiceServer).SetExternalDataSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalDataSourceService_SetExternalDataSource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalDataSourceServiceServer).SetExternalDataSource(ctx, req.(*SetExternalDataSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExternalDataSourceService_ServiceDesc is the grpc.ServiceDesc for ExternalDataSourceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExternalDataSourceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pomerium.dashboard.ExternalDataSourceService",
	HandlerType: (*ExternalDataSourceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteExternalDataSource",
			Handler:    _ExternalDataSourceService_DeleteExternalDataSource_Handler,
		},
		{
			MethodName: "GetExternalDataSource",
			Handler:    _ExternalDataSourceService_GetExternalDataSource_Handler,
		},
		{
			MethodName: "ListExternalDataSources",
			Handler:    _ExternalDataSourceService_ListExternalDataSources_Handler,
		},
		{
			MethodName: "ListExternalDataSourceRecordTypes",
			Handler:    _ExternalDataSourceService_ListExternalDataSourceRecordTypes_Handler,
		},
		{
			MethodName: "ListExternalDataSourceRecordFields",
			Handler:    _ExternalDataSourceService_ListExternalDataSourceRecordFields_Handler,
		},
		{
			MethodName: "SetExternalDataSource",
			Handler:    _ExternalDataSourceService_SetExternalDataSource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "external_data_sources.proto",
}
