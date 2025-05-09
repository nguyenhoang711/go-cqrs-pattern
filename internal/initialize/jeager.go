package initialize

import (
	"github.com/nguyenhoang711/go-cqrs-pattern/global"
	"github.com/nguyenhoang711/go-cqrs-pattern/pkg/tracing"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

// Approach 1: Using jaeger
func InitJeagerTracer() {
	if global.Config.JaegerConfig.Enable {
		tracer, closer, err := tracing.NewJaegerTracer(&global.Config.JaegerConfig)
		if err != nil {
			global.Logger.Error("error to new jaeger tracer")
			return
		}
		defer closer.Close()
		opentracing.SetGlobalTracer(tracer)
	}
}

// Approach 2: Using open-telemtry
func StartTracing() {
	_, err := tracing.StartTracing()
	if err != nil {
		global.Logger.Error("error to start tracing", zap.Error(err))
		return
	}
	global.Logger.Info("start tracing successfully!")
}
