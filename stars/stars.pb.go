// Code generated by protoc-gen-go.
// source: stars/stars.proto
// DO NOT EDIT!

/*
Package stars is a generated protocol buffer package.

It is generated from these files:
	stars/stars.proto

It has these top-level messages:
	GetStarsRequest
	GetStarsResponse
	Star
*/
package stars

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "core"

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

type GetStarsRequest struct {
}

func (m *GetStarsRequest) Reset()                    { *m = GetStarsRequest{} }
func (m *GetStarsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetStarsRequest) ProtoMessage()               {}
func (*GetStarsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetStarsResponse struct {
	UnixTime int64   `protobuf:"varint,1,opt,name=unixTime" json:"unixTime,omitempty"`
	Stars    []*Star `protobuf:"bytes,2,rep,name=stars" json:"stars,omitempty"`
}

func (m *GetStarsResponse) Reset()                    { *m = GetStarsResponse{} }
func (m *GetStarsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetStarsResponse) ProtoMessage()               {}
func (*GetStarsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetStarsResponse) GetUnixTime() int64 {
	if m != nil {
		return m.UnixTime
	}
	return 0
}

func (m *GetStarsResponse) GetStars() []*Star {
	if m != nil {
		return m.Stars
	}
	return nil
}

type Star struct {
	Id             int32             `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name           string            `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Coordinates    *core.Coordinates `protobuf:"bytes,3,opt,name=coordinates" json:"coordinates,omitempty"`
	Classification string            `protobuf:"bytes,4,opt,name=classification" json:"classification,omitempty"`
	Temperature    float64           `protobuf:"fixed64,5,opt,name=temperature" json:"temperature,omitempty"`
	Mass           float64           `protobuf:"fixed64,6,opt,name=mass" json:"mass,omitempty"`
	Radius         float64           `protobuf:"fixed64,7,opt,name=radius" json:"radius,omitempty"`
	Color          *core.Color       `protobuf:"bytes,8,opt,name=color" json:"color,omitempty"`
}

func (m *Star) Reset()                    { *m = Star{} }
func (m *Star) String() string            { return proto.CompactTextString(m) }
func (*Star) ProtoMessage()               {}
func (*Star) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Star) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Star) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Star) GetCoordinates() *core.Coordinates {
	if m != nil {
		return m.Coordinates
	}
	return nil
}

func (m *Star) GetClassification() string {
	if m != nil {
		return m.Classification
	}
	return ""
}

func (m *Star) GetTemperature() float64 {
	if m != nil {
		return m.Temperature
	}
	return 0
}

func (m *Star) GetMass() float64 {
	if m != nil {
		return m.Mass
	}
	return 0
}

func (m *Star) GetRadius() float64 {
	if m != nil {
		return m.Radius
	}
	return 0
}

func (m *Star) GetColor() *core.Color {
	if m != nil {
		return m.Color
	}
	return nil
}

