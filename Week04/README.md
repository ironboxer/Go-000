按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。

以上作业，要求提交到Github上面，Week04作业提交地址：
https://github.com/Go-000/Go-000/issues/76

```bash
➜  Week04 git:(main) tree
.
├── Makefile
├── README.md
├── api
│   └── tag
│       └── v1
│           ├── tag.pb.go
│           ├── tag.proto
│           └── tag_grpc.pb.go
├── bin
│   └── server
├── cmd
│   └── server
│       ├── main.go
│       ├── wire.go
│       └── wire_gen.go
├── conf
│   ├── setting.go
│   └── setting.ini
├── go.mod
├── go.sum
└── internal
    ├── biz
    │   └── tag.go
    ├── data
    │   └── tag.go
    ├── pkg
    │   └── grpc
    │       └── server.go
    ├── server
    │   └── server.go
    └── service
        └── tag.go

14 directories, 18 files
```

编译
```bash
make build
```

运行
```bash
make run
```

清理
```bash
make clean
```