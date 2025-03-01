// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.7
// source: report.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// PolicyReportRequest may either specify a list of routes,
// or request to report all routes of the namespace
type PolicyReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouteIds    []string `protobuf:"bytes,1,rep,name=route_ids,json=routeIds,proto3" json:"route_ids,omitempty"`
	NamespaceId string   `protobuf:"bytes,2,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
}

func (x *PolicyReportRequest) Reset() {
	*x = PolicyReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyReportRequest) ProtoMessage() {}

func (x *PolicyReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_report_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyReportRequest.ProtoReflect.Descriptor instead.
func (*PolicyReportRequest) Descriptor() ([]byte, []int) {
	return file_report_proto_rawDescGZIP(), []int{0}
}

func (x *PolicyReportRequest) GetRouteIds() []string {
	if x != nil {
		return x.RouteIds
	}
	return nil
}

func (x *PolicyReportRequest) GetNamespaceId() string {
	if x != nil {
		return x.NamespaceId
	}
	return ""
}

type PolicyReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Routes   []*Route  `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	Policies []*Policy `protobuf:"bytes,2,rep,name=policies,proto3" json:"policies,omitempty"`
}

func (x *PolicyReportResponse) Reset() {
	*x = PolicyReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyReportResponse) ProtoMessage() {}

func (x *PolicyReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_report_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyReportResponse.ProtoReflect.Descriptor instead.
func (*PolicyReportResponse) Descriptor() ([]byte, []int) {
	return file_report_proto_rawDescGZIP(), []int{1}
}

func (x *PolicyReportResponse) GetRoutes() []*Route {
	if x != nil {
		return x.Routes
	}
	return nil
}

func (x *PolicyReportResponse) GetPolicies() []*Policy {
	if x != nil {
		return x.Policies
	}
	return nil
}

var File_report_proto protoreflect.FileDescriptor

var file_report_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12,
	0x70, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x1a, 0x0c, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55,
	0x0a, 0x13, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x49,
	0x64, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0x81, 0x01, 0x0a, 0x14, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31,
	0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x70, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x12, 0x36, 0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x32, 0x6b, 0x0a, 0x06, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x61, 0x0a, 0x0c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x27, 0x2e, 0x70, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x70,
	0x6f, 0x6d, 0x65, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x75, 0x6d, 0x2f, 0x70, 0x6f,
	0x6d, 0x65, 0x72, 0x69, 0x75, 0x6d, 0x2d, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_report_proto_rawDescOnce sync.Once
	file_report_proto_rawDescData = file_report_proto_rawDesc
)

func file_report_proto_rawDescGZIP() []byte {
	file_report_proto_rawDescOnce.Do(func() {
		file_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_report_proto_rawDescData)
	})
	return file_report_proto_rawDescData
}

var file_report_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_report_proto_goTypes = []any{
	(*PolicyReportRequest)(nil),  // 0: pomerium.dashboard.PolicyReportRequest
	(*PolicyReportResponse)(nil), // 1: pomerium.dashboard.PolicyReportResponse
	(*Route)(nil),                // 2: pomerium.dashboard.Route
	(*Policy)(nil),               // 3: pomerium.dashboard.Policy
}
var file_report_proto_depIdxs = []int32{
	2, // 0: pomerium.dashboard.PolicyReportResponse.routes:type_name -> pomerium.dashboard.Route
	3, // 1: pomerium.dashboard.PolicyReportResponse.policies:type_name -> pomerium.dashboard.Policy
	0, // 2: pomerium.dashboard.Report.PolicyReport:input_type -> pomerium.dashboard.PolicyReportRequest
	1, // 3: pomerium.dashboard.Report.PolicyReport:output_type -> pomerium.dashboard.PolicyReportResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_report_proto_init() }
func file_report_proto_init() {
	if File_report_proto != nil {
		return
	}
	file_policy_proto_init()
	file_routes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_report_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PolicyReportRequest); i {
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
		file_report_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PolicyReportResponse); i {
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
			RawDescriptor: file_report_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_report_proto_goTypes,
		DependencyIndexes: file_report_proto_depIdxs,
		MessageInfos:      file_report_proto_msgTypes,
	}.Build()
	File_report_proto = out.File
	file_report_proto_rawDesc = nil
	file_report_proto_goTypes = nil
	file_report_proto_depIdxs = nil
}
