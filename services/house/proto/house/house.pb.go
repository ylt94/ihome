// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/house/house.proto

package go_micro_service_house

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

type ListRequest struct {
	Aid                  uint32   `protobuf:"varint,1,opt,name=aid,proto3" json:"aid,omitempty"`
	StartDate            string   `protobuf:"bytes,2,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate              string   `protobuf:"bytes,3,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	Page                 uint32   `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{0}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetAid() uint32 {
	if m != nil {
		return m.Aid
	}
	return 0
}

func (m *ListRequest) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *ListRequest) GetEndDate() string {
	if m != nil {
		return m.EndDate
	}
	return ""
}

func (m *ListRequest) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

type ListItem struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	AreaName             string   `protobuf:"bytes,2,opt,name=area_name,json=areaName,proto3" json:"area_name,omitempty"`
	Ctime                string   `protobuf:"bytes,3,opt,name=ctime,proto3" json:"ctime,omitempty"`
	HouseId              uint32   `protobuf:"varint,4,opt,name=house_id,json=houseId,proto3" json:"house_id,omitempty"`
	ImageUrl             string   `protobuf:"bytes,5,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	OrderCount           uint32   `protobuf:"varint,6,opt,name=order_count,json=orderCount,proto3" json:"order_count,omitempty"`
	Price                uint32   `protobuf:"varint,7,opt,name=price,proto3" json:"price,omitempty"`
	RoomCount            uint32   `protobuf:"varint,8,opt,name=room_count,json=roomCount,proto3" json:"room_count,omitempty"`
	Title                string   `protobuf:"bytes,9,opt,name=title,proto3" json:"title,omitempty"`
	UserAvatar           string   `protobuf:"bytes,10,opt,name=user_avatar,json=userAvatar,proto3" json:"user_avatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListItem) Reset()         { *m = ListItem{} }
func (m *ListItem) String() string { return proto.CompactTextString(m) }
func (*ListItem) ProtoMessage()    {}
func (*ListItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{1}
}

func (m *ListItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListItem.Unmarshal(m, b)
}
func (m *ListItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListItem.Marshal(b, m, deterministic)
}
func (m *ListItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListItem.Merge(m, src)
}
func (m *ListItem) XXX_Size() int {
	return xxx_messageInfo_ListItem.Size(m)
}
func (m *ListItem) XXX_DiscardUnknown() {
	xxx_messageInfo_ListItem.DiscardUnknown(m)
}

var xxx_messageInfo_ListItem proto.InternalMessageInfo

func (m *ListItem) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ListItem) GetAreaName() string {
	if m != nil {
		return m.AreaName
	}
	return ""
}

func (m *ListItem) GetCtime() string {
	if m != nil {
		return m.Ctime
	}
	return ""
}

func (m *ListItem) GetHouseId() uint32 {
	if m != nil {
		return m.HouseId
	}
	return 0
}

func (m *ListItem) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func (m *ListItem) GetOrderCount() uint32 {
	if m != nil {
		return m.OrderCount
	}
	return 0
}

func (m *ListItem) GetPrice() uint32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *ListItem) GetRoomCount() uint32 {
	if m != nil {
		return m.RoomCount
	}
	return 0
}

func (m *ListItem) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ListItem) GetUserAvatar() string {
	if m != nil {
		return m.UserAvatar
	}
	return ""
}

type ListResponse struct {
	CurrentPage          uint32      `protobuf:"varint,1,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty"`
	Houses               []*ListItem `protobuf:"bytes,2,rep,name=houses,proto3" json:"houses,omitempty"`
	TotalPage            uint32      `protobuf:"varint,3,opt,name=total_page,json=totalPage,proto3" json:"total_page,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{2}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetCurrentPage() uint32 {
	if m != nil {
		return m.CurrentPage
	}
	return 0
}

func (m *ListResponse) GetHouses() []*ListItem {
	if m != nil {
		return m.Houses
	}
	return nil
}

func (m *ListResponse) GetTotalPage() uint32 {
	if m != nil {
		return m.TotalPage
	}
	return 0
}

type CreateRequest struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Price                uint32   `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	AreaId               uint32   `protobuf:"varint,3,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	Address              string   `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	RoomCount            uint32   `protobuf:"varint,5,opt,name=room_count,json=roomCount,proto3" json:"room_count,omitempty"`
	Acreage              uint32   `protobuf:"varint,6,opt,name=acreage,proto3" json:"acreage,omitempty"`
	Unit                 string   `protobuf:"bytes,7,opt,name=unit,proto3" json:"unit,omitempty"`
	Capacity             uint32   `protobuf:"varint,8,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Beds                 string   `protobuf:"bytes,9,opt,name=beds,proto3" json:"beds,omitempty"`
	Deposit              uint32   `protobuf:"varint,10,opt,name=deposit,proto3" json:"deposit,omitempty"`
	MinDays              uint32   `protobuf:"varint,11,opt,name=min_days,json=minDays,proto3" json:"min_days,omitempty"`
	MaxDays              uint32   `protobuf:"varint,12,opt,name=max_days,json=maxDays,proto3" json:"max_days,omitempty"`
	Facility             []uint32 `protobuf:"varint,13,rep,packed,name=facility,proto3" json:"facility,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{3}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreateRequest) GetPrice() uint32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *CreateRequest) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *CreateRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *CreateRequest) GetRoomCount() uint32 {
	if m != nil {
		return m.RoomCount
	}
	return 0
}

