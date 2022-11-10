### 生成证书

`https://eddycjy.gitbook.io/golang/di-4-ke-grpc/grpc-tls`

openssl ecparam -genkey -name secp384r1 -out server.key
openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650

### 教程

`https://www.bookstack.cn/read/eddycjy-grpc/aed171745d70160c.md`
