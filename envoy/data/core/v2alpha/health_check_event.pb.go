// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.29.3
// source: envoy/data/core/v2alpha/health_check_event.proto

package v2alpha

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type HealthCheckFailureType int32

const (
	HealthCheckFailureType_ACTIVE  HealthCheckFailureType = 0
	HealthCheckFailureType_PASSIVE HealthCheckFailureType = 1
	HealthCheckFailureType_NETWORK HealthCheckFailureType = 2
)

// Enum value maps for HealthCheckFailureType.
var (
	HealthCheckFailureType_name = map[int32]string{
		0: "ACTIVE",
		1: "PASSIVE",
		2: "NETWORK",
	}
	HealthCheckFailureType_value = map[string]int32{
		"ACTIVE":  0,
		"PASSIVE": 1,
		"NETWORK": 2,
	}
)

func (x HealthCheckFailureType) Enum() *HealthCheckFailureType {
	p := new(HealthCheckFailureType)
	*p = x
	return p
}

func (x HealthCheckFailureType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HealthCheckFailureType) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_data_core_v2alpha_health_check_event_proto_enumTypes[0].Descriptor()
}

func (HealthCheckFailureType) Type() protoreflect.EnumType {
	return &file_envoy_data_core_v2alpha_health_check_event_proto_enumTypes[0]
}

func (x HealthCheckFailureType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HealthCheckFailureType.Descriptor instead.
func (HealthCheckFailureType) EnumDescriptor() ([]byte, []int) {
	return file_envoy_data_core_v2alpha_health_check_event_proto_rawDescGZIP(), []int{0}
}

type HealthCheckerType int32

const (
	HealthCheckerType_HTTP  HealthCheckerType = 0
	HealthCheckerType_TCP   HealthCheckerType = 1
	HealthCheckerType_GRPC  HealthCheckerType = 2
	HealthCheckerType_REDIS HealthCheckerType = 3
)

// Enum value maps for HealthCheckerType.
var (
	HealthCheckerType_name = map[int32]string{
		0: "HTTP",
		1: "TCP",
		2: "GRPC",
		3: "REDIS",
	}
	HealthCheckerType_value = map[string]int32{
		"HTTP":  0,
		"TCP":   1,
		"GRPC":  2,
		"REDIS": 3,
	}
)

func (x HealthCheckerType) Enum() *HealthCheckerType {
	p := new(HealthCheckerType)
	*p = x
	return p
}

func (x HealthCheckerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HealthCheckerType) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_data_core_v2alpha_health_check_event_proto_enumTypes[1].Descriptor()
}

func (HealthCheckerType) Type() protoreflect.EnumType {
	return &file_envoy_data_core_v2alpha_health_check_event_proto_enumTypes[1]
}

func (x HealthCheckerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HealthCheckerType.Descriptor instead.
func (HealthCheckerType) EnumDescriptor() ([]byte, []int) {
	return file_envoy_data_core_v2alpha_health_check_event_proto_rawDescGZIP(), []int{1}
}

// [#next-free-field: 10]
type HealthCheckEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HealthCheckerType HealthCheckerType `protobuf:"varint,1,opt,name=health_checker_type,json=healthCheckerType,proto3,enum=envoy.data.core.v2alpha.HealthCheckerType" json:"health_checker_type,omitempty"`
	Host              *core.Address     `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	ClusterName       string            `protobuf:"bytes,3,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// Types that are assignable to Event:
	//
	//	*HealthCheckEvent_EjectUnhealthyEvent
	//	*HealthCheckEvent_AddHealthyEvent
	//	*HealthCheckEvent_HealthCheckFailureEvent
	//	*HealthCheckEvent_DegradedHealthyHost
	//	*HealthCheckEvent_NoLongerDegradedHost
	Event isHealthCheckEvent_Event `protobuf_oneof:"event"`
	// Timestamp for event.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *HealthCheckEvent) Reset() {
	*x = HealthCheckEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_core_v2alpha_health_check_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckEvent) ProtoMessage() {}

func (x *HealthCheckEvent) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_core_v2alpha_health_check_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckEvent.ProtoReflect.Descriptor instead.
func (*HealthCheckEvent) Descriptor() ([]byte, []int) {
	return file_envoy_data_core_v2alpha_health_check_event_proto_rawDescGZIP(), []int{0}
}

