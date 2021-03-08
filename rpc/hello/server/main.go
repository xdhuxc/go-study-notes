package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/xdhuxc/go-study-notes/rpc/hello/proto"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Received: %v", req.GetName())

	return &pb.GreetResponse{Message: "Hello " + req.GetName()}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Received %v again", req.GetName())

	return &pb.GreetResponse{Message: "Hello " + req.GetName() + " Again"}, nil
}

func main() {
	conn, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	log.Printf("listening at %s", port)
	if err := s.Serve(conn); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
