package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	log.Printf("Pid: %d\n", os.Getpid())
	ctx := context.Background()
	g, cancel := errgroup.WithContext(ctx)
	// 启动后通过Ctrl-C(或Kill pid)触发SIGINT使进程退出
	g.Go(func() error {
		return RegisterSignalHandler(cancel)
	})
	g.Go(func() error {
		return CreateHttpServer(cancel)
	})
	// 通过重复绑定统一端口来触发报错 进而退出
	// g.Go(func() error {
	// 	return CreateHttpServer(cancel)
	// })
	if err := g.Wait(); err != nil {
		log.Println(err)
	}
	log.Println("Exit!")
}

func RegisterSignalHandler(ctx context.Context) error {
	sig := make(chan os.Signal)
	// 注册了SIGINT SIGUQIT信号
	// 通过 kill -2 $(pid) kill -3 $(pid) 即可停止进程
	signal.Notify(sig, syscall.SIGINT, syscall.SIGQUIT)
	for {
		select {
		case sig := <-sig:
			return fmt.Errorf("Signal Interrupted! %+v", sig)
		case <-ctx.Done():
			return nil
		}
	}
}

func CreateHttpServer(ctx context.Context) error {
	addr := ":8034"
	handler := &HttpHandler{}
	server := http.Server{
		Addr:    addr,
		Handler: handler,
	}
	go func(ctx context.Context) {
		<-ctx.Done()
		server.Shutdown(context.Background())
	}(ctx)
	err := server.ListenAndServe()
	return err
}

type HttpHandler struct {
}

func (h *HttpHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	time.Sleep(time.Second * 1)
	method := req.Method
	url := req.URL
	addr := req.RemoteAddr
	fmt.Fprintf(resp, "URL: %v\n", url)
	log.Printf("[%v] %v %v\n", addr, method, url)
}
