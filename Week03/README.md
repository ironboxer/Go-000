### Week03 作业题目：
基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。


对于http server,通过启动两个http server,触发重复绑定端口的错误,导致程序退出

```go
	g.Go(func() error {
		return CreateHttpServer(cancel)
	})
	// 通过重复绑定统一端口来触发报错 进而退出
	g.Go(func() error {
		return CreateHttpServer(cancel)
	})
```


注册了两个信号量```SIGINT```, ```SIGQUIT```, 分别使用 kill -2 pid, kill -3 pid 即可杀死进程
```go
signal.Notify(sig, syscall.SIGINT, syscall.SIGQUIT)
```
