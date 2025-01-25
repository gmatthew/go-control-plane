// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.29.2
// source: envoy/extensions/bootstrap/internal_listener/v3/internal_listener.proto

package internal_listenerv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Configuration for internal listener.
type InternalListener struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The internal listener client connection buffer size in KiB.
	// For example, if “buffer_size_kb“ is set to 5, then the actual buffer size is
	// 5 KiB = 5 * 1024 bytes.
	// If the “buffer_size_kb“ is not specified, the buffer size is set to 1024 KiB.
	BufferSizeKb *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=buffer_size_kb,json=bufferSizeKb,proto3" json:"buffer_size_kb,omitempty"`
}

func (x *InternalListener) Reset() {
	*x = InternalListener{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_bootstrap_internal_listener_v3_internal_listener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalListener) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalListener) ProtoMessage() {}

func (x *InternalListener) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_bootstrap_internal_listener_v3_internal_listener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalListener.ProtoReflect.Descriptor instead.
func (*InternalListener) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_bootstrap_internal_listener_v3_internal_listener_proto_rawDescGZIP(), []int{0}
}

func (x *InternalListener) GetBufferSizeKb() *wrapperspb.UInt32Value {
	if x != nil {
		return x.BufferSizeKb
	}
	return nil
}

var File_envoy_extensions_bootstrap_internal_listener_v3_internal_listener_proto protoreflect.FileDescriptor

var file_envoy_extensions_bootstrap_internal_listener_v3_internal_listener_proto_rawDesc = []byte{
	0x0a, 0x47, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x76,
	0x33, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x62, 0x6f, 0x6f, 0x74,
	0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x62, 0x0a, 0x10, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4c, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x4e, 0x0a, 0x0e, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6b, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0a, 0xfa, 0x42,
	0x07, 0x2a, 0x05, 0x18, 0x80, 0x40, 0x28, 0x01, 0x52, 0x0c, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72,
	0x53, 0x69, 0x7a, 0x65, 0x4b, 0x62, 0x42, 0xcc, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10,
	0x02, 0x0a, 0x3d, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x33,
	0x42, 0x15, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x76,
	0x33, 0x3b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_bootstrap_internal_listener_v3_internal_listener_proto_rawDescOnce sync.Once
	file_envoy_extensions_bootstrap_internal_listener_v3_internal_listener_proto_rawDescData = file_envoy_extensions_bootstrap_internal_listener_v3_internal_listener_proto_rawDesc
)

func file_envoy_extensions_bootstrap_internal_listener_v3_internal_listener_proto_rawDescGZIP() []byte {
	file_envoy_extensions_bootstrap_internal_listener_v3_internal_listener_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_bootstrap_internal_listener_v3_internal_listener_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_bootstrap_internal_listener_v3_internal_listener_proto_rawDescData)
	})
	return file_envoy_extensions_bootstrap_internal_listener_v3_internal_listener_proto_rawDescData
}

var file_envoy_extensions_bootstrap_internal_listener_v3_internal_listener_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_bootstrap_internal_listener_v3_internal_listener_proto_goTypes = []interface{}{
	(*InternalListener)(nil),       // 0: envoy.extensions.bootstrap.internal_listener.v3.InternalListener
	(*wrapperspb.UInt32Value)(nil), // 1: google.protobuf.UInt32Value
}
var file_envoy_extensions_bootstrap_internal_listener_v3_internal_listener_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.bootstrap.internal_listener.v3.InternalListener.buffer_size_kb:type_name -> google.protobuf.UInt32Value
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_extensions_bootstrap_internal_listener_v3_internal_listener_proto_init() }
func file_envoy_extensions_bootstrap_internal_listener_v3_internal_listener_proto_init() {
	if File_envoy_extensions_bootstrap_internal_listener_v3_internal_listener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_bootstrap_internal_listener_v3_internal_listener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalListener); i {
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
			RawDescriptor: file_envoy_extensions_bootstrap_internal_listener_v3_internal_listener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_bootstrap_internal_listener_v3_internal_listener_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_bootstrap_internal_listener_v3_internal_listener_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_bootstrap_internal_listener_v3_internal_listener_proto_msgTypes,
	}.Build()
	File_envoy_extensions_bootstrap_internal_listener_v3_internal_listener_proto = out.File
	file_envoy_extensions_bootstrap_internal_listener_v3_internal_listener_proto_rawDesc = nil
	file_envoy_extensions_bootstrap_internal_listener_v3_internal_listener_proto_goTypes = nil
	file_envoy_extensions_bootstrap_internal_listener_v3_internal_listener_proto_depIdxs = nil
}
