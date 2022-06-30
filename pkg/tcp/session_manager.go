package server

import (
	"net"
	"sync"
)

type Session struct {
	id            int64         // 连接id,可以使用一个原子整数单调自增的方式生成
	conn          net.Conn      //用户实际的socket连接
	packer        Packer        // 包格式
	responseQueue chan struct{} // 响应队列
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

func (sm *SessionManager) AddSession(session *Session) {
	sm.lock.Lock()
	sm.NextID++
	sm.sessionMap[sm.NextID] = session
	sm.lock.Unlock()
}
