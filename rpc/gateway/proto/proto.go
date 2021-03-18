package proto

import (
	"context"
	"fmt"
)

type myService struct{}

func (m *myService) mustEmbedUnimplementedEchoServiceServer() {}

func (m *myService) Echo(c context.Context, s *EchoMessage) (*EchoMessage, error) {
	fmt.Printf("rpc request Echo(%q)\n", s.Value)
	return s, nil
}

func NewServer() *myService {
	return new(myService)
}
