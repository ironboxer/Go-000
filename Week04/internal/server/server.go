package server

import (
	"context"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	*grpc.Server
	address string
}

func NewServer(address string) *Server {
	srv := grpc.NewServer()
	return &Server{Server: srv, address: address}
}

func (s *Server) Start(ctx context.Context) error {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	// wait until complete
	go func() {
		<-ctx.Done()
		s.GracefulStop()
	}()
	err = s.Serve(listener)
	return err
}
