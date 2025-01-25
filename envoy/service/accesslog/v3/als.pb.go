// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.29.2
// source: envoy/service/accesslog/v3/als.proto

package accesslogv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v31 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Empty response for the StreamAccessLogs API. Will never be sent. See below.
type StreamAccessLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StreamAccessLogsResponse) Reset() {
	*x = StreamAccessLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_accesslog_v3_als_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamAccessLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamAccessLogsResponse) ProtoMessage() {}

func (x *StreamAccessLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_accesslog_v3_als_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamAccessLogsResponse.ProtoReflect.Descriptor instead.
func (*StreamAccessLogsResponse) Descriptor() ([]byte, []int) {
	return file_envoy_service_accesslog_v3_als_proto_rawDescGZIP(), []int{0}
}

// Stream message for the StreamAccessLogs API. Envoy will open a stream to the server and stream
// access logs without ever expecting a response.
type StreamAccessLogsMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier data that will only be sent in the first message on the stream. This is effectively
	// structured metadata and is a performance optimization.
	Identifier *StreamAccessLogsMessage_Identifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// Batches of log entries of a single type. Generally speaking, a given stream should only
	// ever include one type of log entry.
	//
	// Types that are assignable to LogEntries:
	//
	//	*StreamAccessLogsMessage_HttpLogs
	//	*StreamAccessLogsMessage_TcpLogs
	LogEntries isStreamAccessLogsMessage_LogEntries `protobuf_oneof:"log_entries"`
}

func (x *StreamAccessLogsMessage) Reset() {
	*x = StreamAccessLogsMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_accesslog_v3_als_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamAccessLogsMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamAccessLogsMessage) ProtoMessage() {}

func (x *StreamAccessLogsMessage) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_accesslog_v3_als_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamAccessLogsMessage.ProtoReflect.Descriptor instead.
func (*StreamAccessLogsMessage) Descriptor() ([]byte, []int) {
	return file_envoy_service_accesslog_v3_als_proto_rawDescGZIP(), []int{1}
}

func (x *StreamAccessLogsMessage) GetIdentifier() *StreamAccessLogsMessage_Identifier {
	if x != nil {
		return x.Identifier
	}
	return nil
}

func (m *StreamAccessLogsMessage) GetLogEntries() isStreamAccessLogsMessage_LogEntries {
	if m != nil {
		return m.LogEntries
	}
	return nil
}

func (x *StreamAccessLogsMessage) GetHttpLogs() *StreamAccessLogsMessage_HTTPAccessLogEntries {
	if x, ok := x.GetLogEntries().(*StreamAccessLogsMessage_HttpLogs); ok {
		return x.HttpLogs
	}
	return nil
}

func (x *StreamAccessLogsMessage) GetTcpLogs() *StreamAccessLogsMessage_TCPAccessLogEntries {
	if x, ok := x.GetLogEntries().(*StreamAccessLogsMessage_TcpLogs); ok {
		return x.TcpLogs
	}
	return nil
}

type isStreamAccessLogsMessage_LogEntries interface {
	isStreamAccessLogsMessage_LogEntries()
}

type StreamAccessLogsMessage_HttpLogs struct {
	HttpLogs *StreamAccessLogsMessage_HTTPAccessLogEntries `protobuf:"bytes,2,opt,name=http_logs,json=httpLogs,proto3,oneof"`
}

type StreamAccessLogsMessage_TcpLogs struct {
	TcpLogs *StreamAccessLogsMessage_TCPAccessLogEntries `protobuf:"bytes,3,opt,name=tcp_logs,json=tcpLogs,proto3,oneof"`
}

func (*StreamAccessLogsMessage_HttpLogs) isStreamAccessLogsMessage_LogEntries() {}

func (*StreamAccessLogsMessage_TcpLogs) isStreamAccessLogsMessage_LogEntries() {}

