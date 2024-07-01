// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: aura/v1/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	blocklist "github.com/ondoprotocol/usdy-noble/x/aura/types/blocklist"
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

type GenesisState struct {
	// blocklist_state is the genesis state of the blocklist submodule.
	BlocklistState blocklist.GenesisState `protobuf:"bytes,1,opt,name=blocklist_state,json=blocklistState,proto3" json:"blocklist_state"`
	// paused is the paused state of this module.
	Paused bool `protobuf:"varint,2,opt,name=paused,proto3" json:"paused,omitempty"`
	// owner is the address that can control this module.
	Owner string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	// pending_owner is the address of the new owner during an ownership transfer.
	PendingOwner string `protobuf:"bytes,4,opt,name=pending_owner,json=pendingOwner,proto3" json:"pending_owner,omitempty"`
	// burners is the list of addresses that can burn USDY.
	Burners []Burner `protobuf:"bytes,5,rep,name=burners,proto3" json:"burners"`
	// minters is the list of addresses that can mint USDY.
	Minters []Minter `protobuf:"bytes,6,rep,name=minters,proto3" json:"minters"`
	// pausers is the list of addresses that can pause USDY.
	Pausers []string `protobuf:"bytes,7,rep,name=pausers,proto3" json:"pausers,omitempty"`
	// blocked_channels is the list of IBC channels where transfers are blocked.
	BlockedChannels []string `protobuf:"bytes,8,rep,name=blocked_channels,json=blockedChannels,proto3" json:"blocked_channels,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_dffbbeb9813c8a98, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetBlocklistState() blocklist.GenesisState {
	if m != nil {
		return m.BlocklistState
	}
	return blocklist.GenesisState{}
}

func (m *GenesisState) GetPaused() bool {
	if m != nil {
		return m.Paused
	}
	return false
}

func (m *GenesisState) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *GenesisState) GetPendingOwner() string {
	if m != nil {
		return m.PendingOwner
	}
	return ""
}

func (m *GenesisState) GetBurners() []Burner {
	if m != nil {
		return m.Burners
	}
	return nil
}

func (m *GenesisState) GetMinters() []Minter {
	if m != nil {
		return m.Minters
	}
	return nil
}

func (m *GenesisState) GetPausers() []string {
	if m != nil {
		return m.Pausers
	}
	return nil
}

func (m *GenesisState) GetBlockedChannels() []string {
	if m != nil {
		return m.BlockedChannels
	}
	return nil
}

type Burner struct {
	Address   string                                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Allowance github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=allowance,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"allowance"`
}

func (m *Burner) Reset()         { *m = Burner{} }
func (m *Burner) String() string { return proto.CompactTextString(m) }
func (*Burner) ProtoMessage()    {}
func (*Burner) Descriptor() ([]byte, []int) {
	return fileDescriptor_dffbbeb9813c8a98, []int{1}
}
func (m *Burner) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Burner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Burner.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Burner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Burner.Merge(m, src)
}
func (m *Burner) XXX_Size() int {
	return m.Size()
}
func (m *Burner) XXX_DiscardUnknown() {
	xxx_messageInfo_Burner.DiscardUnknown(m)
}

var xxx_messageInfo_Burner proto.InternalMessageInfo

