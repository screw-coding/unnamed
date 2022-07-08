package server

import (
	"github.com/stretchr/testify/assert"
	"net"
	"sync"
	"testing"
)

func TestNewSession(t *testing.T) {
	type args struct {
		conn   net.Conn
		packer Packer
	}
	tests := []struct {
		name string
		args args
		want *Session
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewSession(tt.args.conn, tt.args.packer), "NewSession(%v, %v)", tt.args.conn, tt.args.packer)
		})
	}
}

func TestNewSessionManager(t *testing.T) {
	tests := []struct {
		name string
		want *SessionManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewSessionManager(), "NewSessionManager()")
		})
	}
}

func TestSessionManager_AddSession(t *testing.T) {
	type fields struct {
		sessionMap map[int32]*Session
		NextID     int32
		lock       sync.Mutex
	}
	type args struct {
		session *Session
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &SessionManager{
				sessionMap: tt.fields.sessionMap,
				NextID:     tt.fields.NextID,
				lock:       tt.fields.lock,
			}
			sm.AddSession(tt.args.session)
		})
	}
}

func TestSessionManager_RemoveSession(t *testing.T) {
	type fields struct {
		sessionMap map[int32]*Session
		NextID     int32
		lock       sync.Mutex
	}
	type args struct {
		session *Session
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &SessionManager{
				sessionMap: tt.fields.sessionMap,
				NextID:     tt.fields.NextID,
				lock:       tt.fields.lock,
			}
			sm.RemoveSession(tt.args.session)
		})
	}
}

func TestSession_Send(t *testing.T) {
	type fields struct {
		id            int32
		conn          net.Conn
		packer        Packer
		responseQueue chan RouteContext
		requestQueue  chan struct{}
		closed        chan struct{}
	}
	type args struct {
		rt *RouteContext
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantOk bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Session{
				id:            tt.fields.id,
				conn:          tt.fields.conn,
				packer:        tt.fields.packer,
				responseQueue: tt.fields.responseQueue,
				requestQueue:  tt.fields.requestQueue,
				closed:        tt.fields.closed,
			}
			assert.Equalf(t, tt.wantOk, s.Send(tt.args.rt), "Send(%v)", tt.args.rt)
		})
	}
}

func TestSession_handleReq(t *testing.T) {
	type fields struct {
		id            int32
		conn          net.Conn
		packer        Packer
		responseQueue chan RouteContext
		requestQueue  chan struct{}
		closed        chan struct{}
	}
	type args struct {
		router map[uint32]HandlerFunc
		msg    *Message
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Session{
				id:            tt.fields.id,
				conn:          tt.fields.conn,
				packer:        tt.fields.packer,
				responseQueue: tt.fields.responseQueue,
				requestQueue:  tt.fields.requestQueue,
				closed:        tt.fields.closed,
			}
			s.handleReq(tt.args.router, tt.args.msg)
		})
	}
}

func TestSession_readInbound(t *testing.T) {
	type fields struct {
		id            int32
		conn          net.Conn
		packer        Packer
		responseQueue chan RouteContext
		requestQueue  chan struct{}
		closed        chan struct{}
	}
	type args struct {
		router map[uint32]HandlerFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Session{
				id:            tt.fields.id,
				conn:          tt.fields.conn,
				packer:        tt.fields.packer,
				responseQueue: tt.fields.responseQueue,
				requestQueue:  tt.fields.requestQueue,
				closed:        tt.fields.closed,
			}
			s.readInbound(tt.args.router)
		})
	}
}

func TestSession_writeOutbound(t *testing.T) {
	type fields struct {
		id            int32
		conn          net.Conn
		packer        Packer
		responseQueue chan RouteContext
		requestQueue  chan struct{}
		closed        chan struct{}
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Session{
				id:            tt.fields.id,
				conn:          tt.fields.conn,
				packer:        tt.fields.packer,
				responseQueue: tt.fields.responseQueue,
				requestQueue:  tt.fields.requestQueue,
				closed:        tt.fields.closed,
			}
			s.writeOutbound()
		})
	}
}
