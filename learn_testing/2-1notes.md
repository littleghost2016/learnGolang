# 2-1

## 1. 每个test文件都必须`import "testing"`。

```go
package main

import (
    "testing"
)

func TestPring(t *testing.T) {
    res := Print1to20()
    if res != 210 {
        t.Errorf("Return value not valid!")
    }
}
```

- `test`文件下每一个test case都必须是`TestXXX`的命名方式，否则go test会跳过测试不执行
- `test case`的入参必须是`t *testing.T`或者`b *testing.B`(后者用于跑`Benchmark测试性能`)
- `t.Errorf()`函数会打印出错误信息，并且退出当前test case
- `t.SkipNow()`为跳过当前test case继续执行下一个test，且必须在当前case的第一行，否则不起作用

```go
func TestMain(m *testing.M) {
    fmt.Println("test main function")
    m.Run()
}
```

- 使用`test main`作为初始化test，并且使用`m.Run()`调用其他tests可以完成一些需要初始化操作的testing，比如数据库连接、文件打开、REST服务登录等
- 如果没有在`TestMain`中执行`m.Run()`，则除了`TestMain`以外的其他test都不会被执行。*即，当文件里没有TestMain函数时，其他test case都会被执行，当文件里有TestMain却没有m.Run()，其他的test case都不会被执行。*

## 2. 交叉编译

> 作者：磐石区
> 来源：CSDN
> 原文：https://blog.csdn.net/panshiqu/article/details/53788067
> 版权声明：本文为博主原创文章，转载请附上博文链接！

### Mac 下编译 Linux 和 Windows 64位可执行程序

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

### Linux 下编译 Mac 和 Windows 64位可执行程序

```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

### Windows 下编译 Mac 和 Linux 64位可执行程序**

```bash
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build main.go

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go
```

- GOOS：目标平台的操作系统（darwin、freebsd、linux、windows）
- GOARCH：目标平台的体系架构（386、amd64、arm）
- 交叉编译不支持 CGO 所以要禁用它
- 上面的命令编译 64 位可执行程序，当然也可以使用 386 编译 32 位可执行程序

## 3.Benchmark

```go
func BenchmarkAll(b *testing.B)  {
    for n := 0; n < b.N; n++ {
        Print1to20()
    }
}
```

- 函数名要以`Benchmark`开头
- 同样受到`TestMain`的限制
- 一般的命令行命令为`go test -bench=.`，此时只会跑以Benchmark开头的函数
- b.N在每次跑的过程中都会变化
- 当测试需要传参的Benchmark时，需要被测函数总能在一定时间后达到running time的稳态