// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.1
// source: proto/gpsgend/v1/tracker_service.proto

package gpsgendproto

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

type GetClientsInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetClientsInfoRequest) Reset() {
	*x = GetClientsInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gpsgend_v1_tracker_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientsInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientsInfoRequest) ProtoMessage() {}

func (x *GetClientsInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gpsgend_v1_tracker_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientsInfoRequest.ProtoReflect.Descriptor instead.
func (*GetClientsInfoRequest) Descriptor() ([]byte, []int) {
	return file_proto_gpsgend_v1_tracker_service_proto_rawDescGZIP(), []int{0}
}

type GetClientsInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clients []*ClientInfo `protobuf:"bytes,1,rep,name=clients,proto3" json:"clients,omitempty"`
}

func (x *GetClientsInfoResponse) Reset() {
	*x = GetClientsInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gpsgend_v1_tracker_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientsInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientsInfoResponse) ProtoMessage() {}

func (x *GetClientsInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gpsgend_v1_tracker_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientsInfoResponse.ProtoReflect.Descriptor instead.
func (*GetClientsInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_gpsgend_v1_tracker_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetClientsInfoResponse) GetClients() []*ClientInfo {
	if x != nil {
		return x.Clients
	}
	return nil
}

type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gpsgend_v1_tracker_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gpsgend_v1_tracker_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_proto_gpsgend_v1_tracker_service_proto_rawDescGZIP(), []int{2}
}

func (x *SubscribeRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type SubscribeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Packet []byte `protobuf:"bytes,1,opt,name=packet,proto3" json:"packet,omitempty"`
}

func (x *SubscribeResponse) Reset() {
	*x = SubscribeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gpsgend_v1_tracker_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeResponse) ProtoMessage() {}

func (x *SubscribeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gpsgend_v1_tracker_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeResponse.ProtoReflect.Descriptor instead.
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return file_proto_gpsgend_v1_tracker_service_proto_rawDescGZIP(), []int{3}
}

func (x *SubscribeResponse) GetPacket() []byte {
	if x != nil {
		return x.Packet
	}
	return nil
}

type UnsubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *UnsubscribeRequest) Reset() {
	*x = UnsubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gpsgend_v1_tracker_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribeRequest) ProtoMessage() {}

