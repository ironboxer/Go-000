package main

import (
	pb "Week04/api/tag/v1"
	"Week04/conf"
	"Week04/internal/pkg/grpc"
	"Week04/internal/service"
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	tagUsecase := InitTagUsecase()
	service := service.NewTagService(tagUsecase)

	s := grpc.NewServer(conf.ServerAddr)
	pb.RegisterTagServer(s, service)
	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return s.Start(ctx)
	})

	g.Go(func() error {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT, syscall.SIGQUIT)
		select {
		case sig := <-signals:
			log.Printf("Received Signal: %v\n", sig)
			cancel()
		case <-ctx.Done():
			return ctx.Err()
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Printf("Error: %v", err)
	}
}
