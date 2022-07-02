package server

import (
	"net"
	"sync"
)

type Session struct {
	id            int32         // 连接id,可以使用一个原子整数单调自增的方式生成
	conn          net.Conn      //用户实际的socket连接
	packer        Packer        // 包格式
	responseQueue chan struct{} // 响应队列
	requestQueue  chan struct{} // 请求队列
}

func (s *Session) Send() {
}

//
// readInbound
// @Description: 读取客户端的数据
// @receiver s
//
func (s *Session) readInbound() {
	for {
		packer := NewDefaultPacker()
		reqMsg := packer.Unpack(s.conn)
		go handleReq(reqMsg)
	}
}

func handleReq(msg *Message) {

}

func (s *Session) writeOutbound() {

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
		id:     0,
		conn:   conn,
		packer: packer,
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
