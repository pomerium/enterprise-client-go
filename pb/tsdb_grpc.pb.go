// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.3
// source: tsdb.proto

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
	TimeSeriesDB_GetRouteMetricChange_FullMethodName          = "/pomerium.dashboard.TimeSeriesDB/GetRouteMetricChange"
	TimeSeriesDB_GetRouteMetricChangeHistogram_FullMethodName = "/pomerium.dashboard.TimeSeriesDB/GetRouteMetricChangeHistogram"
	TimeSeriesDB_GetRouteMetricSeries_FullMethodName          = "/pomerium.dashboard.TimeSeriesDB/GetRouteMetricSeries"
	TimeSeriesDB_GetRouteMetricSeriesHistogram_FullMethodName = "/pomerium.dashboard.TimeSeriesDB/GetRouteMetricSeriesHistogram"
	TimeSeriesDB_GetRouteMetricSeriesMulti_FullMethodName     = "/pomerium.dashboard.TimeSeriesDB/GetRouteMetricSeriesMulti"
	TimeSeriesDB_GetUptime_FullMethodName                     = "/pomerium.dashboard.TimeSeriesDB/GetUptime"
	TimeSeriesDB_GetInstances_FullMethodName                  = "/pomerium.dashboard.TimeSeriesDB/GetInstances"
	TimeSeriesDB_GetServerMetricSeries_FullMethodName         = "/pomerium.dashboard.TimeSeriesDB/GetServerMetricSeries"
	TimeSeriesDB_GetServerMetric_FullMethodName               = "/pomerium.dashboard.TimeSeriesDB/GetServerMetric"
	TimeSeriesDB_GetStatus_FullMethodName                     = "/pomerium.dashboard.TimeSeriesDB/GetStatus"
	TimeSeriesDB_GetLastMetricError_FullMethodName            = "/pomerium.dashboard.TimeSeriesDB/GetLastMetricError"
	TimeSeriesDB_GetUsageReport_FullMethodName                = "/pomerium.dashboard.TimeSeriesDB/GetUsageReport"
)

// TimeSeriesDBClient is the client API for TimeSeriesDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// TimeSeriesDB is a generic service that is meant to be able to query for
// historical metrics and should provide a sufficient abstraction between the UI
// and underlying time series service, would it be Prometheus, embedded TSDB or
// other 3rd party provider
type TimeSeriesDBClient interface {
	// returns metric change for a period of time
	GetRouteMetricChange(ctx context.Context, in *RouteMetricChangeRequest, opts ...grpc.CallOption) (*Scalar, error)
	// returns buckets of values for a given metric
	GetRouteMetricChangeHistogram(ctx context.Context, in *RouteMetricChangeRequest, opts ...grpc.CallOption) (*ScalarBuckets, error)
	// returns metric change as time series
	GetRouteMetricSeries(ctx context.Context, in *RouteMetricSeriesRequest, opts ...grpc.CallOption) (*TimeSeriesResponse, error)
	// returns metric change as time series
	GetRouteMetricSeriesHistogram(ctx context.Context, in *RouteMetricSeriesHistogramRequest, opts ...grpc.CallOption) (*TimeSeriesResponse, error)
	// returns multiple annotated time series
	GetRouteMetricSeriesMulti(ctx context.Context, in *RouteMetricSeriesRequest, opts ...grpc.CallOption) (*TimeSeriesResponseMulti, error)
	// returns service uptime statistics
	GetUptime(ctx context.Context, in *UptimeRequest, opts ...grpc.CallOption) (*UptimeResponse, error)
	// returns list of system services with metrics
	GetInstances(ctx context.Context, in *GetInstancesRequest, opts ...grpc.CallOption) (*Instances, error)
	// returns server queries
	GetServerMetricSeries(ctx context.Context, in *ServerMetricSeriesRequest, opts ...grpc.CallOption) (*TimeSeriesResponse, error)
	// returns current metric value
	GetServerMetric(ctx context.Context, in *ServerMetricRequest, opts ...grpc.CallOption) (*Sample, error)
	// returns current status of scraping targets
	GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error)
	// returns last known error for a metric, if available
	GetLastMetricError(ctx context.Context, in *LastErrorRequest, opts ...grpc.CallOption) (*LastErrorResponse, error)
	// returns usage report
	GetUsageReport(ctx context.Context, in *UsageReportRequest, opts ...grpc.CallOption) (*UsageReportResponse, error)
}

