// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/transfer/transfer.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// MsgTransfer defines a msg to transfer fungible tokens (i.e Coins) between ICS20 enabled chains.
// See ICS Spec here: https://github.com/cosmos/ics/tree/master/spec/ics-020-fungible-token-transfer#data-structures
type MsgTransfer struct {
	// the port on which the packet will be sent
	SourcePort string `protobuf:"bytes,1,opt,name=source_port,json=sourcePort,proto3" json:"source_port,omitempty" yaml:"source_port"`
	// the channel by which the packet will be sent
	SourceChannel string `protobuf:"bytes,2,opt,name=source_channel,json=sourceChannel,proto3" json:"source_channel,omitempty" yaml:"source_channel"`
	// the tokens to be transferred
	Amount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
	// the sender address
	Sender github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,4,opt,name=sender,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"sender,omitempty"`
	// the recipient address on the destination chain
	Receiver string `protobuf:"bytes,5,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// Timeout height relative to the current block height.
	// The timeout is disabled when set to 0.
	TimeoutHeight uint64 `protobuf:"varint,6,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height,omitempty" yaml:"timeout_height"`
	// Timeout timestamp (in nanoseconds) relative to the current block timestamp.
	// The timeout is disabled when set to 0.
	TimeoutTimestamp uint64 `protobuf:"varint,7,opt,name=timeout_timestamp,json=timeoutTimestamp,proto3" json:"timeout_timestamp,omitempty" yaml:"timeout_timestamp"`
}

func (m *MsgTransfer) Reset()         { *m = MsgTransfer{} }
func (m *MsgTransfer) String() string { return proto.CompactTextString(m) }
func (*MsgTransfer) ProtoMessage()    {}
func (*MsgTransfer) Descriptor() ([]byte, []int) {
	return fileDescriptor_08134a70fd29e656, []int{0}
}
func (m *MsgTransfer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTransfer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTransfer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTransfer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTransfer.Merge(m, src)
}
func (m *MsgTransfer) XXX_Size() int {
	return m.Size()
}
func (m *MsgTransfer) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTransfer.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTransfer proto.InternalMessageInfo

func (m *MsgTransfer) GetSourcePort() string {
	if m != nil {
		return m.SourcePort
	}
	return ""
}

func (m *MsgTransfer) GetSourceChannel() string {
	if m != nil {
		return m.SourceChannel
	}
	return ""
}

func (m *MsgTransfer) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *MsgTransfer) GetSender() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *MsgTransfer) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *MsgTransfer) GetTimeoutHeight() uint64 {
	if m != nil {
		return m.TimeoutHeight
	}
	return 0
}

func (m *MsgTransfer) GetTimeoutTimestamp() uint64 {
	if m != nil {
		return m.TimeoutTimestamp
	}
	return 0
}

// FungibleTokenPacketData defines a struct for the packet payload
// See FungibleTokenPacketData spec: https://github.com/cosmos/ics/tree/master/spec/ics-020-fungible-token-transfer#data-structures
type FungibleTokenPacketData struct {
	// the tokens to be transferred
	Amount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
	// the sender address
	Sender string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	// the recipient address on the destination chain
	Receiver string `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver,omitempty"`
}

func (m *FungibleTokenPacketData) Reset()         { *m = FungibleTokenPacketData{} }
func (m *FungibleTokenPacketData) String() string { return proto.CompactTextString(m) }
func (*FungibleTokenPacketData) ProtoMessage()    {}
func (*FungibleTokenPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_08134a70fd29e656, []int{1}
}
func (m *FungibleTokenPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FungibleTokenPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FungibleTokenPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FungibleTokenPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FungibleTokenPacketData.Merge(m, src)
}
func (m *FungibleTokenPacketData) XXX_Size() int {
	return m.Size()
}
func (m *FungibleTokenPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_FungibleTokenPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_FungibleTokenPacketData proto.InternalMessageInfo

func (m *FungibleTokenPacketData) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *FungibleTokenPacketData) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *FungibleTokenPacketData) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

