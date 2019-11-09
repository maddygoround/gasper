// Code generated by protoc-gen-go. DO NOT EDIT.
// source: database.proto

package database

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RequestBody struct {
	Language             string   `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestBody) Reset()         { *m = RequestBody{} }
func (m *RequestBody) String() string { return proto.CompactTextString(m) }
func (*RequestBody) ProtoMessage()    {}
func (*RequestBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_b90fe3356ea5df07, []int{0}
}

func (m *RequestBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestBody.Unmarshal(m, b)
}
func (m *RequestBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestBody.Marshal(b, m, deterministic)
}
func (m *RequestBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestBody.Merge(m, src)
}
func (m *RequestBody) XXX_Size() int {
	return xxx_messageInfo_RequestBody.Size(m)
}
func (m *RequestBody) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestBody.DiscardUnknown(m)
}

var xxx_messageInfo_RequestBody proto.InternalMessageInfo

func (m *RequestBody) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *RequestBody) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *RequestBody) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ResponseBody struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseBody) Reset()         { *m = ResponseBody{} }
func (m *ResponseBody) String() string { return proto.CompactTextString(m) }
func (*ResponseBody) ProtoMessage()    {}
func (*ResponseBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_b90fe3356ea5df07, []int{1}
}

func (m *ResponseBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseBody.Unmarshal(m, b)
}
func (m *ResponseBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseBody.Marshal(b, m, deterministic)
}
func (m *ResponseBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseBody.Merge(m, src)
}
func (m *ResponseBody) XXX_Size() int {
	return xxx_messageInfo_ResponseBody.Size(m)
}
func (m *ResponseBody) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseBody.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseBody proto.InternalMessageInfo

func (m *ResponseBody) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type NameHolder struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NameHolder) Reset()         { *m = NameHolder{} }
func (m *NameHolder) String() string { return proto.CompactTextString(m) }
func (*NameHolder) ProtoMessage()    {}
func (*NameHolder) Descriptor() ([]byte, []int) {
	return fileDescriptor_b90fe3356ea5df07, []int{2}
}

func (m *NameHolder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameHolder.Unmarshal(m, b)
}
func (m *NameHolder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameHolder.Marshal(b, m, deterministic)
}
func (m *NameHolder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameHolder.Merge(m, src)
}
func (m *NameHolder) XXX_Size() int {
	return xxx_messageInfo_NameHolder.Size(m)
}
func (m *NameHolder) XXX_DiscardUnknown() {
	xxx_messageInfo_NameHolder.DiscardUnknown(m)
}

var xxx_messageInfo_NameHolder proto.InternalMessageInfo

func (m *NameHolder) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type LanguageHolder struct {
	Language             string   `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LanguageHolder) Reset()         { *m = LanguageHolder{} }
func (m *LanguageHolder) String() string { return proto.CompactTextString(m) }
func (*LanguageHolder) ProtoMessage()    {}
func (*LanguageHolder) Descriptor() ([]byte, []int) {
	return fileDescriptor_b90fe3356ea5df07, []int{3}
}

func (m *LanguageHolder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LanguageHolder.Unmarshal(m, b)
}
func (m *LanguageHolder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LanguageHolder.Marshal(b, m, deterministic)
}
func (m *LanguageHolder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LanguageHolder.Merge(m, src)
}
func (m *LanguageHolder) XXX_Size() int {
	return xxx_messageInfo_LanguageHolder.Size(m)
}
func (m *LanguageHolder) XXX_DiscardUnknown() {
	xxx_messageInfo_LanguageHolder.DiscardUnknown(m)
}

var xxx_messageInfo_LanguageHolder proto.InternalMessageInfo

func (m *LanguageHolder) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

type GenericResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenericResponse) Reset()         { *m = GenericResponse{} }
func (m *GenericResponse) String() string { return proto.CompactTextString(m) }
func (*GenericResponse) ProtoMessage()    {}
func (*GenericResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b90fe3356ea5df07, []int{4}
}

