// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/core.proto

/*
Package core is a generated protocol buffer package.

It is generated from these files:
	core/core.proto

It has these top-level messages:
	Coordinates
	Color
*/
package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Coordinates struct {
	X float64 `protobuf:"fixed64,1,opt,name=x" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y" json:"y,omitempty"`
	Z float64 `protobuf:"fixed64,3,opt,name=z" json:"z,omitempty"`
}

func (m *Coordinates) Reset()                    { *m = Coordinates{} }
func (m *Coordinates) String() string            { return proto.CompactTextString(m) }
func (*Coordinates) ProtoMessage()               {}
func (*Coordinates) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Coordinates) GetX() float64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Coordinates) GetY() float64 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Coordinates) GetZ() float64 {
	if m != nil {
		return m.Z
	}
	return 0
}

type Color struct {
	R float32 `protobuf:"fixed32,1,opt,name=r" json:"r,omitempty"`
	G float32 `protobuf:"fixed32,2,opt,name=g" json:"g,omitempty"`
	B float32 `protobuf:"fixed32,3,opt,name=b" json:"b,omitempty"`
}

func (m *Color) Reset()                    { *m = Color{} }
func (m *Color) String() string            { return proto.CompactTextString(m) }
func (*Color) ProtoMessage()               {}
func (*Color) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Color) GetR() float32 {
	if m != nil {
		return m.R
	}
	return 0
}

func (m *Color) GetG() float32 {
	if m != nil {
		return m.G
	}
	return 0
}

func (m *Color) GetB() float32 {
	if m != nil {
		return m.B
	}
	return 0
}

func init() {
	proto.RegisterType((*Coordinates)(nil), "core.Coordinates")
	proto.RegisterType((*Color)(nil), "core.Color")
}

func init() { proto.RegisterFile("core/core.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xce, 0x2f, 0x4a,
	0xd5, 0x07, 0x11, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x2c, 0x20, 0xb6, 0x92, 0x39, 0x17,
	0xb7, 0x73, 0x7e, 0x7e, 0x51, 0x4a, 0x66, 0x5e, 0x62, 0x49, 0x6a, 0xb1, 0x10, 0x0f, 0x17, 0x63,
	0x85, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x63, 0x10, 0x63, 0x05, 0x88, 0x57, 0x29, 0xc1, 0x04, 0xe1,
	0x55, 0x82, 0x78, 0x55, 0x12, 0xcc, 0x10, 0x5e, 0x95, 0x92, 0x21, 0x17, 0xab, 0x73, 0x7e, 0x4e,
	0x7e, 0x11, 0x48, 0xb8, 0x08, 0xac, 0x85, 0x29, 0x88, 0x11, 0xcc, 0x4b, 0x07, 0x6b, 0x61, 0x0a,
	0x62, 0x4c, 0x07, 0xf1, 0x92, 0xc0, 0x5a, 0x98, 0x82, 0x18, 0x93, 0x92, 0xd8, 0xc0, 0x16, 0x1b,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x66, 0x4c, 0xae, 0x8b, 0x00, 0x00, 0x00,
}
