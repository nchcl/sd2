// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: chat.proto

package chat

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Parts int32  `protobuf:"varint,2,opt,name=parts,proto3" json:"parts,omitempty"`
	Data  []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Chunk) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Chunk) GetParts() int32 {
	if x != nil {
		return x.Parts
	}
	return 0
}

func (x *Chunk) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Signal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Body     string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Response bool   `protobuf:"varint,3,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *Signal) Reset() {
	*x = Signal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signal) ProtoMessage() {}

func (x *Signal) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signal.ProtoReflect.Descriptor instead.
func (*Signal) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *Signal) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Signal) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Signal) GetResponse() bool {
	if x != nil {
		return x.Response
	}
	return false
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x68,
	0x61, 0x74, 0x22, 0x45, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x70, 0x61, 0x72, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x48, 0x0a, 0x06, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0xcf, 0x01, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x73, 0x12, 0x0b, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x0c,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x22, 0x00, 0x12, 0x32,
	0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x4c,
	0x69, 0x73, 0x74, 0x61, 0x12, 0x0c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x6c, 0x1a, 0x0c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x22, 0x00, 0x12, 0x2f, 0x0a, 0x0f, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x70,
	0x75, 0x65, 0x73, 0x74, 0x61, 0x12, 0x0c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x1a, 0x0c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x10, 0x52, 0x65, 0x63, 0x69, 0x62, 0x69, 0x72, 0x50, 0x72,
	0x6f, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x12, 0x0c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x6c, 0x1a, 0x0c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x63, 0x68, 0x63, 0x6c, 0x2f, 0x73, 0x64, 0x3b, 0x63, 0x68, 0x61,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_chat_proto_goTypes = []interface{}{
	(*Chunk)(nil),  // 0: chat.Chunk
	(*Signal)(nil), // 1: chat.Signal
}
var file_chat_proto_depIdxs = []int32{
	0, // 0: chat.ChatService.SendChunks:input_type -> chat.Chunk
	1, // 1: chat.ChatService.TransferenciaLista:input_type -> chat.Signal
	1, // 2: chat.ChatService.EnviarPropuesta:input_type -> chat.Signal
	1, // 3: chat.ChatService.RecibirPropuesta:input_type -> chat.Signal
	1, // 4: chat.ChatService.SendChunks:output_type -> chat.Signal
	1, // 5: chat.ChatService.TransferenciaLista:output_type -> chat.Signal
	1, // 6: chat.ChatService.EnviarPropuesta:output_type -> chat.Signal
	1, // 7: chat.ChatService.RecibirPropuesta:output_type -> chat.Signal
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signal); i {
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
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	SendChunks(ctx context.Context, in *Chunk, opts ...grpc.CallOption) (*Signal, error)
	TransferenciaLista(ctx context.Context, in *Signal, opts ...grpc.CallOption) (*Signal, error)
	EnviarPropuesta(ctx context.Context, in *Signal, opts ...grpc.CallOption) (*Signal, error)
	RecibirPropuesta(ctx context.Context, in *Signal, opts ...grpc.CallOption) (*Signal, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) SendChunks(ctx context.Context, in *Chunk, opts ...grpc.CallOption) (*Signal, error) {
	out := new(Signal)
	err := c.cc.Invoke(ctx, "/chat.ChatService/SendChunks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) TransferenciaLista(ctx context.Context, in *Signal, opts ...grpc.CallOption) (*Signal, error) {
	out := new(Signal)
	err := c.cc.Invoke(ctx, "/chat.ChatService/TransferenciaLista", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) EnviarPropuesta(ctx context.Context, in *Signal, opts ...grpc.CallOption) (*Signal, error) {
	out := new(Signal)
	err := c.cc.Invoke(ctx, "/chat.ChatService/EnviarPropuesta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) RecibirPropuesta(ctx context.Context, in *Signal, opts ...grpc.CallOption) (*Signal, error) {
	out := new(Signal)
	err := c.cc.Invoke(ctx, "/chat.ChatService/RecibirPropuesta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	SendChunks(context.Context, *Chunk) (*Signal, error)
	TransferenciaLista(context.Context, *Signal) (*Signal, error)
	EnviarPropuesta(context.Context, *Signal) (*Signal, error)
	RecibirPropuesta(context.Context, *Signal) (*Signal, error)
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) SendChunks(context.Context, *Chunk) (*Signal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendChunks not implemented")
}
func (*UnimplementedChatServiceServer) TransferenciaLista(context.Context, *Signal) (*Signal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferenciaLista not implemented")
}
func (*UnimplementedChatServiceServer) EnviarPropuesta(context.Context, *Signal) (*Signal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarPropuesta not implemented")
}
func (*UnimplementedChatServiceServer) RecibirPropuesta(context.Context, *Signal) (*Signal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecibirPropuesta not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_SendChunks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Chunk)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SendChunks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/SendChunks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SendChunks(ctx, req.(*Chunk))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_TransferenciaLista_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Signal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).TransferenciaLista(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/TransferenciaLista",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).TransferenciaLista(ctx, req.(*Signal))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_EnviarPropuesta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Signal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).EnviarPropuesta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/EnviarPropuesta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).EnviarPropuesta(ctx, req.(*Signal))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_RecibirPropuesta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Signal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).RecibirPropuesta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/RecibirPropuesta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).RecibirPropuesta(ctx, req.(*Signal))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendChunks",
			Handler:    _ChatService_SendChunks_Handler,
		},
		{
			MethodName: "TransferenciaLista",
			Handler:    _ChatService_TransferenciaLista_Handler,
		},
		{
			MethodName: "EnviarPropuesta",
			Handler:    _ChatService_EnviarPropuesta_Handler,
		},
		{
			MethodName: "RecibirPropuesta",
			Handler:    _ChatService_RecibirPropuesta_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}