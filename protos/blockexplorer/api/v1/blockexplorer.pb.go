// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: blockexplorer/api/v1/blockexplorer.proto

package v1

import (
	v1 "code.vegaprotocol.io/vega/protos/vega/commands/v1"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// node information
type InfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InfoRequest) Reset() {
	*x = InfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockexplorer_api_v1_blockexplorer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoRequest) ProtoMessage() {}

func (x *InfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blockexplorer_api_v1_blockexplorer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoRequest.ProtoReflect.Descriptor instead.
func (*InfoRequest) Descriptor() ([]byte, []int) {
	return file_blockexplorer_api_v1_blockexplorer_proto_rawDescGZIP(), []int{0}
}

type InfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Semver formatted version of the data node
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// Commit hash from which the data node was built
	CommitHash string `protobuf:"bytes,2,opt,name=commit_hash,json=commitHash,proto3" json:"commit_hash,omitempty"`
}

func (x *InfoResponse) Reset() {
	*x = InfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockexplorer_api_v1_blockexplorer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResponse) ProtoMessage() {}

func (x *InfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blockexplorer_api_v1_blockexplorer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResponse.ProtoReflect.Descriptor instead.
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return file_blockexplorer_api_v1_blockexplorer_proto_rawDescGZIP(), []int{1}
}

func (x *InfoResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *InfoResponse) GetCommitHash() string {
	if x != nil {
		return x.CommitHash
	}
	return ""
}

type GetTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Hash of the transaction
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *GetTransactionRequest) Reset() {
	*x = GetTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockexplorer_api_v1_blockexplorer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionRequest) ProtoMessage() {}

func (x *GetTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blockexplorer_api_v1_blockexplorer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionRequest.ProtoReflect.Descriptor instead.
func (*GetTransactionRequest) Descriptor() ([]byte, []int) {
	return file_blockexplorer_api_v1_blockexplorer_proto_rawDescGZIP(), []int{2}
}

func (x *GetTransactionRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type GetTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Transaction corresponding to the hash
	Transaction *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *GetTransactionResponse) Reset() {
	*x = GetTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockexplorer_api_v1_blockexplorer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionResponse) ProtoMessage() {}

func (x *GetTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blockexplorer_api_v1_blockexplorer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionResponse.ProtoReflect.Descriptor instead.
func (*GetTransactionResponse) Descriptor() ([]byte, []int) {
	return file_blockexplorer_api_v1_blockexplorer_proto_rawDescGZIP(), []int{3}
}

func (x *GetTransactionResponse) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

type ListTransactionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of transactions to be returned from the blockchain.
	// This is deprecated, use first and last instead.
	//
	// Deprecated: Do not use.
	Limit uint32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// Optional cursor to paginate the request
	Before *string `protobuf:"bytes,2,opt,name=before,proto3,oneof" json:"before,omitempty"`
	// Optional cursor to paginate the request
	After *string `protobuf:"bytes,3,opt,name=after,proto3,oneof" json:"after,omitempty"`
	// Filters to apply to the request
	Filters map[string]string `protobuf:"bytes,4,rep,name=filters,proto3" json:"filters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Transaction command types filter, for listing transactions with specified command types
	CmdTypes []string `protobuf:"bytes,5,rep,name=cmd_types,json=cmdTypes,proto3" json:"cmd_types,omitempty"`
	// Transaction command types exclusion filter, for listing all the transactions except the ones with specified command types
	ExcludeCmdTypes []string `protobuf:"bytes,6,rep,name=exclude_cmd_types,json=excludeCmdTypes,proto3" json:"exclude_cmd_types,omitempty"`
	// Party IDs filter, can be sender or receiver
	Parties []string `protobuf:"bytes,7,rep,name=parties,proto3" json:"parties,omitempty"`
	// Number of transactions to be returned from the blockchain. Use in conjunction with the `after` cursor to paginate forwards.
	// On its own, this will return the first `first` transactions.
	First uint32 `protobuf:"varint,8,opt,name=first,proto3" json:"first,omitempty"`
	// Number of transactions to be returned from the blockchain. Use in conjunction with the `before` cursor to paginate backwards.
	// On its own, this will return the last `last` transactions.
	Last uint32 `protobuf:"varint,9,opt,name=last,proto3" json:"last,omitempty"`
}

func (x *ListTransactionsRequest) Reset() {
	*x = ListTransactionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockexplorer_api_v1_blockexplorer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTransactionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTransactionsRequest) ProtoMessage() {}

func (x *ListTransactionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blockexplorer_api_v1_blockexplorer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTransactionsRequest.ProtoReflect.Descriptor instead.
func (*ListTransactionsRequest) Descriptor() ([]byte, []int) {
	return file_blockexplorer_api_v1_blockexplorer_proto_rawDescGZIP(), []int{4}
}

// Deprecated: Do not use.
func (x *ListTransactionsRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListTransactionsRequest) GetBefore() string {
	if x != nil && x.Before != nil {
		return *x.Before
	}
	return ""
}

func (x *ListTransactionsRequest) GetAfter() string {
	if x != nil && x.After != nil {
		return *x.After
	}
	return ""
}

func (x *ListTransactionsRequest) GetFilters() map[string]string {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *ListTransactionsRequest) GetCmdTypes() []string {
	if x != nil {
		return x.CmdTypes
	}
	return nil
}

func (x *ListTransactionsRequest) GetExcludeCmdTypes() []string {
	if x != nil {
		return x.ExcludeCmdTypes
	}
	return nil
}

func (x *ListTransactionsRequest) GetParties() []string {
	if x != nil {
		return x.Parties
	}
	return nil
}

func (x *ListTransactionsRequest) GetFirst() uint32 {
	if x != nil {
		return x.First
	}
	return 0
}

func (x *ListTransactionsRequest) GetLast() uint32 {
	if x != nil {
		return x.Last
	}
	return 0
}

type ListTransactionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Transaction corresponding to the specific request and filters
	Transactions []*Transaction `protobuf:"bytes,3,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *ListTransactionsResponse) Reset() {
	*x = ListTransactionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockexplorer_api_v1_blockexplorer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTransactionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTransactionsResponse) ProtoMessage() {}

func (x *ListTransactionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blockexplorer_api_v1_blockexplorer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTransactionsResponse.ProtoReflect.Descriptor instead.
func (*ListTransactionsResponse) Descriptor() ([]byte, []int) {
	return file_blockexplorer_api_v1_blockexplorer_proto_rawDescGZIP(), []int{5}
}

func (x *ListTransactionsResponse) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Height of the block the transaction was found in
	Block uint64 `protobuf:"varint,1,opt,name=block,proto3" json:"block,omitempty"`
	// Index of the transaction in the block
	Index uint32 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	// Hash of the transaction
	Hash string `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	// Vega public key of the transaction's submitter
	Submitter string `protobuf:"bytes,4,opt,name=submitter,proto3" json:"submitter,omitempty"`
	// Type of transaction
	Type string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	// Results code of the transaction. 0 indicates the transaction was successful
	Code uint32 `protobuf:"varint,6,opt,name=code,proto3" json:"code,omitempty"`
	// Cursor for this transaction. This is used for paginating results
	Cursor string `protobuf:"bytes,7,opt,name=cursor,proto3" json:"cursor,omitempty"`
	// Actual command of the transaction
	Command *v1.InputData `protobuf:"bytes,8,opt,name=command,proto3" json:"command,omitempty"`
	// Signature generated by the submitter for the transaction
	Signature *v1.Signature `protobuf:"bytes,9,opt,name=signature,proto3" json:"signature,omitempty"`
	// Optional error happening when processing / checking the transaction
	// This should be set if error code is not 0
	Error *string `protobuf:"bytes,10,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockexplorer_api_v1_blockexplorer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_blockexplorer_api_v1_blockexplorer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_blockexplorer_api_v1_blockexplorer_proto_rawDescGZIP(), []int{6}
}

func (x *Transaction) GetBlock() uint64 {
	if x != nil {
		return x.Block
	}
	return 0
}

func (x *Transaction) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Transaction) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Transaction) GetSubmitter() string {
	if x != nil {
		return x.Submitter
	}
	return ""
}

