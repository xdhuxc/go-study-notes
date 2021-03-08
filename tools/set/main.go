package main

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
)

type User struct {
	Id   string
	Name string
}

func main() {
	users := []User{
		{
			Id:   "1",
			Name: "1",
		},
		{
			Id:   "1",
			Name: "1",
		},
		{
			Id:   "1",
			Name: "1",
		},
		{
			Id:   "2",
			Name: "1",
		},
		{
			Id:   "2",
			Name: "1",
		},
		{
			Id:   "3",
			Name: "1",
		},
	}
	m := mapset.NewSet()
	var x []User
	for _, user := range users {
		if m.Contains(user.Id) {
			continue
		} else {
			fmt.Printf("%#v \n", user)
			m.Add(user.Id)
			x = append(x, user)
		}
	}

	fmt.Printf("the length is %d", len(x))
}
