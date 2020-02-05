// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kafmesh/deviceId/details.proto

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

type Details struct {
	Time                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Details) Reset()         { *m = Details{} }
func (m *Details) String() string { return proto.CompactTextString(m) }
func (*Details) ProtoMessage()    {}
func (*Details) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a07319a7aba1c74, []int{0}
}

func (m *Details) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Details.Unmarshal(m, b)
}
func (m *Details) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Details.Marshal(b, m, deterministic)
}
func (m *Details) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Details.Merge(m, src)
}
func (m *Details) XXX_Size() int {
	return xxx_messageInfo_Details.Size(m)
}
func (m *Details) XXX_DiscardUnknown() {
	xxx_messageInfo_Details.DiscardUnknown(m)
}

var xxx_messageInfo_Details proto.InternalMessageInfo

func (m *Details) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *Details) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Details)(nil), "kafmesh.deviceId.Details")
}

func init() { proto.RegisterFile("kafmesh/deviceId/details.proto", fileDescriptor_0a07319a7aba1c74) }

var fileDescriptor_0a07319a7aba1c74 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcb, 0x4e, 0x4c, 0xcb,
	0x4d, 0x2d, 0xce, 0xd0, 0x4f, 0x49, 0x2d, 0xcb, 0x4c, 0x4e, 0xf5, 0x4c, 0xd1, 0x4f, 0x49, 0x2d,
	0x49, 0xcc, 0xcc, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x80, 0xca, 0xeb, 0xc1,
	0xe4, 0xa5, 0xe4, 0xd3, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0xf2, 0x49, 0xa5, 0x69, 0xfa,
	0x25, 0x99, 0xb9, 0xa9, 0xc5, 0x25, 0x89, 0xb9, 0x05, 0x10, 0x2d, 0x4a, 0xbe, 0x5c, 0xec, 0x2e,
	0x10, 0x33, 0x84, 0xf4, 0xb8, 0x58, 0x40, 0xb2, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x52,
	0x7a, 0x10, 0xad, 0x7a, 0x30, 0xad, 0x7a, 0x21, 0x30, 0xad, 0x41, 0x60, 0x75, 0x42, 0x42, 0x5c,
	0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x13, 0x57,
	0x14, 0x07, 0xcc, 0xee, 0x24, 0x36, 0xb0, 0x4e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd3,
	0xef, 0xb0, 0xaa, 0xb6, 0x00, 0x00, 0x00,
}
