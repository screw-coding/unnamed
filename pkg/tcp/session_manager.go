package server

import (
	"log"
	"net"
	"sync"
)

type Session struct {
	id            int32             // 连接id,可以使用一个原子整数单调自增的方式生成
	conn          net.Conn          //用户实际的socket连接
	packer        Packer            // 包格式
	responseQueue chan RouteContext // 响应队列
	requestQueue  chan struct{}     // 请求队列
	closed        chan struct{}
}

func (s *Session) Send(rt *RouteContext) (ok bool) {
	select {
	case s.responseQueue <- *rt:
		return true
	}

}

//
// readInbound
// @Description: 读取客户端的数据
// @receiver s
//
func (s *Session) readInbound(router map[uint32]HandlerFunc) {
	for {
		packer := NewDefaultPacker()
		//err := s.conn.SetReadDeadline(time.Now().Add(time.Second * 3))
		//if err != nil {
		//	log.Printf("session %d set ReadDeadline err:%s", s.id, err)
		//}

		reqMsg, err := packer.Unpack(s.conn)
		log.Printf("receive {%d,%s} from %d:", reqMsg.Id, string(reqMsg.Data), s.id)
		if err != nil {
			log.Printf("session unpack err,session id %d,err:%s", s.id, err)
		}
		if reqMsg == nil {
			continue
		}
		go s.handleReq(router, reqMsg)
	}
}

func (s *Session) handleReq(router map[uint32]HandlerFunc, msg *Message) {
	//TODO:中间件处理,可以链式执行中间件
	currentRouterFunc := router[msg.Id]
	rt := &RouteContext{
		reqMsg:      msg,
		session:     s,
		responseMsg: &Message{},
	}
	currentRouterFunc(*rt)
	send := s.Send(rt)
	log.Printf("send to outbound result:%t", send)
}

func (s *Session) writeOutbound() {
	for {

		rt, ok := <-s.responseQueue
		if !ok {
			break
		}

		log.Println("writeOutbound receive ctx:", rt)
		bytes := s.packer.Pack(rt.Response())
		_, _ = s.conn.Write(bytes)

	}

}

type SessionManager struct {
	sessionMap map[int32]*Session
	NextID     int32
	lock       sync.Mutex
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessionMap: make(map[int32]*Session),
	}
}

func NewSession(conn net.Conn, packer Packer) *Session {
	return &Session{
		id:            0,
		conn:          conn,
		packer:        packer,
		responseQueue: make(chan RouteContext, 1024),
	}
}

//
// AddSession
// @Description: 添加一个session到session管理器
// @receiver sm
// @param session
//
func (sm *SessionManager) AddSession(session *Session) {
	sm.lock.Lock()
	sm.NextID++
	session.id = sm.NextID
	sm.sessionMap[sm.NextID] = session
	sm.lock.Unlock()
}

//
// RemoveSession
// @Description: 从session管理器中移除一个session
// @receiver sm
// @param session
//
func (sm *SessionManager) RemoveSession(session *Session) {
	sm.lock.Lock()
	sm.sessionMap[sm.NextID] = session
	delete(sm.sessionMap, session.id)
	sm.lock.Unlock()
}
