### 基准测试

1、基准测试的代码文件必须以 _test.go 结尾
2、基准测试的函数必须以 Benchmark 开头，必须是可导出的
3、基准测试函数必须接受一个指向 Benchmark 类型的指针作为唯一参数
4、基准测试函数不能有返回值
5、b.ResetTimer 是重置计时器，这样可以避免 for 循环之前的初始化代码的干扰
6、最后的 for 循环很重要，被测试的代码要放到循环里
7、b.N 是基准测试框架提供的，表示循环的次数，因为需要反复调用测试的代码，才可以评估性能

使用如下命令运行基准测试：
```markdown
wanghuan@wanghuans-MacBook-Pro benchmark % go test -bench=. -run=none
goos: darwin
goarch: amd64
pkg: github.com/xdhuxc/go-study-notes/tools/test/benchmark
BenchmarkSprintf-4      11860180                93.4 ns/op
PASS
ok      github.com/xdhuxc/go-study-notes/tools/test/benchmark   1.342s
```

运行基准测试也是使用 `go test` 命令，但是需要加上 `-bench=` 参数，接收一个表达式以匹配基准测试的函数。`.` 表示运行所有的基准测试

默认情况下，`go test` 会运行单元测试，为了防止单元测试的输出影响我们查看基准测试的结果，可以使用 `-run=` 匹配一个从来没有的单元测试方法，过滤掉单元测试的输出。
这里使用 `none`，因为我们无需创建这个名字的单元测试方法。

函数名称后的 `-4` 表示运行时对应的 `GOMAXPROCS` 的值；`11860180` 表示运行 for 循环的测试，也就是调用被测试代码的次数；`93.4 ns/op` 表示每次需要耗时 93.4 纳秒。

如果想让测试运行的时间更长，可以通过 -benchtime 指定，比如 3 秒。
```markdown
wanghuan@wanghuans-MacBook-Pro benchmark % go test -bench=. -benchtime=3s -run=none
goos: darwin
goarch: amd64
pkg: github.com/xdhuxc/go-study-notes/tools/test/benchmark
BenchmarkSprintf-4      36750073                92.0 ns/op
PASS
ok      github.com/xdhuxc/go-study-notes/tools/test/benchmark   4.559s
```
可以看到，在加长测试时间后，每次执行代码耗费的时间并没有明显的变化。


### 比较

比较 fmt.Sprintf()，strconv.FormatInt()，strconv.Itoa() 的性能

```markdown
wanghuan@wanghuans-MacBook-Pro benchmark % go test -bench=. -run=none              
goos: darwin
goarch: amd64
pkg: github.com/xdhuxc/go-study-notes/tools/test/benchmark
BenchmarkSprintf-4      12029704                92.9 ns/op
BenchmarkFormat-4       270879309                4.44 ns/op
BenchmarkItoa-4         252876846                4.71 ns/op
PASS
ok      github.com/xdhuxc/go-study-notes/tools/test/benchmark   4.632s
```

从结果可以看出，strconv.FormatInt() 函数是最快的，其次是 strconv.Itoa()，最慢的是 fmt.Sprintf()。

但是差距居然达到了 20 倍左右，那么这个差距为何如此之大呢？我们再通过 -benchmem 找到原因，-benchmem 可以提供每次操作分配内存的次数，以及每次操作分配的字节数。

```markdown
wanghuan@wanghuans-MacBook-Pro benchmark % go test -bench=. -benchmem -run=none
goos: darwin
goarch: amd64
pkg: github.com/xdhuxc/go-study-notes/tools/test/benchmark
BenchmarkSprintf-4      12470200                91.9 ns/op            16 B/op          2 allocs/op
BenchmarkFormat-4       264188895                4.48 ns/op            0 B/op          0 allocs/op
BenchmarkItoa-4         249210574                4.84 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/xdhuxc/go-study-notes/tools/test/benchmark   4.674s
```
从结果可以看出，性能高的 strconv.FormatInt() 和 strconv.Itoa()，每次操作都不进行内存分配，而 fmt.Sprintf() 要分配两次；
性能高的 strconv.FormatInt() 和 strconv.Itoa()，每次操作不分配内存，而慢的 fmt.Sprintf() 每次分配内存 16 个字节。
由此，我们就知道 fmt.Sprintf() 为什么这么慢了，每次操作都进行了内存分配和内存占用太高。

### 其他

在开发过程中，我们经常会遇到遍历数组的情况，对于最基本的遍历，在终止条件的位置上怎么写？很多时候，我们还是需要注意下的。

对于如下所示的两种基本遍历：
```markdown
for i := 0; i < len(x); i++ {
    _ = i
}
```
和
```markdown
len := len(x)
for i := 0; i < len; i++ {
    _ = i
}
```
在运行上，有区别吗？

为了验证这个问题，编写如下基准测试：
```markdown
package benchmark

import (
	"math/rand"
	"testing"
)

const a = 10^10
var x [a]int
func init() {
	for i := 0; i < len(x); i++ {
		x[i] = rand.Intn(a)
	}
}

func length() {
	for i := 0; i < len(x); i++ {
		_ = i
	}
}

func lengthOne() {
	length := len(x)
	for i := 0; i < length; i++ {
		_ = i
	}
}

func BenchmarkLen(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		length()
	}
}

func BenchmarkLenOne(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lengthOne()
	}
}
```
使用如下命令运行基准测试：
```markdown
go test -bench=. -benchmem -run=none length_test.go
```
结果如下所示：
```markdown
goos: darwin
goarch: amd64
BenchmarkLen-4          835416472                1.41 ns/op            0 B/op          0 allocs/op
BenchmarkLenOne-4       853285748                1.40 ns/op            0 B/op          0 allocs/op
PASS
ok      command-line-arguments  2.787s
```
多次运行，结果基本类似。

所以，看起来，这两种写法不会有性能方面的差别。

实际上，我们知道 golang 的数组和切片，在底层存储时，都会带有长度和容量字段，想必 len() 直接返回了数组或切片的长度。

在 [The Go Programming Language Specification](https://golang.org/ref/spec#Length_and_capacity) 中，关于 len() 函数有如下说明：
```markdown
The expression len(s) is constant if s is a string constant. The expressions len(s) and cap(s) are constants if the type of s is an array or pointer to an array and the expression s does not contain channel receives or (non-constant) function calls; in this case s is not evaluated. Otherwise, invocations of len and cap are not constant and s is evaluated.
```
翻译过来就是：
如果 s 是字符串常量，那么表达式 len(s) 就是常数。如果 s 的类型是数组或指向数组的指针，并且表达式 s 不包含 channel 接收或（非常量的）函数调用，那么表达式 len(s) 和 cap(s) 是常数。
在这种情况下，不会去计算 s。否则，len 和 cap 的调用不是常数，会对 s 进行计算。












