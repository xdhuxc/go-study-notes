package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/xdhuxc/go-study-notes/rpc/cert/proto"
)

const (
	cert = "../cert/server.crt"
	key  = "../cert/server.key"
	addr = "localhost:8080"
)

var (
	KeyPair  *tls.Certificate
	CertPool *x509.CertPool
	Addr     string
)

func init() {
	var err error
	certInBytes, err := ioutil.ReadFile(cert)
	if err != nil {
		panic(err)
	}
	keyInBytes, err := ioutil.ReadFile(key)
	if err != nil {
		panic(err)
	}
	pair, err := tls.X509KeyPair(certInBytes, keyInBytes)
	if err != nil {
		log.Errorf("x509 key pair: %s", err)
		panic(err)
	}
	KeyPair = &pair
}

// grpcHandlerFunc returns an http.Handler that delegates to grpcServer on incoming gRPC
// connections or otherHandler otherwise. Copied from cockroachdb.
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO(tamird): point to merged gRPC code rather than a PR.
		// This is a partial recreation of gRPC's internal checks https://github.com/grpc/grpc-go/pull/514/files#diff-95e9a25b738459a2d3030e1e6fa2a718R61
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}

func main() {
	creds, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
		log.Errorf("new server from file error: %s", err)
		return
	}

	opts := []grpc.ServerOption{grpc.Creds(creds)}
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterEchoServiceServer(grpcServer, pb.NewServer())

	conn, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Errorf("listen error: %s", err)
		return
	}

	fmt.Printf("grpc on %s \n", ":8080")
	_ = grpcServer.Serve(conn)
}
