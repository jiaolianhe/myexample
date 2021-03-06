// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greeter.proto

package greeter

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

type MyRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MyRequest) Reset()         { *m = MyRequest{} }
func (m *MyRequest) String() string { return proto.CompactTextString(m) }
func (*MyRequest) ProtoMessage()    {}
func (*MyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e585294ab3f34af5, []int{0}
}

func (m *MyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MyRequest.Unmarshal(m, b)
}
func (m *MyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MyRequest.Marshal(b, m, deterministic)
}
func (m *MyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyRequest.Merge(m, src)
}
func (m *MyRequest) XXX_Size() int {
	return xxx_messageInfo_MyRequest.Size(m)
}
func (m *MyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MyRequest proto.InternalMessageInfo

func (m *MyRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type MyResponse struct {
	Greeting             string   `protobuf:"bytes,2,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MyResponse) Reset()         { *m = MyResponse{} }
func (m *MyResponse) String() string { return proto.CompactTextString(m) }
func (*MyResponse) ProtoMessage()    {}
func (*MyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e585294ab3f34af5, []int{1}
}

func (m *MyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MyResponse.Unmarshal(m, b)
}
func (m *MyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MyResponse.Marshal(b, m, deterministic)
}
func (m *MyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyResponse.Merge(m, src)
}
func (m *MyResponse) XXX_Size() int {
	return xxx_messageInfo_MyResponse.Size(m)
}
func (m *MyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MyResponse proto.InternalMessageInfo

func (m *MyResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto.RegisterType((*MyRequest)(nil), "MyRequest")
	proto.RegisterType((*MyResponse)(nil), "MyResponse")
}

func init() { proto.RegisterFile("greeter.proto", fileDescriptor_e585294ab3f34af5) }

var fileDescriptor_e585294ab3f34af5 = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x2f, 0x4a, 0x4d,
	0x2d, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe7, 0xe2, 0xf4, 0xad, 0x0c,
	0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0x34, 0xb8, 0xb8, 0x40, 0x0a, 0x8a, 0x0b, 0xf2,
	0xf3, 0x8a, 0x53, 0x85, 0xa4, 0xb8, 0x38, 0xc0, 0xfa, 0x33, 0xf3, 0xd2, 0x25, 0x98, 0xc0, 0xaa,
	0xe0, 0x7c, 0x23, 0x5d, 0x2e, 0x76, 0x77, 0x88, 0xd9, 0x42, 0x4a, 0x5c, 0xac, 0x1e, 0xa9, 0x39,
	0x39, 0xf9, 0x42, 0x5c, 0x7a, 0x70, 0xd3, 0xa5, 0xb8, 0xf5, 0x10, 0x06, 0x29, 0x31, 0x24, 0xb1,
	0x81, 0x1d, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x08, 0xa0, 0xf3, 0x4c, 0x91, 0x00, 0x00,
	0x00,
}
