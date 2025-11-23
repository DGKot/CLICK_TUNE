package rest

import (
	"click_tune/internal/service"
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

type ServerDeps struct {
	Service      *service.Service
	Host         string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func NewServer(deps ServerDeps) *Server {
	addr := net.JoinHostPort(deps.Host, deps.Port)
	server := &http.Server{
		Addr:         addr,
		ReadTimeout:  deps.ReadTimeout,
		WriteTimeout: deps.WriteTimeout,
		IdleTimeout:  deps.IdleTimeout,
		Handler:      NewHandler(deps.Service).Routes(),
	}
	return &Server{
		server: server,
	}
}

func (s *Server) Start() error {
	fmt.Println("server start")
	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	fmt.Println("server close")

	return s.server.Shutdown(ctx)
}

// For test.
func (s *Server) Handler() http.Handler {
	return s.server.Handler
}
