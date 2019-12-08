// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/cart.proto

package pb

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_7deb92d3dea1f277, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type HealthMessage struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthMessage) Reset()         { *m = HealthMessage{} }
func (m *HealthMessage) String() string { return proto.CompactTextString(m) }
func (*HealthMessage) ProtoMessage()    {}
func (*HealthMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_7deb92d3dea1f277, []int{1}
}

func (m *HealthMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthMessage.Unmarshal(m, b)
}
func (m *HealthMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthMessage.Marshal(b, m, deterministic)
}
func (m *HealthMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthMessage.Merge(m, src)
}
func (m *HealthMessage) XXX_Size() int {
	return xxx_messageInfo_HealthMessage.Size(m)
}
func (m *HealthMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthMessage.DiscardUnknown(m)
}

var xxx_messageInfo_HealthMessage proto.InternalMessageInfo

func (m *HealthMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type StatusMessage struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Code                 int32    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Id                   int32    `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusMessage) Reset()         { *m = StatusMessage{} }
func (m *StatusMessage) String() string { return proto.CompactTextString(m) }
func (*StatusMessage) ProtoMessage()    {}
func (*StatusMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_7deb92d3dea1f277, []int{2}
}

func (m *StatusMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusMessage.Unmarshal(m, b)
}
func (m *StatusMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusMessage.Marshal(b, m, deterministic)
}
func (m *StatusMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusMessage.Merge(m, src)
}
func (m *StatusMessage) XXX_Size() int {
	return xxx_messageInfo_StatusMessage.Size(m)
}
func (m *StatusMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StatusMessage proto.InternalMessageInfo

func (m *StatusMessage) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *StatusMessage) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *StatusMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *StatusMessage) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Cart struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cart) Reset()         { *m = Cart{} }
func (m *Cart) String() string { return proto.CompactTextString(m) }
func (*Cart) ProtoMessage()    {}
func (*Cart) Descriptor() ([]byte, []int) {
	return fileDescriptor_7deb92d3dea1f277, []int{3}
}

func (m *Cart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cart.Unmarshal(m, b)
}
func (m *Cart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cart.Marshal(b, m, deterministic)
}
func (m *Cart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cart.Merge(m, src)
}
func (m *Cart) XXX_Size() int {
	return xxx_messageInfo_Cart.Size(m)
}
func (m *Cart) XXX_DiscardUnknown() {
	xxx_messageInfo_Cart.DiscardUnknown(m)
}

var xxx_messageInfo_Cart proto.InternalMessageInfo

func (m *Cart) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Product struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price                int32    `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_7deb92d3dea1f277, []int{4}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

type ProductItem struct {
	Product              *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	Quantity             int32    `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Amount               int32    `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductItem) Reset()         { *m = ProductItem{} }
func (m *ProductItem) String() string { return proto.CompactTextString(m) }
func (*ProductItem) ProtoMessage()    {}
func (*ProductItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_7deb92d3dea1f277, []int{5}
}

func (m *ProductItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductItem.Unmarshal(m, b)
}
func (m *ProductItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductItem.Marshal(b, m, deterministic)
}
func (m *ProductItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductItem.Merge(m, src)
}
func (m *ProductItem) XXX_Size() int {
	return xxx_messageInfo_ProductItem.Size(m)
}
func (m *ProductItem) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductItem.DiscardUnknown(m)
}

var xxx_messageInfo_ProductItem proto.InternalMessageInfo

func (m *ProductItem) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
}

func (m *ProductItem) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *ProductItem) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type AddRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Product              *Product `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	Quantity             int32    `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7deb92d3dea1f277, []int{6}
}

func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AddRequest) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
}

func (m *AddRequest) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type DeleteRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId            int32    `protobuf:"varint,2,opt,name=productId,proto3" json:"productId,omitempty"`
	Quantity             int32    `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7deb92d3dea1f277, []int{7}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DeleteRequest) GetProductId() int32 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *DeleteRequest) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type GetResponse struct {
	Id                   int32          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Products             []*ProductItem `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"`
	Quantity             int32          `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Amount               int32          `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7deb92d3dea1f277, []int{8}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetResponse) GetProducts() []*ProductItem {
	if m != nil {
		return m.Products
	}
	return nil
}

func (m *GetResponse) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *GetResponse) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "pb.Empty")
	proto.RegisterType((*HealthMessage)(nil), "pb.HealthMessage")
	proto.RegisterType((*StatusMessage)(nil), "pb.StatusMessage")
	proto.RegisterType((*Cart)(nil), "pb.Cart")
	proto.RegisterType((*Product)(nil), "pb.Product")
	proto.RegisterType((*ProductItem)(nil), "pb.ProductItem")
	proto.RegisterType((*AddRequest)(nil), "pb.AddRequest")
	proto.RegisterType((*DeleteRequest)(nil), "pb.DeleteRequest")
	proto.RegisterType((*GetResponse)(nil), "pb.GetResponse")
}

