// Code generated by protoc-gen-go.
// source: userdb.proto
// DO NOT EDIT!

/*
Package userdb is a generated protocol buffer package.

It is generated from these files:
	userdb.proto

It has these top-level messages:
	ProtoDB
	Password
	Scrypt
	Plain
*/
package userdb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProtoDB struct {
	Users map[string]*Password `protobuf:"bytes,1,rep,name=users" json:"users,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ProtoDB) Reset()                    { *m = ProtoDB{} }
func (m *ProtoDB) String() string            { return proto.CompactTextString(m) }
func (*ProtoDB) ProtoMessage()               {}
func (*ProtoDB) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProtoDB) GetUsers() map[string]*Password {
	if m != nil {
		return m.Users
	}
	return nil
}

type Password struct {
	// Types that are valid to be assigned to Scheme:
	//	*Password_Scrypt
	//	*Password_Plain
	Scheme isPassword_Scheme `protobuf_oneof:"scheme"`
}

func (m *Password) Reset()                    { *m = Password{} }
func (m *Password) String() string            { return proto.CompactTextString(m) }
func (*Password) ProtoMessage()               {}
func (*Password) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isPassword_Scheme interface {
	isPassword_Scheme()
}

type Password_Scrypt struct {
	Scrypt *Scrypt `protobuf:"bytes,2,opt,name=scrypt,oneof"`
}
type Password_Plain struct {
	Plain *Plain `protobuf:"bytes,3,opt,name=plain,oneof"`
}

func (*Password_Scrypt) isPassword_Scheme() {}
func (*Password_Plain) isPassword_Scheme()  {}

func (m *Password) GetScheme() isPassword_Scheme {
	if m != nil {
		return m.Scheme
	}
	return nil
}

func (m *Password) GetScrypt() *Scrypt {
	if x, ok := m.GetScheme().(*Password_Scrypt); ok {
		return x.Scrypt
	}
	return nil
}

func (m *Password) GetPlain() *Plain {
	if x, ok := m.GetScheme().(*Password_Plain); ok {
		return x.Plain
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Password) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Password_OneofMarshaler, _Password_OneofUnmarshaler, _Password_OneofSizer, []interface{}{
		(*Password_Scrypt)(nil),
		(*Password_Plain)(nil),
	}
}

func _Password_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Password)
	// scheme
	switch x := m.Scheme.(type) {
	case *Password_Scrypt:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Scrypt); err != nil {
			return err
		}
	case *Password_Plain:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Plain); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Password.Scheme has unexpected type %T", x)
	}
	return nil
}

func _Password_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Password)
	switch tag {
	case 2: // scheme.scrypt
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Scrypt)
		err := b.DecodeMessage(msg)
		m.Scheme = &Password_Scrypt{msg}
		return true, err
	case 3: // scheme.plain
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Plain)
		err := b.DecodeMessage(msg)
		m.Scheme = &Password_Plain{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Password_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Password)
	// scheme
	switch x := m.Scheme.(type) {
	case *Password_Scrypt:
		s := proto.Size(x.Scrypt)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Password_Plain:
		s := proto.Size(x.Plain)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Scrypt struct {
	LogN      uint64 `protobuf:"varint,1,opt,name=logN" json:"logN,omitempty"`
	R         int32  `protobuf:"varint,2,opt,name=r" json:"r,omitempty"`
	P         int32  `protobuf:"varint,3,opt,name=p" json:"p,omitempty"`
	KeyLen    int32  `protobuf:"varint,4,opt,name=keyLen" json:"keyLen,omitempty"`
	Salt      []byte `protobuf:"bytes,5,opt,name=salt,proto3" json:"salt,omitempty"`
	Encrypted []byte `protobuf:"bytes,6,opt,name=encrypted,proto3" json:"encrypted,omitempty"`
}

func (m *Scrypt) Reset()                    { *m = Scrypt{} }
func (m *Scrypt) String() string            { return proto.CompactTextString(m) }
func (*Scrypt) ProtoMessage()               {}
func (*Scrypt) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Plain struct {
	Password []byte `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
}

