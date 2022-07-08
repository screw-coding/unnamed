package util

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
)

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

func structToBytes(inter interface{}) (result []byte) {
	var buf bytes.Buffer
	_ = gob.NewEncoder(&buf).Encode(inter)
	return buf.Bytes()

}

func BytesToStruct(data []byte, inter interface{}) {
	buf := bytes.NewBuffer(data)
	_ = gob.NewDecoder(buf).Decode(inter)
	return
}
