package serversuite

import (
	"os"
	"strings"

	"common/mtl"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/config-etcd/etcd"
	etcdServer "github.com/kitex-contrib/config-etcd/server"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

type CommonServerSuite struct {
	CurrentServiceName string
}

func (s CommonServerSuite) Options() []server.Option {

	opts := []server.Option{
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
	}

	if os.Getenv("CONFIG_CENTER_ENABLED") == "true" {
		etcdNodes := os.Getenv("CONFIG_CENTER_NODES")
		if etcdNodes != "" {
			etcdClient, err := etcd.NewClient(etcd.Options{Node: strings.Split(etcdNodes, ",")})
			if err != nil {
				klog.Error(err)
			} else {
				opts = append(opts, server.WithSuite(etcdServer.NewSuite(s.CurrentServiceName, etcdClient)))
			}
		}
	}

	_ = provider.NewOpenTelemetryProvider(provider.WithSdkTracerProvider(mtl.TracerProvider), provider.WithEnableMetrics(false))

	opts = append(opts,
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithTracer(prometheus.NewServerTracer("", "", prometheus.WithDisableServer(true), prometheus.WithRegistry(mtl.Registry))),
	)

	return opts
}
