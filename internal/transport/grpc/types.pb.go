// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.1
// source: internal/transport/grpc/types.proto

package grpc

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

type Tracker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomId    string     `protobuf:"bytes,2,opt,name=custom_id,json=customId,proto3" json:"custom_id,omitempty"`
	Status      *Status    `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Model       string     `protobuf:"bytes,4,opt,name=model,proto3" json:"model,omitempty"`
	Color       string     `protobuf:"bytes,5,opt,name=color,proto3" json:"color,omitempty"`
	Descr       string     `protobuf:"bytes,6,opt,name=descr,proto3" json:"descr,omitempty"`
	Offline     *Offline   `protobuf:"bytes,7,opt,name=offline,proto3" json:"offline,omitempty"`
	Elevation   *Elevation `protobuf:"bytes,8,opt,name=elevation,proto3" json:"elevation,omitempty"`
	Battery     *Battery   `protobuf:"bytes,9,opt,name=battery,proto3" json:"battery,omitempty"`
	Speed       *Speed     `protobuf:"bytes,10,opt,name=speed,proto3" json:"speed,omitempty"`
	Props       []byte     `protobuf:"bytes,11,opt,name=props,proto3" json:"props,omitempty"`
	NumSensors  int64      `protobuf:"varint,12,opt,name=num_sensors,json=numSensors,proto3" json:"num_sensors,omitempty"`
	NumRoutes   int64      `protobuf:"varint,13,opt,name=num_routes,json=numRoutes,proto3" json:"num_routes,omitempty"`
	SkipOffline bool       `protobuf:"varint,14,opt,name=skip_offline,json=skipOffline,proto3" json:"skip_offline,omitempty"`
	CreatedAt   int64      `protobuf:"varint,15,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   int64      `protobuf:"varint,16,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	RunningAt   int64      `protobuf:"varint,17,opt,name=running_at,json=runningAt,proto3" json:"running_at,omitempty"`
	StoppedAt   int64      `protobuf:"varint,18,opt,name=stopped_at,json=stoppedAt,proto3" json:"stopped_at,omitempty"`
}

func (x *Tracker) Reset() {
	*x = Tracker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_transport_grpc_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tracker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tracker) ProtoMessage() {}

func (x *Tracker) ProtoReflect() protoreflect.Message {
	mi := &file_internal_transport_grpc_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tracker.ProtoReflect.Descriptor instead.
func (*Tracker) Descriptor() ([]byte, []int) {
	return file_internal_transport_grpc_types_proto_rawDescGZIP(), []int{0}
}

func (x *Tracker) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tracker) GetCustomId() string {
	if x != nil {
		return x.CustomId
	}
	return ""
}

func (x *Tracker) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *Tracker) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Tracker) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *Tracker) GetDescr() string {
	if x != nil {
		return x.Descr
	}
	return ""
}

func (x *Tracker) GetOffline() *Offline {
	if x != nil {
		return x.Offline
	}
	return nil
}

func (x *Tracker) GetElevation() *Elevation {
	if x != nil {
		return x.Elevation
	}
	return nil
}

func (x *Tracker) GetBattery() *Battery {
	if x != nil {
		return x.Battery
	}
	return nil
}

func (x *Tracker) GetSpeed() *Speed {
	if x != nil {
		return x.Speed
	}
	return nil
}

func (x *Tracker) GetProps() []byte {
	if x != nil {
		return x.Props
	}
	return nil
}

func (x *Tracker) GetNumSensors() int64 {
	if x != nil {
		return x.NumSensors
	}
	return 0
}

func (x *Tracker) GetNumRoutes() int64 {
	if x != nil {
		return x.NumRoutes
	}
	return 0
}

func (x *Tracker) GetSkipOffline() bool {
	if x != nil {
		return x.SkipOffline
	}
	return false
}

func (x *Tracker) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Tracker) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Tracker) GetRunningAt() int64 {
	if x != nil {
		return x.RunningAt
	}
	return 0
}

func (x *Tracker) GetStoppedAt() int64 {
	if x != nil {
		return x.StoppedAt
	}
	return 0
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_transport_grpc_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_internal_transport_grpc_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_internal_transport_grpc_types_proto_rawDescGZIP(), []int{1}
}

