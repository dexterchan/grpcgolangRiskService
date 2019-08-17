// Code generated by protoc-gen-go. DO NOT EDIT.
// source: risk.proto

package riskservice

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

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN     HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING     HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING HealthCheckResponse_ServingStatus = 2
)

var HealthCheckResponse_ServingStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SERVING",
	2: "NOT_SERVING",
}

var HealthCheckResponse_ServingStatus_value = map[string]int32{
	"UNKNOWN":     0,
	"SERVING":     1,
	"NOT_SERVING": 2,
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return proto.EnumName(HealthCheckResponse_ServingStatus_name, int32(x))
}

func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f3e21a9c43a64eed, []int{1, 0}
}

type RiskResponse_Status int32

const (
	RiskResponse_SUCCESS RiskResponse_Status = 0
	RiskResponse_FAIL    RiskResponse_Status = 1
)

var RiskResponse_Status_name = map[int32]string{
	0: "SUCCESS",
	1: "FAIL",
}

var RiskResponse_Status_value = map[string]int32{
	"SUCCESS": 0,
	"FAIL":    1,
}

func (x RiskResponse_Status) String() string {
	return proto.EnumName(RiskResponse_Status_name, int32(x))
}

func (RiskResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f3e21a9c43a64eed, []int{3, 0}
}