type timeSeriesDBClient struct {
	cc grpc.ClientConnInterface
}

func NewTimeSeriesDBClient(cc grpc.ClientConnInterface) TimeSeriesDBClient {
	return &timeSeriesDBClient{cc}
}

func (c *timeSeriesDBClient) GetRouteMetricChange(ctx context.Context, in *RouteMetricChangeRequest, opts ...grpc.CallOption) (*Scalar, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Scalar)
	err := c.cc.Invoke(ctx, TimeSeriesDB_GetRouteMetricChange_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeSeriesDBClient) GetRouteMetricChangeHistogram(ctx context.Context, in *RouteMetricChangeRequest, opts ...grpc.CallOption) (*ScalarBuckets, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ScalarBuckets)
	err := c.cc.Invoke(ctx, TimeSeriesDB_GetRouteMetricChangeHistogram_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeSeriesDBClient) GetRouteMetricSeries(ctx context.Context, in *RouteMetricSeriesRequest, opts ...grpc.CallOption) (*TimeSeriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TimeSeriesResponse)
	err := c.cc.Invoke(ctx, TimeSeriesDB_GetRouteMetricSeries_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeSeriesDBClient) GetRouteMetricSeriesHistogram(ctx context.Context, in *RouteMetricSeriesHistogramRequest, opts ...grpc.CallOption) (*TimeSeriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TimeSeriesResponse)
	err := c.cc.Invoke(ctx, TimeSeriesDB_GetRouteMetricSeriesHistogram_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeSeriesDBClient) GetRouteMetricSeriesMulti(ctx context.Context, in *RouteMetricSeriesRequest, opts ...grpc.CallOption) (*TimeSeriesResponseMulti, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TimeSeriesResponseMulti)
	err := c.cc.Invoke(ctx, TimeSeriesDB_GetRouteMetricSeriesMulti_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeSeriesDBClient) GetUptime(ctx context.Context, in *UptimeRequest, opts ...grpc.CallOption) (*UptimeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UptimeResponse)
	err := c.cc.Invoke(ctx, TimeSeriesDB_GetUptime_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeSeriesDBClient) GetInstances(ctx context.Context, in *GetInstancesRequest, opts ...grpc.CallOption) (*Instances, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Instances)
	err := c.cc.Invoke(ctx, TimeSeriesDB_GetInstances_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeSeriesDBClient) GetServerMetricSeries(ctx context.Context, in *ServerMetricSeriesRequest, opts ...grpc.CallOption) (*TimeSeriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TimeSeriesResponse)
	err := c.cc.Invoke(ctx, TimeSeriesDB_GetServerMetricSeries_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeSeriesDBClient) GetServerMetric(ctx context.Context, in *ServerMetricRequest, opts ...grpc.CallOption) (*Sample, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Sample)
	err := c.cc.Invoke(ctx, TimeSeriesDB_GetServerMetric_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeSeriesDBClient) GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStatusResponse)
	err := c.cc.Invoke(ctx, TimeSeriesDB_GetStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeSeriesDBClient) GetLastMetricError(ctx context.Context, in *LastErrorRequest, opts ...grpc.CallOption) (*LastErrorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LastErrorResponse)
	err := c.cc.Invoke(ctx, TimeSeriesDB_GetLastMetricError_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeSeriesDBClient) GetUsageReport(ctx context.Context, in *UsageReportRequest, opts ...grpc.CallOption) (*UsageReportResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UsageReportResponse)
	err := c.cc.Invoke(ctx, TimeSeriesDB_GetUsageReport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimeSeriesDBServer is the server API for TimeSeriesDB service.
// All implementations should embed UnimplementedTimeSeriesDBServer
// for forward compatibility.
//
// TimeSeriesDB is a generic service that is meant to be able to query for
// historical metrics and should provide a sufficient abstraction between the UI
// and underlying time series service, would it be Prometheus, embedded TSDB or
// other 3rd party provider
type TimeSeriesDBServer interface {
	// returns metric change for a period of time
	GetRouteMetricChange(context.Context, *RouteMetricChangeRequest) (*Scalar, error)
	// returns buckets of values for a given metric
	GetRouteMetricChangeHistogram(context.Context, *RouteMetricChangeRequest) (*ScalarBuckets, error)
	// returns metric change as time series
	GetRouteMetricSeries(context.Context, *RouteMetricSeriesRequest) (*TimeSeriesResponse, error)
	// returns metric change as time series
	GetRouteMetricSeriesHistogram(context.Context, *RouteMetricSeriesHistogramRequest) (*TimeSeriesResponse, error)
	// returns multiple annotated time series
	GetRouteMetricSeriesMulti(context.Context, *RouteMetricSeriesRequest) (*TimeSeriesResponseMulti, error)
	// returns service uptime statistics
	GetUptime(context.Context, *UptimeRequest) (*UptimeResponse, error)
	// returns list of system services with metrics
	GetInstances(context.Context, *GetInstancesRequest) (*Instances, error)
	// returns server queries
	GetServerMetricSeries(context.Context, *ServerMetricSeriesRequest) (*TimeSeriesResponse, error)
	// returns current metric value
	GetServerMetric(context.Context, *ServerMetricRequest) (*Sample, error)
	// returns current status of scraping targets
	GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error)
	// returns last known error for a metric, if available
	GetLastMetricError(context.Context, *LastErrorRequest) (*LastErrorResponse, error)
	// returns usage report
	GetUsageReport(context.Context, *UsageReportRequest) (*UsageReportResponse, error)
}

// UnimplementedTimeSeriesDBServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTimeSeriesDBServer struct{}

func (UnimplementedTimeSeriesDBServer) GetRouteMetricChange(context.Context, *RouteMetricChangeRequest) (*Scalar, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRouteMetricChange not implemented")
}
func (UnimplementedTimeSeriesDBServer) GetRouteMetricChangeHistogram(context.Context, *RouteMetricChangeRequest) (*ScalarBuckets, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRouteMetricChangeHistogram not implemented")
}
func (UnimplementedTimeSeriesDBServer) GetRouteMetricSeries(context.Context, *RouteMetricSeriesRequest) (*TimeSeriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRouteMetricSeries not implemented")
}
func (UnimplementedTimeSeriesDBServer) GetRouteMetricSeriesHistogram(context.Context, *RouteMetricSeriesHistogramRequest) (*TimeSeriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRouteMetricSeriesHistogram not implemented")
}
func (UnimplementedTimeSeriesDBServer) GetRouteMetricSeriesMulti(context.Context, *RouteMetricSeriesRequest) (*TimeSeriesResponseMulti, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRouteMetricSeriesMulti not implemented")
}
func (UnimplementedTimeSeriesDBServer) GetUptime(context.Context, *UptimeRequest) (*UptimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUptime not implemented")
}
func (UnimplementedTimeSeriesDBServer) GetInstances(context.Context, *GetInstancesRequest) (*Instances, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstances not implemented")
}
func (UnimplementedTimeSeriesDBServer) GetServerMetricSeries(context.Context, *ServerMetricSeriesRequest) (*TimeSeriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerMetricSeries not implemented")
}
func (UnimplementedTimeSeriesDBServer) GetServerMetric(context.Context, *ServerMetricRequest) (*Sample, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerMetric not implemented")
}
func (UnimplementedTimeSeriesDBServer) GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedTimeSeriesDBServer) GetLastMetricError(context.Context, *LastErrorRequest) (*LastErrorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastMetricError not implemented")
}
func (UnimplementedTimeSeriesDBServer) GetUsageReport(context.Context, *UsageReportRequest) (*UsageReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsageReport not implemented")
}
func (UnimplementedTimeSeriesDBServer) testEmbeddedByValue() {}

// UnsafeTimeSeriesDBServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TimeSeriesDBServer will
// result in compilation errors.
type UnsafeTimeSeriesDBServer interface {
	mustEmbedUnimplementedTimeSeriesDBServer()
}

func RegisterTimeSeriesDBServer(s grpc.ServiceRegistrar, srv TimeSeriesDBServer) {
	// If the following call pancis, it indicates UnimplementedTimeSeriesDBServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TimeSeriesDB_ServiceDesc, srv)
}

