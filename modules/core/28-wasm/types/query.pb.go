// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/core/wasm/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Latest wasm code query
type LatestWASMCodeQuery struct {
	ClientType string `protobuf:"bytes,1,opt,name=client_type,json=clientType,proto3" json:"client_type,omitempty"`
}

func (m *LatestWASMCodeQuery) Reset()         { *m = LatestWASMCodeQuery{} }
func (m *LatestWASMCodeQuery) String() string { return proto.CompactTextString(m) }
func (*LatestWASMCodeQuery) ProtoMessage()    {}
func (*LatestWASMCodeQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_482bc5ce660a9729, []int{0}
}
func (m *LatestWASMCodeQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LatestWASMCodeQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LatestWASMCodeQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LatestWASMCodeQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LatestWASMCodeQuery.Merge(m, src)
}
func (m *LatestWASMCodeQuery) XXX_Size() int {
	return m.Size()
}
func (m *LatestWASMCodeQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_LatestWASMCodeQuery.DiscardUnknown(m)
}

var xxx_messageInfo_LatestWASMCodeQuery proto.InternalMessageInfo

func (m *LatestWASMCodeQuery) GetClientType() string {
	if m != nil {
		return m.ClientType
	}
	return ""
}

// Latest wasm code response
type LatestWASMCodeResponse struct {
	Code []byte `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (m *LatestWASMCodeResponse) Reset()         { *m = LatestWASMCodeResponse{} }
func (m *LatestWASMCodeResponse) String() string { return proto.CompactTextString(m) }
func (*LatestWASMCodeResponse) ProtoMessage()    {}
func (*LatestWASMCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_482bc5ce660a9729, []int{1}
}
func (m *LatestWASMCodeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LatestWASMCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LatestWASMCodeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LatestWASMCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LatestWASMCodeResponse.Merge(m, src)
}
func (m *LatestWASMCodeResponse) XXX_Size() int {
	return m.Size()
}
func (m *LatestWASMCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LatestWASMCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LatestWASMCodeResponse proto.InternalMessageInfo

func (m *LatestWASMCodeResponse) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

// Latest wasm code entry query
type LatestWASMCodeEntryQuery struct {
	ClientType string `protobuf:"bytes,1,opt,name=client_type,json=clientType,proto3" json:"client_type,omitempty"`
}

func (m *LatestWASMCodeEntryQuery) Reset()         { *m = LatestWASMCodeEntryQuery{} }
func (m *LatestWASMCodeEntryQuery) String() string { return proto.CompactTextString(m) }
func (*LatestWASMCodeEntryQuery) ProtoMessage()    {}
func (*LatestWASMCodeEntryQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_482bc5ce660a9729, []int{2}
}
func (m *LatestWASMCodeEntryQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LatestWASMCodeEntryQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LatestWASMCodeEntryQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LatestWASMCodeEntryQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LatestWASMCodeEntryQuery.Merge(m, src)
}
func (m *LatestWASMCodeEntryQuery) XXX_Size() int {
	return m.Size()
}
func (m *LatestWASMCodeEntryQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_LatestWASMCodeEntryQuery.DiscardUnknown(m)
}

var xxx_messageInfo_LatestWASMCodeEntryQuery proto.InternalMessageInfo

func (m *LatestWASMCodeEntryQuery) GetClientType() string {
	if m != nil {
		return m.ClientType
	}
	return ""
}

// Latest wasm code entry response
type LatestWASMCodeEntryResponse struct {
	CodeId string         `protobuf:"bytes,1,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty"`
	Entry  *WasmCodeEntry `protobuf:"bytes,2,opt,name=entry,proto3" json:"entry,omitempty"`
}

func (m *LatestWASMCodeEntryResponse) Reset()         { *m = LatestWASMCodeEntryResponse{} }
func (m *LatestWASMCodeEntryResponse) String() string { return proto.CompactTextString(m) }
func (*LatestWASMCodeEntryResponse) ProtoMessage()    {}
func (*LatestWASMCodeEntryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_482bc5ce660a9729, []int{3}
}
func (m *LatestWASMCodeEntryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LatestWASMCodeEntryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LatestWASMCodeEntryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LatestWASMCodeEntryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LatestWASMCodeEntryResponse.Merge(m, src)
}
func (m *LatestWASMCodeEntryResponse) XXX_Size() int {
	return m.Size()
}
func (m *LatestWASMCodeEntryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LatestWASMCodeEntryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LatestWASMCodeEntryResponse proto.InternalMessageInfo

func (m *LatestWASMCodeEntryResponse) GetCodeId() string {
	if m != nil {
		return m.CodeId
	}
	return ""
}

func (m *LatestWASMCodeEntryResponse) GetEntry() *WasmCodeEntry {
	if m != nil {
		return m.Entry
	}
	return nil
}

func init() {
	proto.RegisterType((*LatestWASMCodeQuery)(nil), "ibc.core.wasm.v1.LatestWASMCodeQuery")
	proto.RegisterType((*LatestWASMCodeResponse)(nil), "ibc.core.wasm.v1.LatestWASMCodeResponse")
	proto.RegisterType((*LatestWASMCodeEntryQuery)(nil), "ibc.core.wasm.v1.LatestWASMCodeEntryQuery")
	proto.RegisterType((*LatestWASMCodeEntryResponse)(nil), "ibc.core.wasm.v1.LatestWASMCodeEntryResponse")
}

func init() { proto.RegisterFile("ibc/core/wasm/v1/query.proto", fileDescriptor_482bc5ce660a9729) }

var fileDescriptor_482bc5ce660a9729 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc9, 0x4c, 0x4a, 0xd6,
	0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0x4f, 0x2c, 0xce, 0xd5, 0x2f, 0x33, 0xd4, 0x2f, 0x2c, 0x4d,
	0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xc8, 0x4c, 0x4a, 0xd6, 0x03, 0xc9,
	0xea, 0x81, 0x64, 0xf5, 0xca, 0x0c, 0xa5, 0xa4, 0x31, 0xd4, 0x83, 0x65, 0xc0, 0xca, 0xa5, 0x64,
	0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b,
	0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0x21, 0xb2, 0x4a, 0x66, 0x5c, 0xc2, 0x3e, 0x89, 0x25, 0xa9,
	0xc5, 0x25, 0xe1, 0x8e, 0xc1, 0xbe, 0xce, 0xf9, 0x29, 0xa9, 0x81, 0x20, 0x9b, 0x84, 0xe4, 0xb9,
	0xb8, 0x93, 0x73, 0x32, 0x53, 0xf3, 0x4a, 0xe2, 0x4b, 0x2a, 0x0b, 0x52, 0x25, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0xb8, 0x20, 0x42, 0x21, 0x95, 0x05, 0xa9, 0x4a, 0x3a, 0x5c, 0x62, 0xa8, 0xfa,
	0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x84, 0xb8, 0x58, 0x92, 0xf3, 0x53, 0x20,
	0x7a, 0x78, 0x82, 0xc0, 0x6c, 0x25, 0x6b, 0x2e, 0x09, 0x54, 0xd5, 0xae, 0x79, 0x25, 0x45, 0x95,
	0x44, 0x5a, 0x95, 0xcb, 0x25, 0x8d, 0x45, 0x33, 0xdc, 0x3e, 0x71, 0x2e, 0x76, 0x90, 0x1d, 0xf1,
	0x99, 0x29, 0x50, 0xbd, 0x6c, 0x20, 0xae, 0x67, 0x8a, 0x90, 0x29, 0x17, 0x6b, 0x2a, 0x48, 0xa5,
	0x04, 0x93, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0xbc, 0x1e, 0x7a, 0xb8, 0xe9, 0x85, 0x27, 0x16, 0xe7,
	0x22, 0x0c, 0x84, 0xa8, 0x36, 0x7a, 0xc4, 0xc4, 0xc5, 0x0a, 0x71, 0xd9, 0x02, 0x46, 0x2e, 0x3e,
	0x54, 0x9b, 0x85, 0x54, 0x31, 0x0d, 0xc1, 0x12, 0x7c, 0x52, 0x1a, 0x84, 0x94, 0xc1, 0x5c, 0xaf,
	0x64, 0xdb, 0x74, 0xf9, 0xc9, 0x64, 0x26, 0x73, 0x21, 0x53, 0x7d, 0xf4, 0x38, 0x4c, 0x4a, 0x2d,
	0x49, 0x34, 0xd4, 0xcf, 0x01, 0x6b, 0x8b, 0x07, 0x89, 0xc5, 0x83, 0xfc, 0xa5, 0x5f, 0x8d, 0x14,
	0x58, 0xb5, 0x42, 0x5b, 0x19, 0xd1, 0xe3, 0x0f, 0xec, 0x17, 0x21, 0x2d, 0x42, 0x0e, 0x40, 0x44,
	0x80, 0x94, 0x2e, 0x51, 0x6a, 0xe1, 0x2e, 0x76, 0x06, 0xbb, 0xd8, 0x56, 0xc8, 0x9a, 0x48, 0x17,
	0xc7, 0x83, 0x83, 0x15, 0xd5, 0xdd, 0x4e, 0xbe, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7,
	0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c,
	0xc7, 0x10, 0x65, 0x9c, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x9c,
	0x5f, 0x9c, 0x9b, 0x5f, 0x0c, 0xb2, 0x47, 0x37, 0x3d, 0x5f, 0x3f, 0x37, 0x3f, 0xa5, 0x34, 0x27,
	0xb5, 0x18, 0x62, 0xa5, 0x91, 0x85, 0x2e, 0xd8, 0x56, 0x90, 0x71, 0xc5, 0x49, 0x6c, 0xe0, 0xc4,
	0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xbc, 0xe8, 0x2d, 0x39, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Query to get latest wasm code for particular client type.
	LatestWASMCode(ctx context.Context, in *LatestWASMCodeQuery, opts ...grpc.CallOption) (*LatestWASMCodeResponse, error)
	// Query for get latest wasm code entry for particular client type
	LatestWASMCodeEntry(ctx context.Context, in *LatestWASMCodeEntryQuery, opts ...grpc.CallOption) (*LatestWASMCodeEntryResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) LatestWASMCode(ctx context.Context, in *LatestWASMCodeQuery, opts ...grpc.CallOption) (*LatestWASMCodeResponse, error) {
	out := new(LatestWASMCodeResponse)
	err := c.cc.Invoke(ctx, "/ibc.core.wasm.v1.Query/LatestWASMCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) LatestWASMCodeEntry(ctx context.Context, in *LatestWASMCodeEntryQuery, opts ...grpc.CallOption) (*LatestWASMCodeEntryResponse, error) {
	out := new(LatestWASMCodeEntryResponse)
	err := c.cc.Invoke(ctx, "/ibc.core.wasm.v1.Query/LatestWASMCodeEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Query to get latest wasm code for particular client type.
	LatestWASMCode(context.Context, *LatestWASMCodeQuery) (*LatestWASMCodeResponse, error)
	// Query for get latest wasm code entry for particular client type
	LatestWASMCodeEntry(context.Context, *LatestWASMCodeEntryQuery) (*LatestWASMCodeEntryResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) LatestWASMCode(ctx context.Context, req *LatestWASMCodeQuery) (*LatestWASMCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LatestWASMCode not implemented")
}
func (*UnimplementedQueryServer) LatestWASMCodeEntry(ctx context.Context, req *LatestWASMCodeEntryQuery) (*LatestWASMCodeEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LatestWASMCodeEntry not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_LatestWASMCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LatestWASMCodeQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).LatestWASMCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.core.wasm.v1.Query/LatestWASMCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).LatestWASMCode(ctx, req.(*LatestWASMCodeQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_LatestWASMCodeEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LatestWASMCodeEntryQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).LatestWASMCodeEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.core.wasm.v1.Query/LatestWASMCodeEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).LatestWASMCodeEntry(ctx, req.(*LatestWASMCodeEntryQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ibc.core.wasm.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LatestWASMCode",
			Handler:    _Query_LatestWASMCode_Handler,
		},
		{
			MethodName: "LatestWASMCodeEntry",
			Handler:    _Query_LatestWASMCodeEntry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc/core/wasm/v1/query.proto",
}

func (m *LatestWASMCodeQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LatestWASMCodeQuery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LatestWASMCodeQuery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClientType) > 0 {
		i -= len(m.ClientType)
		copy(dAtA[i:], m.ClientType)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ClientType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LatestWASMCodeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LatestWASMCodeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LatestWASMCodeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Code) > 0 {
		i -= len(m.Code)
		copy(dAtA[i:], m.Code)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Code)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LatestWASMCodeEntryQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LatestWASMCodeEntryQuery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LatestWASMCodeEntryQuery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClientType) > 0 {
		i -= len(m.ClientType)
		copy(dAtA[i:], m.ClientType)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ClientType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LatestWASMCodeEntryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LatestWASMCodeEntryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LatestWASMCodeEntryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Entry != nil {
		{
			size, err := m.Entry.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.CodeId) > 0 {
		i -= len(m.CodeId)
		copy(dAtA[i:], m.CodeId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.CodeId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LatestWASMCodeQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientType)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *LatestWASMCodeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *LatestWASMCodeEntryQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientType)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *LatestWASMCodeEntryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CodeId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Entry != nil {
		l = m.Entry.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LatestWASMCodeQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LatestWASMCodeQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LatestWASMCodeQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LatestWASMCodeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LatestWASMCodeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LatestWASMCodeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = append(m.Code[:0], dAtA[iNdEx:postIndex]...)
			if m.Code == nil {
				m.Code = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LatestWASMCodeEntryQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LatestWASMCodeEntryQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LatestWASMCodeEntryQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LatestWASMCodeEntryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LatestWASMCodeEntryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LatestWASMCodeEntryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CodeId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Entry == nil {
				m.Entry = &WasmCodeEntry{}
			}
			if err := m.Entry.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)