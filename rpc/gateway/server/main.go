package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
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

func newGateway() (http.Handler, error) {
	ctx := context.Background()
	creds, err := credentials.NewClientTLSFromFile(cert, "localhost")
	if err != nil {
		log.Errorf("new client tls error: %s", err)
		return nil, err
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}

	gwmux := runtime.NewServeMux()
	err = pb.RegisterEchoServiceHandlerFromEndpoint(ctx, gwmux, addr, opts)
	if err != nil {
		return nil, err
	}

	return gwmux, nil
}

func NewTLSListener(inner net.Listener, config *tls.Config) net.Listener {
	return tls.NewListener(inner, config)
}

func getTLSConfig() *tls.Config {
	certInBytes, err := ioutil.ReadFile(cert)
	if err != nil {
		panic(err)
	}
	keyInBytes, err := ioutil.ReadFile(key)
	if err != nil {
		panic(err)
	}

	var tlsKeyPair *tls.Certificate
	pair, err := tls.X509KeyPair(certInBytes, keyInBytes)
	if err != nil {
		log.Errorf("TLS KeyPair err: %v\n", err)
		panic(err)
	}
	tlsKeyPair = &pair
	return &tls.Config{
		Certificates:       []tls.Certificate{*tlsKeyPair},
		NextProtos:         []string{http2.NextProtoTLS}, // HTTP2 TLS支持
		InsecureSkipVerify: true,
		ServerName:         "localhost",
	}
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

	gwmux, err := newGateway()
	if err != nil {
		log.Errorf("create gateway error: %s", err)
		return
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	srv := &http.Server{
		Addr:      addr,
		Handler:   grpcHandlerFunc(grpcServer, mux),
		TLSConfig: getTLSConfig(),
	}

	fmt.Printf("grpc on %s \n", ":8080")
	_ = srv.Serve(NewTLSListener(conn, getTLSConfig()))
}
