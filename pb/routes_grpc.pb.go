// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.3
// source: routes.proto

package pb

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
	RouteService_DeleteRoute_FullMethodName  = "/pomerium.dashboard.RouteService/DeleteRoute"
	RouteService_DeleteRoutes_FullMethodName = "/pomerium.dashboard.RouteService/DeleteRoutes"
	RouteService_GetRoute_FullMethodName     = "/pomerium.dashboard.RouteService/GetRoute"
	RouteService_ListRoutes_FullMethodName   = "/pomerium.dashboard.RouteService/ListRoutes"
	RouteService_LoadRoutes_FullMethodName   = "/pomerium.dashboard.RouteService/LoadRoutes"
	RouteService_SetRoute_FullMethodName     = "/pomerium.dashboard.RouteService/SetRoute"
	RouteService_SetRoutes_FullMethodName    = "/pomerium.dashboard.RouteService/SetRoutes"
	RouteService_MoveRoutes_FullMethodName   = "/pomerium.dashboard.RouteService/MoveRoutes"
)

// RouteServiceClient is the client API for RouteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// RouteService manages proxy route definitions
type RouteServiceClient interface {
	// DeleteRoute removes an existing route
	DeleteRoute(ctx context.Context, in *DeleteRouteRequest, opts ...grpc.CallOption) (*DeleteRouteResponse, error)
	// DeleteRoutes removes existing routes.
	DeleteRoutes(ctx context.Context, in *DeleteRoutesRequest, opts ...grpc.CallOption) (*DeleteRoutesResponse, error)
	// GetRoute retrieves an existing route
	GetRoute(ctx context.Context, in *GetRouteRequest, opts ...grpc.CallOption) (*GetRouteResponse, error)
	// ListRoutes lists routes based on ListRoutesRequest
	ListRoutes(ctx context.Context, in *ListRoutesRequest, opts ...grpc.CallOption) (*ListRoutesResponse, error)
	// LoadRoutes imports routes from an existing OSS configuration
	LoadRoutes(ctx context.Context, in *LoadRoutesRequest, opts ...grpc.CallOption) (*LoadRoutesResponse, error)
	// SetRoute creates or, if id is defined, updates an existing route
	SetRoute(ctx context.Context, in *SetRouteRequest, opts ...grpc.CallOption) (*SetRouteResponse, error)
	// SetRoutes creates or, if id is defined, updates existing routes
	SetRoutes(ctx context.Context, in *SetRoutesRequest, opts ...grpc.CallOption) (*SetRoutesResponse, error)
	// MoveRoutes takes an array of routeIds and moves them to a new namespace
	MoveRoutes(ctx context.Context, in *MoveRoutesRequest, opts ...grpc.CallOption) (*MoveRoutesResponse, error)
}

type routeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteServiceClient(cc grpc.ClientConnInterface) RouteServiceClient {
	return &routeServiceClient{cc}
}

