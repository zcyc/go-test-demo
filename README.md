# Go Test Demo

这个项目展示了 Go 单元测试的多种方式，包括 Testing、Testify、Ginkgo、GoMock 和 Mockery 的使用案例。

## 项目结构

```
├── pkg                 # 项目源码包
│   └── user            # 用户模块
│       ├── dao         # 数据访问层（使用内存存储）
│       ├── mocks       # 生成的 mock 文件
│       ├── service     # 业务逻辑层
│       └── tests       # 测试文件
├── go.mod              # Go 模块文件
├── go.sum              # Go 依赖版本记录
└── README.md           # 项目说明文档
```

## 技术栈

- **Testing**: Go 标准测试库
- **Testify**: 更丰富的断言和 mock 功能
- **Ginkgo**: BDD 风格测试框架
- **GoMock**: 自动生成 mock 对象
- **Mockery**: 自动生成 Testify mock 对象

## 安装依赖工具

```bash
go install github.com/onsi/ginkgo/v2/ginkgo
go install go.uber.org/mock/mockgen@latest
go install github.com/vektra/mockery/v2@latest
```

## Mock 文件生成

### 使用命令生成

GoMock:
```bash
mockgen -source=pkg/user/dao/struct.go -destination=pkg/user/mocks/gomock_struct.go -package=mocks
```

Mockery:
```bash
mockery --name UserDao --output pkg/user/mocks --filename testify_struct.go --outpkg mocks
```

### 使用注释自动生成

在接口定义文件中添加注释:

```go
//go:generate mockgen -source=struct.go -destination=../mocks/gomock_struct.go -package=mocks
//go:generate mockery --name UserDao --filename testify_struct.go --outpkg mocks --output ../mocks
```

然后执行:

```bash
go generate ./...
```

## 测试示例

本项目包含三种测试方式的示例:

1. **标准 testing 包**: 简单直接的测试方式
2. **Testify**: 使用 assert 提供更丰富的断言功能
3. **Ginkgo**: BDD 风格的测试，提供 Describe/It 语法和性能测量功能

## 运行测试

```bash
# 运行所有测试
go test ./...

# 显示详细输出
go test -v ./...

# 显示测试覆盖率
go test -v -cover ./...

# 生成覆盖率报告
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

## Ginkgo 使用指南

### 创建测试套件

```bash
cd pkg/user/tests
ginkgo bootstrap
```

### 生成测试文件

```bash
ginkgo generate user
```

## 相关文档

- [Testing Package](https://golang.org/pkg/testing/)
- [Testify](https://github.com/stretchr/testify)
- [Ginkgo](https://onsi.github.io/ginkgo/)
- [Gomega](https://onsi.github.io/gomega/)
- [GoMock](https://github.com/uber-go/mock)
- [Mockery](https://vektra.github.io/mockery/)

