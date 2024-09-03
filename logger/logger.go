package logger

import (
	"context"

	"github.com/blendle/zapdriver"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *otelzap.Logger
	sugar  *otelzap.SugaredLogger
)

func GetLogger() *otelzap.Logger {
	return logger
}

func Sync() {
	logger.Sync()
}

func InitLogger(env string) {
	log := newLogger(env)
	logger = otelzap.New(log,
		otelzap.WithCallerDepth(1),
		otelzap.WithTraceIDField(true),          //Log trace_id to zap logging
		otelzap.WithMinLevel(zapcore.InfoLevel), //Set min level to debug for print trace_id field)
	)
	sugar = logger.Sugar()
}

func newLogger(env string) *zap.Logger {
	callerSkip1 := zap.AddCallerSkip(1)
	stackTrace := zap.AddStacktrace(zapcore.DPanicLevel)
	if env == "local" {
		return zap.Must(zap.NewDevelopment(callerSkip1))
	} else {
		return zap.Must(zapdriver.NewProductionConfig().Build(callerSkip1, stackTrace))
	}
}

func Info(ctx context.Context, msg string, fields ...zapcore.Field) {
	logger.Ctx(ctx).Info(msg, fields...)
}

func Infof(ctx context.Context, template string, args ...interface{}) {
	sugar.Ctx(ctx).Infof(template, args)
}

func InfoWithTraceId(ctx context.Context, projectId string, msg string, fields ...zapcore.Field) {
	spanContext := trace.SpanFromContext(ctx).SpanContext()

	traceId := spanContext.TraceID().String()
	spanId := spanContext.SpanID().String()
	isSample := spanContext.TraceFlags().IsSampled()

	fields = append(fields, zapdriver.TraceContext(traceId, spanId, isSample, projectId)...)

	logger.Ctx(ctx).Info(msg, fields...)
}

func Debug(ctx context.Context, msg string, fields ...zapcore.Field) {
	logger.Ctx(ctx).Debug(msg, fields...)
}

func Debugf(ctx context.Context, template string, args ...interface{}) {
	sugar.Ctx(ctx).Debugf(template, args)
}

func Fatal(ctx context.Context, msg string, fields ...zapcore.Field) {
	logger.Ctx(ctx).Fatal(msg, fields...)
}

func Fatalf(ctx context.Context, template string, args ...interface{}) {
	sugar.Ctx(ctx).Fatalf(template, args)
}

func Error(ctx context.Context, msg string, fields ...zapcore.Field) {
	logger.Ctx(ctx).Error(msg, fields...)
}

func Errorf(ctx context.Context, template string, args ...interface{}) {
	sugar.Ctx(ctx).Errorf(template, args)
}

func Warn(ctx context.Context, msg string, fields ...zapcore.Field) {
	logger.Ctx(ctx).Warn(msg, fields...)
}
func Warnf(ctx context.Context, template string, args ...interface{}) {
	sugar.Ctx(ctx).Warnf(template, args)
}

func DPanic(ctx context.Context, msg string, fields ...zapcore.Field) {
	logger.Ctx(ctx).DPanic(msg, fields...)
}

func Panic(ctx context.Context, msg string, fields ...zapcore.Field) {
	logger.Ctx(ctx).Panic(msg, fields...)
}