type StreamAccessLogsMessage_Identifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The node sending the access log messages over the stream.
	Node *v3.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// The friendly name of the log configured in :ref:`CommonGrpcAccessLogConfig
	// <envoy_v3_api_msg_extensions.access_loggers.grpc.v3.CommonGrpcAccessLogConfig>`.
	LogName string `protobuf:"bytes,2,opt,name=log_name,json=logName,proto3" json:"log_name,omitempty"`
}

func (x *StreamAccessLogsMessage_Identifier) Reset() {
	*x = StreamAccessLogsMessage_Identifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_accesslog_v3_als_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamAccessLogsMessage_Identifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamAccessLogsMessage_Identifier) ProtoMessage() {}

func (x *StreamAccessLogsMessage_Identifier) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_accesslog_v3_als_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamAccessLogsMessage_Identifier.ProtoReflect.Descriptor instead.
func (*StreamAccessLogsMessage_Identifier) Descriptor() ([]byte, []int) {
	return file_envoy_service_accesslog_v3_als_proto_rawDescGZIP(), []int{1, 0}
}

func (x *StreamAccessLogsMessage_Identifier) GetNode() *v3.Node {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *StreamAccessLogsMessage_Identifier) GetLogName() string {
	if x != nil {
		return x.LogName
	}
	return ""
}

// Wrapper for batches of HTTP access log entries.
type StreamAccessLogsMessage_HTTPAccessLogEntries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogEntry []*v31.HTTPAccessLogEntry `protobuf:"bytes,1,rep,name=log_entry,json=logEntry,proto3" json:"log_entry,omitempty"`
}

func (x *StreamAccessLogsMessage_HTTPAccessLogEntries) Reset() {
	*x = StreamAccessLogsMessage_HTTPAccessLogEntries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_accesslog_v3_als_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamAccessLogsMessage_HTTPAccessLogEntries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamAccessLogsMessage_HTTPAccessLogEntries) ProtoMessage() {}

func (x *StreamAccessLogsMessage_HTTPAccessLogEntries) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_accesslog_v3_als_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamAccessLogsMessage_HTTPAccessLogEntries.ProtoReflect.Descriptor instead.
func (*StreamAccessLogsMessage_HTTPAccessLogEntries) Descriptor() ([]byte, []int) {
	return file_envoy_service_accesslog_v3_als_proto_rawDescGZIP(), []int{1, 1}
}

func (x *StreamAccessLogsMessage_HTTPAccessLogEntries) GetLogEntry() []*v31.HTTPAccessLogEntry {
	if x != nil {
		return x.LogEntry
	}
	return nil
}

// Wrapper for batches of TCP access log entries.
type StreamAccessLogsMessage_TCPAccessLogEntries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogEntry []*v31.TCPAccessLogEntry `protobuf:"bytes,1,rep,name=log_entry,json=logEntry,proto3" json:"log_entry,omitempty"`
}

func (x *StreamAccessLogsMessage_TCPAccessLogEntries) Reset() {
	*x = StreamAccessLogsMessage_TCPAccessLogEntries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_accesslog_v3_als_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamAccessLogsMessage_TCPAccessLogEntries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamAccessLogsMessage_TCPAccessLogEntries) ProtoMessage() {}

func (x *StreamAccessLogsMessage_TCPAccessLogEntries) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_accesslog_v3_als_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamAccessLogsMessage_TCPAccessLogEntries.ProtoReflect.Descriptor instead.
func (*StreamAccessLogsMessage_TCPAccessLogEntries) Descriptor() ([]byte, []int) {
	return file_envoy_service_accesslog_v3_als_proto_rawDescGZIP(), []int{1, 2}
}

func (x *StreamAccessLogsMessage_TCPAccessLogEntries) GetLogEntry() []*v31.TCPAccessLogEntry {
	if x != nil {
		return x.LogEntry
	}
	return nil
}

var File_envoy_service_accesslog_v3_als_proto protoreflect.FileDescriptor

