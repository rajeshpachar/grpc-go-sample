// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greeter.proto

package greeter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HelloRequest struct {
	Name                 string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Details              []*any.Any `protobuf:"bytes,2,rep,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_greeter_64d10d3f4bf98284, []int{0}
}
func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (dst *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(dst, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HelloRequest) GetDetails() []*any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

type Person struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_greeter_64d10d3f4bf98284, []int{1}
}
func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (dst *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(dst, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Person) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type HelloReply struct {
	Response             string   `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Details              *any.Any `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_greeter_64d10d3f4bf98284, []int{2}
}
func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (dst *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(dst, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func (m *HelloReply) GetDetails() *any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "greeter.HelloRequest")
	proto.RegisterType((*Person)(nil), "greeter.Person")
	proto.RegisterType((*HelloReply)(nil), "greeter.HelloReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/greeter.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greeter.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greeter.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greeter.proto",
}

func init() { proto.RegisterFile("greeter.proto", fileDescriptor_greeter_64d10d3f4bf98284) }

var fileDescriptor_greeter_64d10d3f4bf98284 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xb1, 0x4b, 0xc4, 0x30,
	0x18, 0xc5, 0xed, 0x29, 0xd7, 0xf6, 0x53, 0x97, 0x4f, 0x85, 0xb3, 0x22, 0x1c, 0x99, 0x6e, 0xca,
	0xc1, 0xb9, 0xb8, 0x8a, 0x82, 0x4e, 0x22, 0x71, 0x71, 0x93, 0x1c, 0x7e, 0x57, 0x0e, 0x62, 0x12,
	0x93, 0x74, 0xc8, 0x7f, 0x2f, 0x4d, 0x1b, 0xa9, 0x2e, 0x6e, 0xf9, 0xde, 0x7b, 0xfc, 0xf2, 0x78,
	0x70, 0xda, 0x3a, 0xa2, 0x40, 0x8e, 0x5b, 0x67, 0x82, 0xc1, 0x72, 0x3c, 0x9b, 0xcb, 0xd6, 0x98,
	0x56, 0xd1, 0x3a, 0xc9, 0xdb, 0x6e, 0xb7, 0x96, 0x3a, 0x0e, 0x19, 0x26, 0xe0, 0xe4, 0x89, 0x94,
	0x32, 0x82, 0xbe, 0x3a, 0xf2, 0x01, 0x11, 0x8e, 0xb4, 0xfc, 0xa4, 0x45, 0xb1, 0x2c, 0x56, 0xb5,
	0x48, 0x6f, 0xe4, 0x50, 0x7e, 0x50, 0x90, 0x7b, 0xe5, 0x17, 0xb3, 0xe5, 0xe1, 0xea, 0x78, 0x73,
	0xce, 0x07, 0x20, 0xcf, 0x40, 0x7e, 0xa7, 0xa3, 0xc8, 0x21, 0xf6, 0x00, 0xf3, 0x17, 0x72, 0xde,
	0x68, 0xbc, 0x06, 0xd8, 0xed, 0x9d, 0x0f, 0xef, 0x13, 0x66, 0x9d, 0x94, 0xe7, 0x1e, 0x7c, 0x05,
	0xb5, 0x92, 0xd9, 0x9d, 0x25, 0xb7, 0xea, 0x85, 0xde, 0x64, 0x6f, 0x00, 0x63, 0x33, 0xab, 0x22,
	0x36, 0x50, 0x39, 0xf2, 0xd6, 0x68, 0x9f, 0x39, 0x3f, 0xf7, 0xef, 0x7e, 0xc5, 0xbf, 0xfd, 0x36,
	0xf7, 0x50, 0x3e, 0x0e, 0xcb, 0xe0, 0x2d, 0x54, 0xaf, 0x32, 0xa6, 0x7f, 0xf0, 0x82, 0xe7, 0xf9,
	0xa6, 0x8b, 0x34, 0x67, 0x7f, 0x65, 0xab, 0x22, 0x3b, 0xd8, 0xce, 0x13, 0xfb, 0xe6, 0x3b, 0x00,
	0x00, 0xff, 0xff, 0x25, 0x1e, 0x5d, 0x73, 0x74, 0x01, 0x00, 0x00,
}
