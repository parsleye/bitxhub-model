// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: offchain_transmission.proto

package pb

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

type GetDataResponse_Type int32

const (
	GetDataResponse_DATA_GET_SUCCESS GetDataResponse_Type = 0
	GetDataResponse_DATA_OUT_OF_SIZE GetDataResponse_Type = 1
)

var GetDataResponse_Type_name = map[int32]string{
	0: "DATA_GET_SUCCESS",
	1: "DATA_OUT_OF_SIZE",
}

var GetDataResponse_Type_value = map[string]int32{
	"DATA_GET_SUCCESS": 0,
	"DATA_OUT_OF_SIZE": 1,
}

func (x GetDataResponse_Type) String() string {
	return proto.EnumName(GetDataResponse_Type_name, int32(x))
}

func (GetDataResponse_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4554c4eea08b0029, []int{1, 0}
}

type GetDataRequest struct {
	Index uint64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	From  string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To    string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	IsSrc bool   `protobuf:"varint,4,opt,name=isSrc,proto3" json:"isSrc,omitempty"`
	Req   []byte `protobuf:"bytes,5,opt,name=req,proto3" json:"req,omitempty"`
}

func (m *GetDataRequest) Reset()         { *m = GetDataRequest{} }
func (m *GetDataRequest) String() string { return proto.CompactTextString(m) }
func (*GetDataRequest) ProtoMessage()    {}
func (*GetDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4554c4eea08b0029, []int{0}
}
func (m *GetDataRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetDataRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDataRequest.Merge(m, src)
}
func (m *GetDataRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDataRequest proto.InternalMessageInfo

func (m *GetDataRequest) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *GetDataRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *GetDataRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *GetDataRequest) GetIsSrc() bool {
	if m != nil {
		return m.IsSrc
	}
	return false
}

func (m *GetDataRequest) GetReq() []byte {
	if m != nil {
		return m.Req
	}
	return nil
}

type GetDataResponse struct {
	Index uint64               `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	From  string               `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To    string               `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Type  GetDataResponse_Type `protobuf:"varint,4,opt,name=type,proto3,enum=pb.GetDataResponse_Type" json:"type,omitempty"`
	Msg   string               `protobuf:"bytes,5,opt,name=msg,proto3" json:"msg,omitempty"`
	Data  []byte               `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *GetDataResponse) Reset()         { *m = GetDataResponse{} }
func (m *GetDataResponse) String() string { return proto.CompactTextString(m) }
func (*GetDataResponse) ProtoMessage()    {}
func (*GetDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4554c4eea08b0029, []int{1}
}
func (m *GetDataResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetDataResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDataResponse.Merge(m, src)
}
func (m *GetDataResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDataResponse proto.InternalMessageInfo

func (m *GetDataResponse) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *GetDataResponse) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *GetDataResponse) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *GetDataResponse) GetType() GetDataResponse_Type {
	if m != nil {
		return m.Type
	}
	return GetDataResponse_DATA_GET_SUCCESS
}

func (m *GetDataResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *GetDataResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("pb.GetDataResponse_Type", GetDataResponse_Type_name, GetDataResponse_Type_value)
	proto.RegisterType((*GetDataRequest)(nil), "pb.GetDataRequest")
	proto.RegisterType((*GetDataResponse)(nil), "pb.GetDataResponse")
}

func init() { proto.RegisterFile("offchain_transmission.proto", fileDescriptor_4554c4eea08b0029) }

var fileDescriptor_4554c4eea08b0029 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x90, 0x31, 0x6a, 0xc3, 0x30,
	0x14, 0x86, 0x2d, 0xc7, 0x09, 0x8d, 0x28, 0xae, 0x11, 0x19, 0x04, 0x05, 0x61, 0x3c, 0x79, 0x28,
	0x1e, 0xd2, 0x13, 0xa4, 0x89, 0x1b, 0x3a, 0x05, 0x64, 0x67, 0xe9, 0x62, 0xe4, 0x44, 0x6e, 0x3d,
	0xd8, 0x52, 0x24, 0x15, 0x9a, 0x5b, 0xf4, 0x58, 0x1d, 0x03, 0x5d, 0x3a, 0x16, 0xfb, 0x22, 0xc5,
	0x0a, 0x64, 0xe8, 0xda, 0xed, 0x7b, 0x0f, 0xfd, 0xfa, 0x7e, 0x1e, 0xbc, 0x15, 0x55, 0xb5, 0x7b,
	0x65, 0x75, 0x5b, 0x18, 0xc5, 0x5a, 0xdd, 0xd4, 0x5a, 0xd7, 0xa2, 0x4d, 0xa4, 0x12, 0x46, 0x20,
	0x57, 0x96, 0x91, 0x82, 0xfe, 0x9a, 0x9b, 0x15, 0x33, 0x8c, 0xf2, 0xc3, 0x1b, 0xd7, 0x06, 0xcd,
	0xe0, 0xb8, 0x6e, 0xf7, 0xfc, 0x1d, 0x83, 0x10, 0xc4, 0x1e, 0x3d, 0x0f, 0x08, 0x41, 0xaf, 0x52,
	0xa2, 0xc1, 0x6e, 0x08, 0xe2, 0x29, 0xb5, 0x8c, 0x7c, 0xe8, 0x1a, 0x81, 0x47, 0x76, 0xe3, 0x1a,
	0x61, 0x93, 0x3a, 0x53, 0x3b, 0xec, 0x85, 0x20, 0xbe, 0xa2, 0xe7, 0x01, 0x05, 0x70, 0xa4, 0xf8,
	0x01, 0x8f, 0x43, 0x10, 0x5f, 0xd3, 0x01, 0xa3, 0x2f, 0x00, 0x6f, 0x2e, 0x52, 0x2d, 0x45, 0xab,
	0xf9, 0x3f, 0xac, 0x77, 0xd0, 0x33, 0x47, 0xc9, 0xad, 0xd4, 0x9f, 0xe3, 0x44, 0x96, 0xc9, 0x9f,
	0xcf, 0x93, 0xfc, 0x28, 0x39, 0xb5, 0xaf, 0x86, 0x36, 0x8d, 0x7e, 0xb1, 0x6d, 0xa6, 0x74, 0xc0,
	0xc1, 0xb1, 0x67, 0x86, 0xe1, 0x89, 0x2d, 0x68, 0x39, 0x9a, 0x43, 0x6f, 0xc8, 0xa0, 0x19, 0x0c,
	0x56, 0x8b, 0x7c, 0x51, 0xac, 0xd3, 0xbc, 0xc8, 0xb6, 0xcb, 0x65, 0x9a, 0x65, 0x81, 0x73, 0xd9,
	0x6e, 0xb6, 0x79, 0xb1, 0x79, 0x2c, 0xb2, 0xa7, 0xe7, 0x34, 0x00, 0x0f, 0xf8, 0xb3, 0x23, 0xe0,
	0xd4, 0x11, 0xf0, 0xd3, 0x11, 0xf0, 0xd1, 0x13, 0xe7, 0xd4, 0x13, 0xe7, 0xbb, 0x27, 0x4e, 0x39,
	0xb1, 0xe7, 0xbe, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x11, 0xcd, 0x75, 0x9d, 0x8d, 0x01, 0x00,
	0x00,
}

