package server

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

const (
	ConstDataLength = 4 // 存储数据的大小是4个字节
	ConstIdLength   = 4 //存储id是4个字节
)

type Packer interface {
	Pack(message *Message) []byte
	Unpack(reader io.Reader) (*Message, error)
}

func NewDefaultPacker() *DefaultPacker {
	return &DefaultPacker{byteOrder: binary.LittleEndian}
}

type DefaultPacker struct {
	//默认小端序
	byteOrder binary.ByteOrder
}

//
// Pack
// @Description: 构建一个包
// dataSize(4)+id(4)+data(n)
// @param message
// @return []byte
//
func (d *DefaultPacker) Pack(message *Message) []byte {
	dataSize := len(message.Data)
	buffer := make([]byte, ConstDataLength+ConstIdLength+dataSize)
	d.byteOrder.PutUint32(buffer[:ConstDataLength], uint32(dataSize))
	d.byteOrder.PutUint32(buffer[ConstDataLength:ConstDataLength+ConstIdLength], message.Id)
	copy(buffer[ConstDataLength+ConstIdLength:], message.Data)
	return buffer
}

func (d *DefaultPacker) Unpack(reader io.Reader) (*Message, error) {
	headerBuffer := make([]byte, ConstDataLength+ConstIdLength)
	_, err := io.ReadFull(reader, headerBuffer)
	if err != nil {
		return nil, fmt.Errorf("read a packet error:%s", err)
	}
	dataSize := d.byteOrder.Uint32(headerBuffer[:ConstDataLength])
	id := d.byteOrder.Uint32(headerBuffer[ConstDataLength : ConstDataLength+ConstIdLength])
	data := make([]byte, dataSize)
	_, err = io.ReadFull(reader, data)
	if err != nil {
		return nil, fmt.Errorf("read data error:%s", err)
	}

	return NewMessage(id, data), nil
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
