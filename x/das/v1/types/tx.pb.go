// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zgc/das/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

type MsgRequestDAS struct {
	Requester       string `protobuf:"bytes,1,opt,name=requester,proto3" json:"requester,omitempty" Requester`
	StreamID        string `protobuf:"bytes,2,opt,name=stream_id,json=streamId,proto3" json:"stream_id,omitempty"`
	BatchHeaderHash string `protobuf:"bytes,3,opt,name=batch_header_hash,json=batchHeaderHash,proto3" json:"batch_header_hash,omitempty"`
	NumBlobs        uint32 `protobuf:"varint,4,opt,name=num_blobs,json=numBlobs,proto3" json:"num_blobs,omitempty"`
}

func (m *MsgRequestDAS) Reset()         { *m = MsgRequestDAS{} }
func (m *MsgRequestDAS) String() string { return proto.CompactTextString(m) }
func (*MsgRequestDAS) ProtoMessage()    {}
func (*MsgRequestDAS) Descriptor() ([]byte, []int) {
	return fileDescriptor_030259cfeac21931, []int{0}
}
func (m *MsgRequestDAS) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRequestDAS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRequestDAS.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRequestDAS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequestDAS.Merge(m, src)
}
func (m *MsgRequestDAS) XXX_Size() int {
	return m.Size()
}
func (m *MsgRequestDAS) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequestDAS.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequestDAS proto.InternalMessageInfo

type MsgRequestDASResponse struct {
	RequestID uint64 `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (m *MsgRequestDASResponse) Reset()         { *m = MsgRequestDASResponse{} }
func (m *MsgRequestDASResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRequestDASResponse) ProtoMessage()    {}
func (*MsgRequestDASResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_030259cfeac21931, []int{1}
}
func (m *MsgRequestDASResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRequestDASResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRequestDASResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRequestDASResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequestDASResponse.Merge(m, src)
}
func (m *MsgRequestDASResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRequestDASResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequestDASResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequestDASResponse proto.InternalMessageInfo

type MsgReportDASResult struct {
	RequestID uint64 `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Sampler   string `protobuf:"bytes,2,opt,name=sampler,proto3" json:"sampler,omitempty"`
	Results   []bool `protobuf:"varint,3,rep,packed,name=results,proto3" json:"results,omitempty"`
}

func (m *MsgReportDASResult) Reset()         { *m = MsgReportDASResult{} }
func (m *MsgReportDASResult) String() string { return proto.CompactTextString(m) }
func (*MsgReportDASResult) ProtoMessage()    {}
func (*MsgReportDASResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_030259cfeac21931, []int{2}
}
func (m *MsgReportDASResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgReportDASResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgReportDASResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgReportDASResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgReportDASResult.Merge(m, src)
}
func (m *MsgReportDASResult) XXX_Size() int {
	return m.Size()
}
func (m *MsgReportDASResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgReportDASResult.DiscardUnknown(m)
}

var xxx_messageInfo_MsgReportDASResult proto.InternalMessageInfo

type MsgReportDASResultResponse struct {
}

func (m *MsgReportDASResultResponse) Reset()         { *m = MsgReportDASResultResponse{} }
func (m *MsgReportDASResultResponse) String() string { return proto.CompactTextString(m) }
func (*MsgReportDASResultResponse) ProtoMessage()    {}
func (*MsgReportDASResultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_030259cfeac21931, []int{3}
}
func (m *MsgReportDASResultResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgReportDASResultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgReportDASResultResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgReportDASResultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgReportDASResultResponse.Merge(m, src)
}
func (m *MsgReportDASResultResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgReportDASResultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgReportDASResultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgReportDASResultResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgRequestDAS)(nil), "zgc.das.v1.MsgRequestDAS")
	proto.RegisterType((*MsgRequestDASResponse)(nil), "zgc.das.v1.MsgRequestDASResponse")
	proto.RegisterType((*MsgReportDASResult)(nil), "zgc.das.v1.MsgReportDASResult")
	proto.RegisterType((*MsgReportDASResultResponse)(nil), "zgc.das.v1.MsgReportDASResultResponse")
}

func init() { proto.RegisterFile("zgc/das/v1/tx.proto", fileDescriptor_030259cfeac21931) }

