package main

import (
	"fmt"
	"reflect"
)

type User struct {
	UserName string
}

type Admin struct {
	User
	title string
}

type User2 struct{}

type Admin2 struct {
	User2
}

func (*User2) ToString() {}

func (Admin2) Test() {}

type User3 struct {
	UserName string
	age      int
}

type TestUserName struct {
	UserName string
	number   int
}

type Admin3 struct {
	User3
	title string
	TestUserName
}

type MyUser struct {
	Name string `field:"username" type:"varchar(20)"`
	Age  int    `field:"age" type:"tinyint"`
}

var (
	Int    = reflect.TypeOf(0)
	String = reflect.TypeOf("")
)

type data3 struct {
}

func (*data3) String() string {
	return ""
}

type data4 struct {
	b byte
	x int32
}

func main() {
	// 某些时候，获取对齐信息对于内存自动分析是很有用的
	var d data4

	t := reflect.TypeOf(d)
	fmt.Println(t.Size(), t.Align()) // sizeof，以及最宽字段的对其模数

	f, _ := t.FieldByName("b")
	fmt.Println(f.Type.FieldAlign()) // 字段对其

}

func reflect8() {
	/**
	方法 Implements 判断是否实现了某个具体接口，AssignableTo，ConvertibleTo 用于赋值和转换判断
	*/

	var d *data3
	t := reflect.TypeOf(d)

	/**
	无法直接获取接口类型，但接口本身是个 struct，创建一个空指针对象，这样传递给 TypeOf 转换成 interface{} 时就有类型信息了
	*/
	it := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

	fmt.Println(t.Implements(it))
}

func reflect7() {
	// 可以使用方法 Elem 返回复合类型的基类型
	t := reflect.TypeOf(make(chan int)).Elem()
	fmt.Println(t)
}

func reflect6() {
	// 可以通过基本类型获取所对应的复合类型
	c := reflect.ChanOf(reflect.SendDir, String)
	fmt.Println(c)

	m := reflect.MapOf(String, Int)
	fmt.Println(m)

	s := reflect.SliceOf(Int)
	fmt.Println(s)

	t := struct{ Name string }{}
	p := reflect.PtrTo(reflect.TypeOf(t))
	fmt.Println(p)
}

func reflect5() {
	// 获取字段标签
	var mu MyUser

	t := reflect.TypeOf(mu)

	f, _ := t.FieldByName("Name")

	fmt.Println(f.Tag)
	fmt.Println(f.Tag.Get("field"))
	fmt.Println(f.Tag.Get("type"))
}

func reflect4() {
	var u Admin3
	t := reflect.TypeOf(u)

	f, _ := t.FieldByName("title")
	fmt.Println(f.Name)

	// 访问嵌入字段
	f, _ = t.FieldByName("User3")
	fmt.Println(f.Name)

	// 直接访问嵌入字段成员，会自动深度查找，若嵌入字段中有多个 UserName，值为空字符串
	f, _ = t.FieldByName("UserName")
	fmt.Println(f.Name)

	// Admin3[0] -> User3[1] => age
	f = t.FieldByIndex([]int{0, 1})
	fmt.Println(f.Name)
}

func reflect3() {
	var u Admin2
	methods := func(t reflect.Type) {
		for i, n := 0, t.NumMethod(); i < n; i++ {
			m := t.Method(i) // 小写的 test() 不能被反射读取
			fmt.Println(m.Name)
		}
	}

	fmt.Println("-- value interface --")
	methods(reflect.TypeOf(u))

	fmt.Println("-- pointer interface --")
	methods(reflect.TypeOf(&u))
}

func reflect2() {
	u := new(Admin)

	t := reflect.TypeOf(u)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type)
	}
}

func reflect1() {
	var a Admin
	t := reflect.TypeOf(a)

	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type)
	}
}
