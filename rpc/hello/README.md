### RPC

1、编写 .proto 文件，定义服务

2、生成 .pb.go 文件
```markdown
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    hello.proto
```

3、编写服务端代码

4、编写客户端代码

5、运行服务端代码

6、运行客户端代码