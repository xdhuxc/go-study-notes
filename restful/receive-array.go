package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
)

type User struct {
	ID   string `json:"id" description:"identifier of the user"`
	Name string `json:"name" description:"name of the user" default:"xdhuxc"`
	Age  int    `json:"age" description:"age of the user" default:"21"`
}

func CreateUsers(req *restful.Request, resp *restful.Response) {
	var users []User
	err := req.ReadEntity(&users)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Add")

	for _, user := range users {
		fmt.Println(user)
	}

	_ = resp.WriteEntity(users)
}

func ReadIntegerArray(req *restful.Request, resp *restful.Response) {
	var is []int
	err := req.ReadEntity(&is)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(is)

	fmt.Println("URL Path: ", req.Request.URL.Path)
	fmt.Println("URL Raw Path: ", req.Request.URL.RawPath)
	fmt.Println("URL USER: ", req.Request.URL.User)
	fmt.Println("URL Host: ", req.Request.URL.Host)
	fmt.Println("URL Scheme: ", req.Request.URL.Scheme)
	fmt.Println(req.Request.URL)

	_ = resp.WriteEntity(is)
}

func main() {
	ws := new(restful.WebService)
	ws.Path("/restful").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	tags := []string{"xdhuxc"}

	ws.Route(ws.POST("/users").
		To(CreateUsers).
		Doc("get all users").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads([]User{}).
		Returns(http.StatusOK, "OK", []User{}))

	ws.Route(ws.POST("/integers").
		To(ReadIntegerArray).Reads([]int{}).
		Doc("get all integers").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(http.StatusOK, "OK", []int{}))

	container := restful.NewContainer()
	container.Add(ws)

	log.Print("Start listening on localhost:8080")
	server := &http.Server{
		Addr:    ":8080",
		Handler: container,
	}
	log.Fatal(server.ListenAndServe())
}
