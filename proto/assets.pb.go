// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/assets.proto

package proto

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

// The Vega representation of an external asset.
type Asset struct {
	// Internal identifier of the asset.
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// Name of the asset (e.g: Great British Pound).
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Symbol of the asset (e.g: GBP).
	Symbol string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// Total circulating supply for the asset.
	TotalSupply string `protobuf:"bytes,4,opt,name=totalSupply,proto3" json:"totalSupply,omitempty"`
	// Number of decimals / precision handled by this asset.
	Decimals uint64 `protobuf:"varint,5,opt,name=decimals,proto3" json:"decimals,omitempty"`
	// The definition of the external source for this asset
	Source               *AssetSource `protobuf:"bytes,7,opt,name=source,proto3" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_c13390389c70361a, []int{0}
}

func (m *Asset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset.Unmarshal(m, b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
}
func (m *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(m, src)
}
func (m *Asset) XXX_Size() int {
	return xxx_messageInfo_Asset.Size(m)
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Asset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Asset) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Asset) GetTotalSupply() string {
	if m != nil {
		return m.TotalSupply
	}
	return ""
}

func (m *Asset) GetDecimals() uint64 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

func (m *Asset) GetSource() *AssetSource {
	if m != nil {
		return m.Source
	}
	return nil
}

// Asset source definition.
type AssetSource struct {
	// The source.
	//
	// Types that are valid to be assigned to Source:
	//	*AssetSource_BuiltinAsset
	//	*AssetSource_Erc20
	Source               isAssetSource_Source `protobuf_oneof:"source"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AssetSource) Reset()         { *m = AssetSource{} }
func (m *AssetSource) String() string { return proto.CompactTextString(m) }
func (*AssetSource) ProtoMessage()    {}
func (*AssetSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_c13390389c70361a, []int{1}
}

func (m *AssetSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetSource.Unmarshal(m, b)
}
func (m *AssetSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetSource.Marshal(b, m, deterministic)
}
func (m *AssetSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetSource.Merge(m, src)
}
func (m *AssetSource) XXX_Size() int {
	return xxx_messageInfo_AssetSource.Size(m)
}
func (m *AssetSource) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetSource.DiscardUnknown(m)
}

var xxx_messageInfo_AssetSource proto.InternalMessageInfo

type isAssetSource_Source interface {
	isAssetSource_Source()
}

type AssetSource_BuiltinAsset struct {
	BuiltinAsset *BuiltinAsset `protobuf:"bytes,1,opt,name=builtinAsset,proto3,oneof"`
}

type AssetSource_Erc20 struct {
	Erc20 *ERC20 `protobuf:"bytes,2,opt,name=erc20,proto3,oneof"`
}

func (*AssetSource_BuiltinAsset) isAssetSource_Source() {}

func (*AssetSource_Erc20) isAssetSource_Source() {}

func (m *AssetSource) GetSource() isAssetSource_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *AssetSource) GetBuiltinAsset() *BuiltinAsset {
	if x, ok := m.GetSource().(*AssetSource_BuiltinAsset); ok {
		return x.BuiltinAsset
	}
	return nil
}

func (m *AssetSource) GetErc20() *ERC20 {
	if x, ok := m.GetSource().(*AssetSource_Erc20); ok {
		return x.Erc20
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AssetSource) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AssetSource_BuiltinAsset)(nil),
		(*AssetSource_Erc20)(nil),
	}
}

