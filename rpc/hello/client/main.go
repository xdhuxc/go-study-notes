package main

import (
	"context"
	pb "github.com/xdhuxc/go-study-notes/rpc/hello/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("connect error: %s", err)
	}

	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	num := 0
	for {
		time.Sleep(5 * time.Second)
		r1, err := client.SayHello(ctx, &pb.GreetRequest{Name: "xdhuxc"})
		if err != nil {
			log.Fatalf("greet error: %v", err)
		}
		log.Printf("Greeting to %s %d times", r1.GetMessage(), num)

		r2, err := client.SayHelloAgain(ctx, &pb.GreetRequest{Name: "xdhuxc-again"})
		if err != nil {
			log.Fatalf("greet again error: %v", err)
		}
		log.Printf("Greeting to %s %d times", r2.GetMessage(), num)
		num++
	}
}
