// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/login/login.proto

package go_micro_service_register

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

type LoginStatus int32

const (
	LoginStatus_LoginFail    LoginStatus = 0
	LoginStatus_LoginSuccess LoginStatus = 1
)

var LoginStatus_name = map[int32]string{
	0: "LoginFail",
	1: "LoginSuccess",
}

var LoginStatus_value = map[string]int32{
	"LoginFail":    0,
	"LoginSuccess": 1,
}

func (x LoginStatus) String() string {
	return proto.EnumName(LoginStatus_name, int32(x))
}

func (LoginStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a27ccefc36a61526, []int{0}
}

type AuthStatus int32

const (
	AuthStatus_AuthFail    AuthStatus = 0
	AuthStatus_AuthSuccess AuthStatus = 1
)

var AuthStatus_name = map[int32]string{
	0: "AuthFail",
	1: "AuthSuccess",
}

var AuthStatus_value = map[string]int32{
	"AuthFail":    0,
	"AuthSuccess": 1,
}

func (x AuthStatus) String() string {
	return proto.EnumName(AuthStatus_name, int32(x))
}

func (AuthStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a27ccefc36a61526, []int{1}
}

type LoginRequest struct {
	Phone                string   `protobuf:"bytes,1,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Pwd                  string   `protobuf:"bytes,2,opt,name=Pwd,proto3" json:"Pwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a27ccefc36a61526, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *LoginRequest) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type LoginResponse struct {
	LoginStatus          LoginStatus `protobuf:"varint,1,opt,name=LoginStatus,proto3,enum=go.micro.service.register.LoginStatus" json:"LoginStatus,omitempty"`
	Token                string      `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a27ccefc36a61526, []int{1}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetLoginStatus() LoginStatus {
	if m != nil {
		return m.LoginStatus
	}
	return LoginStatus_LoginFail
}

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AuthRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRequest) Reset()         { *m = AuthRequest{} }
func (m *AuthRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()    {}
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a27ccefc36a61526, []int{2}
}

func (m *AuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRequest.Unmarshal(m, b)
}
func (m *AuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRequest.Marshal(b, m, deterministic)
}
func (m *AuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRequest.Merge(m, src)
}
func (m *AuthRequest) XXX_Size() int {
	return xxx_messageInfo_AuthRequest.Size(m)
}
func (m *AuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRequest proto.InternalMessageInfo

func (m *AuthRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type UserInfo struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Mobile               string   `protobuf:"bytes,3,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	AvatarUrl            string   `protobuf:"bytes,4,opt,name=Avatar_url,json=AvatarUrl,proto3" json:"Avatar_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a27ccefc36a61526, []int{3}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfo) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *UserInfo) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

