// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: pkg/footballer/pb/footablaller.proto

package pb

import (
	context "context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

type CreateFootballerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Surname string `protobuf:"bytes,2,opt,name=surname,proto3" json:"surname,omitempty"`
	Price   int64  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	Club    string `protobuf:"bytes,4,opt,name=club,proto3" json:"club,omitempty"`
}

func (x *CreateFootballerRequest) Reset() {
	*x = CreateFootballerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_footballer_pb_footablaller_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFootballerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFootballerRequest) ProtoMessage() {}

func (x *CreateFootballerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_footballer_pb_footablaller_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFootballerRequest.ProtoReflect.Descriptor instead.
func (*CreateFootballerRequest) Descriptor() ([]byte, []int) {
	return file_pkg_footballer_pb_footablaller_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFootballerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateFootballerRequest) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *CreateFootballerRequest) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateFootballerRequest) GetClub() string {
	if x != nil {
		return x.Club
	}
	return ""
}

type CreateFootballerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Id     int64  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateFootballerResponse) Reset() {
	*x = CreateFootballerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_footballer_pb_footablaller_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFootballerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFootballerResponse) ProtoMessage() {}

func (x *CreateFootballerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_footballer_pb_footablaller_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFootballerResponse.ProtoReflect.Descriptor instead.
func (*CreateFootballerResponse) Descriptor() ([]byte, []int) {
	return file_pkg_footballer_pb_footablaller_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFootballerResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateFootballerResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CreateFootballerResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FindOneData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Surname string `protobuf:"bytes,3,opt,name=surname,proto3" json:"surname,omitempty"`
	Price   int64  `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	Club    string `protobuf:"bytes,5,opt,name=club,proto3" json:"club,omitempty"`
}

func (x *FindOneData) Reset() {
	*x = FindOneData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_footballer_pb_footablaller_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneData) ProtoMessage() {}

func (x *FindOneData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_footballer_pb_footablaller_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneData.ProtoReflect.Descriptor instead.
func (*FindOneData) Descriptor() ([]byte, []int) {
	return file_pkg_footballer_pb_footablaller_proto_rawDescGZIP(), []int{2}
}

func (x *FindOneData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FindOneData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FindOneData) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *FindOneData) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *FindOneData) GetClub() string {
	if x != nil {
		return x.Club
	}
	return ""
}

type FindOneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindOneRequest) Reset() {
	*x = FindOneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_footballer_pb_footablaller_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneRequest) ProtoMessage() {}

func (x *FindOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_footballer_pb_footablaller_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneRequest.ProtoReflect.Descriptor instead.
func (*FindOneRequest) Descriptor() ([]byte, []int) {
	return file_pkg_footballer_pb_footablaller_proto_rawDescGZIP(), []int{3}
}

func (x *FindOneRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FindOneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64        `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string       `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data   *FindOneData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *FindOneResponse) Reset() {
	*x = FindOneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_footballer_pb_footablaller_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneResponse) ProtoMessage() {}

func (x *FindOneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_footballer_pb_footablaller_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneResponse.ProtoReflect.Descriptor instead.
func (*FindOneResponse) Descriptor() ([]byte, []int) {
	return file_pkg_footballer_pb_footablaller_proto_rawDescGZIP(), []int{4}
}

func (x *FindOneResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *FindOneResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *FindOneResponse) GetData() *FindOneData {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeleteFootballerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteFootballerRequest) Reset() {
	*x = DeleteFootballerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_footballer_pb_footablaller_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFootballerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFootballerRequest) ProtoMessage() {}

func (x *DeleteFootballerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_footballer_pb_footablaller_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFootballerRequest.ProtoReflect.Descriptor instead.
func (*DeleteFootballerRequest) Descriptor() ([]byte, []int) {
	return file_pkg_footballer_pb_footablaller_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteFootballerRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteFootballerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeleteFootballerResponse) Reset() {
	*x = DeleteFootballerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_footballer_pb_footablaller_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFootballerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFootballerResponse) ProtoMessage() {}

func (x *DeleteFootballerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_footballer_pb_footablaller_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFootballerResponse.ProtoReflect.Descriptor instead.
func (*DeleteFootballerResponse) Descriptor() ([]byte, []int) {
	return file_pkg_footballer_pb_footablaller_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteFootballerResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DeleteFootballerResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_pkg_footballer_pb_footablaller_proto protoreflect.FileDescriptor

var file_pkg_footballer_pb_footablaller_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x6f, 0x6f, 0x74, 0x62, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x2f, 0x70, 0x62, 0x2f, 0x66, 0x6f, 0x6f, 0x74, 0x61, 0x62, 0x6c, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x71, 0x0a, 0x17,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x74, 0x62, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6c, 0x75, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6c, 0x75, 0x62, 0x22,
	0x58, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x74, 0x62, 0x61, 0x6c,
	0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x75, 0x0a, 0x0b, 0x46, 0x69, 0x6e,
	0x64, 0x4f, 0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6c, 0x75, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6c, 0x75, 0x62,
	0x22, 0x20, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x66, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x29, 0x0a, 0x17, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x74, 0x62, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x48, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46,
	0x6f, 0x6f, 0x74, 0x62, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32,
	0xf7, 0x01, 0x0a, 0x11, 0x46, 0x6f, 0x6f, 0x74, 0x62, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46,
	0x6f, 0x6f, 0x74, 0x62, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x74, 0x62, 0x61, 0x6c, 0x6c, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x74, 0x62, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x07, 0x46, 0x69,
	0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x12, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f,
	0x6f, 0x74, 0x62, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x74, 0x62, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x74, 0x62, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x66, 0x6f, 0x6f, 0x74, 0x62, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_footballer_pb_footablaller_proto_rawDescOnce sync.Once
	file_pkg_footballer_pb_footablaller_proto_rawDescData = file_pkg_footballer_pb_footablaller_proto_rawDesc
)

func file_pkg_footballer_pb_footablaller_proto_rawDescGZIP() []byte {
	file_pkg_footballer_pb_footablaller_proto_rawDescOnce.Do(func() {
		file_pkg_footballer_pb_footablaller_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_footballer_pb_footablaller_proto_rawDescData)
	})
	return file_pkg_footballer_pb_footablaller_proto_rawDescData
}

var file_pkg_footballer_pb_footablaller_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_footballer_pb_footablaller_proto_goTypes = []interface{}{
	(*CreateFootballerRequest)(nil),  // 0: auth.CreateFootballerRequest
	(*CreateFootballerResponse)(nil), // 1: auth.CreateFootballerResponse
	(*FindOneData)(nil),              // 2: auth.FindOneData
	(*FindOneRequest)(nil),           // 3: auth.FindOneRequest
	(*FindOneResponse)(nil),          // 4: auth.FindOneResponse
	(*DeleteFootballerRequest)(nil),  // 5: auth.DeleteFootballerRequest
	(*DeleteFootballerResponse)(nil), // 6: auth.DeleteFootballerResponse
}
var file_pkg_footballer_pb_footablaller_proto_depIdxs = []int32{
	2, // 0: auth.FindOneResponse.data:type_name -> auth.FindOneData
	0, // 1: auth.FootballerService.CreateFootballer:input_type -> auth.CreateFootballerRequest
	3, // 2: auth.FootballerService.FindOne:input_type -> auth.FindOneRequest
	5, // 3: auth.FootballerService.DeleteFootballer:input_type -> auth.DeleteFootballerRequest
	1, // 4: auth.FootballerService.CreateFootballer:output_type -> auth.CreateFootballerResponse
	4, // 5: auth.FootballerService.FindOne:output_type -> auth.FindOneResponse
	6, // 6: auth.FootballerService.DeleteFootballer:output_type -> auth.DeleteFootballerResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_footballer_pb_footablaller_proto_init() }
func file_pkg_footballer_pb_footablaller_proto_init() {
	if File_pkg_footballer_pb_footablaller_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_footballer_pb_footablaller_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFootballerRequest); i {
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
		file_pkg_footballer_pb_footablaller_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFootballerResponse); i {
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
		file_pkg_footballer_pb_footablaller_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneData); i {
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
		file_pkg_footballer_pb_footablaller_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneRequest); i {
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
		file_pkg_footballer_pb_footablaller_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneResponse); i {
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
		file_pkg_footballer_pb_footablaller_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFootballerRequest); i {
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
		file_pkg_footballer_pb_footablaller_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFootballerResponse); i {
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
			RawDescriptor: file_pkg_footballer_pb_footablaller_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_footballer_pb_footablaller_proto_goTypes,
		DependencyIndexes: file_pkg_footballer_pb_footablaller_proto_depIdxs,
		MessageInfos:      file_pkg_footballer_pb_footablaller_proto_msgTypes,
	}.Build()
	File_pkg_footballer_pb_footablaller_proto = out.File
	file_pkg_footballer_pb_footablaller_proto_rawDesc = nil
	file_pkg_footballer_pb_footablaller_proto_goTypes = nil
	file_pkg_footballer_pb_footablaller_proto_depIdxs = nil
}

const _ = grpc.SupportPackageIsVersion7

const (
	FootballerService_CreateFootballer_FullMethodName = "/auth.FootballerService/CreateFootballer"
	FootballerService_FindOne_FullMethodName          = "/auth.FootballerService/FindOne"
	FootballerService_DeleteFootballer_FullMethodName = "/auth.FootballerService/DeleteFootballer"
)

// FootballerServiceClient is the client API for FootballerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FootballerServiceClient interface {
	CreateFootballer(ctx context.Context, in *CreateFootballerRequest, opts ...grpc.CallOption) (*CreateFootballerResponse, error)
	FindOne(ctx context.Context, in *FindOneRequest, opts ...grpc.CallOption) (*FindOneResponse, error)
	DeleteFootballer(ctx context.Context, in *DeleteFootballerRequest, opts ...grpc.CallOption) (*DeleteFootballerResponse, error)
}

type footballerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFootballerServiceClient(cc grpc.ClientConnInterface) FootballerServiceClient {
	return &footballerServiceClient{cc}
}

func (c *footballerServiceClient) CreateFootballer(ctx context.Context, in *CreateFootballerRequest, opts ...grpc.CallOption) (*CreateFootballerResponse, error) {
	out := new(CreateFootballerResponse)
	err := c.cc.Invoke(ctx, FootballerService_CreateFootballer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *footballerServiceClient) FindOne(ctx context.Context, in *FindOneRequest, opts ...grpc.CallOption) (*FindOneResponse, error) {
	out := new(FindOneResponse)
	err := c.cc.Invoke(ctx, FootballerService_FindOne_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *footballerServiceClient) DeleteFootballer(ctx context.Context, in *DeleteFootballerRequest, opts ...grpc.CallOption) (*DeleteFootballerResponse, error) {
	out := new(DeleteFootballerResponse)
	err := c.cc.Invoke(ctx, FootballerService_DeleteFootballer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FootballerServiceServer is the server API for FootballerService service.
// All implementations must embed UnimplementedFootballerServiceServer
// for forward compatibility
type FootballerServiceServer interface {
	CreateFootballer(context.Context, *CreateFootballerRequest) (*CreateFootballerResponse, error)
	FindOne(context.Context, *FindOneRequest) (*FindOneResponse, error)
	DeleteFootballer(context.Context, *DeleteFootballerRequest) (*DeleteFootballerResponse, error)
	mustEmbedUnimplementedFootballerServiceServer()
}

// UnimplementedFootballerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFootballerServiceServer struct {
}

func (UnimplementedFootballerServiceServer) CreateFootballer(context.Context, *CreateFootballerRequest) (*CreateFootballerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFootballer not implemented")
}
func (UnimplementedFootballerServiceServer) FindOne(context.Context, *FindOneRequest) (*FindOneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOne not implemented")
}
func (UnimplementedFootballerServiceServer) DeleteFootballer(context.Context, *DeleteFootballerRequest) (*DeleteFootballerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFootballer not implemented")
}
func (UnimplementedFootballerServiceServer) mustEmbedUnimplementedFootballerServiceServer() {}

// UnsafeFootballerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FootballerServiceServer will
// result in compilation errors.
type UnsafeFootballerServiceServer interface {
	mustEmbedUnimplementedFootballerServiceServer()
}

func RegisterFootballerServiceServer(s grpc.ServiceRegistrar, srv FootballerServiceServer) {
	s.RegisterService(&FootballerService_ServiceDesc, srv)
}

func _FootballerService_CreateFootballer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFootballerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FootballerServiceServer).CreateFootballer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FootballerService_CreateFootballer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FootballerServiceServer).CreateFootballer(ctx, req.(*CreateFootballerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FootballerService_FindOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FootballerServiceServer).FindOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FootballerService_FindOne_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FootballerServiceServer).FindOne(ctx, req.(*FindOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FootballerService_DeleteFootballer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFootballerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FootballerServiceServer).DeleteFootballer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FootballerService_DeleteFootballer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FootballerServiceServer).DeleteFootballer(ctx, req.(*DeleteFootballerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FootballerService_ServiceDesc is the grpc.ServiceDesc for FootballerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FootballerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.FootballerService",
	HandlerType: (*FootballerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFootballer",
			Handler:    _FootballerService_CreateFootballer_Handler,
		},
		{
			MethodName: "FindOne",
			Handler:    _FootballerService_FindOne_Handler,
		},
		{
			MethodName: "DeleteFootballer",
			Handler:    _FootballerService_DeleteFootballer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/footballer/pb/footablaller.proto",
}
