// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: octoback/core/v1/groceries.proto

package corev1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

// GroceryList represents a list of grocery items.
type GroceryList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Items []*GroceryItem `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GroceryList) Reset() {
	*x = GroceryList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_octoback_core_v1_groceries_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroceryList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroceryList) ProtoMessage() {}

func (x *GroceryList) ProtoReflect() protoreflect.Message {
	mi := &file_octoback_core_v1_groceries_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroceryList.ProtoReflect.Descriptor instead.
func (*GroceryList) Descriptor() ([]byte, []int) {
	return file_octoback_core_v1_groceries_proto_rawDescGZIP(), []int{0}
}

func (x *GroceryList) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GroceryList) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GroceryList) GetItems() []*GroceryItem {
	if x != nil {
		return x.Items
	}
	return nil
}

// GroceryItem represents a single grocery item.
type GroceryItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Quantity int32  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Checked  bool   `protobuf:"varint,4,opt,name=checked,proto3" json:"checked,omitempty"`
}

func (x *GroceryItem) Reset() {
	*x = GroceryItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_octoback_core_v1_groceries_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroceryItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroceryItem) ProtoMessage() {}

func (x *GroceryItem) ProtoReflect() protoreflect.Message {
	mi := &file_octoback_core_v1_groceries_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroceryItem.ProtoReflect.Descriptor instead.
func (*GroceryItem) Descriptor() ([]byte, []int) {
	return file_octoback_core_v1_groceries_proto_rawDescGZIP(), []int{1}
}

func (x *GroceryItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GroceryItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GroceryItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *GroceryItem) GetChecked() bool {
	if x != nil {
		return x.Checked
	}
	return false
}

var File_octoback_core_v1_groceries_proto protoreflect.FileDescriptor

var file_octoback_core_v1_groceries_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6f, 0x63, 0x74, 0x6f, 0x62, 0x61, 0x63, 0x6b, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x6f, 0x63, 0x74, 0x6f, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x82, 0x01, 0x0a, 0x0b, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x2a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1a, 0xba,
	0x48, 0x17, 0x72, 0x15, 0x32, 0x13, 0x5e, 0x67, 0x6c, 0x5f, 0x5b, 0x5b, 0x3a, 0x61, 0x6c, 0x6e,
	0x75, 0x6d, 0x3a, 0x5d, 0x5d, 0x7b, 0x38, 0x7d, 0x24, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x33, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x6f, 0x63, 0x74, 0x6f, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x8e, 0x01, 0x0a, 0x0b, 0x47, 0x72, 0x6f, 0x63, 0x65,
	0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x2a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x1a, 0xba, 0x48, 0x17, 0x72, 0x15, 0x32, 0x13, 0x5e, 0x67, 0x69, 0x5f, 0x5b,
	0x5b, 0x3a, 0x61, 0x6c, 0x6e, 0x75, 0x6d, 0x3a, 0x5d, 0x5d, 0x7b, 0x38, 0x7d, 0x24, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x09, 0xba, 0x48, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x64, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x42, 0xd1, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e,
	0x6f, 0x63, 0x74, 0x6f, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x42, 0x0e, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x69, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b,
	0x69, 0x6c, 0x69, 0x61, 0x6e, 0x64, 0x62, 0x69, 0x67, 0x62, 0x6c, 0x75, 0x65, 0x2f, 0x6f, 0x63,
	0x74, 0x6f, 0x62, 0x61, 0x63, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x2f, 0x6f, 0x63, 0x74, 0x6f, 0x62, 0x61, 0x63, 0x6b, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4f, 0x43,
	0x58, 0xaa, 0x02, 0x10, 0x4f, 0x63, 0x74, 0x6f, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x43, 0x6f, 0x72,
	0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x4f, 0x63, 0x74, 0x6f, 0x62, 0x61, 0x63, 0x6b, 0x5c,
	0x43, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x4f, 0x63, 0x74, 0x6f, 0x62, 0x61,
	0x63, 0x6b, 0x5c, 0x43, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x4f, 0x63, 0x74, 0x6f, 0x62, 0x61, 0x63,
	0x6b, 0x3a, 0x3a, 0x43, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_octoback_core_v1_groceries_proto_rawDescOnce sync.Once
	file_octoback_core_v1_groceries_proto_rawDescData = file_octoback_core_v1_groceries_proto_rawDesc
)

func file_octoback_core_v1_groceries_proto_rawDescGZIP() []byte {
	file_octoback_core_v1_groceries_proto_rawDescOnce.Do(func() {
		file_octoback_core_v1_groceries_proto_rawDescData = protoimpl.X.CompressGZIP(file_octoback_core_v1_groceries_proto_rawDescData)
	})
	return file_octoback_core_v1_groceries_proto_rawDescData
}

var file_octoback_core_v1_groceries_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_octoback_core_v1_groceries_proto_goTypes = []interface{}{
	(*GroceryList)(nil), // 0: octoback.core.v1.GroceryList
	(*GroceryItem)(nil), // 1: octoback.core.v1.GroceryItem
}
var file_octoback_core_v1_groceries_proto_depIdxs = []int32{
	1, // 0: octoback.core.v1.GroceryList.items:type_name -> octoback.core.v1.GroceryItem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_octoback_core_v1_groceries_proto_init() }
func file_octoback_core_v1_groceries_proto_init() {
	if File_octoback_core_v1_groceries_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_octoback_core_v1_groceries_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroceryList); i {
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
		file_octoback_core_v1_groceries_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroceryItem); i {
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
			RawDescriptor: file_octoback_core_v1_groceries_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_octoback_core_v1_groceries_proto_goTypes,
		DependencyIndexes: file_octoback_core_v1_groceries_proto_depIdxs,
		MessageInfos:      file_octoback_core_v1_groceries_proto_msgTypes,
	}.Build()
	File_octoback_core_v1_groceries_proto = out.File
	file_octoback_core_v1_groceries_proto_rawDesc = nil
	file_octoback_core_v1_groceries_proto_goTypes = nil
	file_octoback_core_v1_groceries_proto_depIdxs = nil
}