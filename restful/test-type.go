package main

import (
	"fmt"
	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"log"
	"net/http"
	"strconv"
	"time"
)

func getTime(req *restful.Request) (int64, int64) {
	end, err := strconv.ParseInt(req.QueryParameter("end"), 10, 64)
	if err != nil {
		end = int64(time.Now().Unix())
		fmt.Println("getTime", end)
	}

	start, err := strconv.ParseInt(req.QueryParameter("start"), 10, 64)
	if err != nil {
		start = int64(end - 60*60)
		fmt.Println("getTime", start)
	}

	return start, end
}

func receiver(req *restful.Request, resp *restful.Response) {
	start, end := getTime(req)
	fmt.Println(start, end)
	fmt.Println("URL Path: ", req.Request.URL.Path)
	fmt.Println("URL Raw Path: ", req.Request.URL.RawPath)
	fmt.Println("URL USER: ", req.Request.URL.User)
	fmt.Println("URL Host: ", req.Request.URL.Host)
	fmt.Println("URL Scheme: ", req.Request.URL.Scheme)
	fmt.Println(req.Request.URL)
}

func main() {
	ws := new(restful.WebService)
	ws.Path("/restful").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	tags := []string{"xdhuxc"}

	ws.Route(ws.GET("/time").
		To(receiver).
		Doc("get all users").
		Param(ws.QueryParameter("start", "start").DataType("integer").Required(false)).
		Param(ws.QueryParameter("end", "end").DataType("integer").Required(false)).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(http.StatusOK, "OK", nil))

	container := restful.NewContainer()
	container.Add(ws)

	log.Print("Start listening on localhost:8080")
	server := &http.Server{
		Addr:    ":8080",
		Handler: container,
	}
	log.Fatal(server.ListenAndServe())
}
