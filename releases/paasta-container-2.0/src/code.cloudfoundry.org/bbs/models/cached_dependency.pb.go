// Code generated by protoc-gen-gogo.
// source: cached_dependency.proto
// DO NOT EDIT!

package models

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CachedDependency struct {
	Name              string `protobuf:"bytes,1,opt,name=name" json:"name"`
	From              string `protobuf:"bytes,2,opt,name=from" json:"from"`
	To                string `protobuf:"bytes,3,opt,name=to" json:"to"`
	CacheKey          string `protobuf:"bytes,4,opt,name=cache_key,json=cacheKey" json:"cache_key"`
	LogSource         string `protobuf:"bytes,5,opt,name=log_source,json=logSource" json:"log_source"`
	ChecksumAlgorithm string `protobuf:"bytes,6,opt,name=checksum_algorithm,json=checksumAlgorithm" json:"checksum_algorithm,omitempty"`
	ChecksumValue     string `protobuf:"bytes,7,opt,name=checksum_value,json=checksumValue" json:"checksum_value,omitempty"`
}

func (m *CachedDependency) Reset()                    { *m = CachedDependency{} }
func (*CachedDependency) ProtoMessage()               {}
func (*CachedDependency) Descriptor() ([]byte, []int) { return fileDescriptorCachedDependency, []int{0} }

func (m *CachedDependency) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CachedDependency) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *CachedDependency) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *CachedDependency) GetCacheKey() string {
	if m != nil {
		return m.CacheKey
	}
	return ""
}

func (m *CachedDependency) GetLogSource() string {
	if m != nil {
		return m.LogSource
	}
	return ""
}

func (m *CachedDependency) GetChecksumAlgorithm() string {
	if m != nil {
		return m.ChecksumAlgorithm
	}
	return ""
}

func (m *CachedDependency) GetChecksumValue() string {
	if m != nil {
		return m.ChecksumValue
	}
	return ""
}

func init() {
	proto.RegisterType((*CachedDependency)(nil), "models.CachedDependency")
}
func (this *CachedDependency) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*CachedDependency)
	if !ok {
		that2, ok := that.(CachedDependency)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.From != that1.From {
		return false
	}
	if this.To != that1.To {
		return false
	}
	if this.CacheKey != that1.CacheKey {
		return false
	}
	if this.LogSource != that1.LogSource {
		return false
	}
	if this.ChecksumAlgorithm != that1.ChecksumAlgorithm {
		return false
	}
	if this.ChecksumValue != that1.ChecksumValue {
		return false
	}
	return true
}
func (this *CachedDependency) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&models.CachedDependency{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "From: "+fmt.Sprintf("%#v", this.From)+",\n")
	s = append(s, "To: "+fmt.Sprintf("%#v", this.To)+",\n")
	s = append(s, "CacheKey: "+fmt.Sprintf("%#v", this.CacheKey)+",\n")
	s = append(s, "LogSource: "+fmt.Sprintf("%#v", this.LogSource)+",\n")
	s = append(s, "ChecksumAlgorithm: "+fmt.Sprintf("%#v", this.ChecksumAlgorithm)+",\n")
	s = append(s, "ChecksumValue: "+fmt.Sprintf("%#v", this.ChecksumValue)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringCachedDependency(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringCachedDependency(m github_com_gogo_protobuf_proto.Message) string {
	e := github_com_gogo_protobuf_proto.GetUnsafeExtensionsMap(m)
	if e == nil {
		return "nil"
	}
	s := "proto.NewUnsafeXXX_InternalExtensions(map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "})"
	return s
}
func (m *CachedDependency) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *CachedDependency) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintCachedDependency(data, i, uint64(len(m.Name)))
	i += copy(data[i:], m.Name)
	data[i] = 0x12
	i++
	i = encodeVarintCachedDependency(data, i, uint64(len(m.From)))
	i += copy(data[i:], m.From)
	data[i] = 0x1a
	i++
	i = encodeVarintCachedDependency(data, i, uint64(len(m.To)))
	i += copy(data[i:], m.To)
	data[i] = 0x22
	i++
	i = encodeVarintCachedDependency(data, i, uint64(len(m.CacheKey)))
	i += copy(data[i:], m.CacheKey)
	data[i] = 0x2a
	i++
	i = encodeVarintCachedDependency(data, i, uint64(len(m.LogSource)))
	i += copy(data[i:], m.LogSource)
	data[i] = 0x32
	i++
	i = encodeVarintCachedDependency(data, i, uint64(len(m.ChecksumAlgorithm)))
	i += copy(data[i:], m.ChecksumAlgorithm)
	data[i] = 0x3a
	i++
	i = encodeVarintCachedDependency(data, i, uint64(len(m.ChecksumValue)))
	i += copy(data[i:], m.ChecksumValue)
	return i, nil
}

