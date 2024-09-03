package retrylogger

import (
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LeveledOtelZap - is logger to use wtih retryablehttp ("github.com/hashicorp/go-retryablehttp")
// this logger use wtih otelzap logger base
type LeveledOtelZap struct {
	*otelzap.Logger
}

func fields(keysAndValues []interface{}) []zapcore.Field {
	fields := []zapcore.Field{}

	for i := 0; i < len(keysAndValues)-1; i += 2 {
		fields = append(fields, zap.Any(keysAndValues[i].(string), keysAndValues[i+1]))
	}

	return fields
}

func (l *LeveledOtelZap) Error(msg string, keysAndValues ...interface{}) {
	l.With(fields(keysAndValues)...).Error(msg)
}

func (l *LeveledOtelZap) Info(msg string, keysAndValues ...interface{}) {
	l.With(fields(keysAndValues)...).Info(msg)
}

func (l *LeveledOtelZap) Debug(msg string, keysAndValues ...interface{}) {
	l.With(fields(keysAndValues)...).Debug(msg)
}

func (l *LeveledOtelZap) Warn(msg string, keysAndValues ...interface{}) {
	l.With(fields(keysAndValues)...).Warn(msg)
}