func (m *GenericResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenericResponse.Unmarshal(m, b)
}
func (m *GenericResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenericResponse.Marshal(b, m, deterministic)
}
func (m *GenericResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericResponse.Merge(m, src)
}
func (m *GenericResponse) XXX_Size() int {
	return xxx_messageInfo_GenericResponse.Size(m)
}
func (m *GenericResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenericResponse proto.InternalMessageInfo

func (m *GenericResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type LogRequest struct {
	Language             string   `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
	Tail                 string   `protobuf:"bytes,2,opt,name=tail,proto3" json:"tail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogRequest) Reset()         { *m = LogRequest{} }
func (m *LogRequest) String() string { return proto.CompactTextString(m) }
func (*LogRequest) ProtoMessage()    {}
func (*LogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b90fe3356ea5df07, []int{5}
}

func (m *LogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogRequest.Unmarshal(m, b)
}
func (m *LogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogRequest.Marshal(b, m, deterministic)
}
func (m *LogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogRequest.Merge(m, src)
}
func (m *LogRequest) XXX_Size() int {
	return xxx_messageInfo_LogRequest.Size(m)
}
func (m *LogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogRequest proto.InternalMessageInfo

func (m *LogRequest) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *LogRequest) GetTail() string {
	if m != nil {
		return m.Tail
	}
	return ""
}

type LogResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Data                 []string `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogResponse) Reset()         { *m = LogResponse{} }
func (m *LogResponse) String() string { return proto.CompactTextString(m) }
func (*LogResponse) ProtoMessage()    {}
func (*LogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b90fe3356ea5df07, []int{6}
}

func (m *LogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogResponse.Unmarshal(m, b)
}
func (m *LogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogResponse.Marshal(b, m, deterministic)
}
func (m *LogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogResponse.Merge(m, src)
}
func (m *LogResponse) XXX_Size() int {
	return xxx_messageInfo_LogResponse.Size(m)
}
func (m *LogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogResponse proto.InternalMessageInfo

func (m *LogResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *LogResponse) GetData() []string {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*RequestBody)(nil), "database.RequestBody")
	proto.RegisterType((*ResponseBody)(nil), "database.ResponseBody")
	proto.RegisterType((*NameHolder)(nil), "database.NameHolder")
	proto.RegisterType((*LanguageHolder)(nil), "database.LanguageHolder")
	proto.RegisterType((*GenericResponse)(nil), "database.GenericResponse")
	proto.RegisterType((*LogRequest)(nil), "database.LogRequest")
	proto.RegisterType((*LogResponse)(nil), "database.LogResponse")
}

func init() { proto.RegisterFile("database.proto", fileDescriptor_b90fe3356ea5df07) }

var fileDescriptor_b90fe3356ea5df07 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4f, 0x4f, 0xfa, 0x40,
	0x10, 0xa5, 0xc0, 0xaf, 0x3f, 0x18, 0x08, 0x24, 0x13, 0x30, 0xb5, 0xa7, 0x66, 0x4f, 0x24, 0x1a,
	0x0e, 0x7a, 0x52, 0x4c, 0x4c, 0x94, 0xa0, 0x87, 0xc6, 0xc3, 0xfa, 0x09, 0x96, 0x76, 0x52, 0x49,
	0x4a, 0x17, 0xbb, 0xdb, 0x18, 0xbe, 0x80, 0x9f, 0xdb, 0x74, 0xe9, 0x3f, 0x39, 0xd4, 0xdb, 0xcc,
	0xf4, 0xbd, 0x79, 0x7d, 0x6f, 0x16, 0x26, 0xa1, 0xd0, 0x62, 0x2b, 0x14, 0x2d, 0x0f, 0xa9, 0xd4,
	0x12, 0x07, 0x65, 0xcf, 0xde, 0x61, 0xc4, 0xe9, 0x33, 0x23, 0xa5, 0x9f, 0x64, 0x78, 0x44, 0x17,
	0x06, 0xb1, 0x48, 0xa2, 0x4c, 0x44, 0xe4, 0x58, 0x9e, 0xb5, 0x18, 0xf2, 0xaa, 0xc7, 0x19, 0xfc,
	0x93, 0x5f, 0x09, 0xa5, 0x4e, 0xd7, 0x7c, 0x38, 0x35, 0x88, 0xd0, 0xcf, 0x97, 0x39, 0x3d, 0xcf,
	0x5a, 0x8c, 0xb9, 0xa9, 0x19, 0x83, 0x31, 0x27, 0x75, 0x90, 0x89, 0x22, 0xb3, 0xb5, 0xc4, 0x58,
	0x0d, 0x8c, 0x07, 0xf0, 0x26, 0xf6, 0xf4, 0x2a, 0xe3, 0xf0, 0xb4, 0x25, 0x11, 0xfb, 0x52, 0xd3,
	0xd4, 0xec, 0x1a, 0x26, 0x7e, 0xa1, 0x5d, 0xa0, 0x5a, 0xfe, 0x8e, 0x5d, 0xc1, 0xf4, 0x85, 0x12,
	0x4a, 0x77, 0x41, 0x29, 0x8d, 0x0e, 0xfc, 0x57, 0x59, 0x10, 0x90, 0x52, 0x06, 0x3d, 0xe0, 0x65,
	0xcb, 0x1e, 0x00, 0x7c, 0x19, 0x15, 0xc6, 0x5b, 0x4d, 0x23, 0xf4, 0xb5, 0xd8, 0xc5, 0x85, 0x67,
	0x53, 0xb3, 0x15, 0x8c, 0x0c, 0xfb, 0x2f, 0x99, 0xca, 0x77, 0xd7, 0xeb, 0xe5, 0xe4, 0xbc, 0xbe,
	0xf9, 0xee, 0xc2, 0x74, 0x5d, 0xa4, 0xbf, 0x11, 0x81, 0x96, 0xe9, 0x11, 0xef, 0xc0, 0x7e, 0x4e,
	0x49, 0x68, 0xc2, 0xf9, 0xb2, 0xba, 0x54, 0xe3, 0x2c, 0xee, 0x45, 0x73, 0x5c, 0x07, 0xcb, 0x3a,
	0xb8, 0x02, 0x7b, 0x4d, 0x31, 0x69, 0xc2, 0x59, 0x8d, 0xa9, 0x83, 0x75, 0x2f, 0xeb, 0xe9, 0x59,
	0x3c, 0xac, 0x83, 0xf7, 0x30, 0xdc, 0x90, 0x0e, 0x3e, 0x7c, 0x19, 0xa9, 0x26, 0xbf, 0xce, 0xc6,
	0x9d, 0x9f, 0x4d, 0x2b, 0xee, 0x23, 0xd8, 0x9c, 0x62, 0x29, 0x42, 0x74, 0x1a, 0x90, 0x5f, 0xf7,
	0x6a, 0x15, 0xdf, 0xda, 0xe6, 0x29, 0xde, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xea, 0xa9, 0x3a,
	0x16, 0x9c, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DatabaseFactoryClient is the client API for DatabaseFactory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DatabaseFactoryClient interface {
	Create(ctx context.Context, in *RequestBody, opts ...grpc.CallOption) (*ResponseBody, error)
	Delete(ctx context.Context, in *NameHolder, opts ...grpc.CallOption) (*GenericResponse, error)
	FetchLogs(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error)
	Reload(ctx context.Context, in *LanguageHolder, opts ...grpc.CallOption) (*GenericResponse, error)
}

type databaseFactoryClient struct {
	cc *grpc.ClientConn
}

func NewDatabaseFactoryClient(cc *grpc.ClientConn) DatabaseFactoryClient {
	return &databaseFactoryClient{cc}
}

func (c *databaseFactoryClient) Create(ctx context.Context, in *RequestBody, opts ...grpc.CallOption) (*ResponseBody, error) {
	out := new(ResponseBody)
	err := c.cc.Invoke(ctx, "/database.DatabaseFactory/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseFactoryClient) Delete(ctx context.Context, in *NameHolder, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/database.DatabaseFactory/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseFactoryClient) FetchLogs(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error) {
	out := new(LogResponse)
	err := c.cc.Invoke(ctx, "/database.DatabaseFactory/FetchLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseFactoryClient) Reload(ctx context.Context, in *LanguageHolder, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/database.DatabaseFactory/Reload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseFactoryServer is the server API for DatabaseFactory service.
type DatabaseFactoryServer interface {
	Create(context.Context, *RequestBody) (*ResponseBody, error)
	Delete(context.Context, *NameHolder) (*GenericResponse, error)
	FetchLogs(context.Context, *LogRequest) (*LogResponse, error)
	Reload(context.Context, *LanguageHolder) (*GenericResponse, error)
}

// UnimplementedDatabaseFactoryServer can be embedded to have forward compatible implementations.
type UnimplementedDatabaseFactoryServer struct {
}

func (*UnimplementedDatabaseFactoryServer) Create(ctx context.Context, req *RequestBody) (*ResponseBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedDatabaseFactoryServer) Delete(ctx context.Context, req *NameHolder) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedDatabaseFactoryServer) FetchLogs(ctx context.Context, req *LogRequest) (*LogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchLogs not implemented")
}
func (*UnimplementedDatabaseFactoryServer) Reload(ctx context.Context, req *LanguageHolder) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reload not implemented")
}

func RegisterDatabaseFactoryServer(s *grpc.Server, srv DatabaseFactoryServer) {
	s.RegisterService(&_DatabaseFactory_serviceDesc, srv)
}

func _DatabaseFactory_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseFactoryServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/database.DatabaseFactory/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseFactoryServer).Create(ctx, req.(*RequestBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseFactory_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameHolder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseFactoryServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/database.DatabaseFactory/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseFactoryServer).Delete(ctx, req.(*NameHolder))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseFactory_FetchLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseFactoryServer).FetchLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/database.DatabaseFactory/FetchLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseFactoryServer).FetchLogs(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseFactory_Reload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LanguageHolder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseFactoryServer).Reload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/database.DatabaseFactory/Reload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseFactoryServer).Reload(ctx, req.(*LanguageHolder))
	}
	return interceptor(ctx, in, info, handler)
}

var _DatabaseFactory_serviceDesc = grpc.ServiceDesc{
	ServiceName: "database.DatabaseFactory",
	HandlerType: (*DatabaseFactoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _DatabaseFactory_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DatabaseFactory_Delete_Handler,
		},
		{
			MethodName: "FetchLogs",
			Handler:    _DatabaseFactory_FetchLogs_Handler,
		},
		{
			MethodName: "Reload",
			Handler:    _DatabaseFactory_Reload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "database.proto",
}