// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kafmesh/deviceId/customer.proto

package deviceId

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

type Customer struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Customer) Reset()         { *m = Customer{} }
func (m *Customer) String() string { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()    {}
func (*Customer) Descriptor() ([]byte, []int) {
	return fileDescriptor_2270b68511d4000b, []int{0}
}

func (m *Customer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Customer.Unmarshal(m, b)
}
func (m *Customer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Customer.Marshal(b, m, deterministic)
}
func (m *Customer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Customer.Merge(m, src)
}
func (m *Customer) XXX_Size() int {
	return xxx_messageInfo_Customer.Size(m)
}
func (m *Customer) XXX_DiscardUnknown() {
	xxx_messageInfo_Customer.DiscardUnknown(m)
}

var xxx_messageInfo_Customer proto.InternalMessageInfo

func (m *Customer) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Customer)(nil), "kafmesh.deviceId.Customer")
}

func init() { proto.RegisterFile("kafmesh/deviceId/customer.proto", fileDescriptor_2270b68511d4000b) }

var fileDescriptor_2270b68511d4000b = []byte{
	// 98 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcf, 0x4e, 0x4c, 0xcb,
	0x4d, 0x2d, 0xce, 0xd0, 0x4f, 0x49, 0x2d, 0xcb, 0x4c, 0x4e, 0xf5, 0x4c, 0xd1, 0x4f, 0x2e, 0x2d,
	0x2e, 0xc9, 0xcf, 0x4d, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x80, 0x2a, 0xd0,
	0x83, 0x29, 0x50, 0x92, 0xe2, 0xe2, 0x70, 0x86, 0xaa, 0x11, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91,
	0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x62, 0xca, 0x4c, 0x71, 0xe2, 0x8a, 0xe2, 0x80, 0xa9, 0x4b,
	0x62, 0x03, 0x1b, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xed, 0x2a, 0x0b, 0x59, 0x63, 0x00,
	0x00, 0x00,
}
