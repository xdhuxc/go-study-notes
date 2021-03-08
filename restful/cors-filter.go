package main

import (
	"io"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/sirupsen/logrus"
)

type UserResource struct{}

func (u UserResource) RegisterTo(container *restful.Container) {
	ws := new(restful.WebService)

	ws.Path("/users").
		Consumes("*/*").
		Produces("*/*")

	ws.Route(ws.GET("/{user-id}").To(u.nop))

	ws.Route(ws.POST("").To(u.nop))

	ws.Route(ws.PUT("/{user-id}").To(u.nop))

	ws.Route(ws.DELETE("/{user-id}").To(u.nop))

	container.Add(ws)
}

func (u UserResource) nop(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp.ResponseWriter, "This would be a normal response")
}

func main() {
	wsContainer := restful.NewContainer()
	u := UserResource{}
	u.RegisterTo(wsContainer)

	// 添加过滤器以允许 COTS
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"X-Header"},
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "POST"},
		CookiesAllowed: false,
		Container:      wsContainer,
	}

	wsContainer.Filter(cors.Filter)

	wsContainer.Filter(wsContainer.OPTIONSFilter)

	logrus.Print("Start listening on localhost:8080")

	server := &http.Server{
		Addr:    ":8080",
		Handler: wsContainer,
	}

	logrus.Fatal(server.ListenAndServe())

}
