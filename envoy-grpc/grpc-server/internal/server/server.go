package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"grpce/gen/helloworld"
	"grpce/internal/config"
	"grpce/internal/server/rpc"

	"google.golang.org/grpc"
)

type Server struct {
	cfg        config.Config
	grpcServer *grpc.Server
}

func NewServer(cfg config.Config) Server {
	grpcServer := grpc.NewServer()
	helloworld.RegisterGreeterService(grpcServer, helloworld.NewGreeterService(&rpc.RPC{}))

	return Server{
		cfg:        cfg,
		grpcServer: grpcServer,
	}
}

func (s Server) Start(_ context.Context) error {
	log.Printf("starting GRPC on %q tcp...", s.cfg.Address)

	lis, err := net.Listen("tcp", s.cfg.Address)
	if err != nil {
		return fmt.Errorf("failed to listen tcp: %w", err)
	}

	if err := s.grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve grpc server: %w", err)
	}

	return nil
}

func (s Server) Stop() error {
	s.grpcServer.Stop()

	return nil
}
