package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/goodjobtech/candado/internal/database/memory"
	"github.com/goodjobtech/candado/internal/database/redis"
	"github.com/goodjobtech/candado/internal/server"
	"github.com/goodjobtech/candado/internal/state"
)

func main() {
	var db state.Locker

	switch os.Getenv("CANDADO_DB_TYPE") {
	case "memory":
		db = memory.New()
	default:
		db = redis.New()
	}

	server := server.New(db)
	errCh := make(chan error)

	go func() {
		if err := server.Serve(); err != nil {
			errCh <- err
		}
		close(errCh)
	}()

	host, ok := os.LookupEnv("CANDADO_SERVER_HOST")
	if ok {
		server.Host = host
	}

	port, ok := os.LookupEnv("CANDADO_SERVER_PORT")
	if ok {
		server.Port = port
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	select {
	case err, ok := <-errCh:
		if ok {
			log.Println("server error:", err)
		}
	case sig := <-sigCh:
		log.Printf("Signal %s received\n", sig)

		os.Exit(0)
		log.Println("server shutdown")
	}
}
