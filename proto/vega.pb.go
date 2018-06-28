// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vega.proto

package msg

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

type Side int32

const (
	Side_Buy  Side = 0
	Side_Sell Side = 1
)

var Side_name = map[int32]string{
	0: "Buy",
	1: "Sell",
}
var Side_value = map[string]int32{
	"Buy":  0,
	"Sell": 1,
}

func (x Side) String() string {
	return proto.EnumName(Side_name, int32(x))
}
func (Side) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_vega_17c872f01e3ae008, []int{0}
}

type OrderError int32

const (
	OrderError_NONE                   OrderError = 0
	OrderError_INVALID_MARKET_ID      OrderError = 1
	OrderError_ORDER_OUT_OF_SEQUENCE  OrderError = 2
	OrderError_INVALID_REMAINING_SIZE OrderError = 3
	OrderError_NON_EMPTY_NEW_ORDER_ID OrderError = 4
)

var OrderError_name = map[int32]string{
	0: "NONE",
	1: "INVALID_MARKET_ID",
	2: "ORDER_OUT_OF_SEQUENCE",
	3: "INVALID_REMAINING_SIZE",
	4: "NON_EMPTY_NEW_ORDER_ID",
}
var OrderError_value = map[string]int32{
	"NONE":                   0,
	"INVALID_MARKET_ID":      1,
	"ORDER_OUT_OF_SEQUENCE":  2,
	"INVALID_REMAINING_SIZE": 3,
	"NON_EMPTY_NEW_ORDER_ID": 4,
}

func (x OrderError) String() string {
	return proto.EnumName(OrderError_name, int32(x))
}
func (OrderError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_vega_17c872f01e3ae008, []int{1}
}

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
func (Order_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_vega_17c872f01e3ae008, []int{0, 0}
}

type Order_Status int32

const (
	Order_NEW       Order_Status = 0
	Order_ACTIVE    Order_Status = 1
	Order_FILLED    Order_Status = 2
	Order_CANCELLED Order_Status = 3
)

var Order_Status_name = map[int32]string{
	0: "NEW",
	1: "ACTIVE",
	2: "FILLED",
	3: "CANCELLED",
}
var Order_Status_value = map[string]int32{
	"NEW":       0,
	"ACTIVE":    1,
	"FILLED":    2,
	"CANCELLED": 3,
}

func (x Order_Status) String() string {
	return proto.EnumName(Order_Status_name, int32(x))
}
func (Order_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_vega_17c872f01e3ae008, []int{0, 1}
}

type Order struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Market               string       `protobuf:"bytes,2,opt,name=market,proto3" json:"market,omitempty"`
	Party                string       `protobuf:"bytes,3,opt,name=party,proto3" json:"party,omitempty"`
	Side                 Side         `protobuf:"varint,4,opt,name=side,proto3,enum=vega.Side" json:"side,omitempty"`
	Price                uint64       `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	Size                 uint64       `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
	Remaining            uint64       `protobuf:"varint,7,opt,name=remaining,proto3" json:"remaining,omitempty"`
	Type                 Order_Type   `protobuf:"varint,8,opt,name=type,proto3,enum=vega.Order_Type" json:"type,omitempty"`
	Status               Order_Status `protobuf:"varint,9,opt,name=status,proto3,enum=vega.Order_Status" json:"status,omitempty"`
	Timestamp            uint64       `protobuf:"varint,10,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_vega_17c872f01e3ae008, []int{0}
}
func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (dst *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(dst, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

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

func (m *Order) GetSide() Side {
	if m != nil {
		return m.Side
	}
	return Side_Buy
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

func (m *Order) GetStatus() Order_Status {
	if m != nil {
		return m.Status
	}
	return Order_NEW
}

func (m *Order) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type OrderConfirmation struct {
	Order                *Order   `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	Trades               []*Trade `protobuf:"bytes,2,rep,name=trades,proto3" json:"trades,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderConfirmation) Reset()         { *m = OrderConfirmation{} }
func (m *OrderConfirmation) String() string { return proto.CompactTextString(m) }
func (*OrderConfirmation) ProtoMessage()    {}
func (*OrderConfirmation) Descriptor() ([]byte, []int) {
	return fileDescriptor_vega_17c872f01e3ae008, []int{1}
}
func (m *OrderConfirmation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderConfirmation.Unmarshal(m, b)
}
func (m *OrderConfirmation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderConfirmation.Marshal(b, m, deterministic)
}
func (dst *OrderConfirmation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderConfirmation.Merge(dst, src)
}
func (m *OrderConfirmation) XXX_Size() int {
	return xxx_messageInfo_OrderConfirmation.Size(m)
}
func (m *OrderConfirmation) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderConfirmation.DiscardUnknown(m)
}

var xxx_messageInfo_OrderConfirmation proto.InternalMessageInfo

func (m *OrderConfirmation) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *OrderConfirmation) GetTrades() []*Trade {
	if m != nil {
		return m.Trades
	}
	return nil
}

type Trade struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Market               string   `protobuf:"bytes,2,opt,name=market,proto3" json:"market,omitempty"`
	Price                uint64   `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	Size                 uint64   `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	Buyer                string   `protobuf:"bytes,5,opt,name=buyer,proto3" json:"buyer,omitempty"`
	Seller               string   `protobuf:"bytes,6,opt,name=seller,proto3" json:"seller,omitempty"`
	Aggressor            Side     `protobuf:"varint,7,opt,name=aggressor,proto3,enum=vega.Side" json:"aggressor,omitempty"`
	Timestamp            uint64   `protobuf:"varint,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Trade) Reset()         { *m = Trade{} }
