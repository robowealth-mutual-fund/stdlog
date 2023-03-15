// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: masterData/country/v1/message/response.proto

package message

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

type ImportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *ImportResponse) Reset() {
	*x = ImportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_masterData_country_v1_message_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportResponse) ProtoMessage() {}

func (x *ImportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_masterData_country_v1_message_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportResponse.ProtoReflect.Descriptor instead.
func (*ImportResponse) Descriptor() ([]byte, []int) {
	return file_masterData_country_v1_message_response_proto_rawDescGZIP(), []int{0}
}

func (x *ImportResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items      []*ListItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Size       int64       `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Total      int64       `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	TotalPages int64       `protobuf:"varint,4,opt,name=total_pages,proto3" json:"total_pages,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_masterData_country_v1_message_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_masterData_country_v1_message_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_masterData_country_v1_message_response_proto_rawDescGZIP(), []int{1}
}

func (x *ListResponse) GetItems() []*ListItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListResponse) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListResponse) GetTotalPages() int64 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

type ListItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	NameTh   string `protobuf:"bytes,2,opt,name=name_th,proto3" json:"name_th,omitempty"`
	NameEn   string `protobuf:"bytes,3,opt,name=name_en,proto3" json:"name_en,omitempty"`
	Version  string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	IsActive bool   `protobuf:"varint,5,opt,name=is_active,proto3" json:"is_active,omitempty"`
}

func (x *ListItem) Reset() {
	*x = ListItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_masterData_country_v1_message_response_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListItem) ProtoMessage() {}

func (x *ListItem) ProtoReflect() protoreflect.Message {
	mi := &file_masterData_country_v1_message_response_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListItem.ProtoReflect.Descriptor instead.
func (*ListItem) Descriptor() ([]byte, []int) {
	return file_masterData_country_v1_message_response_proto_rawDescGZIP(), []int{2}
}

func (x *ListItem) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ListItem) GetNameTh() string {
	if x != nil {
		return x.NameTh
	}
	return ""
}

func (x *ListItem) GetNameEn() string {
	if x != nil {
		return x.NameEn
	}
	return ""
}

func (x *ListItem) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ListItem) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

var File_masterData_country_v1_message_response_proto protoreflect.FileDescriptor

var file_masterData_country_v1_message_response_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2a, 0x0a, 0x0e, 0x49,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x97, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x73, 0x22, 0x8a, 0x01, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07,
	0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e,
	0x61, 0x6d, 0x65, 0x5f, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x42, 0x4c,
	0x5a, 0x4a, 0x67, 0x69, 0x74, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x64, 0x65, 0x76, 0x2e, 0x63, 0x6f,
	0x2f, 0x72, 0x6f, 0x61, 0x2d, 0x76, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_masterData_country_v1_message_response_proto_rawDescOnce sync.Once
	file_masterData_country_v1_message_response_proto_rawDescData = file_masterData_country_v1_message_response_proto_rawDesc
)

func file_masterData_country_v1_message_response_proto_rawDescGZIP() []byte {
	file_masterData_country_v1_message_response_proto_rawDescOnce.Do(func() {
		file_masterData_country_v1_message_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_masterData_country_v1_message_response_proto_rawDescData)
	})
	return file_masterData_country_v1_message_response_proto_rawDescData
}

var file_masterData_country_v1_message_response_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_masterData_country_v1_message_response_proto_goTypes = []interface{}{
	(*ImportResponse)(nil), // 0: master.data.country.message.ImportResponse
	(*ListResponse)(nil),   // 1: master.data.country.message.ListResponse
	(*ListItem)(nil),       // 2: master.data.country.message.ListItem
}
var file_masterData_country_v1_message_response_proto_depIdxs = []int32{
	2, // 0: master.data.country.message.ListResponse.items:type_name -> master.data.country.message.ListItem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_masterData_country_v1_message_response_proto_init() }
func file_masterData_country_v1_message_response_proto_init() {
	if File_masterData_country_v1_message_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_masterData_country_v1_message_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportResponse); i {
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
		file_masterData_country_v1_message_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_masterData_country_v1_message_response_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListItem); i {
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
			RawDescriptor: file_masterData_country_v1_message_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_masterData_country_v1_message_response_proto_goTypes,
		DependencyIndexes: file_masterData_country_v1_message_response_proto_depIdxs,
		MessageInfos:      file_masterData_country_v1_message_response_proto_msgTypes,
	}.Build()
	File_masterData_country_v1_message_response_proto = out.File
	file_masterData_country_v1_message_response_proto_rawDesc = nil
	file_masterData_country_v1_message_response_proto_goTypes = nil
	file_masterData_country_v1_message_response_proto_depIdxs = nil
}
