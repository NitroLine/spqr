// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/router.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type RouterStatus int32

const (
	RouterStatus_CLOSED RouterStatus = 0
	RouterStatus_OPENED RouterStatus = 1
)

var RouterStatus_name = map[int32]string{
	0: "CLOSED",
	1: "OPENED",
}

var RouterStatus_value = map[string]int32{
	"CLOSED": 0,
	"OPENED": 1,
}

func (x RouterStatus) String() string {
	return proto.EnumName(RouterStatus_name, int32(x))
}

func (RouterStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e481873083316576, []int{0}
}

type Router struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address              string       `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Status               RouterStatus `protobuf:"varint,3,opt,name=status,proto3,enum=spqr.RouterStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Router) Reset()         { *m = Router{} }
func (m *Router) String() string { return proto.CompactTextString(m) }
func (*Router) ProtoMessage()    {}
func (*Router) Descriptor() ([]byte, []int) {
	return fileDescriptor_e481873083316576, []int{0}
}

func (m *Router) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Router.Unmarshal(m, b)
}
func (m *Router) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Router.Marshal(b, m, deterministic)
}
func (m *Router) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Router.Merge(m, src)
}
func (m *Router) XXX_Size() int {
	return xxx_messageInfo_Router.Size(m)
}
func (m *Router) XXX_DiscardUnknown() {
	xxx_messageInfo_Router.DiscardUnknown(m)
}

var xxx_messageInfo_Router proto.InternalMessageInfo

func (m *Router) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Router) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Router) GetStatus() RouterStatus {
	if m != nil {
		return m.Status
	}
	return RouterStatus_CLOSED
}

type ListRoutersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRoutersRequest) Reset()         { *m = ListRoutersRequest{} }
func (m *ListRoutersRequest) String() string { return proto.CompactTextString(m) }
func (*ListRoutersRequest) ProtoMessage()    {}
func (*ListRoutersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e481873083316576, []int{1}
}

