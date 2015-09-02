// Code generated by protoc-gen-gogo.
// source: tags.proto
// DO NOT EDIT!

/*
Package tags is a generated protocol buffer package.

It is generated from these files:
	tags.proto

It has these top-level messages:
	Outside
	Inside
*/
package tags

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Outside struct {
	*Inside          `protobuf:"bytes,1,opt,embedded=Inside" json:""`
	Field2           *string `protobuf:"bytes,2,opt" json:"MyField2" xml:",comment"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Outside) Reset()         { *m = Outside{} }
func (m *Outside) String() string { return proto.CompactTextString(m) }
func (*Outside) ProtoMessage()    {}

func (m *Outside) GetField2() string {
	if m != nil && m.Field2 != nil {
		return *m.Field2
	}
	return ""
}

type Inside struct {
	Field1           *string `protobuf:"bytes,1,opt" json:"MyField1" xml:",chardata"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Inside) Reset()         { *m = Inside{} }
func (m *Inside) String() string { return proto.CompactTextString(m) }
func (*Inside) ProtoMessage()    {}

func (m *Inside) GetField1() string {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return ""
}

func NewPopulatedOutside(r randyTags, easy bool) *Outside {
	this := &Outside{}
	if r.Intn(10) != 0 {
		this.Inside = NewPopulatedInside(r, easy)
	}
	if r.Intn(10) != 0 {
		v1 := randStringTags(r)
		this.Field2 = &v1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTags(r, 3)
	}
	return this
}

func NewPopulatedInside(r randyTags, easy bool) *Inside {
	this := &Inside{}
	if r.Intn(10) != 0 {
		v2 := randStringTags(r)
		this.Field1 = &v2
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTags(r, 2)
	}
	return this
}

type randyTags interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTags(r randyTags) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringTags(r randyTags) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneTags(r)
	}
	return string(tmps)
}
func randUnrecognizedTags(r randyTags, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldTags(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldTags(data []byte, r randyTags, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateTags(data, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		data = encodeVarintPopulateTags(data, uint64(v4))
	case 1:
		data = encodeVarintPopulateTags(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateTags(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateTags(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateTags(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateTags(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
