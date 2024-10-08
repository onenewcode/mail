package main

import (
	"context"
	"io"
	"os"

	"frontend/biz/router"
	"frontend/conf"
	"frontend/infra/mtl"
	"frontend/infra/rpc"
	"frontend/middleware"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	hertzprom "github.com/hertz-contrib/monitor-prometheus"
	hertzotelprovider "github.com/hertz-contrib/obs-opentelemetry/provider"
	hertzoteltracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/pprof"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"
	"github.com/joho/godotenv"
	oteltrace "go.opentelemetry.io/otel/trace"
)

func main() {
	_ = godotenv.Load()
	// 可观测性初始化
	// mtl.InitMtl()
	rpc.InitClient()
	address := conf.GetConf().Hertz.Address
	// 存储日志
	f, err := os.OpenFile(conf.GetConf().Hertz.LogFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fileWriter := io.MultiWriter(f, os.Stdout)
	hlog.SetOutput(fileWriter)

	// 启动追踪，但不启动度量
	_ = hertzotelprovider.NewOpenTelemetryProvider(
		hertzotelprovider.WithSdkTracerProvider(mtl.TracerProvider),
		hertzotelprovider.WithEnableMetrics(false),
	)

	tracer, cfg := hertzoteltracing.NewServerTracer(hertzoteltracing.WithCustomResponseHandler(func(ctx context.Context, c *app.RequestContext) {
		c.Header("shop-trace-id", oteltrace.SpanFromContext(ctx).SpanContext().TraceID().String())
	}))

	h := server.New(server.WithHostPorts(address), server.WithTracer(hertzprom.NewServerTracer(
		"",
		"",
		hertzprom.WithRegistry(mtl.Registry),
		hertzprom.WithDisableServer(true),
	)), tracer)
	h.LoadHTMLGlob("template/*")
	h.Delims("{{", "}}")
	h.Use(hertzoteltracing.ServerMiddleware(cfg))

	registerMiddleware(h)

	router.GeneratedRegister(h)

	h.GET("sign-in", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "sign-in", utils.H{
			"title": "Sign in",
			"next":  c.Query("next"),
		})
	})
	h.GET("sign-up", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "sign-up", utils.H{
			"title": "Sign up",
		})
	})

	h.Static("/static", "./")

	h.Spin()
}

// 注册一些通用组件
func registerMiddleware(h *server.Hertz) {
	// pprof
	if conf.GetConf().Hertz.EnablePprof {
		pprof.Register(h)
	}
	store, err := redis.NewStore(100, "tcp", conf.GetConf().Redis.Address, "", []byte(os.Getenv("SESSION_SECRET")))
	if err != nil {
		panic(err)
	}
	store.Options(sessions.Options{MaxAge: 86400, Path: "/"})
	rs, err := redis.GetRedisStore(store)
	if err == nil {
		rs.SetSerializer(sessions.JSONSerializer{})
	}
	h.Use(sessions.New("cloudwego-shop", store))
	// gzip
	if conf.GetConf().Hertz.EnableGzip {
		h.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	// access log
	if conf.GetConf().Hertz.EnableAccessLog {
		h.Use(accesslog.New())
	}

	// recovery
	h.Use(recovery.Recovery())

	h.OnShutdown = append(h.OnShutdown, mtl.Hooks...)

	// cores
	h.Use(cors.Default())
	middleware.RegisterMiddleware(h)
}
