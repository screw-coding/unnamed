package main

import (
	server "github.com/screw-coding/tcp"
	"log"
)

func main() {
	opt := &server.Option{
		SocketWriteBufferSize: 1024,
		SocketReadBufferSize:  1024,
	}
	newServer := server.NewServer(opt)
	err := newServer.Serve("127.0.0.1:5200")
	if err != nil {
		log.Fatal(err)
	}
}