func (x *Transaction) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Transaction) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Transaction) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

func (x *Transaction) GetCommand() *v1.InputData {
	if x != nil {
		return x.Command
	}
	return nil
}

func (x *Transaction) GetSignature() *v1.Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *Transaction) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

var File_blockexplorer_api_v1_blockexplorer_proto protoreflect.FileDescriptor

var file_blockexplorer_api_v1_blockexplorer_proto_rawDesc = []byte{
	0x0a, 0x28, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x78, 0x70, 0x6c,
	0x6f, 0x72, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x76, 0x65, 0x67, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x76, 0x65, 0x67, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x49, 0x0a, 0x0c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x48, 0x61, 0x73,
	0x68, 0x22, 0x31, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x22, 0x5d, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43,
	0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x78, 0x70, 0x6c, 0x6f,
	0x72, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x9f, 0x03, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x02,
	0x18, 0x01, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x62, 0x65, 0x66,
	0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x62, 0x65, 0x66,
	0x6f, 0x72, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x88, 0x01,
	0x01, 0x12, 0x54, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6d, 0x64, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6d, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f,
	0x63, 0x6d, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x43, 0x6d, 0x64, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x6c, 0x61, 0x73, 0x74, 0x1a, 0x3a, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x61, 0x66, 0x74, 0x65, 0x72, 0x22, 0x61, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x45, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65,
	0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xc2, 0x02, 0x0a, 0x0b, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x35, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x39, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xc9, 0x02,
	0x0a, 0x14, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x78, 0x70,
	0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x73, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2d, 0x2e, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x21, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x78, 0x70,
	0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x78, 0x5a, 0x35, 0x63, 0x6f, 0x64,
	0x65, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x69,
	0x6f, 0x2f, 0x76, 0x65, 0x67, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x92, 0x41, 0x3e, 0x12, 0x23, 0x0a, 0x18, 0x56, 0x65, 0x67, 0x61, 0x20, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x20, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x20, 0x41, 0x50, 0x49,
	0x73, 0x32, 0x07, 0x76, 0x30, 0x2e, 0x37, 0x31, 0x2e, 0x30, 0x1a, 0x13, 0x6c, 0x62, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x6e, 0x65, 0x74, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x78, 0x79, 0x7a, 0x2a,
	0x02, 0x01, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blockexplorer_api_v1_blockexplorer_proto_rawDescOnce sync.Once
	file_blockexplorer_api_v1_blockexplorer_proto_rawDescData = file_blockexplorer_api_v1_blockexplorer_proto_rawDesc
)

