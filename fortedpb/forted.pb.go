// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: forted.proto

package fortedpb

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
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
type HelloFortedRequest struct {
	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}
func (m *HelloFortedRequest) Reset()      { *m = HelloFortedRequest{} }
func (*HelloFortedRequest) ProtoMessage() {}
func (*HelloFortedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_16a4d244bf92f218, []int{0}
}
func (m *HelloFortedRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HelloFortedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HelloFortedRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HelloFortedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloFortedRequest.Merge(m, src)
}
func (m *HelloFortedRequest) XXX_Size() int {
	return m.Size()
}
func (m *HelloFortedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloFortedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloFortedRequest proto.InternalMessageInfo

func (m *HelloFortedRequest) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type HelloFortedResponse struct {
	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *HelloFortedResponse) Reset()      { *m = HelloFortedResponse{} }
func (*HelloFortedResponse) ProtoMessage() {}
func (*HelloFortedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_16a4d244bf92f218, []int{1}
}
func (m *HelloFortedResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HelloFortedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HelloFortedResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HelloFortedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloFortedResponse.Merge(m, src)
}
func (m *HelloFortedResponse) XXX_Size() int {
	return m.Size()
}
func (m *HelloFortedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloFortedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloFortedResponse proto.InternalMessageInfo

func (m *HelloFortedResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloFortedRequest)(nil), "fortedpb.HelloFortedRequest")
	proto.RegisterType((*HelloFortedResponse)(nil), "fortedpb.HelloFortedResponse")
}

func init() { proto.RegisterFile("forted.proto", fileDescriptor_16a4d244bf92f218) }

var fileDescriptor_16a4d244bf92f218 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xcb, 0x2f, 0x2a,
	0x49, 0x4d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf0, 0x0a, 0x92, 0xa4, 0xe4,
	0xd3, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0xe2, 0x49, 0xa5, 0x69, 0xfa, 0x25, 0x99, 0xb9,
	0xa9, 0xc5, 0x25, 0x89, 0xb9, 0x05, 0x10, 0xa5, 0x52, 0xba, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49,
	0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xe9, 0xf9, 0xe9, 0xf9, 0x08, 0x95, 0x20, 0x1e, 0x98, 0x03, 0x66,
	0x41, 0x94, 0x2b, 0xa9, 0x71, 0x09, 0x79, 0xa4, 0xe6, 0xe4, 0xe4, 0xbb, 0x81, 0x2d, 0x08, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe0, 0x62, 0xce, 0x2d, 0x4e, 0x97, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x95, 0xd4, 0xb9, 0x84, 0x51, 0xd4, 0x15, 0x17, 0xe4, 0xe7, 0x15,
	0xa7, 0x62, 0x2a, 0x34, 0x0a, 0xe5, 0x62, 0x0d, 0x28, 0xca, 0xaf, 0xa8, 0x14, 0xf2, 0xe1, 0xe2,
	0x46, 0xd2, 0x21, 0x24, 0xa3, 0x07, 0xf3, 0x83, 0x1e, 0xa6, 0x85, 0x52, 0xb2, 0x38, 0x64, 0x21,
	0xd6, 0x28, 0x31, 0x38, 0x59, 0xdd, 0x78, 0x28, 0xc7, 0xf0, 0xe1, 0xa1, 0x1c, 0xe3, 0x8f, 0x87,
	0x72, 0x8c, 0x0d, 0x8f, 0xe4, 0x18, 0x57, 0x3c, 0x92, 0x63, 0xdc, 0xf1, 0x48, 0x8e, 0xf1, 0xc4,
	0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1,
	0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xe0, 0x61, 0x96, 0xc4, 0x06, 0xf6, 0xaa,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xc1, 0x06, 0x26, 0x54, 0x01, 0x00, 0x00,
}

