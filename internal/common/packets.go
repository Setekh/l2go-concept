package common

// TODO make a pool of buffers!!

import (
	"bytes"
	"encoding/binary"
)

type Buffer struct {
	bytes.Buffer
}

func NewBuffer() *Buffer {
	return &Buffer{}
}

func (b *Buffer) WriteC(value byte) {
	b.WriteByte(value)
}

func (b *Buffer) WriteF(value float64) {
	binary.Write(b, binary.LittleEndian, value)
}

func (b *Buffer) WriteD(value uint32) {
	binary.Write(b, binary.LittleEndian, value)
}

func (b *Buffer) WriteSD(value int32) {
	binary.Write(b, binary.LittleEndian, value)
}

func (b *Buffer) WriteH(value uint16) {
	binary.Write(b, binary.LittleEndian, value)
}

func (b *Buffer) WriteUInt8(value uint8) {
	binary.Write(b, binary.LittleEndian, value)
}

func (b *Buffer) WriteFloat64(value float64) {
	binary.Write(b, binary.LittleEndian, value)
}

func (b *Buffer) WriteFloat32(value float32) {
	binary.Write(b, binary.LittleEndian, value)
}

func (b *Buffer) WriteBytes(value []byte) {
	b.Write(value)
}

func (b *Buffer) WriteS(value string) {
	for i := 0; i < len(value); i++ {
		char := value[i]
		binary.Write(b, binary.LittleEndian, uint16(char))
	}
	binary.Write(b, binary.LittleEndian, uint16(0x00))
}

type Reader struct {
	*bytes.Reader
}

func NewReader(buffer []byte) *Reader {
	return &Reader{bytes.NewReader(buffer)}
}

func (r *Reader) ReadBytes(number int) []byte {
	buffer := make([]byte, number)
	n, _ := r.Read(buffer)
	if n < number {
		return []byte{}
	}

	return buffer
}

func (r *Reader) ReadF() uint64 {
	var result uint64

	buffer := make([]byte, 8)
	n, _ := r.Read(buffer)
	if n < 8 {
		return 0
	}

	buf := bytes.NewBuffer(buffer)
	binary.Read(buf, binary.LittleEndian, &result)

	return result
}

func (r *Reader) ReadD() uint32 {
	var result uint32

	buffer := make([]byte, 4)
	n, _ := r.Read(buffer)
	if n < 4 {
		return 0
	}

	buf := bytes.NewBuffer(buffer)

	binary.Read(buf, binary.LittleEndian, &result)

	return result
}

func (r *Reader) ReadSD() int32 {
	var result int32

	buffer := make([]byte, 4)
	n, _ := r.Read(buffer)
	if n < 4 {
		return 0
	}

	buf := bytes.NewBuffer(buffer)

	binary.Read(buf, binary.LittleEndian, &result)

	return result
}

func (r *Reader) ReadH() uint16 {
	var result uint16

	buffer := make([]byte, 2)
	n, _ := r.Read(buffer)
	if n < 2 {
		return 0
	}

	buf := bytes.NewBuffer(buffer)

	binary.Read(buf, binary.LittleEndian, &result)

	return result
}

func (r *Reader) ReadUInt8() uint8 {
	var result uint8

	buffer := make([]byte, 1)
	n, _ := r.Read(buffer)
	if n < 1 {
		return 0
	}

	buf := bytes.NewBuffer(buffer)

	binary.Read(buf, binary.LittleEndian, &result)

	return result
}

func (r *Reader) ReadString() string {
	var result []byte
	var firstByte, secondByte byte

	for {
		firstByte, _ = r.ReadByte()
		secondByte, _ = r.ReadByte()
		if firstByte == 0x00 && secondByte == 0x00 {
			break
		} else {
			result = append(result, firstByte)
		}
	}

	return string(result)
}
