// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/governance.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
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

// Proposal state transition:
// Open ->
//   - Passed -> Enacted.
//   - Passed -> Failed.
//   - Declined
// Rejected
// Proposal can enter Failed state from any other state.
type Proposal_State int32

const (
	// Proposal could not be enacted after being accepted by the network
	Proposal_FAILED Proposal_State = 0
	// Proposal is open for voting.
	Proposal_OPEN Proposal_State = 1
	// Proposal has gained enough support to be executed.
	Proposal_PASSED Proposal_State = 2
	// Proposal wasn't accepted (validation failed, author not allowed to submit proposals)
	Proposal_REJECTED Proposal_State = 3
	// Proposal didn't get enough votes
	Proposal_DECLINED Proposal_State = 4
	// Proposal has been executed and the changes under this proposal have now been applied.
	Proposal_ENACTED Proposal_State = 5
)

var Proposal_State_name = map[int32]string{
	0: "FAILED",
	1: "OPEN",
	2: "PASSED",
	3: "REJECTED",
	4: "DECLINED",
	5: "ENACTED",
}

var Proposal_State_value = map[string]int32{
	"FAILED":   0,
	"OPEN":     1,
	"PASSED":   2,
	"REJECTED": 3,
	"DECLINED": 4,
	"ENACTED":  5,
}

func (x Proposal_State) String() string {
	return proto.EnumName(Proposal_State_name, int32(x))
}

func (Proposal_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c891e73c7d2524a3, []int{5, 0}
}

type Vote_Value int32

const (
	Vote_NO  Vote_Value = 0
	Vote_YES Vote_Value = 1
)

var Vote_Value_name = map[int32]string{
	0: "NO",
	1: "YES",
}

var Vote_Value_value = map[string]int32{
	"NO":  0,
	"YES": 1,
}

func (x Vote_Value) String() string {
	return proto.EnumName(Vote_Value_name, int32(x))
}

func (Vote_Value) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c891e73c7d2524a3, []int{6, 0}
}