func _TimeSeriesDB_GetRouteMetricChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteMetricChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeSeriesDBServer).GetRouteMetricChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeSeriesDB_GetRouteMetricChange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeSeriesDBServer).GetRouteMetricChange(ctx, req.(*RouteMetricChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeSeriesDB_GetRouteMetricChangeHistogram_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteMetricChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeSeriesDBServer).GetRouteMetricChangeHistogram(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeSeriesDB_GetRouteMetricChangeHistogram_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeSeriesDBServer).GetRouteMetricChangeHistogram(ctx, req.(*RouteMetricChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeSeriesDB_GetRouteMetricSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteMetricSeriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeSeriesDBServer).GetRouteMetricSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeSeriesDB_GetRouteMetricSeries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeSeriesDBServer).GetRouteMetricSeries(ctx, req.(*RouteMetricSeriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeSeriesDB_GetRouteMetricSeriesHistogram_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteMetricSeriesHistogramRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeSeriesDBServer).GetRouteMetricSeriesHistogram(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeSeriesDB_GetRouteMetricSeriesHistogram_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeSeriesDBServer).GetRouteMetricSeriesHistogram(ctx, req.(*RouteMetricSeriesHistogramRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeSeriesDB_GetRouteMetricSeriesMulti_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteMetricSeriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeSeriesDBServer).GetRouteMetricSeriesMulti(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeSeriesDB_GetRouteMetricSeriesMulti_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeSeriesDBServer).GetRouteMetricSeriesMulti(ctx, req.(*RouteMetricSeriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeSeriesDB_GetUptime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UptimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeSeriesDBServer).GetUptime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeSeriesDB_GetUptime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeSeriesDBServer).GetUptime(ctx, req.(*UptimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeSeriesDB_GetInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeSeriesDBServer).GetInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeSeriesDB_GetInstances_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeSeriesDBServer).GetInstances(ctx, req.(*GetInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeSeriesDB_GetServerMetricSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerMetricSeriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeSeriesDBServer).GetServerMetricSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeSeriesDB_GetServerMetricSeries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeSeriesDBServer).GetServerMetricSeries(ctx, req.(*ServerMetricSeriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeSeriesDB_GetServerMetric_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerMetricRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeSeriesDBServer).GetServerMetric(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeSeriesDB_GetServerMetric_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeSeriesDBServer).GetServerMetric(ctx, req.(*ServerMetricRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeSeriesDB_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeSeriesDBServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeSeriesDB_GetStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeSeriesDBServer).GetStatus(ctx, req.(*GetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeSeriesDB_GetLastMetricError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LastErrorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeSeriesDBServer).GetLastMetricError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeSeriesDB_GetLastMetricError_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeSeriesDBServer).GetLastMetricError(ctx, req.(*LastErrorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeSeriesDB_GetUsageReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsageReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeSeriesDBServer).GetUsageReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeSeriesDB_GetUsageReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeSeriesDBServer).GetUsageReport(ctx, req.(*UsageReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TimeSeriesDB_ServiceDesc is the grpc.ServiceDesc for TimeSeriesDB service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TimeSeriesDB_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pomerium.dashboard.TimeSeriesDB",
	HandlerType: (*TimeSeriesDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRouteMetricChange",
			Handler:    _TimeSeriesDB_GetRouteMetricChange_Handler,
		},
		{
			MethodName: "GetRouteMetricChangeHistogram",
			Handler:    _TimeSeriesDB_GetRouteMetricChangeHistogram_Handler,
		},
		{
			MethodName: "GetRouteMetricSeries",
			Handler:    _TimeSeriesDB_GetRouteMetricSeries_Handler,
		},
		{
			MethodName: "GetRouteMetricSeriesHistogram",
			Handler:    _TimeSeriesDB_GetRouteMetricSeriesHistogram_Handler,
		},
		{
			MethodName: "GetRouteMetricSeriesMulti",
			Handler:    _TimeSeriesDB_GetRouteMetricSeriesMulti_Handler,
		},
		{
			MethodName: "GetUptime",
			Handler:    _TimeSeriesDB_GetUptime_Handler,
		},
		{
			MethodName: "GetInstances",
			Handler:    _TimeSeriesDB_GetInstances_Handler,
		},
		{
			MethodName: "GetServerMetricSeries",
			Handler:    _TimeSeriesDB_GetServerMetricSeries_Handler,
		},
		{
			MethodName: "GetServerMetric",
			Handler:    _TimeSeriesDB_GetServerMetric_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _TimeSeriesDB_GetStatus_Handler,
		},
		{
			MethodName: "GetLastMetricError",
			Handler:    _TimeSeriesDB_GetLastMetricError_Handler,
		},
		{
			MethodName: "GetUsageReport",
			Handler:    _TimeSeriesDB_GetUsageReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tsdb.proto",
}