func (m *GetDataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetDataRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetDataRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Req) > 0 {
		i -= len(m.Req)
		copy(dAtA[i:], m.Req)
		i = encodeVarintOffchainTransmission(dAtA, i, uint64(len(m.Req)))
		i--
		dAtA[i] = 0x2a
	}
	if m.IsSrc {
		i--
		if m.IsSrc {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintOffchainTransmission(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintOffchainTransmission(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0x12
	}
	if m.Index != 0 {
		i = encodeVarintOffchainTransmission(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetDataResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetDataResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetDataResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintOffchainTransmission(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Msg) > 0 {
		i -= len(m.Msg)
		copy(dAtA[i:], m.Msg)
		i = encodeVarintOffchainTransmission(dAtA, i, uint64(len(m.Msg)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Type != 0 {
		i = encodeVarintOffchainTransmission(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x20
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintOffchainTransmission(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintOffchainTransmission(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0x12
	}
	if m.Index != 0 {
		i = encodeVarintOffchainTransmission(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintOffchainTransmission(dAtA []byte, offset int, v uint64) int {
	offset -= sovOffchainTransmission(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetDataRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovOffchainTransmission(uint64(m.Index))
	}
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovOffchainTransmission(uint64(l))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovOffchainTransmission(uint64(l))
	}
	if m.IsSrc {
		n += 2
	}
	l = len(m.Req)
	if l > 0 {
		n += 1 + l + sovOffchainTransmission(uint64(l))
	}
	return n
}

func (m *GetDataResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovOffchainTransmission(uint64(m.Index))
	}
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovOffchainTransmission(uint64(l))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovOffchainTransmission(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovOffchainTransmission(uint64(m.Type))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovOffchainTransmission(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovOffchainTransmission(uint64(l))
	}
	return n
}

func sovOffchainTransmission(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOffchainTransmission(x uint64) (n int) {
	return sovOffchainTransmission(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetDataRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOffchainTransmission
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
			return fmt.Errorf("proto: GetDataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetDataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffchainTransmission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffchainTransmission
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
				return ErrInvalidLengthOffchainTransmission
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOffchainTransmission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffchainTransmission
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
				return ErrInvalidLengthOffchainTransmission
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOffchainTransmission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsSrc", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffchainTransmission
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
			m.IsSrc = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Req", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffchainTransmission
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
				return ErrInvalidLengthOffchainTransmission
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthOffchainTransmission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Req = append(m.Req[:0], dAtA[iNdEx:postIndex]...)
			if m.Req == nil {
				m.Req = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOffchainTransmission(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOffchainTransmission
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
func (m *GetDataResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOffchainTransmission
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
			return fmt.Errorf("proto: GetDataResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetDataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffchainTransmission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffchainTransmission
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
				return ErrInvalidLengthOffchainTransmission
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOffchainTransmission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffchainTransmission
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
				return ErrInvalidLengthOffchainTransmission
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOffchainTransmission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffchainTransmission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= GetDataResponse_Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffchainTransmission
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
				return ErrInvalidLengthOffchainTransmission
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOffchainTransmission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffchainTransmission
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
				return ErrInvalidLengthOffchainTransmission
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthOffchainTransmission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOffchainTransmission(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOffchainTransmission
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
func skipOffchainTransmission(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOffchainTransmission
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
					return 0, ErrIntOverflowOffchainTransmission
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
					return 0, ErrIntOverflowOffchainTransmission
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
				return 0, ErrInvalidLengthOffchainTransmission
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOffchainTransmission
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOffchainTransmission
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOffchainTransmission        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOffchainTransmission          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOffchainTransmission = fmt.Errorf("proto: unexpected end of group")
)