type NetworkConfiguration struct {
	MinCloseInSeconds     int64    `protobuf:"varint,1,opt,name=minCloseInSeconds,proto3" json:"minCloseInSeconds,omitempty"`
	MaxCloseInSeconds     int64    `protobuf:"varint,2,opt,name=maxCloseInSeconds,proto3" json:"maxCloseInSeconds,omitempty"`
	MinEnactInSeconds     int64    `protobuf:"varint,3,opt,name=minEnactInSeconds,proto3" json:"minEnactInSeconds,omitempty"`
	MaxEnactInSeconds     int64    `protobuf:"varint,4,opt,name=maxEnactInSeconds,proto3" json:"maxEnactInSeconds,omitempty"`
	MinParticipationStake uint64   `protobuf:"varint,5,opt,name=minParticipationStake,proto3" json:"minParticipationStake,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *NetworkConfiguration) Reset()         { *m = NetworkConfiguration{} }
func (m *NetworkConfiguration) String() string { return proto.CompactTextString(m) }
func (*NetworkConfiguration) ProtoMessage()    {}
func (*NetworkConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_c891e73c7d2524a3, []int{0}
}

func (m *NetworkConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkConfiguration.Unmarshal(m, b)
}
func (m *NetworkConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkConfiguration.Marshal(b, m, deterministic)
}
func (m *NetworkConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkConfiguration.Merge(m, src)
}
func (m *NetworkConfiguration) XXX_Size() int {
	return xxx_messageInfo_NetworkConfiguration.Size(m)
}
func (m *NetworkConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkConfiguration proto.InternalMessageInfo

func (m *NetworkConfiguration) GetMinCloseInSeconds() int64 {
	if m != nil {
		return m.MinCloseInSeconds
	}
	return 0
}

func (m *NetworkConfiguration) GetMaxCloseInSeconds() int64 {
	if m != nil {
		return m.MaxCloseInSeconds
	}
	return 0
}

func (m *NetworkConfiguration) GetMinEnactInSeconds() int64 {
	if m != nil {
		return m.MinEnactInSeconds
	}
	return 0
}

func (m *NetworkConfiguration) GetMaxEnactInSeconds() int64 {
	if m != nil {
		return m.MaxEnactInSeconds
	}
	return 0
}

func (m *NetworkConfiguration) GetMinParticipationStake() uint64 {
	if m != nil {
		return m.MinParticipationStake
	}
	return 0
}

type UpdateMarket struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateMarket) Reset()         { *m = UpdateMarket{} }
func (m *UpdateMarket) String() string { return proto.CompactTextString(m) }
func (*UpdateMarket) ProtoMessage()    {}
func (*UpdateMarket) Descriptor() ([]byte, []int) {
	return fileDescriptor_c891e73c7d2524a3, []int{1}
}

func (m *UpdateMarket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateMarket.Unmarshal(m, b)
}
func (m *UpdateMarket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateMarket.Marshal(b, m, deterministic)
}
func (m *UpdateMarket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMarket.Merge(m, src)
}
func (m *UpdateMarket) XXX_Size() int {
	return xxx_messageInfo_UpdateMarket.Size(m)
}
func (m *UpdateMarket) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMarket.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMarket proto.InternalMessageInfo

type NewMarket struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewMarket) Reset()         { *m = NewMarket{} }
func (m *NewMarket) String() string { return proto.CompactTextString(m) }
func (*NewMarket) ProtoMessage()    {}
func (*NewMarket) Descriptor() ([]byte, []int) {
	return fileDescriptor_c891e73c7d2524a3, []int{2}
}

func (m *NewMarket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewMarket.Unmarshal(m, b)
}
func (m *NewMarket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewMarket.Marshal(b, m, deterministic)
}
func (m *NewMarket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewMarket.Merge(m, src)
}
func (m *NewMarket) XXX_Size() int {
	return xxx_messageInfo_NewMarket.Size(m)
}
func (m *NewMarket) XXX_DiscardUnknown() {
	xxx_messageInfo_NewMarket.DiscardUnknown(m)
}

var xxx_messageInfo_NewMarket proto.InternalMessageInfo

type UpdateNetwork struct {
	Changes              *NetworkConfiguration `protobuf:"bytes,1,opt,name=changes,proto3" json:"changes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateNetwork) Reset()         { *m = UpdateNetwork{} }
func (m *UpdateNetwork) String() string { return proto.CompactTextString(m) }
func (*UpdateNetwork) ProtoMessage()    {}
func (*UpdateNetwork) Descriptor() ([]byte, []int) {
	return fileDescriptor_c891e73c7d2524a3, []int{3}
}

func (m *UpdateNetwork) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateNetwork.Unmarshal(m, b)
}
func (m *UpdateNetwork) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateNetwork.Marshal(b, m, deterministic)
}
func (m *UpdateNetwork) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateNetwork.Merge(m, src)
}
func (m *UpdateNetwork) XXX_Size() int {
	return xxx_messageInfo_UpdateNetwork.Size(m)
}
func (m *UpdateNetwork) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateNetwork.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateNetwork proto.InternalMessageInfo

func (m *UpdateNetwork) GetChanges() *NetworkConfiguration {
	if m != nil {
		return m.Changes
	}
	return nil
}

