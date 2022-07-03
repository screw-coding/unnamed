package server

import (
	"bytes"
	"encoding/gob"
	"log"
	"net"
)

type Server struct {
	socketReadBufferSize  int
	socketWriteBufferSize int
	Packer                Packer
	Listener              net.Listener
	SessionManager        *SessionManager
	Router                map[uint32]HandlerFunc
}
type Option struct {
	SocketReadBufferSize  int
	SocketWriteBufferSize int
	Packer                Packer
}

func NewServer(opt *Option) *Server {
	if opt.Packer == nil {
		opt.Packer = NewDefaultPacker()
	}

	sessionM := NewSessionManager()

	return &Server{
		socketReadBufferSize:  opt.SocketReadBufferSize,
		socketWriteBufferSize: opt.SocketWriteBufferSize,
		Packer:                opt.Packer,
		SessionManager:        sessionM,
		Router:                make(map[uint32]HandlerFunc),
	}
}
func (s *Server) Serve(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	log.Println("服务端监听地址:", listener.Addr())
	s.Listener = listener
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Println("accept err:", err)
			return err
		}

		log.Println("A client connected.")
		go s.handleConn(connection)
	}

}

func (s *Server) handleConn(conn net.Conn) {
	session := NewSession(conn, NewDefaultPacker())
	s.SessionManager.AddSession(session)
	go session.readInbound(s.Router)
	go session.writeOutbound()
	select {
	case <-session.closed:
		log.Println("session closed")
	}
}

type HandlerFunc func(rt RouteContext)

//
// AddRoute
// @Description: 添加路由,针对某类MsgID的消息会被路由到某个处理函数
// @receiver s
// @param msgID
// @param handler
//
func (s *Server) AddRoute(msgID uint32, handler HandlerFunc) {
	s.Router[msgID] = handler
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