func (x *Error) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Model  string `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
	Descr  string `protobuf:"bytes,2,opt,name=descr,proto3" json:"descr,omitempty"`
	Color  string `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
	Status int64  `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	Limit  int64  `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int64  `protobuf:"varint,6,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_transport_grpc_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_internal_transport_grpc_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_internal_transport_grpc_types_proto_rawDescGZIP(), []int{2}
}

func (x *Filter) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Filter) GetDescr() string {
	if x != nil {
		return x.Descr
	}
	return ""
}

func (x *Filter) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *Filter) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Filter) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Filter) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_transport_grpc_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_internal_transport_grpc_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_internal_transport_grpc_types_proto_rawDescGZIP(), []int{3}
}

func (x *Status) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Status) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Offline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min int64 `protobuf:"varint,1,opt,name=min,proto3" json:"min,omitempty"`
	Max int64 `protobuf:"varint,2,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *Offline) Reset() {
	*x = Offline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_transport_grpc_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Offline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Offline) ProtoMessage() {}

func (x *Offline) ProtoReflect() protoreflect.Message {
	mi := &file_internal_transport_grpc_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Offline.ProtoReflect.Descriptor instead.
func (*Offline) Descriptor() ([]byte, []int) {
	return file_internal_transport_grpc_types_proto_rawDescGZIP(), []int{4}
}

func (x *Offline) GetMin() int64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *Offline) GetMax() int64 {
	if x != nil {
		return x.Max
	}
	return 0
}

type Elevation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min       float64 `protobuf:"fixed64,1,opt,name=min,proto3" json:"min,omitempty"`
	Max       float64 `protobuf:"fixed64,2,opt,name=max,proto3" json:"max,omitempty"`
	Amplitude int64   `protobuf:"varint,3,opt,name=amplitude,proto3" json:"amplitude,omitempty"`
	Mode      int64   `protobuf:"varint,4,opt,name=mode,proto3" json:"mode,omitempty"`
}

func (x *Elevation) Reset() {
	*x = Elevation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_transport_grpc_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Elevation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Elevation) ProtoMessage() {}

func (x *Elevation) ProtoReflect() protoreflect.Message {
	mi := &file_internal_transport_grpc_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Elevation.ProtoReflect.Descriptor instead.
func (*Elevation) Descriptor() ([]byte, []int) {
	return file_internal_transport_grpc_types_proto_rawDescGZIP(), []int{5}
}

func (x *Elevation) GetMin() float64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *Elevation) GetMax() float64 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *Elevation) GetAmplitude() int64 {
	if x != nil {
		return x.Amplitude
	}
	return 0
}

func (x *Elevation) GetMode() int64 {
	if x != nil {
		return x.Mode
	}
	return 0
}

type Battery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min        float64 `protobuf:"fixed64,1,opt,name=min,proto3" json:"min,omitempty"`
	Max        float64 `protobuf:"fixed64,2,opt,name=max,proto3" json:"max,omitempty"`
	ChargeTime int64   `protobuf:"varint,3,opt,name=charge_time,json=chargeTime,proto3" json:"charge_time,omitempty"`
}

func (x *Battery) Reset() {
	*x = Battery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_transport_grpc_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Battery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Battery) ProtoMessage() {}

func (x *Battery) ProtoReflect() protoreflect.Message {
	mi := &file_internal_transport_grpc_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Battery.ProtoReflect.Descriptor instead.
func (*Battery) Descriptor() ([]byte, []int) {
	return file_internal_transport_grpc_types_proto_rawDescGZIP(), []int{6}
}

func (x *Battery) GetMin() float64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *Battery) GetMax() float64 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *Battery) GetChargeTime() int64 {
	if x != nil {
		return x.ChargeTime
	}
	return 0
}

type Speed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min       float64 `protobuf:"fixed64,1,opt,name=min,proto3" json:"min,omitempty"`
	Max       float64 `protobuf:"fixed64,2,opt,name=max,proto3" json:"max,omitempty"`
	Amplitude int64   `protobuf:"varint,3,opt,name=amplitude,proto3" json:"amplitude,omitempty"`
}

func (x *Speed) Reset() {
	*x = Speed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_transport_grpc_types_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Speed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Speed) ProtoMessage() {}

func (x *Speed) ProtoReflect() protoreflect.Message {
	mi := &file_internal_transport_grpc_types_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Speed.ProtoReflect.Descriptor instead.
func (*Speed) Descriptor() ([]byte, []int) {
	return file_internal_transport_grpc_types_proto_rawDescGZIP(), []int{7}
}

func (x *Speed) GetMin() float64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *Speed) GetMax() float64 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *Speed) GetAmplitude() int64 {
	if x != nil {
		return x.Amplitude
	}
	return 0
}

