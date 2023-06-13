// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: pkg/pb/club.proto

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

type CreateClubRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Year   int64  `protobuf:"varint,2,opt,name=year,proto3" json:"year,omitempty"`
	Trophy int64  `protobuf:"varint,3,opt,name=trophy,proto3" json:"trophy,omitempty"`
}

func (x *CreateClubRequest) Reset() {
	*x = CreateClubRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_club_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClubRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClubRequest) ProtoMessage() {}

func (x *CreateClubRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_club_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClubRequest.ProtoReflect.Descriptor instead.
func (*CreateClubRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_club_proto_rawDescGZIP(), []int{0}
}

func (x *CreateClubRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateClubRequest) GetYear() int64 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *CreateClubRequest) GetTrophy() int64 {
	if x != nil {
		return x.Trophy
	}
	return 0
}

type CreateClubResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Id     int64  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateClubResponse) Reset() {
	*x = CreateClubResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_club_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClubResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClubResponse) ProtoMessage() {}

func (x *CreateClubResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_club_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClubResponse.ProtoReflect.Descriptor instead.
func (*CreateClubResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_club_proto_rawDescGZIP(), []int{1}
}

func (x *CreateClubResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateClubResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CreateClubResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FindClubData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Year   int64  `protobuf:"varint,3,opt,name=year,proto3" json:"year,omitempty"`
	Trophy int64  `protobuf:"varint,4,opt,name=trophy,proto3" json:"trophy,omitempty"`
}

func (x *FindClubData) Reset() {
	*x = FindClubData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_club_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindClubData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindClubData) ProtoMessage() {}

func (x *FindClubData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_club_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindClubData.ProtoReflect.Descriptor instead.
func (*FindClubData) Descriptor() ([]byte, []int) {
	return file_pkg_pb_club_proto_rawDescGZIP(), []int{2}
}

func (x *FindClubData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FindClubData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FindClubData) GetYear() int64 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *FindClubData) GetTrophy() int64 {
	if x != nil {
		return x.Trophy
	}
	return 0
}

type FindClubRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindClubRequest) Reset() {
	*x = FindClubRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_club_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindClubRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindClubRequest) ProtoMessage() {}

func (x *FindClubRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_club_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindClubRequest.ProtoReflect.Descriptor instead.
func (*FindClubRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_club_proto_rawDescGZIP(), []int{3}
}

func (x *FindClubRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FindClubResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64         `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string        `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data   *FindClubData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *FindClubResponse) Reset() {
	*x = FindClubResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_club_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindClubResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindClubResponse) ProtoMessage() {}

func (x *FindClubResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_club_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindClubResponse.ProtoReflect.Descriptor instead.
func (*FindClubResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_club_proto_rawDescGZIP(), []int{4}
}

func (x *FindClubResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *FindClubResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *FindClubResponse) GetData() *FindClubData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_pkg_pb_club_proto protoreflect.FileDescriptor

var file_pkg_pb_club_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x53, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x72, 0x6f, 0x70, 0x68, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x72, 0x6f, 0x70, 0x68, 0x79, 0x22, 0x52,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x5e, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x6c, 0x75, 0x62, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x72,
	0x6f, 0x70, 0x68, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x72, 0x6f, 0x70,
	0x68, 0x79, 0x22, 0x21, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x6c, 0x75, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x68, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x6c, 0x75,
	0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x43, 0x6c, 0x75, 0x62, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32,
	0x8d, 0x01, 0x0a, 0x0b, 0x43, 0x6c, 0x75, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x41, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x62, 0x12, 0x17, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x62, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3b, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x6c, 0x75, 0x62, 0x12, 0x15,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x6c, 0x75, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x43, 0x6c, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_club_proto_rawDescOnce sync.Once
	file_pkg_pb_club_proto_rawDescData = file_pkg_pb_club_proto_rawDesc
)

func file_pkg_pb_club_proto_rawDescGZIP() []byte {
	file_pkg_pb_club_proto_rawDescOnce.Do(func() {
		file_pkg_pb_club_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_club_proto_rawDescData)
	})
	return file_pkg_pb_club_proto_rawDescData
}

