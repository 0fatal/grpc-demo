// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: pp_spider.proto

package pp

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SectionType int32

const (
	SectionType__             SectionType = 0
	SectionType_ST_TEACH      SectionType = 36
	SectionType_ST_NEWS       SectionType = 239
	SectionType_ST_EQUIPMENT  SectionType = 227
	SectionType_ST_GAME_VIDEO SectionType = 2
	SectionType_ST_ASK        SectionType = 329
)

// Enum value maps for SectionType.
var (
	SectionType_name = map[int32]string{
		0:   "_",
		36:  "ST_TEACH",
		239: "ST_NEWS",
		227: "ST_EQUIPMENT",
		2:   "ST_GAME_VIDEO",
		329: "ST_ASK",
	}
	SectionType_value = map[string]int32{
		"_":             0,
		"ST_TEACH":      36,
		"ST_NEWS":       239,
		"ST_EQUIPMENT":  227,
		"ST_GAME_VIDEO": 2,
		"ST_ASK":        329,
	}
)

func (x SectionType) Enum() *SectionType {
	p := new(SectionType)
	*p = x
	return p
}

func (x SectionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SectionType) Descriptor() protoreflect.EnumDescriptor {
	return file_pp_spider_proto_enumTypes[0].Descriptor()
}

func (SectionType) Type() protoreflect.EnumType {
	return &file_pp_spider_proto_enumTypes[0]
}

