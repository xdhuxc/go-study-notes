package main

import (
	"github.com/emicklei/go-restful"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var logger *log.Logger = log.New(os.Stdout, "", 0)

func NCSACommonLogFormatLogger() restful.FilterFunction {
	return func(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
		username := "-"
		if req.Request.URL.User != nil {
			if name := req.Request.URL.User.Username(); name != "" {
				username = name
			}
		}
		chain.ProcessFilter(req, resp)
		logger.Printf("%s - %s [%s] \"%s %s %s\" %d %d",
			strings.Split(req.Request.RemoteAddr, ":")[0],
			username,
			time.Now().Format("02/Jan/2006:15:04:05 -0700"),
			req.Request.Method,
			req.Request.URL.RequestURI(),
			req.Request.Proto,
			resp.StatusCode(),
			resp.ContentLength())
	}
}

func rhello(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "pong")
}

func main() {
	// GET http://localhost:8080/ping，但是只返回指定 URI 的日志，其他 URI 没有日志输出
	ws := new(restful.WebService)
	ws.Filter(NCSACommonLogFormatLogger())
	ws.Route(ws.GET("/ping").To(rhello))
	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
