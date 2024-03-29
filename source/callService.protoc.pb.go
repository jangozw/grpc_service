// Code generated by protoc-gen-go. DO NOT EDIT.
// source: callService.protoc

package grpc_source

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
//const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
	ServiceName          string   `protobuf:"bytes,1,opt,name=ServiceName,proto3" json:"ServiceName,omitempty"`
	MethodName           string   `protobuf:"bytes,2,opt,name=MethodName,proto3" json:"MethodName,omitempty"`
	Params               string   `protobuf:"bytes,3,opt,name=Params,proto3" json:"Params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_778d87365d6beb98, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *Request) GetMethodName() string {
	if m != nil {
		return m.MethodName
	}
	return ""
}

func (m *Request) GetParams() string {
	if m != nil {
		return m.Params
	}
	return ""
}

type Response struct {
	Code                 int32    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=Success,proto3" json:"Success,omitempty"`
	Msg                  string   `protobuf:"bytes,3,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Data                 string   `protobuf:"bytes,4,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_778d87365d6beb98, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Response) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "grpc_source.Request")
	proto.RegisterType((*Response)(nil), "grpc_source.Response")
}

func init() { proto.RegisterFile("callService.protoc", fileDescriptor_778d87365d6beb98) }

var fileDescriptor_778d87365d6beb98 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x29, 0x2d, 0x69, 0xb9, 0x2e, 0x70, 0x02, 0x64, 0x31, 0xa0, 0x2a, 0x13, 0x53, 0x06,
	0x18, 0xf8, 0x01, 0x65, 0x41, 0xa2, 0x08, 0xb9, 0x3b, 0xc8, 0x5c, 0x4f, 0x05, 0x91, 0xc4, 0xc1,
	0xe7, 0xe4, 0xf7, 0xa3, 0x5c, 0x8c, 0x88, 0xba, 0xbd, 0xfb, 0xfc, 0xfc, 0x74, 0xef, 0x00, 0xc9,
	0x95, 0xe5, 0x96, 0x43, 0xf7, 0x45, 0x5c, 0x34, 0xc1, 0x47, 0x4f, 0xb8, 0xdc, 0x87, 0x86, 0xde,
	0xc5, 0xb7, 0x81, 0x38, 0x27, 0x98, 0x5b, 0xfe, 0x69, 0x59, 0x22, 0xae, 0x60, 0x99, 0x9c, 0x2f,
	0xae, 0x62, 0x33, 0x59, 0x4d, 0x6e, 0x4f, 0xed, 0x18, 0xe1, 0x0d, 0xc0, 0x86, 0xe3, 0xa7, 0xdf,
	0xa9, 0xe1, 0x58, 0x0d, 0x23, 0x82, 0x57, 0x90, 0xbd, 0xba, 0xe0, 0x2a, 0x31, 0x53, 0x7d, 0x4b,
	0x53, 0xfe, 0x06, 0x0b, 0xcb, 0xd2, 0xf8, 0x5a, 0x18, 0x11, 0x66, 0x6b, 0xbf, 0x1b, 0xe2, 0x4f,
	0xac, 0x6a, 0x34, 0x30, 0xdf, 0xb6, 0x44, 0x2c, 0xa2, 0xa1, 0x0b, 0xfb, 0x37, 0xe2, 0x19, 0x4c,
	0x37, 0xb2, 0x4f, 0x71, 0xbd, 0xec, 0xff, 0x3f, 0xba, 0xe8, 0xcc, 0x4c, 0x91, 0xea, 0xbb, 0x67,
	0x38, 0x5f, 0xff, 0xf7, 0x1c, 0x16, 0xc2, 0x07, 0xc8, 0x9e, 0xea, 0xce, 0x7f, 0x33, 0x5e, 0x14,
	0xa3, 0xc6, 0x45, 0xaa, 0x7b, 0x7d, 0x79, 0x40, 0x87, 0xfd, 0xf2, 0xa3, 0x8f, 0x4c, 0xef, 0x74,
	0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x07, 0x66, 0x1f, 0x04, 0x3c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CallServiceMethodClient is the client API for CallServiceMethod service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CallServiceMethodClient interface {
	Invoke(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type callServiceMethodClient struct {
	cc *grpc.ClientConn
}

func NewCallServiceMethodClient(cc *grpc.ClientConn) CallServiceMethodClient {
	return &callServiceMethodClient{cc}
}

func (c *callServiceMethodClient) Invoke(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc_source.CallServiceMethod/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CallServiceMethodServer is the server API for CallServiceMethod service.
type CallServiceMethodServer interface {
	Invoke(context.Context, *Request) (*Response, error)
}

func RegisterCallServiceMethodServer(s *grpc.Server, srv CallServiceMethodServer) {
	s.RegisterService(&_CallServiceMethod_serviceDesc, srv)
}

func _CallServiceMethod_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServiceMethodServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_source.CallServiceMethod/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServiceMethodServer).Invoke(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _CallServiceMethod_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_source.CallServiceMethod",
	HandlerType: (*CallServiceMethodServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _CallServiceMethod_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "callService.protoc",
}