func init() { proto.RegisterFile("pb/cart.proto", fileDescriptor_7deb92d3dea1f277) }

var fileDescriptor_7deb92d3dea1f277 = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0x26, 0x49, 0xd3, 0x34, 0x27, 0x64, 0x57, 0x07, 0x59, 0x43, 0xd8, 0x8b, 0x12, 0x10, 0xea,
	0x0a, 0x0d, 0xd4, 0x0b, 0x61, 0xc1, 0x8b, 0x65, 0x95, 0xba, 0x17, 0x82, 0x8c, 0x57, 0x5e, 0x88,
	0x4c, 0x3b, 0x43, 0x1b, 0xb6, 0xc9, 0x64, 0x93, 0x13, 0x61, 0x6f, 0xbc, 0xf0, 0x15, 0x7c, 0x34,
	0x5f, 0xc1, 0x17, 0xf0, 0x0d, 0x24, 0x33, 0x93, 0x6e, 0x42, 0xb5, 0x78, 0x77, 0x7e, 0xbe, 0xf3,
	0x7d, 0x27, 0x5f, 0xce, 0x40, 0x58, 0xae, 0xd2, 0x35, 0xab, 0x70, 0x5e, 0x56, 0x12, 0x25, 0xb1,
	0xcb, 0x55, 0x7c, 0xbe, 0x91, 0x72, 0xb3, 0x13, 0x29, 0x2b, 0xb3, 0x94, 0x15, 0x85, 0x44, 0x86,
	0x99, 0x2c, 0x6a, 0x8d, 0x88, 0x9f, 0x7e, 0x65, 0xbb, 0x8c, 0x33, 0x14, 0x69, 0x17, 0xe8, 0x46,
	0xe2, 0x81, 0xfb, 0x36, 0x2f, 0xf1, 0x3e, 0x79, 0x0e, 0xe1, 0x3b, 0xc1, 0x76, 0xb8, 0x7d, 0x2f,
	0xea, 0x9a, 0x6d, 0x04, 0x89, 0xc0, 0xcb, 0x75, 0x18, 0x59, 0x53, 0x6b, 0xe6, 0xd3, 0x2e, 0x4d,
	0x3e, 0x43, 0xf8, 0x11, 0x19, 0x36, 0x75, 0x07, 0x3d, 0x01, 0x5b, 0xde, 0x2a, 0xd4, 0x84, 0xda,
	0xf2, 0x96, 0x10, 0x18, 0xad, 0x25, 0x17, 0x91, 0x3d, 0xb5, 0x66, 0x2e, 0x55, 0x71, 0x9f, 0xce,
	0x19, 0xd0, 0xb5, 0xd3, 0x19, 0x8f, 0x46, 0x0a, 0x6b, 0x67, 0x3c, 0x39, 0x83, 0xd1, 0x35, 0xab,
	0xd0, 0xd4, 0xad, 0x7d, 0xfd, 0x1a, 0xbc, 0x0f, 0x95, 0xe4, 0xcd, 0xfa, 0xa0, 0xd5, 0x0a, 0x16,
	0x2c, 0xd7, 0x82, 0x3e, 0x55, 0x31, 0x79, 0x02, 0x6e, 0x59, 0x65, 0x6b, 0x2d, 0xe7, 0x52, 0x9d,
	0x24, 0x5b, 0x08, 0x0c, 0xc9, 0x0d, 0x8a, 0x9c, 0x3c, 0x03, 0xaf, 0xd4, 0xa9, 0x62, 0x0b, 0x16,
	0xc1, 0xbc, 0x5c, 0xcd, 0x0d, 0x82, 0x76, 0x3d, 0x12, 0xc3, 0xe4, 0xae, 0x61, 0x05, 0x66, 0x78,
	0x6f, 0x3e, 0x6a, 0x9f, 0x93, 0x33, 0x18, 0xb3, 0x5c, 0x36, 0x05, 0x1a, 0x21, 0x93, 0x25, 0x5f,
	0x00, 0xae, 0x38, 0xa7, 0xe2, 0xae, 0x11, 0xf5, 0xe1, 0xc6, 0x3d, 0x61, 0xfb, 0x3f, 0x85, 0x9d,
	0xa1, 0x70, 0xf2, 0x09, 0xc2, 0x37, 0x62, 0x27, 0x50, 0xfc, 0x4b, 0xe3, 0x1c, 0x7c, 0xc3, 0x73,
	0xc3, 0xcd, 0xda, 0x0f, 0x85, 0xa3, 0xd4, 0xdf, 0x20, 0x58, 0x0a, 0xa4, 0xa2, 0x2e, 0x65, 0x51,
	0x8b, 0x03, 0xe2, 0x17, 0x30, 0x31, 0x3c, 0x75, 0x64, 0x4f, 0x9d, 0x59, 0xb0, 0x38, 0xed, 0x6d,
	0xdf, 0x1a, 0x4b, 0xf7, 0x80, 0x63, 0x3a, 0x3d, 0xef, 0x46, 0x7d, 0xef, 0x16, 0xbf, 0x2d, 0x00,
	0x2a, 0x1b, 0x14, 0xcb, 0x26, 0xe3, 0x82, 0xbc, 0x82, 0xb1, 0xbe, 0x4d, 0xe2, 0xb7, 0x3a, 0xea,
	0x60, 0xe3, 0xc7, 0x6d, 0x38, 0x38, 0xd9, 0xe4, 0xf4, 0xfb, 0xcf, 0x5f, 0x3f, 0x6c, 0x9f, 0x78,
	0xe9, 0x56, 0xc3, 0x2f, 0xc1, 0xb9, 0xe2, 0x9c, 0x9c, 0xb4, 0xd0, 0x87, 0x9f, 0xa1, 0x47, 0x07,
	0x27, 0x9c, 0x3c, 0x52, 0xa3, 0x10, 0xbb, 0xea, 0x5d, 0x5d, 0x5a, 0x17, 0xe4, 0x35, 0x8c, 0xb5,
	0xbd, 0x44, 0xc1, 0x07, 0x56, 0xff, 0x8d, 0x21, 0x54, 0x0c, 0xde, 0x85, 0x66, 0x20, 0x29, 0x38,
	0x4b, 0x81, 0x64, 0xd2, 0x02, 0xdb, 0x73, 0x8e, 0x95, 0x45, 0x3d, 0x57, 0xbb, 0x01, 0xa2, 0x07,
	0x56, 0x63, 0xf5, 0x20, 0x5f, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x8f, 0x06, 0x0f, 0xdc,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RouteGuideClient is the client API for RouteGuide service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouteGuideClient interface {
	Health(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HealthMessage, error)
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*StatusMessage, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*StatusMessage, error)
	Get(ctx context.Context, in *Cart, opts ...grpc.CallOption) (*GetResponse, error)
}

