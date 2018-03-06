// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/orderbook.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	proto/orderbook.proto

It has these top-level messages:
	Order
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Order_Side int32

const (
	Order_Buy  Order_Side = 0
	Order_Sell Order_Side = 1
)

var Order_Side_name = map[int32]string{
	0: "Buy",
	1: "Sell",
}
var Order_Side_value = map[string]int32{
	"Buy":  0,
	"Sell": 1,
}

func (x Order_Side) String() string {
	return proto.EnumName(Order_Side_name, int32(x))
}
func (Order_Side) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Order_Type int32

const (
	Order_GTC Order_Type = 0
	Order_GTT Order_Type = 1
	Order_ENE Order_Type = 2
	Order_FOK Order_Type = 3
)

var Order_Type_name = map[int32]string{
	0: "GTC",
	1: "GTT",
	2: "ENE",
	3: "FOK",
}
var Order_Type_value = map[string]int32{
	"GTC": 0,
	"GTT": 1,
	"ENE": 2,
	"FOK": 3,
}

func (x Order_Type) String() string {
	return proto.EnumName(Order_Type_name, int32(x))
}
func (Order_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

type Order struct {
	Market    string     `protobuf:"bytes,1,opt,name=market" json:"market,omitempty"`
	Party     string     `protobuf:"bytes,2,opt,name=party" json:"party,omitempty"`
	Side      Order_Side `protobuf:"varint,3,opt,name=side,enum=vega.Order_Side" json:"side,omitempty"`
	Price     uint64     `protobuf:"varint,4,opt,name=price" json:"price,omitempty"`
	Size      uint64     `protobuf:"varint,5,opt,name=size" json:"size,omitempty"`
	Remaining uint64     `protobuf:"varint,6,opt,name=remaining" json:"remaining,omitempty"`
	Type      Order_Type `protobuf:"varint,7,opt,name=type,enum=vega.Order_Type" json:"type,omitempty"`
	Sequence  uint64     `protobuf:"varint,8,opt,name=sequence" json:"sequence,omitempty"`
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Order) GetMarket() string {
	if m != nil {
		return m.Market
	}
	return ""
}

func (m *Order) GetParty() string {
	if m != nil {
		return m.Party
	}
	return ""
}

func (m *Order) GetSide() Order_Side {
	if m != nil {
		return m.Side
	}
	return Order_Buy
}

func (m *Order) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Order) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Order) GetRemaining() uint64 {
	if m != nil {
		return m.Remaining
	}
	return 0
}

func (m *Order) GetType() Order_Type {
	if m != nil {
		return m.Type
	}
	return Order_GTC
}

func (m *Order) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func init() {
	proto.RegisterType((*Order)(nil), "vega.Order")
	proto.RegisterEnum("vega.Order_Side", Order_Side_name, Order_Side_value)
	proto.RegisterEnum("vega.Order_Type", Order_Type_name, Order_Type_value)
}

func init() { proto.RegisterFile("proto/orderbook.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xbb, 0xc9, 0x26, 0x4d, 0xe7, 0x20, 0xcb, 0xa0, 0xb2, 0x8a, 0x87, 0x10, 0x3c, 0x04,
	0x0f, 0x11, 0xf4, 0x0d, 0x2a, 0xd5, 0x83, 0x60, 0x21, 0xcd, 0xc9, 0x5b, 0xd2, 0x0c, 0x25, 0xb4,
	0x4d, 0xd6, 0x6d, 0x2a, 0xc4, 0x67, 0xf2, 0x21, 0x65, 0x27, 0xa2, 0x87, 0xde, 0xfe, 0x6f, 0x3e,
	0xe6, 0x9f, 0x65, 0xe1, 0xc2, 0xd8, 0xae, 0xef, 0xee, 0x3b, 0x5b, 0x93, 0xad, 0xba, 0x6e, 0x9b,
	0x31, 0xa3, 0xfc, 0xa4, 0x4d, 0x99, 0x7c, 0x7b, 0x10, 0x2c, 0x9d, 0xc1, 0x4b, 0x08, 0xf7, 0xa5,
	0xdd, 0x52, 0xaf, 0x45, 0x2c, 0xd2, 0x59, 0xfe, 0x4b, 0x78, 0x0e, 0x81, 0x29, 0x6d, 0x3f, 0x68,
	0x8f, 0xc7, 0x23, 0xe0, 0x2d, 0xc8, 0x43, 0x53, 0x93, 0xf6, 0x63, 0x91, 0x9e, 0x3d, 0xa8, 0xcc,
	0x95, 0x65, 0x5c, 0x94, 0xad, 0x9a, 0x9a, 0x72, 0xb6, 0xbc, 0x6b, 0x9b, 0x35, 0x69, 0x19, 0x8b,
	0x54, 0xe6, 0x23, 0x20, 0xba, 0xdd, 0x2f, 0xd2, 0x01, 0x0f, 0x39, 0xe3, 0x0d, 0xcc, 0x2c, 0xed,
	0xcb, 0xa6, 0x6d, 0xda, 0x8d, 0x0e, 0x59, 0xfc, 0x0f, 0xdc, 0xb5, 0x7e, 0x30, 0xa4, 0xa7, 0xa7,
	0xd7, 0x8a, 0xc1, 0x50, 0xce, 0x16, 0xaf, 0x21, 0x3a, 0xd0, 0xc7, 0x91, 0xda, 0x35, 0xe9, 0x88,
	0x2b, 0xfe, 0x38, 0xb9, 0x02, 0xe9, 0xde, 0x85, 0x53, 0xf0, 0xe7, 0xc7, 0x41, 0x4d, 0x30, 0x02,
	0xb9, 0xa2, 0xdd, 0x4e, 0x89, 0xe4, 0x0e, 0xa4, 0x2b, 0x71, 0xea, 0xa5, 0x78, 0x52, 0x93, 0x31,
	0x14, 0x4a, 0xb8, 0xb0, 0x78, 0x5b, 0x28, 0xcf, 0x85, 0xe7, 0xe5, 0xab, 0xf2, 0xe7, 0xf2, 0xdd,
	0x33, 0x55, 0x15, 0xf2, 0x0f, 0x3e, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x51, 0x5d, 0x51, 0x44,
	0x5a, 0x01, 0x00, 0x00,
}
