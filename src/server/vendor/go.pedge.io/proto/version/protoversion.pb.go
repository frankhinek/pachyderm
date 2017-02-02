// Code generated by protoc-gen-go.
// source: version/protoversion.proto
// DO NOT EDIT!

package protoversion

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import google_protobuf1 "go.pedge.io/pb/go/google/protobuf"

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

type Version struct {
	Major      uint32 `protobuf:"varint,1,opt,name=major" json:"major,omitempty"`
	Minor      uint32 `protobuf:"varint,2,opt,name=minor" json:"minor,omitempty"`
	Micro      uint32 `protobuf:"varint,3,opt,name=micro" json:"micro,omitempty"`
	Additional string `protobuf:"bytes,4,opt,name=additional" json:"additional,omitempty"`
}

func (m *Version) Reset()                    { *m = Version{} }
func (m *Version) String() string            { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()               {}
func (*Version) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Version)(nil), "protoversion.Version")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for API service

type APIClient interface {
	GetVersion(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*Version, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) GetVersion(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*Version, error) {
	out := new(Version)
	err := grpc.Invoke(ctx, "/protoversion.API/GetVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	GetVersion(context.Context, *google_protobuf1.Empty) (*Version, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoversion.API/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetVersion(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protoversion.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _API_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("version/protoversion.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x4b, 0x2d, 0x2a,
	0xce, 0xcc, 0xcf, 0xd3, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x87, 0x72, 0xf4, 0xc0, 0x1c, 0x21, 0x1e,
	0x64, 0x31, 0x29, 0x99, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0xfd, 0xc4, 0x82, 0x4c, 0xfd, 0xc4,
	0xbc, 0xbc, 0xfc, 0x92, 0xc4, 0x12, 0xa0, 0x70, 0x31, 0x44, 0xad, 0x94, 0x34, 0x54, 0x16, 0xcc,
	0x4b, 0x2a, 0x4d, 0xd3, 0x4f, 0xcd, 0x2d, 0x28, 0xa9, 0x84, 0x48, 0x2a, 0x65, 0x73, 0xb1, 0x87,
	0x41, 0x4c, 0x11, 0x12, 0xe1, 0x62, 0xcd, 0x4d, 0xcc, 0xca, 0x2f, 0x92, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0d, 0x82, 0x70, 0xc0, 0xa2, 0x99, 0x79, 0x40, 0x51, 0x26, 0xa8, 0x28, 0x88, 0x03, 0x11,
	0x4d, 0x2e, 0xca, 0x97, 0x60, 0x86, 0x89, 0x02, 0x39, 0x42, 0x72, 0x5c, 0x5c, 0x89, 0x29, 0x29,
	0x99, 0x20, 0xcb, 0x13, 0x73, 0x24, 0x58, 0x80, 0x52, 0x9c, 0x41, 0x48, 0x22, 0x46, 0x21, 0x5c,
	0xcc, 0x8e, 0x01, 0x9e, 0x42, 0xbe, 0x5c, 0x5c, 0xee, 0xa9, 0x25, 0x30, 0x6b, 0xc5, 0xf4, 0x20,
	0xee, 0xd3, 0x83, 0xb9, 0x4f, 0xcf, 0x15, 0xe4, 0x3e, 0x29, 0x51, 0x3d, 0x14, 0x7f, 0x43, 0x95,
	0x2b, 0x09, 0x34, 0x5d, 0x7e, 0x32, 0x99, 0x89, 0x4b, 0x88, 0x43, 0x1f, 0x2a, 0x93, 0xc4, 0x06,
	0x56, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x03, 0x2c, 0xc9, 0x7d, 0x30, 0x01, 0x00, 0x00,
}