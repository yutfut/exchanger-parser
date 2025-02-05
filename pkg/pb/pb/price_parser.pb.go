// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: price_parser.proto

package exchanger_parser_pb

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

type GetRateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rate string `protobuf:"bytes,1,opt,name=rate,proto3" json:"rate,omitempty"`
}

func (x *GetRateResponse) Reset() {
	*x = GetRateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_price_parser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRateResponse) ProtoMessage() {}

func (x *GetRateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_price_parser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRateResponse.ProtoReflect.Descriptor instead.
func (*GetRateResponse) Descriptor() ([]byte, []int) {
	return file_price_parser_proto_rawDescGZIP(), []int{0}
}

func (x *GetRateResponse) GetRate() string {
	if x != nil {
		return x.Rate
	}
	return ""
}

type GetRateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exchange              uint32 `protobuf:"varint,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	ExchangersConditional uint32 `protobuf:"varint,2,opt,name=exchangersConditional,proto3" json:"exchangersConditional,omitempty"`
}

func (x *GetRateRequest) Reset() {
	*x = GetRateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_price_parser_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRateRequest) ProtoMessage() {}

func (x *GetRateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_price_parser_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRateRequest.ProtoReflect.Descriptor instead.
func (*GetRateRequest) Descriptor() ([]byte, []int) {
	return file_price_parser_proto_rawDescGZIP(), []int{1}
}

func (x *GetRateRequest) GetExchange() uint32 {
	if x != nil {
		return x.Exchange
	}
	return 0
}

func (x *GetRateRequest) GetExchangersConditional() uint32 {
	if x != nil {
		return x.ExchangersConditional
	}
	return 0
}

var File_price_parser_proto protoreflect.FileDescriptor

var file_price_parser_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x73,
	0x65, 0x72, 0x22, 0x25, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x22, 0x62, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x15, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x72, 0x73, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x72, 0x73, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x32, 0x5c, 0x0a,
	0x12, 0x50, 0x72, 0x69, 0x63, 0x65, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1c,
	0x2e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x18, 0x5a, 0x16, 0x2e,
	0x2f, 0x3b, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x72, 0x73,
	0x65, 0x72, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_price_parser_proto_rawDescOnce sync.Once
	file_price_parser_proto_rawDescData = file_price_parser_proto_rawDesc
)

func file_price_parser_proto_rawDescGZIP() []byte {
	file_price_parser_proto_rawDescOnce.Do(func() {
		file_price_parser_proto_rawDescData = protoimpl.X.CompressGZIP(file_price_parser_proto_rawDescData)
	})
	return file_price_parser_proto_rawDescData
}

var file_price_parser_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_price_parser_proto_goTypes = []interface{}{
	(*GetRateResponse)(nil), // 0: price_parser.GetRateResponse
	(*GetRateRequest)(nil),  // 1: price_parser.GetRateRequest
}
var file_price_parser_proto_depIdxs = []int32{
	1, // 0: price_parser.PriceParserService.GetRate:input_type -> price_parser.GetRateRequest
	0, // 1: price_parser.PriceParserService.GetRate:output_type -> price_parser.GetRateResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_price_parser_proto_init() }
func file_price_parser_proto_init() {
	if File_price_parser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_price_parser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRateResponse); i {
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
		file_price_parser_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRateRequest); i {
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
			RawDescriptor: file_price_parser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_price_parser_proto_goTypes,
		DependencyIndexes: file_price_parser_proto_depIdxs,
		MessageInfos:      file_price_parser_proto_msgTypes,
	}.Build()
	File_price_parser_proto = out.File
	file_price_parser_proto_rawDesc = nil
	file_price_parser_proto_goTypes = nil
	file_price_parser_proto_depIdxs = nil
}