func (x *UnsubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gpsgend_v1_tracker_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribeRequest.ProtoReflect.Descriptor instead.
func (*UnsubscribeRequest) Descriptor() ([]byte, []int) {
	return file_proto_gpsgend_v1_tracker_service_proto_rawDescGZIP(), []int{4}
}

func (x *UnsubscribeRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type UnsubscribeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *UnsubscribeResponse) Reset() {
	*x = UnsubscribeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gpsgend_v1_tracker_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsubscribeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribeResponse) ProtoMessage() {}

func (x *UnsubscribeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gpsgend_v1_tracker_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribeResponse.ProtoReflect.Descriptor instead.
func (*UnsubscribeResponse) Descriptor() ([]byte, []int) {
	return file_proto_gpsgend_v1_tracker_service_proto_rawDescGZIP(), []int{5}
}

func (x *UnsubscribeResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_proto_gpsgend_v1_tracker_service_proto protoreflect.FileDescriptor

var file_proto_gpsgend_v1_tracker_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x70, 0x73, 0x67, 0x65, 0x6e, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x67, 0x70, 0x73, 0x67, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x70, 0x73, 0x67, 0x65, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x50, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x70, 0x73, 0x67, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x22, 0x2f, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x22, 0x31, 0x0a, 0x12, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0x44, 0x0a, 0x13, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x67, 0x70, 0x73, 0x67, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xaf, 0x02, 0x0a, 0x0e, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a,
	0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x22, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x67, 0x70, 0x73, 0x67, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x70, 0x73, 0x67, 0x65, 0x6e, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x5c, 0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67,
	0x70, 0x73, 0x67, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x70, 0x73, 0x67, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x67, 0x70, 0x73, 0x67, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x70, 0x73, 0x67, 0x65, 0x6e, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1f, 0x5a, 0x1d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x70, 0x73, 0x67, 0x65, 0x6e, 0x64, 0x2f, 0x76, 0x31,
	0x3b, 0x67, 0x70, 0x73, 0x67, 0x65, 0x6e, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_gpsgend_v1_tracker_service_proto_rawDescOnce sync.Once
	file_proto_gpsgend_v1_tracker_service_proto_rawDescData = file_proto_gpsgend_v1_tracker_service_proto_rawDesc
)

func file_proto_gpsgend_v1_tracker_service_proto_rawDescGZIP() []byte {
	file_proto_gpsgend_v1_tracker_service_proto_rawDescOnce.Do(func() {
		file_proto_gpsgend_v1_tracker_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_gpsgend_v1_tracker_service_proto_rawDescData)
	})
	return file_proto_gpsgend_v1_tracker_service_proto_rawDescData
}

var file_proto_gpsgend_v1_tracker_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_gpsgend_v1_tracker_service_proto_goTypes = []interface{}{
	(*GetClientsInfoRequest)(nil),  // 0: proto.gpsgend.v1.GetClientsInfoRequest
	(*GetClientsInfoResponse)(nil), // 1: proto.gpsgend.v1.GetClientsInfoResponse
	(*SubscribeRequest)(nil),       // 2: proto.gpsgend.v1.SubscribeRequest
	(*SubscribeResponse)(nil),      // 3: proto.gpsgend.v1.SubscribeResponse
	(*UnsubscribeRequest)(nil),     // 4: proto.gpsgend.v1.UnsubscribeRequest
	(*UnsubscribeResponse)(nil),    // 5: proto.gpsgend.v1.UnsubscribeResponse
	(*ClientInfo)(nil),             // 6: proto.gpsgend.v1.ClientInfo
	(*Error)(nil),                  // 7: proto.gpsgend.v1.Error
}
var file_proto_gpsgend_v1_tracker_service_proto_depIdxs = []int32{
	6, // 0: proto.gpsgend.v1.GetClientsInfoResponse.clients:type_name -> proto.gpsgend.v1.ClientInfo
	7, // 1: proto.gpsgend.v1.UnsubscribeResponse.error:type_name -> proto.gpsgend.v1.Error
	2, // 2: proto.gpsgend.v1.TrackerService.Subscribe:input_type -> proto.gpsgend.v1.SubscribeRequest
	4, // 3: proto.gpsgend.v1.TrackerService.Unsubscribe:input_type -> proto.gpsgend.v1.UnsubscribeRequest
	0, // 4: proto.gpsgend.v1.TrackerService.GetClientsInfo:input_type -> proto.gpsgend.v1.GetClientsInfoRequest
	3, // 5: proto.gpsgend.v1.TrackerService.Subscribe:output_type -> proto.gpsgend.v1.SubscribeResponse
	5, // 6: proto.gpsgend.v1.TrackerService.Unsubscribe:output_type -> proto.gpsgend.v1.UnsubscribeResponse
	1, // 7: proto.gpsgend.v1.TrackerService.GetClientsInfo:output_type -> proto.gpsgend.v1.GetClientsInfoResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_gpsgend_v1_tracker_service_proto_init() }
func file_proto_gpsgend_v1_tracker_service_proto_init() {
	if File_proto_gpsgend_v1_tracker_service_proto != nil {
		return
	}
	file_proto_gpsgend_v1_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_gpsgend_v1_tracker_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClientsInfoRequest); i {
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
		file_proto_gpsgend_v1_tracker_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClientsInfoResponse); i {
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
		file_proto_gpsgend_v1_tracker_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
		file_proto_gpsgend_v1_tracker_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeResponse); i {
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
		file_proto_gpsgend_v1_tracker_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsubscribeRequest); i {
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
		file_proto_gpsgend_v1_tracker_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsubscribeResponse); i {
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
			RawDescriptor: file_proto_gpsgend_v1_tracker_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_gpsgend_v1_tracker_service_proto_goTypes,
		DependencyIndexes: file_proto_gpsgend_v1_tracker_service_proto_depIdxs,
		MessageInfos:      file_proto_gpsgend_v1_tracker_service_proto_msgTypes,
	}.Build()
	File_proto_gpsgend_v1_tracker_service_proto = out.File
	file_proto_gpsgend_v1_tracker_service_proto_rawDesc = nil
	file_proto_gpsgend_v1_tracker_service_proto_goTypes = nil
	file_proto_gpsgend_v1_tracker_service_proto_depIdxs = nil
}