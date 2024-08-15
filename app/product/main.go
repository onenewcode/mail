package main

import (
	"common/mtl"
	"common/serversuite"
	"net"
	"product/biz/dal"
	"product/conf"
	"rpc_gen/kitex_gen/product/productcatalogservice"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
)

var serviceName string

func main() {
	godotenv.Load()
	serviceName = conf.GetConf().Kitex.Service
	mtl.InitMetric(serviceName, conf.GetConf().Kitex.MetricsPort, conf.GetConf().Registry.RegistryAddress[0])
	mtl.InitTracing(serviceName)
	mtl.InitLog(conf.GetConf().Kitex.LogFileName)

	dal.Init()
	opts := kitexInit()

	svr := productcatalogservice.NewServer(new(ProductCatalogServiceImpl), opts...)
	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	serviceName := conf.GetConf().Kitex.Service

	opts = append(opts,
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
		server.WithSuite(serversuite.CommonServerSuite{CurrentServiceName: serviceName}),
	)

	r, err := consul.NewConsulRegister(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		klog.Fatal(err)
	}
	opts = append(opts, server.WithRegistry(r))

	return
}
