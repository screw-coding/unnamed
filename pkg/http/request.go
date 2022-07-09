package http

import (
	"context"
	"io"
)

type Request struct {
	Method        string //GET
	URL           *URL   //http://www.google.com
	Proto         string //"HTTP/1.1"
	ProtoMajor    int    //1
	ProtoMinor    int    //1
	Header        Header //map[string][]string
	Body          io.ReadCloser
	Response      *response
	RequestURI    string
	ContentLength int64
	ctx           *context.Context
	RemoteAddr    string
	Trailer       Header
	Host          string
	Close         bool
}

func NewRequest(method, url string) *Request {
	if method == "" {
		method = "GET"
	}

	return &Request{
		Method:     method,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
	}
}
