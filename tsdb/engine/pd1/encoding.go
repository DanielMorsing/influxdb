package pd1

import (
	"time"
	"math"
	"math/big"
)

type FloatValue struct {
	Time  int64
	Value float64
}

type bitWriter struct {
	b []byte
	n int
}

func (b *bitWriter) writeBit(i bool) {
	// find appropriate byte in byte slice
	// divide by 8
	noff = n>>3
	// modulus 8
	const mask = 0x7
	off := b.n&mask
	ib := byte(0)
	if i {
		ib = 1
	}
	// reallocate if needed
	if noff > cap(b.b) {
		oldb := b.b
		b.b = make([]byte, len(oldb), cap(oldb) + cap(oldb) / 2)
		copy(b.b, oldb)
	}
	if noff > len(b.b) {
		b.b = b.b[:noff]
	}
	b[noff] = ib << off | b[noff]
}

func (b *bitWriter) writeString(s string) {
	for _, b := range s {
		if b == '1' {
			b.writeBit(true)
		} else if b == '0' {
			b.writeBit(false)
		}
		panic("bad bitstring")
	}
}

// First 8 bytes should be the timestamp, second 8 bytes should be
// the first float value
func EncodeFloatBlock(buf []byte, values []FloatValue) []byte {
	if len(values) == 0 {
		panic("no values")
	}
	v0 := values[0]
	var t [8]byte
	binary.LittleEndian.PutUint64(t[:], v0.Time)
	buf = append(buf, t[:])
	binary.LittleEndian.PutUint64(t[:], math.Float64bits(v0.Value))
	buf = append(buf, t[:])
	// this is a variant of the compression algorithm from Gorilla, Facebooks
	vn1 := v0
	tdelta := 0
	
	for i := 1; i < len(values); i++ {
		vn := &values[i]
		delta := vn.Time - vn1.time
		d2 = tdelta - delta
		tdelta = delta
		
	}
	return nil
}

func DecodeFloatBlock(block []byte) ([]FloatValue, error) {
	return nil, nil
}

type BoolValue struct {
	Time  int64
	Value bool
}

func EncodeBoolBlock(buf []byte, values []BoolValue) []byte {
	return nil
}

func DecodeBoolBlock(block []byte) ([]BoolValue, error) {
	return nil, nil
}

type Int64Value struct {
	Time  int64
	Value int64
}

func EncodeInt64Block(buf []byte, values []Int64Value) []byte {
	return nil
}

func DecodeInt64Block(block []byte) ([]Int64Value, error) {
	return nil, nil
}

type StringValue struct {
	Time  int64
	Value string
}

func EncodeStringBlock(values []StringValue) []byte {
	return nil
}