func (c *routeServiceClient) DeleteRoute(ctx context.Context, in *DeleteRouteRequest, opts ...grpc.CallOption) (*DeleteRouteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteRouteResponse)
	err := c.cc.Invoke(ctx, RouteService_DeleteRoute_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeServiceClient) DeleteRoutes(ctx context.Context, in *DeleteRoutesRequest, opts ...grpc.CallOption) (*DeleteRoutesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteRoutesResponse)
	err := c.cc.Invoke(ctx, RouteService_DeleteRoutes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeServiceClient) GetRoute(ctx context.Context, in *GetRouteRequest, opts ...grpc.CallOption) (*GetRouteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRouteResponse)
	err := c.cc.Invoke(ctx, RouteService_GetRoute_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeServiceClient) ListRoutes(ctx context.Context, in *ListRoutesRequest, opts ...grpc.CallOption) (*ListRoutesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRoutesResponse)
	err := c.cc.Invoke(ctx, RouteService_ListRoutes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeServiceClient) LoadRoutes(ctx context.Context, in *LoadRoutesRequest, opts ...grpc.CallOption) (*LoadRoutesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoadRoutesResponse)
	err := c.cc.Invoke(ctx, RouteService_LoadRoutes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeServiceClient) SetRoute(ctx context.Context, in *SetRouteRequest, opts ...grpc.CallOption) (*SetRouteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetRouteResponse)
	err := c.cc.Invoke(ctx, RouteService_SetRoute_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeServiceClient) SetRoutes(ctx context.Context, in *SetRoutesRequest, opts ...grpc.CallOption) (*SetRoutesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetRoutesResponse)
	err := c.cc.Invoke(ctx, RouteService_SetRoutes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeServiceClient) MoveRoutes(ctx context.Context, in *MoveRoutesRequest, opts ...grpc.CallOption) (*MoveRoutesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MoveRoutesResponse)
	err := c.cc.Invoke(ctx, RouteService_MoveRoutes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouteServiceServer is the server API for RouteService service.
// All implementations should embed UnimplementedRouteServiceServer
// for forward compatibility.
//
// RouteService manages proxy route definitions
type RouteServiceServer interface {
	// DeleteRoute removes an existing route
	DeleteRoute(context.Context, *DeleteRouteRequest) (*DeleteRouteResponse, error)
	// DeleteRoutes removes existing routes.
	DeleteRoutes(context.Context, *DeleteRoutesRequest) (*DeleteRoutesResponse, error)
	// GetRoute retrieves an existing route
	GetRoute(context.Context, *GetRouteRequest) (*GetRouteResponse, error)
	// ListRoutes lists routes based on ListRoutesRequest
	ListRoutes(context.Context, *ListRoutesRequest) (*ListRoutesResponse, error)
	// LoadRoutes imports routes from an existing OSS configuration
	LoadRoutes(context.Context, *LoadRoutesRequest) (*LoadRoutesResponse, error)
	// SetRoute creates or, if id is defined, updates an existing route
	SetRoute(context.Context, *SetRouteRequest) (*SetRouteResponse, error)
	// SetRoutes creates or, if id is defined, updates existing routes
	SetRoutes(context.Context, *SetRoutesRequest) (*SetRoutesResponse, error)
	// MoveRoutes takes an array of routeIds and moves them to a new namespace
	MoveRoutes(context.Context, *MoveRoutesRequest) (*MoveRoutesResponse, error)
}

// UnimplementedRouteServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRouteServiceServer struct{}

func (UnimplementedRouteServiceServer) DeleteRoute(context.Context, *DeleteRouteRequest) (*DeleteRouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoute not implemented")
}
func (UnimplementedRouteServiceServer) DeleteRoutes(context.Context, *DeleteRoutesRequest) (*DeleteRoutesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoutes not implemented")
}
func (UnimplementedRouteServiceServer) GetRoute(context.Context, *GetRouteRequest) (*GetRouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoute not implemented")
}
func (UnimplementedRouteServiceServer) ListRoutes(context.Context, *ListRoutesRequest) (*ListRoutesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoutes not implemented")
}
func (UnimplementedRouteServiceServer) LoadRoutes(context.Context, *LoadRoutesRequest) (*LoadRoutesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadRoutes not implemented")
}
func (UnimplementedRouteServiceServer) SetRoute(context.Context, *SetRouteRequest) (*SetRouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRoute not implemented")
}
func (UnimplementedRouteServiceServer) SetRoutes(context.Context, *SetRoutesRequest) (*SetRoutesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRoutes not implemented")
}
func (UnimplementedRouteServiceServer) MoveRoutes(context.Context, *MoveRoutesRequest) (*MoveRoutesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveRoutes not implemented")
}
func (UnimplementedRouteServiceServer) testEmbeddedByValue() {}

// UnsafeRouteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouteServiceServer will
// result in compilation errors.
type UnsafeRouteServiceServer interface {
	mustEmbedUnimplementedRouteServiceServer()
}

func RegisterRouteServiceServer(s grpc.ServiceRegistrar, srv RouteServiceServer) {
	// If the following call pancis, it indicates UnimplementedRouteServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RouteService_ServiceDesc, srv)
}

func _RouteService_DeleteRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServiceServer).DeleteRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RouteService_DeleteRoute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServiceServer).DeleteRoute(ctx, req.(*DeleteRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteService_DeleteRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoutesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServiceServer).DeleteRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RouteService_DeleteRoutes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServiceServer).DeleteRoutes(ctx, req.(*DeleteRoutesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteService_GetRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServiceServer).GetRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RouteService_GetRoute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServiceServer).GetRoute(ctx, req.(*GetRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteService_ListRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRoutesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServiceServer).ListRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RouteService_ListRoutes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServiceServer).ListRoutes(ctx, req.(*ListRoutesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteService_LoadRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadRoutesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServiceServer).LoadRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RouteService_LoadRoutes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServiceServer).LoadRoutes(ctx, req.(*LoadRoutesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteService_SetRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServiceServer).SetRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RouteService_SetRoute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServiceServer).SetRoute(ctx, req.(*SetRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteService_SetRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRoutesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServiceServer).SetRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RouteService_SetRoutes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServiceServer).SetRoutes(ctx, req.(*SetRoutesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteService_MoveRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveRoutesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServiceServer).MoveRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RouteService_MoveRoutes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServiceServer).MoveRoutes(ctx, req.(*MoveRoutesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RouteService_ServiceDesc is the grpc.ServiceDesc for RouteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RouteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pomerium.dashboard.RouteService",
	HandlerType: (*RouteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteRoute",
			Handler:    _RouteService_DeleteRoute_Handler,
		},
		{
			MethodName: "DeleteRoutes",
			Handler:    _RouteService_DeleteRoutes_Handler,
		},
		{
			MethodName: "GetRoute",
			Handler:    _RouteService_GetRoute_Handler,
		},
		{
			MethodName: "ListRoutes",
			Handler:    _RouteService_ListRoutes_Handler,
		},
		{
			MethodName: "LoadRoutes",
			Handler:    _RouteService_LoadRoutes_Handler,
		},
		{
			MethodName: "SetRoute",
			Handler:    _RouteService_SetRoute_Handler,
		},
		{
			MethodName: "SetRoutes",
			Handler:    _RouteService_SetRoutes_Handler,
		},
		{
			MethodName: "MoveRoutes",
			Handler:    _RouteService_MoveRoutes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "routes.proto",
}