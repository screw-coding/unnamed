package server

type Context interface {
	Response() *Message
	Request() *Message
	Session() *Session
}

type RouteContext struct {
	session     *Session
	reqMsg      *Message
	responseMsg *Message
}

func (r RouteContext) Response() *Message {
	return r.responseMsg
}

func (r RouteContext) Request() *Message {
	return r.reqMsg
}

func (r RouteContext) Session() *Session {
	return r.session
}

func (r RouteContext) SetSession(sess *Session) {
	r.session = sess
}
