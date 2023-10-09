// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/protocolpool/v1/protocolpool.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type MsgSetBudget struct {
	Proposer string  `protobuf:"bytes,1,opt,name=proposer,proto3" json:"proposer,omitempty"`
	Budget   *Budget `protobuf:"bytes,2,opt,name=budget,proto3" json:"budget,omitempty"`
}

func (m *MsgSetBudget) Reset()         { *m = MsgSetBudget{} }
func (m *MsgSetBudget) String() string { return proto.CompactTextString(m) }
func (*MsgSetBudget) ProtoMessage()    {}
func (*MsgSetBudget) Descriptor() ([]byte, []int) {
	return fileDescriptor_b291839d0a414c0a, []int{0}
}
func (m *MsgSetBudget) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetBudget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetBudget.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetBudget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetBudget.Merge(m, src)
}
func (m *MsgSetBudget) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetBudget) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetBudget.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetBudget proto.InternalMessageInfo

func (m *MsgSetBudget) GetProposer() string {
	if m != nil {
		return m.Proposer
	}
	return ""
}

func (m *MsgSetBudget) GetBudget() *Budget {
	if m != nil {
		return m.Budget
	}
	return nil
}

type Budget struct {
	Items []*BudgetItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (m *Budget) Reset()         { *m = Budget{} }
func (m *Budget) String() string { return proto.CompactTextString(m) }
func (*Budget) ProtoMessage()    {}
func (*Budget) Descriptor() ([]byte, []int) {
	return fileDescriptor_b291839d0a414c0a, []int{1}
}
func (m *Budget) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Budget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Budget.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Budget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Budget.Merge(m, src)
}
func (m *Budget) XXX_Size() int {
	return m.Size()
}
func (m *Budget) XXX_DiscardUnknown() {
	xxx_messageInfo_Budget.DiscardUnknown(m)
}

var xxx_messageInfo_Budget proto.InternalMessageInfo

func (m *Budget) GetItems() []*BudgetItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type BudgetItem struct {
	Address            string                      `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Weight             cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=weight,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"weight"`
	ProtocolPoolConfig *ProtocolPoolConfig         `protobuf:"bytes,3,opt,name=protocol_pool_config,json=protocolPoolConfig,proto3" json:"protocol_pool_config,omitempty"`
}

func (m *BudgetItem) Reset()         { *m = BudgetItem{} }
func (m *BudgetItem) String() string { return proto.CompactTextString(m) }
func (*BudgetItem) ProtoMessage()    {}
func (*BudgetItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_b291839d0a414c0a, []int{2}
}
func (m *BudgetItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BudgetItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BudgetItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BudgetItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BudgetItem.Merge(m, src)
}
func (m *BudgetItem) XXX_Size() int {
	return m.Size()
}
func (m *BudgetItem) XXX_DiscardUnknown() {
	xxx_messageInfo_BudgetItem.DiscardUnknown(m)
}

var xxx_messageInfo_BudgetItem proto.InternalMessageInfo

func (m *BudgetItem) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *BudgetItem) GetProtocolPoolConfig() *ProtocolPoolConfig {
	if m != nil {
		return m.ProtocolPoolConfig
	}
	return nil
}

