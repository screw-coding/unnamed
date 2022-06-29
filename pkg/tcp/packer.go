package server

import (
	"bytes"
	"encoding/binary"
)

const (
	ConstHeader       = "myHeader"
	ConstHeaderLength = 8 //存储header的大小是8个字母,8个字节
	ConstDataLength   = 4 // 存储数据的大小是4个字节
)

type Packer interface {
	Pack(message []byte) []byte
	Unpack(buf []byte, readerChannel chan []byte) []byte
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
func (d *DefaultPacker) Pack(message []byte) []byte {
	return append(append([]byte(ConstHeader), intToBytes(len(message))...), message...)
}

func (d *DefaultPacker) Unpack(buf []byte, readerChannel chan []byte) []byte {
	length := len(buf)
	var i int
	for i = 0; i < length; i++ {
		if length < i+ConstHeaderLength+ConstDataLength {
			break
		}

		if string(buf[i:i+ConstHeaderLength]) == ConstHeader {
			//it's a packet
			messageLength := BytesToInt(buf[i+ConstHeaderLength : i+ConstHeaderLength+ConstDataLength])
			if length < i+ConstHeaderLength+ConstDataLength+messageLength {
				break
			}
			messageBytes := buf[i+ConstHeaderLength+ConstDataLength : i+ConstHeaderLength+ConstDataLength+messageLength]
			readerChannel <- messageBytes
			i = i + ConstHeaderLength + ConstDataLength + messageLength - 1
		}

	}
	if i == length {
		// 整个buf读完了,返回个空字节集
		return make([]byte, 0)
	}
	return buf[i:]
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
