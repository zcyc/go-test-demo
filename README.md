# go-test-demo

Go 使用 Ginkgo、Gomega、GoMock 做单元测试的例子。

## 介绍

- Ginkgo 是一个 Go 语言的 BDD 测试框架，它提供了一个类似于 RSpec 的 DSL，让开发者可以用更自然的语言来描述测试用例。 Ginkgo
  还可以帮助开发者组织测试用例，并提供一些方便的功能，例如并行测试、测试结果报告等等。
- Gomega 是一个 Go 语言的断言库，它提供了一系列的匹配器，让开发者可以用更自然、更易读的方式来编写断言语句。 Gomega 的匹配器比
  Go 标准库中的 assert 函数更强大、更灵活，也更易于阅读。
- GoMock 是一个 Go 语言的模拟框架，可以帮助开发者在测试过程中创建模拟对象，从而隔离代码，更有效地测试代码的逻辑。


## 安装

```bash
go install github.com/onsi/ginkgo/v2/ginkgo
go install go.uber.org/mock/mockgen@latest
```

## 生成 mock

在需要 mock 的 interface 所在目录执行生成命令，记得修改命令中的 xxx，**每次修改 interface 都需要执行**。

```bash
mockgen -source=xxx.go -destination=mock_xxx.go -package=xxx
```
## 批量生成 mock

在需要 mock 的 interface 上添加注释，记得修改命令中的 xxx。

```bash
//go:generate mockgen -source=xxx.go -destination=mock_xxx.go -package=xxx
```

在项目根目录执行生成命令，**每次修改 interface 都需要执行**。

```bash
go generate ./...
```

## 生成 suite

在需要测试的文件所在目录执行。必须生成 suite，否则 specs 无法运行。

```bash
ginkgo bootstrap
```

## 生成 specs

记得修改命令中的 xxx，使用 [Ginkgo](https://plugins.jetbrains.com/plugin/17554-ginkgo) 插件要将 suite 和 specs 合并为一个文件。

```bash
ginkgo generate xxx
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

