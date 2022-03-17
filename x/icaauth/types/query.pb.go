// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: icaauth/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
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

// QueryInterchainAccountAddressRequest defines the request for the InterchainAccountAddress query.
type QueryInterchainAccountAddressRequest struct {
	ConnectionId string `protobuf:"bytes,1,opt,name=connectionId,proto3" json:"connectionId,omitempty"`
	Owner        string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (m *QueryInterchainAccountAddressRequest) Reset()         { *m = QueryInterchainAccountAddressRequest{} }
func (m *QueryInterchainAccountAddressRequest) String() string { return proto.CompactTextString(m) }
func (*QueryInterchainAccountAddressRequest) ProtoMessage()    {}
func (*QueryInterchainAccountAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4ef18af0a32af09, []int{0}
}
func (m *QueryInterchainAccountAddressRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryInterchainAccountAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryInterchainAccountAddressRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryInterchainAccountAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryInterchainAccountAddressRequest.Merge(m, src)
}
func (m *QueryInterchainAccountAddressRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryInterchainAccountAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryInterchainAccountAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryInterchainAccountAddressRequest proto.InternalMessageInfo

func (m *QueryInterchainAccountAddressRequest) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *QueryInterchainAccountAddressRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

// QueryInterchainAccountAddressResponse defines the response for the InterchainAccountAddress query.
type QueryInterchainAccountAddressResponse struct {
	InterchainAccountAddress string `protobuf:"bytes,1,opt,name=interchainAccountAddress,proto3" json:"interchainAccountAddress,omitempty"`
}

func (m *QueryInterchainAccountAddressResponse) Reset()         { *m = QueryInterchainAccountAddressResponse{} }
func (m *QueryInterchainAccountAddressResponse) String() string { return proto.CompactTextString(m) }
func (*QueryInterchainAccountAddressResponse) ProtoMessage()    {}
func (*QueryInterchainAccountAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4ef18af0a32af09, []int{1}
}
func (m *QueryInterchainAccountAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryInterchainAccountAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryInterchainAccountAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryInterchainAccountAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryInterchainAccountAddressResponse.Merge(m, src)
}
func (m *QueryInterchainAccountAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryInterchainAccountAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryInterchainAccountAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryInterchainAccountAddressResponse proto.InternalMessageInfo

func (m *QueryInterchainAccountAddressResponse) GetInterchainAccountAddress() string {
	if m != nil {
		return m.InterchainAccountAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryInterchainAccountAddressRequest)(nil), "chainmain.icaauth.v1.QueryInterchainAccountAddressRequest")
	proto.RegisterType((*QueryInterchainAccountAddressResponse)(nil), "chainmain.icaauth.v1.QueryInterchainAccountAddressResponse")
}

func init() { proto.RegisterFile("icaauth/v1/query.proto", fileDescriptor_b4ef18af0a32af09) }

var fileDescriptor_b4ef18af0a32af09 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x6a, 0xdb, 0x40,
	0x10, 0xc6, 0x2d, 0x83, 0x0b, 0x5d, 0x7a, 0x12, 0xa6, 0x08, 0x53, 0x44, 0x11, 0x2d, 0x94, 0x82,
	0x35, 0xa8, 0x85, 0x1e, 0xdc, 0x93, 0x7b, 0x33, 0x94, 0xd2, 0x9a, 0x9e, 0x72, 0x71, 0x56, 0xab,
	0x45, 0x5a, 0x88, 0x76, 0xe4, 0xdd, 0x95, 0x62, 0x63, 0x7c, 0xc9, 0x13, 0x04, 0xf2, 0x52, 0x39,
	0x1a, 0x72, 0xc9, 0x31, 0xd8, 0xb9, 0xe4, 0x9c, 0x17, 0x08, 0x5a, 0x39, 0xff, 0x20, 0x26, 0x21,
	0x17, 0xa1, 0xdd, 0xf9, 0xcd, 0x7c, 0xfb, 0x7d, 0x0c, 0x79, 0x2f, 0x18, 0xa5, 0xa5, 0xc9, 0xa0,
	0x8a, 0x60, 0x5a, 0x72, 0x35, 0x0f, 0x0b, 0x85, 0x06, 0xdd, 0x2e, 0xcb, 0xa8, 0x90, 0x39, 0x15,
	0x32, 0xdc, 0x12, 0x61, 0x15, 0xf5, 0xba, 0x29, 0xa6, 0x68, 0x01, 0xa8, 0xff, 0x1a, 0xb6, 0xf7,
	0x21, 0x45, 0x4c, 0x0f, 0x38, 0xd0, 0x42, 0x00, 0x95, 0x12, 0x0d, 0x35, 0x02, 0xa5, 0xde, 0x56,
	0xbf, 0x32, 0xd4, 0x39, 0x6a, 0x88, 0xa9, 0xe6, 0x8d, 0x04, 0x54, 0x51, 0xcc, 0x0d, 0x8d, 0xa0,
	0xa0, 0xa9, 0x90, 0x16, 0x6e, 0xd8, 0x60, 0x9f, 0x7c, 0xfa, 0x57, 0x13, 0x23, 0x69, 0xb8, 0xb2,
	0x2f, 0x18, 0x32, 0x86, 0xa5, 0x34, 0xc3, 0x24, 0x51, 0x5c, 0xeb, 0x31, 0x9f, 0x96, 0x5c, 0x1b,
	0x37, 0x20, 0xef, 0x18, 0x4a, 0xc9, 0x59, 0xdd, 0x3b, 0x4a, 0x3c, 0xe7, 0xa3, 0xf3, 0xe5, 0xed,
	0xf8, 0xd1, 0x9d, 0xdb, 0x25, 0x1d, 0x3c, 0x94, 0x5c, 0x79, 0x6d, 0x5b, 0x6c, 0x0e, 0x01, 0x23,
	0x9f, 0x9f, 0x51, 0xd0, 0x05, 0x4a, 0xcd, 0xdd, 0x01, 0xf1, 0xc4, 0x0e, 0x66, 0x2b, 0xb7, 0xb3,
	0xfe, 0xed, 0xda, 0x21, 0x1d, 0xab, 0xe2, 0x5e, 0x39, 0xc4, 0xdb, 0x25, 0xe5, 0x0e, 0xc2, 0xa7,
	0x42, 0x0e, 0x5f, 0x92, 0x40, 0xef, 0xe7, 0xab, 0x7a, 0x1b, 0x6f, 0xc1, 0xff, 0xa3, 0xb3, 0xcb,
	0x93, 0xf6, 0x1f, 0xf7, 0x37, 0x24, 0xbc, 0xa2, 0x3a, 0x13, 0x3a, 0x4b, 0x66, 0x06, 0x72, 0xaa,
	0x34, 0xdc, 0xee, 0xc3, 0xbd, 0xb9, 0x09, 0x6d, 0xa6, 0x4c, 0x68, 0x33, 0x06, 0x16, 0x0f, 0x73,
	0x5e, 0xc2, 0xc2, 0x26, 0xbb, 0xfc, 0xf5, 0xf7, 0x74, 0xed, 0x3b, 0xab, 0xb5, 0xef, 0x5c, 0xac,
	0x7d, 0xe7, 0x78, 0xe3, 0xb7, 0x56, 0x1b, 0xbf, 0x75, 0xbe, 0xf1, 0x5b, 0x7b, 0x3f, 0x52, 0x61,
	0xb2, 0x32, 0x0e, 0x19, 0xe6, 0xc0, 0xd4, 0xbc, 0x30, 0xd8, 0x47, 0x95, 0xf6, 0xed, 0x74, 0xb0,
	0xdf, 0x7e, 0x6d, 0x04, 0x66, 0x77, 0xea, 0x66, 0x5e, 0x70, 0x1d, 0xbf, 0xb1, 0x5b, 0xf1, 0xfd,
	0x26, 0x00, 0x00, 0xff, 0xff, 0x71, 0xfb, 0x26, 0x62, 0xa5, 0x02, 0x00, 0x00,
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
	// InterchainAccountAddress queries the interchain account address for given `connectionId` and `owner`
	InterchainAccountAddress(ctx context.Context, in *QueryInterchainAccountAddressRequest, opts ...grpc.CallOption) (*QueryInterchainAccountAddressResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) InterchainAccountAddress(ctx context.Context, in *QueryInterchainAccountAddressRequest, opts ...grpc.CallOption) (*QueryInterchainAccountAddressResponse, error) {
	out := new(QueryInterchainAccountAddressResponse)
	err := c.cc.Invoke(ctx, "/chainmain.icaauth.v1.Query/InterchainAccountAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// InterchainAccountAddress queries the interchain account address for given `connectionId` and `owner`
	InterchainAccountAddress(context.Context, *QueryInterchainAccountAddressRequest) (*QueryInterchainAccountAddressResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) InterchainAccountAddress(ctx context.Context, req *QueryInterchainAccountAddressRequest) (*QueryInterchainAccountAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InterchainAccountAddress not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_InterchainAccountAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryInterchainAccountAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).InterchainAccountAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chainmain.icaauth.v1.Query/InterchainAccountAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).InterchainAccountAddress(ctx, req.(*QueryInterchainAccountAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chainmain.icaauth.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InterchainAccountAddress",
			Handler:    _Query_InterchainAccountAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "icaauth/v1/query.proto",
}

func (m *QueryInterchainAccountAddressRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryInterchainAccountAddressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryInterchainAccountAddressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryInterchainAccountAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryInterchainAccountAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryInterchainAccountAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.InterchainAccountAddress) > 0 {
		i -= len(m.InterchainAccountAddress)
		copy(dAtA[i:], m.InterchainAccountAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.InterchainAccountAddress)))
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
func (m *QueryInterchainAccountAddressRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryInterchainAccountAddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.InterchainAccountAddress)
	if l > 0 {
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
func (m *QueryInterchainAccountAddressRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryInterchainAccountAddressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryInterchainAccountAddressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
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
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
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
			m.Owner = string(dAtA[iNdEx:postIndex])
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
func (m *QueryInterchainAccountAddressResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryInterchainAccountAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryInterchainAccountAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InterchainAccountAddress", wireType)
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
			m.InterchainAccountAddress = string(dAtA[iNdEx:postIndex])
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