var fileDescriptor_030259cfeac21931 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x6e, 0xd3, 0x40,
	0x14, 0xc6, 0x63, 0x52, 0x41, 0xf2, 0x44, 0x54, 0x31, 0x80, 0xe4, 0x18, 0xe4, 0x86, 0x2c, 0x50,
	0xca, 0x1f, 0x4f, 0x0b, 0x27, 0x20, 0x0a, 0xa2, 0x41, 0xea, 0x66, 0xba, 0x82, 0x8d, 0x35, 0xb6,
	0x87, 0x71, 0x24, 0xdb, 0x63, 0xfc, 0xec, 0xa8, 0xed, 0x29, 0x38, 0x08, 0x0b, 0x8e, 0xd1, 0x65,
	0x97, 0xac, 0x2a, 0x70, 0x6e, 0xc0, 0x09, 0x90, 0xc7, 0x76, 0xd2, 0x50, 0x81, 0xc4, 0x2e, 0xdf,
	0xf7, 0x9b, 0xf9, 0xe6, 0x7b, 0xf1, 0x83, 0xfb, 0xe7, 0xd2, 0xa7, 0x01, 0x47, 0xba, 0x3c, 0xa4,
	0xf9, 0xa9, 0x93, 0x66, 0x2a, 0x57, 0x04, 0xce, 0xa5, 0xef, 0x04, 0x1c, 0x9d, 0xe5, 0xa1, 0x35,
	0xf4, 0x15, 0xc6, 0x0a, 0x5d, 0x4d, 0x68, 0x2d, 0xea, 0x63, 0xd6, 0x03, 0xa9, 0xa4, 0xaa, 0xfd,
	0xea, 0x57, 0xe3, 0x0e, 0xa5, 0x52, 0x32, 0x12, 0x54, 0x2b, 0xaf, 0xf8, 0x44, 0x79, 0x72, 0xd6,
	0x20, 0xf3, 0xda, 0x63, 0x52, 0x24, 0x02, 0x17, 0x4d, 0xd4, 0xf8, 0x9b, 0x01, 0x83, 0x63, 0x94,
	0x4c, 0x7c, 0x2e, 0x04, 0xe6, 0xb3, 0x37, 0x27, 0xe4, 0x39, 0xf4, 0xb3, 0x5a, 0x89, 0xcc, 0x34,
	0x46, 0xc6, 0xa4, 0x3f, 0x1d, 0xfc, 0xba, 0xda, 0xeb, 0xb3, 0xd6, 0x64, 0x1b, 0x4e, 0xf6, 0xa1,
	0x8f, 0x79, 0x26, 0x78, 0xec, 0x2e, 0x02, 0xf3, 0x96, 0x3e, 0x7c, 0xb7, 0xbc, 0xda, 0xeb, 0x9d,
	0x68, 0x73, 0x3e, 0x63, 0xbd, 0x1a, 0xcf, 0x03, 0xf2, 0x0c, 0xee, 0x79, 0x3c, 0xf7, 0x43, 0x37,
	0x14, 0x3c, 0x10, 0x99, 0x1b, 0x72, 0x0c, 0xcd, 0x6e, 0x75, 0x85, 0xed, 0x6a, 0x70, 0xa4, 0xfd,
	0x23, 0x8e, 0x21, 0x79, 0x04, 0xfd, 0xa4, 0x88, 0x5d, 0x2f, 0x52, 0x1e, 0x9a, 0x3b, 0x23, 0x63,
	0x32, 0x60, 0xbd, 0xa4, 0x88, 0xa7, 0x95, 0x1e, 0xbf, 0x85, 0x87, 0x5b, 0x8d, 0x99, 0xc0, 0x54,
	0x25, 0x28, 0xc8, 0x0b, 0x80, 0xa6, 0x59, 0xd5, 0xa6, 0xaa, 0xbe, 0x33, 0x1d, 0x94, 0x9b, 0xea,
	0xf3, 0xd9, 0xba, 0xfa, 0x3c, 0x18, 0x2f, 0x81, 0xe8, 0x98, 0x54, 0x65, 0x4d, 0x4a, 0x11, 0xe5,
	0xff, 0x97, 0x41, 0x4c, 0xb8, 0x83, 0x3c, 0x4e, 0x23, 0x91, 0xd5, 0xc3, 0xb3, 0x56, 0x56, 0x24,
	0xd3, 0x89, 0x68, 0x76, 0x47, 0xdd, 0x49, 0x8f, 0xb5, 0x72, 0xfc, 0x18, 0xac, 0x9b, 0xef, 0xb6,
	0x33, 0xbc, 0xfa, 0x6a, 0x40, 0xf7, 0x18, 0x25, 0x79, 0x0f, 0x70, 0xed, 0x9b, 0x0c, 0x9d, 0xcd,
	0x62, 0x38, 0x5b, 0xc3, 0x5b, 0x4f, 0xfe, 0x8a, 0xd6, 0xff, 0xcb, 0x07, 0xd8, 0xfd, 0x73, 0x4c,
	0xfb, 0xc6, 0xad, 0x2d, 0x6e, 0x3d, 0xfd, 0x37, 0x6f, 0xa3, 0xa7, 0xef, 0x2e, 0x7e, 0xda, 0x9d,
	0x8b, 0xd2, 0x36, 0x2e, 0x4b, 0xdb, 0xf8, 0x51, 0xda, 0xc6, 0x97, 0x95, 0xdd, 0xb9, 0x5c, 0xd9,
	0x9d, 0xef, 0x2b, 0xbb, 0xf3, 0x71, 0x5f, 0x2e, 0xf2, 0xb0, 0xf0, 0x1c, 0x5f, 0xc5, 0xf4, 0x40,
	0x46, 0xdc, 0x43, 0x7a, 0x20, 0x5f, 0xfa, 0x21, 0x5f, 0x24, 0xf4, 0x74, 0xbd, 0xfc, 0x67, 0xa9,
	0x40, 0xef, 0xb6, 0x5e, 0xc7, 0xd7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc2, 0xba, 0x08, 0x98,
	0x17, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	RequestDAS(ctx context.Context, in *MsgRequestDAS, opts ...grpc.CallOption) (*MsgRequestDASResponse, error)
	ReportDASResult(ctx context.Context, in *MsgReportDASResult, opts ...grpc.CallOption) (*MsgReportDASResultResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RequestDAS(ctx context.Context, in *MsgRequestDAS, opts ...grpc.CallOption) (*MsgRequestDASResponse, error) {
	out := new(MsgRequestDASResponse)
	err := c.cc.Invoke(ctx, "/zgc.das.v1.Msg/RequestDAS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ReportDASResult(ctx context.Context, in *MsgReportDASResult, opts ...grpc.CallOption) (*MsgReportDASResultResponse, error) {
	out := new(MsgReportDASResultResponse)
	err := c.cc.Invoke(ctx, "/zgc.das.v1.Msg/ReportDASResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	RequestDAS(context.Context, *MsgRequestDAS) (*MsgRequestDASResponse, error)
	ReportDASResult(context.Context, *MsgReportDASResult) (*MsgReportDASResultResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) RequestDAS(ctx context.Context, req *MsgRequestDAS) (*MsgRequestDASResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestDAS not implemented")
}
func (*UnimplementedMsgServer) ReportDASResult(ctx context.Context, req *MsgReportDASResult) (*MsgReportDASResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportDASResult not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_RequestDAS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequestDAS)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RequestDAS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zgc.das.v1.Msg/RequestDAS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RequestDAS(ctx, req.(*MsgRequestDAS))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ReportDASResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgReportDASResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ReportDASResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zgc.das.v1.Msg/ReportDASResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ReportDASResult(ctx, req.(*MsgReportDASResult))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zgc.das.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestDAS",
			Handler:    _Msg_RequestDAS_Handler,
		},
		{
			MethodName: "ReportDASResult",
			Handler:    _Msg_ReportDASResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zgc/das/v1/tx.proto",
}

func (m *MsgRequestDAS) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRequestDAS) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRequestDAS) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NumBlobs != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.NumBlobs))
		i--
		dAtA[i] = 0x20
	}
	if len(m.BatchHeaderHash) > 0 {
		i -= len(m.BatchHeaderHash)
		copy(dAtA[i:], m.BatchHeaderHash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.BatchHeaderHash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.StreamID) > 0 {
		i -= len(m.StreamID)
		copy(dAtA[i:], m.StreamID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.StreamID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Requester) > 0 {
		i -= len(m.Requester)
		copy(dAtA[i:], m.Requester)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Requester)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRequestDASResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRequestDASResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRequestDASResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RequestID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.RequestID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgReportDASResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgReportDASResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgReportDASResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Results) > 0 {
		for iNdEx := len(m.Results) - 1; iNdEx >= 0; iNdEx-- {
			i--
			if m.Results[iNdEx] {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
		}
		i = encodeVarintTx(dAtA, i, uint64(len(m.Results)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Sampler) > 0 {
		i -= len(m.Sampler)
		copy(dAtA[i:], m.Sampler)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sampler)))
		i--
		dAtA[i] = 0x12
	}
	if m.RequestID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.RequestID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgReportDASResultResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgReportDASResultResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgReportDASResultResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgRequestDAS) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Requester)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.StreamID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.BatchHeaderHash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.NumBlobs != 0 {
		n += 1 + sovTx(uint64(m.NumBlobs))
	}
	return n
}

