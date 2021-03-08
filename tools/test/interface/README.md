
```markdown
go tool compile -S -N -l main.go
```
参数含义：
* -S：打印汇编列表
* -N：禁用编译器优化
* -l：禁用内联编译

使用结构体实现接口时：
```markdown
wanghuan@wanghuans-MBP interface % go test -gcflags=-N -benchmem -test.count=3 -test.cpu=1 -test.benchtime=1s -bench=.
goos: darwin
goarch: amd64
pkg: github.com/xdhuxc/go-study-notes/tools/test/interface
BenchmarkDirectCall             508449457                2.31 ns/op            0 B/op          0 allocs/op
BenchmarkDirectCall             506028432                2.29 ns/op            0 B/op          0 allocs/op
BenchmarkDirectCall             524361830                2.30 ns/op            0 B/op          0 allocs/op
BenchmarkDynamicDispatch        246200607                4.84 ns/op            0 B/op          0 allocs/op
BenchmarkDynamicDispatch        250580403                4.87 ns/op            0 B/op          0 allocs/op
BenchmarkDynamicDispatch        246231108                4.85 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/xdhuxc/go-study-notes/tools/test/interface   9.672s
```

使用结构体指针实现接口时：
```markdown
wanghuan@wanghuans-MBP interface % go test -gcflags=-N -benchmem -test.count=3 -test.cpu=1 -test.benchtime=1s -bench=.
goos: darwin
goarch: amd64
pkg: github.com/xdhuxc/go-study-notes/tools/test/interface
BenchmarkDirectCall             462147699                2.36 ns/op            0 B/op          0 allocs/op
BenchmarkDirectCall             475629063                2.30 ns/op            0 B/op          0 allocs/op
BenchmarkDirectCall             515122164                2.30 ns/op            0 B/op          0 allocs/op
BenchmarkDynamicDispatch        466204557                2.59 ns/op            0 B/op          0 allocs/op
BenchmarkDynamicDispatch        462035305                2.61 ns/op            0 B/op          0 allocs/op
BenchmarkDynamicDispatch        458656573                2.60 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/xdhuxc/go-study-notes/tools/test/interface   9.058s
```