// A Vega internal asset.
type BuiltinAsset struct {
	// Name of the asset (e.g: Great British Pound).
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Symbol of the asset (e.g: GBP).
	Symbol string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// Total circulating supply for the asset.
	TotalSupply string `protobuf:"bytes,3,opt,name=totalSupply,proto3" json:"totalSupply,omitempty"`
	// Number of decimal / precision handled by this asset.
	Decimals uint64 `protobuf:"varint,4,opt,name=decimals,proto3" json:"decimals,omitempty"`
	// Maximum amount that can be requested by a party through the built-in asset faucet at a time.
	MaxFaucetAmountMint  string   `protobuf:"bytes,5,opt,name=maxFaucetAmountMint,proto3" json:"maxFaucetAmountMint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuiltinAsset) Reset()         { *m = BuiltinAsset{} }
func (m *BuiltinAsset) String() string { return proto.CompactTextString(m) }
func (*BuiltinAsset) ProtoMessage()    {}
func (*BuiltinAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_c13390389c70361a, []int{2}
}

func (m *BuiltinAsset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuiltinAsset.Unmarshal(m, b)
}
func (m *BuiltinAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuiltinAsset.Marshal(b, m, deterministic)
}
func (m *BuiltinAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuiltinAsset.Merge(m, src)
}
func (m *BuiltinAsset) XXX_Size() int {
	return xxx_messageInfo_BuiltinAsset.Size(m)
}
func (m *BuiltinAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_BuiltinAsset.DiscardUnknown(m)
}

var xxx_messageInfo_BuiltinAsset proto.InternalMessageInfo

func (m *BuiltinAsset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BuiltinAsset) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *BuiltinAsset) GetTotalSupply() string {
	if m != nil {
		return m.TotalSupply
	}
	return ""
}

func (m *BuiltinAsset) GetDecimals() uint64 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

func (m *BuiltinAsset) GetMaxFaucetAmountMint() string {
	if m != nil {
		return m.MaxFaucetAmountMint
	}
	return ""
}

// An ERC20 token based asset, living on the ethereum network.
type ERC20 struct {
	// The address of the contract for the token, on the ethereum network
	ContractAddress      string   `protobuf:"bytes,1,opt,name=contractAddress,proto3" json:"contractAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ERC20) Reset()         { *m = ERC20{} }
func (m *ERC20) String() string { return proto.CompactTextString(m) }
func (*ERC20) ProtoMessage()    {}
func (*ERC20) Descriptor() ([]byte, []int) {
	return fileDescriptor_c13390389c70361a, []int{3}
}

func (m *ERC20) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ERC20.Unmarshal(m, b)
}
func (m *ERC20) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ERC20.Marshal(b, m, deterministic)
}
func (m *ERC20) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ERC20.Merge(m, src)
}
func (m *ERC20) XXX_Size() int {
	return xxx_messageInfo_ERC20.Size(m)
}
func (m *ERC20) XXX_DiscardUnknown() {
	xxx_messageInfo_ERC20.DiscardUnknown(m)
}

var xxx_messageInfo_ERC20 proto.InternalMessageInfo

func (m *ERC20) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