var file_envoy_service_accesslog_v3_als_proto_rawDesc = []byte{
	0x0a, 0x24, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x6f, 0x67, 0x2e,
	0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64,
	0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x18, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x3a, 0x3a, 0x9a, 0xc5, 0x88, 0x1e, 0x35, 0x0a, 0x33, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0xc1, 0x07, 0x0a, 0x17, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4c, 0x6f, 0x67, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x5e, 0x0a, 0x0a, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x3e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52,
	0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x67, 0x0a, 0x09, 0x68,
	0x74, 0x74, 0x70, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x48,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f,
	0x67, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x48, 0x00, 0x52, 0x08, 0x68, 0x74, 0x74, 0x70,
	0x4c, 0x6f, 0x67, 0x73, 0x12, 0x64, 0x0a, 0x08, 0x74, 0x63, 0x70, 0x5f, 0x6c, 0x6f, 0x67, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x6f, 0x67,
	0x2e, 0x76, 0x33, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4c, 0x6f, 0x67, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x43, 0x50, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x48,
	0x00, 0x52, 0x07, 0x74, 0x63, 0x70, 0x4c, 0x6f, 0x67, 0x73, 0x1a, 0xb0, 0x01, 0x0a, 0x0a, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x04, 0x6e, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e,
	0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07,
	0x6c, 0x6f, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x3a, 0x44, 0x9a, 0xc5, 0x88, 0x1e, 0x3f, 0x0a, 0x3d,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x1a, 0xba, 0x01,
	0x0a, 0x14, 0x48, 0x54, 0x54, 0x50, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x45,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x52, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x6f, 0x67,
	0x2e, 0x76, 0x33, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f,
	0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01,
	0x52, 0x08, 0x6c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x3a, 0x4e, 0x9a, 0xc5, 0x88, 0x1e,
	0x49, 0x0a, 0x47, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x1a, 0xb7, 0x01, 0x0a, 0x13, 0x54,
	0x43, 0x50, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x51, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x33, 0x2e,
	0x54, 0x43, 0x50, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x08, 0x6c, 0x6f, 0x67,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x3a, 0x4d, 0x9a, 0xc5, 0x88, 0x1e, 0x48, 0x0a, 0x46, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x54, 0x43, 0x50, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x3a, 0x39, 0x9a, 0xc5, 0x88, 0x1e, 0x34, 0x0a, 0x32, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42,
	0x12, 0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x03,
	0xf8, 0x42, 0x01, 0x32, 0x96, 0x01, 0x0a, 0x10, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x10, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x33, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x34, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x33, 0x2e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42, 0x8d, 0x01, 0xba,
	0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x28, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x76,
	0x33, 0x42, 0x08, 0x41, 0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x33,
	0x3b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x6f, 0x67, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_service_accesslog_v3_als_proto_rawDescOnce sync.Once
	file_envoy_service_accesslog_v3_als_proto_rawDescData = file_envoy_service_accesslog_v3_als_proto_rawDesc
)

func file_envoy_service_accesslog_v3_als_proto_rawDescGZIP() []byte {
	file_envoy_service_accesslog_v3_als_proto_rawDescOnce.Do(func() {
		file_envoy_service_accesslog_v3_als_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_service_accesslog_v3_als_proto_rawDescData)
	})
	return file_envoy_service_accesslog_v3_als_proto_rawDescData
}