func (x SectionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SectionType.Descriptor instead.
func (SectionType) EnumDescriptor() ([]byte, []int) {
	return file_pp_spider_proto_rawDescGZIP(), []int{0}
}

type NewsType int32

const (
	NewsType_NT_ALL       NewsType = 0
	NewsType_NT_GAME      NewsType = 229
	NewsType_NT_PLAYER    NewsType = 231
	NewsType_NT_EQUIPMENT NewsType = 232
	NewsType_NT_FUN       NewsType = 230
)

// Enum value maps for NewsType.
var (
	NewsType_name = map[int32]string{
		0:   "NT_ALL",
		229: "NT_GAME",
		231: "NT_PLAYER",
		232: "NT_EQUIPMENT",
		230: "NT_FUN",
	}
	NewsType_value = map[string]int32{
		"NT_ALL":       0,
		"NT_GAME":      229,
		"NT_PLAYER":    231,
		"NT_EQUIPMENT": 232,
		"NT_FUN":       230,
	}
)

func (x NewsType) Enum() *NewsType {
	p := new(NewsType)
	*p = x
	return p
}

func (x NewsType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NewsType) Descriptor() protoreflect.EnumDescriptor {
	return file_pp_spider_proto_enumTypes[1].Descriptor()
}

func (NewsType) Type() protoreflect.EnumType {
	return &file_pp_spider_proto_enumTypes[1]
}

func (x NewsType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NewsType.Descriptor instead.
func (NewsType) EnumDescriptor() ([]byte, []int) {
	return file_pp_spider_proto_rawDescGZIP(), []int{1}
}

type PPSpiderResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*PPSpiderBody `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *PPSpiderResp) Reset() {
	*x = PPSpiderResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pp_spider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PPSpiderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PPSpiderResp) ProtoMessage() {}

func (x *PPSpiderResp) ProtoReflect() protoreflect.Message {
	mi := &file_pp_spider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PPSpiderResp.ProtoReflect.Descriptor instead.
func (*PPSpiderResp) Descriptor() ([]byte, []int) {
	return file_pp_spider_proto_rawDescGZIP(), []int{0}
}

func (x *PPSpiderResp) GetData() []*PPSpiderBody {
	if x != nil {
		return x.Data
	}
	return nil
}

type PPSpiderBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Link  string `protobuf:"bytes,2,opt,name=link,proto3" json:"link,omitempty"`
	Time  string `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	Date  string `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *PPSpiderBody) Reset() {
	*x = PPSpiderBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pp_spider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PPSpiderBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PPSpiderBody) ProtoMessage() {}

func (x *PPSpiderBody) ProtoReflect() protoreflect.Message {
	mi := &file_pp_spider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PPSpiderBody.ProtoReflect.Descriptor instead.
func (*PPSpiderBody) Descriptor() ([]byte, []int) {
	return file_pp_spider_proto_rawDescGZIP(), []int{1}
}

func (x *PPSpiderBody) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PPSpiderBody) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *PPSpiderBody) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *PPSpiderBody) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type PPSpiderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SectionType SectionType `protobuf:"varint,1,opt,name=section_type,json=sectionType,proto3,enum=pp.SectionType" json:"section_type,omitempty"`
	Page        uint32      `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	NewsType    NewsType    `protobuf:"varint,3,opt,name=news_type,json=newsType,proto3,enum=pp.NewsType" json:"news_type,omitempty"`
}

func (x *PPSpiderReq) Reset() {
	*x = PPSpiderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pp_spider_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PPSpiderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PPSpiderReq) ProtoMessage() {}

func (x *PPSpiderReq) ProtoReflect() protoreflect.Message {
	mi := &file_pp_spider_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PPSpiderReq.ProtoReflect.Descriptor instead.
func (*PPSpiderReq) Descriptor() ([]byte, []int) {
	return file_pp_spider_proto_rawDescGZIP(), []int{2}
}

func (x *PPSpiderReq) GetSectionType() SectionType {
	if x != nil {
		return x.SectionType
	}
	return SectionType__
}

func (x *PPSpiderReq) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PPSpiderReq) GetNewsType() NewsType {
	if x != nil {
		return x.NewsType
	}
	return NewsType_NT_ALL
}

var File_pp_spider_proto protoreflect.FileDescriptor

var file_pp_spider_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x70, 0x5f, 0x73, 0x70, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x70, 0x22, 0x34, 0x0a, 0x0c, 0x50, 0x50, 0x53, 0x70, 0x69, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x70, 0x2e, 0x50, 0x50, 0x53, 0x70, 0x69, 0x64, 0x65,
	0x72, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x60, 0x0a, 0x0c, 0x50,
	0x50, 0x53, 0x70, 0x69, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x80, 0x01,
	0x0a, 0x0b, 0x50, 0x50, 0x53, 0x70, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x32, 0x0a,
	0x0c, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x70, 0x70, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x73, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x70, 0x70, 0x2e, 0x4e, 0x65,
	0x77, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x73, 0x54, 0x79, 0x70, 0x65,
	0x2a, 0x63, 0x0a, 0x0b, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x05, 0x0a, 0x01, 0x5f, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x5f, 0x54, 0x45, 0x41,
	0x43, 0x48, 0x10, 0x24, 0x12, 0x0c, 0x0a, 0x07, 0x53, 0x54, 0x5f, 0x4e, 0x45, 0x57, 0x53, 0x10,
	0xef, 0x01, 0x12, 0x11, 0x0a, 0x0c, 0x53, 0x54, 0x5f, 0x45, 0x51, 0x55, 0x49, 0x50, 0x4d, 0x45,
	0x4e, 0x54, 0x10, 0xe3, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x5f, 0x47, 0x41, 0x4d, 0x45,
	0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x06, 0x53, 0x54, 0x5f, 0x41,
	0x53, 0x4b, 0x10, 0xc9, 0x02, 0x2a, 0x54, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x73, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x07, 0x4e, 0x54, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x10, 0xe5, 0x01, 0x12, 0x0e, 0x0a, 0x09, 0x4e,
	0x54, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x10, 0xe7, 0x01, 0x12, 0x11, 0x0a, 0x0c, 0x4e,
	0x54, 0x5f, 0x45, 0x51, 0x55, 0x49, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0xe8, 0x01, 0x12, 0x0b,
	0x0a, 0x06, 0x4e, 0x54, 0x5f, 0x46, 0x55, 0x4e, 0x10, 0xe6, 0x01, 0x32, 0x3b, 0x0a, 0x08, 0x50,
	0x50, 0x53, 0x70, 0x69, 0x64, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x08, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x50, 0x50, 0x57, 0x12, 0x0f, 0x2e, 0x70, 0x70, 0x2e, 0x50, 0x50, 0x53, 0x70, 0x69, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x70, 0x70, 0x2e, 0x50, 0x50, 0x53, 0x70, 0x69, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pp_spider_proto_rawDescOnce sync.Once
	file_pp_spider_proto_rawDescData = file_pp_spider_proto_rawDesc
)

