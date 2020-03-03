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
