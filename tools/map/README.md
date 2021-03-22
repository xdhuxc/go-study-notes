### map 


测试结果为：
```markdown
wanghuan@wanghuans-MacBook-Pro map % go run main.go 1
With map[int32]*int32, GC took 656.157603ms
wanghuan@wanghuans-MacBook-Pro map % go run main.go 2
With map[int32]int32, GC took 12.819972ms
wanghuan@wanghuans-MacBook-Pro map % go run main.go 3
With map shards ([]map[int32]int32), GC took 2.039869ms
wanghuan@wanghuans-MacBook-Pro map % go run main.go 4
With a plain slice ([]main.t), GC took 188.96µs
```

基准测试情况如下：
```markdown
wanghuan@wanghuans-MacBook-Pro map % go test -bench=. -benchmem -run=none benchmark_test.go
goos: darwin
goarch: amd64
BenchmarkPointer-4       4916869               223 ns/op              58 B/op          1 allocs/op
BenchmarkInt32-4         9333214               174 ns/op              43 B/op          0 allocs/op
BenchmarkShared-4        9730831               164 ns/op              33 B/op          0 allocs/op
BenchmarkSlice-4        100000000               31.2 ns/op            49 B/op          0 allocs/op
PASS
ok      command-line-arguments  8.296s
```
使用数组结构相当高效，结合 GC 时间，多 map 读写优于单 map，使用指针会导致 GC 时间过长，并且每次都会分配内存，是最耗费资源的方式。

