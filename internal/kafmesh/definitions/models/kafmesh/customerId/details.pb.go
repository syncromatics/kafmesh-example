// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kafmesh/customerId/details.proto

package customerId

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Details) Reset()         { *m = Details{} }
func (m *Details) String() string { return proto.CompactTextString(m) }
func (*Details) ProtoMessage()    {}
func (*Details) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1a96df94b11d63a, []int{0}
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

func (m *Details) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Details)(nil), "kafmesh.customerId.Details")
}

func init() { proto.RegisterFile("kafmesh/customerId/details.proto", fileDescriptor_b1a96df94b11d63a) }

var fileDescriptor_b1a96df94b11d63a = []byte{
	// 100 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc8, 0x4e, 0x4c, 0xcb,
	0x4d, 0x2d, 0xce, 0xd0, 0x4f, 0x2e, 0x2d, 0x2e, 0xc9, 0xcf, 0x4d, 0x2d, 0xf2, 0x4c, 0xd1, 0x4f,
	0x49, 0x2d, 0x49, 0xcc, 0xcc, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x82, 0xaa,
	0xd0, 0x43, 0xa8, 0x50, 0x92, 0xe5, 0x62, 0x77, 0x81, 0x28, 0x12, 0x12, 0xe2, 0x62, 0xc9, 0x4b,
	0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x9d, 0x78, 0xa2, 0xb8, 0x10,
	0x8a, 0x93, 0xd8, 0xc0, 0xe6, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x8f, 0xe7, 0xf1,
	0x6b, 0x00, 0x00, 0x00,
}
