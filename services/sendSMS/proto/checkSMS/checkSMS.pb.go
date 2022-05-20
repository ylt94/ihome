// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/checkSMS/checkSMS.proto

package go_micro_service_sendSMS

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

type SMSType int32

const (
	SMSType_Register SMSType = 0
)

var SMSType_name = map[int32]string{
	0: "Register",
}

var SMSType_value = map[string]int32{
	"Register": 0,
}

func (x SMSType) String() string {
	return proto.EnumName(SMSType_name, int32(x))
}

func (SMSType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d819467b42ac65e7, []int{0}
}

type EheckStatus int32

const (
	EheckStatus_Fail   EheckStatus = 0
	EheckStatus_Succ   EheckStatus = 1
	EheckStatus_Expire EheckStatus = 2
)

var EheckStatus_name = map[int32]string{
	0: "Fail",
	1: "Succ",
	2: "Expire",
}

var EheckStatus_value = map[string]int32{
	"Fail":   0,
	"Succ":   1,
	"Expire": 2,
}

func (x EheckStatus) String() string {
	return proto.EnumName(EheckStatus_name, int32(x))
}

func (EheckStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d819467b42ac65e7, []int{1}
}

type Message struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_d819467b42ac65e7, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

type Request struct {
	Phone                string   `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Type                 SMSType  `protobuf:"varint,2,opt,name=Type,proto3,enum=go.micro.service.sendSMS.SMSType" json:"Type,omitempty"`
	Code                 string   `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_d819467b42ac65e7, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Request) GetType() SMSType {
	if m != nil {
		return m.Type
	}
	return SMSType_Register
}

func (m *Request) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type Response struct {
	Status               EheckStatus `protobuf:"varint,1,opt,name=status,proto3,enum=go.micro.service.sendSMS.EheckStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_d819467b42ac65e7, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() EheckStatus {
	if m != nil {
		return m.Status
	}
	return EheckStatus_Fail
}

func init() {
	proto.RegisterEnum("go.micro.service.sendSMS.SMSType", SMSType_name, SMSType_value)
	proto.RegisterEnum("go.micro.service.sendSMS.EheckStatus", EheckStatus_name, EheckStatus_value)
	proto.RegisterType((*Message)(nil), "go.micro.service.sendSMS.Message")
	proto.RegisterType((*Request)(nil), "go.micro.service.sendSMS.Request")
	proto.RegisterType((*Response)(nil), "go.micro.service.sendSMS.Response")
}

func init() { proto.RegisterFile("proto/checkSMS/checkSMS.proto", fileDescriptor_d819467b42ac65e7) }

var fileDescriptor_d819467b42ac65e7 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x14, 0xc4, 0x9b, 0x9a, 0x26, 0xe9, 0x53, 0x4a, 0x78, 0x08, 0x06, 0x41, 0xa8, 0x01, 0xa1, 0x14,
	0x5c, 0x21, 0xe2, 0xd1, 0x93, 0x54, 0xf0, 0x10, 0x90, 0xac, 0x47, 0x2f, 0x71, 0xfb, 0x48, 0xe3,
	0x9f, 0x6c, 0xdc, 0xb7, 0x11, 0xfd, 0xf6, 0x92, 0x6d, 0x10, 0x2f, 0xed, 0x6d, 0x76, 0xf9, 0xed,
	0x0c, 0x33, 0x0b, 0x67, 0xad, 0xd1, 0x56, 0x5f, 0xa9, 0x0d, 0xa9, 0x37, 0x99, 0xcb, 0x3f, 0x21,
	0xdc, 0x3d, 0x26, 0x95, 0x16, 0x1f, 0xb5, 0x32, 0x5a, 0x30, 0x99, 0xaf, 0x5a, 0x91, 0x60, 0x6a,
	0xd6, 0x32, 0x97, 0xe9, 0x14, 0xc2, 0x9c, 0x98, 0xcb, 0x8a, 0xd2, 0x57, 0x08, 0x0b, 0xfa, 0xec,
	0x88, 0x2d, 0x1e, 0xc3, 0xa4, 0xdd, 0xe8, 0x86, 0x12, 0x6f, 0xee, 0x2d, 0xa6, 0xc5, 0xf6, 0x80,
	0x37, 0xe0, 0x3f, 0xfd, 0xb4, 0x94, 0x8c, 0xe7, 0xde, 0x62, 0x96, 0x9d, 0x8b, 0x5d, 0xa6, 0x42,
	0xe6, 0xb2, 0x07, 0x0b, 0x87, 0x23, 0x82, 0xaf, 0xf4, 0x9a, 0x92, 0x03, 0xe7, 0xe5, 0x74, 0xfa,
	0x00, 0x51, 0x41, 0xdc, 0xea, 0x86, 0x09, 0x6f, 0x21, 0x60, 0x5b, 0xda, 0x8e, 0x5d, 0xda, 0x2c,
	0xbb, 0xd8, 0x6d, 0xbc, 0x72, 0xb5, 0x1c, 0x5c, 0x0c, 0x8f, 0x96, 0x27, 0x10, 0x0e, 0x79, 0x78,
	0xd4, 0xbb, 0x56, 0x35, 0x5b, 0x32, 0xf1, 0x68, 0x79, 0x09, 0x87, 0xff, 0x78, 0x8c, 0xc0, 0xbf,
	0x2f, 0xeb, 0xf7, 0x78, 0xd4, 0x2b, 0xd9, 0x29, 0x15, 0x7b, 0x08, 0x10, 0xac, 0xbe, 0xdb, 0xda,
	0x50, 0x3c, 0xce, 0x9e, 0x21, 0xba, 0x1b, 0x56, 0xc3, 0x47, 0x98, 0xb8, 0x05, 0x71, 0x4f, 0xc9,
	0x61, 0xab, 0xd3, 0x74, 0x1f, 0xb2, 0xad, 0x98, 0x8e, 0x5e, 0x02, 0xf7, 0x11, 0xd7, 0xbf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x20, 0x03, 0x86, 0xb2, 0xa9, 0x01, 0x00, 0x00,
}
