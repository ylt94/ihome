// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/order/order.proto

package go_micro_service_order

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

type OrderStatus int32

const (
	OrderStatus_WAIT   OrderStatus = 0
	OrderStatus_ACCEPT OrderStatus = 1
	OrderStatus_REJECT OrderStatus = 2
)

var OrderStatus_name = map[int32]string{
	0: "WAIT",
	1: "ACCEPT",
	2: "REJECT",
}

var OrderStatus_value = map[string]int32{
	"WAIT":   0,
	"ACCEPT": 1,
	"REJECT": 2,
}

func (x OrderStatus) String() string {
	return proto.EnumName(OrderStatus_name, int32(x))
}

func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{0}
}

type CreateRequest struct {
	HouseId              uint32   `protobuf:"varint,1,opt,name=house_id,json=houseId,proto3" json:"house_id,omitempty"`
	StartDate            string   `protobuf:"bytes,2,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate              string   `protobuf:"bytes,3,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	UserId               uint32   `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{0}
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

func (m *CreateRequest) GetHouseId() uint32 {
	if m != nil {
		return m.HouseId
	}
	return 0
}

func (m *CreateRequest) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *CreateRequest) GetEndDate() string {
	if m != nil {
		return m.EndDate
	}
	return ""
}

func (m *CreateRequest) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type CreateResponse struct {
	OrderId              uint64   `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{1}
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

func (m *CreateResponse) GetOrderId() uint64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

type ListItem struct {
	Amount               uint32   `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Comment              string   `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	Ctime                string   `protobuf:"bytes,3,opt,name=ctime,proto3" json:"ctime,omitempty"`
	Days                 uint32   `protobuf:"varint,4,opt,name=days,proto3" json:"days,omitempty"`
	EndDate              string   `protobuf:"bytes,5,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	ImgUrl               string   `protobuf:"bytes,6,opt,name=img_url,json=imgUrl,proto3" json:"img_url,omitempty"`
	OrderId              uint32   `protobuf:"varint,7,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	StartDate            string   `protobuf:"bytes,8,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	Status               string   `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	Title                string   `protobuf:"bytes,10,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListItem) Reset()         { *m = ListItem{} }
func (m *ListItem) String() string { return proto.CompactTextString(m) }
func (*ListItem) ProtoMessage()    {}
func (*ListItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{2}
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

func (m *ListItem) GetAmount() uint32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ListItem) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *ListItem) GetCtime() string {
	if m != nil {
		return m.Ctime
	}
	return ""
}

func (m *ListItem) GetDays() uint32 {
	if m != nil {
		return m.Days
	}
	return 0
}

func (m *ListItem) GetEndDate() string {
	if m != nil {
		return m.EndDate
	}
	return ""
}

func (m *ListItem) GetImgUrl() string {
	if m != nil {
		return m.ImgUrl
	}
	return ""
}

func (m *ListItem) GetOrderId() uint32 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *ListItem) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *ListItem) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ListItem) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type ListRequest struct {
	Role                 string   `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	UserId               uint32   `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{3}
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

func (m *ListRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *ListRequest) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type ListResponse struct {
	Orders               []*ListItem `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{4}
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

func (m *ListResponse) GetOrders() []*ListItem {
	if m != nil {
		return m.Orders
	}
	return nil
}

type StatusRequest struct {
	Status               OrderStatus `protobuf:"varint,1,opt,name=status,proto3,enum=go.micro.service.order.OrderStatus" json:"status,omitempty"`
	UserId               uint32      `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderId              uint32      `protobuf:"varint,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *StatusRequest) Reset()         { *m = StatusRequest{} }
func (m *StatusRequest) String() string { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()    {}
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{5}
}

func (m *StatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusRequest.Unmarshal(m, b)
}
func (m *StatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusRequest.Marshal(b, m, deterministic)
}
func (m *StatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRequest.Merge(m, src)
}
func (m *StatusRequest) XXX_Size() int {
	return xxx_messageInfo_StatusRequest.Size(m)
}
func (m *StatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRequest proto.InternalMessageInfo

func (m *StatusRequest) GetStatus() OrderStatus {
	if m != nil {
		return m.Status
	}
	return OrderStatus_WAIT
}

func (m *StatusRequest) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *StatusRequest) GetOrderId() uint32 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

type StatusResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()    {}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{6}
}

func (m *StatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusResponse.Unmarshal(m, b)
}
func (m *StatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusResponse.Marshal(b, m, deterministic)
}
func (m *StatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponse.Merge(m, src)
}
func (m *StatusResponse) XXX_Size() int {
	return xxx_messageInfo_StatusResponse.Size(m)
}
func (m *StatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponse proto.InternalMessageInfo

type CommentRequest struct {
	OrderId              uint32   `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	UserId               uint32   `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Comment              string   `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommentRequest) Reset()         { *m = CommentRequest{} }
func (m *CommentRequest) String() string { return proto.CompactTextString(m) }
func (*CommentRequest) ProtoMessage()    {}
func (*CommentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{7}
}

func (m *CommentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommentRequest.Unmarshal(m, b)
}
func (m *CommentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommentRequest.Marshal(b, m, deterministic)
}
func (m *CommentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommentRequest.Merge(m, src)
}
func (m *CommentRequest) XXX_Size() int {
	return xxx_messageInfo_CommentRequest.Size(m)
}
func (m *CommentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommentRequest proto.InternalMessageInfo

func (m *CommentRequest) GetOrderId() uint32 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *CommentRequest) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CommentRequest) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

type CommentResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommentResponse) Reset()         { *m = CommentResponse{} }
func (m *CommentResponse) String() string { return proto.CompactTextString(m) }
func (*CommentResponse) ProtoMessage()    {}
func (*CommentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{8}
}

func (m *CommentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommentResponse.Unmarshal(m, b)
}
func (m *CommentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommentResponse.Marshal(b, m, deterministic)
}
func (m *CommentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommentResponse.Merge(m, src)
}
func (m *CommentResponse) XXX_Size() int {
	return xxx_messageInfo_CommentResponse.Size(m)
}
func (m *CommentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommentResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("go.micro.service.order.OrderStatus", OrderStatus_name, OrderStatus_value)
	proto.RegisterType((*CreateRequest)(nil), "go.micro.service.order.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "go.micro.service.order.CreateResponse")
	proto.RegisterType((*ListItem)(nil), "go.micro.service.order.ListItem")
	proto.RegisterType((*ListRequest)(nil), "go.micro.service.order.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "go.micro.service.order.ListResponse")
	proto.RegisterType((*StatusRequest)(nil), "go.micro.service.order.StatusRequest")
	proto.RegisterType((*StatusResponse)(nil), "go.micro.service.order.StatusResponse")
	proto.RegisterType((*CommentRequest)(nil), "go.micro.service.order.CommentRequest")
	proto.RegisterType((*CommentResponse)(nil), "go.micro.service.order.CommentResponse")
}

func init() { proto.RegisterFile("proto/order/order.proto", fileDescriptor_986e030a471601a2) }

var fileDescriptor_986e030a471601a2 = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xdd, 0x8b, 0xd3, 0x4c,
	0x14, 0xc6, 0x9b, 0x36, 0x4d, 0xda, 0xd3, 0x6d, 0xdf, 0xbe, 0x83, 0x74, 0xe3, 0x82, 0x50, 0xb2,
	0x7e, 0x14, 0x85, 0x2c, 0xd4, 0x1b, 0xd1, 0xab, 0xa5, 0x16, 0x8c, 0x08, 0x4a, 0xb6, 0x8b, 0x08,
	0x62, 0x89, 0xcd, 0x50, 0x03, 0x4d, 0xa6, 0xce, 0x4c, 0x04, 0x2f, 0x04, 0xf1, 0x8f, 0xf6, 0x5a,
	0xe6, 0xab, 0x64, 0xaa, 0xd9, 0xbd, 0x09, 0x73, 0xe6, 0xcc, 0x3c, 0xf3, 0x3b, 0xcf, 0x39, 0x2d,
	0x9c, 0xee, 0x29, 0xe1, 0xe4, 0x82, 0xd0, 0x0c, 0x53, 0xf5, 0x8d, 0xe4, 0x0e, 0x9a, 0x6c, 0x49,
	0x54, 0xe4, 0x1b, 0x4a, 0x22, 0x86, 0xe9, 0xb7, 0x7c, 0x83, 0x23, 0x99, 0x0d, 0x7f, 0xc0, 0x70,
	0x41, 0x71, 0xca, 0x71, 0x82, 0xbf, 0x56, 0x98, 0x71, 0x74, 0x17, 0x7a, 0x5f, 0x48, 0xc5, 0xf0,
	0x3a, 0xcf, 0x02, 0x67, 0xea, 0xcc, 0x86, 0x89, 0x2f, 0xe3, 0x38, 0x43, 0xf7, 0x00, 0x18, 0x4f,
	0x29, 0x5f, 0x67, 0x29, 0xc7, 0x41, 0x7b, 0xea, 0xcc, 0xfa, 0x49, 0x5f, 0xee, 0xbc, 0x4c, 0x39,
	0x16, 0x37, 0x71, 0x99, 0xa9, 0x64, 0x47, 0x26, 0x7d, 0x5c, 0x66, 0x32, 0x75, 0x0a, 0x7e, 0xc5,
	0x30, 0x15, 0x9a, 0xae, 0xd4, 0xf4, 0x44, 0x18, 0x67, 0xe1, 0x13, 0x18, 0x99, 0xe7, 0xd9, 0x9e,
	0x94, 0x4c, 0xaa, 0x48, 0x32, 0xf3, 0xbe, 0x9b, 0xf8, 0x32, 0x8e, 0xb3, 0xf0, 0x57, 0x1b, 0x7a,
	0x6f, 0x72, 0xc6, 0x63, 0x8e, 0x0b, 0x34, 0x01, 0x2f, 0x2d, 0x48, 0x55, 0x72, 0x4d, 0xa9, 0x23,
	0x14, 0x80, 0xbf, 0x21, 0x45, 0x81, 0x4b, 0xae, 0x09, 0x4d, 0x88, 0xee, 0x40, 0x77, 0xc3, 0xf3,
	0xc2, 0xc0, 0xa9, 0x00, 0x21, 0x70, 0xb3, 0xf4, 0x3b, 0xd3, 0x5c, 0x72, 0x6d, 0x55, 0xd2, 0xfd,
	0xab, 0x92, 0xbc, 0xd8, 0xae, 0x2b, 0xba, 0x0b, 0x3c, 0x99, 0xf1, 0xf2, 0x62, 0x7b, 0x4d, 0x77,
	0x16, 0xb7, 0xaf, 0x7c, 0xd3, 0xdc, 0x47, 0xbe, 0xf5, 0x8e, 0x7d, 0x9b, 0x80, 0xc7, 0x78, 0xca,
	0x2b, 0x16, 0xf4, 0x95, 0xa2, 0x8a, 0x04, 0x2f, 0xcf, 0xf9, 0x0e, 0x07, 0xa0, 0x78, 0x65, 0x10,
	0x3e, 0x87, 0x81, 0xf0, 0xc0, 0xb4, 0x0b, 0x81, 0x4b, 0xc9, 0x0e, 0x4b, 0x13, 0xfa, 0x89, 0x5c,
	0xd7, 0xdd, 0x6e, 0x5b, 0x6e, 0xbf, 0x82, 0x13, 0x75, 0x57, 0x7b, 0xfd, 0x0c, 0x3c, 0xc9, 0xc8,
	0x02, 0x67, 0xda, 0x99, 0x0d, 0xe6, 0xd3, 0xe8, 0xdf, 0x53, 0x12, 0x19, 0xd7, 0x13, 0x7d, 0x3e,
	0xfc, 0xe9, 0xc0, 0xf0, 0x4a, 0x62, 0x1a, 0x90, 0x17, 0x87, 0x2a, 0x04, 0xca, 0x68, 0x7e, 0xde,
	0xa4, 0xf5, 0x56, 0x7c, 0xf5, 0x5d, 0x53, 0x6a, 0x13, 0xb1, 0xe5, 0x6a, 0xc7, 0x72, 0x35, 0x1c,
	0xc3, 0xc8, 0x10, 0xa8, 0x72, 0xc2, 0x4f, 0x30, 0x5a, 0xa8, 0x5e, 0xd7, 0x86, 0xd9, 0x1a, 0xa6,
	0x5a, 0x53, 0x1a, 0x9f, 0xac, 0x0d, 0x50, 0xc7, 0x1a, 0xa0, 0xf0, 0x7f, 0xf8, 0xef, 0xa0, 0xaf,
	0x9e, 0x7c, 0x7c, 0x01, 0x83, 0x5a, 0x3d, 0xa8, 0x07, 0xee, 0xfb, 0xcb, 0x78, 0x35, 0x6e, 0x21,
	0x00, 0xef, 0x72, 0xb1, 0x58, 0xbe, 0x5b, 0x8d, 0x1d, 0xb1, 0x4e, 0x96, 0xaf, 0x97, 0x8b, 0xd5,
	0xb8, 0x3d, 0xff, 0xdd, 0x86, 0xae, 0xbc, 0x81, 0x3e, 0x80, 0xa7, 0x46, 0x1f, 0x3d, 0x68, 0xb2,
	0xca, 0xfa, 0x65, 0x9e, 0x3d, 0xbc, 0xed, 0x98, 0xb6, 0xa1, 0x85, 0xae, 0xc0, 0x15, 0x1d, 0x43,
	0xe7, 0x37, 0xf5, 0xd3, 0xc8, 0xde, 0xbf, 0xf9, 0xd0, 0x41, 0x74, 0x0d, 0x27, 0xd7, 0x7b, 0x31,
	0xc1, 0xba, 0xd6, 0x46, 0x6a, 0x6b, 0x2e, 0x9a, 0xa9, 0x8f, 0x9a, 0xd7, 0x42, 0x1f, 0xc1, 0xd7,
	0xf6, 0xa2, 0xe6, 0x52, 0xad, 0xfe, 0x9e, 0x3d, 0xba, 0xf5, 0x9c, 0x51, 0xff, 0xec, 0xc9, 0xff,
	0xc1, 0xa7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x67, 0x78, 0x55, 0x1c, 0x22, 0x05, 0x00, 0x00,
}
