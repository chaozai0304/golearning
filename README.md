# golearning

Go 语言学习仓库，记录了从零开始的 Go 学习与练习代码。

## 目录说明

```
golearning
├── README.md                     # 项目说明（本文件）
├── Jenkinsfile                   # Jenkins CI 流水线配置
├── main/                         # 根目录下预留的 main 目录（当前为空）
└── src/
    └── main/
        ├── DockerFile            # 用于构建 testOne 应用的 Dockerfile
        ├── k8s.yml               # Kubernetes 部署配置
        ├── example.txt           # 文件操作示例生成的文本文件
        ├── 入口.html              # 炫酷工作台首页（HTML 演示页面）
        ├── testOne.go            # 并发示例：HTTP 服务 + goroutine（CPU 密集型任务 / 网络请求 / 文件操作）
        ├── day01/
        │   ├── echo1/
        │   │   └── echoTest1.go  # 命令行参数回显：使用 for 循环拼接
        │   ├── echo2/
        │   │   └── echoTest2.go  # 命令行参数回显：使用 range 遍历
        │   └── echo3/
        │       └── echoTest3.go  # 命令行参数回显：使用 strings.Join
        ├── day02/
        │   └── goconnectpg.go    # 连接 PostgreSQL 并执行查询、创建目录、复制文件
        └── day03/
            └── gofileserver.go   # 文件服务器：支持文件上传与下载
```

## 目录内容说明

| 路径 | 说明 |
| --- | --- |
| `src/main/day01/echo1/echoTest1.go` | 通过 `for` 循环和字符串拼接实现命令行参数回显（echo） |
| `src/main/day01/echo2/echoTest2.go` | 通过 `range` 遍历参数实现命令行参数回显（echo） |
| `src/main/day01/echo3/echoTest3.go` | 通过 `strings.Join` 实现命令行参数回显（echo） |
| `src/main/day02/goconnectpg.go` | 连接 PostgreSQL 数据库，按查询结果创建目录并复制文件 |
| `src/main/day03/gofileserver.go` | 基于 `net/http` 的文件服务器，支持文件上传与下载 |
| `src/main/testOne.go` | 并发与 HTTP 服务示例，使用 `sync.WaitGroup` 编排多个 goroutine |
| `src/main/DockerFile` | 将 `testOne.go` 编译为二进制并构建轻量级运行镜像 |
| `src/main/k8s.yml` | Kubernetes 部署清单 |
| `src/main/入口.html` | 炫酷工作台首页演示页面 |
| `Jenkinsfile` | Jenkins 流水线配置，用于拉取代码并执行基础命令 |
