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
		session := NewSession(connection, NewDefaultPacker())
		s.SessionManager.AddSession(session)
		go s.handleConn(session)
	}

}

func (s *Server) handleConn(sess *Session) {
	go sess.readInbound()
	go sess.writeOutbound()
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

func (s *Server) readInbound(readerChannel chan []byte, conn net.Conn) {
	defaultPacker := NewDefaultPacker()
	for {
		select {
		case data := <-readerChannel:
			msg := &Message{}
			BytesToStruct(data, msg)
			log.Println("服务端收到数据:", msg.Id, string(msg.Data))

			msg.Data = append(msg.Data, []byte("处理过了")...)
			_, err := conn.Write(defaultPacker.Pack(msg))
			if err != nil {
				log.Println("返回数据失败:", err)
			}
		}
	}
}

func (s *Server) writeOutbound(channel chan []byte, connection net.Conn) {

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
