// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.29.2
// source: envoy/extensions/load_balancing_policies/client_side_weighted_round_robin/v3/client_side_weighted_round_robin.proto

package client_side_weighted_round_robinv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

// Configuration for the client_side_weighted_round_robin LB policy.
//
// This policy differs from the built-in ROUND_ROBIN policy in terms of
// how the endpoint weights are determined. In the ROUND_ROBIN policy,
// the endpoint weights are sent by the control plane via EDS. However,
// in this policy, the endpoint weights are instead determined via qps (queries
// per second), eps (errors per second), and utilization metrics sent by the
// endpoint using the Open Request Cost Aggregation (ORCA) protocol. Utilization
// is determined by using the ORCA application_utilization field, if set, or
// else falling back to the cpu_utilization field. All queries count toward qps,
// regardless of result. Only failed queries count toward eps. A config
// parameter error_utilization_penalty controls the penalty to adjust endpoint
// weights using eps and qps. The weight of a given endpoint is computed as:
// “qps / (utilization + eps/qps * error_utilization_penalty)“.
//
// See the :ref:`load balancing architecture
// overview<arch_overview_load_balancing_types>` for more information.
//
// [#next-free-field: 8]
type ClientSideWeightedRoundRobin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether to enable out-of-band utilization reporting collection from
	// the endpoints. By default, per-request utilization reporting is used.
	EnableOobLoadReport *wrapperspb.BoolValue `protobuf:"bytes,1,opt,name=enable_oob_load_report,json=enableOobLoadReport,proto3" json:"enable_oob_load_report,omitempty"`
	// Load reporting interval to request from the server. Note that the
	// server may not provide reports as frequently as the client requests.
	// Used only when enable_oob_load_report is true. Default is 10 seconds.
	OobReportingPeriod *durationpb.Duration `protobuf:"bytes,2,opt,name=oob_reporting_period,json=oobReportingPeriod,proto3" json:"oob_reporting_period,omitempty"`
	// A given endpoint must report load metrics continuously for at least
	// this long before the endpoint weight will be used. This avoids
	// churn when the set of endpoint addresses changes. Takes effect
	// both immediately after we establish a connection to an endpoint and
	// after weight_expiration_period has caused us to stop using the most
	// recent load metrics. Default is 10 seconds.
	BlackoutPeriod *durationpb.Duration `protobuf:"bytes,3,opt,name=blackout_period,json=blackoutPeriod,proto3" json:"blackout_period,omitempty"`
	// If a given endpoint has not reported load metrics in this long,
	// then we stop using the reported weight. This ensures that we do
	// not continue to use very stale weights. Once we stop using a stale
	// value, if we later start seeing fresh reports again, the
	// blackout_period applies. Defaults to 3 minutes.
	WeightExpirationPeriod *durationpb.Duration `protobuf:"bytes,4,opt,name=weight_expiration_period,json=weightExpirationPeriod,proto3" json:"weight_expiration_period,omitempty"`
	// How often endpoint weights are recalculated. Values less than 100ms are
	// capped at 100ms. Default is 1 second.
	WeightUpdatePeriod *durationpb.Duration `protobuf:"bytes,5,opt,name=weight_update_period,json=weightUpdatePeriod,proto3" json:"weight_update_period,omitempty"`
	// The multiplier used to adjust endpoint weights with the error rate
	// calculated as eps/qps. Configuration is rejected if this value is negative.
	// Default is 1.0.
	ErrorUtilizationPenalty *wrapperspb.FloatValue `protobuf:"bytes,6,opt,name=error_utilization_penalty,json=errorUtilizationPenalty,proto3" json:"error_utilization_penalty,omitempty"`
	// By default, endpoint weight is computed based on the :ref:`application_utilization <envoy_v3_api_field_.xds.data.orca.v3.OrcaLoadReport.application_utilization>` field reported by the endpoint.
	// If that field is not set, then utilization will instead be computed by taking the max of the values of the metrics specified here.
	// For map fields in the ORCA proto, the string will be of the form “<map_field_name>.<map_key>“. For example, the string “named_metrics.foo“ will mean to look for the key “foo“ in the ORCA :ref:`named_metrics <envoy_v3_api_field_.xds.data.orca.v3.OrcaLoadReport.named_metrics>` field.
	// If none of the specified metrics are present in the load report, then :ref:`cpu_utilization <envoy_v3_api_field_.xds.data.orca.v3.OrcaLoadReport.cpu_utilization>` is used instead.
	MetricNamesForComputingUtilization []string `protobuf:"bytes,7,rep,name=metric_names_for_computing_utilization,json=metricNamesForComputingUtilization,proto3" json:"metric_names_for_computing_utilization,omitempty"`
}