type ProposalTerms struct {
	ClosingTimestamp      int64  `protobuf:"varint,1,opt,name=closingTimestamp,proto3" json:"closingTimestamp,omitempty"`
	EnactmentTimestamp    int64  `protobuf:"varint,2,opt,name=enactmentTimestamp,proto3" json:"enactmentTimestamp,omitempty"`
	MinParticipationStake uint64 `protobuf:"varint,3,opt,name=minParticipationStake,proto3" json:"minParticipationStake,omitempty"`
	// Types that are valid to be assigned to Change:
	//	*ProposalTerms_UpdateMarket
	//	*ProposalTerms_NewMarket
	//	*ProposalTerms_UpdateNetwork
	Change               isProposalTerms_Change `protobuf_oneof:"change"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ProposalTerms) Reset()         { *m = ProposalTerms{} }
func (m *ProposalTerms) String() string { return proto.CompactTextString(m) }
func (*ProposalTerms) ProtoMessage()    {}
func (*ProposalTerms) Descriptor() ([]byte, []int) {
	return fileDescriptor_c891e73c7d2524a3, []int{4}
}

func (m *ProposalTerms) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProposalTerms.Unmarshal(m, b)
}
func (m *ProposalTerms) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProposalTerms.Marshal(b, m, deterministic)
}
func (m *ProposalTerms) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalTerms.Merge(m, src)
}
func (m *ProposalTerms) XXX_Size() int {
	return xxx_messageInfo_ProposalTerms.Size(m)
}
func (m *ProposalTerms) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalTerms.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalTerms proto.InternalMessageInfo

func (m *ProposalTerms) GetClosingTimestamp() int64 {
	if m != nil {
		return m.ClosingTimestamp
	}
	return 0
}

func (m *ProposalTerms) GetEnactmentTimestamp() int64 {
	if m != nil {
		return m.EnactmentTimestamp
	}
	return 0
}

func (m *ProposalTerms) GetMinParticipationStake() uint64 {
	if m != nil {
		return m.MinParticipationStake
	}
	return 0
}

type isProposalTerms_Change interface {
	isProposalTerms_Change()
}

type ProposalTerms_UpdateMarket struct {
	UpdateMarket *UpdateMarket `protobuf:"bytes,101,opt,name=updateMarket,proto3,oneof"`
}

type ProposalTerms_NewMarket struct {
	NewMarket *NewMarket `protobuf:"bytes,102,opt,name=newMarket,proto3,oneof"`
}

type ProposalTerms_UpdateNetwork struct {
	UpdateNetwork *UpdateNetwork `protobuf:"bytes,103,opt,name=updateNetwork,proto3,oneof"`
}

func (*ProposalTerms_UpdateMarket) isProposalTerms_Change() {}

func (*ProposalTerms_NewMarket) isProposalTerms_Change() {}

func (*ProposalTerms_UpdateNetwork) isProposalTerms_Change() {}

func (m *ProposalTerms) GetChange() isProposalTerms_Change {
	if m != nil {
		return m.Change
	}
	return nil
}

func (m *ProposalTerms) GetUpdateMarket() *UpdateMarket {
	if x, ok := m.GetChange().(*ProposalTerms_UpdateMarket); ok {
		return x.UpdateMarket
	}
	return nil
}

func (m *ProposalTerms) GetNewMarket() *NewMarket {
	if x, ok := m.GetChange().(*ProposalTerms_NewMarket); ok {
		return x.NewMarket
	}
	return nil
}

func (m *ProposalTerms) GetUpdateNetwork() *UpdateNetwork {
	if x, ok := m.GetChange().(*ProposalTerms_UpdateNetwork); ok {
		return x.UpdateNetwork
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ProposalTerms) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ProposalTerms_UpdateMarket)(nil),
		(*ProposalTerms_NewMarket)(nil),
		(*ProposalTerms_UpdateNetwork)(nil),
	}
}

type Proposal struct {
	ID                   string         `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Reference            string         `protobuf:"bytes,2,opt,name=reference,proto3" json:"reference,omitempty"`
	PartyID              string         `protobuf:"bytes,3,opt,name=partyID,proto3" json:"partyID,omitempty"`
	State                Proposal_State `protobuf:"varint,4,opt,name=state,proto3,enum=vega.Proposal_State" json:"state,omitempty"`
	Timestamp            int64          `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Terms                *ProposalTerms `protobuf:"bytes,6,opt,name=terms,proto3" json:"terms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Proposal) Reset()         { *m = Proposal{} }
func (m *Proposal) String() string { return proto.CompactTextString(m) }
func (*Proposal) ProtoMessage()    {}
func (*Proposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_c891e73c7d2524a3, []int{5}
}

func (m *Proposal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proposal.Unmarshal(m, b)
}
func (m *Proposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proposal.Marshal(b, m, deterministic)
}
func (m *Proposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proposal.Merge(m, src)
}
func (m *Proposal) XXX_Size() int {
	return xxx_messageInfo_Proposal.Size(m)
}
func (m *Proposal) XXX_DiscardUnknown() {
	xxx_messageInfo_Proposal.DiscardUnknown(m)
}

var xxx_messageInfo_Proposal proto.InternalMessageInfo

func (m *Proposal) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Proposal) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *Proposal) GetPartyID() string {
	if m != nil {
		return m.PartyID
	}
	return ""
}