// FungibleTokenPacketAcknowledgement contains a boolean success flag and an optional error msg
// error msg is empty string on success
// See spec for onAcknowledgePacket: https://github.com/cosmos/ics/tree/master/spec/ics-020-fungible-token-transfer#packet-relay
type FungibleTokenPacketAcknowledgement struct {
	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *FungibleTokenPacketAcknowledgement) Reset()         { *m = FungibleTokenPacketAcknowledgement{} }
func (m *FungibleTokenPacketAcknowledgement) String() string { return proto.CompactTextString(m) }
func (*FungibleTokenPacketAcknowledgement) ProtoMessage()    {}
func (*FungibleTokenPacketAcknowledgement) Descriptor() ([]byte, []int) {
	return fileDescriptor_08134a70fd29e656, []int{2}
}
func (m *FungibleTokenPacketAcknowledgement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FungibleTokenPacketAcknowledgement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FungibleTokenPacketAcknowledgement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FungibleTokenPacketAcknowledgement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FungibleTokenPacketAcknowledgement.Merge(m, src)
}
func (m *FungibleTokenPacketAcknowledgement) XXX_Size() int {
	return m.Size()
}
func (m *FungibleTokenPacketAcknowledgement) XXX_DiscardUnknown() {
	xxx_messageInfo_FungibleTokenPacketAcknowledgement.DiscardUnknown(m)
}

var xxx_messageInfo_FungibleTokenPacketAcknowledgement proto.InternalMessageInfo

func (m *FungibleTokenPacketAcknowledgement) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *FungibleTokenPacketAcknowledgement) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgTransfer)(nil), "ibc.transfer.MsgTransfer")
	proto.RegisterType((*FungibleTokenPacketData)(nil), "ibc.transfer.FungibleTokenPacketData")
	proto.RegisterType((*FungibleTokenPacketAcknowledgement)(nil), "ibc.transfer.FungibleTokenPacketAcknowledgement")
}

func init() { proto.RegisterFile("ibc/transfer/transfer.proto", fileDescriptor_08134a70fd29e656) }

