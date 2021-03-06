// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/template/checknothing/template_handler_service.proto

// The `checknothing` template represents an empty block of data, which can be useful
// in different testing scenarios.
//
// Example config:
//
// ```yaml
// apiVersion: "config.istio.io/v1alpha2"
// kind: instance
// metadata:
//   name: denyrequest
//   namespace: istio-system
// spec:
//   compiledTemplate: checknothing
// ```
//
// CheckNothing represents an empty block of data that is used for Check-capable
// adapters which don't require any parameters. This is primarily intended for testing
// scenarios.

package checknothing

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	v1beta1 "istio.io/api/mixer/adapter/model/v1beta1"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

// Request message for HandleCheckNothing method.
type HandleCheckNothingRequest struct {
	// 'checknothing' instance.
	Instance *InstanceMsg `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	// Adapter specific handler configuration.
	//
	// Note: Backends can also implement [InfrastructureBackend][https://istio.io/docs/reference/config/mixer/istio.mixer.adapter.model.v1beta1.html#InfrastructureBackend]
	// service and therefore opt to receive handler configuration during session creation through [InfrastructureBackend.CreateSession][TODO: Link to this fragment]
	// call. In that case, adapter_config will have type_url as 'google.protobuf.Any.type_url' and would contain string
	// value of session_id (returned from InfrastructureBackend.CreateSession).
	AdapterConfig *types.Any `protobuf:"bytes,2,opt,name=adapter_config,json=adapterConfig,proto3" json:"adapter_config,omitempty"`
	// Id to dedupe identical requests from Mixer.
	DedupId string `protobuf:"bytes,3,opt,name=dedup_id,json=dedupId,proto3" json:"dedup_id,omitempty"`
}

func (m *HandleCheckNothingRequest) Reset()      { *m = HandleCheckNothingRequest{} }
func (*HandleCheckNothingRequest) ProtoMessage() {}
func (*HandleCheckNothingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8608595bf414489a, []int{0}
}
func (m *HandleCheckNothingRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HandleCheckNothingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *HandleCheckNothingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandleCheckNothingRequest.Merge(m, src)
}
func (m *HandleCheckNothingRequest) XXX_Size() int {
	return m.Size()
}
func (m *HandleCheckNothingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HandleCheckNothingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HandleCheckNothingRequest proto.InternalMessageInfo

// Contains instance payload for 'checknothing' template. This is passed to infrastructure backends during request-time
// through HandleCheckNothingService.HandleCheckNothing.
type InstanceMsg struct {
	// Name of the instance as specified in configuration.
	Name string `protobuf:"bytes,72295727,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *InstanceMsg) Reset()      { *m = InstanceMsg{} }
func (*InstanceMsg) ProtoMessage() {}
func (*InstanceMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_8608595bf414489a, []int{1}
}
func (m *InstanceMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InstanceMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *InstanceMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstanceMsg.Merge(m, src)
}
func (m *InstanceMsg) XXX_Size() int {
	return m.Size()
}
func (m *InstanceMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_InstanceMsg.DiscardUnknown(m)
}

var xxx_messageInfo_InstanceMsg proto.InternalMessageInfo

// Contains inferred type information about specific instance of 'checknothing' template. This is passed to
// infrastructure backends during configuration-time through [InfrastructureBackend.CreateSession][TODO: Link to this fragment].
type Type struct {
}

func (m *Type) Reset()      { *m = Type{} }
func (*Type) ProtoMessage() {}
func (*Type) Descriptor() ([]byte, []int) {
	return fileDescriptor_8608595bf414489a, []int{2}
}
func (m *Type) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Type) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Type) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Type.Merge(m, src)
}
func (m *Type) XXX_Size() int {
	return m.Size()
}
func (m *Type) XXX_DiscardUnknown() {
	xxx_messageInfo_Type.DiscardUnknown(m)
}

var xxx_messageInfo_Type proto.InternalMessageInfo

// Represents instance configuration schema for 'checknothing' template.
type InstanceParam struct {
}

func (m *InstanceParam) Reset()      { *m = InstanceParam{} }
func (*InstanceParam) ProtoMessage() {}
func (*InstanceParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_8608595bf414489a, []int{3}
}
func (m *InstanceParam) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InstanceParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *InstanceParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstanceParam.Merge(m, src)
}
func (m *InstanceParam) XXX_Size() int {
	return m.Size()
}
func (m *InstanceParam) XXX_DiscardUnknown() {
	xxx_messageInfo_InstanceParam.DiscardUnknown(m)
}

