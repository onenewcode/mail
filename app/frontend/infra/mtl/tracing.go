package mtl

import (
	"context"

	"frontend/utils"
	"github.com/cloudwego/hertz/pkg/route"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

var TracerProvider *tracesdk.TracerProvider

func InitTracing() route.CtxCallback {
	// 创建一个grpc导出器
	exporter, err := otlptracegrpc.New(context.Background())
	if err != nil {
		panic(err)
	}
	// 创建一个span处理器
	processor := tracesdk.NewBatchSpanProcessor(exporter)
	// 创建一个资源，定义我们自己服务的值
	res, err := resource.New(context.Background(), resource.WithAttributes(semconv.ServiceNameKey.String(utils.ServiceName)))
	if err != nil {
		res = resource.Default()
	}
	// 创建一个tracer provider
	TracerProvider = tracesdk.NewTracerProvider(tracesdk.WithSpanProcessor(processor), tracesdk.WithResource(res))
	otel.SetTracerProvider(TracerProvider)
	// 返回一个关闭函数，用于关闭导出器
	return route.CtxCallback(func(ctx context.Context) {
		exporter.Shutdown(ctx) //nolint:errcheck
	})
}