func (this *HelloFortedRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HelloFortedRequest)
	if !ok {
		that2, ok := that.(HelloFortedRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Msg != that1.Msg {
		return false
	}
	return true
}
func (this *HelloFortedResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HelloFortedResponse)
	if !ok {
		that2, ok := that.(HelloFortedResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Msg != that1.Msg {
		return false
	}
	return true
}
func (this *HelloFortedRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&fortedpb.HelloFortedRequest{")
	s = append(s, "Msg: "+fmt.Sprintf("%#v", this.Msg)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *HelloFortedResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&fortedpb.HelloFortedResponse{")
	s = append(s, "Msg: "+fmt.Sprintf("%#v", this.Msg)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringForted(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProxyClient is the client API for Proxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProxyClient interface {
	HelloForted(ctx context.Context, in *HelloFortedRequest, opts ...grpc.CallOption) (*HelloFortedResponse, error)
}

type proxyClient struct {
	cc *grpc.ClientConn
}

func NewProxyClient(cc *grpc.ClientConn) ProxyClient {
	return &proxyClient{cc}
}

func (c *proxyClient) HelloForted(ctx context.Context, in *HelloFortedRequest, opts ...grpc.CallOption) (*HelloFortedResponse, error) {
	out := new(HelloFortedResponse)
	err := c.cc.Invoke(ctx, "/fortedpb.Proxy/HelloForted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProxyServer is the server API for Proxy service.
type ProxyServer interface {
	HelloForted(context.Context, *HelloFortedRequest) (*HelloFortedResponse, error)
}

// UnimplementedProxyServer can be embedded to have forward compatible implementations.
type UnimplementedProxyServer struct {
}

func (*UnimplementedProxyServer) HelloForted(ctx context.Context, req *HelloFortedRequest) (*HelloFortedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloForted not implemented")
}

func RegisterProxyServer(s *grpc.Server, srv ProxyServer) {
	s.RegisterService(&_Proxy_serviceDesc, srv)
}

func _Proxy_HelloForted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloFortedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).HelloForted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fortedpb.Proxy/HelloForted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).HelloForted(ctx, req.(*HelloFortedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Proxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fortedpb.Proxy",
	HandlerType: (*ProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloForted",
			Handler:    _Proxy_HelloForted_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "forted.proto",
}

func (m *HelloFortedRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HelloFortedRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HelloFortedRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Msg) > 0 {
		i -= len(m.Msg)
		copy(dAtA[i:], m.Msg)
		i = encodeVarintForted(dAtA, i, uint64(len(m.Msg)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HelloFortedResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HelloFortedResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HelloFortedResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Msg) > 0 {
		i -= len(m.Msg)
		copy(dAtA[i:], m.Msg)
		i = encodeVarintForted(dAtA, i, uint64(len(m.Msg)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintForted(dAtA []byte, offset int, v uint64) int {
	offset -= sovForted(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedHelloFortedRequest(r randyForted, easy bool) *HelloFortedRequest {
	this := &HelloFortedRequest{}
	this.Msg = string(randStringForted(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedHelloFortedResponse(r randyForted, easy bool) *HelloFortedResponse {
	this := &HelloFortedResponse{}
	this.Msg = string(randStringForted(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyForted interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneForted(r randyForted) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringForted(r randyForted) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneForted(r)
	}
	return string(tmps)
}
func randUnrecognizedForted(r randyForted, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldForted(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldForted(dAtA []byte, r randyForted, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateForted(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateForted(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateForted(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateForted(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateForted(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateForted(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateForted(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *HelloFortedRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovForted(uint64(l))
	}
	return n
}

func (m *HelloFortedResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovForted(uint64(l))
	}
	return n
}

func sovForted(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozForted(x uint64) (n int) {
	return sovForted(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *HelloFortedRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HelloFortedRequest{`,
		`Msg:` + fmt.Sprintf("%v", this.Msg) + `,`,
		`}`,
	}, "")
	return s
}
func (this *HelloFortedResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HelloFortedResponse{`,
		`Msg:` + fmt.Sprintf("%v", this.Msg) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringForted(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *HelloFortedRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowForted
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
			return fmt.Errorf("proto: HelloFortedRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HelloFortedRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowForted
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
				return ErrInvalidLengthForted
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthForted
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipForted(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthForted
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthForted
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
func (m *HelloFortedResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowForted
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
			return fmt.Errorf("proto: HelloFortedResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HelloFortedResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowForted
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
				return ErrInvalidLengthForted
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthForted
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipForted(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthForted
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthForted
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
func skipForted(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowForted
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
					return 0, ErrIntOverflowForted
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
					return 0, ErrIntOverflowForted
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
				return 0, ErrInvalidLengthForted
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupForted
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthForted
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthForted        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowForted          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupForted = fmt.Errorf("proto: unexpected end of group")
)
