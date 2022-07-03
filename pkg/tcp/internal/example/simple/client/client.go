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
		//input := bufio.NewReader(os.Stdin)

		go func() {
			packer := NewDefaultPacker()
			msg, err := packer.Unpack(conn)
			if err != nil {
				log.Printf("unpack err:%s", err)
				return
			}
			log.Printf("rec <<< | id:(%d) size:(%d) data: %s", msg.Id, len(msg.Data), msg.Data)
		}()
		for {
			err := play(conn)
			if err != nil {
				log.Printf("play err:%s", err)
				conn, _ = net.Dial("tcp", "127.0.0.1:5200")

			}
			time.Sleep(time.Second * 3)
			//readString, err := input.ReadString('\n')
			//if err != nil {
			//	return
			//}
			//readString = strings.TrimSpace(readString)
			//
			//if readString == "quit" {
			//	log.Println("quit")
			//	return
			//}
			//
			//if readString == "play" {
			//	_ = play(conn)
			//}
			//
			//if readString == "pause" {
			//	_ = pause(conn)
			//}

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
	log.Printf("发送消息")
	return
}

func pause(conn net.Conn) (err error) {
	packer := NewDefaultPacker()
	msg := &Message{
		Id:   PAUSE,
		Data: []byte("someshshshs"),
	}

	_, err = conn.Write(packer.Pack(msg))
	if err != nil {
		log.Println("write err:", err)
		return
	}
	return
}
