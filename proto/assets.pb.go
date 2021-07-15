// Code generated by protoc-gen-go. DO NOT EDIT.
// source: assets.proto

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

// The Vega representation of an external asset
type Asset struct {
	// Internal identifier of the asset
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The definition of the external source for this asset
	Details              *AssetDetails `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ca40ce07a87fe, []int{0}
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

func (m *Asset) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Asset) GetDetails() *AssetDetails {
	if m != nil {
		return m.Details
	}
	return nil
}

// The Vega representation of an external asset
type AssetDetails struct {
	// Name of the asset (e.g: Great British Pound)
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Symbol of the asset (e.g: GBP)
	Symbol string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// Total circulating supply for the asset
	TotalSupply string `protobuf:"bytes,3,opt,name=total_supply,json=totalSupply,proto3" json:"total_supply,omitempty"`
	// Number of decimal / precision handled by this asset
	Decimals uint64 `protobuf:"varint,4,opt,name=decimals,proto3" json:"decimals,omitempty"`
	// Min stake required for this asset from liquidity providers
	MinLpStake string `protobuf:"bytes,5,opt,name=min_lp_stake,json=minLpStake,proto3" json:"min_lp_stake,omitempty"`
	// The source
	//
	// Types that are valid to be assigned to Source:
	//	*AssetDetails_BuiltinAsset
	//	*AssetDetails_Erc20
	Source               isAssetDetails_Source `protobuf_oneof:"source"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AssetDetails) Reset()         { *m = AssetDetails{} }
func (m *AssetDetails) String() string { return proto.CompactTextString(m) }
func (*AssetDetails) ProtoMessage()    {}
func (*AssetDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ca40ce07a87fe, []int{1}
}

func (m *AssetDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetDetails.Unmarshal(m, b)
}
func (m *AssetDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetDetails.Marshal(b, m, deterministic)
}
func (m *AssetDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetDetails.Merge(m, src)
}
func (m *AssetDetails) XXX_Size() int {
	return xxx_messageInfo_AssetDetails.Size(m)
}
func (m *AssetDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetDetails.DiscardUnknown(m)
}

var xxx_messageInfo_AssetDetails proto.InternalMessageInfo

func (m *AssetDetails) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AssetDetails) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *AssetDetails) GetTotalSupply() string {
	if m != nil {
		return m.TotalSupply
	}
	return ""
}

func (m *AssetDetails) GetDecimals() uint64 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

func (m *AssetDetails) GetMinLpStake() string {
	if m != nil {
		return m.MinLpStake
	}
	return ""
}

type isAssetDetails_Source interface {
	isAssetDetails_Source()
}

type AssetDetails_BuiltinAsset struct {
	BuiltinAsset *BuiltinAsset `protobuf:"bytes,101,opt,name=builtin_asset,json=builtinAsset,proto3,oneof"`
}

type AssetDetails_Erc20 struct {
	Erc20 *ERC20 `protobuf:"bytes,102,opt,name=erc20,proto3,oneof"`
}

func (*AssetDetails_BuiltinAsset) isAssetDetails_Source() {}

func (*AssetDetails_Erc20) isAssetDetails_Source() {}

func (m *AssetDetails) GetSource() isAssetDetails_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *AssetDetails) GetBuiltinAsset() *BuiltinAsset {
	if x, ok := m.GetSource().(*AssetDetails_BuiltinAsset); ok {
		return x.BuiltinAsset
	}
	return nil
}

func (m *AssetDetails) GetErc20() *ERC20 {
	if x, ok := m.GetSource().(*AssetDetails_Erc20); ok {
		return x.Erc20
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AssetDetails) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AssetDetails_BuiltinAsset)(nil),
		(*AssetDetails_Erc20)(nil),
	}
}

// A Vega internal asset
type BuiltinAsset struct {
	// Maximum amount that can be requested by a party through the built-in asset faucet at a time
	MaxFaucetAmountMint  string   `protobuf:"bytes,1,opt,name=max_faucet_amount_mint,json=maxFaucetAmountMint,proto3" json:"max_faucet_amount_mint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuiltinAsset) Reset()         { *m = BuiltinAsset{} }
func (m *BuiltinAsset) String() string { return proto.CompactTextString(m) }
func (*BuiltinAsset) ProtoMessage()    {}
func (*BuiltinAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ca40ce07a87fe, []int{2}
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

func (m *BuiltinAsset) GetMaxFaucetAmountMint() string {
	if m != nil {
		return m.MaxFaucetAmountMint
	}
	return ""
}

// An ERC20 token based asset, living on the ethereum network
type ERC20 struct {
	// The address of the contract for the token, on the ethereum network
	ContractAddress      string   `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ERC20) Reset()         { *m = ERC20{} }
func (m *ERC20) String() string { return proto.CompactTextString(m) }
func (*ERC20) ProtoMessage()    {}
func (*ERC20) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ca40ce07a87fe, []int{3}
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

