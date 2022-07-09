package main

import (
	"github.com/screw-coding/http"
	"log"
)

func main() {
	server := new(http.Server)
	server.Addr = "127.0.0.1:12345"
	serverMux := &http.ServerMux{
		Handlers: make(map[string]http.Handler),
	}

	serverMux.Handle("/", new(HelloHandler))
	server.Handler = serverMux
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("listen error")
	}
}

type HelloHandler struct {
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
