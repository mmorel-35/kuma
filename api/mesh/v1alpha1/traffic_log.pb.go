// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mesh/v1alpha1/traffic_log.proto

package v1alpha1

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

type TrafficLog struct {
	Rules                []*TrafficLog_Rule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TrafficLog) Reset()         { *m = TrafficLog{} }
func (m *TrafficLog) String() string { return proto.CompactTextString(m) }
func (*TrafficLog) ProtoMessage()    {}
func (*TrafficLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_47c4f4c9c894eeed, []int{0}
}

func (m *TrafficLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficLog.Unmarshal(m, b)
}
func (m *TrafficLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficLog.Marshal(b, m, deterministic)
}
func (m *TrafficLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficLog.Merge(m, src)
}
func (m *TrafficLog) XXX_Size() int {
	return xxx_messageInfo_TrafficLog.Size(m)
}
func (m *TrafficLog) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficLog.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficLog proto.InternalMessageInfo

func (m *TrafficLog) GetRules() []*TrafficLog_Rule {
	if m != nil {
		return m.Rules
	}
	return nil
}

type TrafficLog_Rule struct {
	Sources              []*Selector           `protobuf:"bytes,1,rep,name=sources,proto3" json:"sources,omitempty"`
	Destinations         []*Selector           `protobuf:"bytes,2,rep,name=destinations,proto3" json:"destinations,omitempty"`
	Conf                 *TrafficLog_Rule_Conf `protobuf:"bytes,3,opt,name=conf,proto3" json:"conf,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TrafficLog_Rule) Reset()         { *m = TrafficLog_Rule{} }
func (m *TrafficLog_Rule) String() string { return proto.CompactTextString(m) }
func (*TrafficLog_Rule) ProtoMessage()    {}
func (*TrafficLog_Rule) Descriptor() ([]byte, []int) {
	return fileDescriptor_47c4f4c9c894eeed, []int{0, 0}
}

func (m *TrafficLog_Rule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficLog_Rule.Unmarshal(m, b)
}
func (m *TrafficLog_Rule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficLog_Rule.Marshal(b, m, deterministic)
}
func (m *TrafficLog_Rule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficLog_Rule.Merge(m, src)
}
func (m *TrafficLog_Rule) XXX_Size() int {
	return xxx_messageInfo_TrafficLog_Rule.Size(m)
}
func (m *TrafficLog_Rule) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficLog_Rule.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficLog_Rule proto.InternalMessageInfo

func (m *TrafficLog_Rule) GetSources() []*Selector {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *TrafficLog_Rule) GetDestinations() []*Selector {
	if m != nil {
		return m.Destinations
	}
	return nil
}

func (m *TrafficLog_Rule) GetConf() *TrafficLog_Rule_Conf {
	if m != nil {
		return m.Conf
	}
	return nil
}

type TrafficLog_Rule_Conf struct {
	Backend              string   `protobuf:"bytes,1,opt,name=backend,proto3" json:"backend,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrafficLog_Rule_Conf) Reset()         { *m = TrafficLog_Rule_Conf{} }
func (m *TrafficLog_Rule_Conf) String() string { return proto.CompactTextString(m) }
func (*TrafficLog_Rule_Conf) ProtoMessage()    {}
func (*TrafficLog_Rule_Conf) Descriptor() ([]byte, []int) {
	return fileDescriptor_47c4f4c9c894eeed, []int{0, 0, 0}
}

func (m *TrafficLog_Rule_Conf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficLog_Rule_Conf.Unmarshal(m, b)
}
func (m *TrafficLog_Rule_Conf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficLog_Rule_Conf.Marshal(b, m, deterministic)
}
func (m *TrafficLog_Rule_Conf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficLog_Rule_Conf.Merge(m, src)
}
func (m *TrafficLog_Rule_Conf) XXX_Size() int {
	return xxx_messageInfo_TrafficLog_Rule_Conf.Size(m)
}
func (m *TrafficLog_Rule_Conf) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficLog_Rule_Conf.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficLog_Rule_Conf proto.InternalMessageInfo

func (m *TrafficLog_Rule_Conf) GetBackend() string {
	if m != nil {
		return m.Backend
	}
	return ""
}

func init() {
	proto.RegisterType((*TrafficLog)(nil), "kuma.mesh.v1alpha1.TrafficLog")
	proto.RegisterType((*TrafficLog_Rule)(nil), "kuma.mesh.v1alpha1.TrafficLog.Rule")
	proto.RegisterType((*TrafficLog_Rule_Conf)(nil), "kuma.mesh.v1alpha1.TrafficLog.Rule.Conf")
}

func init() { proto.RegisterFile("mesh/v1alpha1/traffic_log.proto", fileDescriptor_47c4f4c9c894eeed) }

var fileDescriptor_47c4f4c9c894eeed = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x46, 0x49, 0x3b, 0x5a, 0xbd, 0x75, 0x95, 0x55, 0x18, 0x0a, 0x0e, 0xba, 0x99, 0x55, 0x4a,
	0x2b, 0x08, 0x82, 0x0b, 0xd1, 0xad, 0xab, 0xe8, 0xca, 0x8d, 0xa4, 0x69, 0xd2, 0x0e, 0x4d, 0x73,
	0x4b, 0x7e, 0x7c, 0x28, 0x9f, 0xca, 0x47, 0x91, 0xe9, 0x18, 0xa5, 0xd8, 0xc5, 0x2c, 0x3f, 0x38,
	0xe7, 0x70, 0xb9, 0x70, 0xb9, 0xd5, 0x61, 0x3d, 0xfd, 0x98, 0x49, 0xbb, 0x5b, 0xcb, 0xd9, 0x34,
	0x7a, 0x69, 0x4c, 0xa3, 0xde, 0x2d, 0xae, 0xf8, 0xce, 0x63, 0x44, 0x4a, 0x37, 0x69, 0x2b, 0x79,
	0x4b, 0xf1, 0x4c, 0x95, 0x93, 0x43, 0x29, 0x68, 0xab, 0x55, 0x44, 0xdf, 0x19, 0x57, 0x9f, 0x03,
	0x80, 0xd7, 0xae, 0xf3, 0x8c, 0x2b, 0x7a, 0x07, 0x27, 0x3e, 0x59, 0x1d, 0x18, 0xa9, 0x86, 0xf5,
	0x78, 0x7e, 0xcd, 0xff, 0x07, 0xf9, 0x1f, 0xce, 0x45, 0xb2, 0x5a, 0x74, 0x46, 0xf9, 0x45, 0xa0,
	0x68, 0x37, 0xbd, 0x85, 0x51, 0xc0, 0xe4, 0xd5, 0x6f, 0x65, 0x72, 0xac, 0xf2, 0xf2, 0x73, 0x87,
	0xc8, 0x30, 0x7d, 0x80, 0x8b, 0xa5, 0x0e, 0xb1, 0x71, 0x32, 0x36, 0xe8, 0x02, 0x1b, 0xf4, 0x90,
	0x0f, 0x0c, 0x7a, 0x0f, 0x85, 0x42, 0x67, 0xd8, 0xb0, 0x22, 0xf5, 0x78, 0x5e, 0xf7, 0x38, 0x9e,
	0x3f, 0xa1, 0x33, 0x62, 0x6f, 0x95, 0x15, 0x14, 0xed, 0xa2, 0x0c, 0x46, 0x0b, 0xa9, 0x36, 0xda,
	0x2d, 0x19, 0xa9, 0x48, 0x7d, 0x2e, 0xf2, 0x7c, 0x84, 0xb7, 0xb3, 0x1c, 0x5a, 0x9c, 0xee, 0xff,
	0x77, 0xf3, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xfb, 0xbb, 0xf7, 0x43, 0x94, 0x01, 0x00, 0x00,
}
