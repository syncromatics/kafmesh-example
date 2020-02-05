// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kafmesh/deviceId/heartbeat.proto

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

type Heartbeat struct {
	Time                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	IsHealthy            bool                 `protobuf:"varint,2,opt,name=is_healthy,json=isHealthy,proto3" json:"is_healthy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Heartbeat) Reset()         { *m = Heartbeat{} }
func (m *Heartbeat) String() string { return proto.CompactTextString(m) }
func (*Heartbeat) ProtoMessage()    {}
func (*Heartbeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a3eef1db8bb1444, []int{0}
}

func (m *Heartbeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Heartbeat.Unmarshal(m, b)
}
func (m *Heartbeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Heartbeat.Marshal(b, m, deterministic)
}
func (m *Heartbeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Heartbeat.Merge(m, src)
}
func (m *Heartbeat) XXX_Size() int {
	return xxx_messageInfo_Heartbeat.Size(m)
}
func (m *Heartbeat) XXX_DiscardUnknown() {
	xxx_messageInfo_Heartbeat.DiscardUnknown(m)
}

var xxx_messageInfo_Heartbeat proto.InternalMessageInfo

func (m *Heartbeat) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *Heartbeat) GetIsHealthy() bool {
	if m != nil {
		return m.IsHealthy
	}
	return false
}

func init() {
	proto.RegisterType((*Heartbeat)(nil), "kafmesh.deviceId.Heartbeat")
}

func init() { proto.RegisterFile("kafmesh/deviceId/heartbeat.proto", fileDescriptor_3a3eef1db8bb1444) }

var fileDescriptor_3a3eef1db8bb1444 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc8, 0x4e, 0x4c, 0xcb,
	0x4d, 0x2d, 0xce, 0xd0, 0x4f, 0x49, 0x2d, 0xcb, 0x4c, 0x4e, 0xf5, 0x4c, 0xd1, 0xcf, 0x48, 0x4d,
	0x2c, 0x2a, 0x49, 0x4a, 0x4d, 0x2c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x80, 0xaa,
	0xd0, 0x83, 0xa9, 0x90, 0x92, 0x4f, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0xcb, 0x27, 0x95,
	0xa6, 0xe9, 0x97, 0x64, 0xe6, 0xa6, 0x16, 0x97, 0x24, 0xe6, 0x16, 0x40, 0xb4, 0x28, 0x45, 0x71,
	0x71, 0x7a, 0xc0, 0x4c, 0x11, 0xd2, 0xe3, 0x62, 0x01, 0xc9, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x1b, 0x49, 0xe9, 0x41, 0x34, 0xeb, 0xc1, 0x34, 0xeb, 0x85, 0xc0, 0x34, 0x07, 0x81, 0xd5, 0x09,
	0xc9, 0x72, 0x71, 0x65, 0x16, 0xc7, 0x67, 0xa4, 0x26, 0xe6, 0x94, 0x64, 0x54, 0x4a, 0x30, 0x29,
	0x30, 0x6a, 0x70, 0x04, 0x71, 0x66, 0x16, 0x7b, 0x40, 0x04, 0x9c, 0xb8, 0xa2, 0x38, 0x60, 0x0e,
	0x49, 0x62, 0x03, 0x1b, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x70, 0x66, 0x77, 0x82, 0xc5,
	0x00, 0x00, 0x00,
}