func (m *CreateRequest) GetAcreage() uint32 {
	if m != nil {
		return m.Acreage
	}
	return 0
}

func (m *CreateRequest) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *CreateRequest) GetCapacity() uint32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *CreateRequest) GetBeds() string {
	if m != nil {
		return m.Beds
	}
	return ""
}

func (m *CreateRequest) GetDeposit() uint32 {
	if m != nil {
		return m.Deposit
	}
	return 0
}

func (m *CreateRequest) GetMinDays() uint32 {
	if m != nil {
		return m.MinDays
	}
	return 0
}

func (m *CreateRequest) GetMaxDays() uint32 {
	if m != nil {
		return m.MaxDays
	}
	return 0
}

func (m *CreateRequest) GetFacility() []uint32 {
	if m != nil {
		return m.Facility
	}
	return nil
}

type CreateResponse struct {
	HouseId              uint32   `protobuf:"varint,1,opt,name=house_id,json=houseId,proto3" json:"house_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{4}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetHouseId() uint32 {
	if m != nil {
		return m.HouseId
	}
	return 0
}

type UploadImageRequest struct {
	HouseId              uint32   `protobuf:"varint,1,opt,name=house_id,json=houseId,proto3" json:"house_id,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadImageRequest) Reset()         { *m = UploadImageRequest{} }
func (m *UploadImageRequest) String() string { return proto.CompactTextString(m) }
func (*UploadImageRequest) ProtoMessage()    {}
func (*UploadImageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{5}
}

func (m *UploadImageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadImageRequest.Unmarshal(m, b)
}
func (m *UploadImageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadImageRequest.Marshal(b, m, deterministic)
}
func (m *UploadImageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadImageRequest.Merge(m, src)
}
func (m *UploadImageRequest) XXX_Size() int {
	return xxx_messageInfo_UploadImageRequest.Size(m)
}
func (m *UploadImageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadImageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadImageRequest proto.InternalMessageInfo

func (m *UploadImageRequest) GetHouseId() uint32 {
	if m != nil {
		return m.HouseId
	}
	return 0
}

func (m *UploadImageRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type UploadImageResponse struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadImageResponse) Reset()         { *m = UploadImageResponse{} }
func (m *UploadImageResponse) String() string { return proto.CompactTextString(m) }
func (*UploadImageResponse) ProtoMessage()    {}
func (*UploadImageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{6}
}

func (m *UploadImageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadImageResponse.Unmarshal(m, b)
}
func (m *UploadImageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadImageResponse.Marshal(b, m, deterministic)
}
func (m *UploadImageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadImageResponse.Merge(m, src)
}
func (m *UploadImageResponse) XXX_Size() int {
	return xxx_messageInfo_UploadImageResponse.Size(m)
}
func (m *UploadImageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadImageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadImageResponse proto.InternalMessageInfo

func (m *UploadImageResponse) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type Comment struct {
	Comment              string   `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
	Ctime                string   `protobuf:"bytes,2,opt,name=ctime,proto3" json:"ctime,omitempty"`
	UserName             string   `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Comment) Reset()         { *m = Comment{} }
func (m *Comment) String() string { return proto.CompactTextString(m) }
func (*Comment) ProtoMessage()    {}
func (*Comment) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{7}
}

func (m *Comment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Comment.Unmarshal(m, b)
}
func (m *Comment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Comment.Marshal(b, m, deterministic)
}
func (m *Comment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Comment.Merge(m, src)
}
func (m *Comment) XXX_Size() int {
	return xxx_messageInfo_Comment.Size(m)
}
func (m *Comment) XXX_DiscardUnknown() {
	xxx_messageInfo_Comment.DiscardUnknown(m)
}

var xxx_messageInfo_Comment proto.InternalMessageInfo

func (m *Comment) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *Comment) GetCtime() string {
	if m != nil {
		return m.Ctime
	}
	return ""
}

func (m *Comment) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type DetailRequest struct {
	HouseId              uint32   `protobuf:"varint,1,opt,name=house_id,json=houseId,proto3" json:"house_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetailRequest) Reset()         { *m = DetailRequest{} }
func (m *DetailRequest) String() string { return proto.CompactTextString(m) }
func (*DetailRequest) ProtoMessage()    {}
func (*DetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{8}
}

func (m *DetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailRequest.Unmarshal(m, b)
}
func (m *DetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailRequest.Marshal(b, m, deterministic)
}
func (m *DetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailRequest.Merge(m, src)
}
func (m *DetailRequest) XXX_Size() int {
	return xxx_messageInfo_DetailRequest.Size(m)
}
func (m *DetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DetailRequest proto.InternalMessageInfo

func (m *DetailRequest) GetHouseId() uint32 {
	if m != nil {
		return m.HouseId
	}
	return 0
}

type DetailResponse struct {
	Acreage              uint32     `protobuf:"varint,1,opt,name=acreage,proto3" json:"acreage,omitempty"`
	Address              string     `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Beds                 string     `protobuf:"bytes,3,opt,name=beds,proto3" json:"beds,omitempty"`
	Capacity             uint32     `protobuf:"varint,4,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Comments             []*Comment `protobuf:"bytes,5,rep,name=comments,proto3" json:"comments,omitempty"`
	Deposit              uint32     `protobuf:"varint,6,opt,name=deposit,proto3" json:"deposit,omitempty"`
	Facilities           []string   `protobuf:"bytes,7,rep,name=facilities,proto3" json:"facilities,omitempty"`
	Hid                  uint32     `protobuf:"varint,8,opt,name=hid,proto3" json:"hid,omitempty"`
	ImgUrls              []string   `protobuf:"bytes,9,rep,name=img_urls,json=imgUrls,proto3" json:"img_urls,omitempty"`
	MinDays              uint32     `protobuf:"varint,10,opt,name=min_days,json=minDays,proto3" json:"min_days,omitempty"`
	MaxDays              uint32     `protobuf:"varint,11,opt,name=max_days,json=maxDays,proto3" json:"max_days,omitempty"`
	Price                uint32     `protobuf:"varint,12,opt,name=price,proto3" json:"price,omitempty"`
	RoomCount            uint32     `protobuf:"varint,13,opt,name=room_count,json=roomCount,proto3" json:"room_count,omitempty"`
	Title                string     `protobuf:"bytes,14,opt,name=title,proto3" json:"title,omitempty"`
	Unit                 string     `protobuf:"bytes,15,opt,name=unit,proto3" json:"unit,omitempty"`
	UserAvatar           string     `protobuf:"bytes,16,opt,name=user_avatar,json=userAvatar,proto3" json:"user_avatar,omitempty"`
	UserId               uint32     `protobuf:"varint,17,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName             string     `protobuf:"bytes,18,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DetailResponse) Reset()         { *m = DetailResponse{} }
func (m *DetailResponse) String() string { return proto.CompactTextString(m) }
func (*DetailResponse) ProtoMessage()    {}
func (*DetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{9}
}

func (m *DetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailResponse.Unmarshal(m, b)
}
func (m *DetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailResponse.Marshal(b, m, deterministic)
}
func (m *DetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailResponse.Merge(m, src)
}
func (m *DetailResponse) XXX_Size() int {
	return xxx_messageInfo_DetailResponse.Size(m)
}
func (m *DetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DetailResponse proto.InternalMessageInfo

func (m *DetailResponse) GetAcreage() uint32 {
	if m != nil {
		return m.Acreage
	}
	return 0
}

func (m *DetailResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *DetailResponse) GetBeds() string {
	if m != nil {
		return m.Beds
	}
	return ""
}

func (m *DetailResponse) GetCapacity() uint32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *DetailResponse) GetComments() []*Comment {
	if m != nil {
		return m.Comments
	}
	return nil
}

func (m *DetailResponse) GetDeposit() uint32 {
	if m != nil {
		return m.Deposit
	}
	return 0
}

func (m *DetailResponse) GetFacilities() []string {
	if m != nil {
		return m.Facilities
	}
	return nil
}

func (m *DetailResponse) GetHid() uint32 {
	if m != nil {
		return m.Hid
	}
	return 0
}

func (m *DetailResponse) GetImgUrls() []string {
	if m != nil {
		return m.ImgUrls
	}
	return nil
}

func (m *DetailResponse) GetMinDays() uint32 {
	if m != nil {
		return m.MinDays
	}
	return 0
}

func (m *DetailResponse) GetMaxDays() uint32 {
	if m != nil {
		return m.MaxDays
	}
	return 0
}

func (m *DetailResponse) GetPrice() uint32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *DetailResponse) GetRoomCount() uint32 {
	if m != nil {
		return m.RoomCount
	}
	return 0
}