type routeGuideClient struct {
	cc *grpc.ClientConn
}

func NewRouteGuideClient(cc *grpc.ClientConn) RouteGuideClient {
	return &routeGuideClient{cc}
}

func (c *routeGuideClient) Health(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HealthMessage, error) {
	out := new(HealthMessage)
	err := c.cc.Invoke(ctx, "/pb.RouteGuide/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeGuideClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*StatusMessage, error) {
	out := new(StatusMessage)
	err := c.cc.Invoke(ctx, "/pb.RouteGuide/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeGuideClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*StatusMessage, error) {
	out := new(StatusMessage)
	err := c.cc.Invoke(ctx, "/pb.RouteGuide/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeGuideClient) Get(ctx context.Context, in *Cart, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/pb.RouteGuide/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouteGuideServer is the server API for RouteGuide service.
type RouteGuideServer interface {
	Health(context.Context, *Empty) (*HealthMessage, error)
	Add(context.Context, *AddRequest) (*StatusMessage, error)
	Delete(context.Context, *DeleteRequest) (*StatusMessage, error)
	Get(context.Context, *Cart) (*GetResponse, error)
}

// UnimplementedRouteGuideServer can be embedded to have forward compatible implementations.
type UnimplementedRouteGuideServer struct {
}

func (*UnimplementedRouteGuideServer) Health(ctx context.Context, req *Empty) (*HealthMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (*UnimplementedRouteGuideServer) Add(ctx context.Context, req *AddRequest) (*StatusMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedRouteGuideServer) Delete(ctx context.Context, req *DeleteRequest) (*StatusMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedRouteGuideServer) Get(ctx context.Context, req *Cart) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterRouteGuideServer(s *grpc.Server, srv RouteGuideServer) {
	s.RegisterService(&_RouteGuide_serviceDesc, srv)
}

func _RouteGuide_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuideServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RouteGuide/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuideServer).Health(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteGuide_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuideServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RouteGuide/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuideServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteGuide_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuideServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RouteGuide/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuideServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteGuide_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cart)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuideServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RouteGuide/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuideServer).Get(ctx, req.(*Cart))
	}
	return interceptor(ctx, in, info, handler)
}

var _RouteGuide_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.RouteGuide",
	HandlerType: (*RouteGuideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _RouteGuide_Health_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _RouteGuide_Add_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RouteGuide_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _RouteGuide_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/cart.proto",
}
