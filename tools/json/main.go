package main

import (
	"encoding/json"
	"fmt"
)

/**
json 包是通过反射机制来实现编解码的，因此结构体必须导出所转换的字段，不导出的字段不会被 json 包解析
json 包在解析结构体时，如果遇到 key 为 json 的字段标签，则会按照一定规则解析该标签，第一个出现的是字段在 JSON 串中使用的名称，之后为其他选项，
例如，omitempty 指定空值字段不出现在 JSON 中，如果整个 value 为 -，则不解析该字段。
*/
type Person struct {
	UserId   int32  `json:"user_id,omitempty"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Password string `json:"name"`
}

func (person Person) String() string {
	return fmt.Sprintf("{UserId: %d, UserName: %s, Password: %s, Email: %s}", person.UserId, person.UserName, person.Password, person.Email)
}

func main() {

	person := Person{UserId: 12, UserName: "xdhuxc", Password: "xdhuxc234", Email: "xdhuxc@163.com"}

	/**
	data, err := json.Marshal(person)
	if err == nil {
		fmt.Printf("%s\n", data)
	}
	*/
	// 将结构体转换为 JSON 格式
	if data, err := json.Marshal(person); err == nil {
		fmt.Printf("%s\n", data)
	}

	personStr := `{"user_id":12,"user_name":"xdhuxc","password":"xdhuxc234","email":"xdhuxc@163.com"}`
	var xperson Person
	_ = json.Unmarshal([]byte(personStr), &xperson)

	fmt.Println(xperson)

	user := User{
		ID:       "abc",
		Name:     "xdhuxc",
		Address:  "s",
		Password: "073#asd",
	}
	dataInBytes, _ := json.Marshal(&user)
	fmt.Println(string(dataInBytes))
}