func (m *MsgRequestDASResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RequestID != 0 {
		n += 1 + sovTx(uint64(m.RequestID))
	}
	return n
}

func (m *MsgReportDASResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RequestID != 0 {
		n += 1 + sovTx(uint64(m.RequestID))
	}
	l = len(m.Sampler)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Results) > 0 {
		n += 1 + sovTx(uint64(len(m.Results))) + len(m.Results)*1
	}
	return n
}

func (m *MsgReportDASResultResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRequestDAS) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRequestDAS: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRequestDAS: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Requester", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Requester = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StreamID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StreamID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchHeaderHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BatchHeaderHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumBlobs", wireType)
			}
			m.NumBlobs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumBlobs |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgRequestDASResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRequestDASResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRequestDASResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestID", wireType)
			}
			m.RequestID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgReportDASResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgReportDASResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgReportDASResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestID", wireType)
			}
			m.RequestID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sampler", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sampler = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTx
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Results = append(m.Results, bool(v != 0))
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTx
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTx
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthTx
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				elementCount = packedLen
				if elementCount != 0 && len(m.Results) == 0 {
					m.Results = make([]bool, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTx
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Results = append(m.Results, bool(v != 0))
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Results", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgReportDASResultResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgReportDASResultResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgReportDASResultResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