func init() {
	proto.RegisterType((*Asset)(nil), "vega.Asset")
	proto.RegisterType((*AssetDetails)(nil), "vega.AssetDetails")
	proto.RegisterType((*BuiltinAsset)(nil), "vega.BuiltinAsset")
	proto.RegisterType((*ERC20)(nil), "vega.ERC20")
}

func init() {
	proto.RegisterFile("assets.proto", fileDescriptor_610ca40ce07a87fe)
}

var fileDescriptor_610ca40ce07a87fe = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x86, 0x37, 0x21, 0x29, 0xbb, 0xd3, 0xf0, 0x21, 0x23, 0xad, 0x22, 0x4e, 0x21, 0x20, 0x54,
	0x24, 0xc8, 0xae, 0xb2, 0x27, 0x8e, 0xed, 0x52, 0xd4, 0x03, 0x5c, 0xd2, 0x1b, 0x17, 0xcb, 0xb1,
	0x5d, 0x64, 0xe1, 0x8f, 0x28, 0x76, 0x50, 0xfb, 0x03, 0xf8, 0xdf, 0xa8, 0x93, 0x74, 0x95, 0x53,
	0x32, 0xcf, 0xe3, 0x77, 0x32, 0x19, 0x43, 0xc6, 0xbc, 0x97, 0xc1, 0x57, 0x5d, 0xef, 0x82, 0x23,
	0xc9, 0x5f, 0xf9, 0x9b, 0x95, 0x5b, 0x48, 0xd7, 0x67, 0x4a, 0x5e, 0x42, 0xac, 0x44, 0x1e, 0x15,
	0xd1, 0xea, 0xa6, 0x89, 0x95, 0x20, 0x9f, 0xe1, 0xb9, 0x90, 0x81, 0x29, 0xed, 0xf3, 0xb8, 0x88,
	0x56, 0xcb, 0x9a, 0x54, 0xe7, 0x40, 0x85, 0xa7, 0xbf, 0x8d, 0xa6, 0xb9, 0x1c, 0x29, 0xff, 0xc5,
	0x90, 0xcd, 0x0d, 0x21, 0x90, 0x58, 0x66, 0xe4, 0xd4, 0x10, 0xdf, 0xc9, 0x2d, 0x2c, 0xfc, 0xc9,
	0xb4, 0x4e, 0x63, 0xc7, 0x9b, 0x66, 0xaa, 0xc8, 0x3b, 0xc8, 0x82, 0x0b, 0x4c, 0x53, 0x3f, 0x74,
	0x9d, 0x3e, 0xe5, 0xcf, 0xd0, 0x2e, 0x91, 0xed, 0x11, 0x91, 0xb7, 0x70, 0x2d, 0x24, 0x57, 0x86,
	0x69, 0x9f, 0x27, 0x45, 0xb4, 0x4a, 0x9a, 0xa7, 0x9a, 0x14, 0x90, 0x19, 0x65, 0xa9, 0xee, 0xa8,
	0x0f, 0xec, 0x8f, 0xcc, 0x53, 0x8c, 0x83, 0x51, 0xf6, 0x47, 0xb7, 0x3f, 0x13, 0xf2, 0x15, 0x5e,
	0xb4, 0x83, 0xd2, 0x41, 0x59, 0x8a, 0x2b, 0xc8, 0xe5, 0xfc, 0x8f, 0x36, 0xa3, 0xc2, 0xf1, 0x77,
	0x57, 0x4d, 0xd6, 0xce, 0x6a, 0xf2, 0x1e, 0x52, 0xd9, 0xf3, 0xfa, 0x3e, 0x3f, 0x60, 0x64, 0x39,
	0x46, 0xb6, 0xcd, 0x63, 0x7d, 0xbf, 0xbb, 0x6a, 0x46, 0xb7, 0xb9, 0x86, 0x85, 0x77, 0x43, 0xcf,
	0x65, 0xf9, 0x08, 0xd9, 0xbc, 0x1d, 0x79, 0x80, 0x5b, 0xc3, 0x8e, 0xf4, 0xc0, 0x06, 0x2e, 0x03,
	0x65, 0xc6, 0x0d, 0x36, 0x50, 0xa3, 0x6c, 0x98, 0x16, 0xf3, 0xc6, 0xb0, 0xe3, 0x77, 0x94, 0x6b,
	0x74, 0x3f, 0x95, 0x0d, 0x65, 0x0d, 0x29, 0x7e, 0x80, 0x7c, 0x82, 0xd7, 0xdc, 0xd9, 0xd0, 0x33,
	0x1e, 0x28, 0x13, 0xa2, 0x97, 0xde, 0x4f, 0xb9, 0x57, 0x17, 0xbe, 0x1e, 0xf1, 0xe6, 0xe3, 0xaf,
	0x0f, 0xdc, 0x09, 0x89, 0xe3, 0xe1, 0xfd, 0x72, 0xa7, 0x2b, 0xe5, 0xee, 0x04, 0x0b, 0xec, 0x8b,
	0x75, 0x42, 0xde, 0x21, 0x6d, 0x17, 0xf8, 0x78, 0xf8, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xb4,
	0xbf, 0x4f, 0x0c, 0x02, 0x00, 0x00,
}