type Navigator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lon             float64 `protobuf:"fixed64,1,opt,name=lon,proto3" json:"lon,omitempty"`
	Lat             float64 `protobuf:"fixed64,2,opt,name=lat,proto3" json:"lat,omitempty"`
	Distance        float64 `protobuf:"fixed64,3,opt,name=distance,proto3" json:"distance,omitempty"`
	RouteDistance   float64 `protobuf:"fixed64,4,opt,name=route_distance,json=routeDistance,proto3" json:"route_distance,omitempty"`
	RouteIndex      int64   `protobuf:"varint,5,opt,name=route_index,json=routeIndex,proto3" json:"route_index,omitempty"`
	TrackDistance   float64 `protobuf:"fixed64,6,opt,name=track_distance,json=trackDistance,proto3" json:"track_distance,omitempty"`
	TrackIndex      int64   `protobuf:"varint,7,opt,name=track_index,json=trackIndex,proto3" json:"track_index,omitempty"`
	SegmentDistance float64 `protobuf:"fixed64,8,opt,name=segment_distance,json=segmentDistance,proto3" json:"segment_distance,omitempty"`
	SegmentIndex    int64   `protobuf:"varint,9,opt,name=segment_index,json=segmentIndex,proto3" json:"segment_index,omitempty"`
	Units           string  `protobuf:"bytes,10,opt,name=units,proto3" json:"units,omitempty"`
}

func (x *Navigator) Reset() {
	*x = Navigator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_transport_grpc_types_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Navigator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Navigator) ProtoMessage() {}

func (x *Navigator) ProtoReflect() protoreflect.Message {
	mi := &file_internal_transport_grpc_types_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Navigator.ProtoReflect.Descriptor instead.
func (*Navigator) Descriptor() ([]byte, []int) {
	return file_internal_transport_grpc_types_proto_rawDescGZIP(), []int{8}
}

func (x *Navigator) GetLon() float64 {
	if x != nil {
		return x.Lon
	}
	return 0
}

func (x *Navigator) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Navigator) GetDistance() float64 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *Navigator) GetRouteDistance() float64 {
	if x != nil {
		return x.RouteDistance
	}
	return 0
}

func (x *Navigator) GetRouteIndex() int64 {
	if x != nil {
		return x.RouteIndex
	}
	return 0
}

func (x *Navigator) GetTrackDistance() float64 {
	if x != nil {
		return x.TrackDistance
	}
	return 0
}

func (x *Navigator) GetTrackIndex() int64 {
	if x != nil {
		return x.TrackIndex
	}
	return 0
}

func (x *Navigator) GetSegmentDistance() float64 {
	if x != nil {
		return x.SegmentDistance
	}
	return 0
}

func (x *Navigator) GetSegmentIndex() int64 {
	if x != nil {
		return x.SegmentIndex
	}
	return 0
}

func (x *Navigator) GetUnits() string {
	if x != nil {
		return x.Units
	}
	return ""
}

type Sensor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Min       float64 `protobuf:"fixed64,3,opt,name=min,proto3" json:"min,omitempty"`
	Max       float64 `protobuf:"fixed64,4,opt,name=max,proto3" json:"max,omitempty"`
	Amplitude int64   `protobuf:"varint,5,opt,name=amplitude,proto3" json:"amplitude,omitempty"`
	Mode      int64   `protobuf:"varint,6,opt,name=mode,proto3" json:"mode,omitempty"`
}

func (x *Sensor) Reset() {
	*x = Sensor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_transport_grpc_types_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sensor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sensor) ProtoMessage() {}

func (x *Sensor) ProtoReflect() protoreflect.Message {
	mi := &file_internal_transport_grpc_types_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sensor.ProtoReflect.Descriptor instead.
func (*Sensor) Descriptor() ([]byte, []int) {
	return file_internal_transport_grpc_types_proto_rawDescGZIP(), []int{9}
}

func (x *Sensor) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Sensor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Sensor) GetMin() float64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *Sensor) GetMax() float64 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *Sensor) GetAmplitude() int64 {
	if x != nil {
		return x.Amplitude
	}
	return 0
}

func (x *Sensor) GetMode() int64 {
	if x != nil {
		return x.Mode
	}
	return 0
}

type ClientInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp int64  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ClientInfo) Reset() {
	*x = ClientInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_transport_grpc_types_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientInfo) ProtoMessage() {}

