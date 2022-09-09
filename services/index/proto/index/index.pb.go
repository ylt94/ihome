// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/index/index.proto

package go_micro_service_index

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

type AreaItem struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AreaItem) Reset()         { *m = AreaItem{} }
func (m *AreaItem) String() string { return proto.CompactTextString(m) }
func (*AreaItem) ProtoMessage()    {}
func (*AreaItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5b8486092d5d3cd, []int{0}
}

func (m *AreaItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AreaItem.Unmarshal(m, b)
}
func (m *AreaItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AreaItem.Marshal(b, m, deterministic)
}
func (m *AreaItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AreaItem.Merge(m, src)
}
func (m *AreaItem) XXX_Size() int {
	return xxx_messageInfo_AreaItem.Size(m)
}
func (m *AreaItem) XXX_DiscardUnknown() {
	xxx_messageInfo_AreaItem.DiscardUnknown(m)
}

var xxx_messageInfo_AreaItem proto.InternalMessageInfo

func (m *AreaItem) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AreaItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AreaRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AreaRequest) Reset()         { *m = AreaRequest{} }
func (m *AreaRequest) String() string { return proto.CompactTextString(m) }
func (*AreaRequest) ProtoMessage()    {}
func (*AreaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5b8486092d5d3cd, []int{1}
}

func (m *AreaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AreaRequest.Unmarshal(m, b)
}
func (m *AreaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AreaRequest.Marshal(b, m, deterministic)
}
func (m *AreaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AreaRequest.Merge(m, src)
}
func (m *AreaRequest) XXX_Size() int {
	return xxx_messageInfo_AreaRequest.Size(m)
}
func (m *AreaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AreaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AreaRequest proto.InternalMessageInfo

type AreaResponse struct {
	Areas                []*AreaItem `protobuf:"bytes,1,rep,name=areas,proto3" json:"areas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AreaResponse) Reset()         { *m = AreaResponse{} }
func (m *AreaResponse) String() string { return proto.CompactTextString(m) }
func (*AreaResponse) ProtoMessage()    {}
func (*AreaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5b8486092d5d3cd, []int{2}
}

func (m *AreaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AreaResponse.Unmarshal(m, b)
}
func (m *AreaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AreaResponse.Marshal(b, m, deterministic)
}
func (m *AreaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AreaResponse.Merge(m, src)
}
func (m *AreaResponse) XXX_Size() int {
	return xxx_messageInfo_AreaResponse.Size(m)
}
func (m *AreaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AreaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AreaResponse proto.InternalMessageInfo

func (m *AreaResponse) GetAreas() []*AreaItem {
	if m != nil {
		return m.Areas
	}
	return nil
}

type House struct {
	HouseId              uint32   `protobuf:"varint,1,opt,name=house_id,json=houseId,proto3" json:"house_id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Price                string   `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	AreaName             string   `protobuf:"bytes,4,opt,name=area_name,json=areaName,proto3" json:"area_name,omitempty"`
	ImageUrl             string   `protobuf:"bytes,5,opt,name=image_url,json=imageUrl,proto3" json:"image_url"`
	RoomCount            uint32   `protobuf:"varint,6,opt,name=room_count,json=roomCount,proto3" json:"room_count,omitempty"`
	OrderCount           uint32   `protobuf:"varint,7,opt,name=order_count,json=orderCount,proto3" json:"order_count,omitempty"`
	Address              string   `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
	UserAvatar           string   `protobuf:"bytes,9,opt,name=user_avatar,json=userAvatar,proto3" json:"user_avatar,omitempty"`
	Ctime                string   `protobuf:"bytes,10,opt,name=ctime,proto3" json:"ctime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *House) Reset()         { *m = House{} }
func (m *House) String() string { return proto.CompactTextString(m) }
func (*House) ProtoMessage()    {}
func (*House) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5b8486092d5d3cd, []int{3}
}

func (m *House) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_House.Unmarshal(m, b)
}
func (m *House) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_House.Marshal(b, m, deterministic)
}
func (m *House) XXX_Merge(src proto.Message) {
	xxx_messageInfo_House.Merge(m, src)
}
func (m *House) XXX_Size() int {
	return xxx_messageInfo_House.Size(m)
}
func (m *House) XXX_DiscardUnknown() {
	xxx_messageInfo_House.DiscardUnknown(m)
}

