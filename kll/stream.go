package kll

import (
	"encoding/binary"
	"encoding/hex"
)

type Stream struct {
	Value []byte
	Pos   uint64
}

func (stream *Stream) Walk(val uint64) {
	stream.Pos += val
}

func (stream *Stream) Write(val []byte) {
	stream.Value = append(stream.Value, val...)
	stream.Walk(uint64(len(val)))
}

func (stream *Stream) WriteInt8(val int8) {
	stream.Write([]byte{byte(val)})
}
func (stream *Stream) WriteInt16(val int16) {
	a := []byte{0, 0}
	binary.LittleEndian.PutUint16(a, uint16(val))
	stream.Write(a)
}
func (stream *Stream) WriteInt32(val int32) {
	a := []byte{0, 0, 0, 0}
	binary.LittleEndian.PutUint32(a, uint32(val))
	stream.Write(a)
}
func (stream *Stream) WriteInt64(val int64) {
	a := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	binary.LittleEndian.PutUint64(a, uint64(val))
	stream.Write(a)
}

func (stream *Stream) WriteInt(val int) {
	a := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	binary.LittleEndian.PutUint64(a, uint64(val))
	stream.Write(a)
}

func (stream *Stream) WriteUInt8(val uint8) {
	stream.Write([]byte{byte(val)})
}
func (stream *Stream) WriteUInt16(val uint16) {
	a := []byte{0, 0}
	binary.LittleEndian.PutUint16(a, uint16(val))
	stream.Write(a)
}
func (stream *Stream) WriteUInt32(val uint32) {
	a := []byte{0, 0, 0, 0}
	binary.LittleEndian.PutUint32(a, uint32(val))
	stream.Write(a)
}
func (stream *Stream) WriteUInt64(val uint64) {
	a := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	binary.LittleEndian.PutUint64(a, uint64(val))
	stream.Write(a)
}

func (stream *Stream) WriteUInt(val uint) {
	a := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	binary.LittleEndian.PutUint64(a, uint64(val))
	stream.Write(a)
}

func (stream *Stream) Read(cout uint64) []byte {
	ret := stream.Value[stream.Pos : stream.Pos+cout]
	stream.Walk(cout)
	return ret
}
func (stream *Stream) ReadInt8() int8 {
	return int8(stream.Read(1)[0])
}
func (stream *Stream) ReadInt16() int16 {
	return int16(binary.LittleEndian.Uint16(stream.Read(2)))
}
func (stream *Stream) ReadInt32() int32 {
	return int32(binary.LittleEndian.Uint32(stream.Read(4)))
}
func (stream *Stream) ReadInt64() int64 {
	return int64(binary.LittleEndian.Uint64(stream.Read(8)))
}

func (stream *Stream) ReadInt() int {
	return int(binary.LittleEndian.Uint64(stream.Read(8)))
}

func (stream *Stream) ReadUInt8() uint8 {
	return uint8(stream.Read(1)[0])
}
func (stream *Stream) ReadUInt16() uint16 {
	return uint16(binary.LittleEndian.Uint16(stream.Read(2)))
}
func (stream *Stream) ReadUInt32() uint32 {
	return uint32(binary.LittleEndian.Uint32(stream.Read(4)))
}
func (stream *Stream) ReadUInt64() uint64 {
	return uint64(binary.LittleEndian.Uint64(stream.Read(8)))
}

func (stream *Stream) ReadUInt() uint {
	return uint(binary.LittleEndian.Uint64(stream.Read(8)))
}

func (stream *Stream) WriteString(st string) {
	stream.WriteUInt32(uint32(len(st)))
	stream.Write([]byte(st))
}
func (stream *Stream) ReadString() string {
	size := stream.ReadUInt32()
	return string(stream.Read(uint64(size)))
}

func (stream *Stream) String() string {
	ret := ""
	for i, v := range stream.Value {
		if i > 0 {
			ret += " "
		}
		ret += hex.EncodeToString([]byte{v})
	}
	return ret
}
func NewStream(pre_allocate uint) *Stream {
	return &Stream{Value: make([]byte, 0, pre_allocate)}
}
