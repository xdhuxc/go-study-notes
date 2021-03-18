package main

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/xdhuxc/go-study-notes/rpc/cert/proto"
)

const (
	cert = "../cert/server.crt"
	key  = "../cert/server.key"
	addr = "localhost:8080"
)

func main() {
	var opts []grpc.DialOption
	// serverNameOverride 的值，必须和生成 server.crt 时的 Common Name 的值一样
	creds, err := credentials.NewClientTLSFromFile(cert, "localhost")
	if err != nil {
		log.Errorf("new client tls error: %s", err)
		return
	}
	opts = append(opts, grpc.WithTransportCredentials(creds))

	conn, err := grpc.Dial(":8080", opts...)
	if err != nil {
		log.Errorf("fail to dial: %v", err)
		return
	}
	defer conn.Close()
	client := pb.NewEchoServiceClient(conn)

	msg, err := client.Echo(context.Background(), &pb.EchoMessage{Value: "this is client"})
	if err != nil {
		log.Errorf("call rpc error: %s", err)
		return
	}

	println(msg.Value)
}
