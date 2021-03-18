grpc:
	@echo ">> generate grpc .pb.go file ..."
	@protoc -I. -I${GOPATH}/src/github.com/googleapis/googleapis --go_out . --go_opt paths=source_relative --go-grpc_out . --go-grpc_opt paths=source_relative cert.proto
	@echo ">> generate grpc .pb.gw.go file ..."
	@protoc -I. -I${GOPATH}/src/github.com/googleapis/googleapis --proto_path=. --grpc-gateway_out=logtostderr=true:. cert.proto
	@echo ">> generate grpc swagger file ..."
    @protoc -I. -I${GOPATH}/src/github.com/googleapis/googleapis --proto_path=. --swagger_out=logtostderr=true:../proto ../proto/cert.proto
