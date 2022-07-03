package main

import (
	. "github.com/screw-coding/tcp"
	"log"
	"net"
	"time"
)

const (
	PAUSE = 2
	PLAY  = 1
)

func main() {
	for {
		log.Printf("开始尝试链接")
		conn, err := net.Dial("tcp", "127.0.0.1:5200")
		if err != nil {
			log.Printf("dial err:%s", err)
			time.Sleep(3 * time.Second)
			continue
		}
		log.Printf("已经连接上")

		go func() {
			packer := NewDefaultPacker()
			for {
				msg, err := packer.Unpack(conn)
				if err != nil {
					log.Printf("unpack err:%s", err)
					return
				}
				log.Printf("收到服务端返回的消息 <<< | id:(%d) size:(%d) data: %s", msg.Id, len(msg.Data), msg.Data)
			}

		}()
		for {
			err := play(conn)
			if err != nil {
				log.Printf("play err:%s", err)
				conn, _ = net.Dial("tcp", "127.0.0.1:5200")

			}
			time.Sleep(time.Second * 3)
		}

	}
}

func play(conn net.Conn) (err error) {
	packer := NewDefaultPacker()
	msg := &Message{
		Id:   PLAY,
		Data: []byte("SUEJksueiskUEUUE"),
	}
	log.Printf("开始发送消息")
	_, err = conn.Write(packer.Pack(msg))
	if err != nil {
		log.Println("write err:", err)
		return
	}
	log.Printf("发送消息成功")
	return
}