var xxx_messageInfo_InstanceParam proto.InternalMessageInfo

func init() {
	proto.RegisterType((*HandleCheckNothingRequest)(nil), "checknothing.HandleCheckNothingRequest")
	proto.RegisterType((*InstanceMsg)(nil), "checknothing.InstanceMsg")
	proto.RegisterType((*Type)(nil), "checknothing.Type")
	proto.RegisterType((*InstanceParam)(nil), "checknothing.InstanceParam")
}

func init() {
	proto.RegisterFile("mixer/template/checknothing/template_handler_service.proto", fileDescriptor_8608595bf414489a)
}

var fileDescriptor_8608595bf414489a = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x31, 0x6f, 0xd4, 0x30,
	0x1c, 0xc5, 0xed, 0xa3, 0x2a, 0xc5, 0xa5, 0x20, 0x59, 0x45, 0xba, 0xbb, 0xc1, 0xaa, 0x22, 0xa4,
	0x76, 0x40, 0xb6, 0x5a, 0xc4, 0x02, 0x13, 0x74, 0xa1, 0x03, 0x08, 0x85, 0xee, 0x27, 0x5f, 0xf2,
	0x6f, 0xce, 0x22, 0xb1, 0x43, 0xec, 0xab, 0x7a, 0x1b, 0x62, 0x64, 0x42, 0xea, 0x97, 0x40, 0x2c,
	0x7c, 0x01, 0x3e, 0x40, 0xc5, 0x74, 0x62, 0xba, 0x05, 0x89, 0xf8, 0x3a, 0x30, 0x76, 0x64, 0x44,
	0x75, 0xd2, 0xea, 0x2a, 0xa0, 0x5b, 0xfe, 0x7e, 0x3f, 0xdb, 0xef, 0xbd, 0x98, 0x3c, 0x2e, 0xd4,
	0x11, 0x54, 0xc2, 0x41, 0x51, 0xe6, 0xd2, 0x81, 0x48, 0x46, 0x90, 0xbc, 0xd1, 0xc6, 0x8d, 0x94,
	0xce, 0x2e, 0x57, 0x07, 0x23, 0xa9, 0xd3, 0x1c, 0xaa, 0x81, 0x85, 0xea, 0x50, 0x25, 0xc0, 0xcb,
	0xca, 0x38, 0x43, 0x6f, 0x2f, 0xc2, 0xfd, 0xf5, 0xcc, 0x64, 0x26, 0x08, 0xe2, 0xfc, 0xab, 0x61,
	0xfa, 0x0f, 0x9a, 0xf3, 0x65, 0x2a, 0x4b, 0x07, 0x95, 0x28, 0x4c, 0x0a, 0xb9, 0x38, 0xdc, 0x1e,
	0x82, 0x93, 0xdb, 0x02, 0x8e, 0x1c, 0x68, 0xab, 0x8c, 0xb6, 0x2d, 0xdd, 0xcb, 0x8c, 0xc9, 0x72,
	0x10, 0x61, 0x1a, 0x8e, 0x0f, 0x84, 0xd4, 0x93, 0x56, 0xda, 0xbc, 0xee, 0xa0, 0x60, 0xa4, 0x01,
	0xa3, 0xcf, 0x98, 0xf4, 0x9e, 0x07, 0xbf, 0xbb, 0xe7, 0xab, 0x2f, 0x1b, 0x7b, 0x31, 0xbc, 0x1d,
	0x83, 0x75, 0xf4, 0x11, 0x59, 0x51, 0xda, 0x3a, 0xa9, 0x13, 0xe8, 0xe2, 0x0d, 0xbc, 0xb5, 0xba,
	0xd3, 0xe3, 0x8b, 0x31, 0xf8, 0x5e, 0xab, 0xbe, 0xb0, 0x59, 0x7c, 0x89, 0xd2, 0x27, 0xe4, 0x4e,
	0x7b, 0xf3, 0x20, 0x31, 0xfa, 0x40, 0x65, 0xdd, 0x4e, 0xd8, 0xbc, 0xce, 0x1b, 0xc7, 0xfc, 0xc2,
	0x31, 0x7f, 0xaa, 0x27, 0xf1, 0x5a, 0xcb, 0xee, 0x06, 0x94, 0xf6, 0xc8, 0x4a, 0x0a, 0xe9, 0xb8,
	0x1c, 0xa8, 0xb4, 0x7b, 0x63, 0x03, 0x6f, 0xdd, 0x8a, 0x6f, 0x86, 0x79, 0x2f, 0x8d, 0xee, 0x93,
	0xd5, 0x85, 0x0b, 0xe9, 0x3d, 0xb2, 0xa4, 0x65, 0x01, 0xdd, 0x2f, 0xdf, 0xbe, 0x46, 0x01, 0x0c,
	0x63, 0xb4, 0x4c, 0x96, 0xf6, 0x27, 0x25, 0x44, 0x77, 0xc9, 0xda, 0x05, 0xfd, 0x4a, 0x56, 0xb2,
	0xd8, 0xf9, 0xf0, 0xcf, 0xac, 0xaf, 0x9b, 0xbf, 0x44, 0x0b, 0x42, 0xff, 0x16, 0xe9, 0xe6, 0xd5,
	0xbc, 0xff, 0xad, 0xaa, 0xcf, 0xb9, 0xb2, 0x4e, 0x19, 0x1e, 0x8a, 0xe7, 0x6d, 0x24, 0x1e, 0x8a,
	0xe7, 0x6d, 0xf1, 0x3c, 0xec, 0x8b, 0xc1, 0x8e, 0x73, 0xf7, 0x6c, 0xff, 0xa4, 0x66, 0x68, 0x5a,
	0x33, 0x34, 0xab, 0x19, 0x3a, 0xab, 0x19, 0x7a, 0xe7, 0x19, 0xfe, 0xe4, 0x19, 0x3a, 0xf1, 0x0c,
	0x4f, 0x3d, 0xc3, 0x33, 0xcf, 0xf0, 0x4f, 0xcf, 0xf0, 0x2f, 0xcf, 0xd0, 0x99, 0x67, 0xf8, 0xe3,
	0x9c, 0xa1, 0xe9, 0x9c, 0xa1, 0xd9, 0x9c, 0xa1, 0xdf, 0xdf, 0x4f, 0x8f, 0x3b, 0xe8, 0xfd, 0x8f,
	0xd3, 0xe3, 0xce, 0x95, 0x67, 0x35, 0x5c, 0x0e, 0xcd, 0x3e, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff,
	0x55, 0xce, 0x5b, 0x88, 0xa9, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HandleCheckNothingServiceClient is the client API for HandleCheckNothingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HandleCheckNothingServiceClient interface {
	// HandleCheckNothing is called by Mixer at request-time to deliver 'checknothing' instances to the backend.
	HandleCheckNothing(ctx context.Context, in *HandleCheckNothingRequest, opts ...grpc.CallOption) (*v1beta1.CheckResult, error)
}

