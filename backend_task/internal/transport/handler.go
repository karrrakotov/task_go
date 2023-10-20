package transport

import "net/http"

var action bool

type ClientHandler interface {
	Init(router *http.ServeMux)
}

type ParserHandler interface {
	Init(router *http.ServeMux)
}
