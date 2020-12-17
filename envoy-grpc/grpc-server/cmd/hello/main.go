package main

import (
	"context"
	"log"
	"os"

	"grpce/internal/config"
	"grpce/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Printf("failed to start: %v", err)
		os.Exit(1)
	}

	s := server.NewServer(cfg)
	err = s.Start(context.Background())
	if err != nil {
		log.Printf("failed to start: %v", err)
		os.Exit(1)
	}

	// TODO: add graceful shutdown
}
