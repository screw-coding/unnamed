package main

import (
	"github.com/screw-coding/tcp"
	"log"
	"net"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	addr, err := net.ResolveTCPAddr("listener", ":5200")
	if err != nil {
		log.Fatal(err)
		return
	}

	listener, err := net.ListenTCP("listener", addr)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Listening on port 5200")

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Println("accept err:", err)
			return
		}

		log.Println("A client connected.")
		go handleConn(connection)
	}

}

func handleConn(connection net.Conn) {
	tmpBuffer := make([]byte, 0)
	readerChannel := make(chan []byte, 16)
	go read(readerChannel, connection)
	buffer := make([]byte, 1024)
	defaultPacker := server.NewDefaultPacker()
	for {
		n, err := connection.Read(buffer)
		if err != nil {
			log.Println(connection.RemoteAddr(), "connection error:", err)
			return
		}

		tmpBuffer = defaultPacker.Unpack(append(tmpBuffer, buffer[:n]...), readerChannel)
		log.Println("read:", tmpBuffer)
	}
}

func read(readerChannel chan []byte, conn net.Conn) {
	defaultPacker := server.NewDefaultPacker()
	for {
		select {
		case data := <-readerChannel:
			log.Println("服务端收到数据:", string(data))
			_, err := conn.Write(defaultPacker.Pack(append([]byte("服务端收到了且返回处理过的数据:"), data...)))
			if err != nil {
				log.Println("返回数据失败:", err)
			}
		}
	}
}