func file_pp_spider_proto_rawDescGZIP() []byte {
	file_pp_spider_proto_rawDescOnce.Do(func() {
		file_pp_spider_proto_rawDescData = protoimpl.X.CompressGZIP(file_pp_spider_proto_rawDescData)
	})
	return file_pp_spider_proto_rawDescData
}

var file_pp_spider_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pp_spider_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pp_spider_proto_goTypes = []interface{}{
	(SectionType)(0),     // 0: pp.SectionType
	(NewsType)(0),        // 1: pp.NewsType
	(*PPSpiderResp)(nil), // 2: pp.PPSpiderResp
	(*PPSpiderBody)(nil), // 3: pp.PPSpiderBody
	(*PPSpiderReq)(nil),  // 4: pp.PPSpiderReq
}
var file_pp_spider_proto_depIdxs = []int32{
	3, // 0: pp.PPSpiderResp.data:type_name -> pp.PPSpiderBody
	0, // 1: pp.PPSpiderReq.section_type:type_name -> pp.SectionType
	1, // 2: pp.PPSpiderReq.news_type:type_name -> pp.NewsType
	4, // 3: pp.PPSpider.FetchPPW:input_type -> pp.PPSpiderReq
	2, // 4: pp.PPSpider.FetchPPW:output_type -> pp.PPSpiderResp
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pp_spider_proto_init() }
func file_pp_spider_proto_init() {
	if File_pp_spider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pp_spider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PPSpiderResp); i {
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
		file_pp_spider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PPSpiderBody); i {
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
		file_pp_spider_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PPSpiderReq); i {
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
			RawDescriptor: file_pp_spider_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pp_spider_proto_goTypes,
		DependencyIndexes: file_pp_spider_proto_depIdxs,
		EnumInfos:         file_pp_spider_proto_enumTypes,
		MessageInfos:      file_pp_spider_proto_msgTypes,
	}.Build()
	File_pp_spider_proto = out.File
	file_pp_spider_proto_rawDesc = nil
	file_pp_spider_proto_goTypes = nil
	file_pp_spider_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PPSpiderClient is the client API for PPSpider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PPSpiderClient interface {
	FetchPPW(ctx context.Context, in *PPSpiderReq, opts ...grpc.CallOption) (*PPSpiderResp, error)
}

type pPSpiderClient struct {
	cc grpc.ClientConnInterface
}

func NewPPSpiderClient(cc grpc.ClientConnInterface) PPSpiderClient {
	return &pPSpiderClient{cc}
}

func (c *pPSpiderClient) FetchPPW(ctx context.Context, in *PPSpiderReq, opts ...grpc.CallOption) (*PPSpiderResp, error) {
	out := new(PPSpiderResp)
	err := c.cc.Invoke(ctx, "/pp.PPSpider/FetchPPW", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PPSpiderServer is the server API for PPSpider service.
type PPSpiderServer interface {
	FetchPPW(context.Context, *PPSpiderReq) (*PPSpiderResp, error)
}

// UnimplementedPPSpiderServer can be embedded to have forward compatible implementations.
type UnimplementedPPSpiderServer struct {
}

func (*UnimplementedPPSpiderServer) FetchPPW(context.Context, *PPSpiderReq) (*PPSpiderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchPPW not implemented")
}

func RegisterPPSpiderServer(s *grpc.Server, srv PPSpiderServer) {
	s.RegisterService(&_PPSpider_serviceDesc, srv)
}

func _PPSpider_FetchPPW_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PPSpiderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PPSpiderServer).FetchPPW(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pp.PPSpider/FetchPPW",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PPSpiderServer).FetchPPW(ctx, req.(*PPSpiderReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _PPSpider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pp.PPSpider",
	HandlerType: (*PPSpiderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchPPW",
			Handler:    _PPSpider_FetchPPW_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pp_spider.proto",
}