var xxx_messageInfo_House proto.InternalMessageInfo

func (m *House) GetHouseId() uint32 {
	if m != nil {
		return m.HouseId
	}
	return 0
}

func (m *House) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *House) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *House) GetAreaName() string {
	if m != nil {
		return m.AreaName
	}
	return ""
}

func (m *House) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func (m *House) GetRoomCount() uint32 {
	if m != nil {
		return m.RoomCount
	}
	return 0
}

func (m *House) GetOrderCount() uint32 {
	if m != nil {
		return m.OrderCount
	}
	return 0
}

func (m *House) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *House) GetUserAvatar() string {
	if m != nil {
		return m.UserAvatar
	}
	return ""
}

func (m *House) GetCtime() string {
	if m != nil {
		return m.Ctime
	}
	return ""
}

type BannerRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BannerRequest) Reset()         { *m = BannerRequest{} }
func (m *BannerRequest) String() string { return proto.CompactTextString(m) }
func (*BannerRequest) ProtoMessage()    {}
func (*BannerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5b8486092d5d3cd, []int{4}
}

func (m *BannerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BannerRequest.Unmarshal(m, b)
}
func (m *BannerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BannerRequest.Marshal(b, m, deterministic)
}
func (m *BannerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BannerRequest.Merge(m, src)
}
func (m *BannerRequest) XXX_Size() int {
	return xxx_messageInfo_BannerRequest.Size(m)
}
func (m *BannerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BannerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BannerRequest proto.InternalMessageInfo

type BannerResponse struct {
	Houses               []*House `protobuf:"bytes,1,rep,name=houses,proto3" json:"houses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BannerResponse) Reset()         { *m = BannerResponse{} }
func (m *BannerResponse) String() string { return proto.CompactTextString(m) }
func (*BannerResponse) ProtoMessage()    {}
func (*BannerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5b8486092d5d3cd, []int{5}
}

func (m *BannerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BannerResponse.Unmarshal(m, b)
}
func (m *BannerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BannerResponse.Marshal(b, m, deterministic)
}
func (m *BannerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BannerResponse.Merge(m, src)
}
func (m *BannerResponse) XXX_Size() int {
	return xxx_messageInfo_BannerResponse.Size(m)
}
func (m *BannerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BannerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BannerResponse proto.InternalMessageInfo

func (m *BannerResponse) GetHouses() []*House {
	if m != nil {
		return m.Houses
	}
	return nil
}

func init() {
	proto.RegisterType((*AreaItem)(nil), "go.micro.service.index.AreaItem")
	proto.RegisterType((*AreaRequest)(nil), "go.micro.service.index.AreaRequest")
	proto.RegisterType((*AreaResponse)(nil), "go.micro.service.index.AreaResponse")
	proto.RegisterType((*House)(nil), "go.micro.service.index.House")
	proto.RegisterType((*BannerRequest)(nil), "go.micro.service.index.BannerRequest")
	proto.RegisterType((*BannerResponse)(nil), "go.micro.service.index.BannerResponse")
}

func init() { proto.RegisterFile("proto/index/index.proto", fileDescriptor_b5b8486092d5d3cd) }

var fileDescriptor_b5b8486092d5d3cd = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xdf, 0x0a, 0xd3, 0x30,
	0x14, 0xc6, 0x6d, 0xd7, 0x76, 0xed, 0x99, 0x9b, 0x10, 0x44, 0xe3, 0x64, 0x58, 0xea, 0x1f, 0x76,
	0x55, 0x61, 0xa2, 0xf7, 0x53, 0x50, 0x77, 0xe3, 0x45, 0xc5, 0x0b, 0xaf, 0x4a, 0x6c, 0x0f, 0x33,
	0xb0, 0x36, 0x33, 0x49, 0x87, 0xaf, 0xe2, 0x93, 0xf8, 0x7a, 0x92, 0x93, 0xf6, 0x42, 0x70, 0xf3,
	0xa6, 0xe4, 0xfc, 0xbe, 0x2f, 0xc9, 0x39, 0x5f, 0x0a, 0x0f, 0xcf, 0x5a, 0x59, 0xf5, 0x52, 0xf6,
	0x2d, 0xfe, 0xf4, 0xdf, 0x92, 0x08, 0x7b, 0x70, 0x54, 0x65, 0x27, 0x1b, 0xad, 0x4a, 0x83, 0xfa,
	0x22, 0x1b, 0x2c, 0x49, 0x2d, 0x4a, 0x48, 0xf7, 0x1a, 0xc5, 0xc1, 0x62, 0xc7, 0x56, 0x10, 0xca,
	0x96, 0x07, 0x79, 0xb0, 0x8d, 0xaa, 0x50, 0xb6, 0x8c, 0x41, 0xd4, 0x8b, 0x0e, 0x79, 0x98, 0x07,
	0xdb, 0xac, 0xa2, 0x75, 0xb1, 0x84, 0x85, 0xf3, 0x57, 0xf8, 0x63, 0x40, 0x63, 0x8b, 0xf7, 0x70,
	0xd7, 0x97, 0xe6, 0xac, 0x7a, 0x83, 0xec, 0x0d, 0xc4, 0x42, 0xa3, 0x30, 0x3c, 0xc8, 0x67, 0xdb,
	0xc5, 0x2e, 0x2f, 0xff, 0x7d, 0x6d, 0x39, 0xdd, 0x59, 0x79, 0x7b, 0xf1, 0x2b, 0x84, 0xf8, 0xa3,
	0x1a, 0x0c, 0xb2, 0x47, 0x90, 0x7e, 0x77, 0x8b, 0x7a, 0x6c, 0x65, 0x59, 0xcd, 0xa9, 0x3e, 0xb4,
	0xec, 0x3e, 0xc4, 0x56, 0xda, 0xd3, 0xd4, 0x90, 0x2f, 0x1c, 0x3d, 0x6b, 0xd9, 0x20, 0x9f, 0x79,
	0x4a, 0x05, 0x7b, 0x0c, 0x99, 0x3b, 0xb9, 0xa6, 0x01, 0x22, 0x52, 0x52, 0x07, 0x3e, 0x89, 0x8e,
	0x44, 0xd9, 0x89, 0x23, 0xd6, 0x83, 0x3e, 0xf1, 0xd8, 0x8b, 0x04, 0xbe, 0xe8, 0x13, 0xdb, 0x00,
	0x68, 0xa5, 0xba, 0xba, 0x51, 0x43, 0x6f, 0x79, 0x42, 0x2d, 0x64, 0x8e, 0xbc, 0x73, 0x80, 0x3d,
	0x81, 0x85, 0xd2, 0x2d, 0xea, 0x51, 0x9f, 0x93, 0x0e, 0x84, 0xbc, 0x81, 0xc3, 0x5c, 0xb4, 0xad,
	0x46, 0x63, 0x78, 0x4a, 0x47, 0x4f, 0xa5, 0xdb, 0x3a, 0x18, 0xd4, 0xb5, 0xb8, 0x08, 0x2b, 0x34,
	0xcf, 0x48, 0x05, 0x87, 0xf6, 0x44, 0xdc, 0x28, 0x8d, 0x95, 0x1d, 0x72, 0xf0, 0xa3, 0x50, 0x51,
	0xdc, 0x83, 0xe5, 0x5b, 0xd1, 0xf7, 0xa8, 0xa7, 0xd0, 0x3f, 0xc0, 0x6a, 0x02, 0x63, 0xec, 0xaf,
	0x21, 0xa1, 0x90, 0xa6, 0xdc, 0x37, 0xd7, 0x72, 0xa7, 0x8c, 0xab, 0xd1, 0xbc, 0xfb, 0x1d, 0x40,
	0x7c, 0x70, 0x9c, 0x7d, 0x86, 0xc8, 0x3d, 0x09, 0x7b, 0x7a, 0xeb, 0xc1, 0xc6, 0xfb, 0xd7, 0xcf,
	0x6e, 0x9b, 0x7c, 0x4f, 0xc5, 0x1d, 0xf6, 0x15, 0x12, 0xdf, 0x27, 0x7b, 0x7e, 0x6d, 0xc7, 0x5f,
	0x83, 0xad, 0x5f, 0xfc, 0xcf, 0x36, 0x1d, 0xfd, 0x2d, 0xa1, 0xbf, 0xfa, 0xd5, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x99, 0xea, 0x95, 0xb5, 0xf0, 0x02, 0x00, 0x00,
}
