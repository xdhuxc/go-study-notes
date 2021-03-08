package main

import (
	"github.com/xdhuxc/go-study-notes/rpc/user/model"
	"net/http"

	"golang.org/x/net/trace"
	"google.golang.org/grpc/grpclog"
)

var users []model.User

func init() {
	users = []model.User{
		{
			Id:    1,
			Name:  "11",
			Email: "11@163.com",
		},
		{
			Id:    2,
			Name:  "22",
			Email: "22@163.com",
		},
		{
			Id:    3,
			Name:  "33",
			Email: "33@163.com",
		},
	}
}

type userServer struct {
	*pb.UnimplementedU
}

func main() {

	// 开启trace
	go startTrace()

}

func startTrace() {
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}

	go http.ListenAndServe(":50051", nil)
	grpclog.Info("Trace listen on 50051")
}
