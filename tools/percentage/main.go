package main

import "fmt"

type User struct {
	Name   string
	Age    int
	UseMFA bool
}

func main() {
	var a int64 = 12
	var b int32 = 12

	fmt.Println(int(a))
	fmt.Println(int(b))

	u := User{
		Name:   "xdhuxc",
		Age:    10,
		UseMFA: false,
	}

	fmt.Printf("1 %v \n", u)
	fmt.Printf("2 %#v \n", u)
	fmt.Printf("3 %+v \n", u)

	users := []User{
		{
			Name:   "xdhuxc1",
			Age:    1,
			UseMFA: false,
		},
		{
			Name:   "xdhuxc2",
			Age:    100,
			UseMFA: true,
		},
		{
			Name:   "xdhuxc3",
			Age:    10,
			UseMFA: false,
		},
		{
			Name:   "xdhuxc4",
			Age:    12,
			UseMFA: true,
		},
		{
			Name:   "xdhuxc5",
			Age:    11,
			UseMFA: true,
		},
	}

	var x User
	for _, u := range users {
		x = u
		fmt.Println(x)

		fmt.Println(x == u)
	}

	for i := 0; i < len(users); i++ {

	}

}