func (x *ClientSideWeightedRoundRobin) Reset() {
	*x = ClientSideWeightedRoundRobin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_load_balancing_policies_client_side_weighted_round_robin_v3_client_side_weighted_round_robin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientSideWeightedRoundRobin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientSideWeightedRoundRobin) ProtoMessage() {}

func (x *ClientSideWeightedRoundRobin) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_load_balancing_policies_client_side_weighted_round_robin_v3_client_side_weighted_round_robin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientSideWeightedRoundRobin.ProtoReflect.Descriptor instead.
func (*ClientSideWeightedRoundRobin) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_load_balancing_policies_client_side_weighted_round_robin_v3_client_side_weighted_round_robin_proto_rawDescGZIP(), []int{0}
}

func (x *ClientSideWeightedRoundRobin) GetEnableOobLoadReport() *wrapperspb.BoolValue {
	if x != nil {
		return x.EnableOobLoadReport
	}
	return nil
}

func (x *ClientSideWeightedRoundRobin) GetOobReportingPeriod() *durationpb.Duration {
	if x != nil {
		return x.OobReportingPeriod
	}
	return nil
}

func (x *ClientSideWeightedRoundRobin) GetBlackoutPeriod() *durationpb.Duration {
	if x != nil {
		return x.BlackoutPeriod
	}
	return nil
}

func (x *ClientSideWeightedRoundRobin) GetWeightExpirationPeriod() *durationpb.Duration {
	if x != nil {
		return x.WeightExpirationPeriod
	}
	return nil
}

func (x *ClientSideWeightedRoundRobin) GetWeightUpdatePeriod() *durationpb.Duration {
	if x != nil {
		return x.WeightUpdatePeriod
	}
	return nil
}

func (x *ClientSideWeightedRoundRobin) GetErrorUtilizationPenalty() *wrapperspb.FloatValue {
	if x != nil {
		return x.ErrorUtilizationPenalty
	}
	return nil
}

func (x *ClientSideWeightedRoundRobin) GetMetricNamesForComputingUtilization() []string {
	if x != nil {
		return x.MetricNamesForComputingUtilization
	}
	return nil
}

var File_envoy_extensions_load_balancing_policies_client_side_weighted_round_robin_v3_client_side_weighted_round_robin_proto protoreflect.FileDescriptor

var file_envoy_extensions_load_balancing_policies_client_side_weighted_round_robin_v3_client_side_weighted_round_robin_proto_rawDesc = []byte{
	0x0a, 0x73, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e,
	0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x5f,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x2f, 0x76, 0x33, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x4c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x72, 0x6f, 0x62, 0x69, 0x6e,
	0x2e, 0x76, 0x33, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x04, 0x0a, 0x1c,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x64, 0x65, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x65, 0x64, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x6f, 0x62, 0x69, 0x6e, 0x12, 0x4f, 0x0a, 0x16,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6f, 0x6f, 0x62, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x5f,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x13, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x4f, 0x6f, 0x62, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x4b, 0x0a,
	0x14, 0x6f, 0x6f, 0x62, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x6f, 0x6f, 0x62, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x42, 0x0a, 0x0f, 0x62, 0x6c,
	0x61, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e,
	0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x53,
	0x0a, 0x18, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x16, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x12, 0x4b, 0x0a, 0x14, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x12, 0x63, 0x0a, 0x19, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x0a, 0xfa, 0x42, 0x07, 0x0a, 0x05, 0x2d, 0x00, 0x00, 0x00, 0x00, 0x52, 0x17, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65,
	0x6e, 0x61, 0x6c, 0x74, 0x79, 0x12, 0x52, 0x0a, 0x26, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x22, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x55, 0x74,
	0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0xa2, 0x02, 0xba, 0x80, 0xc8, 0xd1,
	0x06, 0x02, 0x10, 0x02, 0x0a, 0x5a, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65,
	0x64, 0x5f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x2e, 0x76, 0x33,
	0x42, 0x21, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x64, 0x65, 0x57, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x6f, 0x62, 0x69, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x96, 0x01, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f,
	0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x69, 0x64, 0x65, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x5f, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x2f, 0x76, 0x33, 0x3b, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64,
	0x5f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x76, 0x33, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_load_balancing_policies_client_side_weighted_round_robin_v3_client_side_weighted_round_robin_proto_rawDescOnce sync.Once
	file_envoy_extensions_load_balancing_policies_client_side_weighted_round_robin_v3_client_side_weighted_round_robin_proto_rawDescData = file_envoy_extensions_load_balancing_policies_client_side_weighted_round_robin_v3_client_side_weighted_round_robin_proto_rawDesc
)

