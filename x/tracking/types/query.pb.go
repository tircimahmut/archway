// DONTCOVER
// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: archway/tracking/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// QueryBlockGasTrackingRequest is the request for Query.BlockGasTracking.
type QueryBlockGasTrackingRequest struct {
}

func (m *QueryBlockGasTrackingRequest) Reset()         { *m = QueryBlockGasTrackingRequest{} }
func (m *QueryBlockGasTrackingRequest) String() string { return proto.CompactTextString(m) }
func (*QueryBlockGasTrackingRequest) ProtoMessage()    {}
func (*QueryBlockGasTrackingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4819b1e6d461893, []int{0}
}
func (m *QueryBlockGasTrackingRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBlockGasTrackingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBlockGasTrackingRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBlockGasTrackingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBlockGasTrackingRequest.Merge(m, src)
}
func (m *QueryBlockGasTrackingRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryBlockGasTrackingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBlockGasTrackingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBlockGasTrackingRequest proto.InternalMessageInfo

// QueryBlockGasTrackingResponse is the response for Query.BlockGasTracking.
type QueryBlockGasTrackingResponse struct {
	Block BlockTracking `protobuf:"bytes,1,opt,name=block,proto3" json:"block"`
}

func (m *QueryBlockGasTrackingResponse) Reset()         { *m = QueryBlockGasTrackingResponse{} }
func (m *QueryBlockGasTrackingResponse) String() string { return proto.CompactTextString(m) }
func (*QueryBlockGasTrackingResponse) ProtoMessage()    {}
func (*QueryBlockGasTrackingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4819b1e6d461893, []int{1}
}
func (m *QueryBlockGasTrackingResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBlockGasTrackingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBlockGasTrackingResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBlockGasTrackingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBlockGasTrackingResponse.Merge(m, src)
}
func (m *QueryBlockGasTrackingResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryBlockGasTrackingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBlockGasTrackingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBlockGasTrackingResponse proto.InternalMessageInfo

func (m *QueryBlockGasTrackingResponse) GetBlock() BlockTracking {
	if m != nil {
		return m.Block
	}
	return BlockTracking{}
}

func init() {
	proto.RegisterType((*QueryBlockGasTrackingRequest)(nil), "archway.tracking.v1beta1.QueryBlockGasTrackingRequest")
	proto.RegisterType((*QueryBlockGasTrackingResponse)(nil), "archway.tracking.v1beta1.QueryBlockGasTrackingResponse")
}

func init() {
	proto.RegisterFile("archway/tracking/v1beta1/query.proto", fileDescriptor_f4819b1e6d461893)
}

var fileDescriptor_f4819b1e6d461893 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0x2c, 0x4a, 0xce,
	0x28, 0x4f, 0xac, 0xd4, 0x2f, 0x29, 0x4a, 0x4c, 0xce, 0xce, 0xcc, 0x4b, 0xd7, 0x2f, 0x33, 0x4c,
	0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x92, 0x80, 0xaa, 0xd2, 0x83, 0xa9, 0xd2, 0x83, 0xaa, 0x92, 0x12, 0x49, 0xcf, 0x4f, 0xcf,
	0x07, 0x2b, 0xd2, 0x07, 0xb1, 0x20, 0xea, 0xa5, 0x64, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5,
	0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0xa1,
	0xb2, 0xea, 0x38, 0xed, 0x84, 0x1b, 0x0f, 0x56, 0xa8, 0x24, 0xc7, 0x25, 0x13, 0x08, 0x72, 0x85,
	0x53, 0x4e, 0x7e, 0x72, 0xb6, 0x7b, 0x62, 0x71, 0x08, 0x54, 0x3a, 0x28, 0xb5, 0xb0, 0x34, 0xb5,
	0xb8, 0x44, 0x29, 0x85, 0x4b, 0x16, 0x87, 0x7c, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x33,
	0x17, 0x6b, 0x12, 0x48, 0x4e, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x5d, 0x0f, 0x97, 0x3f,
	0xf4, 0xc0, 0x46, 0xc0, 0xf4, 0x3b, 0xb1, 0x9c, 0xb8, 0x27, 0xcf, 0x10, 0x04, 0xd1, 0x6b, 0xb4,
	0x8f, 0x91, 0x8b, 0x15, 0x6c, 0x8d, 0xd0, 0x16, 0x46, 0x2e, 0x01, 0x74, 0xbb, 0x84, 0xcc, 0x70,
	0x1b, 0x8a, 0xcf, 0xf1, 0x52, 0xe6, 0x24, 0xeb, 0x83, 0x78, 0x4a, 0x49, 0xbf, 0xe9, 0xf2, 0x93,
	0xc9, 0x4c, 0x9a, 0x42, 0xea, 0xfa, 0x58, 0xc2, 0x51, 0x1f, 0xec, 0xe6, 0xf8, 0xf4, 0xc4, 0xe2,
	0x78, 0x98, 0xa8, 0x93, 0xef, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24,
	0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x19,
	0xa7, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xc2, 0x0c, 0xd3, 0xcd, 0x4b, 0x2d,
	0x29, 0xcf, 0x2f, 0xca, 0x86, 0x1b, 0x5e, 0x81, 0x30, 0xbe, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89,
	0x0d, 0x1c, 0x39, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc3, 0xcc, 0xbc, 0x3a, 0x3b, 0x02,
	0x00, 0x00,
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
	// BlockGasTracking returns block gas tracking for the current block
	BlockGasTracking(ctx context.Context, in *QueryBlockGasTrackingRequest, opts ...grpc.CallOption) (*QueryBlockGasTrackingResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) BlockGasTracking(ctx context.Context, in *QueryBlockGasTrackingRequest, opts ...grpc.CallOption) (*QueryBlockGasTrackingResponse, error) {
	out := new(QueryBlockGasTrackingResponse)
	err := c.cc.Invoke(ctx, "/archway.tracking.v1beta1.Query/BlockGasTracking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// BlockGasTracking returns block gas tracking for the current block
	BlockGasTracking(context.Context, *QueryBlockGasTrackingRequest) (*QueryBlockGasTrackingResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) BlockGasTracking(ctx context.Context, req *QueryBlockGasTrackingRequest) (*QueryBlockGasTrackingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockGasTracking not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_BlockGasTracking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBlockGasTrackingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BlockGasTracking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/archway.tracking.v1beta1.Query/BlockGasTracking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BlockGasTracking(ctx, req.(*QueryBlockGasTrackingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "archway.tracking.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BlockGasTracking",
			Handler:    _Query_BlockGasTracking_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "archway/tracking/v1beta1/query.proto",
}

func (m *QueryBlockGasTrackingRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBlockGasTrackingRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBlockGasTrackingRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryBlockGasTrackingResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBlockGasTrackingResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBlockGasTrackingResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Block.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
func (m *QueryBlockGasTrackingRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryBlockGasTrackingResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Block.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryBlockGasTrackingRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryBlockGasTrackingRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBlockGasTrackingRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *QueryBlockGasTrackingResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryBlockGasTrackingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBlockGasTrackingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
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
			if err := m.Block.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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