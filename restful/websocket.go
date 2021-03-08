package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var (
	entropy      *rand.Rand
	entropyMutex sync.Mutex
)

func init() {
	entropy = rand.New(rand.NewSource(time.Now().UnixNano()))
}

type Info struct {
	Websocket    bool     `json:"websocket"`
	CookieNeeded bool     `json:"cookie_needed"`
	Origins      []string `json:"origins"`
	Entropy      int32    `json:"entropy"`
}

func generateEntropy() int32 {
	entropyMutex.Lock()
	entropy := entropy.Int31()
	entropyMutex.Unlock()
	return entropy
}

var (
	addr = flag.String("addr", "localhost:8080", "http service address")
)

func init() {
	flag.Parse()
}

func echo(req *restful.Request, resp *restful.Response) {
	fmt.Println("-------------------------")
	fmt.Println(req.PathParameter("type"))

	upgrader := websocket.Upgrader{}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	c, err := upgrader.Upgrade(resp.ResponseWriter, req.Request, nil)
	if err != nil {
		log.Println("upgrade: ", err)
		return
	}
	defer c.Close()
	for {
		messageType, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read: ", err)
		}
		fmt.Println(string(message))
		log.Printf("receive: %s", string(message))
		println(messageType)

		err = c.WriteMessage(messageType, message)
		if err != nil {
			log.Println("write: ", err)
		}
	}
}

func info(req *restful.Request, resp *restful.Response) {

	switch req.Request.Method {
	case "GET":
		resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
		// resp.Header().Set("Access-Control-Allow-Credentials", "true")
		_ = json.NewEncoder(resp).Encode(Info{
			Websocket:    true,
			CookieNeeded: false,
			Origins:      []string{"*:*"},
			Entropy:      generateEntropy(),
		})
	case "OPTIONS":
		resp.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET")
		resp.Header().Set("Access-Control-Max-Age", fmt.Sprintf("%d", 365*24*60*60))
		resp.WriteHeader(http.StatusNoContent) // 204
	default:
		http.NotFound(resp, req.Request)
	}
}

func main() {
	ws := new(restful.WebService)
	ws.Path("/apps").
		Consumes(restful.MIME_JSON, restful.MIME_XML).
		Produces(restful.MIME_JSON, restful.MIME_XML)
	tags := []string{"xdhuxc"}

	ws.Route(ws.GET("/echo").
		To(echo).
		Doc("the echo").
		Param(ws.QueryParameter("type", "the type of connection").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(http.StatusOK, "OK", nil).
		Returns(http.StatusBadRequest, "ERROR", nil))

	ws.Route(ws.GET("/info").
		To(info).
		Doc("info").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(http.StatusOK, "OK", nil))

	container := restful.NewContainer()
	container.Add(ws)

	// 添加过滤器以允许 COTS
	cors := restful.CrossOriginResourceSharing{
		AllowedDomains: []string{},
		AllowedHeaders: []string{"Accept", "Access-Control-Allow-Origin", "X-CSRF-Token", "Authorization"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS", "DELETE", "PUT"},
		CookiesAllowed: false,
		Container:      container,
	}

	container.Filter(cors.Filter)
	container.Filter(container.OPTIONSFilter)

	log.Print("Start listening on localhost:8888")
	server := &http.Server{
		Addr:    ":8888",
		Handler: container,
	}
	log.Fatal(server.ListenAndServe())
}
