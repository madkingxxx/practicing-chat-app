package httpserver

import (
	"context"
	"log"
	"net/http"
	"time"
)

const (
	_httpDefaultReadTimeout                = time.Second * 5
	_httpDefaultWriteTimeout               = time.Second * 5
	_defaultAddr             string        = ":8090"
	_defaultShutdownTimeout  time.Duration = time.Second * 3
)

// Server ...
type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

// New ...
func New(handler http.Handler, opts ...Option) *Server {
	httpServer := &http.Server{
		Addr:         _defaultAddr,
		Handler:      handler,
		ReadTimeout:  _httpDefaultReadTimeout,
		WriteTimeout: _httpDefaultWriteTimeout,
	}

	s := &Server{
		server:          httpServer,
		notify:          make(chan error),
		shutdownTimeout: _defaultShutdownTimeout,
	}

	// Custom options
	for _, opt := range opts {
		opt(s)
	}
	s.start()

	return s
}

// Start ...
func (s *Server) start() {
	go func() {
		log.Print("server started at: ", _defaultAddr)
		err := s.server.ListenAndServe()
		if err != nil {
			s.notify <- err
		}
	}()
}

// Shtudown ...
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()
	return s.server.Shutdown(ctx)
}

func (s *Server) Notify() <-chan error {
	return s.notify
}