func (m *Burner) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type Minter struct {
	Address   string                                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Allowance github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=allowance,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"allowance"`
}

func (m *Minter) Reset()         { *m = Minter{} }
func (m *Minter) String() string { return proto.CompactTextString(m) }
func (*Minter) ProtoMessage()    {}
func (*Minter) Descriptor() ([]byte, []int) {
	return fileDescriptor_dffbbeb9813c8a98, []int{2}
}
func (m *Minter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Minter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Minter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Minter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Minter.Merge(m, src)
}
func (m *Minter) XXX_Size() int {
	return m.Size()
}
func (m *Minter) XXX_DiscardUnknown() {
	xxx_messageInfo_Minter.DiscardUnknown(m)
}

var xxx_messageInfo_Minter proto.InternalMessageInfo

func (m *Minter) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "aura.v1.GenesisState")
	proto.RegisterType((*Burner)(nil), "aura.v1.Burner")
	proto.RegisterType((*Minter)(nil), "aura.v1.Minter")
}

func init() { proto.RegisterFile("aura/v1/genesis.proto", fileDescriptor_dffbbeb9813c8a98) }

var fileDescriptor_dffbbeb9813c8a98 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x52, 0xcb, 0x6a, 0xdc, 0x30,
	0x14, 0xb5, 0x33, 0x89, 0x1d, 0x2b, 0x69, 0xa7, 0x98, 0xb4, 0x98, 0x2c, 0x6c, 0x93, 0x42, 0x71,
	0x17, 0x91, 0x98, 0xf4, 0x0f, 0xdc, 0x45, 0x09, 0xf4, 0x01, 0xee, 0xae, 0x9b, 0x41, 0xb6, 0x84,
	0x63, 0xe2, 0x91, 0x8c, 0x25, 0x4f, 0x9a, 0xbf, 0xe8, 0x67, 0x65, 0x39, 0xcb, 0xa1, 0x8b, 0xa1,
	0xcc, 0xfc, 0x48, 0xd1, 0xc3, 0xd3, 0xa1, 0xdd, 0x77, 0x25, 0xdd, 0x73, 0xcf, 0x39, 0xf7, 0xe8,
	0x22, 0xf0, 0x12, 0x0f, 0x3d, 0x46, 0xcb, 0x19, 0xaa, 0x29, 0xa3, 0xa2, 0x11, 0xb0, 0xeb, 0xb9,
	0xe4, 0xa1, 0xaf, 0x60, 0xb8, 0x9c, 0x5d, 0x26, 0xba, 0x5f, 0xb6, 0xbc, 0xba, 0x6f, 0x1b, 0x21,
	0xff, 0x61, 0x5e, 0x5e, 0xd4, 0xbc, 0xe6, 0xfa, 0x8a, 0xd4, 0xcd, 0xa0, 0x57, 0xeb, 0x23, 0x70,
	0xfe, 0xc1, 0xf0, 0xbe, 0x4a, 0x2c, 0x69, 0xf8, 0x19, 0x4c, 0xf7, 0x26, 0x73, 0xa1, 0xa0, 0xc8,
	0x4d, 0xdd, 0xec, 0xec, 0x26, 0x81, 0x7a, 0xd4, 0xbe, 0x09, 0x97, 0x33, 0x78, 0xa8, 0xcc, 0x8f,
	0x9f, 0x36, 0x89, 0x53, 0x3c, 0xdf, 0x13, 0x8c, 0xdf, 0x2b, 0xe0, 0x75, 0x78, 0x10, 0x94, 0x44,
	0x47, 0xa9, 0x9b, 0x9d, 0x16, 0xb6, 0x0a, 0x2f, 0xc0, 0x09, 0x7f, 0x60, 0xb4, 0x8f, 0x26, 0xa9,
	0x9b, 0x05, 0x85, 0x29, 0xc2, 0xd7, 0xe0, 0x59, 0x47, 0x19, 0x69, 0x58, 0x3d, 0x37, 0xdd, 0x63,
	0xdd, 0x3d, 0xb7, 0xe0, 0x17, 0x4d, 0x42, 0xc0, 0x2f, 0x87, 0x9e, 0xd1, 0x5e, 0x44, 0x27, 0xe9,
	0x24, 0x3b, 0xbb, 0x99, 0x42, 0xbb, 0x05, 0x98, 0x6b, 0xdc, 0x46, 0x19, 0x59, 0x4a, 0xb0, 0x68,
	0x98, 0x54, 0x02, 0xef, 0x2f, 0xc1, 0x27, 0x8d, 0x8f, 0x02, 0xcb, 0x0a, 0x23, 0xe0, 0xeb, 0x98,
	0xbd, 0x88, 0xfc, 0x74, 0x92, 0x05, 0xc5, 0x58, 0x86, 0x6f, 0xc1, 0x0b, 0xfd, 0x40, 0x4a, 0xe6,
	0xd5, 0x1d, 0x66, 0x8c, 0xb6, 0x22, 0x3a, 0xd5, 0x94, 0xa9, 0xc5, 0xdf, 0x5b, 0xf8, 0xaa, 0x03,
	0x9e, 0x89, 0xa3, 0xec, 0x30, 0x21, 0x3d, 0x15, 0x42, 0xef, 0x32, 0x28, 0xc6, 0x32, 0xfc, 0x08,
	0x02, 0xdc, 0xb6, 0xfc, 0x01, 0xb3, 0x8a, 0xea, 0x05, 0x05, 0x39, 0x54, 0x51, 0x7e, 0x6e, 0x92,
	0x37, 0x75, 0x23, 0xef, 0x86, 0x12, 0x56, 0x7c, 0x81, 0x2a, 0x2e, 0x16, 0x5c, 0xd8, 0xe3, 0x5a,
	0x90, 0x7b, 0x24, 0x1f, 0x3b, 0x2a, 0xe0, 0x2d, 0x93, 0xc5, 0x1f, 0x03, 0x35, 0xd1, 0xbc, 0xe7,
	0x7f, 0x4d, 0xcc, 0x6f, 0x9f, 0xb6, 0xb1, 0xbb, 0xda, 0xc6, 0xee, 0xaf, 0x6d, 0xec, 0xfe, 0xd8,
	0xc5, 0xce, 0x6a, 0x17, 0x3b, 0xeb, 0x5d, 0xec, 0x7c, 0x43, 0x07, 0x66, 0x9c, 0x11, 0xf3, 0xf3,
	0x2a, 0xde, 0xa2, 0x41, 0x90, 0xc7, 0x6b, 0xc6, 0xcb, 0x96, 0xa2, 0xef, 0x48, 0x7f, 0x5a, 0xed,
	0x5c, 0x7a, 0x9a, 0xf1, 0xee, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x20, 0x92, 0xe8, 0xb8, 0xe9,
	0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BlockedChannels) > 0 {
		for iNdEx := len(m.BlockedChannels) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.BlockedChannels[iNdEx])
			copy(dAtA[i:], m.BlockedChannels[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.BlockedChannels[iNdEx])))
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.Pausers) > 0 {
		for iNdEx := len(m.Pausers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Pausers[iNdEx])
			copy(dAtA[i:], m.Pausers[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.Pausers[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.Minters) > 0 {
		for iNdEx := len(m.Minters) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Minters[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Burners) > 0 {
		for iNdEx := len(m.Burners) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Burners[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.PendingOwner) > 0 {
		i -= len(m.PendingOwner)
		copy(dAtA[i:], m.PendingOwner)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PendingOwner)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Paused {
		i--
		if m.Paused {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.BlocklistState.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Burner) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Burner) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Burner) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Allowance.Size()
		i -= size
		if _, err := m.Allowance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Minter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Minter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Minter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Allowance.Size()
		i -= size
		if _, err := m.Allowance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.BlocklistState.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.Paused {
		n += 2
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.PendingOwner)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Burners) > 0 {
		for _, e := range m.Burners {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Minters) > 0 {
		for _, e := range m.Minters {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Pausers) > 0 {
		for _, s := range m.Pausers {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BlockedChannels) > 0 {
		for _, s := range m.BlockedChannels {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Burner) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.Allowance.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *Minter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.Allowance.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlocklistState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BlocklistState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Paused", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
			m.Paused = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PendingOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Burners", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Burners = append(m.Burners, Burner{})
			if err := m.Burners[len(m.Burners)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Minters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Minters = append(m.Minters, Minter{})
			if err := m.Minters[len(m.Minters)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pausers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pausers = append(m.Pausers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockedChannels", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockedChannels = append(m.BlockedChannels, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Burner) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Burner: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Burner: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allowance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Allowance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Minter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Minter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Minter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allowance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Allowance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
