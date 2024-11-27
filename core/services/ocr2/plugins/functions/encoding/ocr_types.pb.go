// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: core/services/ocr2/plugins/functions/encoding/ocr_types.proto

package encoding

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

// These protos are used internally by the OCR2 reporting plugin to
// pass data between initial phases. Report is ABI-encoded.
type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestIDs [][]byte `protobuf:"bytes,1,rep,name=requestIDs,proto3" json:"requestIDs,omitempty"`
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_rawDescGZIP(), []int{0}
}

func (x *Query) GetRequestIDs() [][]byte {
	if x != nil {
		return x.RequestIDs
	}
	return nil
}

type Observation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProcessedRequests []*ProcessedRequest `protobuf:"bytes,1,rep,name=processedRequests,proto3" json:"processedRequests,omitempty"`
}

func (x *Observation) Reset() {
	*x = Observation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Observation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Observation) ProtoMessage() {}

func (x *Observation) ProtoReflect() protoreflect.Message {
	mi := &file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Observation.ProtoReflect.Descriptor instead.
func (*Observation) Descriptor() ([]byte, []int) {
	return file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_rawDescGZIP(), []int{1}
}

func (x *Observation) GetProcessedRequests() []*ProcessedRequest {
	if x != nil {
		return x.ProcessedRequests
	}
	return nil
}

type ProcessedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestID           []byte `protobuf:"bytes,1,opt,name=requestID,proto3" json:"requestID,omitempty"`
	Result              []byte `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	Error               []byte `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	CallbackGasLimit    uint32 `protobuf:"varint,4,opt,name=callbackGasLimit,proto3" json:"callbackGasLimit,omitempty"`
	CoordinatorContract []byte `protobuf:"bytes,5,opt,name=coordinatorContract,proto3" json:"coordinatorContract,omitempty"`
	OnchainMetadata     []byte `protobuf:"bytes,6,opt,name=onchainMetadata,proto3" json:"onchainMetadata,omitempty"`
}

func (x *ProcessedRequest) Reset() {
	*x = ProcessedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessedRequest) ProtoMessage() {}

func (x *ProcessedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessedRequest.ProtoReflect.Descriptor instead.
func (*ProcessedRequest) Descriptor() ([]byte, []int) {
	return file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_rawDescGZIP(), []int{2}
}

func (x *ProcessedRequest) GetRequestID() []byte {
	if x != nil {
		return x.RequestID
	}
	return nil
}

func (x *ProcessedRequest) GetResult() []byte {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *ProcessedRequest) GetError() []byte {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *ProcessedRequest) GetCallbackGasLimit() uint32 {
	if x != nil {
		return x.CallbackGasLimit
	}
	return 0
}

func (x *ProcessedRequest) GetCoordinatorContract() []byte {
	if x != nil {
		return x.CoordinatorContract
	}
	return nil
}

func (x *ProcessedRequest) GetOnchainMetadata() []byte {
	if x != nil {
		return x.OnchainMetadata
	}
	return nil
}

var File_core_services_ocr2_plugins_functions_encoding_ocr_types_proto protoreflect.FileDescriptor

var file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x6f, 0x63, 0x72, 0x32, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x66, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f,
	0x6f, 0x63, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x27, 0x0a, 0x05, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x44, 0x73, 0x22, 0x57, 0x0a, 0x0b, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x48, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65,
	0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0xe6, 0x01, 0x0a, 0x10,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2a, 0x0a, 0x10,
	0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x47, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x47, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x63, 0x6f, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x13, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74,
	0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x6f, 0x6e,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0f, 0x6f, 0x6e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x42, 0x2f, 0x5a, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x63, 0x72, 0x32, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x73, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x6e, 0x63,
	0x6f, 0x64, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_rawDescOnce sync.Once
	file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_rawDescData = file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_rawDesc
)

func file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_rawDescGZIP() []byte {
	file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_rawDescOnce.Do(func() {
		file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_rawDescData)
	})
	return file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_rawDescData
}

var file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_goTypes = []interface{}{
	(*Query)(nil),            // 0: encoding.Query
	(*Observation)(nil),      // 1: encoding.Observation
	(*ProcessedRequest)(nil), // 2: encoding.ProcessedRequest
}
var file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_depIdxs = []int32{
	2, // 0: encoding.Observation.processedRequests:type_name -> encoding.ProcessedRequest
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_init() }
func file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_init() {
	if File_core_services_ocr2_plugins_functions_encoding_ocr_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query); i {
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
		file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Observation); i {
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
		file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessedRequest); i {
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
			RawDescriptor: file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_goTypes,
		DependencyIndexes: file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_depIdxs,
		MessageInfos:      file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_msgTypes,
	}.Build()
	File_core_services_ocr2_plugins_functions_encoding_ocr_types_proto = out.File
	file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_rawDesc = nil
	file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_goTypes = nil
	file_core_services_ocr2_plugins_functions_encoding_ocr_types_proto_depIdxs = nil
}
