# go-test-demo

Go 使用 Testing、Ginkgo、GoMock、Gomega、Testify、Mockery 做单元测试的例子。

## 介绍

- Ginkgo 是一个 Go 语言的 BDD 测试框架，它提供了一个类似于 RSpec 的 DSL，让开发者可以用更自然的语言来描述测试用例。
- GoMock 是一个 Go 语言的 Mock 生成器，可以帮助开发者创建 Ginkgo Mock 对象。
- Gomega 是一个 Go 语言的断言库，它提供了一系列的匹配器，让开发者可以用更自然、更易读的方式来编写断言语句。
- Testify 是一个 Go 语言测试框架，可以很好地与标准库配合使用。
- Mockery 是一个 Go 语言的 Mock 生成器，可以帮助开发者创建 Testify Mock 对象。

## 安装

```bash
go install github.com/onsi/ginkgo/v2/ginkgo
go install go.uber.org/mock/mockgen@latest
go install github.com/vektra/mockery/v2@v2.44.1
```

## 使用命令生成 mock 文件

在需要 mock 的 interface 所在目录执行生成命令，记得修改命令中的 xxx，**每次修改 interface 都需要执行**。

### GoMock
xxx 为 interface 所在的文件名称
```bash
mockgen -source=xxx.go -destination=mocks/gomock_struct.go -package=mocks
```
### Mockery
xxx 为 interface 名称
```bash
mockery --name xxx --filename testify_struct.go --outpkg mocks
```

## 使用注释批量生成 mock 文件

在需要 mock 的 interface 上添加注释，然后在项目根目录执行生成命令。**每次修改 interface 都需要执行**。

```bash
go generate ./...
```

### GoMock 注释
xxx 为 interface 所在的文件名称
```bash
//go:generate mockgen -source=xxx.go -destination=mocks/gomock_struct.go -package=mocks
```

### Mockery 注释
xxx 为 interface 名称
```bash
//go:generate mockery --name xxx --filename testify_struct.go --outpkg mocks
```

## 使用 Ginkgo 进行测试

1. 生成 suite

在需要测试的文件所在目录执行。必须生成 suite，否则 specs 无法运行。

```bash
ginkgo bootstrap
```

2. 生成 specs

记得修改命令中的 xxx，使用 [Ginkgo](https://plugins.jetbrains.com/plugin/17554-ginkgo) 插件要将 suite 和 specs 合并为一个文件。

```bash
ginkgo generate xxx
```

## 使用 Testify 进行测试
GoLand 可以一键生成测试文件，手动创建步骤如下：
1. 创建一个名为 *_test.go 的文件，例如 mypackage_test.go。
2. 在文件中定义测试函数，函数名必须以 Test 开头，例如 TestMyFunction。
3. 在测试函数中使用 testing.T 类型来进行断言和错误处理。
```go
package xx_test

import (
    "testing"
)

func TestMyFunction(t *testing.T) {
}
```

## 执行测试

```shell
go test ./...
```

## 执行测试并展示详细信息

```shell
go test -v ./...
```

## 执行测试并展示详细信息和覆盖率

```shell
go test -v -cover ./...
```

## 执行测试并导出覆盖率报告

```shell
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

## 文档

- [Ginkgo](https://onsi.github.io/ginkgo/#bootstrapping-a-suite)
- [Gomega](https://onsi.github.io/gomega/)
- [gomock](https://github.com/uber-go/mock)
- [Testify](https://github.com/stretchr/testify)
- [Mockery](https://vektra.github.io/mockery/latest)

