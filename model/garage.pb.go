// Code generated by protoc-gen-go. DO NOT EDIT.
// source: garage.proto

package model

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

type GarageCoordinate struct {
	Latitude             float32  `protobuf:"fixed32,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longtitude           float32  `protobuf:"fixed32,2,opt,name=longtitude,proto3" json:"longtitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GarageCoordinate) Reset()         { *m = GarageCoordinate{} }
func (m *GarageCoordinate) String() string { return proto.CompactTextString(m) }
func (*GarageCoordinate) ProtoMessage()    {}
func (*GarageCoordinate) Descriptor() ([]byte, []int) {
	return fileDescriptor_d35248283e3db3c1, []int{0}
}

func (m *GarageCoordinate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GarageCoordinate.Unmarshal(m, b)
}
func (m *GarageCoordinate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GarageCoordinate.Marshal(b, m, deterministic)
}
func (m *GarageCoordinate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GarageCoordinate.Merge(m, src)
}
func (m *GarageCoordinate) XXX_Size() int {
	return xxx_messageInfo_GarageCoordinate.Size(m)
}
func (m *GarageCoordinate) XXX_DiscardUnknown() {
	xxx_messageInfo_GarageCoordinate.DiscardUnknown(m)
}

var xxx_messageInfo_GarageCoordinate proto.InternalMessageInfo

func (m *GarageCoordinate) GetLatitude() float32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *GarageCoordinate) GetLongtitude() float32 {
	if m != nil {
		return m.Longtitude
	}
	return 0
}

type Garage struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Coordinate           *GarageCoordinate `protobuf:"bytes,3,opt,name=coordinate,proto3" json:"coordinate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Garage) Reset()         { *m = Garage{} }
func (m *Garage) String() string { return proto.CompactTextString(m) }
func (*Garage) ProtoMessage()    {}
func (*Garage) Descriptor() ([]byte, []int) {
	return fileDescriptor_d35248283e3db3c1, []int{1}
}

func (m *Garage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Garage.Unmarshal(m, b)
}
func (m *Garage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Garage.Marshal(b, m, deterministic)
}
func (m *Garage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Garage.Merge(m, src)
}
func (m *Garage) XXX_Size() int {
	return xxx_messageInfo_Garage.Size(m)
}
func (m *Garage) XXX_DiscardUnknown() {
	xxx_messageInfo_Garage.DiscardUnknown(m)
}

var xxx_messageInfo_Garage proto.InternalMessageInfo

func (m *Garage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Garage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Garage) GetCoordinate() *GarageCoordinate {
	if m != nil {
		return m.Coordinate
	}
	return nil
}

type GarageList struct {
	List                 []*Garage `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GarageList) Reset()         { *m = GarageList{} }
func (m *GarageList) String() string { return proto.CompactTextString(m) }
func (*GarageList) ProtoMessage()    {}
func (*GarageList) Descriptor() ([]byte, []int) {
	return fileDescriptor_d35248283e3db3c1, []int{2}
}

func (m *GarageList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GarageList.Unmarshal(m, b)
}
func (m *GarageList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GarageList.Marshal(b, m, deterministic)
}
func (m *GarageList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GarageList.Merge(m, src)
}
func (m *GarageList) XXX_Size() int {
	return xxx_messageInfo_GarageList.Size(m)
}
func (m *GarageList) XXX_DiscardUnknown() {
	xxx_messageInfo_GarageList.DiscardUnknown(m)
}

var xxx_messageInfo_GarageList proto.InternalMessageInfo

func (m *GarageList) GetList() []*Garage {
	if m != nil {
		return m.List
	}
	return nil
}

