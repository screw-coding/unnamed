package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	. "github.com/screw-coding/tcp"
	"log"
	"net"
	"os"
	"strings"
)

const (
	PAUSE = 2
	PLAY  = 1
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:5200")
	if err != nil {
		log.Fatal("dial err:", err)
		return
	}
	input := bufio.NewReader(os.Stdin)

	for {
		readString, err := input.ReadString('\n')
		if err != nil {
			return
		}
		readString = strings.TrimSpace(readString)

		if readString == "quit" {
			log.Println("quit")
			return
		}

		if readString == "play" {
			_ = play(conn)
		}

		if readString == "pause" {
			_ = pause(conn)
		}

		readerChannel := make(chan []byte, 16)
		go clientRead(readerChannel)
		go receiveAndUnpack(conn, readerChannel)

	}
}

func play(conn net.Conn) (err error) {
	packer := NewDefaultPacker()
	msg := Message{
		Id:   PLAY,
		Data: []byte("sjsjsjsjss"),
	}
	_, err = conn.Write(packer.Pack(structToBytes(msg)))
	if err != nil {
		log.Println("write err:", err)
		return
	}
	return
}

func pause(conn net.Conn) (err error) {
	packer := NewDefaultPacker()
	msg := &Message{
		Id:   PAUSE,
		Data: []byte("someshshshs"),
	}

	_, err = conn.Write(packer.Pack(structToBytes(msg)))
	if err != nil {
		log.Println("write err:", err)
		return
	}
	return
}

func receiveAndUnpack(conn net.Conn, readerChannel chan []byte) {
	buffer := make([]byte, 1024)
	tmpBuffer := make([]byte, 0)
	packer := NewDefaultPacker()
	for {
		size, err := conn.Read(buffer)
		if err != nil {
			return
		}
		if err != nil {
			log.Println("read err:", err)
			return
		}
		packer.Unpack(append(tmpBuffer, buffer[:size]...), readerChannel)
	}
}

func clientRead(readerChannel chan []byte) {
	for {
		select {
		case data := <-readerChannel:
			msg := &Message{}
			BytesToStruct(data, msg)
			log.Printf("客户端收到数据id: %d,data:%s", msg.Id, string(msg.Data))
		}
	}
}

func structToBytes(inter interface{}) (result []byte) {
	var buf bytes.Buffer
	_ = gob.NewEncoder(&buf).Encode(inter)
	return buf.Bytes()

}
