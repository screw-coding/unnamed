package server

import (
	"log"
	"net"
)

type Server struct {
	socketReadBufferSize  int
	socketWriteBufferSize int
	Packer                Packer
	Listener              net.Listener
	SessionManager        *SessionManager
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
		go s.handleConn(connection)
	}

}

func (s *Server) handleConn(connection net.Conn) {
	tmpBuffer := make([]byte, 0)
	readerChannel := make(chan []byte, 16)
	go s.read(readerChannel, connection)
	buffer := make([]byte, 1024)
	defaultPacker := NewDefaultPacker()
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

func (s *Server) read(readerChannel chan []byte, conn net.Conn) {
	defaultPacker := NewDefaultPacker()
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
