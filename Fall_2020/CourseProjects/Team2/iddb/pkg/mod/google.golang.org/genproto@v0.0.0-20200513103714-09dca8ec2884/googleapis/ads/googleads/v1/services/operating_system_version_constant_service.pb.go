// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/operating_system_version_constant_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for
// [OperatingSystemVersionConstantService.GetOperatingSystemVersionConstant][google.ads.googleads.v1.services.OperatingSystemVersionConstantService.GetOperatingSystemVersionConstant].
type GetOperatingSystemVersionConstantRequest struct {
	// Required. Resource name of the OS version to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOperatingSystemVersionConstantRequest) Reset() {
	*m = GetOperatingSystemVersionConstantRequest{}
}
func (m *GetOperatingSystemVersionConstantRequest) String() string { return proto.CompactTextString(m) }
func (*GetOperatingSystemVersionConstantRequest) ProtoMessage()    {}
func (*GetOperatingSystemVersionConstantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_97a0893223528a69, []int{0}
}

func (m *GetOperatingSystemVersionConstantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOperatingSystemVersionConstantRequest.Unmarshal(m, b)
}
func (m *GetOperatingSystemVersionConstantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOperatingSystemVersionConstantRequest.Marshal(b, m, deterministic)
}
func (m *GetOperatingSystemVersionConstantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOperatingSystemVersionConstantRequest.Merge(m, src)
}
func (m *GetOperatingSystemVersionConstantRequest) XXX_Size() int {
	return xxx_messageInfo_GetOperatingSystemVersionConstantRequest.Size(m)
}
func (m *GetOperatingSystemVersionConstantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOperatingSystemVersionConstantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOperatingSystemVersionConstantRequest proto.InternalMessageInfo

func (m *GetOperatingSystemVersionConstantRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetOperatingSystemVersionConstantRequest)(nil), "google.ads.googleads.v1.services.GetOperatingSystemVersionConstantRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/operating_system_version_constant_service.proto", fileDescriptor_97a0893223528a69)
}

var fileDescriptor_97a0893223528a69 = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x3f, 0x8b, 0x13, 0x41,
	0x18, 0xc6, 0xc9, 0x0a, 0x82, 0x8b, 0x36, 0xdb, 0x78, 0x46, 0xc1, 0x78, 0x9c, 0x70, 0x5c, 0x31,
	0xc3, 0x2a, 0x72, 0x38, 0x22, 0x32, 0xb1, 0x88, 0x0a, 0x6a, 0xf0, 0x20, 0x85, 0x04, 0x96, 0xb9,
	0xdd, 0xd7, 0x75, 0x20, 0x3b, 0x6f, 0x9c, 0x77, 0x6e, 0x41, 0xc4, 0x42, 0x3f, 0x81, 0xe2, 0x37,
	0xb0, 0xf4, 0x93, 0xc8, 0xb5, 0x76, 0x56, 0x16, 0x56, 0x7e, 0x04, 0xab, 0x23, 0x99, 0x9d, 0xbd,
	0xa4, 0xc8, 0x9f, 0xee, 0x61, 0xdf, 0x67, 0x7e, 0xef, 0xc3, 0x33, 0xb3, 0xf1, 0xb0, 0x44, 0x2c,
	0x27, 0xc0, 0x55, 0x41, 0xdc, 0xcb, 0x99, 0xaa, 0x53, 0x4e, 0x60, 0x6b, 0x9d, 0x03, 0x71, 0x9c,
	0x82, 0x55, 0x4e, 0x9b, 0x32, 0xa3, 0xf7, 0xe4, 0xa0, 0xca, 0x6a, 0xb0, 0xa4, 0xd1, 0x64, 0x39,
	0x1a, 0x72, 0xca, 0xb8, 0xac, 0xb1, 0xb2, 0xa9, 0x45, 0x87, 0x49, 0xcf, 0x63, 0x98, 0x2a, 0x88,
	0xb5, 0x44, 0x56, 0xa7, 0x2c, 0x10, 0xbb, 0x4f, 0x57, 0xed, 0xb4, 0x40, 0x78, 0x62, 0xb7, 0x5a,
	0xea, 0x97, 0x75, 0x6f, 0x04, 0xd4, 0x54, 0x73, 0x65, 0x0c, 0x3a, 0xe5, 0x34, 0x1a, 0x6a, 0xa6,
	0x57, 0x17, 0xa6, 0xf9, 0x44, 0x43, 0x7b, 0xec, 0xe6, 0xc2, 0xe0, 0x8d, 0x86, 0x49, 0x91, 0x1d,
	0xc3, 0x5b, 0x55, 0x6b, 0xb4, 0x8d, 0xe1, 0xda, 0x82, 0x21, 0xa4, 0xf2, 0xa3, 0xdd, 0x2f, 0x9d,
	0x78, 0x7f, 0x00, 0xee, 0x65, 0x48, 0x78, 0x34, 0x0f, 0x38, 0xf2, 0xf9, 0x1e, 0x37, 0xf1, 0x5e,
	0xc1, 0xbb, 0x13, 0x20, 0x97, 0x14, 0xf1, 0x95, 0x70, 0x3c, 0x33, 0xaa, 0x82, 0x9d, 0x4e, 0xaf,
	0xb3, 0x7f, 0xa9, 0xff, 0xe8, 0x8f, 0x8c, 0xfe, 0xcb, 0xfb, 0xf1, 0xe1, 0x79, 0x41, 0x8d, 0x9a,
	0x6a, 0x62, 0x39, 0x56, 0x7c, 0x03, 0xfe, 0x72, 0xa0, 0xbe, 0x50, 0x15, 0xdc, 0xf9, 0x19, 0xc5,
	0xb7, 0xd7, 0x1f, 0x38, 0xf2, 0xdd, 0x27, 0x9f, 0xa2, 0xf8, 0xd6, 0xc6, 0xf0, 0xc9, 0x33, 0xb6,
	0xe9, 0x0e, 0xd9, 0xb6, 0x0d, 0x74, 0xe5, 0x4a, 0x56, 0x7b, 0xdb, 0x6c, 0x3d, 0x69, 0xf7, 0xf9,
	0x6f, 0xb9, 0xdc, 0xe2, 0xe7, 0x5f, 0x7f, 0xbf, 0x45, 0x87, 0xc9, 0xbd, 0xd9, 0x9b, 0xf9, 0xb0,
	0x34, 0x79, 0x88, 0x6b, 0x51, 0xc4, 0x0f, 0x3e, 0x76, 0xaf, 0x9f, 0xca, 0x9d, 0x55, 0xbd, 0xf7,
	0xbf, 0x46, 0xf1, 0x5e, 0x8e, 0xd5, 0xc6, 0x02, 0xfa, 0x07, 0x5b, 0x15, 0x3e, 0x9c, 0x3d, 0x99,
	0x61, 0xe7, 0xf5, 0x93, 0x86, 0x57, 0xe2, 0x44, 0x99, 0x92, 0xa1, 0x2d, 0x79, 0x09, 0x66, 0xfe,
	0xa0, 0xf8, 0x79, 0x82, 0xd5, 0x7f, 0xe1, 0x83, 0x20, 0xbe, 0x47, 0x17, 0x06, 0x52, 0xfe, 0x88,
	0x7a, 0x03, 0x0f, 0x94, 0x05, 0x31, 0x2f, 0x67, 0x6a, 0x94, 0xb2, 0x66, 0x31, 0x9d, 0x06, 0xcb,
	0x58, 0x16, 0x34, 0x6e, 0x2d, 0xe3, 0x51, 0x3a, 0x0e, 0x96, 0x7f, 0xd1, 0x9e, 0xff, 0x2e, 0x84,
	0x2c, 0x48, 0x88, 0xd6, 0x24, 0xc4, 0x28, 0x15, 0x22, 0xd8, 0x8e, 0x2f, 0xce, 0x73, 0xde, 0x3d,
	0x0b, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x57, 0xee, 0xd8, 0x2c, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OperatingSystemVersionConstantServiceClient is the client API for OperatingSystemVersionConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OperatingSystemVersionConstantServiceClient interface {
	// Returns the requested OS version constant in full detail.
	GetOperatingSystemVersionConstant(ctx context.Context, in *GetOperatingSystemVersionConstantRequest, opts ...grpc.CallOption) (*resources.OperatingSystemVersionConstant, error)
}

type operatingSystemVersionConstantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOperatingSystemVersionConstantServiceClient(cc grpc.ClientConnInterface) OperatingSystemVersionConstantServiceClient {
	return &operatingSystemVersionConstantServiceClient{cc}
}

func (c *operatingSystemVersionConstantServiceClient) GetOperatingSystemVersionConstant(ctx context.Context, in *GetOperatingSystemVersionConstantRequest, opts ...grpc.CallOption) (*resources.OperatingSystemVersionConstant, error) {
	out := new(resources.OperatingSystemVersionConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.OperatingSystemVersionConstantService/GetOperatingSystemVersionConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperatingSystemVersionConstantServiceServer is the server API for OperatingSystemVersionConstantService service.
type OperatingSystemVersionConstantServiceServer interface {
	// Returns the requested OS version constant in full detail.
	GetOperatingSystemVersionConstant(context.Context, *GetOperatingSystemVersionConstantRequest) (*resources.OperatingSystemVersionConstant, error)
}

// UnimplementedOperatingSystemVersionConstantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOperatingSystemVersionConstantServiceServer struct {
}

func (*UnimplementedOperatingSystemVersionConstantServiceServer) GetOperatingSystemVersionConstant(ctx context.Context, req *GetOperatingSystemVersionConstantRequest) (*resources.OperatingSystemVersionConstant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOperatingSystemVersionConstant not implemented")
}

func RegisterOperatingSystemVersionConstantServiceServer(s *grpc.Server, srv OperatingSystemVersionConstantServiceServer) {
	s.RegisterService(&_OperatingSystemVersionConstantService_serviceDesc, srv)
}

func _OperatingSystemVersionConstantService_GetOperatingSystemVersionConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperatingSystemVersionConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatingSystemVersionConstantServiceServer).GetOperatingSystemVersionConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.OperatingSystemVersionConstantService/GetOperatingSystemVersionConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatingSystemVersionConstantServiceServer).GetOperatingSystemVersionConstant(ctx, req.(*GetOperatingSystemVersionConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OperatingSystemVersionConstantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.OperatingSystemVersionConstantService",
	HandlerType: (*OperatingSystemVersionConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOperatingSystemVersionConstant",
			Handler:    _OperatingSystemVersionConstantService_GetOperatingSystemVersionConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/operating_system_version_constant_service.proto",
}
