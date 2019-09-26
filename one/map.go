package main

import "fmt"

type user struct {
	name string
}

func main() {
	m := map[string]string{
		"a": "a",
		"b": "b",
	}

	// 对于不存在的 key，返回 value 的零值，不会出错
	fmt.Println(m["c"])

	// 判断 key 是否存在
	v, ok := m["d"]
	fmt.Println(ok)
	fmt.Println(v)
	if ok {
		fmt.Println(v)
	}

	m["c"] = "c"

	// 删除，如果 key 不存在，不会出错
	delete(m, "c")

	fmt.Println(len(m))

	for k, v := range m {
		fmt.Println(k, "--->", v)
	}

	m1 := map[int]user{
		1: {"user1"},
	}

	// m1[1].name = "xdhuxc" can not assign to m1[1].name
	u := m1[1]
	u.name = "xdhuxc"
	m1[1] = u

	m2 := map[int]*user{
		1: &user{"xdhuxc"},
	}

	m2[1].name = "xdhuxc"

}
