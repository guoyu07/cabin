// Code generated by protoc-gen-go. DO NOT EDIT.
// source: reply.proto

package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Reply struct {
	C int32 `protobuf:"varint,1,opt,name=c" json:"c,omitempty"`
}

func (m *Reply) Reset()                    { *m = Reply{} }
func (m *Reply) String() string            { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()               {}
func (*Reply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Reply) GetC() int32 {
	if m != nil {
		return m.C
	}
	return 0
}

func init() {
	proto.RegisterType((*Reply)(nil), "main.Reply")
}

func init() { proto.RegisterFile("reply.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 69 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x4a, 0x2d, 0xc8,
	0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53, 0x12, 0xe5,
	0x62, 0x0d, 0x02, 0x09, 0x0a, 0xf1, 0x70, 0x31, 0x26, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06,
	0x31, 0x26, 0x27, 0xb1, 0x81, 0xd5, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x2e, 0x8f,
	0x6d, 0x32, 0x00, 0x00, 0x00,
}
