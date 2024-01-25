// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: octoback/core/v1/service.proto

package corev1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetGroceryListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the grocery list to get.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetGroceryListRequest) Reset() {
	*x = GetGroceryListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_octoback_core_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroceryListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroceryListRequest) ProtoMessage() {}

func (x *GetGroceryListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_octoback_core_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroceryListRequest.ProtoReflect.Descriptor instead.
func (*GetGroceryListRequest) Descriptor() ([]byte, []int) {
	return file_octoback_core_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetGroceryListRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetGroceryListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The grocery list.
	GroceryList *GroceryList `protobuf:"bytes,1,opt,name=grocery_list,json=groceryList,proto3" json:"grocery_list,omitempty"`
}

func (x *GetGroceryListResponse) Reset() {
	*x = GetGroceryListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_octoback_core_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroceryListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroceryListResponse) ProtoMessage() {}

func (x *GetGroceryListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_octoback_core_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroceryListResponse.ProtoReflect.Descriptor instead.
func (*GetGroceryListResponse) Descriptor() ([]byte, []int) {
	return file_octoback_core_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetGroceryListResponse) GetGroceryList() *GroceryList {
	if x != nil {
		return x.GroceryList
	}
	return nil
}

type CreateGroceryListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The grocery list name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateGroceryListRequest) Reset() {
	*x = CreateGroceryListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_octoback_core_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroceryListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroceryListRequest) ProtoMessage() {}

func (x *CreateGroceryListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_octoback_core_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroceryListRequest.ProtoReflect.Descriptor instead.
func (*CreateGroceryListRequest) Descriptor() ([]byte, []int) {
	return file_octoback_core_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateGroceryListRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateGroceryListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The grocery list.
	GroceryList *GroceryList `protobuf:"bytes,1,opt,name=grocery_list,json=groceryList,proto3" json:"grocery_list,omitempty"`
}

func (x *CreateGroceryListResponse) Reset() {
	*x = CreateGroceryListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_octoback_core_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroceryListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroceryListResponse) ProtoMessage() {}

func (x *CreateGroceryListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_octoback_core_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroceryListResponse.ProtoReflect.Descriptor instead.
func (*CreateGroceryListResponse) Descriptor() ([]byte, []int) {
	return file_octoback_core_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateGroceryListResponse) GetGroceryList() *GroceryList {
	if x != nil {
		return x.GroceryList
	}
	return nil
}

type UpdateGroceryListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The grocery list.
	GroceryList *GroceryList `protobuf:"bytes,1,opt,name=grocery_list,json=groceryList,proto3" json:"grocery_list,omitempty"`
	// The update mask.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateGroceryListRequest) Reset() {
	*x = UpdateGroceryListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_octoback_core_v1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroceryListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroceryListRequest) ProtoMessage() {}

func (x *UpdateGroceryListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_octoback_core_v1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroceryListRequest.ProtoReflect.Descriptor instead.
func (*UpdateGroceryListRequest) Descriptor() ([]byte, []int) {
	return file_octoback_core_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateGroceryListRequest) GetGroceryList() *GroceryList {
	if x != nil {
		return x.GroceryList
	}
	return nil
}

func (x *UpdateGroceryListRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type UpdateGroceryListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The grocery list.
	GroceryList *GroceryList `protobuf:"bytes,1,opt,name=grocery_list,json=groceryList,proto3" json:"grocery_list,omitempty"`
}

func (x *UpdateGroceryListResponse) Reset() {
	*x = UpdateGroceryListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_octoback_core_v1_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroceryListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroceryListResponse) ProtoMessage() {}

func (x *UpdateGroceryListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_octoback_core_v1_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroceryListResponse.ProtoReflect.Descriptor instead.
func (*UpdateGroceryListResponse) Descriptor() ([]byte, []int) {
	return file_octoback_core_v1_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateGroceryListResponse) GetGroceryList() *GroceryList {
	if x != nil {
		return x.GroceryList
	}
	return nil
}

type ListGroceryListsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListGroceryListsRequest) Reset() {
	*x = ListGroceryListsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_octoback_core_v1_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGroceryListsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGroceryListsRequest) ProtoMessage() {}

func (x *ListGroceryListsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_octoback_core_v1_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGroceryListsRequest.ProtoReflect.Descriptor instead.
func (*ListGroceryListsRequest) Descriptor() ([]byte, []int) {
	return file_octoback_core_v1_service_proto_rawDescGZIP(), []int{6}
}

type ListGroceryListsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroceryLists []*GroceryList `protobuf:"bytes,1,rep,name=grocery_lists,json=groceryLists,proto3" json:"grocery_lists,omitempty"`
}

func (x *ListGroceryListsResponse) Reset() {
	*x = ListGroceryListsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_octoback_core_v1_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGroceryListsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGroceryListsResponse) ProtoMessage() {}

func (x *ListGroceryListsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_octoback_core_v1_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGroceryListsResponse.ProtoReflect.Descriptor instead.
func (*ListGroceryListsResponse) Descriptor() ([]byte, []int) {
	return file_octoback_core_v1_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListGroceryListsResponse) GetGroceryLists() []*GroceryList {
	if x != nil {
		return x.GroceryLists
	}
	return nil
}

type DeleteGroceryListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the grocery list to delete.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteGroceryListRequest) Reset() {
	*x = DeleteGroceryListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_octoback_core_v1_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGroceryListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroceryListRequest) ProtoMessage() {}

func (x *DeleteGroceryListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_octoback_core_v1_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroceryListRequest.ProtoReflect.Descriptor instead.
func (*DeleteGroceryListRequest) Descriptor() ([]byte, []int) {
	return file_octoback_core_v1_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteGroceryListRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteGroceryListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteGroceryListResponse) Reset() {
	*x = DeleteGroceryListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_octoback_core_v1_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGroceryListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroceryListResponse) ProtoMessage() {}

func (x *DeleteGroceryListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_octoback_core_v1_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroceryListResponse.ProtoReflect.Descriptor instead.
func (*DeleteGroceryListResponse) Descriptor() ([]byte, []int) {
	return file_octoback_core_v1_service_proto_rawDescGZIP(), []int{9}
}

var File_octoback_core_v1_service_proto protoreflect.FileDescriptor

var file_octoback_core_v1_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6f, 0x63, 0x74, 0x6f, 0x62, 0x61, 0x63, 0x6b, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x6f, 0x63, 0x74, 0x6f, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x6f, 0x63, 0x74, 0x6f, 0x62, 0x61, 0x63, 0x6b, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72,
	0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5a, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0c, 0x67, 0x72, 0x6f, 0x63, 0x65, 0x72,
	0x79, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f,
	0x63, 0x74, 0x6f, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0b, 0x67, 0x72, 0x6f,
	0x63, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x37, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x5d, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x63, 0x65,
	0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40,
	0x0a, 0x0c, 0x67, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x63, 0x74, 0x6f, 0x62, 0x61, 0x63, 0x6b, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x0b, 0x67, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x99, 0x01, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x63, 0x65,
	0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a,
	0x0c, 0x67, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x63, 0x74, 0x6f, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x0b, 0x67, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x5d, 0x0a, 0x19,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0c, 0x67, 0x72, 0x6f,
	0x63, 0x65, 0x72, 0x79, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x6f, 0x63, 0x74, 0x6f, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0b,
	0x67, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x19, 0x0a, 0x17, 0x4c,
	0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5e, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72,
	0x6f, 0x63, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x42, 0x0a, 0x0d, 0x67, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x63, 0x74, 0x6f,
	0x62, 0x61, 0x63, 0x6b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f,
	0x63, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0c, 0x67, 0x72, 0x6f, 0x63, 0x65, 0x72,
	0x79, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x22, 0x2a, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x1b, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x63,
	0x65, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xa3, 0x04, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x2e,
	0x6f, 0x63, 0x74, 0x6f, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6f, 0x63, 0x74, 0x6f, 0x62, 0x61, 0x63,
	0x6b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f,
	0x63, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x6c, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72,
	0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x2e, 0x6f, 0x63, 0x74, 0x6f, 0x62, 0x61, 0x63, 0x6b,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47,
	0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2b, 0x2e, 0x6f, 0x63, 0x74, 0x6f, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x63, 0x65,
	0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6c,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x2a, 0x2e, 0x6f, 0x63, 0x74, 0x6f, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f,
	0x63, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x6f, 0x63, 0x74, 0x6f, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x69, 0x0a, 0x10,
	0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x73,
	0x12, 0x29, 0x2e, 0x6f, 0x63, 0x74, 0x6f, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x4c,
	0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6f, 0x63,
	0x74, 0x6f, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6c, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x2e, 0x6f,
	0x63, 0x74, 0x6f, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6f, 0x63, 0x74, 0x6f, 0x62,
	0x61, 0x63, 0x6b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xcf, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x63,
	0x74, 0x6f, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x47,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x69, 0x6c, 0x69, 0x61,
	0x6e, 0x64, 0x62, 0x69, 0x67, 0x62, 0x6c, 0x75, 0x65, 0x2f, 0x6f, 0x63, 0x74, 0x6f, 0x62, 0x61,
	0x63, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f,
	0x6f, 0x63, 0x74, 0x6f, 0x62, 0x61, 0x63, 0x6b, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31,
	0x3b, 0x63, 0x6f, 0x72, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4f, 0x43, 0x58, 0xaa, 0x02, 0x10,
	0x4f, 0x63, 0x74, 0x6f, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x10, 0x4f, 0x63, 0x74, 0x6f, 0x62, 0x61, 0x63, 0x6b, 0x5c, 0x43, 0x6f, 0x72, 0x65,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x4f, 0x63, 0x74, 0x6f, 0x62, 0x61, 0x63, 0x6b, 0x5c, 0x43,
	0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x12, 0x4f, 0x63, 0x74, 0x6f, 0x62, 0x61, 0x63, 0x6b, 0x3a, 0x3a, 0x43,
	0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_octoback_core_v1_service_proto_rawDescOnce sync.Once
	file_octoback_core_v1_service_proto_rawDescData = file_octoback_core_v1_service_proto_rawDesc
)

func file_octoback_core_v1_service_proto_rawDescGZIP() []byte {
	file_octoback_core_v1_service_proto_rawDescOnce.Do(func() {
		file_octoback_core_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_octoback_core_v1_service_proto_rawDescData)
	})
	return file_octoback_core_v1_service_proto_rawDescData
}

var file_octoback_core_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_octoback_core_v1_service_proto_goTypes = []interface{}{
	(*GetGroceryListRequest)(nil),     // 0: octoback.core.v1.GetGroceryListRequest
	(*GetGroceryListResponse)(nil),    // 1: octoback.core.v1.GetGroceryListResponse
	(*CreateGroceryListRequest)(nil),  // 2: octoback.core.v1.CreateGroceryListRequest
	(*CreateGroceryListResponse)(nil), // 3: octoback.core.v1.CreateGroceryListResponse
	(*UpdateGroceryListRequest)(nil),  // 4: octoback.core.v1.UpdateGroceryListRequest
	(*UpdateGroceryListResponse)(nil), // 5: octoback.core.v1.UpdateGroceryListResponse
	(*ListGroceryListsRequest)(nil),   // 6: octoback.core.v1.ListGroceryListsRequest
	(*ListGroceryListsResponse)(nil),  // 7: octoback.core.v1.ListGroceryListsResponse
	(*DeleteGroceryListRequest)(nil),  // 8: octoback.core.v1.DeleteGroceryListRequest
	(*DeleteGroceryListResponse)(nil), // 9: octoback.core.v1.DeleteGroceryListResponse
	(*GroceryList)(nil),               // 10: octoback.core.v1.GroceryList
	(*fieldmaskpb.FieldMask)(nil),     // 11: google.protobuf.FieldMask
}
var file_octoback_core_v1_service_proto_depIdxs = []int32{
	10, // 0: octoback.core.v1.GetGroceryListResponse.grocery_list:type_name -> octoback.core.v1.GroceryList
	10, // 1: octoback.core.v1.CreateGroceryListResponse.grocery_list:type_name -> octoback.core.v1.GroceryList
	10, // 2: octoback.core.v1.UpdateGroceryListRequest.grocery_list:type_name -> octoback.core.v1.GroceryList
	11, // 3: octoback.core.v1.UpdateGroceryListRequest.update_mask:type_name -> google.protobuf.FieldMask
	10, // 4: octoback.core.v1.UpdateGroceryListResponse.grocery_list:type_name -> octoback.core.v1.GroceryList
	10, // 5: octoback.core.v1.ListGroceryListsResponse.grocery_lists:type_name -> octoback.core.v1.GroceryList
	0,  // 6: octoback.core.v1.Service.GetGroceryList:input_type -> octoback.core.v1.GetGroceryListRequest
	2,  // 7: octoback.core.v1.Service.CreateGroceryList:input_type -> octoback.core.v1.CreateGroceryListRequest
	4,  // 8: octoback.core.v1.Service.UpdateGroceryList:input_type -> octoback.core.v1.UpdateGroceryListRequest
	6,  // 9: octoback.core.v1.Service.ListGroceryLists:input_type -> octoback.core.v1.ListGroceryListsRequest
	8,  // 10: octoback.core.v1.Service.DeleteGroceryList:input_type -> octoback.core.v1.DeleteGroceryListRequest
	1,  // 11: octoback.core.v1.Service.GetGroceryList:output_type -> octoback.core.v1.GetGroceryListResponse
	3,  // 12: octoback.core.v1.Service.CreateGroceryList:output_type -> octoback.core.v1.CreateGroceryListResponse
	5,  // 13: octoback.core.v1.Service.UpdateGroceryList:output_type -> octoback.core.v1.UpdateGroceryListResponse
	7,  // 14: octoback.core.v1.Service.ListGroceryLists:output_type -> octoback.core.v1.ListGroceryListsResponse
	9,  // 15: octoback.core.v1.Service.DeleteGroceryList:output_type -> octoback.core.v1.DeleteGroceryListResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_octoback_core_v1_service_proto_init() }
func file_octoback_core_v1_service_proto_init() {
	if File_octoback_core_v1_service_proto != nil {
		return
	}
	file_octoback_core_v1_groceries_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_octoback_core_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroceryListRequest); i {
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
		file_octoback_core_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroceryListResponse); i {
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
		file_octoback_core_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroceryListRequest); i {
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
		file_octoback_core_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroceryListResponse); i {
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
		file_octoback_core_v1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroceryListRequest); i {
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
		file_octoback_core_v1_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroceryListResponse); i {
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
		file_octoback_core_v1_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGroceryListsRequest); i {
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
		file_octoback_core_v1_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGroceryListsResponse); i {
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
		file_octoback_core_v1_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGroceryListRequest); i {
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
		file_octoback_core_v1_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGroceryListResponse); i {
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
			RawDescriptor: file_octoback_core_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_octoback_core_v1_service_proto_goTypes,
		DependencyIndexes: file_octoback_core_v1_service_proto_depIdxs,
		MessageInfos:      file_octoback_core_v1_service_proto_msgTypes,
	}.Build()
	File_octoback_core_v1_service_proto = out.File
	file_octoback_core_v1_service_proto_rawDesc = nil
	file_octoback_core_v1_service_proto_goTypes = nil
	file_octoback_core_v1_service_proto_depIdxs = nil
}
