// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package go_micro_service_user

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

type AvatarStatus int32

const (
	AvatarStatus_AvatarStatusFail AvatarStatus = 0
	AvatarStatus_AvatarStatusSucc AvatarStatus = 1
)

var AvatarStatus_name = map[int32]string{
	0: "AvatarStatusFail",
	1: "AvatarStatusSucc",
}

var AvatarStatus_value = map[string]int32{
	"AvatarStatusFail": 0,
	"AvatarStatusSucc": 1,
}

func (x AvatarStatus) String() string {
	return proto.EnumName(AvatarStatus_name, int32(x))
}

func (AvatarStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

type UpNameStatus int32

const (
	UpNameStatus_UpNameFail UpNameStatus = 0
	UpNameStatus_UpNameSucc UpNameStatus = 1
)

var UpNameStatus_name = map[int32]string{
	0: "UpNameFail",
	1: "UpNameSucc",
}

var UpNameStatus_value = map[string]int32{
	"UpNameFail": 0,
	"UpNameSucc": 1,
}

func (x UpNameStatus) String() string {
	return proto.EnumName(UpNameStatus_name, int32(x))
}

func (UpNameStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

type InfoRequest struct {
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InfoRequest) Reset()         { *m = InfoRequest{} }
func (m *InfoRequest) String() string { return proto.CompactTextString(m) }
func (*InfoRequest) ProtoMessage()    {}
func (*InfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

func (m *InfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfoRequest.Unmarshal(m, b)
}
func (m *InfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfoRequest.Marshal(b, m, deterministic)
}
func (m *InfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoRequest.Merge(m, src)
}
func (m *InfoRequest) XXX_Size() int {
	return xxx_messageInfo_InfoRequest.Size(m)
}
func (m *InfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InfoRequest proto.InternalMessageInfo

func (m *InfoRequest) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type InfoResponse struct {
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Mobile               string   `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	RealName             string   `protobuf:"bytes,4,opt,name=real_name,json=realName,proto3" json:"real_name,omitempty"`
	IdCard               string   `protobuf:"bytes,5,opt,name=id_card,json=idCard,proto3" json:"id_card,omitempty"`
	AvatarUrl            string   `protobuf:"bytes,6,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InfoResponse) Reset()         { *m = InfoResponse{} }
func (m *InfoResponse) String() string { return proto.CompactTextString(m) }
func (*InfoResponse) ProtoMessage()    {}
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

func (m *InfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfoResponse.Unmarshal(m, b)
}
func (m *InfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfoResponse.Marshal(b, m, deterministic)
}
func (m *InfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoResponse.Merge(m, src)
}
func (m *InfoResponse) XXX_Size() int {
	return xxx_messageInfo_InfoResponse.Size(m)
}
func (m *InfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InfoResponse proto.InternalMessageInfo

func (m *InfoResponse) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *InfoResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InfoResponse) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *InfoResponse) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

func (m *InfoResponse) GetIdCard() string {
	if m != nil {
		return m.IdCard
	}
	return ""
}

func (m *InfoResponse) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

type AvatarRequest struct {
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AvatarUrl            string   `protobuf:"bytes,2,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AvatarRequest) Reset()         { *m = AvatarRequest{} }
func (m *AvatarRequest) String() string { return proto.CompactTextString(m) }
func (*AvatarRequest) ProtoMessage()    {}
func (*AvatarRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *AvatarRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvatarRequest.Unmarshal(m, b)
}
func (m *AvatarRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvatarRequest.Marshal(b, m, deterministic)
}
func (m *AvatarRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvatarRequest.Merge(m, src)
}
func (m *AvatarRequest) XXX_Size() int {
	return xxx_messageInfo_AvatarRequest.Size(m)
}
func (m *AvatarRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AvatarRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AvatarRequest proto.InternalMessageInfo

func (m *AvatarRequest) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *AvatarRequest) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

type AvatarResponse struct {
	Status               AvatarStatus `protobuf:"varint,1,opt,name=status,proto3,enum=go.micro.service.user.AvatarStatus" json:"status,omitempty"`
	AvatarUrl            string       `protobuf:"bytes,2,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AvatarResponse) Reset()         { *m = AvatarResponse{} }
func (m *AvatarResponse) String() string { return proto.CompactTextString(m) }
func (*AvatarResponse) ProtoMessage()    {}
func (*AvatarResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
}

func (m *AvatarResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvatarResponse.Unmarshal(m, b)
}
func (m *AvatarResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvatarResponse.Marshal(b, m, deterministic)
}
func (m *AvatarResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvatarResponse.Merge(m, src)
}
func (m *AvatarResponse) XXX_Size() int {
	return xxx_messageInfo_AvatarResponse.Size(m)
}
func (m *AvatarResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AvatarResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AvatarResponse proto.InternalMessageInfo

func (m *AvatarResponse) GetStatus() AvatarStatus {
	if m != nil {
		return m.Status
	}
	return AvatarStatus_AvatarStatusFail
}

func (m *AvatarResponse) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

type UpNameRequest struct {
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpNameRequest) Reset()         { *m = UpNameRequest{} }
func (m *UpNameRequest) String() string { return proto.CompactTextString(m) }
func (*UpNameRequest) ProtoMessage()    {}
func (*UpNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{4}
}

func (m *UpNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpNameRequest.Unmarshal(m, b)
}
func (m *UpNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpNameRequest.Marshal(b, m, deterministic)
}
func (m *UpNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpNameRequest.Merge(m, src)
}
func (m *UpNameRequest) XXX_Size() int {
	return xxx_messageInfo_UpNameRequest.Size(m)
}
func (m *UpNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpNameRequest proto.InternalMessageInfo

func (m *UpNameRequest) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UpNameRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UpNameResponse struct {
	Status               UpNameStatus `protobuf:"varint,1,opt,name=status,proto3,enum=go.micro.service.user.UpNameStatus" json:"status,omitempty"`
	Name                 string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpNameResponse) Reset()         { *m = UpNameResponse{} }
func (m *UpNameResponse) String() string { return proto.CompactTextString(m) }
func (*UpNameResponse) ProtoMessage()    {}
func (*UpNameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{5}
}

func (m *UpNameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpNameResponse.Unmarshal(m, b)
}
func (m *UpNameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpNameResponse.Marshal(b, m, deterministic)
}
func (m *UpNameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpNameResponse.Merge(m, src)
}
func (m *UpNameResponse) XXX_Size() int {
	return xxx_messageInfo_UpNameResponse.Size(m)
}
func (m *UpNameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpNameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpNameResponse proto.InternalMessageInfo

func (m *UpNameResponse) GetStatus() UpNameStatus {
	if m != nil {
		return m.Status
	}
	return UpNameStatus_UpNameFail
}

func (m *UpNameResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterEnum("go.micro.service.user.AvatarStatus", AvatarStatus_name, AvatarStatus_value)
	proto.RegisterEnum("go.micro.service.user.UpNameStatus", UpNameStatus_name, UpNameStatus_value)
	proto.RegisterType((*InfoRequest)(nil), "go.micro.service.user.InfoRequest")
	proto.RegisterType((*InfoResponse)(nil), "go.micro.service.user.InfoResponse")
	proto.RegisterType((*AvatarRequest)(nil), "go.micro.service.user.AvatarRequest")
	proto.RegisterType((*AvatarResponse)(nil), "go.micro.service.user.AvatarResponse")
	proto.RegisterType((*UpNameRequest)(nil), "go.micro.service.user.UpNameRequest")
	proto.RegisterType((*UpNameResponse)(nil), "go.micro.service.user.UpNameResponse")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xd1, 0x4e, 0xc2, 0x30,
	0x14, 0x65, 0x38, 0x87, 0x5c, 0x61, 0x21, 0x8d, 0xc8, 0xa2, 0x31, 0x21, 0x03, 0x0d, 0xe1, 0x61,
	0x26, 0xf8, 0xa6, 0xbe, 0x18, 0x13, 0x0d, 0x2f, 0x26, 0x8e, 0xf0, 0x64, 0xcc, 0x52, 0xd6, 0x6a,
	0x96, 0x0c, 0x8a, 0xed, 0xc6, 0x9f, 0xf8, 0x0f, 0x7e, 0xa6, 0x69, 0x3b, 0xc8, 0x5c, 0x80, 0xc5,
	0x17, 0xc2, 0xb9, 0x3d, 0xe7, 0xf4, 0xdc, 0xdb, 0x3b, 0x68, 0x2f, 0x39, 0x4b, 0xd8, 0x75, 0x2a,
	0x28, 0x57, 0x3f, 0x9e, 0xc2, 0xa8, 0xfd, 0xc9, 0xbc, 0x79, 0x14, 0x72, 0xe6, 0x09, 0xca, 0x57,
	0x51, 0x48, 0x3d, 0x79, 0xe8, 0x5e, 0xc1, 0xf1, 0x78, 0xf1, 0xc1, 0x7c, 0xfa, 0x95, 0x52, 0x91,
	0xa0, 0x0e, 0xd4, 0x64, 0x39, 0x88, 0x88, 0x63, 0x74, 0x8d, 0x41, 0xd3, 0xb7, 0x24, 0x1c, 0x13,
	0xf7, 0xc7, 0x80, 0x86, 0x26, 0x8a, 0x25, 0x5b, 0x08, 0xba, 0x93, 0x89, 0x10, 0x98, 0x0b, 0x3c,
	0xa7, 0x4e, 0xb5, 0x6b, 0x0c, 0xea, 0xbe, 0xfa, 0x8f, 0x4e, 0xc1, 0x9a, 0xb3, 0x59, 0x14, 0x53,
	0xe7, 0x40, 0x55, 0x33, 0x84, 0xce, 0xa1, 0xce, 0x29, 0x8e, 0x03, 0x25, 0x30, 0xd5, 0xd1, 0x91,
	0x2c, 0xbc, 0x48, 0x51, 0x07, 0x6a, 0x11, 0x09, 0x42, 0xcc, 0x89, 0x73, 0xa8, 0x55, 0x11, 0x79,
	0xc4, 0x9c, 0xa0, 0x0b, 0x00, 0xbc, 0xc2, 0x09, 0xe6, 0x41, 0xca, 0x63, 0xc7, 0x52, 0x67, 0x75,
	0x5d, 0x99, 0xf2, 0xd8, 0x7d, 0x86, 0xe6, 0x83, 0x02, 0x65, 0x4d, 0x15, 0x8c, 0xaa, 0x45, 0xa3,
	0x18, 0xec, 0xb5, 0x51, 0xd6, 0xf4, 0x1d, 0x58, 0x22, 0xc1, 0x49, 0x2a, 0x94, 0x91, 0x3d, 0xea,
	0x79, 0x5b, 0xa7, 0xea, 0x69, 0xd9, 0x44, 0x51, 0xfd, 0x4c, 0x52, 0x76, 0xdb, 0x3d, 0x34, 0xa7,
	0x4b, 0xd9, 0x78, 0x69, 0xec, 0x2d, 0x13, 0x76, 0x31, 0xd8, 0x6b, 0xf5, 0x3f, 0xb3, 0x6a, 0x59,
	0x21, 0xeb, 0x96, 0x2b, 0x86, 0xb7, 0xd0, 0xc8, 0xf7, 0x85, 0x4e, 0xa0, 0x95, 0xc7, 0x4f, 0x38,
	0x8a, 0x5b, 0x95, 0x62, 0x75, 0x92, 0x86, 0x61, 0xcb, 0x18, 0x7a, 0xd0, 0xc8, 0xdf, 0x83, 0x6c,
	0x00, 0x8d, 0x33, 0xd5, 0x06, 0x6b, 0xfe, 0xe8, 0xbb, 0x0a, 0xe6, 0x54, 0x50, 0x8e, 0x5e, 0xc1,
	0x94, 0x6b, 0x87, 0xdc, 0x1d, 0xe9, 0x73, 0xcb, 0x7b, 0xd6, 0xdb, 0xcb, 0xd1, 0x63, 0x71, 0x2b,
	0xe8, 0x5d, 0x66, 0x89, 0x19, 0x26, 0x3a, 0x27, 0xea, 0xef, 0x7d, 0xc4, 0xb5, 0xf9, 0x65, 0x09,
	0x6b, 0x63, 0xff, 0x26, 0x5b, 0x21, 0x38, 0xa1, 0x6a, 0x89, 0xfb, 0x7b, 0xa7, 0x5e, 0x66, 0xfe,
	0xf7, 0x49, 0xdd, 0xca, 0xcc, 0x52, 0x1f, 0xf3, 0xcd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x47,
	0x00, 0xe5, 0xa3, 0xe5, 0x03, 0x00, 0x00,
}
