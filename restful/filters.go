package main

import (
	"fmt"
	"github.com/emicklei/go-restful"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type Auth struct {
	Name                string   `yaml:"name"`
	SigningKey          string   `yaml:"signing_key"`
	EnableAuthOnOptions bool     `yaml:"enable_auth_on_options"`
	ContextKey          string   `yaml:"context_key"`
	HeaderKey           string   `yaml:"header_key"`
	ExcludeURL          []string `yaml:"exclude_url"`
	ExcludePrefix       []string `yaml:"exclude_prefix"`
}

func OK(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte("ok"))
}

func Enforce(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	auth := Auth{
		SigningKey:          "",
		EnableAuthOnOptions: false,
		ContextKey:          "",
		ExcludeURL:          []string{"/auth/hi"},
	}
	r := req.Request
	if !auth.EnableAuthOnOptions {
		if req.Request.Method == "OPTIONS" {
			// chain.ProcessFilter(req, resp)
			return
		}
	}

	// check exclude url
	if len(auth.ExcludeURL) > 0 {
		for _, url := range auth.ExcludeURL {
			fmt.Println(auth.ExcludeURL)
			if url == r.URL.Path {
				fmt.Println(r.URL.Path)
				// chain.ProcessFilter(req, resp)
				return
			}
		}
	}
	fmt.Println("fdsfdgfd")
	// check exclude url prefix
	if len(auth.ExcludePrefix) > 0 {
		for _, prefix := range auth.ExcludePrefix {
			if strings.HasPrefix(r.URL.Path, prefix) {
				// chain.ProcessFilter(req, resp)
				return
			}
		}
	}
}

func main() {
	restful.Filter(Enforce)

	restful.Filter(OK)

	ws := new(restful.WebService)
	restful.Add(ws)
	log.Infof("start listening on localhost %s", fmt.Sprintf(":%s", "8080"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8080"), nil))
}