func (m *ListRoutersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRoutersRequest.Unmarshal(m, b)
}
func (m *ListRoutersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRoutersRequest.Marshal(b, m, deterministic)
}
func (m *ListRoutersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRoutersRequest.Merge(m, src)
}
func (m *ListRoutersRequest) XXX_Size() int {
	return xxx_messageInfo_ListRoutersRequest.Size(m)
}
func (m *ListRoutersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRoutersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRoutersRequest proto.InternalMessageInfo

type ListRoutersReply struct {
	Routers              []*Router `protobuf:"bytes,1,rep,name=routers,proto3" json:"routers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListRoutersReply) Reset()         { *m = ListRoutersReply{} }
func (m *ListRoutersReply) String() string { return proto.CompactTextString(m) }
func (*ListRoutersReply) ProtoMessage()    {}
func (*ListRoutersReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e481873083316576, []int{2}
}

func (m *ListRoutersReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRoutersReply.Unmarshal(m, b)
}
func (m *ListRoutersReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRoutersReply.Marshal(b, m, deterministic)
}
func (m *ListRoutersReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRoutersReply.Merge(m, src)
}
func (m *ListRoutersReply) XXX_Size() int {
	return xxx_messageInfo_ListRoutersReply.Size(m)
}
func (m *ListRoutersReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRoutersReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListRoutersReply proto.InternalMessageInfo

func (m *ListRoutersReply) GetRouters() []*Router {
	if m != nil {
		return m.Routers
	}
	return nil
}

type AddRouterRequest struct {
	Router               *Router  `protobuf:"bytes,1,opt,name=router,proto3" json:"router,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRouterRequest) Reset()         { *m = AddRouterRequest{} }
func (m *AddRouterRequest) String() string { return proto.CompactTextString(m) }
func (*AddRouterRequest) ProtoMessage()    {}
func (*AddRouterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e481873083316576, []int{3}
}

func (m *AddRouterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRouterRequest.Unmarshal(m, b)
}
func (m *AddRouterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRouterRequest.Marshal(b, m, deterministic)
}
func (m *AddRouterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRouterRequest.Merge(m, src)
}
func (m *AddRouterRequest) XXX_Size() int {
	return xxx_messageInfo_AddRouterRequest.Size(m)
}
func (m *AddRouterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRouterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRouterRequest proto.InternalMessageInfo

func (m *AddRouterRequest) GetRouter() *Router {
	if m != nil {
		return m.Router
	}
	return nil
}

type AddRouterReply struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRouterReply) Reset()         { *m = AddRouterReply{} }
func (m *AddRouterReply) String() string { return proto.CompactTextString(m) }
func (*AddRouterReply) ProtoMessage()    {}
func (*AddRouterReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e481873083316576, []int{4}
}

func (m *AddRouterReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRouterReply.Unmarshal(m, b)
}
func (m *AddRouterReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRouterReply.Marshal(b, m, deterministic)
}
func (m *AddRouterReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRouterReply.Merge(m, src)
}
func (m *AddRouterReply) XXX_Size() int {
	return xxx_messageInfo_AddRouterReply.Size(m)
}
func (m *AddRouterReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRouterReply.DiscardUnknown(m)
}

var xxx_messageInfo_AddRouterReply proto.InternalMessageInfo

func (m *AddRouterReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterEnum("spqr.RouterStatus", RouterStatus_name, RouterStatus_value)
	proto.RegisterType((*Router)(nil), "spqr.Router")
	proto.RegisterType((*ListRoutersRequest)(nil), "spqr.ListRoutersRequest")
	proto.RegisterType((*ListRoutersReply)(nil), "spqr.ListRoutersReply")
	proto.RegisterType((*AddRouterRequest)(nil), "spqr.AddRouterRequest")
	proto.RegisterType((*AddRouterReply)(nil), "spqr.AddRouterReply")
}

func init() { proto.RegisterFile("protos/router.proto", fileDescriptor_e481873083316576) }

var fileDescriptor_e481873083316576 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x6b, 0xbb, 0x40,
	0x10, 0xc5, 0xb3, 0xe6, 0x8f, 0x21, 0x13, 0xff, 0x22, 0xd3, 0x50, 0x96, 0x9c, 0x64, 0x29, 0x41,
	0x72, 0x30, 0x60, 0x2f, 0xa5, 0x3d, 0xa5, 0x8d, 0xb7, 0xd0, 0x14, 0x73, 0xeb, 0xa1, 0x90, 0x76,
	0xf7, 0x20, 0x04, 0x34, 0xbb, 0x6b, 0x21, 0x9f, 0xa2, 0x5f, 0xb9, 0xb8, 0xbb, 0x29, 0xd6, 0xf4,
	0xa4, 0xf3, 0x7e, 0x33, 0xcf, 0xf7, 0x10, 0xae, 0x6a, 0x59, 0xe9, 0x4a, 0x2d, 0x65, 0xd5, 0x68,
	0x21, 0x53, 0x33, 0xe1, 0x3f, 0x55, 0x1f, 0x25, 0x7b, 0x03, 0xbf, 0x30, 0x2a, 0x86, 0xe0, 0x95,
	0x9c, 0x92, 0x98, 0x24, 0xe3, 0xc2, 0x2b, 0x39, 0x52, 0x18, 0xed, 0x39, 0x97, 0x42, 0x29, 0xea,
	0x19, 0xf1, 0x3c, 0xe2, 0x02, 0x7c, 0xa5, 0xf7, 0xba, 0x51, 0x74, 0x18, 0x93, 0x24, 0xcc, 0x30,
	0x6d, 0xad, 0x52, 0xeb, 0xb3, 0x33, 0xa4, 0x70, 0x1b, 0x6c, 0x0a, 0xb8, 0x29, 0x95, 0xb6, 0x4c,
	0x15, 0xe2, 0xd8, 0x08, 0xa5, 0xd9, 0x3d, 0x44, 0xbf, 0xd4, 0xfa, 0x70, 0xc2, 0x39, 0x8c, 0x6c,
	0x3e, 0x45, 0x49, 0x3c, 0x4c, 0x26, 0x59, 0xd0, 0xb5, 0x2d, 0xce, 0x90, 0xdd, 0x41, 0xb4, 0xe2,
	0xdc, 0xa9, 0xd6, 0x0f, 0x6f, 0xc0, 0xb7, 0xd8, 0xe4, 0xef, 0x9f, 0x3a, 0xc6, 0x62, 0x08, 0x3b,
	0x97, 0xed, 0x37, 0x7b, 0x9d, 0x17, 0x73, 0x08, 0xba, 0x2d, 0x10, 0xc0, 0x7f, 0xda, 0x6c, 0x77,
	0xf9, 0x3a, 0x1a, 0xb4, 0xef, 0xdb, 0x97, 0xfc, 0x39, 0x5f, 0x47, 0x24, 0xfb, 0x22, 0xf0, 0xdf,
	0x2d, 0x0a, 0xf9, 0x59, 0x7e, 0x08, 0x5c, 0xc1, 0xa4, 0xd3, 0x08, 0xa9, 0x0d, 0x70, 0x59, 0x7d,
	0x76, 0xfd, 0x07, 0xa9, 0x0f, 0x27, 0x36, 0xc0, 0x07, 0x18, 0xff, 0xc4, 0x43, 0xb7, 0xd6, 0x6f,
	0x3a, 0x9b, 0x5e, 0xe8, 0xe6, 0xf8, 0x31, 0x78, 0x85, 0x16, 0x2c, 0xcd, 0xbf, 0x7d, 0xf7, 0xcd,
	0xe3, 0xf6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xf9, 0xe2, 0x9e, 0x2f, 0xf9, 0x01, 0x00, 0x00,
}