func (m *DetailResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *DetailResponse) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *DetailResponse) GetUserAvatar() string {
	if m != nil {
		return m.UserAvatar
	}
	return ""
}

func (m *DetailResponse) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *DetailResponse) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type DataRequest struct {
	HouseIds             []uint32 `protobuf:"varint,1,rep,packed,name=house_ids,json=houseIds,proto3" json:"house_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataRequest) Reset()         { *m = DataRequest{} }
func (m *DataRequest) String() string { return proto.CompactTextString(m) }
func (*DataRequest) ProtoMessage()    {}
func (*DataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{10}
}

func (m *DataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataRequest.Unmarshal(m, b)
}
func (m *DataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataRequest.Marshal(b, m, deterministic)
}
func (m *DataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataRequest.Merge(m, src)
}
func (m *DataRequest) XXX_Size() int {
	return xxx_messageInfo_DataRequest.Size(m)
}
func (m *DataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataRequest proto.InternalMessageInfo

func (m *DataRequest) GetHouseIds() []uint32 {
	if m != nil {
		return m.HouseIds
	}
	return nil
}

type DataItem struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               uint32   `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AreaId               uint32   `protobuf:"varint,3,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	Title                string   `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Address              string   `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	RoomCount            uint32   `protobuf:"varint,6,opt,name=room_count,json=roomCount,proto3" json:"room_count,omitempty"`
	Acreage              uint32   `protobuf:"varint,7,opt,name=acreage,proto3" json:"acreage,omitempty"`
	Price                uint32   `protobuf:"varint,8,opt,name=price,proto3" json:"price,omitempty"`
	Unit                 string   `protobuf:"bytes,9,opt,name=unit,proto3" json:"unit,omitempty"`
	Capacity             uint32   `protobuf:"varint,10,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Beds                 string   `protobuf:"bytes,11,opt,name=beds,proto3" json:"beds,omitempty"`
	Deposit              uint32   `protobuf:"varint,12,opt,name=deposit,proto3" json:"deposit,omitempty"`
	MinDays              uint32   `protobuf:"varint,13,opt,name=min_days,json=minDays,proto3" json:"min_days,omitempty"`
	MaxDays              uint32   `protobuf:"varint,14,opt,name=max_days,json=maxDays,proto3" json:"max_days,omitempty"`
	OrderCount           uint32   `protobuf:"varint,15,opt,name=order_count,json=orderCount,proto3" json:"order_count,omitempty"`
	IndexImageUrl        string   `protobuf:"bytes,16,opt,name=index_image_url,json=indexImageUrl,proto3" json:"index_image_url,omitempty"`
	CreatedAt            string   `protobuf:"bytes,17,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,18,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            string   `protobuf:"bytes,19,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataItem) Reset()         { *m = DataItem{} }
func (m *DataItem) String() string { return proto.CompactTextString(m) }
func (*DataItem) ProtoMessage()    {}
func (*DataItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{11}
}

func (m *DataItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataItem.Unmarshal(m, b)
}
func (m *DataItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataItem.Marshal(b, m, deterministic)
}
func (m *DataItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataItem.Merge(m, src)
}
func (m *DataItem) XXX_Size() int {
	return xxx_messageInfo_DataItem.Size(m)
}
func (m *DataItem) XXX_DiscardUnknown() {
	xxx_messageInfo_DataItem.DiscardUnknown(m)
}

var xxx_messageInfo_DataItem proto.InternalMessageInfo

func (m *DataItem) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DataItem) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *DataItem) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *DataItem) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *DataItem) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *DataItem) GetRoomCount() uint32 {
	if m != nil {
		return m.RoomCount
	}
	return 0
}

