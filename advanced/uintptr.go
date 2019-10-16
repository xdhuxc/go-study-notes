package main

import (
	"fmt"
	"runtime"
	"time"
	"unsafe"
)

type data struct {
	x [1024 * 100]byte
}

func testUintptr() uintptr {
	p := &data{}
	return uintptr(unsafe.Pointer(p))
}

func mainUintptr() {
	/**
		uintptr 被 GC 当做普通整数对象，所引用的对象会被回收
	 */
	const N = 10000
	cache := new([N]uintptr)

	for i := 0; i < N; i++ {
		cache[i] = testUintptr()
		time.Sleep(time.Millisecond)
	}
}

func testUnsafePointer() unsafe.Pointer {
	/**
		合法的 unsafe.Pointer 被当做普通指针对待，确保对象不被回收
	 */
	p := &data{}
	return unsafe.Pointer(p)
}

func mainUnsafePointer() {
	const N = 10000
	cache := new([N]unsafe.Pointer)

	for i := 0; i < N; i++ {
		cache[i] = testUnsafePointer()
		time.Sleep(time.Millisecond)
	}
}

type data2 struct {
	x [1024 * 100]byte
	y int
}

func testUnsafePointer2() unsafe.Pointer {
	/**
		指向对象成员的 unsafe.Pointer，同样能确保对象不被回收
	 */
	d := data2{}
	return unsafe.Pointer(&d.y)
}

func mainUnsafePointer2() {
	const N = 10000
	cache := new([N]unsafe.Pointer)

	for i := 0; i < N; i++ {
		cache[i] = testUnsafePointer2()
		time.Sleep(time.Millisecond)
	}

}

type Data struct {
	d [1024 * 100]byte
	o *Data
}

func testPointer() {
	var a, b Data

	a.o = &b
	b.o = &a

	/**
		垃圾回收器能正确处理"指针循环引用"，但无法确定 Finalizer 依赖次序，也就无法调用 Finalizer 函数，这会导致目标对象无法变成不可达状态，其所占用内存无法被回收。
	 */
	runtime.SetFinalizer(&a, func(d *Data) { fmt.Printf("a %p final.\n", d) })
	runtime.SetFinalizer(&b, func(d *Data) { fmt.Printf("b %p final.\n", d) })
}

func main() {

	// go build -o uintptr uintptr.go && GODEBUG="gctrace=1" ./uintptr

	// mainUintptr()

	// mainUnsafePointer()

	// mainUnsafePointer2()

	// go build -gcflags "-N -l" -o uintptr uintptr.go && GODEBUG="gctrace=1" ./uintptr
	for {
		testPointer()
		time.Sleep(time.Millisecond)
	}
}