func (x *HealthCheckEvent) GetHealthCheckerType() HealthCheckerType {
	if x != nil {
		return x.HealthCheckerType
	}
	return HealthCheckerType_HTTP
}

func (x *HealthCheckEvent) GetHost() *core.Address {
	if x != nil {
		return x.Host
	}
	return nil
}

func (x *HealthCheckEvent) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (m *HealthCheckEvent) GetEvent() isHealthCheckEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *HealthCheckEvent) GetEjectUnhealthyEvent() *HealthCheckEjectUnhealthy {
	if x, ok := x.GetEvent().(*HealthCheckEvent_EjectUnhealthyEvent); ok {
		return x.EjectUnhealthyEvent
	}
	return nil
}

func (x *HealthCheckEvent) GetAddHealthyEvent() *HealthCheckAddHealthy {
	if x, ok := x.GetEvent().(*HealthCheckEvent_AddHealthyEvent); ok {
		return x.AddHealthyEvent
	}
	return nil
}

func (x *HealthCheckEvent) GetHealthCheckFailureEvent() *HealthCheckFailure {
	if x, ok := x.GetEvent().(*HealthCheckEvent_HealthCheckFailureEvent); ok {
		return x.HealthCheckFailureEvent
	}
	return nil
}

func (x *HealthCheckEvent) GetDegradedHealthyHost() *DegradedHealthyHost {
	if x, ok := x.GetEvent().(*HealthCheckEvent_DegradedHealthyHost); ok {
		return x.DegradedHealthyHost
	}
	return nil
}

func (x *HealthCheckEvent) GetNoLongerDegradedHost() *NoLongerDegradedHost {
	if x, ok := x.GetEvent().(*HealthCheckEvent_NoLongerDegradedHost); ok {
		return x.NoLongerDegradedHost
	}
	return nil
}

func (x *HealthCheckEvent) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type isHealthCheckEvent_Event interface {
	isHealthCheckEvent_Event()
}

type HealthCheckEvent_EjectUnhealthyEvent struct {
	// Host ejection.
	EjectUnhealthyEvent *HealthCheckEjectUnhealthy `protobuf:"bytes,4,opt,name=eject_unhealthy_event,json=ejectUnhealthyEvent,proto3,oneof"`
}

type HealthCheckEvent_AddHealthyEvent struct {
	// Host addition.
	AddHealthyEvent *HealthCheckAddHealthy `protobuf:"bytes,5,opt,name=add_healthy_event,json=addHealthyEvent,proto3,oneof"`
}

type HealthCheckEvent_HealthCheckFailureEvent struct {
	// Host failure.
	HealthCheckFailureEvent *HealthCheckFailure `protobuf:"bytes,7,opt,name=health_check_failure_event,json=healthCheckFailureEvent,proto3,oneof"`
}

type HealthCheckEvent_DegradedHealthyHost struct {
	// Healthy host became degraded.
	DegradedHealthyHost *DegradedHealthyHost `protobuf:"bytes,8,opt,name=degraded_healthy_host,json=degradedHealthyHost,proto3,oneof"`
}

type HealthCheckEvent_NoLongerDegradedHost struct {
	// A degraded host returned to being healthy.
	NoLongerDegradedHost *NoLongerDegradedHost `protobuf:"bytes,9,opt,name=no_longer_degraded_host,json=noLongerDegradedHost,proto3,oneof"`
}

func (*HealthCheckEvent_EjectUnhealthyEvent) isHealthCheckEvent_Event() {}

func (*HealthCheckEvent_AddHealthyEvent) isHealthCheckEvent_Event() {}

func (*HealthCheckEvent_HealthCheckFailureEvent) isHealthCheckEvent_Event() {}

func (*HealthCheckEvent_DegradedHealthyHost) isHealthCheckEvent_Event() {}

func (*HealthCheckEvent_NoLongerDegradedHost) isHealthCheckEvent_Event() {}

type HealthCheckEjectUnhealthy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of failure that caused this ejection.
	FailureType HealthCheckFailureType `protobuf:"varint,1,opt,name=failure_type,json=failureType,proto3,enum=envoy.data.core.v2alpha.HealthCheckFailureType" json:"failure_type,omitempty"`
}

func (x *HealthCheckEjectUnhealthy) Reset() {
	*x = HealthCheckEjectUnhealthy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_core_v2alpha_health_check_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckEjectUnhealthy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckEjectUnhealthy) ProtoMessage() {}

