package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type User4 struct {
	UserName string
	age      int
}

type Admin4 struct {
	User4
	title string
}

type User5 struct {
	UserName string
	age      int
}

func main() {
	s := make([]int, 0, 10)
	v := reflect.ValueOf(&s).Elem()

	v.SetLen(2)
	v.Index(0).SetInt(100)
	v.Index(1).SetInt(200)

	fmt.Println(v.Interface(), s)
	fmt.Println("======================")

	v2 := reflect.Append(v, reflect.ValueOf(300))
	v2 = reflect.AppendSlice(v2, reflect.ValueOf([]int{400, 500}))
	fmt.Println(v2.Interface())
	fmt.Println("----------------------")

	m := map[string]int{"a": 1}
	v = reflect.ValueOf(&m).Elem()

	v.SetMapIndex(reflect.ValueOf("a"), reflect.ValueOf(100))
	v.SetMapIndex(reflect.ValueOf("b"), reflect.ValueOf(200))
	fmt.Println(v.Interface(), m)
}

func value6() {
	u := User5{"Jack", 25}
	p := reflect.ValueOf(&u).Elem()

	p.FieldByName("UserName").SetString("xdhuxc")

	f := p.FieldByName("age")
	fmt.Println(f.CanSet())

	// 判断是否能获取地址
	if f.CanAddr() {
		age := (*int)(unsafe.Pointer(f.UnsafeAddr()))
		age = (*int)(unsafe.Pointer(f.Addr().Pointer()))
		*age = 90
	}

	// 注意，p 是 Value 类型，需要还原成接口才能转型
	fmt.Println(u, p.Interface().(User5))
}

func value5() {
	u := User5{"Jack", 25}

	v := reflect.ValueOf(u)
	p := reflect.ValueOf(&u)

	fmt.Println(v.CanSet(), v.FieldByName("UserName").CanSet())
	fmt.Println(p.CanSet(), p.Elem().FieldByName("UserName").CanSet())
}

func value4() {
	u := User4{}
	v := reflect.ValueOf(u)

	f := v.FieldByName("a")
	fmt.Println(f.Kind(), f.IsValid())
}

func value3() {
	v := reflect.ValueOf([]int{1, 2, 3})

	for i, n := 0, v.Len(); i < n; i++ {
		fmt.Println(v.Index(i).Int())
	}
	fmt.Println("---------------------")

	v = reflect.ValueOf(map[string]int{"a": 1, "b": 2, "c": 3})
	for _, k := range v.MapKeys() {
		fmt.Println(k.String(), v.MapIndex(k).Int())
	}
}

func value2() {
	u1 := User4{"xdhuxc", 25}
	v1 := reflect.ValueOf(u1)

	f := v1.FieldByName("age")

	fmt.Println(v1.FieldByName("UserName").Interface())
	fmt.Println(v1.FieldByName("age").Interface())

	if f.CanInterface() {
		fmt.Println(f.Interface())
	} else {
		fmt.Println(f.Int())
	}
}

func value1() {
	u := &Admin4{
		User4{"xdhuxc", 23},
		"NT",
	}

	v := reflect.ValueOf(u).Elem()

	// 用转换方法获取字段值
	fmt.Println(v.FieldByName("title").String())
	// 直接访问嵌入字段
	fmt.Println(v.FieldByName("age").Int())
	// 用多级序号访问嵌入字段成员
	fmt.Println(v.FieldByIndex([]int{0, 1}).Int())

}
