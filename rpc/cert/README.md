### 带认证的 RPC


### 生成证书的方法

1、使用如下命令安装 certstrap
```markdown
brew install certstrap
```

2、生成证书文件
```markdown
certstrap init --common-name "localhost"
```


### 生成证书  
1、生成私钥
```markdown
openssl genrsa -out server.key 2048
```

2、生成公钥
```markdown
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

```markdown
wanghuan@wanghuans-MacBook-Pro cert % openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) []:
State or Province Name (full name) []:
Locality Name (eg, city) []:
Organization Name (eg, company) []:
Organizational Unit Name (eg, section) []:
Common Name (eg, fully qualified host name) []:xdhuxc
Email Address []:
```

### 常见问题及解决
1、启动客户端时，报错如下：
```markdown
wanghuan@wanghuans-MacBook-Pro client % go run main.go
ERRO[0000] call rpc error: rpc error: code = Unavailable desc = connection error: desc = "transport: authentication handshake failed: x509: certificate relies on legacy Common Name field, use SANs or temporarily enable Common Name matching with GODEBUG=x509ignoreCN=0"
```
改为如下方式执行启动命令：
```markdown
wanghuan@wanghuans-MacBook-Pro client % GODEBUG=x509ignoreCN=0 go run main.go 
this is client
```
即可执行成功

2、在初始化客户端 TransportCredentials 时，serverNameOverride 的值必须和生成 server.crt 时的 Common Name 的值一样，否则会报如下错误：
```markdown
wanghuan@wanghuans-MacBook-Pro client % GODEBUG=x509ignoreCN=0 go run main.go
ERRO[0000] call rpc error: rpc error: code = Unavailable desc = connection error: desc = "transport: authentication handshake failed: x509: certificate is valid for xdhuxc, not localhost"
```
