package server

type Context interface {
	Response() *Message
	Request() *Message
	Session() *Session
}

type RouteContext struct {
	session *Session
}

func (r RouteContext) Response() *Message {
	//TODO implement me
	panic("implement me")
}

func (r RouteContext) Request() *Message {
	//TODO implement me
	panic("implement me")
}

func (r RouteContext) Session() *Session {

	return r.session
}

func (r RouteContext) SetSession(sess *Session) {
	r.session = sess
}
