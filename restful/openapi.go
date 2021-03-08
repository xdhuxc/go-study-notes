package main

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
	"net/http"
)

type User struct {
	ID   string `json:"id" description:"identifier of the user"`
	Name string `json:"name" description:"name of the user" default:"xdhuxc"`
	Age  int    `json:"age" description:"age of the user" default:"21"`
}

type UserResource struct {
	users map[string]User
}

func (u UserResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)
	tags := []string{"xdhuxc"}

	ws.Route(ws.GET("/").
		To(u.findAllUsers).
		Doc("get all users").
		Metadata(restfulspec.KeyOpenAPITags, tags).Writes([]User{}).
		Returns(http.StatusOK, "OK", []User{}))

	return ws
}

func (u UserResource) findAllUsers(req *restful.Request, resp *restful.Response) {
	users := []User{}
	for _, user := range u.users {
		users = append(users, user)
	}
	resp.WriteEntity(users)
}

func main() {
	u := UserResource{map[string]User{}}

	restful.DefaultContainer.Add(u.WebService())

	config := restfulspec.Config{
		WebServices:                   restful.RegisteredWebServices(),
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject,
	}

	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))

}
