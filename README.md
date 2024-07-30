# day2
- 安装cwgo
  - > go install github.com/cloudwego/cwgo@latest
- 安装protobuf
  - https://github.com/protocolbuffers/protobuf
- 安装protoc-gen-go
  - > go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  - > go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# day3
- 安装etcd
  - > sudo apt update
  - > sudo apt install etcd
  - `
    docker run -d --name etcd-server \
    --publish 2379:2379 \
    --publish 2380:2380 \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    --env ETCD_ADVERTISE_CLIENT_URLS=http://etcd-server:2379 \
    bitnami/etcd:latest
`