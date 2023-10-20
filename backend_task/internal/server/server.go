package server

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,          // 1 MB
		ReadTimeout:    15 * time.Second, // 15 SEC
		WriteTimeout:   15 * time.Second, // 15 SEC
	}

	return s.httpServer.ListenAndServe()
}