type HealthCheckRequest struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckRequest) Reset()         { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()    {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3e21a9c43a64eed, []int{0}
}

func (m *HealthCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckRequest.Unmarshal(m, b)
}
func (m *HealthCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckRequest.Marshal(b, m, deterministic)
}
func (m *HealthCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckRequest.Merge(m, src)
}
func (m *HealthCheckRequest) XXX_Size() int {
	return xxx_messageInfo_HealthCheckRequest.Size(m)
}
func (m *HealthCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckRequest proto.InternalMessageInfo

func (m *HealthCheckRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type HealthCheckResponse struct {
	Status               HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,proto3,enum=riskservice.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *HealthCheckResponse) Reset()         { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()    {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3e21a9c43a64eed, []int{1}
}

func (m *HealthCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckResponse.Unmarshal(m, b)
}
func (m *HealthCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckResponse.Marshal(b, m, deterministic)
}
func (m *HealthCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckResponse.Merge(m, src)
}
func (m *HealthCheckResponse) XXX_Size() int {
	return xxx_messageInfo_HealthCheckResponse.Size(m)
}
func (m *HealthCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckResponse proto.InternalMessageInfo

func (m *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	if m != nil {
		return m.Status
	}
	return HealthCheckResponse_UNKNOWN
}

type RiskRequest struct {
	SystemDate           string   `protobuf:"bytes,1,opt,name=systemDate,proto3" json:"systemDate,omitempty"`
	TradeId              string   `protobuf:"bytes,2,opt,name=tradeId,proto3" json:"tradeId,omitempty"`
	TradeMessage         string   `protobuf:"bytes,3,opt,name=tradeMessage,proto3" json:"tradeMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RiskRequest) Reset()         { *m = RiskRequest{} }
func (m *RiskRequest) String() string { return proto.CompactTextString(m) }
func (*RiskRequest) ProtoMessage()    {}
func (*RiskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3e21a9c43a64eed, []int{2}
}

func (m *RiskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RiskRequest.Unmarshal(m, b)
}
func (m *RiskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RiskRequest.Marshal(b, m, deterministic)
}
func (m *RiskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RiskRequest.Merge(m, src)
}
func (m *RiskRequest) XXX_Size() int {
	return xxx_messageInfo_RiskRequest.Size(m)
}
func (m *RiskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RiskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RiskRequest proto.InternalMessageInfo

func (m *RiskRequest) GetSystemDate() string {
	if m != nil {
		return m.SystemDate
	}
	return ""
}

func (m *RiskRequest) GetTradeId() string {
	if m != nil {
		return m.TradeId
	}
	return ""
}

func (m *RiskRequest) GetTradeMessage() string {
	if m != nil {
		return m.TradeMessage
	}
	return ""
}

type RiskResponse struct {
	Status               RiskResponse_Status   `protobuf:"varint,1,opt,name=status,proto3,enum=riskservice.RiskResponse_Status" json:"status,omitempty"`
	Time                 string                `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	TradeId              string                `protobuf:"bytes,3,opt,name=tradeId,proto3" json:"tradeId,omitempty"`
	Ccy                  string                `protobuf:"bytes,4,opt,name=ccy,proto3" json:"ccy,omitempty"`
	Npv                  float64               `protobuf:"fixed64,5,opt,name=npv,proto3" json:"npv,omitempty"`
	Risk                 []*RiskResponse_Tenor `protobuf:"bytes,6,rep,name=risk,proto3" json:"risk,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RiskResponse) Reset()         { *m = RiskResponse{} }
func (m *RiskResponse) String() string { return proto.CompactTextString(m) }
func (*RiskResponse) ProtoMessage()    {}
func (*RiskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3e21a9c43a64eed, []int{3}
}

func (m *RiskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RiskResponse.Unmarshal(m, b)
}
func (m *RiskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RiskResponse.Marshal(b, m, deterministic)
}
func (m *RiskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RiskResponse.Merge(m, src)
}
func (m *RiskResponse) XXX_Size() int {
	return xxx_messageInfo_RiskResponse.Size(m)
}
func (m *RiskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RiskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RiskResponse proto.InternalMessageInfo

func (m *RiskResponse) GetStatus() RiskResponse_Status {
	if m != nil {
		return m.Status
	}
	return RiskResponse_SUCCESS
}

func (m *RiskResponse) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *RiskResponse) GetTradeId() string {
	if m != nil {
		return m.TradeId
	}
	return ""
}

func (m *RiskResponse) GetCcy() string {
	if m != nil {
		return m.Ccy
	}
	return ""
}

func (m *RiskResponse) GetNpv() float64 {
	if m != nil {
		return m.Npv
	}
	return 0
}

func (m *RiskResponse) GetRisk() []*RiskResponse_Tenor {
	if m != nil {
		return m.Risk
	}
	return nil
}

type RiskResponse_Tenor struct {
	Label                string   `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Value                float64  `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RiskResponse_Tenor) Reset()         { *m = RiskResponse_Tenor{} }
func (m *RiskResponse_Tenor) String() string { return proto.CompactTextString(m) }
func (*RiskResponse_Tenor) ProtoMessage()    {}
func (*RiskResponse_Tenor) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3e21a9c43a64eed, []int{3, 0}
}

func (m *RiskResponse_Tenor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RiskResponse_Tenor.Unmarshal(m, b)
}
func (m *RiskResponse_Tenor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RiskResponse_Tenor.Marshal(b, m, deterministic)
}
func (m *RiskResponse_Tenor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RiskResponse_Tenor.Merge(m, src)
}
func (m *RiskResponse_Tenor) XXX_Size() int {
	return xxx_messageInfo_RiskResponse_Tenor.Size(m)
}
func (m *RiskResponse_Tenor) XXX_DiscardUnknown() {
	xxx_messageInfo_RiskResponse_Tenor.DiscardUnknown(m)
}

var xxx_messageInfo_RiskResponse_Tenor proto.InternalMessageInfo

func (m *RiskResponse_Tenor) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *RiskResponse_Tenor) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterEnum("riskservice.HealthCheckResponse_ServingStatus", HealthCheckResponse_ServingStatus_name, HealthCheckResponse_ServingStatus_value)
	proto.RegisterEnum("riskservice.RiskResponse_Status", RiskResponse_Status_name, RiskResponse_Status_value)
	proto.RegisterType((*HealthCheckRequest)(nil), "riskservice.HealthCheckRequest")
	proto.RegisterType((*HealthCheckResponse)(nil), "riskservice.HealthCheckResponse")
	proto.RegisterType((*RiskRequest)(nil), "riskservice.RiskRequest")
	proto.RegisterType((*RiskResponse)(nil), "riskservice.RiskResponse")
	proto.RegisterType((*RiskResponse_Tenor)(nil), "riskservice.RiskResponse.Tenor")
}

func init() { proto.RegisterFile("risk.proto", fileDescriptor_f3e21a9c43a64eed) }

var fileDescriptor_f3e21a9c43a64eed = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xed, 0x3a, 0x1f, 0xa5, 0xe3, 0x16, 0xac, 0x81, 0x83, 0xe9, 0x81, 0x58, 0x16, 0x87, 0x9c,
	0xf6, 0x90, 0x5c, 0x10, 0x37, 0x9a, 0xb6, 0x34, 0x7c, 0xb8, 0xd5, 0xba, 0x85, 0x23, 0xda, 0xba,
	0xa3, 0xd4, 0x8a, 0x1b, 0x1b, 0xef, 0x26, 0xa2, 0x3f, 0x85, 0x13, 0x12, 0x12, 0xff, 0x13, 0xed,
	0xae, 0x03, 0x31, 0x6a, 0xc4, 0x6d, 0xde, 0x9b, 0x99, 0x9d, 0x79, 0x33, 0xb3, 0x00, 0x75, 0xae,
	0xe6, 0xbc, 0xaa, 0x4b, 0x5d, 0xa2, 0x6f, 0x6c, 0x45, 0xf5, 0x2a, 0xcf, 0x28, 0xe6, 0x80, 0x67,
	0x24, 0x0b, 0x7d, 0x3b, 0xb9, 0xa5, 0x6c, 0x2e, 0xe8, 0xeb, 0x92, 0x94, 0xc6, 0x10, 0x76, 0x9b,
	0x80, 0x90, 0x45, 0x6c, 0xb8, 0x27, 0xd6, 0x30, 0xfe, 0xce, 0xe0, 0x69, 0x2b, 0x41, 0x55, 0xe5,
	0x42, 0x11, 0x9e, 0x42, 0x5f, 0x69, 0xa9, 0x97, 0xca, 0x26, 0x3c, 0x1e, 0x71, 0xbe, 0x51, 0x85,
	0x3f, 0x90, 0xc1, 0x53, 0xe3, 0x5b, 0xcc, 0x52, 0x9b, 0x25, 0x9a, 0xec, 0xf8, 0x35, 0x1c, 0xb4,
	0x1c, 0xe8, 0xc3, 0xee, 0x55, 0xf2, 0x3e, 0x39, 0xff, 0x9c, 0x04, 0x3b, 0x06, 0xa4, 0x27, 0xe2,
	0xd3, 0x34, 0x79, 0x1b, 0x30, 0x7c, 0x02, 0x7e, 0x72, 0x7e, 0xf9, 0x65, 0x4d, 0x78, 0xf1, 0x1c,
	0x7c, 0x91, 0xab, 0x3f, 0x22, 0x5e, 0x00, 0xa8, 0x7b, 0xa5, 0xe9, 0xee, 0x58, 0xea, 0xb5, 0x8e,
	0x0d, 0xc6, 0x88, 0xd4, 0xb5, 0xbc, 0xa1, 0xe9, 0x4d, 0xe8, 0x39, 0x91, 0x0d, 0xc4, 0x18, 0xf6,
	0xad, 0xf9, 0x91, 0x94, 0x92, 0x33, 0x0a, 0x3b, 0xd6, 0xdd, 0xe2, 0xe2, 0x5f, 0x1e, 0xec, 0xbb,
	0x6a, 0xcd, 0x04, 0x5e, 0xfd, 0x33, 0x81, 0xa8, 0x35, 0x81, 0xcd, 0x50, 0xde, 0xd6, 0x8c, 0x08,
	0x5d, 0x9d, 0xdf, 0x51, 0xd3, 0x85, 0xb5, 0x37, 0x9b, 0xeb, 0xb4, 0x9b, 0x0b, 0xa0, 0x93, 0x65,
	0xf7, 0x61, 0xd7, 0xb2, 0xc6, 0x34, 0xcc, 0xa2, 0x5a, 0x85, 0xbd, 0x88, 0x0d, 0x99, 0x30, 0x26,
	0x8e, 0xa1, 0x6b, 0x8a, 0x87, 0xfd, 0xa8, 0x33, 0xf4, 0x47, 0x83, 0xed, 0x9d, 0x5c, 0xd2, 0xa2,
	0xac, 0x85, 0x0d, 0x3e, 0x1c, 0x43, 0xcf, 0x42, 0x7c, 0x06, 0xbd, 0x42, 0x5e, 0x53, 0xd1, 0xcc,
	0xcc, 0x01, 0xc3, 0xae, 0x64, 0xb1, 0x74, 0x6d, 0x32, 0xe1, 0x40, 0x3c, 0x80, 0xfe, 0xdf, 0x45,
	0xa5, 0x57, 0x93, 0xc9, 0x49, 0x9a, 0x06, 0x3b, 0xf8, 0x08, 0xba, 0xa7, 0x6f, 0xa6, 0x1f, 0x02,
	0x36, 0xfa, 0xc1, 0xdc, 0x56, 0x52, 0x57, 0x1e, 0xdf, 0x41, 0xcf, 0xde, 0x01, 0x0e, 0xb6, 0x5f,
	0x88, 0xdd, 0xdf, 0x61, 0xf4, 0xbf, 0x13, 0xc2, 0x63, 0x38, 0xc8, 0x64, 0x91, 0x2d, 0x0b, 0xa9,
	0xc9, 0xd4, 0xc0, 0xf0, 0x01, 0xa5, 0xee, 0xb1, 0xe7, 0x5b, 0x67, 0x70, 0xf4, 0x12, 0x82, 0xbc,
	0xe4, 0xf4, 0xad, 0xe2, 0xb3, 0xba, 0xca, 0x6c, 0xdc, 0xd1, 0x9e, 0x89, 0xb8, 0x30, 0xdf, 0xe5,
	0x82, 0xfd, 0xf4, 0xbc, 0xb3, 0xc9, 0x75, 0xdf, 0x7e, 0x9e, 0xf1, 0xef, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x8c, 0x52, 0x16, 0x49, 0x4a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RiskServiceClient is the client API for RiskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RiskServiceClient interface {
	// Sends a healthCheck
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	CalculateRisk(ctx context.Context, in *RiskRequest, opts ...grpc.CallOption) (*RiskResponse, error)
}

type riskServiceClient struct {
	cc *grpc.ClientConn
}

func NewRiskServiceClient(cc *grpc.ClientConn) RiskServiceClient {
	return &riskServiceClient{cc}
}

func (c *riskServiceClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/riskservice.RiskService/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *riskServiceClient) CalculateRisk(ctx context.Context, in *RiskRequest, opts ...grpc.CallOption) (*RiskResponse, error) {
	out := new(RiskResponse)
	err := c.cc.Invoke(ctx, "/riskservice.RiskService/calculateRisk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RiskServiceServer is the server API for RiskService service.
type RiskServiceServer interface {
	// Sends a healthCheck
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	CalculateRisk(context.Context, *RiskRequest) (*RiskResponse, error)
}

// UnimplementedRiskServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRiskServiceServer struct {
}

func (*UnimplementedRiskServiceServer) Check(ctx context.Context, req *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (*UnimplementedRiskServiceServer) CalculateRisk(ctx context.Context, req *RiskRequest) (*RiskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateRisk not implemented")
}

func RegisterRiskServiceServer(s *grpc.Server, srv RiskServiceServer) {
	s.RegisterService(&_RiskService_serviceDesc, srv)
}

func _RiskService_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiskServiceServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/riskservice.RiskService/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiskServiceServer).Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RiskService_CalculateRisk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiskServiceServer).CalculateRisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/riskservice.RiskService/CalculateRisk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiskServiceServer).CalculateRisk(ctx, req.(*RiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RiskService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "riskservice.RiskService",
	HandlerType: (*RiskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _RiskService_Check_Handler,
		},
		{
			MethodName: "calculateRisk",
			Handler:    _RiskService_CalculateRisk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "risk.proto",
}