type AuthResponse struct {
	AuthStatus           AuthStatus `protobuf:"varint,1,opt,name=AuthStatus,proto3,enum=go.micro.service.register.AuthStatus" json:"AuthStatus,omitempty"`
	User                 *UserInfo  `protobuf:"bytes,2,opt,name=User,proto3" json:"User,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a27ccefc36a61526, []int{4}
}

func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (m *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(m, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetAuthStatus() AuthStatus {
	if m != nil {
		return m.AuthStatus
	}
	return AuthStatus_AuthFail
}

func (m *AuthResponse) GetUser() *UserInfo {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterEnum("go.micro.service.register.LoginStatus", LoginStatus_name, LoginStatus_value)
	proto.RegisterEnum("go.micro.service.register.AuthStatus", AuthStatus_name, AuthStatus_value)
	proto.RegisterType((*LoginRequest)(nil), "go.micro.service.register.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "go.micro.service.register.LoginResponse")
	proto.RegisterType((*AuthRequest)(nil), "go.micro.service.register.AuthRequest")
	proto.RegisterType((*UserInfo)(nil), "go.micro.service.register.UserInfo")
	proto.RegisterType((*AuthResponse)(nil), "go.micro.service.register.AuthResponse")
}

func init() { proto.RegisterFile("proto/login/login.proto", fileDescriptor_a27ccefc36a61526) }

var fileDescriptor_a27ccefc36a61526 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x6d, 0xda, 0xb4, 0xb4, 0xd3, 0x0f, 0xc3, 0x20, 0x1a, 0x05, 0x41, 0x52, 0x6c, 0x4b, 0x85,
	0x08, 0x15, 0xf4, 0xdc, 0x83, 0x62, 0x41, 0xa5, 0xa4, 0xf6, 0x20, 0x08, 0x92, 0xa6, 0x6b, 0x1b,
	0x4c, 0xb3, 0x75, 0x37, 0xa9, 0xff, 0xc2, 0xdf, 0xe4, 0x4f, 0x93, 0xec, 0x6e, 0x34, 0x3d, 0xd8,
	0x5c, 0xc2, 0xbc, 0xd9, 0xf7, 0x66, 0xde, 0xcc, 0x10, 0x38, 0x5c, 0x33, 0x1a, 0xd1, 0x8b, 0x80,
	0x2e, 0xfc, 0x50, 0x7e, 0x6d, 0x91, 0xc1, 0xa3, 0x05, 0xb5, 0x57, 0xbe, 0xc7, 0xa8, 0xcd, 0x09,
	0xdb, 0xf8, 0x1e, 0xb1, 0x19, 0x59, 0xf8, 0x3c, 0x22, 0xcc, 0xba, 0x82, 0xc6, 0x7d, 0xc2, 0x74,
	0xc8, 0x47, 0x4c, 0x78, 0x84, 0xfb, 0x50, 0x1e, 0x2f, 0x69, 0x48, 0x4c, 0xed, 0x54, 0xeb, 0xd5,
	0x1c, 0x09, 0xd0, 0x80, 0xd2, 0xf8, 0x73, 0x6e, 0x16, 0x45, 0x2e, 0x09, 0x2d, 0x0a, 0x4d, 0xa5,
	0xe3, 0x6b, 0x1a, 0x72, 0x82, 0x77, 0x50, 0x17, 0x89, 0x49, 0xe4, 0x46, 0x31, 0x17, 0xf2, 0xd6,
	0xa0, 0x63, 0xff, 0xdb, 0xd9, 0xce, 0xb0, 0x9d, 0xac, 0x34, 0xb1, 0xf0, 0x44, 0xdf, 0x49, 0xa8,
	0xda, 0x49, 0x60, 0xb5, 0xa1, 0x3e, 0x8c, 0xa3, 0x65, 0xc6, 0xa7, 0x24, 0x69, 0x59, 0x12, 0x81,
	0xea, 0x94, 0x13, 0x36, 0x0a, 0xdf, 0x28, 0xb6, 0xa0, 0x38, 0x9a, 0x8b, 0xe7, 0xb2, 0x53, 0x1c,
	0xcd, 0x11, 0x41, 0x7f, 0x74, 0x57, 0x44, 0x55, 0x15, 0x31, 0x1e, 0x40, 0xe5, 0x81, 0xce, 0xfc,
	0x80, 0x98, 0x25, 0x91, 0x55, 0x08, 0x4f, 0x00, 0x86, 0x1b, 0x37, 0x72, 0xd9, 0x6b, 0xcc, 0x02,
	0x53, 0x17, 0x6f, 0x35, 0x99, 0x99, 0xb2, 0xc0, 0xfa, 0xd2, 0xa0, 0x21, 0xcd, 0xa8, 0xe1, 0x6f,
	0x00, 0x12, 0xbc, 0x35, 0xfb, 0xd9, 0x8e, 0xd9, 0xff, 0xc8, 0x4e, 0x46, 0x88, 0xd7, 0xa0, 0x27,
	0xf6, 0x85, 0xc5, 0xfa, 0xa0, 0xbd, 0xa3, 0x40, 0x3a, 0xa5, 0x23, 0x04, 0x7d, 0x7b, 0x6b, 0xf9,
	0xd8, 0x84, 0x9a, 0x80, 0xb7, 0xae, 0x1f, 0x18, 0x05, 0x34, 0xd4, 0x8d, 0x27, 0xb1, 0xe7, 0x11,
	0xce, 0x0d, 0xad, 0x7f, 0x9e, 0xf5, 0x8b, 0x0d, 0xa8, 0x26, 0x48, 0xb1, 0xf7, 0xe4, 0xa2, 0x7f,
	0xc9, 0x83, 0x6f, 0x0d, 0xca, 0x42, 0x8f, 0x2f, 0x69, 0xd0, 0xcd, 0xbb, 0xab, 0x3a, 0xd3, 0x71,
	0x2f, 0x9f, 0x28, 0x57, 0x68, 0x15, 0xf0, 0x19, 0xf4, 0xa4, 0x31, 0x76, 0x72, 0x16, 0x97, 0xd6,
	0xee, 0xe6, 0xf2, 0xd2, 0xd2, 0xb3, 0x8a, 0xf8, 0x0f, 0x2e, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xd4, 0x4b, 0x1c, 0x3b, 0x22, 0x03, 0x00, 0x00,
}
