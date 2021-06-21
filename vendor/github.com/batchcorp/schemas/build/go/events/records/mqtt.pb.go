// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events/records/mqtt.proto

package records

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

type MQTTRecord struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Topic                string   `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Value                []byte   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Duplicate            bool     `protobuf:"varint,4,opt,name=duplicate,proto3" json:"duplicate,omitempty"`
	Retained             bool     `protobuf:"varint,5,opt,name=retained,proto3" json:"retained,omitempty"`
	Qos                  uint32   `protobuf:"varint,6,opt,name=qos,proto3" json:"qos,omitempty"`
	Timestamp            int64    `protobuf:"varint,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MQTTRecord) Reset()         { *m = MQTTRecord{} }
func (m *MQTTRecord) String() string { return proto.CompactTextString(m) }
func (*MQTTRecord) ProtoMessage()    {}
func (*MQTTRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_c391d35771173aef, []int{0}
}

func (m *MQTTRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MQTTRecord.Unmarshal(m, b)
}
func (m *MQTTRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MQTTRecord.Marshal(b, m, deterministic)
}
func (m *MQTTRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MQTTRecord.Merge(m, src)
}
func (m *MQTTRecord) XXX_Size() int {
	return xxx_messageInfo_MQTTRecord.Size(m)
}
func (m *MQTTRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_MQTTRecord.DiscardUnknown(m)
}

var xxx_messageInfo_MQTTRecord proto.InternalMessageInfo

func (m *MQTTRecord) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MQTTRecord) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *MQTTRecord) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *MQTTRecord) GetDuplicate() bool {
	if m != nil {
		return m.Duplicate
	}
	return false
}

func (m *MQTTRecord) GetRetained() bool {
	if m != nil {
		return m.Retained
	}
	return false
}

func (m *MQTTRecord) GetQos() uint32 {
	if m != nil {
		return m.Qos
	}
	return 0
}

func (m *MQTTRecord) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*MQTTRecord)(nil), "records.MQTTRecord")
}

func init() { proto.RegisterFile("events/records/mqtt.proto", fileDescriptor_c391d35771173aef) }

var fileDescriptor_c391d35771173aef = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xcd, 0x4e, 0xc4, 0x20,
	0x14, 0x85, 0x43, 0xeb, 0xfc, 0x11, 0x35, 0x86, 0x98, 0x88, 0xc6, 0x05, 0x71, 0xc5, 0x0a, 0x16,
	0xfa, 0x04, 0xee, 0x8d, 0x91, 0xcc, 0xca, 0x1d, 0x05, 0x32, 0x25, 0x29, 0x03, 0x03, 0xb7, 0xf3,
	0x60, 0x3e, 0xa1, 0x29, 0x9d, 0xd8, 0xb8, 0xe3, 0xfb, 0xb8, 0xb9, 0x37, 0xe7, 0xe0, 0x47, 0x77,
	0x76, 0x47, 0x28, 0x32, 0x3b, 0x13, 0xb3, 0x2d, 0x32, 0x9c, 0x00, 0x44, 0xca, 0x11, 0x22, 0xd9,
	0x5c, 0xdc, 0xcb, 0x0f, 0xc2, 0xf8, 0xe3, 0x6b, 0xbf, 0x57, 0x95, 0xc9, 0x2d, 0x6e, 0xbc, 0xa5,
	0x88, 0x21, 0x7e, 0xa3, 0x1a, 0x6f, 0xc9, 0x3d, 0x5e, 0x41, 0x4c, 0xde, 0xd0, 0x86, 0x21, 0xbe,
	0x53, 0x33, 0x4c, 0xf6, 0xac, 0x87, 0xd1, 0xd1, 0x96, 0x21, 0x7e, 0xad, 0x66, 0x20, 0xcf, 0x78,
	0x67, 0xc7, 0x34, 0x78, 0xa3, 0xc1, 0xd1, 0x2b, 0x86, 0xf8, 0x56, 0x2d, 0x82, 0x3c, 0xe1, 0x6d,
	0x76, 0xa0, 0xfd, 0xd1, 0x59, 0xba, 0xaa, 0x9f, 0x7f, 0x4c, 0xee, 0x70, 0x7b, 0x8a, 0x85, 0xae,
	0xeb, 0xd9, 0xe9, 0x39, 0xed, 0x02, 0x1f, 0x5c, 0x01, 0x1d, 0x12, 0xdd, 0x30, 0xc4, 0x5b, 0xb5,
	0x88, 0xf7, 0x4f, 0xfc, 0x50, 0x7a, 0xd1, 0x69, 0x30, 0xbd, 0x98, 0x33, 0x8a, 0x4b, 0x9e, 0xef,
	0xb7, 0x83, 0x87, 0x7e, 0xec, 0x84, 0x89, 0x41, 0xd6, 0x01, 0x13, 0x73, 0x92, 0xc5, 0xf4, 0x2e,
	0xe8, 0x22, 0xbb, 0xd1, 0x0f, 0x56, 0x1e, 0xa2, 0xfc, 0xdf, 0x4c, 0xb7, 0xae, 0xad, 0xbc, 0xfe,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x2c, 0xa9, 0x20, 0x36, 0x32, 0x01, 0x00, 0x00,
}