func (x *HealthCheckEjectUnhealthy) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_core_v2alpha_health_check_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckEjectUnhealthy.ProtoReflect.Descriptor instead.
func (*HealthCheckEjectUnhealthy) Descriptor() ([]byte, []int) {
	return file_envoy_data_core_v2alpha_health_check_event_proto_rawDescGZIP(), []int{1}
}

func (x *HealthCheckEjectUnhealthy) GetFailureType() HealthCheckFailureType {
	if x != nil {
		return x.FailureType
	}
	return HealthCheckFailureType_ACTIVE
}

type HealthCheckAddHealthy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether this addition is the result of the first ever health check on a host, in which case
	// the configured :ref:`healthy threshold <envoy_api_field_core.HealthCheck.healthy_threshold>`
	// is bypassed and the host is immediately added.
	FirstCheck bool `protobuf:"varint,1,opt,name=first_check,json=firstCheck,proto3" json:"first_check,omitempty"`
}

func (x *HealthCheckAddHealthy) Reset() {
	*x = HealthCheckAddHealthy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_core_v2alpha_health_check_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckAddHealthy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckAddHealthy) ProtoMessage() {}

func (x *HealthCheckAddHealthy) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_core_v2alpha_health_check_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckAddHealthy.ProtoReflect.Descriptor instead.
func (*HealthCheckAddHealthy) Descriptor() ([]byte, []int) {
	return file_envoy_data_core_v2alpha_health_check_event_proto_rawDescGZIP(), []int{2}
}

func (x *HealthCheckAddHealthy) GetFirstCheck() bool {
	if x != nil {
		return x.FirstCheck
	}
	return false
}

type HealthCheckFailure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of failure that caused this event.
	FailureType HealthCheckFailureType `protobuf:"varint,1,opt,name=failure_type,json=failureType,proto3,enum=envoy.data.core.v2alpha.HealthCheckFailureType" json:"failure_type,omitempty"`
	// Whether this event is the result of the first ever health check on a host.
	FirstCheck bool `protobuf:"varint,2,opt,name=first_check,json=firstCheck,proto3" json:"first_check,omitempty"`
}

func (x *HealthCheckFailure) Reset() {
	*x = HealthCheckFailure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_core_v2alpha_health_check_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckFailure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckFailure) ProtoMessage() {}

func (x *HealthCheckFailure) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_core_v2alpha_health_check_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckFailure.ProtoReflect.Descriptor instead.
func (*HealthCheckFailure) Descriptor() ([]byte, []int) {
	return file_envoy_data_core_v2alpha_health_check_event_proto_rawDescGZIP(), []int{3}
}

func (x *HealthCheckFailure) GetFailureType() HealthCheckFailureType {
	if x != nil {
		return x.FailureType
	}
	return HealthCheckFailureType_ACTIVE
}

func (x *HealthCheckFailure) GetFirstCheck() bool {
	if x != nil {
		return x.FirstCheck
	}
	return false
}

type DegradedHealthyHost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DegradedHealthyHost) Reset() {
	*x = DegradedHealthyHost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_core_v2alpha_health_check_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DegradedHealthyHost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DegradedHealthyHost) ProtoMessage() {}

func (x *DegradedHealthyHost) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_core_v2alpha_health_check_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DegradedHealthyHost.ProtoReflect.Descriptor instead.
func (*DegradedHealthyHost) Descriptor() ([]byte, []int) {
	return file_envoy_data_core_v2alpha_health_check_event_proto_rawDescGZIP(), []int{4}
}

type NoLongerDegradedHost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoLongerDegradedHost) Reset() {
	*x = NoLongerDegradedHost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_core_v2alpha_health_check_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoLongerDegradedHost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoLongerDegradedHost) ProtoMessage() {}

func (x *NoLongerDegradedHost) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_core_v2alpha_health_check_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoLongerDegradedHost.ProtoReflect.Descriptor instead.
func (*NoLongerDegradedHost) Descriptor() ([]byte, []int) {
	return file_envoy_data_core_v2alpha_health_check_event_proto_rawDescGZIP(), []int{5}
}

var File_envoy_data_core_v2alpha_health_check_event_proto protoreflect.FileDescriptor

