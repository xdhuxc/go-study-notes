package main

import "fmt"

var (
	SigningKey          string = "xdhuxc.com"
	EnableAuthOnOptions bool   = false
	ContextKey          string = "user"
)

type Auth struct {
	Name                string   `yaml:"name"`
	SigningKey          string   `yaml:"signing_key"`
	EnableAuthOnOptions bool     `yaml:"enable_auth_on_options"`
	ContextKey          string   `yaml:"context_key"`
	HeaderKey           string   `yaml:"header_key"`
	ExcludeURL          []string `yaml:"exclude_url"`
	ExcludePrefix       []string `yaml:"exclude_prefix"`
	// Default value "x-scmp-shareid".
	XShareidKey string `yaml:"x_shareid_key"`
	// Default value "x-scmp-group".
	XGroupKey string `yaml:"x_group_key"`
	// Default value "x-scmp-role".
	XRoleKey string `yaml:"x_role_key"`
}

func main() {
	a := Auth{
		ExcludeURL:    nil,
		ExcludePrefix: []string{"x", "y", "z"},
	}
	for _, item := range a.ExcludeURL {
		fmt.Println("url")
		fmt.Println(item)
	}

	for _, item := range a.ExcludePrefix {
		fmt.Println(item)
	}

}
