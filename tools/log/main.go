package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name    string
	Address string
}

func (u *User) String() string {
	if dataInBytes, err := json.Marshal(u); err == nil {
		return string(dataInBytes)
	}

	return ""
}

func main() {
	err := fmt.Errorf("it is %s", "error")
	if err != nil {
		fmt.Println(err)
	}

	err = nil
	if err == nil {
		fmt.Println(err)
	}

	u := &User{
		Name:    "xdhuxc",
		Address: "xdhuxc@163.com",
	}

	fmt.Println("the user is ", u.String())

	cfg := &Config{
		Stdout: true,
	}
	InitLogger(cfg)
	defer Logger.Sync()

	Logger.Info("it is a info message")
	Logger.Error("it is a error message")

	// Logger.Debug("it is a debug message")

	// Logger.Fatal("it is a fatal message")

	a := 100
	b := 300

	x := fmt.Sprintf("%.2f", float32(b)/float32(a))
	fmt.Println(x)
}
