package mtl

import (
	"os"
	"time"

	"github.com/cloudwego/kitex/server"

	"github.com/cloudwego/kitex/pkg/klog"
	kitexzap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLog(logFileName string) {
	var opts []kitexzap.Option
	var output zapcore.WriteSyncer
	if os.Getenv("GO_ENV") != "online" {
		opts = append(opts, kitexzap.WithCoreEnc(zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())))
		output = os.Stdout
	} else {
		opts = append(opts, kitexzap.WithCoreEnc(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())))
		fileio, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic("open log file failed")
		}
		// 使用 MultiWriteSyncer 将标准输出和文件写入器组合起来
		writeSyncer := zapcore.NewMultiWriteSyncer(os.Stdout, fileio)
		// async log
		output = &zapcore.BufferedWriteSyncer{
			WS:            writeSyncer,
			FlushInterval: time.Minute,
		}
		server.RegisterShutdownHook(func() {
			output.Sync() //nolint:errcheck
		})
	}
	log := kitexzap.NewLogger(opts...)
	klog.SetLogger(log)
	klog.SetLevel(klog.LevelTrace)
	klog.SetOutput(output)
}
