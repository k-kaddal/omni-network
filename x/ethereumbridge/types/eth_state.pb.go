// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ethereumbridge/eth_state.proto

package types

import (
	fmt "fmt"
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

type EthState struct {
	Index       string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Address     string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Slot        string `protobuf:"bytes,3,opt,name=slot,proto3" json:"slot,omitempty"`
	Balance     string `protobuf:"bytes,4,opt,name=balance,proto3" json:"balance,omitempty"`
	Storage     string `protobuf:"bytes,5,opt,name=storage,proto3" json:"storage,omitempty"`
	Blocknumber string `protobuf:"bytes,6,opt,name=blocknumber,proto3" json:"blocknumber,omitempty"`
}

func (m *EthState) Reset()         { *m = EthState{} }
func (m *EthState) String() string { return proto.CompactTextString(m) }
func (*EthState) ProtoMessage()    {}
func (*EthState) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a36628c0771aaa1, []int{0}
}
func (m *EthState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EthState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EthState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EthState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthState.Merge(m, src)
}
func (m *EthState) XXX_Size() int {
	return m.Size()
}
func (m *EthState) XXX_DiscardUnknown() {
	xxx_messageInfo_EthState.DiscardUnknown(m)
}

var xxx_messageInfo_EthState proto.InternalMessageInfo

func (m *EthState) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *EthState) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *EthState) GetSlot() string {
	if m != nil {
		return m.Slot
	}
	return ""
}

func (m *EthState) GetBalance() string {
	if m != nil {
		return m.Balance
	}
	return ""
}

func (m *EthState) GetStorage() string {
	if m != nil {
		return m.Storage
	}
	return ""
}

func (m *EthState) GetBlocknumber() string {
	if m != nil {
		return m.Blocknumber
	}
	return ""
}

func init() {
	proto.RegisterType((*EthState)(nil), "kkaddal.omni.ethereumbridge.EthState")
}

func init() { proto.RegisterFile("ethereumbridge/eth_state.proto", fileDescriptor_5a36628c0771aaa1) }

var fileDescriptor_5a36628c0771aaa1 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x31, 0x4e, 0xc4, 0x30,
	0x10, 0x45, 0x63, 0xd8, 0x5d, 0xc0, 0x74, 0x16, 0x85, 0x25, 0x24, 0x6b, 0x45, 0x45, 0x43, 0x5c,
	0x70, 0x03, 0x24, 0x0a, 0x5a, 0xe8, 0x68, 0x90, 0x1d, 0x8f, 0x92, 0x28, 0x89, 0xbd, 0xb2, 0x27,
	0xd2, 0x72, 0x0b, 0x0e, 0xc1, 0x61, 0x28, 0xb7, 0xa4, 0x44, 0xc9, 0x45, 0x90, 0xed, 0x20, 0x01,
	0xdd, 0xbc, 0xff, 0x5f, 0x33, 0x9f, 0x0a, 0xc0, 0x06, 0x3c, 0x8c, 0x83, 0xf6, 0xad, 0xa9, 0x41,
	0x02, 0x36, 0x2f, 0x01, 0x15, 0x42, 0xb9, 0xf3, 0x0e, 0x1d, 0xbb, 0xec, 0x3a, 0x65, 0x8c, 0xea,
	0x4b, 0x37, 0xd8, 0xb6, 0xfc, 0x2b, 0x5f, 0xbd, 0x13, 0x7a, 0x7a, 0x8f, 0xcd, 0x53, 0xf4, 0xd9,
	0x05, 0x5d, 0xb7, 0xd6, 0xc0, 0x9e, 0x93, 0x2d, 0xb9, 0x3e, 0x7b, 0xcc, 0xc0, 0x38, 0x3d, 0x51,
	0xc6, 0x78, 0x08, 0x81, 0x1f, 0xa5, 0xfc, 0x07, 0x19, 0xa3, 0xab, 0xd0, 0x3b, 0xe4, 0xc7, 0x29,
	0x4e, 0x77, 0xb4, 0xb5, 0xea, 0x95, 0xad, 0x80, 0xaf, 0xb2, 0xbd, 0x60, 0x6c, 0x02, 0x3a, 0xaf,
	0x6a, 0xe0, 0xeb, 0xdc, 0x2c, 0xc8, 0xb6, 0xf4, 0x5c, 0xf7, 0xae, 0xea, 0xec, 0x38, 0x68, 0xf0,
	0x7c, 0x93, 0xda, 0xdf, 0xd1, 0xdd, 0xc3, 0xc7, 0x24, 0xc8, 0x61, 0x12, 0xe4, 0x6b, 0x12, 0xe4,
	0x6d, 0x16, 0xc5, 0x61, 0x16, 0xc5, 0xe7, 0x2c, 0x8a, 0x67, 0x59, 0xb7, 0xd8, 0x8c, 0xba, 0xac,
	0xdc, 0x20, 0xbb, 0x9b, 0xfc, 0xa9, 0x8c, 0x9f, 0xca, 0xbd, 0xfc, 0x37, 0x0c, 0xbe, 0xee, 0x20,
	0xe8, 0x4d, 0x5a, 0xe5, 0xf6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x8f, 0x89, 0x31, 0x37, 0x01,
	0x00, 0x00,
}

func (m *EthState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EthState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EthState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Blocknumber) > 0 {
		i -= len(m.Blocknumber)
		copy(dAtA[i:], m.Blocknumber)
		i = encodeVarintEthState(dAtA, i, uint64(len(m.Blocknumber)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Storage) > 0 {
		i -= len(m.Storage)
		copy(dAtA[i:], m.Storage)
		i = encodeVarintEthState(dAtA, i, uint64(len(m.Storage)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Balance) > 0 {
		i -= len(m.Balance)
		copy(dAtA[i:], m.Balance)
		i = encodeVarintEthState(dAtA, i, uint64(len(m.Balance)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Slot) > 0 {
		i -= len(m.Slot)
		copy(dAtA[i:], m.Slot)
		i = encodeVarintEthState(dAtA, i, uint64(len(m.Slot)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEthState(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintEthState(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEthState(dAtA []byte, offset int, v uint64) int {
	offset -= sovEthState(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EthState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovEthState(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEthState(uint64(l))
	}
	l = len(m.Slot)
	if l > 0 {
		n += 1 + l + sovEthState(uint64(l))
	}
	l = len(m.Balance)
	if l > 0 {
		n += 1 + l + sovEthState(uint64(l))
	}
	l = len(m.Storage)
	if l > 0 {
		n += 1 + l + sovEthState(uint64(l))
	}
	l = len(m.Blocknumber)
	if l > 0 {
		n += 1 + l + sovEthState(uint64(l))
	}
	return n
}

func sovEthState(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEthState(x uint64) (n int) {
	return sovEthState(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EthState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEthState
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
			return fmt.Errorf("proto: EthState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EthState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthState
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
				return ErrInvalidLengthEthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthState
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
				return ErrInvalidLengthEthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slot", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthState
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
				return ErrInvalidLengthEthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Slot = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthState
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
				return ErrInvalidLengthEthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Balance = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Storage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthState
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
				return ErrInvalidLengthEthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Storage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blocknumber", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthState
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
				return ErrInvalidLengthEthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blocknumber = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEthState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEthState
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
func skipEthState(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEthState
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
					return 0, ErrIntOverflowEthState
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
					return 0, ErrIntOverflowEthState
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
				return 0, ErrInvalidLengthEthState
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEthState
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEthState
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEthState        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEthState          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEthState = fmt.Errorf("proto: unexpected end of group")
)