type GarageListByUser struct {
	List                 map[string]*GarageList `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GarageListByUser) Reset()         { *m = GarageListByUser{} }
func (m *GarageListByUser) String() string { return proto.CompactTextString(m) }
func (*GarageListByUser) ProtoMessage()    {}
func (*GarageListByUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_d35248283e3db3c1, []int{3}
}

func (m *GarageListByUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GarageListByUser.Unmarshal(m, b)
}
func (m *GarageListByUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GarageListByUser.Marshal(b, m, deterministic)
}
func (m *GarageListByUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GarageListByUser.Merge(m, src)
}
func (m *GarageListByUser) XXX_Size() int {
	return xxx_messageInfo_GarageListByUser.Size(m)
}
func (m *GarageListByUser) XXX_DiscardUnknown() {
	xxx_messageInfo_GarageListByUser.DiscardUnknown(m)
}

var xxx_messageInfo_GarageListByUser proto.InternalMessageInfo

func (m *GarageListByUser) GetList() map[string]*GarageList {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*GarageCoordinate)(nil), "model.GarageCoordinate")
	proto.RegisterType((*Garage)(nil), "model.Garage")
	proto.RegisterType((*GarageList)(nil), "model.GarageList")
	proto.RegisterType((*GarageListByUser)(nil), "model.GarageListByUser")
	proto.RegisterMapType((map[string]*GarageList)(nil), "model.GarageListByUser.ListEntry")
}

func init() {
	proto.RegisterFile("garage.proto", fileDescriptor_d35248283e3db3c1)
}

var fileDescriptor_d35248283e3db3c1 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0xc1, 0x4a, 0xc4, 0x30,
	0x14, 0x24, 0xe9, 0xee, 0x62, 0x5f, 0x55, 0xd6, 0x77, 0xb1, 0xec, 0x41, 0xba, 0xbd, 0xd8, 0x53,
	0x85, 0x8a, 0x28, 0x1e, 0x15, 0x11, 0x44, 0x3c, 0x04, 0xfc, 0x80, 0x68, 0x42, 0x09, 0x66, 0x1b,
	0x49, 0xb3, 0x42, 0x3f, 0xc4, 0xff, 0x95, 0x26, 0x6b, 0x6d, 0xdd, 0xdb, 0x64, 0xe6, 0xcd, 0xe4,
	0x0d, 0x0f, 0x0e, 0x6b, 0x6e, 0x79, 0x2d, 0xcb, 0x4f, 0x6b, 0x9c, 0xc1, 0xf9, 0xc6, 0x08, 0xa9,
	0xf3, 0x17, 0x58, 0x3e, 0x7a, 0xfa, 0xde, 0x18, 0x2b, 0x54, 0xc3, 0x9d, 0xc4, 0x15, 0x1c, 0x68,
	0xee, 0x94, 0xdb, 0x0a, 0x99, 0x92, 0x8c, 0x14, 0x94, 0x0d, 0x6f, 0x3c, 0x03, 0xd0, 0xa6, 0xa9,
	0x77, 0x2a, 0xf5, 0xea, 0x88, 0xc9, 0x25, 0x2c, 0x42, 0x1e, 0x1e, 0x03, 0x55, 0xc2, 0xfb, 0x63,
	0x46, 0x95, 0x40, 0x84, 0x59, 0xc3, 0x37, 0xc1, 0x13, 0x33, 0x8f, 0xf1, 0x1a, 0xe0, 0x7d, 0xf8,
	0x37, 0x8d, 0x32, 0x52, 0x24, 0xd5, 0x69, 0xe9, 0x37, 0x2b, 0xff, 0xaf, 0xc5, 0x46, 0xa3, 0xf9,
	0x05, 0x40, 0xd0, 0x9f, 0x55, 0xeb, 0x70, 0x0d, 0x33, 0xad, 0x5a, 0x97, 0x92, 0x2c, 0x2a, 0x92,
	0xea, 0x68, 0x12, 0xc0, 0xbc, 0x94, 0x7f, 0x93, 0xdf, 0xa2, 0xbd, 0xe3, 0xae, 0x7b, 0x6d, 0xa5,
	0xc5, 0xab, 0x89, 0x6f, 0x3d, 0xf1, 0xfd, 0x8d, 0x95, 0x3d, 0x7c, 0x68, 0x9c, 0xed, 0x42, 0xd6,
	0xea, 0x09, 0xe2, 0x81, 0xc2, 0x25, 0x44, 0x1f, 0xb2, 0xdb, 0xf5, 0xec, 0x21, 0x9e, 0xc3, 0xfc,
	0x8b, 0xeb, 0x6d, 0x68, 0x9a, 0x54, 0x27, 0x7b, 0xb1, 0x2c, 0xe8, 0xb7, 0xf4, 0x86, 0xbc, 0x2d,
	0xfc, 0x35, 0x2e, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x24, 0xa2, 0xe0, 0x85, 0x9d, 0x01, 0x00,
	0x00,
}