var file_envoy_data_core_v2alpha_health_check_event_proto_rawDesc = []byte{
	0x0a, 0x30, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x17, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x06, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x64, 0x0a, 0x13, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x11, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x2e, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x12, 0x2a, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52,
	0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x68, 0x0a, 0x15,
	0x65, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x6e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x5f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x45, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x6e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x48,
	0x00, 0x52, 0x13, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x6e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x5c, 0x0a, 0x11, 0x61, 0x64, 0x64, 0x5f, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x79, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x64, 0x64, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x79, 0x48, 0x00, 0x52, 0x0f, 0x61, 0x64, 0x64, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x6a, 0x0a, 0x1a, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x46, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x48, 0x00, 0x52, 0x17, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x62, 0x0a, 0x15, 0x64, 0x65, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64, 0x5f, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x79, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x65, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x64, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x48, 0x6f, 0x73, 0x74, 0x48, 0x00, 0x52,
	0x13, 0x64, 0x65, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79,
	0x48, 0x6f, 0x73, 0x74, 0x12, 0x66, 0x0a, 0x17, 0x6e, 0x6f, 0x5f, 0x6c, 0x6f, 0x6e, 0x67, 0x65,
	0x72, 0x5f, 0x64, 0x65, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x4e, 0x6f, 0x4c, 0x6f, 0x6e, 0x67, 0x65, 0x72, 0x44, 0x65, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64,
	0x48, 0x6f, 0x73, 0x74, 0x48, 0x00, 0x52, 0x14, 0x6e, 0x6f, 0x4c, 0x6f, 0x6e, 0x67, 0x65, 0x72,
	0x44, 0x65, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0c, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x03, 0xf8, 0x42, 0x01, 0x22, 0x79, 0x0a, 0x19, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x45, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x6e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x79, 0x12, 0x5c, 0x0a, 0x0c, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x46, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x38, 0x0a, 0x15, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x64,
	0x64, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x22, 0x93, 0x01, 0x0a, 0x12, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x12, 0x5c, 0x0a, 0x0c, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x46, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x22,
	0x15, 0x0a, 0x13, 0x44, 0x65, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x79, 0x48, 0x6f, 0x73, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x4e, 0x6f, 0x4c, 0x6f, 0x6e, 0x67,
	0x65, 0x72, 0x44, 0x65, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64, 0x48, 0x6f, 0x73, 0x74, 0x2a, 0x3e,
	0x0a, 0x16, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x46, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49,
	0x56, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x41, 0x53, 0x53, 0x49, 0x56, 0x45, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x02, 0x2a, 0x3b,
	0x0a, 0x11, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x54, 0x54, 0x50, 0x10, 0x00, 0x12, 0x07, 0x0a,
	0x03, 0x54, 0x43, 0x50, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x47, 0x52, 0x50, 0x43, 0x10, 0x02,
	0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x44, 0x49, 0x53, 0x10, 0x03, 0x42, 0x88, 0x01, 0xba, 0x80,
	0xc8, 0xd1, 0x06, 0x02, 0x10, 0x01, 0x0a, 0x25, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x15, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f,
	0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_data_core_v2alpha_health_check_event_proto_rawDescOnce sync.Once
	file_envoy_data_core_v2alpha_health_check_event_proto_rawDescData = file_envoy_data_core_v2alpha_health_check_event_proto_rawDesc
)

func file_envoy_data_core_v2alpha_health_check_event_proto_rawDescGZIP() []byte {
	file_envoy_data_core_v2alpha_health_check_event_proto_rawDescOnce.Do(func() {
		file_envoy_data_core_v2alpha_health_check_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_data_core_v2alpha_health_check_event_proto_rawDescData)
	})
	return file_envoy_data_core_v2alpha_health_check_event_proto_rawDescData
}

