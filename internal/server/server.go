package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/goodjobtech/candado/internal/state"
	"github.com/gorilla/mux"
)

type Server struct {
	router    *mux.Router
	db        state.Locker
	server    *http.Server
	Host      string
	Port      string
	Addr      string
	MustClose bool
}

func New(locker state.Locker) *Server {
	return &Server{
		router: mux.NewRouter(),
		db:     locker,
	}
}

func (s *Server) Serve() error {
	s.routes()

	s.server = &http.Server{
		Handler:      s.router,
		Addr:         s.Host + ":" + s.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Starting server at %s\n", s.Addr)
	if err := s.server.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			return err
		}
	}

	return nil
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.ShutdownWithContext(ctx)
}

func (s *Server) ShutdownWithContext(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		if s.MustClose {
			s.server.Close()
		}
		return err
	}

	return nil
}