// Dev assets are for use in development networks only.
type DevAssets struct {
	// Asset sources for development networks.
	Sources              []*AssetSource `protobuf:"bytes,1,rep,name=sources,proto3" json:"sources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DevAssets) Reset()         { *m = DevAssets{} }
func (m *DevAssets) String() string { return proto.CompactTextString(m) }
func (*DevAssets) ProtoMessage()    {}
func (*DevAssets) Descriptor() ([]byte, []int) {
	return fileDescriptor_c13390389c70361a, []int{4}
}

func (m *DevAssets) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DevAssets.Unmarshal(m, b)
}
func (m *DevAssets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DevAssets.Marshal(b, m, deterministic)
}
func (m *DevAssets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DevAssets.Merge(m, src)
}
func (m *DevAssets) XXX_Size() int {
	return xxx_messageInfo_DevAssets.Size(m)
}
func (m *DevAssets) XXX_DiscardUnknown() {
	xxx_messageInfo_DevAssets.DiscardUnknown(m)
}

var xxx_messageInfo_DevAssets proto.InternalMessageInfo

func (m *DevAssets) GetSources() []*AssetSource {
	if m != nil {
		return m.Sources
	}
	return nil
}

func init() {
	proto.RegisterType((*Asset)(nil), "vega.Asset")
	proto.RegisterType((*AssetSource)(nil), "vega.AssetSource")
	proto.RegisterType((*BuiltinAsset)(nil), "vega.BuiltinAsset")
	proto.RegisterType((*ERC20)(nil), "vega.ERC20")
	proto.RegisterType((*DevAssets)(nil), "vega.DevAssets")
}

func init() { proto.RegisterFile("proto/assets.proto", fileDescriptor_c13390389c70361a) }

var fileDescriptor_c13390389c70361a = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0x4f, 0x83, 0x40,
	0x10, 0xc5, 0x0b, 0x85, 0xfe, 0x19, 0x1a, 0x8d, 0x6b, 0x62, 0x88, 0x17, 0x11, 0x2f, 0x18, 0x13,
	0x5a, 0xf1, 0xd2, 0x6b, 0x6b, 0x35, 0xed, 0xc1, 0xcb, 0xf6, 0xe6, 0x6d, 0xbb, 0x6c, 0x0c, 0x09,
	0xb0, 0x0d, 0xbb, 0x34, 0xf6, 0x2b, 0x19, 0x3f, 0xa4, 0x61, 0x68, 0x1b, 0x6c, 0x9a, 0x9e, 0x60,
	0x7e, 0xef, 0x2d, 0x79, 0xf3, 0x58, 0x20, 0xeb, 0x42, 0x6a, 0x39, 0x64, 0x4a, 0x09, 0xad, 0x42,
	0x1c, 0x88, 0xb5, 0x11, 0x5f, 0xcc, 0xff, 0x35, 0xc0, 0x9e, 0x54, 0x98, 0x5c, 0x80, 0xb9, 0x98,
	0xb9, 0x86, 0x67, 0x04, 0x7d, 0x6a, 0x2e, 0x66, 0x84, 0x80, 0x95, 0xb3, 0x4c, 0xb8, 0x26, 0x12,
	0x7c, 0x27, 0x37, 0xd0, 0x51, 0xdb, 0x6c, 0x25, 0x53, 0xb7, 0x8d, 0x74, 0x37, 0x11, 0x0f, 0x1c,
	0x2d, 0x35, 0x4b, 0x97, 0xe5, 0x7a, 0x9d, 0x6e, 0x5d, 0x0b, 0xc5, 0x26, 0x22, 0xb7, 0xd0, 0x8b,
	0x05, 0x4f, 0x32, 0x96, 0x2a, 0xd7, 0xf6, 0x8c, 0xc0, 0xa2, 0x87, 0x99, 0x3c, 0x42, 0x47, 0xc9,
	0xb2, 0xe0, 0xc2, 0xed, 0x7a, 0x46, 0xe0, 0x44, 0x57, 0x61, 0x15, 0x2d, 0xc4, 0x58, 0x4b, 0x14,
	0xe8, 0xce, 0xe0, 0x6f, 0xc0, 0x69, 0x60, 0x32, 0x86, 0xc1, 0xaa, 0x4c, 0x52, 0x9d, 0xe4, 0x48,
	0x31, 0xbd, 0x13, 0x91, 0xfa, 0xfc, 0xb4, 0xa1, 0xcc, 0x5b, 0xf4, 0x9f, 0x93, 0x3c, 0x80, 0x2d,
	0x0a, 0x1e, 0x8d, 0x70, 0x3d, 0x27, 0x72, 0xea, 0x23, 0x6f, 0xf4, 0x35, 0x1a, 0xcd, 0x5b, 0xb4,
	0xd6, 0xa6, 0xbd, 0x7d, 0x30, 0xff, 0xc7, 0x80, 0x41, 0xf3, 0x7b, 0x87, 0x76, 0x8c, 0x93, 0xed,
	0x98, 0xe7, 0xda, 0x69, 0x9f, 0x6f, 0xc7, 0x3a, 0x6a, 0x67, 0x04, 0xd7, 0x19, 0xfb, 0x7e, 0x67,
	0x25, 0x17, 0x7a, 0x92, 0xc9, 0x32, 0xd7, 0x1f, 0x49, 0xae, 0xb1, 0xc4, 0x3e, 0x3d, 0x25, 0xf9,
	0xcf, 0x60, 0xe3, 0x22, 0x24, 0x80, 0x4b, 0x2e, 0x73, 0x5d, 0x30, 0xae, 0x27, 0x71, 0x5c, 0x08,
	0xa5, 0x76, 0x79, 0x8f, 0xb1, 0x3f, 0x86, 0xfe, 0x4c, 0x6c, 0x70, 0x35, 0x45, 0x9e, 0xa0, 0x5b,
	0xaf, 0x5d, 0xd9, 0xdb, 0xa7, 0x7f, 0xc8, 0xde, 0x31, 0xbd, 0xff, 0xbc, 0xe3, 0x32, 0x16, 0xe8,
	0xc0, 0x8b, 0xc5, 0x65, 0x1a, 0x26, 0x72, 0x58, 0xcd, 0x43, 0x04, 0xab, 0x0e, 0x3e, 0x5e, 0xfe,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xfe, 0x68, 0x3c, 0x86, 0x02, 0x00, 0x00,
}
