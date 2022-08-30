// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.14.0
// source: events.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Event_EventKind int32

const (
	Event_EVENT_KIND_UNDEFINED Event_EventKind = 0
	// envoy_service_discovery_v3.DeltaDiscoveryRequest
	Event_EVENT_DISCOVERY_REQUEST_ACK  Event_EventKind = 1
	Event_EVENT_DISCOVERY_REQUEST_NACK Event_EventKind = 2
	// envoy_service_discovery_v3.DeltaDiscoveryResponse
	Event_EVENT_DISCOVERY_RESPONSE Event_EventKind = 3
)

// Enum value maps for Event_EventKind.
var (
	Event_EventKind_name = map[int32]string{
		0: "EVENT_KIND_UNDEFINED",
		1: "EVENT_DISCOVERY_REQUEST_ACK",
		2: "EVENT_DISCOVERY_REQUEST_NACK",
		3: "EVENT_DISCOVERY_RESPONSE",
	}
	Event_EventKind_value = map[string]int32{
		"EVENT_KIND_UNDEFINED":         0,
		"EVENT_DISCOVERY_REQUEST_ACK":  1,
		"EVENT_DISCOVERY_REQUEST_NACK": 2,
		"EVENT_DISCOVERY_RESPONSE":     3,
	}
)

func (x Event_EventKind) Enum() *Event_EventKind {
	p := new(Event_EventKind)
	*p = x
	return p
}

func (x Event_EventKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event_EventKind) Descriptor() protoreflect.EnumDescriptor {
	return file_events_proto_enumTypes[0].Descriptor()
}

func (Event_EventKind) Type() protoreflect.EnumType {
	return &file_events_proto_enumTypes[0]
}

func (x Event_EventKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event_EventKind.Descriptor instead.
func (Event_EventKind) EnumDescriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{0, 0}
}