func file_envoy_extensions_load_balancing_policies_client_side_weighted_round_robin_v3_client_side_weighted_round_robin_proto_rawDescGZIP() []byte {
	file_envoy_extensions_load_balancing_policies_client_side_weighted_round_robin_v3_client_side_weighted_round_robin_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_load_balancing_policies_client_side_weighted_round_robin_v3_client_side_weighted_round_robin_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_load_balancing_policies_client_side_weighted_round_robin_v3_client_side_weighted_round_robin_proto_rawDescData)
	})
	return file_envoy_extensions_load_balancing_policies_client_side_weighted_round_robin_v3_client_side_weighted_round_robin_proto_rawDescData
}

var file_envoy_extensions_load_balancing_policies_client_side_weighted_round_robin_v3_client_side_weighted_round_robin_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_load_balancing_policies_client_side_weighted_round_robin_v3_client_side_weighted_round_robin_proto_goTypes = []interface{}{
	(*ClientSideWeightedRoundRobin)(nil), // 0: envoy.extensions.load_balancing_policies.client_side_weighted_round_robin.v3.ClientSideWeightedRoundRobin
	(*wrapperspb.BoolValue)(nil),         // 1: google.protobuf.BoolValue
	(*durationpb.Duration)(nil),          // 2: google.protobuf.Duration
	(*wrapperspb.FloatValue)(nil),        // 3: google.protobuf.FloatValue
}
var file_envoy_extensions_load_balancing_policies_client_side_weighted_round_robin_v3_client_side_weighted_round_robin_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.load_balancing_policies.client_side_weighted_round_robin.v3.ClientSideWeightedRoundRobin.enable_oob_load_report:type_name -> google.protobuf.BoolValue
	2, // 1: envoy.extensions.load_balancing_policies.client_side_weighted_round_robin.v3.ClientSideWeightedRoundRobin.oob_reporting_period:type_name -> google.protobuf.Duration
	2, // 2: envoy.extensions.load_balancing_policies.client_side_weighted_round_robin.v3.ClientSideWeightedRoundRobin.blackout_period:type_name -> google.protobuf.Duration
	2, // 3: envoy.extensions.load_balancing_policies.client_side_weighted_round_robin.v3.ClientSideWeightedRoundRobin.weight_expiration_period:type_name -> google.protobuf.Duration
	2, // 4: envoy.extensions.load_balancing_policies.client_side_weighted_round_robin.v3.ClientSideWeightedRoundRobin.weight_update_period:type_name -> google.protobuf.Duration
	3, // 5: envoy.extensions.load_balancing_policies.client_side_weighted_round_robin.v3.ClientSideWeightedRoundRobin.error_utilization_penalty:type_name -> google.protobuf.FloatValue
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() {
	file_envoy_extensions_load_balancing_policies_client_side_weighted_round_robin_v3_client_side_weighted_round_robin_proto_init()
}
func file_envoy_extensions_load_balancing_policies_client_side_weighted_round_robin_v3_client_side_weighted_round_robin_proto_init() {
	if File_envoy_extensions_load_balancing_policies_client_side_weighted_round_robin_v3_client_side_weighted_round_robin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_load_balancing_policies_client_side_weighted_round_robin_v3_client_side_weighted_round_robin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientSideWeightedRoundRobin); i {
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
			RawDescriptor: file_envoy_extensions_load_balancing_policies_client_side_weighted_round_robin_v3_client_side_weighted_round_robin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_load_balancing_policies_client_side_weighted_round_robin_v3_client_side_weighted_round_robin_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_load_balancing_policies_client_side_weighted_round_robin_v3_client_side_weighted_round_robin_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_load_balancing_policies_client_side_weighted_round_robin_v3_client_side_weighted_round_robin_proto_msgTypes,
	}.Build()
	File_envoy_extensions_load_balancing_policies_client_side_weighted_round_robin_v3_client_side_weighted_round_robin_proto = out.File
	file_envoy_extensions_load_balancing_policies_client_side_weighted_round_robin_v3_client_side_weighted_round_robin_proto_rawDesc = nil
	file_envoy_extensions_load_balancing_policies_client_side_weighted_round_robin_v3_client_side_weighted_round_robin_proto_goTypes = nil
	file_envoy_extensions_load_balancing_policies_client_side_weighted_round_robin_v3_client_side_weighted_round_robin_proto_depIdxs = nil
}
