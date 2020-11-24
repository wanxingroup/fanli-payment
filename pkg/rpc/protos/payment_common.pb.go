// Code generated by protoc-gen-go. DO NOT EDIT.
// source: payment_common.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Error struct {
	Code    int64  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Stack   *Error `protobuf:"bytes,3,opt,name=stack" json:"stack,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Error) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetStack() *Error {
	if m != nil {
		return m.Stack
	}
	return nil
}

func init() {
	proto.RegisterType((*Error)(nil), "payment.pkg.rpc.protos.Error")
}

func init() { proto.RegisterFile("payment_common.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x48, 0xac, 0xcc,
	0x4d, 0xcd, 0x2b, 0x89, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x12, 0x83, 0x8a, 0xea, 0x15, 0x64, 0xa7, 0xeb, 0x15, 0x15, 0x24, 0x43, 0x84, 0x8b, 0x95,
	0xb2, 0xb8, 0x58, 0x5d, 0x8b, 0x8a, 0xf2, 0x8b, 0x84, 0x84, 0xb8, 0x58, 0x92, 0xf3, 0x53, 0x52,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0xc0, 0x6c, 0x21, 0x09, 0x2e, 0xf6, 0xdc, 0xd4, 0xe2,
	0xe2, 0xc4, 0xf4, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x57, 0xc8, 0x98, 0x8b,
	0xb5, 0xb8, 0x24, 0x31, 0x39, 0x5b, 0x82, 0x59, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x56, 0x0f, 0xbb,
	0xf1, 0x7a, 0x60, 0xb3, 0x83, 0x20, 0x6a, 0x9d, 0x38, 0xa2, 0xd8, 0x20, 0xc2, 0x49, 0x10, 0xda,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x01, 0x39, 0x80, 0x34, 0xac, 0x00, 0x00, 0x00,
}
