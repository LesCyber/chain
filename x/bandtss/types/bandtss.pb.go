// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bandtss/v1beta1/bandtss.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MemberStatus is an enumeration of the possible statuses of a member.
type MemberStatus int32

const (
	// MEMBER_STATUS_UNSPECIFIED is unknown status.
	MEMBER_STATUS_UNSPECIFIED MemberStatus = 0
	// MEMBER_STATUS_ACTIVE is the active status.
	MEMBER_STATUS_ACTIVE MemberStatus = 1
	// MEMBER_STATUS_INACTIVE is the inactive status.
	MEMBER_STATUS_INACTIVE MemberStatus = 2
	// MEMBER_STATUS_JAIL is the status when member is jailed.
	MEMBER_STATUS_JAIL MemberStatus = 3
)

var MemberStatus_name = map[int32]string{
	0: "MEMBER_STATUS_UNSPECIFIED",
	1: "MEMBER_STATUS_ACTIVE",
	2: "MEMBER_STATUS_INACTIVE",
	3: "MEMBER_STATUS_JAIL",
}

var MemberStatus_value = map[string]int32{
	"MEMBER_STATUS_UNSPECIFIED": 0,
	"MEMBER_STATUS_ACTIVE":      1,
	"MEMBER_STATUS_INACTIVE":    2,
	"MEMBER_STATUS_JAIL":        3,
}

func (x MemberStatus) String() string {
	return proto.EnumName(MemberStatus_name, int32(x))
}

func (MemberStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2effaef066b71284, []int{0}
}

