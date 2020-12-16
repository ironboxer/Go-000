package grpc


import "context"
import "log"
import "net"
import "google.golang.org/grpc"



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
    log.Printf("grpc server start: %s\n", s.address)

    go func() {
        <-ctx.Done()
        s.GracefulStop()
        log.Println("grpc server gracefully stopped")
    }()

    return s.Serve(listener)
}

