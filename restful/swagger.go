package main

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
	"log"
	"net/http"
)

type Book struct {
	Title  string
	Author string
}

func swagger(c *restful.Container) {
	config := restfulspec.Config{
		WebServices:                   c.RegisteredWebServices(),
		WebServicesURL:                "http://0.0.0.0:8081",
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject,
	}

	c.Handle("/apidocs", http.StripPrefix("/apidocs", http.FileServer(http.Dir("dist"))))
	c.Add(restfulspec.NewOpenAPIService(config))
}

func enrichSwaggerObject(swagger *spec.Swagger) {
	swagger.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "Xdhuxc APIServer",
			Description: "Resource for managing xdhuxc API",
			Contact: &spec.ContactInfo{
				Name:  "xdhuxc",
				Email: "xdhuxc@163.com",
				URL:   "http://xdhuxc.club",
			},
			License: &spec.License{
				Name: "MIT",
				URL:  "http://mit.org",
			},
			Version: "1.0.0",
		},
	}
	swagger.Tags = []spec.Tag{
		spec.Tag{
			TagProps: spec.TagProps{
				Name:        "users",
				Description: "Managing users",
			},
		},
	}
}

func main() {
	ws := new(restful.WebService)
	ws.Path("/books").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	container := restful.NewContainer()
	container.Add(ws)

	ws.Route(ws.GET("/{medium}").
		To(noop).
		Doc("Search all books").
		Param(ws.PathParameter("medium", "digital or paperback").DataType("string")).
		Param(ws.QueryParameter("language", "en,nl,de").DataType("string")).
		Param(ws.HeaderParameter("If-Modified-Since", "last known timestamp").DataType("datetime")).
		Do(returns200, returns500))

	ws.Route(ws.PUT("/{medium}").
		To(noop).Doc("Add a new book").
		Param(ws.PathParameter("medium", "digital or paperback").DataType("string")).
		Reads(Book{}))

	swagger(container)
	//swagger.RegisterSwaggerService(config, restful.DefaultContainer)
	log.Print("Start listening on localhost:8080")
	server := &http.Server{
		Addr:    ":8080",
		Handler: container,
	}
	log.Fatal(server.ListenAndServe())

}

func noop(req *restful.Request, resp *restful.Response) {

}

func returns200(b *restful.RouteBuilder) {
	b.Returns(http.StatusOK, "OK", Book{})
}

func returns500(b *restful.RouteBuilder) {
	b.Returns(http.StatusInternalServerError, "There is something wrong", nil)
}
