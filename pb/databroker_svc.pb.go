// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0-devel
// 	protoc        v3.20.3
// source: databroker_svc.proto

package pb

import (
	context "context"
	databroker "github.com/pomerium/pomerium/pkg/grpc/databroker"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListDataBrokerRecordsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecordType string `protobuf:"bytes,1,opt,name=record_type,json=recordType,proto3" json:"record_type,omitempty"`
}

func (x *ListDataBrokerRecordsRequest) Reset() {
	*x = ListDataBrokerRecordsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_databroker_svc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDataBrokerRecordsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDataBrokerRecordsRequest) ProtoMessage() {}

func (x *ListDataBrokerRecordsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_databroker_svc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDataBrokerRecordsRequest.ProtoReflect.Descriptor instead.
func (*ListDataBrokerRecordsRequest) Descriptor() ([]byte, []int) {
	return file_databroker_svc_proto_rawDescGZIP(), []int{0}
}

func (x *ListDataBrokerRecordsRequest) GetRecordType() string {
	if x != nil {
		return x.RecordType
	}
	return ""
}

type ListDataBrokerRecordsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []*databroker.Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *ListDataBrokerRecordsResponse) Reset() {
	*x = ListDataBrokerRecordsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_databroker_svc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDataBrokerRecordsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDataBrokerRecordsResponse) ProtoMessage() {}

func (x *ListDataBrokerRecordsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_databroker_svc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDataBrokerRecordsResponse.ProtoReflect.Descriptor instead.
func (*ListDataBrokerRecordsResponse) Descriptor() ([]byte, []int) {
	return file_databroker_svc_proto_rawDescGZIP(), []int{1}
}

func (x *ListDataBrokerRecordsResponse) GetRecords() []*databroker.Record {
	if x != nil {
		return x.Records
	}
	return nil
}

type ListDataBrokerRecordTypesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecordTypes []string `protobuf:"bytes,1,rep,name=record_types,json=recordTypes,proto3" json:"record_types,omitempty"`
}

func (x *ListDataBrokerRecordTypesResponse) Reset() {
	*x = ListDataBrokerRecordTypesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_databroker_svc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDataBrokerRecordTypesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDataBrokerRecordTypesResponse) ProtoMessage() {}

func (x *ListDataBrokerRecordTypesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_databroker_svc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDataBrokerRecordTypesResponse.ProtoReflect.Descriptor instead.
func (*ListDataBrokerRecordTypesResponse) Descriptor() ([]byte, []int) {
	return file_databroker_svc_proto_rawDescGZIP(), []int{2}
}

func (x *ListDataBrokerRecordTypesResponse) GetRecordTypes() []string {
	if x != nil {
		return x.RecordTypes
	}
	return nil
}

var File_databroker_svc_proto protoreflect.FileDescriptor

var file_databroker_svc_proto_rawDesc = []byte{
	0x0a, 0x14, 0x64, 0x61, 0x74, 0x61, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x5f, 0x73, 0x76, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x70, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x75, 0x6d,
	0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x1a, 0x41, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x75, 0x6d, 0x2f,
	0x70, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x75, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x1c, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4d, 0x0a, 0x1d, 0x4c,
	0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x46, 0x0a, 0x21, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x32, 0xf6, 0x01, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x42, 0x72, 0x6f, 0x6b, 0x65,
	0x72, 0x12, 0x7c, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x42, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x30, 0x2e, 0x70, 0x6f, 0x6d,
	0x65, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x70,
	0x6f, 0x6d, 0x65, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x6a, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x42, 0x72, 0x6f, 0x6b, 0x65,
	0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x35, 0x2e, 0x70, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x75, 0x6d, 0x2e,
	0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x6d, 0x65, 0x72, 0x69,
	0x75, 0x6d, 0x2f, 0x70, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x75, 0x6d, 0x2d, 0x63, 0x6f, 0x6e, 0x73,
	0x6f, 0x6c, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_databroker_svc_proto_rawDescOnce sync.Once
	file_databroker_svc_proto_rawDescData = file_databroker_svc_proto_rawDesc
)

func file_databroker_svc_proto_rawDescGZIP() []byte {
	file_databroker_svc_proto_rawDescOnce.Do(func() {
		file_databroker_svc_proto_rawDescData = protoimpl.X.CompressGZIP(file_databroker_svc_proto_rawDescData)
	})
	return file_databroker_svc_proto_rawDescData
}

var file_databroker_svc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_databroker_svc_proto_goTypes = []interface{}{
	(*ListDataBrokerRecordsRequest)(nil),      // 0: pomerium.dashboard.ListDataBrokerRecordsRequest
	(*ListDataBrokerRecordsResponse)(nil),     // 1: pomerium.dashboard.ListDataBrokerRecordsResponse
	(*ListDataBrokerRecordTypesResponse)(nil), // 2: pomerium.dashboard.ListDataBrokerRecordTypesResponse
	(*databroker.Record)(nil),                 // 3: databroker.Record
	(*emptypb.Empty)(nil),                     // 4: google.protobuf.Empty
}
var file_databroker_svc_proto_depIdxs = []int32{
	3, // 0: pomerium.dashboard.ListDataBrokerRecordsResponse.records:type_name -> databroker.Record
	0, // 1: pomerium.dashboard.DataBroker.ListDataBrokerRecords:input_type -> pomerium.dashboard.ListDataBrokerRecordsRequest
	4, // 2: pomerium.dashboard.DataBroker.ListDataBrokerRecordTypes:input_type -> google.protobuf.Empty
	1, // 3: pomerium.dashboard.DataBroker.ListDataBrokerRecords:output_type -> pomerium.dashboard.ListDataBrokerRecordsResponse
	2, // 4: pomerium.dashboard.DataBroker.ListDataBrokerRecordTypes:output_type -> pomerium.dashboard.ListDataBrokerRecordTypesResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_databroker_svc_proto_init() }
func file_databroker_svc_proto_init() {
	if File_databroker_svc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_databroker_svc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDataBrokerRecordsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_databroker_svc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDataBrokerRecordsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_databroker_svc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDataBrokerRecordTypesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_databroker_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_databroker_svc_proto_goTypes,
		DependencyIndexes: file_databroker_svc_proto_depIdxs,
		MessageInfos:      file_databroker_svc_proto_msgTypes,
	}.Build()
	File_databroker_svc_proto = out.File
	file_databroker_svc_proto_rawDesc = nil
	file_databroker_svc_proto_goTypes = nil
	file_databroker_svc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DataBrokerClient is the client API for DataBroker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataBrokerClient interface {
	ListDataBrokerRecords(ctx context.Context, in *ListDataBrokerRecordsRequest, opts ...grpc.CallOption) (*ListDataBrokerRecordsResponse, error)
	ListDataBrokerRecordTypes(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListDataBrokerRecordTypesResponse, error)
}

type dataBrokerClient struct {
	cc grpc.ClientConnInterface
}

func NewDataBrokerClient(cc grpc.ClientConnInterface) DataBrokerClient {
	return &dataBrokerClient{cc}
}

func (c *dataBrokerClient) ListDataBrokerRecords(ctx context.Context, in *ListDataBrokerRecordsRequest, opts ...grpc.CallOption) (*ListDataBrokerRecordsResponse, error) {
	out := new(ListDataBrokerRecordsResponse)
	err := c.cc.Invoke(ctx, "/pomerium.dashboard.DataBroker/ListDataBrokerRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataBrokerClient) ListDataBrokerRecordTypes(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListDataBrokerRecordTypesResponse, error) {
	out := new(ListDataBrokerRecordTypesResponse)
	err := c.cc.Invoke(ctx, "/pomerium.dashboard.DataBroker/ListDataBrokerRecordTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataBrokerServer is the server API for DataBroker service.
type DataBrokerServer interface {
	ListDataBrokerRecords(context.Context, *ListDataBrokerRecordsRequest) (*ListDataBrokerRecordsResponse, error)
	ListDataBrokerRecordTypes(context.Context, *emptypb.Empty) (*ListDataBrokerRecordTypesResponse, error)
}

// UnimplementedDataBrokerServer can be embedded to have forward compatible implementations.
type UnimplementedDataBrokerServer struct {
}

func (*UnimplementedDataBrokerServer) ListDataBrokerRecords(context.Context, *ListDataBrokerRecordsRequest) (*ListDataBrokerRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDataBrokerRecords not implemented")
}
func (*UnimplementedDataBrokerServer) ListDataBrokerRecordTypes(context.Context, *emptypb.Empty) (*ListDataBrokerRecordTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDataBrokerRecordTypes not implemented")
}

func RegisterDataBrokerServer(s *grpc.Server, srv DataBrokerServer) {
	s.RegisterService(&_DataBroker_serviceDesc, srv)
}

func _DataBroker_ListDataBrokerRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDataBrokerRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataBrokerServer).ListDataBrokerRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pomerium.dashboard.DataBroker/ListDataBrokerRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataBrokerServer).ListDataBrokerRecords(ctx, req.(*ListDataBrokerRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataBroker_ListDataBrokerRecordTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataBrokerServer).ListDataBrokerRecordTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pomerium.dashboard.DataBroker/ListDataBrokerRecordTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataBrokerServer).ListDataBrokerRecordTypes(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataBroker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pomerium.dashboard.DataBroker",
	HandlerType: (*DataBrokerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDataBrokerRecords",
			Handler:    _DataBroker_ListDataBrokerRecords_Handler,
		},
		{
			MethodName: "ListDataBrokerRecordTypes",
			Handler:    _DataBroker_ListDataBrokerRecordTypes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "databroker_svc.proto",
}