type ProtocolPoolConfig struct {
	AddressPercentage map[string]string `protobuf:"bytes,1,rep,name=address_percentage,json=addressPercentage,proto3" json:"address_percentage,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *ProtocolPoolConfig) Reset()         { *m = ProtocolPoolConfig{} }
func (m *ProtocolPoolConfig) String() string { return proto.CompactTextString(m) }
func (*ProtocolPoolConfig) ProtoMessage()    {}
func (*ProtocolPoolConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_b291839d0a414c0a, []int{3}
}
func (m *ProtocolPoolConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProtocolPoolConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProtocolPoolConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProtocolPoolConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolPoolConfig.Merge(m, src)
}
func (m *ProtocolPoolConfig) XXX_Size() int {
	return m.Size()
}
func (m *ProtocolPoolConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolPoolConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolPoolConfig proto.InternalMessageInfo

func (m *ProtocolPoolConfig) GetAddressPercentage() map[string]string {
	if m != nil {
		return m.AddressPercentage
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgSetBudget)(nil), "cosmos.protocolpool.v1.MsgSetBudget")
	proto.RegisterType((*Budget)(nil), "cosmos.protocolpool.v1.Budget")
	proto.RegisterType((*BudgetItem)(nil), "cosmos.protocolpool.v1.BudgetItem")
	proto.RegisterType((*ProtocolPoolConfig)(nil), "cosmos.protocolpool.v1.ProtocolPoolConfig")
	proto.RegisterMapType((map[string]string)(nil), "cosmos.protocolpool.v1.ProtocolPoolConfig.AddressPercentageEntry")
}

func init() {
	proto.RegisterFile("cosmos/protocolpool/v1/protocolpool.proto", fileDescriptor_b291839d0a414c0a)
}

var fileDescriptor_b291839d0a414c0a = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xce, 0x34, 0x34, 0xda, 0x57, 0x0f, 0x3a, 0x84, 0xb2, 0x46, 0xd8, 0xc6, 0x3d, 0x45, 0xa1,
	0x1b, 0x12, 0x45, 0x8a, 0x9e, 0xba, 0xc6, 0x43, 0x41, 0x21, 0x4c, 0x6f, 0x22, 0x84, 0xed, 0xee,
	0x73, 0xba, 0x74, 0x77, 0x67, 0xd8, 0x99, 0x46, 0x17, 0xfc, 0x11, 0xfe, 0x98, 0xfe, 0x88, 0x1e,
	0x3c, 0x94, 0x9e, 0xc4, 0x43, 0x91, 0xe4, 0xe4, 0xbf, 0x90, 0xec, 0x4c, 0xaa, 0x6b, 0xa3, 0xf6,
	0x36, 0xef, 0x7b, 0xdf, 0xf7, 0xde, 0xf7, 0x3e, 0x18, 0x78, 0x14, 0x09, 0x95, 0x09, 0xd5, 0x97,
	0x85, 0xd0, 0x22, 0x12, 0xa9, 0x14, 0x22, 0xed, 0x4f, 0x07, 0xb5, 0xda, 0xaf, 0x0a, 0xba, 0x65,
	0xa8, 0x7e, 0xad, 0x35, 0x1d, 0x74, 0xda, 0x5c, 0x70, 0x51, 0x81, 0xfd, 0xc5, 0xcb, 0xf4, 0x3b,
	0xf7, 0x0d, 0x7b, 0x62, 0x1a, 0xbf, 0x4b, 0xbd, 0x4f, 0x70, 0xe7, 0x8d, 0xe2, 0x07, 0xa8, 0x83,
	0x93, 0x98, 0xa3, 0xa6, 0x4f, 0xe1, 0xb6, 0x2c, 0x84, 0x14, 0x0a, 0x0b, 0x87, 0x74, 0x49, 0x6f,
	0x23, 0x70, 0x2e, 0x4e, 0x77, 0xda, 0x56, 0xb3, 0x17, 0xc7, 0x05, 0x2a, 0x75, 0xa0, 0x8b, 0x24,
	0xe7, 0xec, 0x8a, 0x49, 0x9f, 0x41, 0xeb, 0xb0, 0xd2, 0x3b, 0x6b, 0x5d, 0xd2, 0xdb, 0x1c, 0xba,
	0xfe, 0x6a, 0x7f, 0xbe, 0xd9, 0xc2, 0x2c, 0xdb, 0x0b, 0xa0, 0x65, 0xf7, 0xee, 0xc2, 0x7a, 0xa2,
	0x31, 0x53, 0x0e, 0xe9, 0x36, 0x7b, 0x9b, 0x43, 0xef, 0xdf, 0x03, 0xf6, 0x35, 0x66, 0xcc, 0x08,
	0xbc, 0x1f, 0x04, 0xe0, 0x17, 0x4a, 0x87, 0x70, 0x2b, 0x34, 0x2e, 0xff, 0xeb, 0x7f, 0x49, 0xa4,
	0xfb, 0xd0, 0xfa, 0x80, 0x09, 0x3f, 0x32, 0xf6, 0x37, 0x82, 0xc1, 0xd9, 0xe5, 0x76, 0xe3, 0xdb,
	0xe5, 0xf6, 0x03, 0x23, 0x53, 0xf1, 0xb1, 0x9f, 0x88, 0x7e, 0x16, 0xea, 0x23, 0xff, 0x35, 0xf2,
	0x30, 0x2a, 0x47, 0x18, 0x5d, 0x9c, 0xee, 0x80, 0x9d, 0x3a, 0xc2, 0x88, 0xd9, 0x01, 0xf4, 0x1d,
	0xb4, 0x97, 0x96, 0x27, 0x0b, 0xcf, 0x93, 0x48, 0xe4, 0xef, 0x13, 0xee, 0x34, 0xab, 0x5c, 0x1e,
	0xff, 0xed, 0xac, 0xb1, 0xad, 0xc7, 0x42, 0xa4, 0x2f, 0x2b, 0x05, 0xa3, 0xf2, 0x1a, 0xe6, 0x7d,
	0x21, 0x40, 0xaf, 0x53, 0xa9, 0x04, 0x6a, 0x4f, 0x99, 0x48, 0x2c, 0x22, 0xcc, 0x75, 0xc8, 0xd1,
	0x26, 0xb9, 0x77, 0xf3, 0x95, 0xcb, 0x78, 0xc6, 0x57, 0x33, 0x5e, 0xe5, 0xba, 0x28, 0xd9, 0xbd,
	0xf0, 0x4f, 0xbc, 0x33, 0x82, 0xad, 0xd5, 0x64, 0x7a, 0x17, 0x9a, 0xc7, 0x58, 0x9a, 0xec, 0xd9,
	0xe2, 0x49, 0xdb, 0xb0, 0x3e, 0x0d, 0xd3, 0x13, 0x34, 0xe1, 0x32, 0x53, 0x3c, 0x5f, 0xdb, 0x25,
	0xc1, 0x8b, 0xb3, 0x99, 0x4b, 0xce, 0x67, 0x2e, 0xf9, 0x3e, 0x73, 0xc9, 0xe7, 0xb9, 0xdb, 0x38,
	0x9f, 0xbb, 0x8d, 0xaf, 0x73, 0xb7, 0xf1, 0xf6, 0x61, 0x2d, 0xf9, 0x8f, 0xf5, 0x2f, 0xa1, 0x4b,
	0x89, 0xea, 0xb0, 0x55, 0x61, 0x4f, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x83, 0x40, 0x7e, 0xd3,
	0x36, 0x03, 0x00, 0x00,
}

func (m *MsgSetBudget) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetBudget) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetBudget) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Budget != nil {
		{
			size, err := m.Budget.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintProtocolpool(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Proposer) > 0 {
		i -= len(m.Proposer)
		copy(dAtA[i:], m.Proposer)
		i = encodeVarintProtocolpool(dAtA, i, uint64(len(m.Proposer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Budget) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Budget) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Budget) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProtocolpool(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *BudgetItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BudgetItem) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BudgetItem) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProtocolPoolConfig != nil {
		{
			size, err := m.ProtocolPoolConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintProtocolpool(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.Weight.Size()
		i -= size
		if _, err := m.Weight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProtocolpool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintProtocolpool(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProtocolPoolConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProtocolPoolConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProtocolPoolConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AddressPercentage) > 0 {
		for k := range m.AddressPercentage {
			v := m.AddressPercentage[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintProtocolpool(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintProtocolpool(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintProtocolpool(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintProtocolpool(dAtA []byte, offset int, v uint64) int {
	offset -= sovProtocolpool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSetBudget) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Proposer)
	if l > 0 {
		n += 1 + l + sovProtocolpool(uint64(l))
	}
	if m.Budget != nil {
		l = m.Budget.Size()
		n += 1 + l + sovProtocolpool(uint64(l))
	}
	return n
}

func (m *Budget) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovProtocolpool(uint64(l))
		}
	}
	return n
}

func (m *BudgetItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovProtocolpool(uint64(l))
	}
	l = m.Weight.Size()
	n += 1 + l + sovProtocolpool(uint64(l))
	if m.ProtocolPoolConfig != nil {
		l = m.ProtocolPoolConfig.Size()
		n += 1 + l + sovProtocolpool(uint64(l))
	}
	return n
}

func (m *ProtocolPoolConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AddressPercentage) > 0 {
		for k, v := range m.AddressPercentage {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovProtocolpool(uint64(len(k))) + 1 + len(v) + sovProtocolpool(uint64(len(v)))
			n += mapEntrySize + 1 + sovProtocolpool(uint64(mapEntrySize))
		}
	}
	return n
}

func sovProtocolpool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProtocolpool(x uint64) (n int) {
	return sovProtocolpool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSetBudget) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocolpool
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
			return fmt.Errorf("proto: MsgSetBudget: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetBudget: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocolpool
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
				return ErrInvalidLengthProtocolpool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtocolpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Budget", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocolpool
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
				return ErrInvalidLengthProtocolpool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProtocolpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Budget == nil {
				m.Budget = &Budget{}
			}
			if err := m.Budget.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtocolpool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProtocolpool
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
func (m *Budget) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocolpool
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
			return fmt.Errorf("proto: Budget: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Budget: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocolpool
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
				return ErrInvalidLengthProtocolpool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProtocolpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &BudgetItem{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtocolpool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProtocolpool
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
func (m *BudgetItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocolpool
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
			return fmt.Errorf("proto: BudgetItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BudgetItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocolpool
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
				return ErrInvalidLengthProtocolpool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtocolpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocolpool
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
				return ErrInvalidLengthProtocolpool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtocolpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Weight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtocolPoolConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocolpool
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
				return ErrInvalidLengthProtocolpool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProtocolpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ProtocolPoolConfig == nil {
				m.ProtocolPoolConfig = &ProtocolPoolConfig{}
			}
			if err := m.ProtocolPoolConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtocolpool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProtocolpool
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
func (m *ProtocolPoolConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocolpool
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
			return fmt.Errorf("proto: ProtocolPoolConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProtocolPoolConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddressPercentage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocolpool
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
				return ErrInvalidLengthProtocolpool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProtocolpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AddressPercentage == nil {
				m.AddressPercentage = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowProtocolpool
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowProtocolpool
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthProtocolpool
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthProtocolpool
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowProtocolpool
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthProtocolpool
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthProtocolpool
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipProtocolpool(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthProtocolpool
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.AddressPercentage[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtocolpool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProtocolpool
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
func skipProtocolpool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProtocolpool
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
					return 0, ErrIntOverflowProtocolpool
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
					return 0, ErrIntOverflowProtocolpool
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
				return 0, ErrInvalidLengthProtocolpool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProtocolpool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProtocolpool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProtocolpool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProtocolpool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProtocolpool = fmt.Errorf("proto: unexpected end of group")
)