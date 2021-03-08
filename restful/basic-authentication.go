package main

import (
	"github.com/emicklei/go-restful"
	"io"
	"log"
	"net/http"
)

func basicAuthenticate(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	u, p, ok := req.Request.BasicAuth()
	if !ok || u != "admin" || p != "admin" {
		resp.AddHeader("WWW-Authenticate", "Basic realm=Protected Area")
		_ = resp.WriteErrorString(401, "401: Not Authorized")
		return
	}

	chain.ProcessFilter(req, resp)
}

func secret(req *restful.Request, resp *restful.Response) {
	_, _ = io.WriteString(resp, "42")
}

func main() {
	/**
	GET http://localhost:8080/secret
	使用 admin/admin 认证
	*/
	ws := new(restful.WebService)
	ws.Route(ws.GET("/secret").Filter(basicAuthenticate).To(secret))
	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
