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
}

func init() { proto.RegisterFile("proto/house/house.proto", fileDescriptor_fc2bc500116e1060) }

var fileDescriptor_fc2bc500116e1060 = []byte{
	// 837 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0x4b, 0x8f, 0xe3, 0x44,
	0x10, 0x5e, 0x3b, 0x0f, 0xdb, 0x95, 0xc9, 0xb0, 0x34, 0x88, 0x35, 0x59, 0xc1, 0x06, 0xf3, 0x1a,
	0xed, 0x4a, 0x41, 0x5a, 0x2e, 0x48, 0x9c, 0x56, 0x33, 0x07, 0x22, 0x21, 0x84, 0x8c, 0x06, 0x89,
	0x93, 0xd5, 0xeb, 0x6e, 0x32, 0x2d, 0xf9, 0x11, 0xba, 0x3b, 0xab, 0x9d, 0x3b, 0x27, 0x7e, 0x03,
	0xff, 0x8b, 0x3f, 0xc3, 0x01, 0x55, 0x75, 0x3b, 0xb1, 0x03, 0xde, 0xe5, 0x12, 0x75, 0x3d, 0xba,
	0xcb, 0x55, 0xdf, 0xf7, 0x95, 0x02, 0x8f, 0xf6, 0xba, 0xb5, 0xed, 0x57, 0x77, 0xed, 0xc1, 0x48,
	0xf7, 0xbb, 0x21, 0x0f, 0xfb, 0x60, 0xd7, 0x6e, 0x6a, 0x55, 0xea, 0x76, 0x63, 0xa4, 0x7e, 0xa5,
	0x4a, 0xb9, 0xa1, 0x68, 0x56, 0xc3, 0xe2, 0x7b, 0x65, 0x6c, 0x2e, 0x7f, 0x3b, 0x48, 0x63, 0xd9,
	0x43, 0x98, 0x70, 0x25, 0xd2, 0x60, 0x1d, 0x5c, 0x2d, 0x73, 0x3c, 0xb2, 0x8f, 0x00, 0x8c, 0xe5,
	0xda, 0x16, 0x82, 0x5b, 0x99, 0x86, 0xeb, 0xe0, 0x2a, 0xc9, 0x13, 0xf2, 0xdc, 0x70, 0x2b, 0xd9,
	0x87, 0x10, 0xcb, 0x46, 0xb8, 0xe0, 0x84, 0x82, 0x91, 0x6c, 0x04, 0x85, 0x18, 0x4c, 0xf7, 0x7c,
	0x27, 0xd3, 0x29, 0x3d, 0x46, 0xe7, 0xec, 0xcf, 0x10, 0x62, 0xac, 0xb7, 0xb5, 0xb2, 0x66, 0x29,
	0x44, 0x5c, 0x08, 0x2d, 0x8d, 0xa1, 0x82, 0x49, 0xde, 0x99, 0xec, 0x31, 0x24, 0x5c, 0x4b, 0x5e,
	0x34, 0xbc, 0xee, 0x6a, 0xc6, 0xe8, 0xf8, 0x81, 0xd7, 0x92, 0xbd, 0x0f, 0xb3, 0xd2, 0xaa, 0xba,
	0xab, 0xe7, 0x0c, 0xfc, 0x10, 0xea, 0xa8, 0x50, 0xc2, 0x57, 0x8c, 0xc8, 0xde, 0x0a, 0x7c, 0x4d,
	0xd5, 0x7c, 0x27, 0x8b, 0x83, 0xae, 0xd2, 0x99, 0x7b, 0x8d, 0x1c, 0xb7, 0xba, 0x62, 0x4f, 0x60,
	0xd1, 0x6a, 0x21, 0x75, 0x51, 0xb6, 0x87, 0xc6, 0xa6, 0x73, 0xba, 0x0a, 0xe4, 0xba, 0x46, 0x0f,
	0x96, 0xdb, 0x6b, 0x55, 0xca, 0x34, 0xa2, 0x90, 0x33, 0x70, 0x2c, 0xba, 0x6d, 0x6b, 0x7f, 0x2b,
	0xa6, 0x50, 0x82, 0x9e, 0xe3, 0x25, 0xab, 0x6c, 0x25, 0xd3, 0xc4, 0x7d, 0x23, 0x19, 0x58, 0xeb,
	0x60, 0xa4, 0x2e, 0xf8, 0x2b, 0x6e, 0xb9, 0x4e, 0x81, 0x62, 0x80, 0xae, 0x17, 0xe4, 0xc9, 0xfe,
	0x08, 0xe0, 0xc2, 0xc1, 0x61, 0xf6, 0x6d, 0x63, 0x24, 0xfb, 0x04, 0x2e, 0xca, 0x83, 0xd6, 0xb2,
	0xb1, 0x05, 0xcd, 0xd2, 0x01, 0xb3, 0xf0, 0xbe, 0x1f, 0xf9, 0x4e, 0xb2, 0x6f, 0x60, 0x4e, 0x8d,
	0x9a, 0x34, 0x5c, 0x4f, 0xae, 0x16, 0xcf, 0xd7, 0x9b, 0xff, 0x86, 0x7a, 0xd3, 0xcd, 0x3d, 0xf7,
	0xf9, 0xd8, 0x83, 0x6d, 0x2d, 0xaf, 0xdc, 0xd3, 0x13, 0xd7, 0x03, 0x79, 0xf0, 0xe1, 0xec, 0xaf,
	0x10, 0x96, 0xd7, 0x5a, 0x72, 0x2b, 0x3b, 0x76, 0x1c, 0xbb, 0x0a, 0xfa, 0x5d, 0x1d, 0x07, 0x14,
	0xf6, 0x07, 0xf4, 0x08, 0x22, 0x82, 0x50, 0x09, 0xff, 0xf2, 0x1c, 0xcd, 0xad, 0xe8, 0xa3, 0x3e,
	0x1d, 0xa2, 0x3e, 0x9c, 0xe9, 0xec, 0x7c, 0xa6, 0x78, 0xb1, 0xd4, 0x12, 0xbf, 0xd5, 0xa1, 0xd4,
	0x99, 0xc8, 0xb4, 0x43, 0xa3, 0x2c, 0x21, 0x94, 0xe4, 0x74, 0x66, 0x2b, 0x88, 0x4b, 0xbe, 0xe7,
	0xa5, 0xb2, 0xf7, 0x1e, 0x9e, 0xa3, 0x8d, 0xf9, 0x2f, 0xa5, 0x30, 0x1e, 0x1c, 0x3a, 0xe3, 0xeb,
	0x42, 0xee, 0x5b, 0xa3, 0x2c, 0xe1, 0xb2, 0xcc, 0x3b, 0x13, 0x99, 0x55, 0xab, 0xa6, 0x10, 0xfc,
	0xde, 0xa4, 0x0b, 0x17, 0xaa, 0x55, 0x73, 0xc3, 0xef, 0x0d, 0x85, 0xf8, 0x6b, 0x17, 0xba, 0xf0,
	0x21, 0xfe, 0x9a, 0x42, 0x2b, 0x88, 0x7f, 0xe5, 0xa5, 0xaa, 0xb0, 0xfe, 0x72, 0x3d, 0xc1, 0xfa,
	0x9d, 0x9d, 0x3d, 0x83, 0xcb, 0x6e, 0xb0, 0x1e, 0xe7, 0x3e, 0x7b, 0x83, 0x01, 0x7b, 0xb3, 0x17,
	0xc0, 0x6e, 0xf7, 0x55, 0xcb, 0xc5, 0x16, 0x29, 0xdb, 0x41, 0x31, 0x7e, 0x01, 0x35, 0x8c, 0x44,
	0x77, 0xb2, 0xc1, 0x63, 0xf6, 0x25, 0xbc, 0x37, 0x78, 0xc2, 0x17, 0xf5, 0x89, 0xc1, 0x29, 0xf1,
	0x67, 0x88, 0xae, 0xdb, 0xba, 0x96, 0x6e, 0xda, 0xa5, 0x3b, 0x76, 0xe2, 0xf4, 0xe6, 0x49, 0x7f,
	0x61, 0x5f, 0x7f, 0x8f, 0x21, 0x21, 0x6e, 0x93, 0x64, 0x9d, 0x32, 0x63, 0x74, 0xa0, 0x64, 0xb3,
	0xa7, 0xb0, 0xbc, 0x91, 0x96, 0xab, 0xea, 0xed, 0x9f, 0x9f, 0xfd, 0x3e, 0x85, 0xcb, 0x2e, 0xd9,
	0x7f, 0x68, 0x0f, 0xf9, 0x60, 0x88, 0x7c, 0x8f, 0x4c, 0xe1, 0x90, 0x4c, 0x1d, 0xc6, 0x93, 0x1e,
	0xc6, 0x7d, 0x4e, 0x4c, 0xcf, 0x38, 0xf1, 0x2d, 0xc4, 0xbe, 0x41, 0x93, 0xce, 0x48, 0x48, 0x4f,
	0xc6, 0x84, 0xe4, 0x47, 0x94, 0x1f, 0x2f, 0xf4, 0xc9, 0x33, 0x1f, 0x92, 0xe7, 0x63, 0x00, 0x0f,
	0xbb, 0x92, 0x26, 0x8d, 0xd6, 0x13, 0x54, 0xfc, 0xc9, 0x83, 0x18, 0xdc, 0x29, 0xe1, 0x19, 0x8a,
	0x47, 0x1c, 0x8d, 0xaa, 0x77, 0xb8, 0xab, 0x90, 0xa0, 0x98, 0x1f, 0xa9, 0x7a, 0x77, 0xab, 0x2b,
	0x33, 0x60, 0x22, 0x8c, 0x33, 0x71, 0x31, 0x64, 0xe2, 0x51, 0x9f, 0x17, 0xe3, 0x0b, 0x6c, 0x39,
	0xba, 0xc0, 0x2e, 0xfb, 0x52, 0xef, 0x84, 0xf6, 0x4e, 0x4f, 0x68, 0x67, 0x4b, 0xed, 0xe1, 0xf9,
	0x52, 0xc3, 0x4d, 0x40, 0x09, 0x4a, 0xa4, 0xef, 0xba, 0x4d, 0x80, 0xa6, 0xdb, 0xcb, 0x27, 0xca,
	0xb0, 0x21, 0x65, 0x9e, 0xff, 0x1d, 0xc2, 0xec, 0x3b, 0x1c, 0x37, 0xfb, 0x09, 0xa6, 0xb8, 0xba,
	0xd8, 0xa7, 0x6f, 0x5a, 0x6c, 0x9e, 0x58, 0xab, 0xcf, 0xde, 0x9c, 0xe4, 0x08, 0x95, 0x3d, 0x60,
	0xbf, 0xc0, 0xdc, 0x49, 0x90, 0x7d, 0x3e, 0x0a, 0x73, 0x7f, 0xf7, 0xad, 0xbe, 0x78, 0x5b, 0xda,
	0xf1, 0xe9, 0x3b, 0x58, 0xf4, 0xd4, 0xc6, 0x9e, 0x8e, 0x5d, 0xfc, 0xb7, 0xaa, 0x57, 0xcf, 0xfe,
	0x57, 0x6e, 0xbf, 0x09, 0xa7, 0x94, 0xf1, 0x26, 0x06, 0xb2, 0x1b, 0x6f, 0x62, 0x28, 0xb8, 0xec,
	0xc1, 0xcb, 0x39, 0xfd, 0x6d, 0xf8, 0xfa, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2c, 0xba, 0x91,
	0x15, 0x51, 0x08, 0x00, 0x00,
}