// Event represents a single envoy DeltaDiscovery event
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time    *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Message string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Code    int32                  `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
	// JSON serialized details
	Details []string `protobuf:"bytes,4,rep,name=details,proto3" json:"details,omitempty"`
	// databroker config version
	ConfigVersion uint64 `protobuf:"varint,5,opt,name=config_version,json=configVersion,proto3" json:"config_version,omitempty"`
	// envoy resource type (i.e. listener, cluster)
	TypeUrl string `protobuf:"bytes,6,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// envoy event kind
	Kind Event_EventKind `protobuf:"varint,7,opt,name=kind,proto3,enum=pomerium.dashboard.Event_EventKind" json:"kind,omitempty"`
	// envoy clusters or listeners that were added to the configuration
	ResourceSubscribed []string `protobuf:"bytes,8,rep,name=resource_subscribed,json=resourceSubscribed,proto3" json:"resource_subscribed,omitempty"`
	// clusters or listeners that were removed from the envoy configuration
	ResourceUnsubscribed []string `protobuf:"bytes,9,rep,name=resource_unsubscribed,json=resourceUnsubscribed,proto3" json:"resource_unsubscribed,omitempty"`
	// pomerium instance this event originated from
	Instance string `protobuf:"bytes,10,opt,name=instance,proto3" json:"instance,omitempty"`
	// databroker record version during this event
	SeqNo uint64 `protobuf:"varint,11,opt,name=seq_no,json=seqNo,proto3" json:"seq_no,omitempty"`
	Nonce string `protobuf:"bytes,12,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Event) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Event) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Event) GetDetails() []string {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *Event) GetConfigVersion() uint64 {
	if x != nil {
		return x.ConfigVersion
	}
	return 0
}

func (x *Event) GetTypeUrl() string {
	if x != nil {
		return x.TypeUrl
	}
	return ""
}

func (x *Event) GetKind() Event_EventKind {
	if x != nil {
		return x.Kind
	}
	return Event_EVENT_KIND_UNDEFINED
}

func (x *Event) GetResourceSubscribed() []string {
	if x != nil {
		return x.ResourceSubscribed
	}
	return nil
}

func (x *Event) GetResourceUnsubscribed() []string {
	if x != nil {
		return x.ResourceUnsubscribed
	}
	return nil
}

func (x *Event) GetInstance() string {
	if x != nil {
		return x.Instance
	}
	return ""
}

func (x *Event) GetSeqNo() uint64 {
	if x != nil {
		return x.SeqNo
	}
	return 0
}

func (x *Event) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

type SyncRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SyncRequest) Reset() {
	*x = SyncRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncRequest) ProtoMessage() {}

func (x *SyncRequest) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncRequest.ProtoReflect.Descriptor instead.
func (*SyncRequest) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{1}
}

type SyncResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *SyncResponse) Reset() {
	*x = SyncResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncResponse) ProtoMessage() {}

func (x *SyncResponse) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncResponse.ProtoReflect.Descriptor instead.
func (*SyncResponse) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{2}
}

func (x *SyncResponse) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

var File_events_proto protoreflect.FileDescriptor

var file_events_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12,
	0x70, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x04, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x74, 0x79, 0x70, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x37, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x70, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x75, 0x6d,
	0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x12, 0x2f, 0x0a, 0x13, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x64, 0x12, 0x33, 0x0a, 0x15, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x6e,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x14, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x65, 0x71, 0x5f, 0x6e, 0x6f, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x73, 0x65, 0x71, 0x4e, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x22,
	0x86, 0x01, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x0a,
	0x14, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x44, 0x45,
	0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45,
	0x53, 0x54, 0x5f, 0x41, 0x43, 0x4b, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x52, 0x45, 0x51, 0x55,
	0x45, 0x53, 0x54, 0x5f, 0x4e, 0x41, 0x43, 0x4b, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x56,
	0x45, 0x4e, 0x54, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x52, 0x45,
	0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x03, 0x22, 0x0d, 0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x0c, 0x53, 0x79, 0x6e, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x75,
	0x6d, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x32, 0x55, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x4b, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x1f, 0x2e, 0x70, 0x6f, 0x6d,
	0x65, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x6f,
	0x6d, 0x65, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42,
	0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f,
	0x6d, 0x65, 0x72, 0x69, 0x75, 0x6d, 0x2f, 0x70, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x75, 0x6d, 0x2d,
	0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_events_proto_rawDescOnce sync.Once
	file_events_proto_rawDescData = file_events_proto_rawDesc
)

func file_events_proto_rawDescGZIP() []byte {
	file_events_proto_rawDescOnce.Do(func() {
		file_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_events_proto_rawDescData)
	})
	return file_events_proto_rawDescData
}

var file_events_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_events_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_events_proto_goTypes = []interface{}{
	(Event_EventKind)(0),          // 0: pomerium.dashboard.Event.EventKind
	(*Event)(nil),                 // 1: pomerium.dashboard.Event
	(*SyncRequest)(nil),           // 2: pomerium.dashboard.SyncRequest
	(*SyncResponse)(nil),          // 3: pomerium.dashboard.SyncResponse
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_events_proto_depIdxs = []int32{
	4, // 0: pomerium.dashboard.Event.time:type_name -> google.protobuf.Timestamp
	0, // 1: pomerium.dashboard.Event.kind:type_name -> pomerium.dashboard.Event.EventKind
	1, // 2: pomerium.dashboard.SyncResponse.event:type_name -> pomerium.dashboard.Event
	2, // 3: pomerium.dashboard.Events.Sync:input_type -> pomerium.dashboard.SyncRequest
	3, // 4: pomerium.dashboard.Events.Sync:output_type -> pomerium.dashboard.SyncResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_events_proto_init() }
func file_events_proto_init() {
	if File_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncRequest); i {
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
		file_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncResponse); i {
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
			RawDescriptor: file_events_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_events_proto_goTypes,
		DependencyIndexes: file_events_proto_depIdxs,
		EnumInfos:         file_events_proto_enumTypes,
		MessageInfos:      file_events_proto_msgTypes,
	}.Build()
	File_events_proto = out.File
	file_events_proto_rawDesc = nil
	file_events_proto_goTypes = nil
	file_events_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EventsClient is the client API for Events service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventsClient interface {
	// Sync sends all current events and then pushes new events as they arrive
	Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (Events_SyncClient, error)
}

type eventsClient struct {
	cc grpc.ClientConnInterface
}

func NewEventsClient(cc grpc.ClientConnInterface) EventsClient {
	return &eventsClient{cc}
}

func (c *eventsClient) Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (Events_SyncClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Events_serviceDesc.Streams[0], "/pomerium.dashboard.Events/Sync", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventsSyncClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Events_SyncClient interface {
	Recv() (*SyncResponse, error)
	grpc.ClientStream
}

type eventsSyncClient struct {
	grpc.ClientStream
}

func (x *eventsSyncClient) Recv() (*SyncResponse, error) {
	m := new(SyncResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventsServer is the server API for Events service.
type EventsServer interface {
	// Sync sends all current events and then pushes new events as they arrive
	Sync(*SyncRequest, Events_SyncServer) error
}

// UnimplementedEventsServer can be embedded to have forward compatible implementations.
type UnimplementedEventsServer struct {
}

func (*UnimplementedEventsServer) Sync(*SyncRequest, Events_SyncServer) error {
	return status.Errorf(codes.Unimplemented, "method Sync not implemented")
}

func RegisterEventsServer(s *grpc.Server, srv EventsServer) {
	s.RegisterService(&_Events_serviceDesc, srv)
}

func _Events_Sync_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SyncRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventsServer).Sync(m, &eventsSyncServer{stream})
}

type Events_SyncServer interface {
	Send(*SyncResponse) error
	grpc.ServerStream
}

type eventsSyncServer struct {
	grpc.ServerStream
}

func (x *eventsSyncServer) Send(m *SyncResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Events_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pomerium.dashboard.Events",
	HandlerType: (*EventsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Sync",
			Handler:       _Events_Sync_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "events.proto",
}
