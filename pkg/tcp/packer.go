package server

import (
	"bytes"
	"encoding/binary"
	"io"
)

const (
	ConstHeader       = "myHeader"
	ConstHeaderLength = 8 //存储header的大小是8个字母,8个字节
	ConstDataLength   = 4 // 存储数据的大小是4个字节
)

type Packer interface {
	Pack(message *Message) []byte
	Unpack(reader io.Reader) *Message
}

func NewDefaultPacker() *DefaultPacker {
	return &DefaultPacker{}
}

type DefaultPacker struct {
}

//
// Pack
// @Description: 构建一个包
// @param message
// @return []byte
//
func (d *DefaultPacker) Pack(message *Message) []byte {
	return nil
}

func (d *DefaultPacker) Unpack(conn io.Reader) *Message {
	return nil
}

func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	_ = binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}

func intToBytes(i int) []byte {
	x := int32(i)
	buffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(buffer, binary.BigEndian, x)
	return buffer.Bytes()

}