func file_blockexplorer_api_v1_blockexplorer_proto_rawDescGZIP() []byte {
	file_blockexplorer_api_v1_blockexplorer_proto_rawDescOnce.Do(func() {
		file_blockexplorer_api_v1_blockexplorer_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockexplorer_api_v1_blockexplorer_proto_rawDescData)
	})
	return file_blockexplorer_api_v1_blockexplorer_proto_rawDescData
}

var file_blockexplorer_api_v1_blockexplorer_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_blockexplorer_api_v1_blockexplorer_proto_goTypes = []interface{}{
	(*InfoRequest)(nil),              // 0: blockexplorer.api.v1.InfoRequest
	(*InfoResponse)(nil),             // 1: blockexplorer.api.v1.InfoResponse
	(*GetTransactionRequest)(nil),    // 2: blockexplorer.api.v1.GetTransactionRequest
	(*GetTransactionResponse)(nil),   // 3: blockexplorer.api.v1.GetTransactionResponse
	(*ListTransactionsRequest)(nil),  // 4: blockexplorer.api.v1.ListTransactionsRequest
	(*ListTransactionsResponse)(nil), // 5: blockexplorer.api.v1.ListTransactionsResponse
	(*Transaction)(nil),              // 6: blockexplorer.api.v1.Transaction
	nil,                              // 7: blockexplorer.api.v1.ListTransactionsRequest.FiltersEntry
	(*v1.InputData)(nil),             // 8: vega.commands.v1.InputData
	(*v1.Signature)(nil),             // 9: vega.commands.v1.Signature
}
var file_blockexplorer_api_v1_blockexplorer_proto_depIdxs = []int32{
	6, // 0: blockexplorer.api.v1.GetTransactionResponse.transaction:type_name -> blockexplorer.api.v1.Transaction
	7, // 1: blockexplorer.api.v1.ListTransactionsRequest.filters:type_name -> blockexplorer.api.v1.ListTransactionsRequest.FiltersEntry
	6, // 2: blockexplorer.api.v1.ListTransactionsResponse.transactions:type_name -> blockexplorer.api.v1.Transaction
	8, // 3: blockexplorer.api.v1.Transaction.command:type_name -> vega.commands.v1.InputData
	9, // 4: blockexplorer.api.v1.Transaction.signature:type_name -> vega.commands.v1.Signature
	2, // 5: blockexplorer.api.v1.BlockExplorerService.GetTransaction:input_type -> blockexplorer.api.v1.GetTransactionRequest
	4, // 6: blockexplorer.api.v1.BlockExplorerService.ListTransactions:input_type -> blockexplorer.api.v1.ListTransactionsRequest
	0, // 7: blockexplorer.api.v1.BlockExplorerService.Info:input_type -> blockexplorer.api.v1.InfoRequest
	3, // 8: blockexplorer.api.v1.BlockExplorerService.GetTransaction:output_type -> blockexplorer.api.v1.GetTransactionResponse
	5, // 9: blockexplorer.api.v1.BlockExplorerService.ListTransactions:output_type -> blockexplorer.api.v1.ListTransactionsResponse
	1, // 10: blockexplorer.api.v1.BlockExplorerService.Info:output_type -> blockexplorer.api.v1.InfoResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_blockexplorer_api_v1_blockexplorer_proto_init() }
func file_blockexplorer_api_v1_blockexplorer_proto_init() {
	if File_blockexplorer_api_v1_blockexplorer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blockexplorer_api_v1_blockexplorer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blockexplorer_api_v1_blockexplorer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blockexplorer_api_v1_blockexplorer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blockexplorer_api_v1_blockexplorer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blockexplorer_api_v1_blockexplorer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTransactionsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blockexplorer_api_v1_blockexplorer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTransactionsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blockexplorer_api_v1_blockexplorer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_blockexplorer_api_v1_blockexplorer_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_blockexplorer_api_v1_blockexplorer_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blockexplorer_api_v1_blockexplorer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blockexplorer_api_v1_blockexplorer_proto_goTypes,
		DependencyIndexes: file_blockexplorer_api_v1_blockexplorer_proto_depIdxs,
		MessageInfos:      file_blockexplorer_api_v1_blockexplorer_proto_msgTypes,
	}.Build()
	File_blockexplorer_api_v1_blockexplorer_proto = out.File
	file_blockexplorer_api_v1_blockexplorer_proto_rawDesc = nil
	file_blockexplorer_api_v1_blockexplorer_proto_goTypes = nil
	file_blockexplorer_api_v1_blockexplorer_proto_depIdxs = nil
}
