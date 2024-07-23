# go-test-demo

A simple demo for testing in Go.

## Dependencies

- ginkgo
  Ginkgo 是一个 Go 语言的 BDD 测试框架，它提供了一个类似于 RSpec 的 DSL，让开发者可以用更自然的语言来描述测试用例。 Ginkgo
  还可以帮助开发者组织测试用例，并提供一些方便的功能，例如并行测试、测试结果报告等等。
- gomega
  Gomega 是一个 Go 语言的断言库，它提供了一系列的匹配器，让开发者可以用更自然、更易读的方式来编写断言语句。 Gomega 的匹配器比
  Go 标准库中的 assert 函数更强大、更灵活，也更易于阅读。
- gomock
  GoMock 是一个 Go 语言的模拟框架，可以帮助开发者在测试过程中创建模拟对象，从而隔离代码，更有效地测试代码的逻辑。

## How to run

```shell
go test ./...
```

## How to run with coverage

```shell
go test -cover ./...
```

## How to run with coverage and generate coverage report

```shell
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```