// Status maintains whether a member is an active member.
type Status struct {
	// address is the address of the member.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// status represents the current status of the member
	Status MemberStatus `protobuf:"varint,2,opt,name=status,proto3,enum=bandtss.v1beta1.MemberStatus" json:"status,omitempty"`
	// since is a block timestamp when a member has been activated/deactivated/jailed
	Since time.Time `protobuf:"bytes,3,opt,name=since,proto3,stdtime" json:"since"`
	// last_active is a latest block timestamp when a member is active
	LastActive time.Time `protobuf:"bytes,4,opt,name=last_active,json=lastActive,proto3,stdtime" json:"last_active"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_2effaef066b71284, []int{0}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Status.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return m.Size()
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Status) GetStatus() MemberStatus {
	if m != nil {
		return m.Status
	}
	return MEMBER_STATUS_UNSPECIFIED
}

func (m *Status) GetSince() time.Time {
	if m != nil {
		return m.Since
	}
	return time.Time{}
}

func (m *Status) GetLastActive() time.Time {
	if m != nil {
		return m.LastActive
	}
	return time.Time{}
}

func init() {
	proto.RegisterEnum("bandtss.v1beta1.MemberStatus", MemberStatus_name, MemberStatus_value)
	proto.RegisterType((*Status)(nil), "bandtss.v1beta1.Status")
}

func init() { proto.RegisterFile("bandtss/v1beta1/bandtss.proto", fileDescriptor_2effaef066b71284) }

var fileDescriptor_2effaef066b71284 = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xb1, 0xae, 0xd3, 0x30,
	0x18, 0x85, 0xe3, 0x7b, 0x4b, 0x01, 0x5f, 0x04, 0x95, 0x75, 0x75, 0x95, 0x46, 0xaa, 0x53, 0x31,
	0x55, 0x0c, 0x31, 0x2d, 0x62, 0xe9, 0x96, 0x96, 0x20, 0x05, 0xd1, 0x0a, 0x25, 0x29, 0x03, 0x4b,
	0xe4, 0xa4, 0x26, 0x8d, 0xd4, 0xc4, 0x55, 0xed, 0x56, 0x30, 0xb2, 0x31, 0xf6, 0x11, 0x90, 0x58,
	0x78, 0x94, 0x8e, 0x1d, 0x99, 0x00, 0xa5, 0x0b, 0x8f, 0x81, 0x1a, 0x27, 0x88, 0xb2, 0xdd, 0xcd,
	0x47, 0xdf, 0xf9, 0x8f, 0xff, 0x23, 0xfd, 0xb0, 0x13, 0xd1, 0x7c, 0x2e, 0x85, 0x20, 0xdb, 0x7e,
	0xc4, 0x24, 0xed, 0x93, 0x4a, 0x5b, 0xab, 0x35, 0x97, 0x1c, 0x3d, 0xaa, 0x65, 0x85, 0x8d, 0xeb,
	0x84, 0x27, 0xbc, 0x64, 0xe4, 0xf4, 0x52, 0x36, 0xc3, 0x4c, 0x38, 0x4f, 0x96, 0x8c, 0x94, 0x2a,
	0xda, 0xbc, 0x27, 0x32, 0xcd, 0x98, 0x90, 0x34, 0x5b, 0x55, 0x06, 0x1c, 0x73, 0x91, 0x71, 0x41,
	0x22, 0x2a, 0xd8, 0xdf, 0xaf, 0x62, 0x9e, 0xe6, 0x15, 0x6f, 0x2b, 0x1e, 0xaa, 0x64, 0x25, 0x14,
	0x7a, 0x5c, 0x00, 0xd8, 0xf4, 0x25, 0x95, 0x1b, 0x81, 0x74, 0x78, 0x97, 0xce, 0xe7, 0x6b, 0x26,
	0x84, 0x0e, 0xba, 0xa0, 0x77, 0xdf, 0xab, 0x25, 0x7a, 0x0e, 0x9b, 0xa2, 0xf4, 0xe8, 0x17, 0x5d,
	0xd0, 0x7b, 0x38, 0xe8, 0x58, 0xff, 0x2d, 0x6e, 0x4d, 0x58, 0x16, 0xb1, 0xb5, 0x0a, 0xf2, 0x2a,
	0x33, 0x1a, 0xc2, 0x3b, 0x22, 0xcd, 0x63, 0xa6, 0x5f, 0x76, 0x41, 0xef, 0x6a, 0x60, 0x58, 0xaa,
	0x87, 0x55, 0xf7, 0xb0, 0x82, 0xba, 0xc7, 0xe8, 0xde, 0xfe, 0x87, 0xa9, 0xed, 0x7e, 0x9a, 0xc0,
	0x53, 0x23, 0xc8, 0x81, 0x57, 0x4b, 0x2a, 0x64, 0x48, 0x63, 0x99, 0x6e, 0x99, 0xde, 0xb8, 0x45,
	0x02, 0x3c, 0x0d, 0xda, 0xe5, 0xdc, 0xb0, 0xf1, 0xfb, 0x8b, 0x09, 0x9e, 0x7c, 0x02, 0xf0, 0xc1,
	0xbf, 0x1b, 0xa2, 0x0e, 0x6c, 0x4f, 0x9c, 0xc9, 0xc8, 0xf1, 0x42, 0x3f, 0xb0, 0x83, 0x99, 0x1f,
	0xce, 0xa6, 0xfe, 0x1b, 0x67, 0xec, 0xbe, 0x74, 0x9d, 0x17, 0x2d, 0x0d, 0xe9, 0xf0, 0xfa, 0x1c,
	0xdb, 0xe3, 0xc0, 0x7d, 0xeb, 0xb4, 0x00, 0x32, 0xe0, 0xcd, 0x39, 0x71, 0xa7, 0x15, 0xbb, 0x40,
	0x37, 0x10, 0x9d, 0xb3, 0x57, 0xb6, 0xfb, 0xba, 0x75, 0x69, 0x34, 0x3e, 0x7f, 0xc5, 0xda, 0x68,
	0xfa, 0xad, 0xc0, 0x60, 0x5f, 0x60, 0x70, 0x28, 0x30, 0xf8, 0x55, 0x60, 0xb0, 0x3b, 0x62, 0xed,
	0x70, 0xc4, 0xda, 0xf7, 0x23, 0xd6, 0xde, 0x3d, 0x4d, 0x52, 0xb9, 0xd8, 0x44, 0x56, 0xcc, 0xb3,
	0xf2, 0x46, 0xca, 0x82, 0x31, 0x5f, 0x92, 0x78, 0x41, 0xd3, 0x9c, 0x6c, 0x07, 0xe4, 0x43, 0x7d,
	0x3b, 0x44, 0x7e, 0x5c, 0x31, 0x11, 0x35, 0x4b, 0xcb, 0xb3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x5d, 0x1a, 0x23, 0xaf, 0x63, 0x02, 0x00, 0x00,
}

func (this *Status) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Status)
	if !ok {
		that2, ok := that.(Status)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	if !this.Since.Equal(that1.Since) {
		return false
	}
	if !this.LastActive.Equal(that1.LastActive) {
		return false
	}
	return true
}
func (m *Status) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Status) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Status) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.LastActive, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.LastActive):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintBandtss(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Since, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Since):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintBandtss(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	if m.Status != 0 {
		i = encodeVarintBandtss(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintBandtss(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBandtss(dAtA []byte, offset int, v uint64) int {
	offset -= sovBandtss(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Status) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovBandtss(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovBandtss(uint64(m.Status))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Since)
	n += 1 + l + sovBandtss(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.LastActive)
	n += 1 + l + sovBandtss(uint64(l))
	return n
}

func sovBandtss(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBandtss(x uint64) (n int) {
	return sovBandtss(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Status) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBandtss
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
			return fmt.Errorf("proto: Status: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Status: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBandtss
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
				return ErrInvalidLengthBandtss
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBandtss
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBandtss
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= MemberStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Since", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBandtss
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
				return ErrInvalidLengthBandtss
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBandtss
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Since, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastActive", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBandtss
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
				return ErrInvalidLengthBandtss
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBandtss
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.LastActive, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBandtss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBandtss
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
func skipBandtss(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBandtss
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
					return 0, ErrIntOverflowBandtss
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
					return 0, ErrIntOverflowBandtss
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
				return 0, ErrInvalidLengthBandtss
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBandtss
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBandtss
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBandtss        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBandtss          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBandtss = fmt.Errorf("proto: unexpected end of group")
)
