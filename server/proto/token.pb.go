// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.1
// source: server/token.proto

package proto

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

type SaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                       // Идентификатор пользователя
	ClubId string `protobuf:"bytes,2,opt,name=club_id,json=clubId,proto3" json:"club_id,omitempty"` // Идентификатор клуба
}

func (x *SaveRequest) Reset() {
	*x = SaveRequest{}
	mi := &file_server_token_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveRequest) ProtoMessage() {}

func (x *SaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_token_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveRequest.ProtoReflect.Descriptor instead.
func (*SaveRequest) Descriptor() ([]byte, []int) {
	return file_server_token_proto_rawDescGZIP(), []int{0}
}

func (x *SaveRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SaveRequest) GetClubId() string {
	if x != nil {
		return x.ClubId
	}
	return ""
}

type SaveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SaveResponse) Reset() {
	*x = SaveResponse{}
	mi := &file_server_token_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SaveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveResponse) ProtoMessage() {}

func (x *SaveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_token_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveResponse.ProtoReflect.Descriptor instead.
func (*SaveResponse) Descriptor() ([]byte, []int) {
	return file_server_token_proto_rawDescGZIP(), []int{1}
}

func (x *SaveResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CheckMultipleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids    []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`                     // Список идентификаторов пользователей
	ClubId string   `protobuf:"bytes,2,opt,name=club_id,json=clubId,proto3" json:"club_id,omitempty"` // Идентификатор клуба для проверки
}

func (x *CheckMultipleRequest) Reset() {
	*x = CheckMultipleRequest{}
	mi := &file_server_token_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckMultipleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckMultipleRequest) ProtoMessage() {}

func (x *CheckMultipleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_token_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckMultipleRequest.ProtoReflect.Descriptor instead.
func (*CheckMultipleRequest) Descriptor() ([]byte, []int) {
	return file_server_token_proto_rawDescGZIP(), []int{2}
}

func (x *CheckMultipleRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *CheckMultipleRequest) GetClubId() string {
	if x != nil {
		return x.ClubId
	}
	return ""
}

type CheckMultipleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"` // Количество совпадений в базе
}

func (x *CheckMultipleResponse) Reset() {
	*x = CheckMultipleResponse{}
	mi := &file_server_token_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckMultipleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckMultipleResponse) ProtoMessage() {}

func (x *CheckMultipleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_token_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckMultipleResponse.ProtoReflect.Descriptor instead.
func (*CheckMultipleResponse) Descriptor() ([]byte, []int) {
	return file_server_token_proto_rawDescGZIP(), []int{3}
}

func (x *CheckMultipleResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_server_token_proto protoreflect.FileDescriptor

var file_server_token_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x36, 0x0a, 0x0b,
	0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63,
	0x6c, 0x75, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6c,
	0x75, 0x62, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x0c, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x41,
	0x0a, 0x14, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x62,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6c, 0x75, 0x62, 0x49,
	0x64, 0x22, 0x2d, 0x0a, 0x15, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x32, 0x8f, 0x01, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x31, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x70, 0x6c, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x14, 0x5a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_token_proto_rawDescOnce sync.Once
	file_server_token_proto_rawDescData = file_server_token_proto_rawDesc
)

func file_server_token_proto_rawDescGZIP() []byte {
	file_server_token_proto_rawDescOnce.Do(func() {
		file_server_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_token_proto_rawDescData)
	})
	return file_server_token_proto_rawDescData
}

var file_server_token_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_server_token_proto_goTypes = []any{
	(*SaveRequest)(nil),           // 0: server.SaveRequest
	(*SaveResponse)(nil),          // 1: server.SaveResponse
	(*CheckMultipleRequest)(nil),  // 2: server.CheckMultipleRequest
	(*CheckMultipleResponse)(nil), // 3: server.CheckMultipleResponse
}
var file_server_token_proto_depIdxs = []int32{
	0, // 0: server.TokenService.Save:input_type -> server.SaveRequest
	2, // 1: server.TokenService.CheckMultiple:input_type -> server.CheckMultipleRequest
	1, // 2: server.TokenService.Save:output_type -> server.SaveResponse
	3, // 3: server.TokenService.CheckMultiple:output_type -> server.CheckMultipleResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_server_token_proto_init() }
func file_server_token_proto_init() {
	if File_server_token_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_server_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_token_proto_goTypes,
		DependencyIndexes: file_server_token_proto_depIdxs,
		MessageInfos:      file_server_token_proto_msgTypes,
	}.Build()
	File_server_token_proto = out.File
	file_server_token_proto_rawDesc = nil
	file_server_token_proto_goTypes = nil
	file_server_token_proto_depIdxs = nil
}