func (m *Trade) String() string { return proto.CompactTextString(m) }
func (*Trade) ProtoMessage()    {}
func (*Trade) Descriptor() ([]byte, []int) {
	return fileDescriptor_vega_17c872f01e3ae008, []int{2}
}
func (m *Trade) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trade.Unmarshal(m, b)
}
func (m *Trade) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trade.Marshal(b, m, deterministic)
}
func (dst *Trade) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trade.Merge(dst, src)
}
func (m *Trade) XXX_Size() int {
	return xxx_messageInfo_Trade.Size(m)
}
func (m *Trade) XXX_DiscardUnknown() {
	xxx_messageInfo_Trade.DiscardUnknown(m)
}

var xxx_messageInfo_Trade proto.InternalMessageInfo

func (m *Trade) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Trade) GetMarket() string {
	if m != nil {
		return m.Market
	}
	return ""
}

func (m *Trade) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Trade) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Trade) GetBuyer() string {
	if m != nil {
		return m.Buyer
	}
	return ""
}

func (m *Trade) GetSeller() string {
	if m != nil {
		return m.Seller
	}
	return ""
}

func (m *Trade) GetAggressor() Side {
	if m != nil {
		return m.Aggressor
	}
	return Side_Buy
}

func (m *Trade) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type TradeSet struct {
	Trades               []*Trade `protobuf:"bytes,1,rep,name=trades,proto3" json:"trades,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TradeSet) Reset()         { *m = TradeSet{} }
func (m *TradeSet) String() string { return proto.CompactTextString(m) }
func (*TradeSet) ProtoMessage()    {}
func (*TradeSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_vega_17c872f01e3ae008, []int{3}
}
func (m *TradeSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradeSet.Unmarshal(m, b)
}
func (m *TradeSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradeSet.Marshal(b, m, deterministic)
}
func (dst *TradeSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeSet.Merge(dst, src)
}
func (m *TradeSet) XXX_Size() int {
	return xxx_messageInfo_TradeSet.Size(m)
}
func (m *TradeSet) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeSet.DiscardUnknown(m)
}

var xxx_messageInfo_TradeSet proto.InternalMessageInfo

func (m *TradeSet) GetTrades() []*Trade {
	if m != nil {
		return m.Trades
	}
	return nil
}

type MarketData struct {
	BestBid              uint64   `protobuf:"varint,1,opt,name=bestBid,proto3" json:"bestBid,omitempty"`
	BestOffer            uint64   `protobuf:"varint,2,opt,name=bestOffer,proto3" json:"bestOffer,omitempty"`
	LastTradedPrice      uint64   `protobuf:"varint,3,opt,name=lastTradedPrice,proto3" json:"lastTradedPrice,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MarketData) Reset()         { *m = MarketData{} }
func (m *MarketData) String() string { return proto.CompactTextString(m) }
func (*MarketData) ProtoMessage()    {}
func (*MarketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_vega_17c872f01e3ae008, []int{4}
}
func (m *MarketData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarketData.Unmarshal(m, b)
}
func (m *MarketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarketData.Marshal(b, m, deterministic)
}
func (dst *MarketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketData.Merge(dst, src)
}
func (m *MarketData) XXX_Size() int {
	return xxx_messageInfo_MarketData.Size(m)
}
func (m *MarketData) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketData.DiscardUnknown(m)
}

