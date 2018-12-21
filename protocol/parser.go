package protocol

import (
	"encoding/binary"
	// "fmt"
	"math"
)

const (
	TYPE_SYSPARA    = 1
	TYPE_CY_JOGUP   = 2
	TYPE_CY_JOGDOWN = 4
	TYPE_CY_RESET   = 8
	TYPE_CY_PARA    = 16
)

var FRAME_HEADER = [4]byte{0xff, 0xff, 0xff, 0xff}

//大端法！
func Float32ToByte(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, bits)

	return bytes
}

func ByteToFloat32(bytes []byte) float32 {
	bits := binary.BigEndian.Uint32(bytes)

	return math.Float32frombits(bits)
}

func Float64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, bits)

	return bytes
}

func ByteToFloat64(bytes []byte) float64 {
	bits := binary.BigEndian.Uint64(bytes)

	return math.Float64frombits(bits)
}

func ByteToUint16(bytes []byte) uint16 {
	return binary.BigEndian.Uint16(bytes)
}
func ByteToUint32(bytes []byte) uint32 {
	return binary.BigEndian.Uint32(bytes)
}

func Uint32ToByte(bits uint32) []byte {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, bits)

	return bytes
}

func GenerateXorValue(bs []byte) byte {
	var xorVal byte
	for _, b := range bs {
		xorVal ^= b
	}

	return xorVal
}