var fileDescriptor_08134a70fd29e656 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x3f, 0x6f, 0xd3, 0x40,
	0x1c, 0x8d, 0x49, 0x9b, 0x96, 0x4b, 0x8a, 0xe0, 0x28, 0xe5, 0x08, 0xc8, 0x8e, 0x3c, 0x79, 0x49,
	0x4c, 0x61, 0x40, 0x62, 0x22, 0x29, 0x42, 0x54, 0x08, 0xa9, 0xb2, 0x22, 0x06, 0x96, 0xca, 0x3e,
	0xff, 0x70, 0x4e, 0x89, 0xef, 0xa2, 0xbb, 0x33, 0xd0, 0x6f, 0xc1, 0x97, 0x60, 0xe1, 0x93, 0x74,
	0xec, 0xc0, 0xc0, 0x64, 0x50, 0xf2, 0x0d, 0x32, 0x32, 0x21, 0xff, 0x89, 0x89, 0xa5, 0x08, 0x31,
	0x74, 0xba, 0x7b, 0xef, 0xfd, 0x7e, 0xcf, 0xf7, 0x7b, 0xbe, 0x43, 0x0f, 0x59, 0x40, 0x5d, 0x2d,
	0x7d, 0xae, 0x3e, 0x80, 0xac, 0x36, 0x83, 0xb9, 0x14, 0x5a, 0xe0, 0x0e, 0x0b, 0xe8, 0x60, 0xcd,
	0x75, 0x0f, 0x23, 0x11, 0x89, 0x5c, 0x70, 0xb3, 0x5d, 0x51, 0xd3, 0xbd, 0x4b, 0x85, 0x8a, 0x85,
	0x72, 0x8b, 0xa5, 0x20, 0xed, 0xef, 0x4d, 0xd4, 0x7e, 0xab, 0xa2, 0x71, 0xd9, 0x8a, 0x9f, 0xa1,
	0xb6, 0x12, 0x89, 0xa4, 0x70, 0x3e, 0x17, 0x52, 0x13, 0xa3, 0x67, 0x38, 0x37, 0x47, 0x47, 0xab,
	0xd4, 0xc2, 0x17, 0x7e, 0x3c, 0x7b, 0x6e, 0x6f, 0x88, 0xb6, 0x87, 0x0a, 0x74, 0x26, 0xa4, 0xc6,
	0x2f, 0xd0, 0xad, 0x52, 0xa3, 0x13, 0x9f, 0x73, 0x98, 0x91, 0x1b, 0x79, 0xef, 0x83, 0x55, 0x6a,
	0xdd, 0xab, 0xf5, 0x96, 0xba, 0xed, 0x1d, 0x14, 0xc4, 0x49, 0x81, 0xf1, 0x3b, 0xd4, 0xf2, 0x63,
	0x91, 0x70, 0x4d, 0x9a, 0xbd, 0xa6, 0xd3, 0x7e, 0xd2, 0x19, 0x94, 0x27, 0x3d, 0x11, 0x8c, 0x8f,
	0x1e, 0x5f, 0xa6, 0x56, 0xe3, 0xdb, 0x4f, 0xcb, 0x89, 0x98, 0x9e, 0x24, 0xc1, 0x80, 0x8a, 0xd8,
	0xad, 0x0d, 0xd4, 0x57, 0xe1, 0xd4, 0xd5, 0x17, 0x73, 0x28, 0x1a, 0x94, 0x57, 0xba, 0xe1, 0x53,
	0xd4, 0x52, 0xc0, 0x43, 0x90, 0x64, 0xa7, 0x67, 0x38, 0x9d, 0xd1, 0xf1, 0xef, 0xd4, 0xea, 0xff,
	0x87, 0xcb, 0x90, 0xd2, 0x61, 0x18, 0x4a, 0x50, 0xca, 0x2b, 0x0d, 0x70, 0x17, 0xed, 0x4b, 0xa0,
	0xc0, 0x3e, 0x82, 0x24, 0xbb, 0xd9, 0x78, 0x5e, 0x85, 0xb3, 0x00, 0x34, 0x8b, 0x41, 0x24, 0xfa,
	0x7c, 0x02, 0x2c, 0x9a, 0x68, 0xd2, 0xea, 0x19, 0xce, 0xce, 0x66, 0x00, 0x75, 0xdd, 0xf6, 0x0e,
	0x4a, 0xe2, 0x75, 0x8e, 0xf1, 0x29, 0xba, 0xb3, 0xae, 0xc8, 0x56, 0xa5, 0xfd, 0x78, 0x4e, 0xf6,
	0x72, 0x93, 0x47, 0xab, 0xd4, 0x22, 0x75, 0x93, 0xaa, 0xc4, 0xf6, 0x6e, 0x97, 0xdc, 0xb8, 0xa2,
	0xbe, 0x1a, 0xe8, 0xfe, 0xab, 0x84, 0x47, 0x2c, 0x98, 0xc1, 0x58, 0x4c, 0x81, 0x9f, 0xf9, 0x74,
	0x0a, 0xfa, 0xa5, 0xaf, 0xfd, 0x8d, 0x9c, 0x8d, 0x6b, 0xcd, 0xf9, 0xa8, 0xca, 0x39, 0xff, 0xf3,
	0x5b, 0x43, 0x6b, 0xd6, 0x43, 0xb3, 0xc7, 0xc8, 0xde, 0x72, 0xcc, 0x21, 0x9d, 0x72, 0xf1, 0x69,
	0x06, 0x61, 0x04, 0x31, 0x70, 0x8d, 0x09, 0xda, 0x53, 0x09, 0xa5, 0xa0, 0x54, 0x7e, 0x21, 0xf7,
	0xbd, 0x35, 0xc4, 0x87, 0x68, 0x17, 0xa4, 0x14, 0xeb, 0x4f, 0x16, 0x60, 0xf4, 0xe6, 0x72, 0x61,
	0x1a, 0x57, 0x0b, 0xd3, 0xf8, 0xb5, 0x30, 0x8d, 0x2f, 0x4b, 0xb3, 0x71, 0xb5, 0x34, 0x1b, 0x3f,
	0x96, 0x66, 0xe3, 0xfd, 0xf1, 0x3f, 0xa7, 0xfa, 0xec, 0xb2, 0x80, 0xf6, 0xff, 0xbe, 0xb1, 0x6c,
	0xc8, 0xa0, 0x95, 0x3f, 0x94, 0xa7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x95, 0x43, 0xad,
	0x80, 0x03, 0x00, 0x00,
}

