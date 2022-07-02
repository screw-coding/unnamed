package main

import (
	server "github.com/screw-coding/tcp"
	"log"
	"strings"
)

func main() {
	opt := &server.Option{
		SocketWriteBufferSize: 1024,
		SocketReadBufferSize:  1024,
	}
	newServer := server.NewServer(opt)
	newServer.AddRoute(1, func(rt server.RouteContext) {
		msg := rt.Request()
		data := msg.Data
		//一些业务处理
		newData := strings.ToLower(string(data))
		rt.Response().Id = 1
		rt.Response().Data = []byte(newData)

	})

	err := newServer.Serve("127.0.0.1:5200")
	if err != nil {
		log.Fatal(err)
	}

}