var file_envoy_service_accesslog_v3_als_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_envoy_service_accesslog_v3_als_proto_goTypes = []interface{}{
	(*StreamAccessLogsResponse)(nil),                     // 0: envoy.service.accesslog.v3.StreamAccessLogsResponse
	(*StreamAccessLogsMessage)(nil),                      // 1: envoy.service.accesslog.v3.StreamAccessLogsMessage
	(*StreamAccessLogsMessage_Identifier)(nil),           // 2: envoy.service.accesslog.v3.StreamAccessLogsMessage.Identifier
	(*StreamAccessLogsMessage_HTTPAccessLogEntries)(nil), // 3: envoy.service.accesslog.v3.StreamAccessLogsMessage.HTTPAccessLogEntries
	(*StreamAccessLogsMessage_TCPAccessLogEntries)(nil),  // 4: envoy.service.accesslog.v3.StreamAccessLogsMessage.TCPAccessLogEntries
	(*v3.Node)(nil),                // 5: envoy.config.core.v3.Node
	(*v31.HTTPAccessLogEntry)(nil), // 6: envoy.data.accesslog.v3.HTTPAccessLogEntry
	(*v31.TCPAccessLogEntry)(nil),  // 7: envoy.data.accesslog.v3.TCPAccessLogEntry
}
var file_envoy_service_accesslog_v3_als_proto_depIdxs = []int32{
	2, // 0: envoy.service.accesslog.v3.StreamAccessLogsMessage.identifier:type_name -> envoy.service.accesslog.v3.StreamAccessLogsMessage.Identifier
	3, // 1: envoy.service.accesslog.v3.StreamAccessLogsMessage.http_logs:type_name -> envoy.service.accesslog.v3.StreamAccessLogsMessage.HTTPAccessLogEntries
	4, // 2: envoy.service.accesslog.v3.StreamAccessLogsMessage.tcp_logs:type_name -> envoy.service.accesslog.v3.StreamAccessLogsMessage.TCPAccessLogEntries
	5, // 3: envoy.service.accesslog.v3.StreamAccessLogsMessage.Identifier.node:type_name -> envoy.config.core.v3.Node
	6, // 4: envoy.service.accesslog.v3.StreamAccessLogsMessage.HTTPAccessLogEntries.log_entry:type_name -> envoy.data.accesslog.v3.HTTPAccessLogEntry
	7, // 5: envoy.service.accesslog.v3.StreamAccessLogsMessage.TCPAccessLogEntries.log_entry:type_name -> envoy.data.accesslog.v3.TCPAccessLogEntry
	1, // 6: envoy.service.accesslog.v3.AccessLogService.StreamAccessLogs:input_type -> envoy.service.accesslog.v3.StreamAccessLogsMessage
	0, // 7: envoy.service.accesslog.v3.AccessLogService.StreamAccessLogs:output_type -> envoy.service.accesslog.v3.StreamAccessLogsResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_envoy_service_accesslog_v3_als_proto_init() }
func file_envoy_service_accesslog_v3_als_proto_init() {
	if File_envoy_service_accesslog_v3_als_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_service_accesslog_v3_als_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamAccessLogsResponse); i {
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
		file_envoy_service_accesslog_v3_als_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamAccessLogsMessage); i {
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
		file_envoy_service_accesslog_v3_als_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamAccessLogsMessage_Identifier); i {
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
		file_envoy_service_accesslog_v3_als_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamAccessLogsMessage_HTTPAccessLogEntries); i {
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
		file_envoy_service_accesslog_v3_als_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamAccessLogsMessage_TCPAccessLogEntries); i {
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
	file_envoy_service_accesslog_v3_als_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*StreamAccessLogsMessage_HttpLogs)(nil),
		(*StreamAccessLogsMessage_TcpLogs)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_service_accesslog_v3_als_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_envoy_service_accesslog_v3_als_proto_goTypes,
		DependencyIndexes: file_envoy_service_accesslog_v3_als_proto_depIdxs,
		MessageInfos:      file_envoy_service_accesslog_v3_als_proto_msgTypes,
	}.Build()
	File_envoy_service_accesslog_v3_als_proto = out.File
	file_envoy_service_accesslog_v3_als_proto_rawDesc = nil
	file_envoy_service_accesslog_v3_als_proto_goTypes = nil
	file_envoy_service_accesslog_v3_als_proto_depIdxs = nil
}
