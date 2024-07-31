# day2
- 安装cwgo
  - > go install github.com/cloudwego/cwgo@latest
- 安装protobuf
  - https://github.com/protocolbuffers/protobuf
- 安装protoc-gen-go
  - > go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  - > go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# day3

[//]: # (- 安装etcd)

[//]: # (  - > sudo apt update)

[//]: # (  - > sudo apt install etcd)

[//]: # (  - `)

[//]: # (    docker run -d --name etcd-server \)

[//]: # (    --publish 2379:2379 \)

[//]: # (    --publish 2380:2380 \)

[//]: # (    --env ALLOW_NONE_AUTHENTICATION=yes \)

[//]: # (    --env ETCD_ADVERTISE_CLIENT_URLS=http://etcd-server:2379 \)

[//]: # (    bitnami/etcd:latest)

[//]: # (    `)

- 安装Consul,只能本机安装，否则会出现健康检查不通过
  - https://developer.hashicorp.com/consul/install#windows
  - 根目录执行以下命令 >consul agent -dev
  
[//]: # (  - `docker run -id --name=consul \)

[//]: # (    -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8500:8500 -p 8600:8600 \)

[//]: # (    -v consul-data:/consul/data \)

[//]: # (    hashicorp/consul agent -server -ui -node=n1 -bootstrap-expect=1 -client=0.0.0.0`)

# day8
自动生成服务端代码，结构目录参考
>https://www.cloudwego.io/zh/docs/cwgo/tutorials/layout/

# day9
完善ui界面