type handleCheckNothingServiceClient struct {
	cc *grpc.ClientConn
}

func NewHandleCheckNothingServiceClient(cc *grpc.ClientConn) HandleCheckNothingServiceClient {
	return &handleCheckNothingServiceClient{cc}
}

func (c *handleCheckNothingServiceClient) HandleCheckNothing(ctx context.Context, in *HandleCheckNothingRequest, opts ...grpc.CallOption) (*v1beta1.CheckResult, error) {
	out := new(v1beta1.CheckResult)
	err := c.cc.Invoke(ctx, "/checknothing.HandleCheckNothingService/HandleCheckNothing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HandleCheckNothingServiceServer is the server API for HandleCheckNothingService service.
type HandleCheckNothingServiceServer interface {
	// HandleCheckNothing is called by Mixer at request-time to deliver 'checknothing' instances to the backend.
	HandleCheckNothing(context.Context, *HandleCheckNothingRequest) (*v1beta1.CheckResult, error)
}

// UnimplementedHandleCheckNothingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHandleCheckNothingServiceServer struct {
}

func (*UnimplementedHandleCheckNothingServiceServer) HandleCheckNothing(ctx context.Context, req *HandleCheckNothingRequest) (*v1beta1.CheckResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleCheckNothing not implemented")
}

func RegisterHandleCheckNothingServiceServer(s *grpc.Server, srv HandleCheckNothingServiceServer) {
	s.RegisterService(&_HandleCheckNothingService_serviceDesc, srv)
}

func _HandleCheckNothingService_HandleCheckNothing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleCheckNothingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandleCheckNothingServiceServer).HandleCheckNothing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/checknothing.HandleCheckNothingService/HandleCheckNothing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandleCheckNothingServiceServer).HandleCheckNothing(ctx, req.(*HandleCheckNothingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HandleCheckNothingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "checknothing.HandleCheckNothingService",
	HandlerType: (*HandleCheckNothingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleCheckNothing",
			Handler:    _HandleCheckNothingService_HandleCheckNothing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mixer/template/checknothing/template_handler_service.proto",
}

func (m *HandleCheckNothingRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HandleCheckNothingRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HandleCheckNothingRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DedupId) > 0 {
		i -= len(m.DedupId)
		copy(dAtA[i:], m.DedupId)
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(len(m.DedupId)))
		i--
		dAtA[i] = 0x1a
	}
	if m.AdapterConfig != nil {
		{
			size, err := m.AdapterConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTemplateHandlerService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Instance != nil {
		{
			size, err := m.Instance.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTemplateHandlerService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *InstanceMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InstanceMsg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InstanceMsg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x2
		i--
		dAtA[i] = 0x93
		i--
		dAtA[i] = 0xe4
		i--
		dAtA[i] = 0xd2
		i--
		dAtA[i] = 0xfa
	}
	return len(dAtA) - i, nil
}

func (m *Type) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Type) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Type) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *InstanceParam) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InstanceParam) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InstanceParam) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTemplateHandlerService(dAtA []byte, offset int, v uint64) int {
	offset -= sovTemplateHandlerService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HandleCheckNothingRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Instance != nil {
		l = m.Instance.Size()
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	if m.AdapterConfig != nil {
		l = m.AdapterConfig.Size()
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	l = len(m.DedupId)
	if l > 0 {
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	return n
}

func (m *InstanceMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 5 + l + sovTemplateHandlerService(uint64(l))
	}
	return n
}

func (m *Type) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *InstanceParam) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTemplateHandlerService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTemplateHandlerService(x uint64) (n int) {
	return sovTemplateHandlerService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *HandleCheckNothingRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HandleCheckNothingRequest{`,
		`Instance:` + strings.Replace(this.Instance.String(), "InstanceMsg", "InstanceMsg", 1) + `,`,
		`AdapterConfig:` + strings.Replace(fmt.Sprintf("%v", this.AdapterConfig), "Any", "types.Any", 1) + `,`,
		`DedupId:` + fmt.Sprintf("%v", this.DedupId) + `,`,
		`}`,
	}, "")
	return s
}
func (this *InstanceMsg) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&InstanceMsg{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Type) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Type{`,
		`}`,
	}, "")
	return s
}
func (this *InstanceParam) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&InstanceParam{`,
		`}`,
	}, "")
	return s
}
func valueToStringTemplateHandlerService(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *HandleCheckNothingRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTemplateHandlerService
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
			return fmt.Errorf("proto: HandleCheckNothingRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HandleCheckNothingRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Instance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
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
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Instance == nil {
				m.Instance = &InstanceMsg{}
			}
			if err := m.Instance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdapterConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
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
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AdapterConfig == nil {
				m.AdapterConfig = &types.Any{}
			}
			if err := m.AdapterConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DedupId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
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
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DedupId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTemplateHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTemplateHandlerService
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
func (m *InstanceMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTemplateHandlerService
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
			return fmt.Errorf("proto: InstanceMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InstanceMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 72295727:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
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
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTemplateHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTemplateHandlerService
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
func (m *Type) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTemplateHandlerService
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
			return fmt.Errorf("proto: Type: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Type: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTemplateHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTemplateHandlerService
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
func (m *InstanceParam) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTemplateHandlerService
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
			return fmt.Errorf("proto: InstanceParam: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InstanceParam: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTemplateHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTemplateHandlerService
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
func skipTemplateHandlerService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTemplateHandlerService
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
					return 0, ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTemplateHandlerService
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
				return 0, ErrInvalidLengthTemplateHandlerService
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTemplateHandlerService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTemplateHandlerService
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTemplateHandlerService(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTemplateHandlerService
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTemplateHandlerService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTemplateHandlerService   = fmt.Errorf("proto: integer overflow")
)