var xxx_messageInfo_MarketData proto.InternalMessageInfo

func (m *MarketData) GetBestBid() uint64 {
	if m != nil {
		return m.BestBid
	}
	return 0
}

func (m *MarketData) GetBestOffer() uint64 {
	if m != nil {
		return m.BestOffer
	}
	return 0
}

func (m *MarketData) GetLastTradedPrice() uint64 {
	if m != nil {
		return m.LastTradedPrice
	}
	return 0
}

type MarketDepth struct {
	BuyOrderCount        uint64   `protobuf:"varint,1,opt,name=buyOrderCount,proto3" json:"buyOrderCount,omitempty"`
	SellOrderCount       uint64   `protobuf:"varint,2,opt,name=sellOrderCount,proto3" json:"sellOrderCount,omitempty"`
	BuyOrderVolume       uint64   `protobuf:"varint,3,opt,name=buyOrderVolume,proto3" json:"buyOrderVolume,omitempty"`
	SellOrderVolume      uint64   `protobuf:"varint,4,opt,name=sellOrderVolume,proto3" json:"sellOrderVolume,omitempty"`
	BuyPriceLevels       uint64   `protobuf:"varint,5,opt,name=buyPriceLevels,proto3" json:"buyPriceLevels,omitempty"`
	SellPriceLevels      uint64   `protobuf:"varint,6,opt,name=sellPriceLevels,proto3" json:"sellPriceLevels,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MarketDepth) Reset()         { *m = MarketDepth{} }
func (m *MarketDepth) String() string { return proto.CompactTextString(m) }
func (*MarketDepth) ProtoMessage()    {}
func (*MarketDepth) Descriptor() ([]byte, []int) {
	return fileDescriptor_vega_17c872f01e3ae008, []int{5}
}
func (m *MarketDepth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarketDepth.Unmarshal(m, b)
}
func (m *MarketDepth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarketDepth.Marshal(b, m, deterministic)
}
func (dst *MarketDepth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketDepth.Merge(dst, src)
}
func (m *MarketDepth) XXX_Size() int {
	return xxx_messageInfo_MarketDepth.Size(m)
}
func (m *MarketDepth) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketDepth.DiscardUnknown(m)
}

var xxx_messageInfo_MarketDepth proto.InternalMessageInfo

func (m *MarketDepth) GetBuyOrderCount() uint64 {
	if m != nil {
		return m.BuyOrderCount
	}
	return 0
}

func (m *MarketDepth) GetSellOrderCount() uint64 {
	if m != nil {
		return m.SellOrderCount
	}
	return 0
}

func (m *MarketDepth) GetBuyOrderVolume() uint64 {
	if m != nil {
		return m.BuyOrderVolume
	}
	return 0
}

func (m *MarketDepth) GetSellOrderVolume() uint64 {
	if m != nil {
		return m.SellOrderVolume
	}
	return 0
}

func (m *MarketDepth) GetBuyPriceLevels() uint64 {
	if m != nil {
		return m.BuyPriceLevels
	}
	return 0
}

func (m *MarketDepth) GetSellPriceLevels() uint64 {
	if m != nil {
		return m.SellPriceLevels
	}
	return 0
}

func init() {
	proto.RegisterType((*Order)(nil), "vega.Order")
	proto.RegisterType((*OrderConfirmation)(nil), "vega.OrderConfirmation")
	proto.RegisterType((*Trade)(nil), "vega.Trade")
	proto.RegisterType((*TradeSet)(nil), "vega.TradeSet")
	proto.RegisterType((*MarketData)(nil), "vega.MarketData")
	proto.RegisterType((*MarketDepth)(nil), "vega.MarketDepth")
	proto.RegisterEnum("vega.Side", Side_name, Side_value)
	proto.RegisterEnum("vega.OrderError", OrderError_name, OrderError_value)
	proto.RegisterEnum("vega.Order_Type", Order_Type_name, Order_Type_value)
	proto.RegisterEnum("vega.Order_Status", Order_Status_name, Order_Status_value)
}

func init() { proto.RegisterFile("vega.proto", fileDescriptor_vega_17c872f01e3ae008) }

var fileDescriptor_vega_17c872f01e3ae008 = []byte{
	// 673 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcf, 0x6f, 0xda, 0x48,
	0x14, 0xc7, 0x31, 0x36, 0x0e, 0xbc, 0x28, 0xd9, 0xc9, 0x28, 0x89, 0x9c, 0xd5, 0x6a, 0xc5, 0xb2,
	0x51, 0x85, 0x38, 0xa4, 0x52, 0x7a, 0xe9, 0x95, 0xc0, 0x24, 0xb2, 0x02, 0x76, 0x3a, 0x38, 0x89,
	0x9a, 0x1e, 0x2c, 0x53, 0x26, 0xd4, 0x2a, 0xc6, 0x68, 0x3c, 0x44, 0xa2, 0xa7, 0x4a, 0xfd, 0x03,
	0xfb, 0x27, 0xb5, 0x9a, 0x67, 0x13, 0x1c, 0xd4, 0x1e, 0x7a, 0x7b, 0xef, 0xf3, 0x7e, 0xce, 0xfb,
	0x62, 0x00, 0x9e, 0xc4, 0x34, 0x3a, 0x5b, 0xc8, 0x54, 0xa5, 0xd4, 0xd2, 0x76, 0xeb, 0x9b, 0x09,
	0x35, 0x5f, 0x4e, 0x84, 0xa4, 0xfb, 0x50, 0x8d, 0x27, 0x8e, 0xd1, 0x34, 0xda, 0x0d, 0x5e, 0x8d,
	0x27, 0xf4, 0x18, 0xec, 0x24, 0x92, 0x9f, 0x85, 0x72, 0xaa, 0xc8, 0x0a, 0x8f, 0x1e, 0x42, 0x6d,
	0x11, 0x49, 0xb5, 0x72, 0x4c, 0xc4, 0xb9, 0x43, 0xff, 0x05, 0x2b, 0x8b, 0x27, 0xc2, 0xb1, 0x9a,
	0x46, 0x7b, 0xff, 0x1c, 0xce, 0x70, 0xd0, 0x28, 0x9e, 0x08, 0x8e, 0x1c, 0xab, 0x64, 0xfc, 0x51,
	0x38, 0xb5, 0xa6, 0xd1, 0xb6, 0x78, 0xee, 0x50, 0xaa, 0xab, 0xbe, 0x08, 0xc7, 0x46, 0x88, 0x36,
	0xfd, 0x07, 0x1a, 0x52, 0x24, 0x51, 0x3c, 0x8f, 0xe7, 0x53, 0x67, 0x07, 0x03, 0x1b, 0x40, 0x4f,
	0xc1, 0x52, 0xab, 0x85, 0x70, 0xea, 0x38, 0x87, 0xe4, 0x73, 0xf0, 0x01, 0x67, 0xc1, 0x6a, 0x21,
	0x38, 0x46, 0x69, 0x07, 0xec, 0x4c, 0x45, 0x6a, 0x99, 0x39, 0x0d, 0xcc, 0xa3, 0xe5, 0xbc, 0x11,
	0x46, 0x78, 0x91, 0xa1, 0xe7, 0xa9, 0x38, 0x11, 0x99, 0x8a, 0x92, 0x85, 0x03, 0xf9, 0xbc, 0x67,
	0xd0, 0xea, 0x80, 0xa5, 0xfb, 0xd2, 0x1d, 0x30, 0xaf, 0x82, 0x1e, 0xa9, 0xe4, 0x46, 0x40, 0x0c,
	0x6d, 0x30, 0x8f, 0x91, 0xaa, 0x36, 0x2e, 0xfd, 0x6b, 0x62, 0xb6, 0xde, 0x82, 0x9d, 0xf7, 0xd6,
	0xc8, 0x63, 0xf7, 0xa4, 0x42, 0x01, 0xec, 0x6e, 0x2f, 0x70, 0xef, 0x18, 0x31, 0xb4, 0x7d, 0xe9,
	0x0e, 0x06, 0xac, 0x4f, 0xaa, 0x74, 0x0f, 0x1a, 0xbd, 0xae, 0xd7, 0x63, 0xe8, 0x9a, 0xad, 0x0f,
	0x70, 0x80, 0xbb, 0xf5, 0xd2, 0xf9, 0x63, 0x2c, 0x93, 0x48, 0xc5, 0xe9, 0x9c, 0xfe, 0x07, 0xb5,
	0x54, 0x43, 0xd4, 0x64, 0xf7, 0x7c, 0xb7, 0xf4, 0x06, 0x9e, 0x47, 0xe8, 0xff, 0x60, 0x2b, 0x19,
	0x4d, 0x44, 0xe6, 0x54, 0x9b, 0xe6, 0x26, 0x27, 0xd0, 0x8c, 0x17, 0xa1, 0xd6, 0x77, 0x03, 0x6a,
	0x48, 0xfe, 0x48, 0x62, 0x14, 0xcb, 0xfc, 0x95, 0x58, 0x56, 0x49, 0xac, 0x43, 0xa8, 0x8d, 0x97,
	0x2b, 0x21, 0x51, 0xd6, 0x06, 0xcf, 0x1d, 0xdd, 0x37, 0x13, 0xb3, 0x99, 0x90, 0x28, 0x6c, 0x83,
	0x17, 0x1e, 0x6d, 0x43, 0x23, 0x9a, 0x4e, 0xa5, 0xc8, 0xb2, 0x54, 0xa2, 0xb4, 0x2f, 0x7f, 0x29,
	0x9b, 0xe0, 0x4b, 0x51, 0xea, 0xdb, 0xa2, 0xbc, 0x86, 0x3a, 0x3e, 0x68, 0x24, 0x54, 0xe9, 0x04,
	0xc6, 0xef, 0x4f, 0x30, 0x07, 0x18, 0xe2, 0xd3, 0xfa, 0x91, 0x8a, 0xa8, 0x03, 0x3b, 0x63, 0x91,
	0xa9, 0x8b, 0xe2, 0x16, 0x16, 0x5f, 0xbb, 0x7a, 0xac, 0x36, 0xfd, 0xc7, 0x47, 0x21, 0xf1, 0x26,
	0x16, 0xdf, 0x00, 0xda, 0x86, 0xbf, 0x66, 0x51, 0xa6, 0xb0, 0xf5, 0xe4, 0xa6, 0x74, 0xa0, 0x6d,
	0xdc, 0xfa, 0x61, 0xc0, 0x6e, 0x31, 0x50, 0x2c, 0xd4, 0x27, 0x7a, 0x0a, 0x7b, 0xe3, 0xe5, 0xaa,
	0x90, 0x78, 0x39, 0x57, 0xc5, 0xdc, 0x97, 0x90, 0xbe, 0x82, 0x7d, 0x7d, 0xa8, 0x52, 0x5a, 0xbe,
	0xc2, 0x16, 0xd5, 0x79, 0xeb, 0xc2, 0xbb, 0x74, 0xb6, 0x4c, 0xd6, 0x6b, 0x6c, 0x51, 0xbd, 0xef,
	0x73, 0x65, 0x91, 0x98, 0x6b, 0xb7, 0x8d, 0x8b, 0x8e, 0xb8, 0xfb, 0x40, 0x3c, 0x89, 0x59, 0x56,
	0x7c, 0xa6, 0x5b, 0x74, 0xdd, 0xb1, 0x9c, 0x68, 0x6f, 0x3a, 0x96, 0x70, 0xe7, 0x04, 0x2c, 0xad,
	0xa9, 0xfe, 0x12, 0x2e, 0x96, 0x2b, 0x52, 0xa1, 0x75, 0xb0, 0x46, 0x62, 0x36, 0x23, 0x46, 0xe7,
	0xab, 0x01, 0x80, 0xc3, 0x99, 0x94, 0xa9, 0xd4, 0x01, 0xcf, 0xf7, 0x18, 0xa9, 0xd0, 0x23, 0x38,
	0x70, 0xbd, 0xbb, 0xee, 0xc0, 0xed, 0x87, 0xc3, 0x2e, 0xbf, 0x66, 0x41, 0xe8, 0xf6, 0x89, 0x41,
	0x4f, 0xe0, 0xc8, 0xe7, 0x7d, 0xc6, 0x43, 0xff, 0x36, 0x08, 0xfd, 0xcb, 0x70, 0xc4, 0xde, 0xdd,
	0x32, 0xaf, 0xa7, 0x3f, 0xbd, 0xbf, 0xe1, 0x78, 0x5d, 0xc1, 0xd9, 0xb0, 0xeb, 0x7a, 0xae, 0x77,
	0x15, 0x8e, 0xdc, 0x07, 0x46, 0x4c, 0x1d, 0xf3, 0x7c, 0x2f, 0x64, 0xc3, 0x9b, 0xe0, 0x7d, 0xe8,
	0xb1, 0xfb, 0x30, 0x6f, 0xe2, 0xf6, 0x89, 0x75, 0x51, 0x7b, 0x30, 0x93, 0x6c, 0x3a, 0xb6, 0xf1,
	0x9f, 0xf0, 0xcd, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xae, 0x0f, 0xa0, 0xe0, 0x17, 0x05, 0x00,
	0x00,
}
