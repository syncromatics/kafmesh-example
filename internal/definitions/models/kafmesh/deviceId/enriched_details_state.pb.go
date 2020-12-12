// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: kafmesh/deviceId/enriched_details_state.proto

package deviceId

import (
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type EnrichedDetailsState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Details    *Details             `protobuf:"bytes,1,opt,name=details,proto3" json:"details,omitempty"`
	CustomerId *wrappers.Int64Value `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *EnrichedDetailsState) Reset() {
	*x = EnrichedDetailsState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kafmesh_deviceId_enriched_details_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrichedDetailsState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrichedDetailsState) ProtoMessage() {}

func (x *EnrichedDetailsState) ProtoReflect() protoreflect.Message {
	mi := &file_kafmesh_deviceId_enriched_details_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrichedDetailsState.ProtoReflect.Descriptor instead.
func (*EnrichedDetailsState) Descriptor() ([]byte, []int) {
	return file_kafmesh_deviceId_enriched_details_state_proto_rawDescGZIP(), []int{0}
}

func (x *EnrichedDetailsState) GetDetails() *Details {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *EnrichedDetailsState) GetCustomerId() *wrappers.Int64Value {
	if x != nil {
		return x.CustomerId
	}
	return nil
}

var File_kafmesh_deviceId_enriched_details_state_proto protoreflect.FileDescriptor

var file_kafmesh_deviceId_enriched_details_state_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x6b, 0x61, 0x66, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x2f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x6b, 0x61, 0x66, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x64, 0x1a, 0x1e, 0x6b, 0x61, 0x66, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x2f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x89, 0x01, 0x0a, 0x14, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x65, 0x64, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6b, 0x61,
	0x66, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x2e, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x3c, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x42, 0x0a, 0x5a,
	0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_kafmesh_deviceId_enriched_details_state_proto_rawDescOnce sync.Once
	file_kafmesh_deviceId_enriched_details_state_proto_rawDescData = file_kafmesh_deviceId_enriched_details_state_proto_rawDesc
)

func file_kafmesh_deviceId_enriched_details_state_proto_rawDescGZIP() []byte {
	file_kafmesh_deviceId_enriched_details_state_proto_rawDescOnce.Do(func() {
		file_kafmesh_deviceId_enriched_details_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_kafmesh_deviceId_enriched_details_state_proto_rawDescData)
	})
	return file_kafmesh_deviceId_enriched_details_state_proto_rawDescData
}

var file_kafmesh_deviceId_enriched_details_state_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_kafmesh_deviceId_enriched_details_state_proto_goTypes = []interface{}{
	(*EnrichedDetailsState)(nil), // 0: kafmesh.deviceId.EnrichedDetailsState
	(*Details)(nil),              // 1: kafmesh.deviceId.Details
	(*wrappers.Int64Value)(nil),  // 2: google.protobuf.Int64Value
}
var file_kafmesh_deviceId_enriched_details_state_proto_depIdxs = []int32{
	1, // 0: kafmesh.deviceId.EnrichedDetailsState.details:type_name -> kafmesh.deviceId.Details
	2, // 1: kafmesh.deviceId.EnrichedDetailsState.customer_id:type_name -> google.protobuf.Int64Value
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_kafmesh_deviceId_enriched_details_state_proto_init() }
func file_kafmesh_deviceId_enriched_details_state_proto_init() {
	if File_kafmesh_deviceId_enriched_details_state_proto != nil {
		return
	}
	file_kafmesh_deviceId_details_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kafmesh_deviceId_enriched_details_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrichedDetailsState); i {
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
			RawDescriptor: file_kafmesh_deviceId_enriched_details_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kafmesh_deviceId_enriched_details_state_proto_goTypes,
		DependencyIndexes: file_kafmesh_deviceId_enriched_details_state_proto_depIdxs,
		MessageInfos:      file_kafmesh_deviceId_enriched_details_state_proto_msgTypes,
	}.Build()
	File_kafmesh_deviceId_enriched_details_state_proto = out.File
	file_kafmesh_deviceId_enriched_details_state_proto_rawDesc = nil
	file_kafmesh_deviceId_enriched_details_state_proto_goTypes = nil
	file_kafmesh_deviceId_enriched_details_state_proto_depIdxs = nil
}