func (m *Plain) Reset()                    { *m = Plain{} }
func (m *Plain) String() string            { return proto.CompactTextString(m) }
func (*Plain) ProtoMessage()               {}
func (*Plain) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*ProtoDB)(nil), "userdb.ProtoDB")
	proto.RegisterType((*Password)(nil), "userdb.Password")
	proto.RegisterType((*Scrypt)(nil), "userdb.Scrypt")
	proto.RegisterType((*Plain)(nil), "userdb.Plain")
}

func init() { proto.RegisterFile("userdb.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x44, 0x91, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x4d, 0xdb, 0xc4, 0x76, 0xba, 0x4a, 0x99, 0x83, 0x84, 0xe2, 0x41, 0x56, 0x94, 0x9e,
	0x16, 0xa9, 0x17, 0xf1, 0x58, 0x14, 0x44, 0x44, 0x24, 0xe2, 0x0f, 0xd8, 0xba, 0x41, 0xc5, 0x75,
	0x37, 0x24, 0x5b, 0x65, 0xaf, 0x5e, 0xfc, 0xdb, 0x26, 0xb3, 0xe9, 0xf6, 0x36, 0xef, 0x7b, 0x33,
	0x6f, 0x26, 0x04, 0x92, 0x8d, 0xd3, 0xb6, 0x58, 0x67, 0xc6, 0xd6, 0x4d, 0x8d, 0xa2, 0x53, 0xe9,
	0x1f, 0x83, 0xfd, 0xa7, 0x40, 0x6e, 0x56, 0x78, 0x01, 0x3c, 0x50, 0x27, 0xd9, 0xc9, 0x70, 0x31,
	0x5d, 0xce, 0xb3, 0x38, 0x11, 0xfd, 0xec, 0x25, 0x98, 0xb7, 0x55, 0x63, 0x5b, 0xd5, 0x35, 0xce,
	0xef, 0x01, 0x76, 0x10, 0x67, 0x30, 0xfc, 0xd4, 0xad, 0x9f, 0x66, 0x8b, 0x89, 0x0a, 0x25, 0x9e,
	0x03, 0xff, 0xce, 0xcb, 0x8d, 0x96, 0x03, 0xcf, 0xa6, 0xcb, 0x59, 0x9f, 0x98, 0x3b, 0xf7, 0x53,
	0xdb, 0x42, 0x75, 0xf6, 0xf5, 0xe0, 0x8a, 0xa5, 0x1a, 0xc6, 0x5b, 0x8c, 0x0b, 0x10, 0xee, 0xd5,
	0xb6, 0xa6, 0x89, 0x83, 0x87, 0xdb, 0xc1, 0x67, 0xa2, 0x77, 0x7b, 0x2a, 0xfa, 0x78, 0x06, 0xdc,
	0x94, 0xf9, 0x47, 0x25, 0x87, 0xd4, 0x78, 0xd0, 0x6f, 0x08, 0xd0, 0xf7, 0x75, 0xee, 0x6a, 0x1c,
	0x02, 0xdf, 0xf5, 0x97, 0x4e, 0x7f, 0x19, 0x88, 0x2e, 0x05, 0x11, 0x46, 0x65, 0xfd, 0xf6, 0x48,
	0x07, 0x8f, 0x14, 0xd5, 0x98, 0x00, 0xb3, 0xb4, 0x94, 0x2b, 0x66, 0x83, 0x32, 0x94, 0xec, 0x95,
	0xc1, 0x23, 0x10, 0xfe, 0x51, 0x0f, 0xba, 0x92, 0x23, 0x42, 0x51, 0x85, 0x1c, 0x97, 0x97, 0x8d,
	0xe4, 0x9e, 0x26, 0x8a, 0x6a, 0x3c, 0x86, 0x89, 0xae, 0x68, 0x8d, 0x2e, 0xa4, 0x20, 0x63, 0x07,
	0xd2, 0x53, 0xe0, 0x74, 0x20, 0xce, 0x61, 0x6c, 0xe2, 0xa3, 0xe9, 0x8c, 0x44, 0xf5, 0x7a, 0x2d,
	0xe8, 0xa7, 0x2e, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x93, 0xae, 0x19, 0xb9, 0x01, 0x00,
	0x00,
}