var file_envoy_data_core_v2alpha_health_check_event_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_envoy_data_core_v2alpha_health_check_event_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_envoy_data_core_v2alpha_health_check_event_proto_goTypes = []interface{}{
	(HealthCheckFailureType)(0),       // 0: envoy.data.core.v2alpha.HealthCheckFailureType
	(HealthCheckerType)(0),            // 1: envoy.data.core.v2alpha.HealthCheckerType
	(*HealthCheckEvent)(nil),          // 2: envoy.data.core.v2alpha.HealthCheckEvent
	(*HealthCheckEjectUnhealthy)(nil), // 3: envoy.data.core.v2alpha.HealthCheckEjectUnhealthy
	(*HealthCheckAddHealthy)(nil),     // 4: envoy.data.core.v2alpha.HealthCheckAddHealthy
	(*HealthCheckFailure)(nil),        // 5: envoy.data.core.v2alpha.HealthCheckFailure
	(*DegradedHealthyHost)(nil),       // 6: envoy.data.core.v2alpha.DegradedHealthyHost
	(*NoLongerDegradedHost)(nil),      // 7: envoy.data.core.v2alpha.NoLongerDegradedHost
	(*core.Address)(nil),              // 8: envoy.api.v2.core.Address
	(*timestamppb.Timestamp)(nil),     // 9: google.protobuf.Timestamp
}
var file_envoy_data_core_v2alpha_health_check_event_proto_depIdxs = []int32{
	1,  // 0: envoy.data.core.v2alpha.HealthCheckEvent.health_checker_type:type_name -> envoy.data.core.v2alpha.HealthCheckerType
	8,  // 1: envoy.data.core.v2alpha.HealthCheckEvent.host:type_name -> envoy.api.v2.core.Address
	3,  // 2: envoy.data.core.v2alpha.HealthCheckEvent.eject_unhealthy_event:type_name -> envoy.data.core.v2alpha.HealthCheckEjectUnhealthy
	4,  // 3: envoy.data.core.v2alpha.HealthCheckEvent.add_healthy_event:type_name -> envoy.data.core.v2alpha.HealthCheckAddHealthy
	5,  // 4: envoy.data.core.v2alpha.HealthCheckEvent.health_check_failure_event:type_name -> envoy.data.core.v2alpha.HealthCheckFailure
	6,  // 5: envoy.data.core.v2alpha.HealthCheckEvent.degraded_healthy_host:type_name -> envoy.data.core.v2alpha.DegradedHealthyHost
	7,  // 6: envoy.data.core.v2alpha.HealthCheckEvent.no_longer_degraded_host:type_name -> envoy.data.core.v2alpha.NoLongerDegradedHost
	9,  // 7: envoy.data.core.v2alpha.HealthCheckEvent.timestamp:type_name -> google.protobuf.Timestamp
	0,  // 8: envoy.data.core.v2alpha.HealthCheckEjectUnhealthy.failure_type:type_name -> envoy.data.core.v2alpha.HealthCheckFailureType
	0,  // 9: envoy.data.core.v2alpha.HealthCheckFailure.failure_type:type_name -> envoy.data.core.v2alpha.HealthCheckFailureType
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_envoy_data_core_v2alpha_health_check_event_proto_init() }
func file_envoy_data_core_v2alpha_health_check_event_proto_init() {
	if File_envoy_data_core_v2alpha_health_check_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_data_core_v2alpha_health_check_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckEvent); i {
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
		file_envoy_data_core_v2alpha_health_check_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckEjectUnhealthy); i {
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
		file_envoy_data_core_v2alpha_health_check_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckAddHealthy); i {
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
		file_envoy_data_core_v2alpha_health_check_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckFailure); i {
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
		file_envoy_data_core_v2alpha_health_check_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DegradedHealthyHost); i {
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
		file_envoy_data_core_v2alpha_health_check_event_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoLongerDegradedHost); i {
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
	file_envoy_data_core_v2alpha_health_check_event_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*HealthCheckEvent_EjectUnhealthyEvent)(nil),
		(*HealthCheckEvent_AddHealthyEvent)(nil),
		(*HealthCheckEvent_HealthCheckFailureEvent)(nil),
		(*HealthCheckEvent_DegradedHealthyHost)(nil),
		(*HealthCheckEvent_NoLongerDegradedHost)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_data_core_v2alpha_health_check_event_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_data_core_v2alpha_health_check_event_proto_goTypes,
		DependencyIndexes: file_envoy_data_core_v2alpha_health_check_event_proto_depIdxs,
		EnumInfos:         file_envoy_data_core_v2alpha_health_check_event_proto_enumTypes,
		MessageInfos:      file_envoy_data_core_v2alpha_health_check_event_proto_msgTypes,
	}.Build()
	File_envoy_data_core_v2alpha_health_check_event_proto = out.File
	file_envoy_data_core_v2alpha_health_check_event_proto_rawDesc = nil
	file_envoy_data_core_v2alpha_health_check_event_proto_goTypes = nil
	file_envoy_data_core_v2alpha_health_check_event_proto_depIdxs = nil
}
