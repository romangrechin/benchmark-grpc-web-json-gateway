// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ping.proto

package pkg

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type PingRequest struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d51d96c3ad891f5, []int{0}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type PingResponse struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d51d96c3ad891f5, []int{1}
}

func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func (m *PingResponse) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *PingResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "sms.core.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "sms.core.PingResponse")
}

func init() { proto.RegisterFile("ping.proto", fileDescriptor_6d51d96c3ad891f5) }

var fileDescriptor_6d51d96c3ad891f5 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4b, 0x03, 0x41,
	0x10, 0x85, 0xb9, 0x28, 0xc6, 0x6c, 0x6c, 0x5c, 0x51, 0xce, 0x90, 0x22, 0x5c, 0x15, 0x10, 0xf6,
	0x40, 0x5b, 0x1b, 0x23, 0x68, 0x7b, 0x9c, 0x9d, 0x8d, 0xac, 0x9b, 0x61, 0xb3, 0x90, 0xdd, 0x19,
	0x77, 0x56, 0x1b, 0xb1, 0xf1, 0x2f, 0xf8, 0xd3, 0xfc, 0x0b, 0xfe, 0x10, 0xb9, 0x2c, 0xa7, 0x62,
	0x91, 0x72, 0x1e, 0x8f, 0xf7, 0xbe, 0x79, 0x42, 0x90, 0x0b, 0x56, 0x51, 0xc4, 0x84, 0x72, 0x9f,
	0x3d, 0x2b, 0x83, 0x11, 0x26, 0x53, 0x8b, 0x68, 0xd7, 0x50, 0x6b, 0x72, 0xb5, 0x0e, 0x01, 0x93,
	0x4e, 0x0e, 0x03, 0x67, 0x5f, 0x75, 0x26, 0xc6, 0x8d, 0x0b, 0xb6, 0x85, 0xa7, 0x67, 0xe0, 0x24,
	0xa7, 0x62, 0x94, 0x9c, 0x07, 0x4e, 0xda, 0x53, 0x59, 0xcc, 0x8a, 0xf9, 0x4e, 0xfb, 0x2b, 0x54,
	0x37, 0xe2, 0x20, 0x9b, 0x99, 0x30, 0x30, 0x6c, 0x77, 0xcb, 0x52, 0x0c, 0x3d, 0x30, 0x6b, 0x0b,
	0xe5, 0x60, 0x56, 0xcc, 0x47, 0x6d, 0x7f, 0x9e, 0x3f, 0x88, 0xf1, 0x35, 0x46, 0xb8, 0x83, 0xf8,
	0xe2, 0x0c, 0xc8, 0x46, 0xec, 0x76, 0xb1, 0xf2, 0x58, 0xf5, 0xd0, 0xea, 0x0f, 0xd3, 0xe4, 0xe4,
	0xbf, 0x9c, 0xdb, 0xab, 0xd3, 0xf7, 0xcf, 0xaf, 0x8f, 0xc1, 0x91, 0x3c, 0xac, 0xbb, 0xbf, 0xeb,
	0xd7, 0x9f, 0xe6, 0xb7, 0xc5, 0xa5, 0xa8, 0x0c, 0x7a, 0x95, 0x56, 0x2e, 0xac, 0x96, 0x5a, 0x31,
	0xc5, 0x6e, 0x19, 0x1b, 0xc9, 0xe4, 0x1c, 0x8f, 0x4b, 0x58, 0x2f, 0x36, 0x10, 0xb7, 0x91, 0xcc,
	0x15, 0xb9, 0xa6, 0xb8, 0x1f, 0x72, 0xe6, 0x79, 0xdc, 0xdb, 0x4c, 0x73, 0xf1, 0x1d, 0x00, 0x00,
	0xff, 0xff, 0x0a, 0x37, 0xeb, 0x16, 0x50, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CoreServiceClient is the client API for CoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CoreServiceClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type coreServiceClient struct {
	cc *grpc.ClientConn
}

func NewCoreServiceClient(cc *grpc.ClientConn) CoreServiceClient {
	return &coreServiceClient{cc}
}

func (c *coreServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/sms.core.CoreService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoreServiceServer is the server API for CoreService service.
type CoreServiceServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
}

// UnimplementedCoreServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCoreServiceServer struct {
}

func (*UnimplementedCoreServiceServer) Ping(ctx context.Context, req *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterCoreServiceServer(s *grpc.Server, srv CoreServiceServer) {
	s.RegisterService(&_CoreService_serviceDesc, srv)
}

func _CoreService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sms.core.CoreService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sms.core.CoreService",
	HandlerType: (*CoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _CoreService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ping.proto",
}