func (m *DataItem) GetAcreage() uint32 {
	if m != nil {
		return m.Acreage
	}
	return 0
}

func (m *DataItem) GetPrice() uint32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *DataItem) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *DataItem) GetCapacity() uint32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *DataItem) GetBeds() string {
	if m != nil {
		return m.Beds
	}
	return ""
}

func (m *DataItem) GetDeposit() uint32 {
	if m != nil {
		return m.Deposit
	}
	return 0
}

func (m *DataItem) GetMinDays() uint32 {
	if m != nil {
		return m.MinDays
	}
	return 0
}

func (m *DataItem) GetMaxDays() uint32 {
	if m != nil {
		return m.MaxDays
	}
	return 0
}

func (m *DataItem) GetOrderCount() uint32 {
	if m != nil {
		return m.OrderCount
	}
	return 0
}

func (m *DataItem) GetIndexImageUrl() string {
	if m != nil {
		return m.IndexImageUrl
	}
	return ""
}

func (m *DataItem) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *DataItem) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *DataItem) GetDeletedAt() string {
	if m != nil {
		return m.DeletedAt
	}
	return ""
}

type DataResponse struct {
	Houses               []*DataItem `protobuf:"bytes,1,rep,name=houses,proto3" json:"houses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DataResponse) Reset()         { *m = DataResponse{} }
func (m *DataResponse) String() string { return proto.CompactTextString(m) }
func (*DataResponse) ProtoMessage()    {}
func (*DataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{12}
}

func (m *DataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataResponse.Unmarshal(m, b)
}
func (m *DataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataResponse.Marshal(b, m, deterministic)
}
func (m *DataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataResponse.Merge(m, src)
}
func (m *DataResponse) XXX_Size() int {
	return xxx_messageInfo_DataResponse.Size(m)
}
func (m *DataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DataResponse proto.InternalMessageInfo

func (m *DataResponse) GetHouses() []*DataItem {
	if m != nil {
		return m.Houses
	}
	return nil
}

func init() {
	proto.RegisterType((*ListRequest)(nil), "go.micro.service.house.ListRequest")
	proto.RegisterType((*ListItem)(nil), "go.micro.service.house.ListItem")
	proto.RegisterType((*ListResponse)(nil), "go.micro.service.house.ListResponse")
	proto.RegisterType((*CreateRequest)(nil), "go.micro.service.house.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "go.micro.service.house.CreateResponse")
	proto.RegisterType((*UploadImageRequest)(nil), "go.micro.service.house.UploadImageRequest")
	proto.RegisterType((*UploadImageResponse)(nil), "go.micro.service.house.UploadImageResponse")
	proto.RegisterType((*Comment)(nil), "go.micro.service.house.Comment")
	proto.RegisterType((*DetailRequest)(nil), "go.micro.service.house.DetailRequest")
	proto.RegisterType((*DetailResponse)(nil), "go.micro.service.house.DetailResponse")
	proto.RegisterType((*DataRequest)(nil), "go.micro.service.house.DataRequest")
	proto.RegisterType((*DataItem)(nil), "go.micro.service.house.DataItem")
	proto.RegisterType((*DataResponse)(nil), "go.micro.service.house.DataResponse")
}

func init() { proto.RegisterFile("proto/house/house.proto", fileDescriptor_fc2bc500116e1060) }

var fileDescriptor_fc2bc500116e1060 = []byte{
	// 1025 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0x4b, 0x8f, 0xe3, 0x44,
	0x10, 0x9e, 0x38, 0x2f, 0xbb, 0xf2, 0x98, 0xa5, 0x17, 0xb1, 0x26, 0xa3, 0x65, 0x83, 0x81, 0x65,
	0xb4, 0x2b, 0x05, 0x69, 0xb9, 0x20, 0x71, 0x1a, 0x66, 0x0e, 0x13, 0x09, 0x21, 0x14, 0x34, 0x48,
	0x1c, 0x90, 0xd5, 0xeb, 0x6e, 0x32, 0x2d, 0xc5, 0x76, 0x70, 0x77, 0x56, 0x33, 0x77, 0xb8, 0xf0,
	0x1b, 0xf8, 0x5f, 0x1c, 0xf8, 0x33, 0xa8, 0xfa, 0xe1, 0xb4, 0x33, 0xeb, 0x0c, 0x97, 0xa8, 0xeb,
	0xd1, 0x5d, 0xae, 0xaa, 0xaf, 0xbe, 0x0a, 0x3c, 0xdb, 0x56, 0xa5, 0x2a, 0xbf, 0xba, 0x2d, 0x77,
	0x92, 0x9b, 0xdf, 0x85, 0xd6, 0x90, 0x8f, 0xd6, 0xe5, 0x22, 0x17, 0x59, 0x55, 0x2e, 0x24, 0xaf,
	0xde, 0x89, 0x8c, 0x2f, 0xb4, 0x35, 0xc9, 0x61, 0xf4, 0xbd, 0x90, 0x6a, 0xc5, 0x7f, 0xdf, 0x71,
	0xa9, 0xc8, 0x13, 0xe8, 0x52, 0xc1, 0xe2, 0xce, 0xbc, 0x73, 0x3e, 0x59, 0xe1, 0x91, 0x3c, 0x07,
	0x90, 0x8a, 0x56, 0x2a, 0x65, 0x54, 0xf1, 0x38, 0x98, 0x77, 0xce, 0xa3, 0x55, 0xa4, 0x35, 0x57,
	0x54, 0x71, 0xf2, 0x31, 0x84, 0xbc, 0x60, 0xc6, 0xd8, 0xd5, 0xc6, 0x21, 0x2f, 0x98, 0x36, 0x11,
	0xe8, 0x6d, 0xe9, 0x9a, 0xc7, 0x3d, 0xfd, 0x98, 0x3e, 0x27, 0x7f, 0x07, 0x10, 0x62, 0xbc, 0xa5,
	0xe2, 0x39, 0x89, 0x61, 0x48, 0x19, 0xab, 0xb8, 0x94, 0x3a, 0x60, 0xb4, 0x72, 0x22, 0x39, 0x83,
	0x88, 0x56, 0x9c, 0xa6, 0x05, 0xcd, 0x5d, 0xcc, 0x10, 0x15, 0x3f, 0xd0, 0x9c, 0x93, 0x0f, 0xa1,
	0x9f, 0x29, 0x91, 0xbb, 0x78, 0x46, 0xc0, 0x0f, 0xd1, 0x19, 0xa5, 0x82, 0xd9, 0x88, 0x43, 0x2d,
	0x2f, 0x19, 0xbe, 0x26, 0x72, 0xba, 0xe6, 0xe9, 0xae, 0xda, 0xc4, 0x7d, 0xf3, 0x9a, 0x56, 0xdc,
	0x54, 0x1b, 0xf2, 0x02, 0x46, 0x65, 0xc5, 0x78, 0x95, 0x66, 0xe5, 0xae, 0x50, 0xf1, 0x40, 0x5f,
	0x05, 0xad, 0xba, 0x44, 0x0d, 0x86, 0xdb, 0x56, 0x22, 0xe3, 0xf1, 0x50, 0x9b, 0x8c, 0x80, 0x65,
	0xa9, 0xca, 0x32, 0xb7, 0xb7, 0x42, 0x6d, 0x8a, 0x50, 0x53, 0x5f, 0x52, 0x42, 0x6d, 0x78, 0x1c,
	0x99, 0x6f, 0xd4, 0x02, 0xc6, 0xda, 0x49, 0x5e, 0xa5, 0xf4, 0x1d, 0x55, 0xb4, 0x8a, 0x41, 0xdb,
	0x00, 0x55, 0x17, 0x5a, 0x93, 0xfc, 0xd5, 0x81, 0xb1, 0x69, 0x87, 0xdc, 0x96, 0x85, 0xe4, 0xe4,
	0x53, 0x18, 0x67, 0xbb, 0xaa, 0xe2, 0x85, 0x4a, 0x75, 0x2d, 0x4d, 0x63, 0x46, 0x56, 0xf7, 0x23,
	0x5d, 0x73, 0xf2, 0x0d, 0x0c, 0x74, 0xa2, 0x32, 0x0e, 0xe6, 0xdd, 0xf3, 0xd1, 0x9b, 0xf9, 0xe2,
	0xfd, 0xad, 0x5e, 0xb8, 0xba, 0xaf, 0xac, 0x3f, 0xe6, 0xa0, 0x4a, 0x45, 0x37, 0xe6, 0xe9, 0xae,
	0xc9, 0x41, 0x6b, 0xf0, 0xe1, 0xe4, 0x9f, 0x00, 0x26, 0x97, 0x15, 0xa7, 0x8a, 0x3b, 0x74, 0xd4,
	0x59, 0x75, 0xfc, 0xac, 0xea, 0x02, 0x05, 0x7e, 0x81, 0x9e, 0xc1, 0x50, 0xb7, 0x50, 0x30, 0xfb,
	0xf2, 0x00, 0xc5, 0x25, 0xf3, 0xbb, 0xde, 0x6b, 0x76, 0xbd, 0x59, 0xd3, 0xfe, 0x61, 0x4d, 0xf1,
	0x62, 0x56, 0x71, 0xfc, 0x56, 0xd3, 0x25, 0x27, 0x22, 0xd2, 0x76, 0x85, 0x50, 0xba, 0x43, 0xd1,
	0x4a, 0x9f, 0xc9, 0x0c, 0xc2, 0x8c, 0x6e, 0x69, 0x26, 0xd4, 0xbd, 0x6d, 0x4f, 0x2d, 0xa3, 0xff,
	0x5b, 0xce, 0xa4, 0x6d, 0x8e, 0x3e, 0xe3, 0xeb, 0x8c, 0x6f, 0x4b, 0x29, 0x94, 0xee, 0xcb, 0x64,
	0xe5, 0x44, 0x44, 0x56, 0x2e, 0x8a, 0x94, 0xd1, 0x7b, 0x19, 0x8f, 0x8c, 0x29, 0x17, 0xc5, 0x15,
	0xbd, 0x97, 0xda, 0x44, 0xef, 0x8c, 0x69, 0x6c, 0x4d, 0xf4, 0x4e, 0x9b, 0x66, 0x10, 0xfe, 0x46,
	0x33, 0xb1, 0xc1, 0xf8, 0x93, 0x79, 0x17, 0xe3, 0x3b, 0x39, 0x79, 0x0d, 0x53, 0x57, 0x58, 0xdb,
	0x67, 0x1f, 0xbd, 0x9d, 0x06, 0x7a, 0x93, 0x0b, 0x20, 0x37, 0xdb, 0x4d, 0x49, 0xd9, 0x12, 0x21,
	0xeb, 0x5a, 0xd1, 0x7e, 0x01, 0x67, 0x18, 0x81, 0x6e, 0xc6, 0x06, 0x8f, 0xc9, 0x97, 0xf0, 0xb4,
	0xf1, 0x84, 0x0d, 0x6a, 0x1d, 0x3b, 0x7b, 0xc7, 0x9f, 0x61, 0x78, 0x59, 0xe6, 0x39, 0x37, 0xd5,
	0xce, 0xcc, 0xd1, 0x0d, 0xa7, 0x15, 0xf7, 0xf3, 0x17, 0xf8, 0xf3, 0x77, 0x06, 0x91, 0xc6, 0xb6,
	0x1e, 0x59, 0x33, 0x99, 0x21, 0x2a, 0x70, 0x64, 0x93, 0x57, 0x30, 0xb9, 0xe2, 0x8a, 0x8a, 0xcd,
	0xe3, 0x9f, 0x9f, 0xfc, 0xd1, 0x83, 0xa9, 0x73, 0xb6, 0x1f, 0xea, 0x75, 0xbe, 0xd3, 0xec, 0xbc,
	0x07, 0xa6, 0xa0, 0x09, 0x26, 0xd7, 0xe3, 0xae, 0xd7, 0x63, 0x1f, 0x13, 0xbd, 0x03, 0x4c, 0x7c,
	0x0b, 0xa1, 0x4d, 0x50, 0xc6, 0x7d, 0x3d, 0x48, 0x2f, 0xda, 0x06, 0xc9, 0x96, 0x68, 0x55, 0x5f,
	0xf0, 0xc1, 0x33, 0x68, 0x82, 0xe7, 0x13, 0x00, 0xdb, 0x76, 0xc1, 0x65, 0x3c, 0x9c, 0x77, 0x71,
	0xe2, 0xf7, 0x1a, 0xec, 0xc1, 0xad, 0x60, 0x16, 0xa1, 0x78, 0xc4, 0xd2, 0x88, 0x7c, 0x8d, 0x5c,
	0x85, 0x00, 0x45, 0xff, 0xa1, 0xc8, 0xd7, 0x37, 0xd5, 0x46, 0x36, 0x90, 0x08, 0xed, 0x48, 0x1c,
	0x35, 0x91, 0x58, 0xcf, 0xe7, 0xb8, 0x9d, 0xc0, 0x26, 0xad, 0x04, 0x36, 0xf5, 0x47, 0xdd, 0x0d,
	0xda, 0xa9, 0x37, 0x68, 0x07, 0xa4, 0xf6, 0xe4, 0x90, 0xd4, 0x90, 0x09, 0xb4, 0x83, 0x60, 0xf1,
	0x07, 0x86, 0x09, 0x50, 0x34, 0xbc, 0xbc, 0x87, 0x0c, 0x79, 0x00, 0x99, 0xd1, 0x15, 0x55, 0xd4,
	0x01, 0xe6, 0x0c, 0x22, 0x07, 0x18, 0xdc, 0x16, 0x7a, 0x9e, 0x2c, 0x62, 0x64, 0xf2, 0x67, 0x0f,
	0x42, 0x74, 0xd6, 0x5b, 0x65, 0x0a, 0x41, 0x0d, 0xaa, 0x40, 0x30, 0x3f, 0x7c, 0xd0, 0x08, 0xdf,
	0xca, 0x50, 0x75, 0xee, 0x3d, 0x3f, 0x77, 0x0f, 0x6a, 0xfd, 0x63, 0xbc, 0x35, 0x38, 0xc2, 0x5b,
	0xc3, 0x26, 0x7a, 0xeb, 0xce, 0x84, 0x7e, 0x67, 0x5c, 0x91, 0xa3, 0x16, 0x36, 0x83, 0x16, 0x36,
	0x1b, 0xbd, 0x9f, 0xcd, 0xc6, 0xed, 0x6c, 0x36, 0x69, 0xc7, 0xd0, 0xb4, 0x89, 0xa1, 0x83, 0x2d,
	0x79, 0xfa, 0x60, 0x4b, 0xbe, 0x84, 0x53, 0x51, 0x30, 0x7e, 0x97, 0xee, 0x37, 0xad, 0x41, 0xc2,
	0x44, 0xab, 0x97, 0x6e, 0xdd, 0x3e, 0x07, 0xc8, 0x34, 0xf5, 0xb1, 0x94, 0x2a, 0x8d, 0x87, 0x68,
	0x15, 0x59, 0xcd, 0x85, 0x42, 0xf3, 0x6e, 0xcb, 0x9c, 0xd9, 0x60, 0x22, 0xb2, 0x1a, 0x63, 0x66,
	0x7c, 0xc3, 0xad, 0xf9, 0xa9, 0x31, 0x5b, 0xcd, 0x85, 0x4a, 0xae, 0x61, 0x6c, 0x30, 0x63, 0x79,
	0x63, 0xbf, 0x1a, 0x3b, 0xc7, 0x57, 0xa3, 0x03, 0x8f, 0x5b, 0x8d, 0x6f, 0xfe, 0xed, 0x42, 0xff,
	0x1a, 0x8f, 0xe4, 0x27, 0xe8, 0xe1, 0xe2, 0x24, 0x9f, 0x1d, 0x5b, 0xab, 0x16, 0xa5, 0xb3, 0xcf,
	0x8f, 0x3b, 0x99, 0xcf, 0x4a, 0x4e, 0xc8, 0x2f, 0x30, 0x30, 0x0b, 0x80, 0x7c, 0xd1, 0x4a, 0x32,
	0xfe, 0xe6, 0x9d, 0xbd, 0x7c, 0xcc, 0xad, 0x7e, 0xfa, 0x16, 0x46, 0x1e, 0xd7, 0x93, 0x57, 0x6d,
	0x17, 0x1f, 0xee, 0x94, 0xd9, 0xeb, 0xff, 0xe5, 0xeb, 0x27, 0x61, 0x78, 0xba, 0x3d, 0x89, 0x06,
	0xe9, 0xb7, 0x27, 0xd1, 0xa4, 0xfb, 0xe4, 0x84, 0xfc, 0x0a, 0x53, 0x6c, 0xc9, 0x77, 0xf7, 0xd7,
	0x76, 0xc4, 0xdb, 0xcb, 0xef, 0x91, 0x44, 0x7b, 0xf9, 0x7d, 0x54, 0x24, 0x27, 0x6f, 0x07, 0xfa,
	0x3f, 0xf1, 0xd7, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xdc, 0xe0, 0x11, 0x59, 0x2e, 0x0b, 0x00,
	0x00,
}