var file_pkg_pb_club_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_pb_club_proto_goTypes = []interface{}{
	(*CreateClubRequest)(nil),  // 0: auth.CreateClubRequest
	(*CreateClubResponse)(nil), // 1: auth.CreateClubResponse
	(*FindClubData)(nil),       // 2: auth.FindClubData
	(*FindClubRequest)(nil),    // 3: auth.FindClubRequest
	(*FindClubResponse)(nil),   // 4: auth.FindClubResponse
}
var file_pkg_pb_club_proto_depIdxs = []int32{
	2, // 0: auth.FindClubResponse.data:type_name -> auth.FindClubData
	0, // 1: auth.ClubService.CreateClub:input_type -> auth.CreateClubRequest
	3, // 2: auth.ClubService.FindClub:input_type -> auth.FindClubRequest
	1, // 3: auth.ClubService.CreateClub:output_type -> auth.CreateClubResponse
	4, // 4: auth.ClubService.FindClub:output_type -> auth.FindClubResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_pb_club_proto_init() }
func file_pkg_pb_club_proto_init() {
	if File_pkg_pb_club_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_club_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClubRequest); i {
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
		file_pkg_pb_club_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClubResponse); i {
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
		file_pkg_pb_club_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindClubData); i {
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
		file_pkg_pb_club_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindClubRequest); i {
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
		file_pkg_pb_club_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindClubResponse); i {
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
			RawDescriptor: file_pkg_pb_club_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_club_proto_goTypes,
		DependencyIndexes: file_pkg_pb_club_proto_depIdxs,
		MessageInfos:      file_pkg_pb_club_proto_msgTypes,
	}.Build()
	File_pkg_pb_club_proto = out.File
	file_pkg_pb_club_proto_rawDesc = nil
	file_pkg_pb_club_proto_goTypes = nil
	file_pkg_pb_club_proto_depIdxs = nil
}

const _ = grpc.SupportPackageIsVersion7

const (
	ClubService_CreateClub_FullMethodName = "/auth.ClubService/CreateClub"
	ClubService_FindClub_FullMethodName   = "/auth.ClubService/FindClub"
)

// ClubServiceClient is the client API for ClubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClubServiceClient interface {
	CreateClub(ctx context.Context, in *CreateClubRequest, opts ...grpc.CallOption) (*CreateClubResponse, error)
	FindClub(ctx context.Context, in *FindClubRequest, opts ...grpc.CallOption) (*FindClubResponse, error)
}

type clubServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClubServiceClient(cc grpc.ClientConnInterface) ClubServiceClient {
	return &clubServiceClient{cc}
}

func (c *clubServiceClient) CreateClub(ctx context.Context, in *CreateClubRequest, opts ...grpc.CallOption) (*CreateClubResponse, error) {
	out := new(CreateClubResponse)
	err := c.cc.Invoke(ctx, ClubService_CreateClub_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubServiceClient) FindClub(ctx context.Context, in *FindClubRequest, opts ...grpc.CallOption) (*FindClubResponse, error) {
	out := new(FindClubResponse)
	err := c.cc.Invoke(ctx, ClubService_FindClub_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClubServiceServer is the server API for ClubService service.
// All implementations must embed UnimplementedClubServiceServer
// for forward compatibility
type ClubServiceServer interface {
	CreateClub(context.Context, *CreateClubRequest) (*CreateClubResponse, error)
	FindClub(context.Context, *FindClubRequest) (*FindClubResponse, error)
}

// UnimplementedClubServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClubServiceServer struct {
}

func (UnimplementedClubServiceServer) CreateClub(context.Context, *CreateClubRequest) (*CreateClubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClub not implemented")
}
func (UnimplementedClubServiceServer) FindClub(context.Context, *FindClubRequest) (*FindClubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindClub not implemented")
}
func (UnimplementedClubServiceServer) mustEmbedUnimplementedClubServiceServer() {}

// UnsafeClubServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClubServiceServer will
// result in compilation errors.
type UnsafeClubServiceServer interface {
	mustEmbedUnimplementedClubServiceServer()
}

func RegisterClubServiceServer(s *grpc.Server, srv ClubServiceServer) {
	s.RegisterService(&ClubService_ServiceDesc, srv)
}

func _ClubService_CreateClub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClubServiceServer).CreateClub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClubService_CreateClub_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClubServiceServer).CreateClub(ctx, req.(*CreateClubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClubService_FindClub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindClubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClubServiceServer).FindClub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClubService_FindClub_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClubServiceServer).FindClub(ctx, req.(*FindClubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClubService_ServiceDesc is the grpc.ServiceDesc for ClubService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClubService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.ClubService",
	HandlerType: (*ClubServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClub",
			Handler:    _ClubService_CreateClub_Handler,
		},
		{
			MethodName: "FindClub",
			Handler:    _ClubService_FindClub_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/club.proto",
}
