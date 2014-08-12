// Code generated by protoc-gen-gogo.
// source: v2_store_cmd.proto
// DO NOT EDIT!

/*
	Package etcd is a generated protocol buffer package.

	It is generated from these files:
		v2_store_cmd.proto

	It has these top-level messages:
		Cmd
*/
package etcd

import proto "code.google.com/p/gogoprotobuf/proto"
import json "encoding/json"
import math "math"

// discarding unused import gogoproto "code.google.com/p/gogoprotobuf/gogoproto/gogo.pb"

import io "io"
import code_google_com_p_gogoprotobuf_proto "code.google.com/p/gogoprotobuf/proto"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Cmd struct {
	Type             int32   `protobuf:"varint,1,req,name=type" json:"type"`
	Key              string  `protobuf:"bytes,2,req,name=key" json:"key"`
	Value            *string `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	PrevValue        *string `protobuf:"bytes,4,opt,name=prevValue" json:"prevValue,omitempty"`
	PrevIndex        *uint64 `protobuf:"varint,5,opt,name=prevIndex" json:"prevIndex,omitempty"`
	Dir              *bool   `protobuf:"varint,6,opt,name=dir" json:"dir,omitempty"`
	Recursive        *bool   `protobuf:"varint,7,opt,name=recursive" json:"recursive,omitempty"`
	Unique           *bool   `protobuf:"varint,8,opt,name=unique" json:"unique,omitempty"`
	Sorted           *bool   `protobuf:"varint,9,opt,name=sorted" json:"sorted,omitempty"`
	Time             []byte  `protobuf:"bytes,10,opt,name=time" json:"time,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Cmd) Reset()         { *m = Cmd{} }
func (m *Cmd) String() string { return proto.CompactTextString(m) }
func (*Cmd) ProtoMessage()    {}

func init() {
}
func (m *Cmd) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				m.Type |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(data[index:postIndex])
			index = postIndex
		case 3:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[index:postIndex])
			m.Value = &s
			index = postIndex
		case 4:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[index:postIndex])
			m.PrevValue = &s
			index = postIndex
		case 5:
			if wireType != 0 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.PrevIndex = &v
		case 6:
			if wireType != 0 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Dir = &b
		case 7:
			if wireType != 0 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Recursive = &b
		case 8:
			if wireType != 0 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Unique = &b
		case 9:
			if wireType != 0 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Sorted = &b
		case 10:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Time = append(m.Time, data[index:postIndex]...)
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *Cmd) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovV2StoreCmd(uint64(uint32(m.Type)))
	l = len(m.Key)
	n += 1 + l + sovV2StoreCmd(uint64(l))
	if m.Value != nil {
		l = len(*m.Value)
		n += 1 + l + sovV2StoreCmd(uint64(l))
	}
	if m.PrevValue != nil {
		l = len(*m.PrevValue)
		n += 1 + l + sovV2StoreCmd(uint64(l))
	}
	if m.PrevIndex != nil {
		n += 1 + sovV2StoreCmd(uint64(*m.PrevIndex))
	}
	if m.Dir != nil {
		n += 2
	}
	if m.Recursive != nil {
		n += 2
	}
	if m.Unique != nil {
		n += 2
	}
	if m.Sorted != nil {
		n += 2
	}
	if m.Time != nil {
		l = len(m.Time)
		n += 1 + l + sovV2StoreCmd(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovV2StoreCmd(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozV2StoreCmd(x uint64) (n int) {
	return sovV2StoreCmd(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Cmd) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Cmd) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintV2StoreCmd(data, i, uint64(uint32(m.Type)))
	data[i] = 0x12
	i++
	i = encodeVarintV2StoreCmd(data, i, uint64(len(m.Key)))
	i += copy(data[i:], m.Key)
	if m.Value != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintV2StoreCmd(data, i, uint64(len(*m.Value)))
		i += copy(data[i:], *m.Value)
	}
	if m.PrevValue != nil {
		data[i] = 0x22
		i++
		i = encodeVarintV2StoreCmd(data, i, uint64(len(*m.PrevValue)))
		i += copy(data[i:], *m.PrevValue)
	}
	if m.PrevIndex != nil {
		data[i] = 0x28
		i++
		i = encodeVarintV2StoreCmd(data, i, uint64(*m.PrevIndex))
	}
	if m.Dir != nil {
		data[i] = 0x30
		i++
		if *m.Dir {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.Recursive != nil {
		data[i] = 0x38
		i++
		if *m.Recursive {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.Unique != nil {
		data[i] = 0x40
		i++
		if *m.Unique {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.Sorted != nil {
		data[i] = 0x48
		i++
		if *m.Sorted {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.Time != nil {
		data[i] = 0x52
		i++
		i = encodeVarintV2StoreCmd(data, i, uint64(len(m.Time)))
		i += copy(data[i:], m.Time)
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func encodeFixed64V2StoreCmd(data []byte, offset int, v uint64) int {
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
func encodeFixed32V2StoreCmd(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintV2StoreCmd(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}