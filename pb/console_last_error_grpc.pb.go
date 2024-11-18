// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.3
// source: console_last_error.proto

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
	LastErrorService_GetLastError_FullMethodName = "/pomerium.dashboard.LastErrorService/GetLastError"
)

// LastErrorServiceClient is the client API for LastErrorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LastErrorServiceClient interface {
	GetLastError(ctx context.Context, in *GetLastErrorRequest, opts ...grpc.CallOption) (*GetLastErrorResponse, error)
}

type lastErrorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLastErrorServiceClient(cc grpc.ClientConnInterface) LastErrorServiceClient {
	return &lastErrorServiceClient{cc}
}

func (c *lastErrorServiceClient) GetLastError(ctx context.Context, in *GetLastErrorRequest, opts ...grpc.CallOption) (*GetLastErrorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLastErrorResponse)
	err := c.cc.Invoke(ctx, LastErrorService_GetLastError_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LastErrorServiceServer is the server API for LastErrorService service.
// All implementations should embed UnimplementedLastErrorServiceServer
// for forward compatibility.
type LastErrorServiceServer interface {
	GetLastError(context.Context, *GetLastErrorRequest) (*GetLastErrorResponse, error)
}

// UnimplementedLastErrorServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLastErrorServiceServer struct{}

func (UnimplementedLastErrorServiceServer) GetLastError(context.Context, *GetLastErrorRequest) (*GetLastErrorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastError not implemented")
}
func (UnimplementedLastErrorServiceServer) testEmbeddedByValue() {}

// UnsafeLastErrorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LastErrorServiceServer will
// result in compilation errors.
type UnsafeLastErrorServiceServer interface {
	mustEmbedUnimplementedLastErrorServiceServer()
}

func RegisterLastErrorServiceServer(s grpc.ServiceRegistrar, srv LastErrorServiceServer) {
	// If the following call pancis, it indicates UnimplementedLastErrorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LastErrorService_ServiceDesc, srv)
}

func _LastErrorService_GetLastError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLastErrorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LastErrorServiceServer).GetLastError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LastErrorService_GetLastError_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LastErrorServiceServer).GetLastError(ctx, req.(*GetLastErrorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LastErrorService_ServiceDesc is the grpc.ServiceDesc for LastErrorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LastErrorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pomerium.dashboard.LastErrorService",
	HandlerType: (*LastErrorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLastError",
			Handler:    _LastErrorService_GetLastError_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "console_last_error.proto",
}