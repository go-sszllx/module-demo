# Go Module Demo

该Project演示了如何使用Go的module功能

## Getting Started

根据如下的步骤,可以创建一个简单的应用Go Module的工程

### Prerequisites

需安装Go v1.14版本
[下载](https://golang.org/dl/)

## 执行步骤
### 创建工程
```
├── main.go
├── Makefile
├── README.md
└── utils
    ├── http
    │   └── http.go
    └── utils.go
```
main.go中引用了utils.go,utils.go引用了http文件夹下的http.go

### 创建module
```
go mod init test
```

test为模块名

### Build
```
go build
```

### go mod命令列表

| 命令 | 作用 |
| ------| ------ |
| go mod init | 生成 go.mod 文件 |
| go mod download	 | 下载 go.mod 文件中指明的所有依赖 |
|go mod tidy	| 整理现有的依赖 |
|go mod graph	| 查看现有的依赖结构 |
|go mod edit	|编辑 go.mod 文件
|go mod vendor	|导出项目所有的依赖到vendor目录
|go mod verify	|校验一个模块是否被篡改过
|go mod why	|查看为什么需要依赖某模块
