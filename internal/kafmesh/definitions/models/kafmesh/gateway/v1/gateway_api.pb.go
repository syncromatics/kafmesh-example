// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kafmesh/gateway/v1/gateway_api.proto

package gatewayv1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type DetailsRequest struct {
	DeviceId             int64    `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetailsRequest) Reset()         { *m = DetailsRequest{} }
func (m *DetailsRequest) String() string { return proto.CompactTextString(m) }
func (*DetailsRequest) ProtoMessage()    {}
func (*DetailsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f356f584d20963a, []int{0}
}

func (m *DetailsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailsRequest.Unmarshal(m, b)
}
func (m *DetailsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailsRequest.Marshal(b, m, deterministic)
}
func (m *DetailsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailsRequest.Merge(m, src)
}
func (m *DetailsRequest) XXX_Size() int {
	return xxx_messageInfo_DetailsRequest.Size(m)
}
func (m *DetailsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DetailsRequest proto.InternalMessageInfo

func (m *DetailsRequest) GetDeviceId() int64 {
	if m != nil {
		return m.DeviceId
	}
	return 0
}

func (m *DetailsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DetailsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetailsResponse) Reset()         { *m = DetailsResponse{} }
func (m *DetailsResponse) String() string { return proto.CompactTextString(m) }
func (*DetailsResponse) ProtoMessage()    {}
func (*DetailsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f356f584d20963a, []int{1}
}

func (m *DetailsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailsResponse.Unmarshal(m, b)
}
func (m *DetailsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailsResponse.Marshal(b, m, deterministic)
}
func (m *DetailsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailsResponse.Merge(m, src)
}
func (m *DetailsResponse) XXX_Size() int {
	return xxx_messageInfo_DetailsResponse.Size(m)
}
func (m *DetailsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DetailsResponse proto.InternalMessageInfo

type HeartbeatRequest struct {
	DeviceId             int64                `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Time                 *timestamp.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	IsHealthy            bool                 `protobuf:"varint,3,opt,name=is_healthy,json=isHealthy,proto3" json:"is_healthy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *HeartbeatRequest) Reset()         { *m = HeartbeatRequest{} }
func (m *HeartbeatRequest) String() string { return proto.CompactTextString(m) }
func (*HeartbeatRequest) ProtoMessage()    {}
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f356f584d20963a, []int{2}
}

func (m *HeartbeatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatRequest.Unmarshal(m, b)
}
func (m *HeartbeatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatRequest.Marshal(b, m, deterministic)
}
func (m *HeartbeatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatRequest.Merge(m, src)
}
func (m *HeartbeatRequest) XXX_Size() int {
	return xxx_messageInfo_HeartbeatRequest.Size(m)
}
func (m *HeartbeatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatRequest proto.InternalMessageInfo

func (m *HeartbeatRequest) GetDeviceId() int64 {
	if m != nil {
		return m.DeviceId
	}
	return 0
}

func (m *HeartbeatRequest) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *HeartbeatRequest) GetIsHealthy() bool {
	if m != nil {
		return m.IsHealthy
	}
	return false
}

type HeartbeatResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatResponse) Reset()         { *m = HeartbeatResponse{} }
func (m *HeartbeatResponse) String() string { return proto.CompactTextString(m) }
func (*HeartbeatResponse) ProtoMessage()    {}
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f356f584d20963a, []int{3}
}

func (m *HeartbeatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatResponse.Unmarshal(m, b)
}
func (m *HeartbeatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatResponse.Marshal(b, m, deterministic)
}
func (m *HeartbeatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatResponse.Merge(m, src)
}
func (m *HeartbeatResponse) XXX_Size() int {
	return xxx_messageInfo_HeartbeatResponse.Size(m)
}
func (m *HeartbeatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DetailsRequest)(nil), "kafmesh.gateway.v1.DetailsRequest")
	proto.RegisterType((*DetailsResponse)(nil), "kafmesh.gateway.v1.DetailsResponse")
	proto.RegisterType((*HeartbeatRequest)(nil), "kafmesh.gateway.v1.HeartbeatRequest")
	proto.RegisterType((*HeartbeatResponse)(nil), "kafmesh.gateway.v1.HeartbeatResponse")
}

func init() {
	proto.RegisterFile("kafmesh/gateway/v1/gateway_api.proto", fileDescriptor_0f356f584d20963a)
}

var fileDescriptor_0f356f584d20963a = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x49, 0x5b, 0xbe, 0xaf, 0xb9, 0x82, 0xda, 0x71, 0x53, 0x22, 0x62, 0x88, 0x15, 0xb2,
	0x9a, 0x90, 0xfa, 0x04, 0x15, 0xc1, 0x16, 0x37, 0x32, 0x88, 0x88, 0x9b, 0x30, 0x31, 0xb7, 0xc9,
	0x60, 0xd2, 0xc4, 0xce, 0x34, 0xd2, 0x8d, 0x0f, 0xe4, 0x23, 0xf8, 0x74, 0x62, 0x66, 0x52, 0xfc,
	0x87, 0xba, 0x1b, 0xee, 0x3d, 0xf7, 0x9c, 0xdf, 0x1c, 0x18, 0xdd, 0xf3, 0x79, 0x81, 0x32, 0x0b,
	0x52, 0xae, 0xf0, 0x91, 0xaf, 0x83, 0x3a, 0x6c, 0x9f, 0x11, 0xaf, 0x04, 0xad, 0x96, 0xa5, 0x2a,
	0x09, 0x31, 0x2a, 0x6a, 0x56, 0xb4, 0x0e, 0x9d, 0xc3, 0xb4, 0x2c, 0xd3, 0x1c, 0x83, 0x46, 0x11,
	0xaf, 0xe6, 0x81, 0x12, 0x05, 0x4a, 0xc5, 0x8b, 0x4a, 0x1f, 0x79, 0x13, 0xd8, 0x3e, 0x43, 0xc5,
	0x45, 0x2e, 0x19, 0x3e, 0xac, 0x50, 0x2a, 0xb2, 0x0f, 0x76, 0x82, 0xb5, 0xb8, 0xc3, 0x48, 0x24,
	0x43, 0xcb, 0xb5, 0xfc, 0x2e, 0xeb, 0xeb, 0xc1, 0x2c, 0x21, 0x04, 0x7a, 0x0b, 0x5e, 0xe0, 0xb0,
	0xe3, 0x5a, 0xbe, 0xcd, 0x9a, 0xb7, 0x37, 0x80, 0x9d, 0x8d, 0x85, 0xac, 0xca, 0x85, 0x44, 0xef,
	0x09, 0x76, 0xa7, 0xc8, 0x97, 0x2a, 0x46, 0xae, 0xfe, 0xe4, 0x4b, 0xa1, 0xf7, 0x46, 0xd6, 0xf8,
	0x6e, 0x8d, 0x1d, 0xaa, 0xb1, 0x69, 0x8b, 0x4d, 0xaf, 0x5a, 0x6c, 0xd6, 0xe8, 0xc8, 0x01, 0x80,
	0x90, 0x51, 0x86, 0x3c, 0x57, 0xd9, 0x7a, 0xd8, 0x75, 0x2d, 0xbf, 0xcf, 0x6c, 0x21, 0xa7, 0x7a,
	0xe0, 0xed, 0xc1, 0xe0, 0x5d, 0xbe, 0x86, 0x1a, 0xbf, 0x58, 0x00, 0xe7, 0xba, 0x9a, 0xc9, 0xe5,
	0x8c, 0x30, 0xf8, 0x6f, 0xb0, 0x89, 0x47, 0xbf, 0x56, 0x47, 0x3f, 0xd6, 0xe2, 0x1c, 0xfd, 0xa8,
	0xd1, 0x11, 0xe4, 0x06, 0xec, 0x4d, 0x2e, 0x19, 0x7d, 0x77, 0xf1, 0xb9, 0x16, 0xe7, 0xf8, 0x17,
	0x95, 0x76, 0x3e, 0x75, 0x6f, 0x6d, 0xb3, 0xaf, 0xc3, 0xe7, 0x0e, 0xb9, 0x30, 0x37, 0xe6, 0x3f,
	0xf4, 0x3a, 0x8c, 0xff, 0x35, 0x65, 0x9d, 0xbc, 0x06, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xce, 0x29,
	0x49, 0x2d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GatewayAPIClient is the client API for GatewayAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GatewayAPIClient interface {
	Details(ctx context.Context, in *DetailsRequest, opts ...grpc.CallOption) (*DetailsResponse, error)
	Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error)
}

type gatewayAPIClient struct {
	cc *grpc.ClientConn
}

func NewGatewayAPIClient(cc *grpc.ClientConn) GatewayAPIClient {
	return &gatewayAPIClient{cc}
}

func (c *gatewayAPIClient) Details(ctx context.Context, in *DetailsRequest, opts ...grpc.CallOption) (*DetailsResponse, error) {
	out := new(DetailsResponse)
	err := c.cc.Invoke(ctx, "/kafmesh.gateway.v1.GatewayAPI/Details", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayAPIClient) Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error) {
	out := new(HeartbeatResponse)
	err := c.cc.Invoke(ctx, "/kafmesh.gateway.v1.GatewayAPI/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayAPIServer is the server API for GatewayAPI service.
type GatewayAPIServer interface {
	Details(context.Context, *DetailsRequest) (*DetailsResponse, error)
	Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error)
}

// UnimplementedGatewayAPIServer can be embedded to have forward compatible implementations.
type UnimplementedGatewayAPIServer struct {
}

func (*UnimplementedGatewayAPIServer) Details(ctx context.Context, req *DetailsRequest) (*DetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Details not implemented")
}
func (*UnimplementedGatewayAPIServer) Heartbeat(ctx context.Context, req *HeartbeatRequest) (*HeartbeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}

func RegisterGatewayAPIServer(s *grpc.Server, srv GatewayAPIServer) {
	s.RegisterService(&_GatewayAPI_serviceDesc, srv)
}

func _GatewayAPI_Details_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayAPIServer).Details(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kafmesh.gateway.v1.GatewayAPI/Details",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayAPIServer).Details(ctx, req.(*DetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayAPI_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayAPIServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kafmesh.gateway.v1.GatewayAPI/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayAPIServer).Heartbeat(ctx, req.(*HeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GatewayAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kafmesh.gateway.v1.GatewayAPI",
	HandlerType: (*GatewayAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Details",
			Handler:    _GatewayAPI_Details_Handler,
		},
		{
			MethodName: "Heartbeat",
			Handler:    _GatewayAPI_Heartbeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kafmesh/gateway/v1/gateway_api.proto",
}
