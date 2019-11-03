package main

import (
	"fmt"
	"reflect"
)

type Data1 struct {
}

func (*Data1) Test(x int, y int) (int, int) {
	return x + 100, y + 100
}

func (*Data1) Sum(s string, x ...int) string {
	c := 0
	for _, n := range x {
		c += n
	}

	return fmt.Sprintf(s, c)
}

func info(m reflect.Method) {
	t := m.Type

	fmt.Println(m.Name)

	for i, n := 0, t.NumIn(); i < n; i++ {
		fmt.Printf("in[%d] %v\n", i, t.In(i))
	}

	for i, n := 0, t.NumOut(); i < n; i++ {
		fmt.Printf("out [%d] %v\n", i, t.Out(i))
	}
}

func test1() {
	d := new(Data1)
	t := reflect.TypeOf(d)

	test, _ := t.MethodByName("Test")
	info(test)

	sum, _ := t.MethodByName("Sum")
	info(sum)
}

func test2() {
	d := new(Data1)
	v := reflect.ValueOf(d)

	exec := func(name string, in []reflect.Value) {
		// 获取函数名称
		m := v.MethodByName(name)
		// 调用函数
		out := m.Call(in)

		for _, v := range out {
			fmt.Println(v.Interface())
		}
	}

	exec("Test", []reflect.Value{
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	})
	fmt.Println("-------------------")

	exec("Sum", []reflect.Value{
		reflect.ValueOf("result=%d"),
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	})
	fmt.Println("-------------------")

	// 将变参打包成 slice
	m := v.MethodByName("Sum")
	in := []reflect.Value{
		reflect.ValueOf("result = %d"),
		reflect.ValueOf([]int{1, 2}),
	}
	out := m.CallSlice(in)

	for _, v := range out {
		fmt.Println(v.Interface())
	}
}

var (
	Int    = reflect.TypeOf(0)
	String = reflect.TypeOf("")
)

func Make(T reflect.Type, fptr interface{}) {
	swap := func(in []reflect.Value) []reflect.Value {
		return []reflect.Value{
			reflect.MakeSlice(
				reflect.SliceOf(T),
				int(in[0].Int()),
				int(in[1].Int()),
			),
		}
	}
	// 传入的是函数变量指针，因为我们要将变量指向 swap 函数
	fn := reflect.ValueOf(fptr).Elem()
	// 获取函数指针类型，生成所需 swap function value
	v := reflect.MakeFunc(fn.Type(), swap)
	// 修改函数指针实际指向，也就是 swap
	fn.Set(v)
}

func main() {
	var makeints func(int, int) []int
	var makestrings func(int, int) []string

	Make(Int, &makeints)
	Make(String, &makestrings)

	x := makeints(5, 10)
	fmt.Printf("%#v\n", x)

	s := makestrings(3, 10)
	fmt.Printf("%#v\n", s)
}
