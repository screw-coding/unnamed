package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRouteContext_Request(t *testing.T) {
	type fields struct {
		session     *Session
		reqMsg      *Message
		responseMsg *Message
	}
	tests := []struct {
		name   string
		fields fields
		want   *Message
	}{
		// Add test cases.
		{name: "test", fields: fields{session: &Session{}, reqMsg: &Message{Id: 1, Data: []byte{1}}}, want: &Message{Id: 1, Data: []byte{1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RouteContext{
				session:     tt.fields.session,
				reqMsg:      tt.fields.reqMsg,
				responseMsg: tt.fields.responseMsg,
			}
			assert.Equalf(t, tt.want, r.Request(), "Request()")
		})
	}
}

func TestRouteContext_Response(t *testing.T) {
	type fields struct {
		session     *Session
		reqMsg      *Message
		responseMsg *Message
	}
	tests := []struct {
		name   string
		fields fields
		want   *Message
	}{
		// Add test cases.
		{name: "test", fields: fields{session: &Session{}, responseMsg: &Message{Id: 1, Data: []byte{1}}}, want: &Message{Id: 1, Data: []byte{1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RouteContext{
				session:     tt.fields.session,
				reqMsg:      tt.fields.reqMsg,
				responseMsg: tt.fields.responseMsg,
			}
			assert.Equalf(t, tt.want, r.Response(), "Response()")
		})
	}
}

func TestRouteContext_Session(t *testing.T) {
	type fields struct {
		session     *Session
		reqMsg      *Message
		responseMsg *Message
	}
	tests := []struct {
		name   string
		fields fields
		want   *Session
	}{
		// Add test cases.
		{name: "test", fields: fields{session: &Session{id: 1, packer: NewDefaultPacker()}}, want: &Session{id: 1, packer: NewDefaultPacker()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RouteContext{
				session:     tt.fields.session,
				reqMsg:      tt.fields.reqMsg,
				responseMsg: tt.fields.responseMsg,
			}
			assert.Equalf(t, tt.want, r.Session(), "Session()")
		})
	}
}

func TestRouteContext_SetSession(t *testing.T) {
	type fields struct {
		session     *Session
		reqMsg      *Message
		responseMsg *Message
	}
	type args struct {
		sess *Session
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// Add test cases.
		{name: "test", fields: fields{session: &Session{id: 1, packer: NewDefaultPacker()}}, args: args{sess: &Session{id: 1, packer: NewDefaultPacker()}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RouteContext{
				session:     tt.fields.session,
				reqMsg:      tt.fields.reqMsg,
				responseMsg: tt.fields.responseMsg,
			}
			r.SetSession(tt.args.sess)
		})
	}
}