func init() {
	proto.RegisterType((*GetStarsRequest)(nil), "stars.GetStarsRequest")
	proto.RegisterType((*GetStarsResponse)(nil), "stars.GetStarsResponse")
	proto.RegisterType((*Star)(nil), "stars.Star")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for StarService service

type StarServiceClient interface {
	GetStars(ctx context.Context, in *GetStarsRequest, opts ...grpc.CallOption) (*GetStarsResponse, error)
}

type starServiceClient struct {
	cc *grpc.ClientConn
}

func NewStarServiceClient(cc *grpc.ClientConn) StarServiceClient {
	return &starServiceClient{cc}
}

func (c *starServiceClient) GetStars(ctx context.Context, in *GetStarsRequest, opts ...grpc.CallOption) (*GetStarsResponse, error) {
	out := new(GetStarsResponse)
	err := grpc.Invoke(ctx, "/stars.StarService/GetStars", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StarService service

type StarServiceServer interface {
	GetStars(context.Context, *GetStarsRequest) (*GetStarsResponse, error)
}

func RegisterStarServiceServer(s *grpc.Server, srv StarServiceServer) {
	s.RegisterService(&_StarService_serviceDesc, srv)
}

func _StarService_GetStars_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStarsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarServiceServer).GetStars(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stars.StarService/GetStars",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarServiceServer).GetStars(ctx, req.(*GetStarsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StarService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stars.StarService",
	HandlerType: (*StarServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStars",
			Handler:    _StarService_GetStars_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stars/stars.proto",
}

func init() { proto.RegisterFile("stars/stars.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x51, 0xbd, 0x4e, 0x33, 0x31,
	0x10, 0x94, 0x2f, 0x3f, 0x5f, 0xbe, 0xb5, 0x94, 0x90, 0x2d, 0x82, 0x95, 0xea, 0x48, 0x81, 0xae,
	0x4a, 0xa4, 0xa4, 0xa4, 0xa4, 0x40, 0xa2, 0xc3, 0xe1, 0x05, 0x8c, 0x6f, 0x91, 0x2c, 0xe5, 0xce,
	0x87, 0xed, 0x43, 0x3c, 0x3a, 0x25, 0xb2, 0x9d, 0x90, 0x28, 0x34, 0xab, 0x9d, 0x19, 0x6b, 0x3c,
	0xbb, 0x0b, 0x73, 0x1f, 0x94, 0xf3, 0x9b, 0x54, 0xd7, 0x9d, 0xb3, 0xc1, 0xe2, 0x28, 0x81, 0xe5,
	0x4c, 0x5b, 0x47, 0x9b, 0x58, 0x32, 0xbf, 0x9a, 0xc3, 0xec, 0x89, 0xc2, 0x3e, 0x8a, 0x92, 0x3e,
	0x7a, 0xf2, 0x61, 0xf5, 0x02, 0x37, 0x67, 0xca, 0x77, 0xb6, 0xf5, 0x84, 0x4b, 0x98, 0xf4, 0xad,
	0xf9, 0x7a, 0x35, 0x0d, 0x09, 0x56, 0xb2, 0x6a, 0x20, 0x7f, 0x31, 0xde, 0x41, 0x36, 0x17, 0x45,
	0x39, 0xa8, 0xf8, 0x96, 0xaf, 0xf3, 0xbf, 0xd1, 0x40, 0x66, 0x65, 0xf5, 0xcd, 0x60, 0x18, 0x31,
	0x4e, 0xa1, 0x30, 0x75, 0x72, 0x18, 0xc9, 0xc2, 0xd4, 0x88, 0x30, 0x6c, 0x55, 0x43, 0xa2, 0x28,
	0x59, 0xf5, 0x5f, 0xa6, 0x1e, 0x77, 0xc0, 0xb5, 0xb5, 0xae, 0x36, 0xad, 0x0a, 0xe4, 0xc5, 0xa0,
	0x64, 0x15, 0xdf, 0xce, 0xd7, 0x29, 0xf4, 0xe3, 0x59, 0x90, 0x97, 0xaf, 0xf0, 0x1e, 0xa6, 0xfa,
	0xa0, 0xbc, 0x37, 0xef, 0x46, 0xab, 0x60, 0x6c, 0x2b, 0x86, 0xc9, 0xf2, 0x8a, 0xc5, 0x12, 0x78,
	0xa0, 0xa6, 0x23, 0xa7, 0x42, 0xef, 0x48, 0x8c, 0x4a, 0x56, 0x31, 0x79, 0x49, 0xc5, 0x48, 0x8d,
	0xf2, 0x5e, 0x8c, 0x93, 0x94, 0x7a, 0x5c, 0xc0, 0xd8, 0xa9, 0xda, 0xf4, 0x5e, 0xfc, 0x4b, 0xec,
	0x11, 0xc5, 0xd1, 0xb5, 0x3d, 0x58, 0x27, 0x26, 0x29, 0x24, 0x3f, 0x85, 0x3c, 0x58, 0x27, 0xb3,
	0xb2, 0x7d, 0x06, 0x1e, 0x27, 0xdf, 0x93, 0xfb, 0x34, 0x9a, 0xf0, 0x01, 0x26, 0xa7, 0xe5, 0xe2,
	0xe2, 0xb8, 0xa9, 0xab, 0x03, 0x2c, 0x6f, 0xff, 0xf0, 0xf9, 0x0a, 0x6f, 0xe3, 0x74, 0xb3, 0xdd,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x59, 0x30, 0x57, 0xe0, 0x01, 0x00, 0x00,
}