func (x *ClientInfo) ProtoReflect() protoreflect.Message {
	mi := &file_internal_transport_grpc_types_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientInfo.ProtoReflect.Descriptor instead.
func (*ClientInfo) Descriptor() ([]byte, []int) {
	return file_internal_transport_grpc_types_proto_rawDescGZIP(), []int{10}
}

func (x *ClientInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ClientInfo) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_internal_transport_grpc_types_proto protoreflect.FileDescriptor

var file_internal_transport_grpc_types_proto_rawDesc = []byte{
	0x0a, 0x23, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0xb7, 0x04, 0x0a, 0x07,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x73, 0x63, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x73, 0x63, 0x72, 0x12, 0x27, 0x0a, 0x07,
	0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x07, 0x6f, 0x66,
	0x66, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x65, 0x6c, 0x65, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x45, 0x6c, 0x65, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x65, 0x6c, 0x65, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x07, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x61, 0x74,
	0x74, 0x65, 0x72, 0x79, 0x52, 0x07, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x12, 0x21, 0x0a,
	0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x5f, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6e, 0x75, 0x6d,
	0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x75, 0x6d, 0x5f, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6e, 0x75, 0x6d,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x6f,
	0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x6b,
	0x69, 0x70, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x75, 0x6e, 0x6e, 0x69,
	0x6e, 0x67, 0x5f, 0x61, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72, 0x75, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x70, 0x70, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x70,
	0x70, 0x65, 0x64, 0x41, 0x74, 0x22, 0x2d, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x22, 0x90, 0x01, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x73, 0x63, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x73, 0x63, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x2c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2d, 0x0a, 0x07, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6d,
	0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x6d, 0x61, 0x78, 0x22, 0x61, 0x0a, 0x09, 0x45, 0x6c, 0x65, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03,
	0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x4e, 0x0a, 0x07, 0x42, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x68, 0x61,
	0x72, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x49, 0x0a, 0x05, 0x53, 0x70, 0x65, 0x65, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6d,
	0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x03, 0x6d, 0x61, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x22, 0xc1, 0x02, 0x0a, 0x09, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c,
	0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x03, 0x6c, 0x61, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x44,
	0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0d, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x29, 0x0a, 0x10, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x73, 0x65, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x14, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6d, 0x70,
	0x6c, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x6d,
	0x70, 0x6c, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x3a, 0x0a, 0x0a, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x1e, 0x5a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_transport_grpc_types_proto_rawDescOnce sync.Once
	file_internal_transport_grpc_types_proto_rawDescData = file_internal_transport_grpc_types_proto_rawDesc
)

func file_internal_transport_grpc_types_proto_rawDescGZIP() []byte {
	file_internal_transport_grpc_types_proto_rawDescOnce.Do(func() {
		file_internal_transport_grpc_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_transport_grpc_types_proto_rawDescData)
	})
	return file_internal_transport_grpc_types_proto_rawDescData
}

var file_internal_transport_grpc_types_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_internal_transport_grpc_types_proto_goTypes = []interface{}{
	(*Tracker)(nil),    // 0: grpc.Tracker
	(*Error)(nil),      // 1: grpc.Error
	(*Filter)(nil),     // 2: grpc.Filter
	(*Status)(nil),     // 3: grpc.Status
	(*Offline)(nil),    // 4: grpc.Offline
	(*Elevation)(nil),  // 5: grpc.Elevation
	(*Battery)(nil),    // 6: grpc.Battery
	(*Speed)(nil),      // 7: grpc.Speed
	(*Navigator)(nil),  // 8: grpc.Navigator
	(*Sensor)(nil),     // 9: grpc.Sensor
	(*ClientInfo)(nil), // 10: grpc.ClientInfo
}
var file_internal_transport_grpc_types_proto_depIdxs = []int32{
	3, // 0: grpc.Tracker.status:type_name -> grpc.Status
	4, // 1: grpc.Tracker.offline:type_name -> grpc.Offline
	5, // 2: grpc.Tracker.elevation:type_name -> grpc.Elevation
	6, // 3: grpc.Tracker.battery:type_name -> grpc.Battery
	7, // 4: grpc.Tracker.speed:type_name -> grpc.Speed
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_internal_transport_grpc_types_proto_init() }
func file_internal_transport_grpc_types_proto_init() {
	if File_internal_transport_grpc_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_transport_grpc_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tracker); i {
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
		file_internal_transport_grpc_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_internal_transport_grpc_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
		file_internal_transport_grpc_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_internal_transport_grpc_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Offline); i {
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
		file_internal_transport_grpc_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Elevation); i {
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
		file_internal_transport_grpc_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Battery); i {
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
		file_internal_transport_grpc_types_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Speed); i {
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
		file_internal_transport_grpc_types_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Navigator); i {
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
		file_internal_transport_grpc_types_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sensor); i {
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
		file_internal_transport_grpc_types_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientInfo); i {
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
			RawDescriptor: file_internal_transport_grpc_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_transport_grpc_types_proto_goTypes,
		DependencyIndexes: file_internal_transport_grpc_types_proto_depIdxs,
		MessageInfos:      file_internal_transport_grpc_types_proto_msgTypes,
	}.Build()
	File_internal_transport_grpc_types_proto = out.File
	file_internal_transport_grpc_types_proto_rawDesc = nil
	file_internal_transport_grpc_types_proto_goTypes = nil
	file_internal_transport_grpc_types_proto_depIdxs = nil
}
