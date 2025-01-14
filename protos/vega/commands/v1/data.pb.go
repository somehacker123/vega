// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: vega/commands/v1/data.proto

package v1

import (
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

// Supported oracle sources
type OracleDataSubmission_OracleSource int32

const (
	// Default value
	OracleDataSubmission_ORACLE_SOURCE_UNSPECIFIED OracleDataSubmission_OracleSource = 0
	// Specifies that the payload will be base64 encoded JSON conforming to the Open Oracle standard
	OracleDataSubmission_ORACLE_SOURCE_OPEN_ORACLE OracleDataSubmission_OracleSource = 1
	// Specifies that the payload will be base64 encoded JSON, but does not specify the shape of the data
	OracleDataSubmission_ORACLE_SOURCE_JSON OracleDataSubmission_OracleSource = 2
	// Specifies that the payload will be base64 encoded JSON conforming to the ETH standard
	OracleDataSubmission_ORACLE_SOURCE_ETHEREUM OracleDataSubmission_OracleSource = 3
)

// Enum value maps for OracleDataSubmission_OracleSource.
var (
	OracleDataSubmission_OracleSource_name = map[int32]string{
		0: "ORACLE_SOURCE_UNSPECIFIED",
		1: "ORACLE_SOURCE_OPEN_ORACLE",
		2: "ORACLE_SOURCE_JSON",
		3: "ORACLE_SOURCE_ETHEREUM",
	}
	OracleDataSubmission_OracleSource_value = map[string]int32{
		"ORACLE_SOURCE_UNSPECIFIED": 0,
		"ORACLE_SOURCE_OPEN_ORACLE": 1,
		"ORACLE_SOURCE_JSON":        2,
		"ORACLE_SOURCE_ETHEREUM":    3,
	}
)

func (x OracleDataSubmission_OracleSource) Enum() *OracleDataSubmission_OracleSource {
	p := new(OracleDataSubmission_OracleSource)
	*p = x
	return p
}

func (x OracleDataSubmission_OracleSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OracleDataSubmission_OracleSource) Descriptor() protoreflect.EnumDescriptor {
	return file_vega_commands_v1_data_proto_enumTypes[0].Descriptor()
}

func (OracleDataSubmission_OracleSource) Type() protoreflect.EnumType {
	return &file_vega_commands_v1_data_proto_enumTypes[0]
}

func (x OracleDataSubmission_OracleSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OracleDataSubmission_OracleSource.Descriptor instead.
func (OracleDataSubmission_OracleSource) EnumDescriptor() ([]byte, []int) {
	return file_vega_commands_v1_data_proto_rawDescGZIP(), []int{0, 0}
}

// Command to submit new Oracle data from third party providers
type OracleDataSubmission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Source from which the data is coming from. Must be base64 encoded.
	// Oracle data is a type of external data source data.
	Source OracleDataSubmission_OracleSource `protobuf:"varint,1,opt,name=source,proto3,enum=vega.commands.v1.OracleDataSubmission_OracleSource" json:"source,omitempty"`
	// Data provided by the data source
	// In the case of Open Oracle - it will be the entire object - it will contain messages, signatures and price data.
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *OracleDataSubmission) Reset() {
	*x = OracleDataSubmission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_commands_v1_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OracleDataSubmission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OracleDataSubmission) ProtoMessage() {}

func (x *OracleDataSubmission) ProtoReflect() protoreflect.Message {
	mi := &file_vega_commands_v1_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OracleDataSubmission.ProtoReflect.Descriptor instead.
func (*OracleDataSubmission) Descriptor() ([]byte, []int) {
	return file_vega_commands_v1_data_proto_rawDescGZIP(), []int{0}
}

func (x *OracleDataSubmission) GetSource() OracleDataSubmission_OracleSource {
	if x != nil {
		return x.Source
	}
	return OracleDataSubmission_ORACLE_SOURCE_UNSPECIFIED
}

func (x *OracleDataSubmission) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_vega_commands_v1_data_proto protoreflect.FileDescriptor

var file_vega_commands_v1_data_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x76, 0x65, 0x67, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x76,
	0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x22,
	0x80, 0x02, 0x0a, 0x14, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x61, 0x63,
	0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22,
	0x80, 0x01, 0x0a, 0x0c, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x1d, 0x0a, 0x19, 0x4f, 0x52, 0x41, 0x43, 0x4c, 0x45, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x1d, 0x0a, 0x19, 0x4f, 0x52, 0x41, 0x43, 0x4c, 0x45, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x4f, 0x52, 0x41, 0x43, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x16,
	0x0a, 0x12, 0x4f, 0x52, 0x41, 0x43, 0x4c, 0x45, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f,
	0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x4f, 0x52, 0x41, 0x43, 0x4c, 0x45,
	0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x45, 0x54, 0x48, 0x45, 0x52, 0x45, 0x55, 0x4d,
	0x10, 0x03, 0x42, 0x33, 0x5a, 0x31, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x65, 0x67, 0x61, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x65, 0x67, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vega_commands_v1_data_proto_rawDescOnce sync.Once
	file_vega_commands_v1_data_proto_rawDescData = file_vega_commands_v1_data_proto_rawDesc
)

func file_vega_commands_v1_data_proto_rawDescGZIP() []byte {
	file_vega_commands_v1_data_proto_rawDescOnce.Do(func() {
		file_vega_commands_v1_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_vega_commands_v1_data_proto_rawDescData)
	})
	return file_vega_commands_v1_data_proto_rawDescData
}

var file_vega_commands_v1_data_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_vega_commands_v1_data_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_vega_commands_v1_data_proto_goTypes = []interface{}{
	(OracleDataSubmission_OracleSource)(0), // 0: vega.commands.v1.OracleDataSubmission.OracleSource
	(*OracleDataSubmission)(nil),           // 1: vega.commands.v1.OracleDataSubmission
}
var file_vega_commands_v1_data_proto_depIdxs = []int32{
	0, // 0: vega.commands.v1.OracleDataSubmission.source:type_name -> vega.commands.v1.OracleDataSubmission.OracleSource
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_vega_commands_v1_data_proto_init() }
func file_vega_commands_v1_data_proto_init() {
	if File_vega_commands_v1_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vega_commands_v1_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OracleDataSubmission); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vega_commands_v1_data_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vega_commands_v1_data_proto_goTypes,
		DependencyIndexes: file_vega_commands_v1_data_proto_depIdxs,
		EnumInfos:         file_vega_commands_v1_data_proto_enumTypes,
		MessageInfos:      file_vega_commands_v1_data_proto_msgTypes,
	}.Build()
	File_vega_commands_v1_data_proto = out.File
	file_vega_commands_v1_data_proto_rawDesc = nil
	file_vega_commands_v1_data_proto_goTypes = nil
	file_vega_commands_v1_data_proto_depIdxs = nil
}
