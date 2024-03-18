// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services/search/proto/search.proto

/*
Package search is a generated protocol buffer package.

It is generated from these files:
	services/search/proto/search.proto

It has these top-level messages:
	NearbyRequest
	SearchResult
*/
package search

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NearbyRequest struct {
	Lat     float32 `protobuf:"fixed32,1,opt,name=lat" json:"lat,omitempty"`
	Lon     float32 `protobuf:"fixed32,2,opt,name=lon" json:"lon,omitempty"`
	InDate  string  `protobuf:"bytes,3,opt,name=inDate" json:"inDate,omitempty"`
	OutDate string  `protobuf:"bytes,4,opt,name=outDate" json:"outDate,omitempty"`
}

func (m *NearbyRequest) Reset()                    { *m = NearbyRequest{} }
func (m *NearbyRequest) String() string            { return proto.CompactTextString(m) }
func (*NearbyRequest) ProtoMessage()               {}
func (*NearbyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NearbyRequest) GetLat() float32 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *NearbyRequest) GetLon() float32 {
	if m != nil {
		return m.Lon
	}
	return 0
}

func (m *NearbyRequest) GetInDate() string {
	if m != nil {
		return m.InDate
	}
	return ""
}

func (m *NearbyRequest) GetOutDate() string {
	if m != nil {
		return m.OutDate
	}
	return ""
}

type SearchResult struct {
	HotelIds []string `protobuf:"bytes,1,rep,name=hotelIds" json:"hotelIds,omitempty"`
}

func (m *SearchResult) Reset()                    { *m = SearchResult{} }
func (m *SearchResult) String() string            { return proto.CompactTextString(m) }
func (*SearchResult) ProtoMessage()               {}
func (*SearchResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SearchResult) GetHotelIds() []string {
	if m != nil {
		return m.HotelIds
	}
	return nil
}

func init() {
	proto.RegisterType((*NearbyRequest)(nil), "search.NearbyRequest")
	proto.RegisterType((*SearchResult)(nil), "search.SearchResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Search service

type SearchClient interface {
	Nearby(ctx context.Context, in *NearbyRequest, opts ...grpc.CallOption) (*SearchResult, error)
}

type searchClient struct {
	cc *grpc.ClientConn
}

func NewSearchClient(cc *grpc.ClientConn) SearchClient {
	return &searchClient{cc}
}

func (c *searchClient) Nearby(ctx context.Context, in *NearbyRequest, opts ...grpc.CallOption) (*SearchResult, error) {
	out := new(SearchResult)
	err := grpc.Invoke(ctx, "/search.Search/Nearby", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Search service

type SearchServer interface {
	Nearby(context.Context, *NearbyRequest) (*SearchResult, error)
}

func RegisterSearchServer(s *grpc.Server, srv SearchServer) {
	s.RegisterService(&_Search_serviceDesc, srv)
}

func _Search_Nearby_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NearbyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).Nearby(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.Search/Nearby",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).Nearby(ctx, req.(*NearbyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Search_serviceDesc = grpc.ServiceDesc{
	ServiceName: "search.Search",
	HandlerType: (*SearchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Nearby",
			Handler:    _Search_Nearby_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/search/proto/search.proto",
}

func init() { proto.RegisterFile("services/search/proto/search.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xb1, 0x0a, 0xc2, 0x30,
	0x10, 0x86, 0x49, 0x2b, 0xd1, 0x1e, 0x0a, 0x12, 0x54, 0x42, 0xa7, 0xd2, 0xa9, 0x38, 0xb4, 0xa0,
	0x38, 0xbb, 0xb8, 0xb8, 0x38, 0xc4, 0x27, 0x48, 0xeb, 0x41, 0x0b, 0xa5, 0xd1, 0x24, 0x15, 0x7c,
	0x7b, 0x31, 0x4d, 0x45, 0xb7, 0x7c, 0xff, 0x1f, 0xee, 0xbe, 0x83, 0xd4, 0xa0, 0x7e, 0x36, 0x15,
	0x9a, 0xc2, 0xa0, 0xd4, 0x55, 0x5d, 0xdc, 0xb5, 0xb2, 0xca, 0x43, 0xee, 0x80, 0xd1, 0x81, 0x52,
	0x84, 0xc5, 0x05, 0xa5, 0x2e, 0x5f, 0x02, 0x1f, 0x3d, 0x1a, 0xcb, 0x96, 0x10, 0xb6, 0xd2, 0x72,
	0x92, 0x90, 0x2c, 0x10, 0x9f, 0xa7, 0x4b, 0x54, 0xc7, 0x03, 0x9f, 0xa8, 0x8e, 0x6d, 0x80, 0x36,
	0xdd, 0x49, 0x5a, 0xe4, 0x61, 0x42, 0xb2, 0x48, 0x78, 0x62, 0x1c, 0xa6, 0xaa, 0xb7, 0xae, 0x98,
	0xb8, 0x62, 0xc4, 0x74, 0x0b, 0xf3, 0xab, 0x5b, 0x28, 0xd0, 0xf4, 0xad, 0x65, 0x31, 0xcc, 0x6a,
	0x65, 0xb1, 0x3d, 0xdf, 0x0c, 0x27, 0x49, 0x98, 0x45, 0xe2, 0xcb, 0xbb, 0x23, 0xd0, 0xe1, 0x2f,
	0x3b, 0x00, 0x1d, 0xe4, 0xd8, 0x3a, 0xf7, 0xf6, 0x7f, 0xb2, 0xf1, 0x6a, 0x8c, 0x7f, 0x87, 0x97,
	0xd4, 0x9d, 0xb8, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x55, 0xd9, 0x68, 0x1c, 0x08, 0x01, 0x00,
	0x00,
}
