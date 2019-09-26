package main

import "fmt"

type Node struct {
	_    int
	id   int
	data *byte
	next *Node
}

type User struct {
	id   int
	name string
}

type Manager struct {
	User
	title string
}

func (u *User) String() string {
	return fmt.Sprintf("User: %p, %v", u, &u)
}

func (m *Manager) String() string {
	return fmt.Sprintf("Manager: %p, %v", m, &m)
}

func (u *User) Test() {
	fmt.Printf("%p, %v\n", &u, u)
}

func main() {
	n1 := Node{
		id:   1,
		data: nil,
	}

	n2 := Node{
		id:   2,
		data: nil,
		next: &n1,
	}

	fmt.Println(n2)

	m := Manager{
		User{2, "xdhuxc"},
		"Administrator",
	}

	fmt.Printf("Manager: %p\n", &m)
	fmt.Println(m.User.String())
	fmt.Println(m.String())

	fmt.Println("=======================")
	u := User{1, "Tom"}
	u.Test()

	// 隐式传递 receiver
	mValue := u.Test
	mValue()

	// 显式传递 receiver
	mExpression := (*User).Test
	mExpression(&u)
}
