// Code generated by protoc-gen-gogo.
// source: session.proto
// DO NOT EDIT!

/*
	Package session is a generated protocol buffer package.

	It is generated from these files:
		session.proto

	It has these top-level messages:
		Session
		Archive
*/
package session

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import sync "github.com/havoc-io/mutagen/sync"
import url "github.com/havoc-io/mutagen/url"

import time "time"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SessionVersion int32

const (
	SessionVersion_Unknown  SessionVersion = 0
	SessionVersion_Version1 SessionVersion = 1
)

var SessionVersion_name = map[int32]string{
	0: "Unknown",
	1: "Version1",
}
var SessionVersion_value = map[string]int32{
	"Unknown":  0,
	"Version1": 1,
}

func (x SessionVersion) String() string {
	return proto.EnumName(SessionVersion_name, int32(x))
}
func (SessionVersion) EnumDescriptor() ([]byte, []int) { return fileDescriptorSession, []int{0} }

type Session struct {
	Identifier           string         `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Version              SessionVersion `protobuf:"varint,2,opt,name=version,proto3,enum=session.SessionVersion" json:"version,omitempty"`
	CreationTime         *time.Time     `protobuf:"bytes,3,opt,name=creationTime,stdtime" json:"creationTime,omitempty"`
	CreatingVersionMajor uint32         `protobuf:"varint,4,opt,name=creatingVersionMajor,proto3" json:"creatingVersionMajor,omitempty"`
	CreatingVersionMinor uint32         `protobuf:"varint,5,opt,name=creatingVersionMinor,proto3" json:"creatingVersionMinor,omitempty"`
	CreatingVersionPatch uint32         `protobuf:"varint,6,opt,name=creatingVersionPatch,proto3" json:"creatingVersionPatch,omitempty"`
	Alpha                *url.URL       `protobuf:"bytes,7,opt,name=alpha" json:"alpha,omitempty"`
	Beta                 *url.URL       `protobuf:"bytes,8,opt,name=beta" json:"beta,omitempty"`
	Paused               bool           `protobuf:"varint,9,opt,name=paused,proto3" json:"paused,omitempty"`
}

func (m *Session) Reset()                    { *m = Session{} }
func (m *Session) String() string            { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()               {}
func (*Session) Descriptor() ([]byte, []int) { return fileDescriptorSession, []int{0} }

func (m *Session) GetCreationTime() *time.Time {
	if m != nil {
		return m.CreationTime
	}
	return nil
}

func (m *Session) GetAlpha() *url.URL {
	if m != nil {
		return m.Alpha
	}
	return nil
}

func (m *Session) GetBeta() *url.URL {
	if m != nil {
		return m.Beta
	}
	return nil
}

type Archive struct {
	Root *sync.Entry `protobuf:"bytes,1,opt,name=root" json:"root,omitempty"`
}

func (m *Archive) Reset()                    { *m = Archive{} }
func (m *Archive) String() string            { return proto.CompactTextString(m) }
func (*Archive) ProtoMessage()               {}
func (*Archive) Descriptor() ([]byte, []int) { return fileDescriptorSession, []int{1} }

func (m *Archive) GetRoot() *sync.Entry {
	if m != nil {
		return m.Root
	}
	return nil
}

func init() {
	proto.RegisterType((*Session)(nil), "session.Session")
	proto.RegisterType((*Archive)(nil), "session.Archive")
	proto.RegisterEnum("session.SessionVersion", SessionVersion_name, SessionVersion_value)
}
func (m *Session) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Session) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Identifier) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSession(dAtA, i, uint64(len(m.Identifier)))
		i += copy(dAtA[i:], m.Identifier)
	}
	if m.Version != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSession(dAtA, i, uint64(m.Version))
	}
	if m.CreationTime != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSession(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreationTime)))
		n1, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.CreationTime, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.CreatingVersionMajor != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintSession(dAtA, i, uint64(m.CreatingVersionMajor))
	}
	if m.CreatingVersionMinor != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintSession(dAtA, i, uint64(m.CreatingVersionMinor))
	}
	if m.CreatingVersionPatch != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintSession(dAtA, i, uint64(m.CreatingVersionPatch))
	}
	if m.Alpha != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintSession(dAtA, i, uint64(m.Alpha.Size()))
		n2, err := m.Alpha.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Beta != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintSession(dAtA, i, uint64(m.Beta.Size()))
		n3, err := m.Beta.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Paused {
		dAtA[i] = 0x48
		i++
		if m.Paused {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *Archive) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Archive) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Root != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSession(dAtA, i, uint64(m.Root.Size()))
		n4, err := m.Root.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func encodeFixed64Session(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Session(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintSession(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Session) Size() (n int) {
	var l int
	_ = l
	l = len(m.Identifier)
	if l > 0 {
		n += 1 + l + sovSession(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovSession(uint64(m.Version))
	}
	if m.CreationTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreationTime)
		n += 1 + l + sovSession(uint64(l))
	}
	if m.CreatingVersionMajor != 0 {
		n += 1 + sovSession(uint64(m.CreatingVersionMajor))
	}
	if m.CreatingVersionMinor != 0 {
		n += 1 + sovSession(uint64(m.CreatingVersionMinor))
	}
	if m.CreatingVersionPatch != 0 {
		n += 1 + sovSession(uint64(m.CreatingVersionPatch))
	}
	if m.Alpha != nil {
		l = m.Alpha.Size()
		n += 1 + l + sovSession(uint64(l))
	}
	if m.Beta != nil {
		l = m.Beta.Size()
		n += 1 + l + sovSession(uint64(l))
	}
	if m.Paused {
		n += 2
	}
	return n
}

func (m *Archive) Size() (n int) {
	var l int
	_ = l
	if m.Root != nil {
		l = m.Root.Size()
		n += 1 + l + sovSession(uint64(l))
	}
	return n
}

func sovSession(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSession(x uint64) (n int) {
	return sovSession(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Session) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSession
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Session: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Session: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= (SessionVersion(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreationTime == nil {
				m.CreationTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.CreationTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatingVersionMajor", wireType)
			}
			m.CreatingVersionMajor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatingVersionMajor |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatingVersionMinor", wireType)
			}
			m.CreatingVersionMinor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatingVersionMinor |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatingVersionPatch", wireType)
			}
			m.CreatingVersionPatch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatingVersionPatch |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alpha", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Alpha == nil {
				m.Alpha = &url.URL{}
			}
			if err := m.Alpha.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Beta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Beta == nil {
				m.Beta = &url.URL{}
			}
			if err := m.Beta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Paused", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Paused = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipSession(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSession
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
func (m *Archive) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSession
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Archive: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Archive: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Root", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Root == nil {
				m.Root = &sync.Entry{}
			}
			if err := m.Root.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSession(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSession
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
func skipSession(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSession
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
					return 0, ErrIntOverflowSession
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSession
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSession
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSession
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSession(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSession = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSession   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("session.proto", fileDescriptorSession) }

var fileDescriptorSession = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x51, 0x4d, 0x6e, 0xd4, 0x30,
	0x14, 0xae, 0xe9, 0x74, 0x92, 0x3a, 0x6d, 0x35, 0xb2, 0x10, 0x58, 0x23, 0x94, 0x89, 0xba, 0x0a,
	0xad, 0xea, 0xa8, 0xe1, 0x04, 0x54, 0xb0, 0x03, 0x09, 0x99, 0x96, 0xbd, 0x93, 0xba, 0x89, 0x61,
	0xe2, 0x17, 0x39, 0xce, 0xa0, 0xde, 0x82, 0x05, 0x0b, 0x8e, 0xc4, 0x92, 0x1b, 0x80, 0x86, 0x8b,
	0xa0, 0x38, 0x89, 0xe8, 0xa0, 0x88, 0x45, 0xa2, 0x7c, 0x7f, 0x8e, 0xbf, 0xf7, 0xf0, 0x71, 0x23,
	0x9b, 0x46, 0x81, 0x66, 0xb5, 0x01, 0x0b, 0xc4, 0x1b, 0xe0, 0xf2, 0xa2, 0x50, 0xb6, 0x6c, 0x33,
	0x96, 0x43, 0x95, 0x14, 0x50, 0x40, 0xe2, 0xf4, 0xac, 0xbd, 0x73, 0xc8, 0x01, 0xf7, 0xd5, 0xe7,
	0x96, 0xab, 0x02, 0xa0, 0x58, 0xcb, 0xbf, 0x2e, 0xab, 0x2a, 0xd9, 0x58, 0x51, 0xd5, 0x83, 0xe1,
	0xfc, 0xc1, 0x79, 0xa5, 0xd8, 0x40, 0x7e, 0xa1, 0x20, 0xa9, 0x5a, 0x2b, 0x0a, 0xa9, 0x93, 0xe6,
	0x5e, 0xe7, 0xee, 0x35, 0x98, 0x9f, 0xff, 0xcf, 0xdc, 0x9a, 0x75, 0xf7, 0xf4, 0xd6, 0xd3, 0xaf,
	0xfb, 0xd8, 0x7b, 0xdf, 0xdf, 0x99, 0x84, 0x18, 0xab, 0x5b, 0xa9, 0xad, 0xba, 0x53, 0xd2, 0x50,
	0x14, 0xa1, 0xf8, 0x90, 0x3f, 0x60, 0xc8, 0x25, 0xf6, 0x36, 0xd2, 0x74, 0x56, 0xfa, 0x28, 0x42,
	0xf1, 0x49, 0xfa, 0x94, 0x8d, 0xed, 0x87, 0x23, 0x3e, 0xf4, 0x32, 0x1f, 0x7d, 0xe4, 0x15, 0x3e,
	0xca, 0x8d, 0x14, 0x56, 0x81, 0xbe, 0x56, 0x95, 0xa4, 0xfb, 0x11, 0x8a, 0x83, 0x74, 0xc9, 0xfa,
	0xba, 0x6c, 0xac, 0xcb, 0xae, 0xc7, 0xba, 0x57, 0xb3, 0x2f, 0x3f, 0x57, 0x88, 0xef, 0xa4, 0x48,
	0x8a, 0x1f, 0xf7, 0x58, 0x17, 0xc3, 0x1f, 0xde, 0x8a, 0x8f, 0x60, 0xe8, 0x2c, 0x42, 0xf1, 0x31,
	0x9f, 0xd4, 0xa6, 0x32, 0x4a, 0x83, 0xa1, 0x07, 0xd3, 0x99, 0x4e, 0x9b, 0xc8, 0xbc, 0x13, 0x36,
	0x2f, 0xe9, 0x7c, 0x32, 0xe3, 0x34, 0x12, 0xe2, 0x03, 0xb1, 0xae, 0x4b, 0x41, 0x3d, 0x57, 0xcd,
	0x67, 0xdd, 0x6c, 0x6f, 0xf8, 0x1b, 0xde, 0xd3, 0xe4, 0x19, 0x9e, 0x65, 0xd2, 0x0a, 0xea, 0xff,
	0x23, 0x3b, 0x96, 0x3c, 0xc1, 0xf3, 0x5a, 0xb4, 0x8d, 0xbc, 0xa5, 0x87, 0x11, 0x8a, 0x7d, 0x3e,
	0xa0, 0xd3, 0x33, 0xec, 0xbd, 0x34, 0x79, 0xa9, 0x36, 0x92, 0xac, 0xf0, 0xcc, 0x00, 0x58, 0xb7,
	0x8f, 0x20, 0x0d, 0x98, 0xdb, 0xf3, 0x6b, 0x6d, 0xcd, 0x3d, 0x77, 0xc2, 0xd9, 0x39, 0x3e, 0xd9,
	0x1d, 0x3f, 0x09, 0xb0, 0x77, 0xa3, 0x3f, 0x69, 0xf8, 0xac, 0x17, 0x7b, 0xe4, 0x08, 0xfb, 0x03,
	0x7f, 0xb9, 0x40, 0x57, 0x8b, 0xef, 0xdb, 0x10, 0xfd, 0xd8, 0x86, 0xe8, 0xd7, 0x36, 0x44, 0xdf,
	0x7e, 0x87, 0x7b, 0xd9, 0xdc, 0x2d, 0xe1, 0xc5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x46,
	0x70, 0xed, 0xca, 0x02, 0x00, 0x00,
}
