// Code generated by protoc-gen-go. DO NOT EDIT.
// source: crypto.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type MsgRequest struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgRequest) Reset()         { *m = MsgRequest{} }
func (m *MsgRequest) String() string { return proto.CompactTextString(m) }
func (*MsgRequest) ProtoMessage()    {}
func (*MsgRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_527278fb02d03321, []int{0}
}

func (m *MsgRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgRequest.Unmarshal(m, b)
}
func (m *MsgRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgRequest.Marshal(b, m, deterministic)
}
func (m *MsgRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequest.Merge(m, src)
}
func (m *MsgRequest) XXX_Size() int {
	return xxx_messageInfo_MsgRequest.Size(m)
}
func (m *MsgRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequest proto.InternalMessageInfo

func (m *MsgRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type MsgResponse struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgResponse) Reset()         { *m = MsgResponse{} }
func (m *MsgResponse) String() string { return proto.CompactTextString(m) }
func (*MsgResponse) ProtoMessage()    {}
func (*MsgResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_527278fb02d03321, []int{1}
}

func (m *MsgResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgResponse.Unmarshal(m, b)
}
func (m *MsgResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgResponse.Marshal(b, m, deterministic)
}
func (m *MsgResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgResponse.Merge(m, src)
}
func (m *MsgResponse) XXX_Size() int {
	return xxx_messageInfo_MsgResponse.Size(m)
}
func (m *MsgResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgResponse proto.InternalMessageInfo

func (m *MsgResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgRequest)(nil), "api.MsgRequest")
	proto.RegisterType((*MsgResponse)(nil), "api.MsgResponse")
}

func init() {
	proto.RegisterFile("crypto.proto", fileDescriptor_527278fb02d03321)
}

var fileDescriptor_527278fb02d03321 = []byte{
	// 127 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x2e, 0xaa, 0x2c,
	0x28, 0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0x52, 0xe0,
	0xe2, 0xf2, 0x2d, 0x4e, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0x29,
	0x49, 0xad, 0x28, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0x14, 0xb9, 0xb8,
	0xc1, 0x2a, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0xb1, 0x29, 0x31, 0x32, 0xe3, 0x62, 0x73, 0x06,
	0x9b, 0x2c, 0xa4, 0xc3, 0xc5, 0x1e, 0x9c, 0x9a, 0x97, 0xe2, 0x5b, 0x9c, 0x2e, 0xc4, 0xaf, 0x97,
	0x58, 0x90, 0xa9, 0x87, 0x30, 0x5c, 0x4a, 0x00, 0x21, 0x00, 0x31, 0x2b, 0x89, 0x0d, 0xec, 0x10,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb9, 0xe8, 0x75, 0xa0, 0x98, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CryptoClient is the client API for Crypto service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CryptoClient interface {
	SendMsg(ctx context.Context, in *MsgRequest, opts ...grpc.CallOption) (*MsgResponse, error)
}

type cryptoClient struct {
	cc grpc.ClientConnInterface
}

func NewCryptoClient(cc grpc.ClientConnInterface) CryptoClient {
	return &cryptoClient{cc}
}

func (c *cryptoClient) SendMsg(ctx context.Context, in *MsgRequest, opts ...grpc.CallOption) (*MsgResponse, error) {
	out := new(MsgResponse)
	err := c.cc.Invoke(ctx, "/api.Crypto/SendMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CryptoServer is the server API for Crypto service.
type CryptoServer interface {
	SendMsg(context.Context, *MsgRequest) (*MsgResponse, error)
}

// UnimplementedCryptoServer can be embedded to have forward compatible implementations.
type UnimplementedCryptoServer struct {
}

func (*UnimplementedCryptoServer) SendMsg(ctx context.Context, req *MsgRequest) (*MsgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMsg not implemented")
}

func RegisterCryptoServer(s *grpc.Server, srv CryptoServer) {
	s.RegisterService(&_Crypto_serviceDesc, srv)
}

func _Crypto_SendMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServer).SendMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Crypto/SendMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServer).SendMsg(ctx, req.(*MsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Crypto_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Crypto",
	HandlerType: (*CryptoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMsg",
			Handler:    _Crypto_SendMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crypto.proto",
}
