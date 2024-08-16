package mtl

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"frontend/conf"

	"github.com/cloudwego/hertz/pkg/common/utils"

	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/route"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/registry/consul"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Registry *prometheus.Registry

func initMetric() route.CtxCallback {
	// 新建一个prometheus的注册器,默认
	Registry = prometheus.NewRegistry()
	// 注册一个go的采集器
	Registry.MustRegister(collectors.NewGoCollector())
	// 注册一个进程的采集器
	Registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	config := consulapi.DefaultConfig()
	config.Address = conf.GetConf().Hertz.RegistryAddr
	consulClient, _ := consulapi.NewClient(config)
	r := consul.NewConsulRegister(consulClient)
	// 获取本机非回环ip
	localIp := utils.LocalIP()
	ip, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", localIp, conf.GetConf().Hertz.MetricsPort))
	if err != nil {
		hlog.Error(err)
	}
	registryInfo := &registry.Info{Addr: ip, ServiceName: "prometheus", Weight: 1, Tags: map[string]string{
		"service": "frontend",
	}}
	// 注册我们的监控服务
	err = r.Register(registryInfo)
	if err != nil {
		hlog.Error(err)
	}

	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))
	go http.ListenAndServe(fmt.Sprintf(":%d", conf.GetConf().Hertz.MetricsPort), nil) //nolint:errcheck
	return func(ctx context.Context) {
		r.Deregister(registryInfo) //nolint:errcheck
	}
}
