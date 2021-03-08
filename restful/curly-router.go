package main

import (
	"github.com/emicklei/go-restful"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type User struct {
	ID   string
	Name string
}

type UserResource struct {
	users map[string]User
}

// GET http://localhost:8080/users/1
func (u UserResource) findUser(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("user_id")
	user := u.users[id]
	if len(user.ID) == 0 {
		resp.AddHeader("Content-Type", "text/plain")
		err := resp.WriteErrorString(http.StatusNoContent, "User could not be found.")
		if err != nil {
			log.Errorln(err)
		}
	} else {
		err := resp.WriteEntity(user)
		if err != nil {
			log.Errorln(err)
		}
	}
}

// PUT http://localhost:8080/users/1
// <User><ID>1</ID><Name>xdhuxc</Name></User>
func (u *UserResource) updateUser(req *restful.Request, resp *restful.Response) {
	user := new(User)
	err := req.ReadEntity(&user)
	if err != nil {
		u.users[user.ID] = *user
		err := resp.WriteEntity(user)
		if err != nil {
			log.Errorln(err)
		}
	} else {
		resp.AddHeader("Content-Type", "text/plain")
		err := resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		if err != nil {
			log.Errorln(err)
		}
	}
}

// POST http://localhost:8080/users
// <User><ID><1/ID><Name>wakaka</Name></User>
func (u *UserResource) createUser(req *restful.Request, resp *restful.Response) {
	user := User{ID: req.PathParameter("user_id")}
	err := req.ReadEntity(&user)
	if err == nil {
		u.users[user.ID] = user
		err := resp.WriteHeaderAndEntity(http.StatusCreated, user)
		if err != nil {
			log.Errorln(err)
		}
	} else {
		resp.AddHeader("Content-Type", "text/plain")
		err := resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		if err != nil {
			log.Errorln(err)
		}
	}
}

func (u *UserResource) removeUser(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("user_id")
	delete(u.users, id)
}

func (u UserResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/users").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	ws.Route(ws.GET("/{user_id}").To(u.findUser))
	ws.Route(ws.POST("").To(u.updateUser))
	ws.Route(ws.PUT("/{user_id}").To(u.createUser))
	ws.Route(ws.DELETE("/{user_id}").To(u.removeUser))

	container.Add(ws)
}

func main() {
	wsc := restful.NewContainer()
	wsc.Router(restful.CurlyRouter{})
	u := UserResource{map[string]User{}}
	u.Register(wsc)

	log.Printf("start listening on localhost:8080")
	server := &http.Server{
		Addr:    ":8080",
		Handler: wsc,
	}
	log.Fatal(server.ListenAndServe())
}