func (m *MsgTransfer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTransfer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTransfer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintTransfer(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x38
	}
	if m.TimeoutHeight != 0 {
		i = encodeVarintTransfer(dAtA, i, uint64(m.TimeoutHeight))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintTransfer(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTransfer(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTransfer(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.SourceChannel) > 0 {
		i -= len(m.SourceChannel)
		copy(dAtA[i:], m.SourceChannel)
		i = encodeVarintTransfer(dAtA, i, uint64(len(m.SourceChannel)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SourcePort) > 0 {
		i -= len(m.SourcePort)
		copy(dAtA[i:], m.SourcePort)
		i = encodeVarintTransfer(dAtA, i, uint64(len(m.SourcePort)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FungibleTokenPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FungibleTokenPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FungibleTokenPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintTransfer(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTransfer(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTransfer(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *FungibleTokenPacketAcknowledgement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FungibleTokenPacketAcknowledgement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FungibleTokenPacketAcknowledgement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintTransfer(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0x12
	}
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTransfer(dAtA []byte, offset int, v uint64) int {
	offset -= sovTransfer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgTransfer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SourcePort)
	if l > 0 {
		n += 1 + l + sovTransfer(uint64(l))
	}
	l = len(m.SourceChannel)
	if l > 0 {
		n += 1 + l + sovTransfer(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTransfer(uint64(l))
		}
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTransfer(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovTransfer(uint64(l))
	}
	if m.TimeoutHeight != 0 {
		n += 1 + sovTransfer(uint64(m.TimeoutHeight))
	}
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovTransfer(uint64(m.TimeoutTimestamp))
	}
	return n
}

func (m *FungibleTokenPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTransfer(uint64(l))
		}
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTransfer(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovTransfer(uint64(l))
	}
	return n
}

func (m *FungibleTokenPacketAcknowledgement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovTransfer(uint64(l))
	}
	return n
}

func sovTransfer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTransfer(x uint64) (n int) {
	return sovTransfer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgTransfer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransfer
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
			return fmt.Errorf("proto: MsgTransfer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTransfer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourcePort", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
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
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourcePort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChannel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
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
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
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
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = append(m.Sender[:0], dAtA[iNdEx:postIndex]...)
			if m.Sender == nil {
				m.Sender = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
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
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			m.TimeoutHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTransfer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTransfer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTransfer
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
func (m *FungibleTokenPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransfer
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
			return fmt.Errorf("proto: FungibleTokenPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FungibleTokenPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
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
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
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
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
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
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTransfer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTransfer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTransfer
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
func (m *FungibleTokenPacketAcknowledgement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransfer
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
			return fmt.Errorf("proto: FungibleTokenPacketAcknowledgement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FungibleTokenPacketAcknowledgement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
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
			m.Success = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
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
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTransfer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTransfer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTransfer
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
func skipTransfer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTransfer
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
					return 0, ErrIntOverflowTransfer
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
					return 0, ErrIntOverflowTransfer
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
				return 0, ErrInvalidLengthTransfer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTransfer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTransfer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTransfer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTransfer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTransfer = fmt.Errorf("proto: unexpected end of group")
)
