// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package graphgrpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
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

type Tenant struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tenant) Reset()         { *m = Tenant{} }
func (m *Tenant) String() string { return proto.CompactTextString(m) }
func (*Tenant) ProtoMessage()    {}
func (*Tenant) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}

func (m *Tenant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tenant.Unmarshal(m, b)
}
func (m *Tenant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tenant.Marshal(b, m, deterministic)
}
func (m *Tenant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tenant.Merge(m, src)
}
func (m *Tenant) XXX_Size() int {
	return xxx_messageInfo_Tenant.Size(m)
}
func (m *Tenant) XXX_DiscardUnknown() {
	xxx_messageInfo_Tenant.DiscardUnknown(m)
}

var xxx_messageInfo_Tenant proto.InternalMessageInfo

func (m *Tenant) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Tenant) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TenantList struct {
	Tenants              []*Tenant `protobuf:"bytes,1,rep,name=tenants,proto3" json:"tenants,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *TenantList) Reset()         { *m = TenantList{} }
func (m *TenantList) String() string { return proto.CompactTextString(m) }
func (*TenantList) ProtoMessage()    {}
func (*TenantList) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{1}
}

func (m *TenantList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TenantList.Unmarshal(m, b)
}
func (m *TenantList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TenantList.Marshal(b, m, deterministic)
}
func (m *TenantList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TenantList.Merge(m, src)
}
func (m *TenantList) XXX_Size() int {
	return xxx_messageInfo_TenantList.Size(m)
}
func (m *TenantList) XXX_DiscardUnknown() {
	xxx_messageInfo_TenantList.DiscardUnknown(m)
}

var xxx_messageInfo_TenantList proto.InternalMessageInfo

func (m *TenantList) GetTenants() []*Tenant {
	if m != nil {
		return m.Tenants
	}
	return nil
}

func init() {
	proto.RegisterType((*Tenant)(nil), "graph.Tenant")
	proto.RegisterType((*TenantList)(nil), "graph.TenantList")
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1) }

var fileDescriptor_77a6da22d6a3feb1 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x8f, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0x49, 0xda, 0x37, 0xaf, 0x1d, 0xa9, 0xe0, 0x1e, 0x24, 0xc4, 0x52, 0x4a, 0x2f, 0xf6,
	0x20, 0x1b, 0x8c, 0xf4, 0x2a, 0xe2, 0x1f, 0xbc, 0x78, 0x6a, 0x8b, 0x07, 0x6f, 0x9b, 0x74, 0x9a,
	0x2e, 0x24, 0xbb, 0xcb, 0x64, 0xa3, 0xe4, 0x53, 0xf9, 0x15, 0xc5, 0x5d, 0x7a, 0xa8, 0x22, 0xa2,
	0x97, 0x90, 0x7d, 0x9e, 0xdf, 0xec, 0xfc, 0x16, 0x06, 0x64, 0x0a, 0x6e, 0x48, 0x5b, 0xcd, 0xfe,
	0x95, 0x24, 0xcc, 0x36, 0x39, 0x2d, 0xb5, 0x2e, 0x2b, 0x4c, 0x5d, 0x98, 0xb7, 0x9b, 0x14, 0x6b,
	0x63, 0x3b, 0xcf, 0x24, 0xe3, 0xcf, 0xe5, 0x2b, 0x09, 0x63, 0x90, 0x1a, 0xdf, 0x4f, 0xcf, 0x21,
	0x5a, 0xa1, 0x12, 0xca, 0xb2, 0x23, 0x08, 0xe5, 0x3a, 0x0e, 0x26, 0xc1, 0x6c, 0xb0, 0x08, 0xe5,
	0x9a, 0x31, 0xe8, 0x2b, 0x51, 0x63, 0x1c, 0xba, 0xc4, 0xfd, 0x4f, 0xe7, 0x00, 0x9e, 0x7e, 0x94,
	0x8d, 0x65, 0x67, 0xf0, 0xdf, 0xba, 0x53, 0x13, 0x07, 0x93, 0xde, 0xec, 0x30, 0x1b, 0x72, 0x67,
	0xc4, 0x3d, 0xb3, 0xd8, 0xb5, 0xd9, 0x5b, 0x08, 0x43, 0x9f, 0x2d, 0x91, 0x5e, 0x64, 0x81, 0x6c,
	0x0e, 0xd1, 0x2d, 0xa1, 0xb0, 0xc8, 0x46, 0xdc, 0x1b, 0xf2, 0x9d, 0x21, 0x5f, 0x5a, 0x92, 0xaa,
	0x7c, 0x12, 0x55, 0x8b, 0xc9, 0xfe, 0x8d, 0xec, 0x02, 0xfa, 0x6e, 0xf3, 0xc9, 0x97, 0xa1, 0xfb,
	0x8f, 0x37, 0x27, 0xc7, 0x7b, 0xb8, 0x43, 0x33, 0xe8, 0x3d, 0xa0, 0xfd, 0xdd, 0x9a, 0x6b, 0x38,
	0x58, 0x51, 0xab, 0x8a, 0x9f, 0xfd, 0xbe, 0x11, 0x61, 0x57, 0x10, 0xdd, 0x61, 0x85, 0x7f, 0x9d,
	0xbf, 0x19, 0x3f, 0x8f, 0x36, 0x79, 0x91, 0x36, 0x5d, 0x6d, 0xb6, 0x5a, 0x75, 0xa9, 0xd3, 0xf3,
	0xdf, 0x92, 0x4c, 0x91, 0x47, 0x8e, 0xbf, 0x7c, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xa5, 0xe1,
	0x66, 0x0e, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TenantServiceClient is the client API for TenantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TenantServiceClient interface {
	Create(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*Tenant, error)
	List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TenantList, error)
	Get(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*Tenant, error)
	Truncate(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*empty.Empty, error)
	Delete(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*empty.Empty, error)
}

type tenantServiceClient struct {
	cc *grpc.ClientConn
}

func NewTenantServiceClient(cc *grpc.ClientConn) TenantServiceClient {
	return &tenantServiceClient{cc}
}

func (c *tenantServiceClient) Create(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, "/graph.TenantService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TenantList, error) {
	out := new(TenantList)
	err := c.cc.Invoke(ctx, "/graph.TenantService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) Get(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, "/graph.TenantService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) Truncate(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/graph.TenantService/Truncate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) Delete(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/graph.TenantService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TenantServiceServer is the server API for TenantService service.
type TenantServiceServer interface {
	Create(context.Context, *wrappers.StringValue) (*Tenant, error)
	List(context.Context, *empty.Empty) (*TenantList, error)
	Get(context.Context, *wrappers.StringValue) (*Tenant, error)
	Truncate(context.Context, *wrappers.StringValue) (*empty.Empty, error)
	Delete(context.Context, *wrappers.StringValue) (*empty.Empty, error)
}

func RegisterTenantServiceServer(s *grpc.Server, srv TenantServiceServer) {
	s.RegisterService(&_TenantService_serviceDesc, srv)
}

func _TenantService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/graph.TenantService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).Create(ctx, req.(*wrappers.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/graph.TenantService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).List(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/graph.TenantService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).Get(ctx, req.(*wrappers.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_Truncate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).Truncate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/graph.TenantService/Truncate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).Truncate(ctx, req.(*wrappers.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/graph.TenantService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).Delete(ctx, req.(*wrappers.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

var _TenantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "graph.TenantService",
	HandlerType: (*TenantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TenantService_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _TenantService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TenantService_Get_Handler,
		},
		{
			MethodName: "Truncate",
			Handler:    _TenantService_Truncate_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TenantService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}