func (m *Proposal) GetState() Proposal_State {
	if m != nil {
		return m.State
	}
	return Proposal_FAILED
}

func (m *Proposal) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Proposal) GetTerms() *ProposalTerms {
	if m != nil {
		return m.Terms
	}
	return nil
}

type Vote struct {
	Voter                string     `protobuf:"bytes,1,opt,name=voter,proto3" json:"voter,omitempty"`
	Value                Vote_Value `protobuf:"varint,2,opt,name=value,proto3,enum=vega.Vote_Value" json:"value,omitempty"`
	ProposalID           string     `protobuf:"bytes,3,opt,name=proposalID,proto3" json:"proposalID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Vote) Reset()         { *m = Vote{} }
func (m *Vote) String() string { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()    {}
func (*Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_c891e73c7d2524a3, []int{6}
}

func (m *Vote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vote.Unmarshal(m, b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
}
func (m *Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vote.Merge(m, src)
}
func (m *Vote) XXX_Size() int {
	return xxx_messageInfo_Vote.Size(m)
}
func (m *Vote) XXX_DiscardUnknown() {
	xxx_messageInfo_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_Vote proto.InternalMessageInfo

func (m *Vote) GetVoter() string {
	if m != nil {
		return m.Voter
	}
	return ""
}

func (m *Vote) GetValue() Vote_Value {
	if m != nil {
		return m.Value
	}
	return Vote_NO
}

func (m *Vote) GetProposalID() string {
	if m != nil {
		return m.ProposalID
	}
	return ""
}

func init() {
	proto.RegisterEnum("vega.Proposal_State", Proposal_State_name, Proposal_State_value)
	proto.RegisterEnum("vega.Vote_Value", Vote_Value_name, Vote_Value_value)
	proto.RegisterType((*NetworkConfiguration)(nil), "vega.NetworkConfiguration")
	proto.RegisterType((*UpdateMarket)(nil), "vega.UpdateMarket")
	proto.RegisterType((*NewMarket)(nil), "vega.NewMarket")
	proto.RegisterType((*UpdateNetwork)(nil), "vega.UpdateNetwork")
	proto.RegisterType((*ProposalTerms)(nil), "vega.ProposalTerms")
	proto.RegisterType((*Proposal)(nil), "vega.Proposal")
	proto.RegisterType((*Vote)(nil), "vega.Vote")
}

func init() { proto.RegisterFile("proto/governance.proto", fileDescriptor_c891e73c7d2524a3) }

var fileDescriptor_c891e73c7d2524a3 = []byte{
	// 672 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x63, 0x27, 0xce, 0x9f, 0x69, 0x1b, 0xcc, 0x52, 0x50, 0x54, 0x21, 0x35, 0xf8, 0x80,
	0x7a, 0xa0, 0xb1, 0x48, 0x51, 0x85, 0x44, 0x2f, 0x4d, 0x6c, 0xd4, 0xa0, 0x36, 0x8d, 0x9c, 0x52,
	0x01, 0xb7, 0xad, 0xb3, 0x75, 0xad, 0xc4, 0xbb, 0x96, 0xbd, 0x49, 0xe0, 0xce, 0x81, 0x37, 0xe0,
	0x45, 0x78, 0x9f, 0x4a, 0xbd, 0xf2, 0x12, 0xc8, 0xbb, 0x8e, 0xe3, 0xa4, 0xe1, 0xd4, 0xee, 0x7c,
	0xdf, 0xb7, 0xb3, 0xf3, 0xf3, 0x28, 0xf0, 0x22, 0x8c, 0x18, 0x67, 0xa6, 0xc7, 0x66, 0x24, 0xa2,
	0x98, 0xba, 0xa4, 0x25, 0x0a, 0xa8, 0x34, 0x23, 0x1e, 0xde, 0x3b, 0xf6, 0x7c, 0x7e, 0x37, 0xbd,
	0x69, 0xb9, 0x2c, 0x30, 0x83, 0xb9, 0xcf, 0xc7, 0x6c, 0x6e, 0x7a, 0xec, 0x50, 0x58, 0x0e, 0x67,
	0x78, 0xe2, 0x8f, 0x30, 0x67, 0x51, 0x6c, 0x66, 0xff, 0xca, 0xb4, 0xf1, 0x53, 0x85, 0xdd, 0x3e,
	0xe1, 0x73, 0x16, 0x8d, 0xbb, 0x8c, 0xde, 0xfa, 0xde, 0x34, 0xc2, 0xdc, 0x67, 0x14, 0xbd, 0x81,
	0xa7, 0x81, 0x4f, 0xbb, 0x13, 0x16, 0x93, 0x1e, 0x1d, 0x12, 0x97, 0xd1, 0x51, 0xdc, 0x50, 0x9a,
	0xca, 0x41, 0xd1, 0x79, 0x2c, 0x08, 0x37, 0xfe, 0xbe, 0xe6, 0x56, 0x53, 0xf7, 0xba, 0x90, 0xde,
	0x6d, 0x53, 0xec, 0xf2, 0xa5, 0xbb, 0x98, 0xdd, 0xbd, 0x2a, 0xa4, 0x77, 0xaf, 0xb9, 0x4b, 0xd9,
	0xdd, 0x6b, 0xee, 0x77, 0xf0, 0x3c, 0xf0, 0xe9, 0x00, 0x47, 0xdc, 0x77, 0xfd, 0x50, 0xcc, 0x32,
	0xe4, 0x78, 0x4c, 0x1a, 0x5a, 0x53, 0x39, 0x28, 0x39, 0x9b, 0x45, 0xa3, 0x0e, 0xdb, 0x9f, 0xc3,
	0x11, 0xe6, 0xe4, 0x02, 0x47, 0x63, 0xc2, 0x8d, 0x2d, 0xa8, 0xf5, 0xc9, 0x3c, 0x3d, 0x5c, 0xc0,
	0x8e, 0x14, 0x53, 0x50, 0xe8, 0x04, 0x2a, 0xee, 0x1d, 0xa6, 0x1e, 0x91, 0x44, 0xb6, 0xda, 0x7b,
	0xad, 0xe4, 0x23, 0xb4, 0x36, 0x81, 0xec, 0x94, 0x1f, 0xee, 0xf7, 0xd5, 0xa6, 0xe2, 0x2c, 0x22,
	0xc6, 0x5f, 0x15, 0x76, 0x06, 0x11, 0x0b, 0x59, 0x8c, 0x27, 0x57, 0x24, 0x0a, 0x62, 0xd4, 0x06,
	0xdd, 0x9d, 0xb0, 0xd8, 0xa7, 0xde, 0x95, 0x1f, 0x90, 0x98, 0xe3, 0x20, 0x94, 0xa8, 0x65, 0x58,
	0x2f, 0x38, 0x8f, 0x74, 0x74, 0x0c, 0x88, 0x24, 0x93, 0x07, 0x84, 0xf2, 0x65, 0x4a, 0x5d, 0x49,
	0x6d, 0x70, 0xa0, 0x93, 0xff, 0xf1, 0x49, 0xf8, 0x97, 0xb2, 0xe8, 0x66, 0x13, 0x7a, 0x0f, 0xdb,
	0xd3, 0x1c, 0xa7, 0x06, 0x11, 0xe3, 0x23, 0x39, 0x7e, 0x9e, 0xe0, 0x59, 0xc1, 0x59, 0x71, 0x22,
	0x13, 0x6a, 0x74, 0x41, 0xb4, 0x71, 0x2b, 0x62, 0x4f, 0x16, 0xd4, 0xe6, 0x59, 0x66, 0xe9, 0x41,
	0x1f, 0x60, 0x67, 0x9a, 0xa7, 0xde, 0xf0, 0x44, 0xe8, 0x59, 0xbe, 0x57, 0x2a, 0x9d, 0x15, 0x9c,
	0x55, 0x6f, 0xa7, 0x0a, 0x65, 0x89, 0xdb, 0xf8, 0xa3, 0x42, 0x75, 0x41, 0x1b, 0xd5, 0x41, 0xed,
	0x59, 0x02, 0x6d, 0xcd, 0x51, 0x7b, 0x16, 0x7a, 0x09, 0xb5, 0x88, 0xdc, 0x92, 0x88, 0x50, 0x97,
	0x08, 0x76, 0x35, 0x67, 0x59, 0x40, 0x4d, 0xa8, 0x84, 0x38, 0xe2, 0x3f, 0x7a, 0x96, 0x80, 0x53,
	0x93, 0x70, 0xbe, 0x28, 0xce, 0xa2, 0x8c, 0x8e, 0x40, 0x8b, 0x39, 0xe6, 0x44, 0xac, 0x63, 0xbd,
	0xbd, 0x2b, 0xdf, 0xb6, 0x68, 0xd7, 0x1a, 0x26, 0x5a, 0xa7, 0xf2, 0x70, 0xbf, 0x5f, 0xfc, 0xa5,
	0x28, 0x8e, 0xf4, 0x26, 0x4d, 0x79, 0xf6, 0xc1, 0x34, 0xb1, 0xc7, 0xcb, 0x02, 0x7a, 0x0b, 0x1a,
	0x4f, 0x96, 0xa2, 0x51, 0xce, 0x8f, 0xbb, 0xb2, 0x2f, 0xd9, 0x4a, 0x49, 0xa7, 0xe1, 0x80, 0x26,
	0x3a, 0x21, 0x80, 0xf2, 0xc7, 0xd3, 0xde, 0xb9, 0x6d, 0xe9, 0x05, 0x54, 0x85, 0xd2, 0xe5, 0xc0,
	0xee, 0xeb, 0x4a, 0x52, 0x1d, 0x9c, 0x0e, 0x87, 0xb6, 0xa5, 0xab, 0x68, 0x1b, 0xaa, 0x8e, 0xfd,
	0xc9, 0xee, 0x5e, 0xd9, 0x96, 0x5e, 0x4c, 0x4e, 0x96, 0xdd, 0x3d, 0xef, 0xf5, 0x6d, 0x4b, 0x2f,
	0xa1, 0x2d, 0xa8, 0xd8, 0xfd, 0x53, 0x21, 0x69, 0xc6, 0x6f, 0x05, 0x4a, 0xd7, 0x4c, 0xbc, 0x56,
	0x9b, 0x31, 0x4e, 0x22, 0x49, 0x2d, 0x43, 0x20, 0x8b, 0xc8, 0x04, 0x6d, 0x86, 0x27, 0x53, 0x09,
	0xaf, 0xde, 0xd6, 0xe5, 0x6b, 0x93, 0x60, 0xeb, 0x3a, 0xa9, 0xe7, 0x86, 0x17, 0x3e, 0xf4, 0x1a,
	0x20, 0x4c, 0x67, 0x79, 0x84, 0x35, 0xa7, 0x18, 0x0d, 0xd0, 0xc4, 0x05, 0xa8, 0x0c, 0x6a, 0xff,
	0x52, 0x2f, 0xa0, 0x0a, 0x14, 0xbf, 0xda, 0x43, 0x5d, 0xe9, 0xbc, 0xfa, 0xb6, 0xef, 0xb2, 0x11,
	0x11, 0x9d, 0xc4, 0x6f, 0x98, 0xcb, 0x26, 0x2d, 0x9f, 0x99, 0xc9, 0xd9, 0x14, 0x85, 0x9b, 0xb2,
	0xf8, 0x73, 0xf4, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x89, 0xa7, 0x12, 0x29, 0x33, 0x05, 0x00, 0x00,
}
