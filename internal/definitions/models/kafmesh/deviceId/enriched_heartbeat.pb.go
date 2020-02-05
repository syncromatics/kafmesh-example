// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kafmesh/deviceId/enriched_heartbeat.proto

package deviceId

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EnrichedHeartbeat struct {
	Time                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	IsHealthy            bool                 `protobuf:"varint,2,opt,name=is_healthy,json=isHealthy,proto3" json:"is_healthy,omitempty"`
	CustomerId           int64                `protobuf:"varint,3,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	CustomerName         string               `protobuf:"bytes,4,opt,name=customer_name,json=customerName,proto3" json:"customer_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *EnrichedHeartbeat) Reset()         { *m = EnrichedHeartbeat{} }
func (m *EnrichedHeartbeat) String() string { return proto.CompactTextString(m) }
func (*EnrichedHeartbeat) ProtoMessage()    {}
func (*EnrichedHeartbeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a30101a0a5fdf64, []int{0}
}

func (m *EnrichedHeartbeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnrichedHeartbeat.Unmarshal(m, b)
}
func (m *EnrichedHeartbeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnrichedHeartbeat.Marshal(b, m, deterministic)
}
func (m *EnrichedHeartbeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnrichedHeartbeat.Merge(m, src)
}
func (m *EnrichedHeartbeat) XXX_Size() int {
	return xxx_messageInfo_EnrichedHeartbeat.Size(m)
}
func (m *EnrichedHeartbeat) XXX_DiscardUnknown() {
	xxx_messageInfo_EnrichedHeartbeat.DiscardUnknown(m)
}

var xxx_messageInfo_EnrichedHeartbeat proto.InternalMessageInfo

func (m *EnrichedHeartbeat) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *EnrichedHeartbeat) GetIsHealthy() bool {
	if m != nil {
		return m.IsHealthy
	}
	return false
}

func (m *EnrichedHeartbeat) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *EnrichedHeartbeat) GetCustomerName() string {
	if m != nil {
		return m.CustomerName
	}
	return ""
}

func init() {
	proto.RegisterType((*EnrichedHeartbeat)(nil), "kafmesh.deviceId.EnrichedHeartbeat")
}

func init() {
	proto.RegisterFile("kafmesh/deviceId/enriched_heartbeat.proto", fileDescriptor_0a30101a0a5fdf64)
}

var fileDescriptor_0a30101a0a5fdf64 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xbf, 0x4a, 0x43, 0x31,
	0x14, 0x87, 0x89, 0x2d, 0xd2, 0x9e, 0x2a, 0x68, 0xa6, 0x50, 0x90, 0x06, 0x5d, 0xe2, 0x92, 0x80,
	0xbe, 0x81, 0x20, 0xb4, 0x8b, 0x43, 0x70, 0x72, 0x29, 0xb9, 0x37, 0xa7, 0x4d, 0xb0, 0x69, 0x4a,
	0x92, 0x0a, 0xbe, 0x91, 0x8f, 0x29, 0xf7, 0x4f, 0xba, 0x7e, 0xf9, 0xf2, 0xf1, 0x3b, 0xf0, 0xfc,
	0x6d, 0x76, 0x01, 0xb3, 0x53, 0x16, 0x7f, 0x7c, 0x8b, 0x1b, 0xab, 0xf0, 0x98, 0x7c, 0xeb, 0xd0,
	0x6e, 0x1d, 0x9a, 0x54, 0x1a, 0x34, 0x45, 0x9e, 0x52, 0x2c, 0x91, 0xde, 0x8d, 0xaa, 0xac, 0xea,
	0x72, 0xb5, 0x8f, 0x71, 0x7f, 0x40, 0xd5, 0xbf, 0x37, 0xe7, 0x9d, 0x2a, 0x3e, 0x60, 0x2e, 0x26,
	0x9c, 0x86, 0x2f, 0x8f, 0x7f, 0x04, 0xee, 0xdf, 0xc7, 0xde, 0xba, 0xe6, 0xa8, 0x84, 0x69, 0x27,
	0x32, 0xc2, 0x89, 0x58, 0xbc, 0x2c, 0xe5, 0x50, 0x91, 0xb5, 0x22, 0x3f, 0x6b, 0x45, 0xf7, 0x1e,
	0x7d, 0x00, 0xf0, 0xb9, 0x9b, 0x73, 0x28, 0xee, 0x97, 0x5d, 0x71, 0x22, 0x66, 0x7a, 0xee, 0xf3,
	0x7a, 0x00, 0x74, 0x05, 0x8b, 0xf6, 0x9c, 0x4b, 0x0c, 0x98, 0xb6, 0xde, 0xb2, 0x09, 0x27, 0x62,
	0xa2, 0xa1, 0xa2, 0x8d, 0xa5, 0x4f, 0x70, 0x7b, 0x11, 0x8e, 0x26, 0x20, 0x9b, 0x72, 0x22, 0xe6,
	0xfa, 0xa6, 0xc2, 0x0f, 0x13, 0xf0, 0x0d, 0xbe, 0x66, 0xf5, 0xae, 0xe6, 0xba, 0x9f, 0xf2, 0xfa,
	0x1f, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xeb, 0x30, 0xe9, 0x1d, 0x01, 0x00, 0x00,
}