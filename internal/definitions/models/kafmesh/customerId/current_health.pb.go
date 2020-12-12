// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: kafmesh/customerId/current_health.proto

package customerId

import (
	proto "github.com/golang/protobuf/proto"
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

type CurrentHealth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalDevices   int32 `protobuf:"varint,1,opt,name=total_devices,json=totalDevices,proto3" json:"total_devices,omitempty"`
	HealthyDevices int32 `protobuf:"varint,2,opt,name=healthy_devices,json=healthyDevices,proto3" json:"healthy_devices,omitempty"`
}

func (x *CurrentHealth) Reset() {
	*x = CurrentHealth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kafmesh_customerId_current_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentHealth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentHealth) ProtoMessage() {}

func (x *CurrentHealth) ProtoReflect() protoreflect.Message {
	mi := &file_kafmesh_customerId_current_health_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentHealth.ProtoReflect.Descriptor instead.
func (*CurrentHealth) Descriptor() ([]byte, []int) {
	return file_kafmesh_customerId_current_health_proto_rawDescGZIP(), []int{0}
}

func (x *CurrentHealth) GetTotalDevices() int32 {
	if x != nil {
		return x.TotalDevices
	}
	return 0
}

func (x *CurrentHealth) GetHealthyDevices() int32 {
	if x != nil {
		return x.HealthyDevices
	}
	return 0
}

var File_kafmesh_customerId_current_health_proto protoreflect.FileDescriptor

var file_kafmesh_customerId_current_health_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6b, 0x61, 0x66, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6b, 0x61, 0x66, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5d, 0x0a,
	0x0d, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x23,
	0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x5f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x79, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x0c, 0x5a, 0x0a,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_kafmesh_customerId_current_health_proto_rawDescOnce sync.Once
	file_kafmesh_customerId_current_health_proto_rawDescData = file_kafmesh_customerId_current_health_proto_rawDesc
)

func file_kafmesh_customerId_current_health_proto_rawDescGZIP() []byte {
	file_kafmesh_customerId_current_health_proto_rawDescOnce.Do(func() {
		file_kafmesh_customerId_current_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_kafmesh_customerId_current_health_proto_rawDescData)
	})
	return file_kafmesh_customerId_current_health_proto_rawDescData
}

var file_kafmesh_customerId_current_health_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_kafmesh_customerId_current_health_proto_goTypes = []interface{}{
	(*CurrentHealth)(nil), // 0: kafmesh.customerId.CurrentHealth
}
var file_kafmesh_customerId_current_health_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kafmesh_customerId_current_health_proto_init() }
func file_kafmesh_customerId_current_health_proto_init() {
	if File_kafmesh_customerId_current_health_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kafmesh_customerId_current_health_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrentHealth); i {
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
			RawDescriptor: file_kafmesh_customerId_current_health_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kafmesh_customerId_current_health_proto_goTypes,
		DependencyIndexes: file_kafmesh_customerId_current_health_proto_depIdxs,
		MessageInfos:      file_kafmesh_customerId_current_health_proto_msgTypes,
	}.Build()
	File_kafmesh_customerId_current_health_proto = out.File
	file_kafmesh_customerId_current_health_proto_rawDesc = nil
	file_kafmesh_customerId_current_health_proto_goTypes = nil
	file_kafmesh_customerId_current_health_proto_depIdxs = nil
}