func encodeFixed64CachedDependency(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32CachedDependency(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintCachedDependency(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *CachedDependency) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovCachedDependency(uint64(l))
	l = len(m.From)
	n += 1 + l + sovCachedDependency(uint64(l))
	l = len(m.To)
	n += 1 + l + sovCachedDependency(uint64(l))
	l = len(m.CacheKey)
	n += 1 + l + sovCachedDependency(uint64(l))
	l = len(m.LogSource)
	n += 1 + l + sovCachedDependency(uint64(l))
	l = len(m.ChecksumAlgorithm)
	n += 1 + l + sovCachedDependency(uint64(l))
	l = len(m.ChecksumValue)
	n += 1 + l + sovCachedDependency(uint64(l))
	return n
}

func sovCachedDependency(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCachedDependency(x uint64) (n int) {
	return sovCachedDependency(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *CachedDependency) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CachedDependency{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`From:` + fmt.Sprintf("%v", this.From) + `,`,
		`To:` + fmt.Sprintf("%v", this.To) + `,`,
		`CacheKey:` + fmt.Sprintf("%v", this.CacheKey) + `,`,
		`LogSource:` + fmt.Sprintf("%v", this.LogSource) + `,`,
		`ChecksumAlgorithm:` + fmt.Sprintf("%v", this.ChecksumAlgorithm) + `,`,
		`ChecksumValue:` + fmt.Sprintf("%v", this.ChecksumValue) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringCachedDependency(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *CachedDependency) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCachedDependency
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CachedDependency: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CachedDependency: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCachedDependency
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCachedDependency
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCachedDependency
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCachedDependency
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCachedDependency
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCachedDependency
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CacheKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCachedDependency
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCachedDependency
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CacheKey = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogSource", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCachedDependency
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCachedDependency
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LogSource = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChecksumAlgorithm", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCachedDependency
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCachedDependency
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChecksumAlgorithm = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChecksumValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCachedDependency
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCachedDependency
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChecksumValue = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCachedDependency(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCachedDependency
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
func skipCachedDependency(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCachedDependency
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowCachedDependency
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowCachedDependency
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCachedDependency
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCachedDependency
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipCachedDependency(data[start:])
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
	ErrInvalidLengthCachedDependency = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCachedDependency   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("cached_dependency.proto", fileDescriptorCachedDependency) }

var fileDescriptorCachedDependency = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x8f, 0xb1, 0x4e, 0x2a, 0x41,
	0x14, 0x86, 0x77, 0xf6, 0x72, 0x51, 0x26, 0x6a, 0x64, 0x1a, 0x27, 0xc4, 0x0c, 0xc4, 0x58, 0x58,
	0x20, 0xc4, 0x47, 0x10, 0x4d, 0x2c, 0xec, 0x96, 0xc4, 0x76, 0xb3, 0xcc, 0x0e, 0xbb, 0x84, 0x1d,
	0x0e, 0x59, 0x66, 0x4d, 0xb6, 0xf3, 0x11, 0x7c, 0x0c, 0x1f, 0x85, 0x92, 0xd2, 0x8a, 0xc8, 0xd8,
	0x18, 0x2b, 0x12, 0x5f, 0xc0, 0x70, 0x5c, 0x10, 0x62, 0x33, 0x39, 0xe7, 0x7c, 0xff, 0xf7, 0x27,
	0x43, 0x4f, 0x64, 0x20, 0x63, 0x15, 0xfa, 0xa1, 0x1a, 0xab, 0x51, 0xa8, 0x46, 0x32, 0x6f, 0x8d,
	0x53, 0x30, 0xc0, 0xca, 0x1a, 0x42, 0x95, 0x4c, 0x6a, 0x97, 0xd1, 0xc0, 0xc4, 0x59, 0xaf, 0x25,
	0x41, 0xb7, 0x23, 0x88, 0xa0, 0x8d, 0xb8, 0x97, 0xf5, 0x71, 0xc3, 0x05, 0xa7, 0x1f, 0xed, 0xec,
	0xcb, 0xa5, 0xc7, 0x37, 0x58, 0x79, 0xbb, 0x69, 0x64, 0x0d, 0x5a, 0x1a, 0x05, 0x5a, 0x71, 0xd2,
	0x20, 0x17, 0x95, 0xce, 0xc1, 0x74, 0x5e, 0x77, 0x3e, 0xe7, 0x75, 0xbc, 0x79, 0xf8, 0xae, 0x12,
	0xfd, 0x14, 0x34, 0x77, 0x77, 0x13, 0xab, 0x9b, 0x87, 0x2f, 0xab, 0x51, 0xd7, 0x00, 0xff, 0x87,
	0x9c, 0x16, 0xdc, 0x35, 0xe0, 0xb9, 0x06, 0x58, 0x8b, 0x56, 0xf0, 0x1b, 0xfe, 0x50, 0xe5, 0xbc,
	0x84, 0x91, 0x6a, 0x11, 0xf9, 0x05, 0xde, 0x3e, 0x8e, 0xf7, 0x2a, 0x67, 0x57, 0x94, 0x26, 0x10,
	0xf9, 0x13, 0xc8, 0x52, 0xa9, 0xf8, 0x7f, 0x14, 0x58, 0x21, 0x6c, 0x11, 0xaf, 0x92, 0x40, 0xd4,
	0xc5, 0x91, 0x75, 0x29, 0x93, 0xb1, 0x92, 0xc3, 0x49, 0xa6, 0xfd, 0x20, 0x89, 0x20, 0x1d, 0x98,
	0x58, 0xf3, 0x32, 0xaa, 0xe7, 0x85, 0x7a, 0xfa, 0x37, 0xd1, 0x04, 0x3d, 0x30, 0x4a, 0x8f, 0x4d,
	0xee, 0x55, 0xd7, 0xf4, 0x7a, 0x0d, 0xd9, 0x1d, 0x3d, 0xda, 0x28, 0x8f, 0x41, 0x92, 0x29, 0xbe,
	0x87, 0x85, 0x8d, 0xa2, 0x90, 0xef, 0xd2, 0xad, 0xb2, 0xc3, 0x35, 0x79, 0x58, 0x81, 0x4e, 0x73,
	0xb6, 0x10, 0xe4, 0x75, 0x21, 0x9c, 0xe5, 0x42, 0x90, 0x27, 0x2b, 0xc8, 0x8b, 0x15, 0x64, 0x6a,
	0x05, 0x99, 0x59, 0x41, 0xde, 0xac, 0x20, 0x1f, 0x56, 0x38, 0x4b, 0x2b, 0xc8, 0xf3, 0xbb, 0x70,
	0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x39, 0xb4, 0xc9, 0x65, 0xf4, 0x01, 0x00, 0x00,
}
