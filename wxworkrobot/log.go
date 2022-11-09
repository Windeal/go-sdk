package wxworkrobot

import (
	"context"
	"fmt"
)

type Logger interface {
	LogDebugContextf(ctx context.Context, format string, args ...interface{})
	LogInfoContextf(ctx context.Context, format string, args ...interface{})
	LogWarnContextf(ctx context.Context, format string, args ...interface{})
	LogErrorContextf(ctx context.Context, format string, args ...interface{})
	LogFatalContextf(ctx context.Context, format string, args ...interface{})
}

var defaultLogger Logger

const ContextTraceKey = "ContextTraceKey"

func init() {
	SetDefaultLogger(&LoggerImpl{})
}

func GetDefaultLogger() Logger {
	return defaultLogger
}

func SetDefaultLogger(logger Logger) {
	defaultLogger = logger
}

type LoggerImpl struct {
}

func (impl *LoggerImpl) LogDebugContextf(ctx context.Context, format string, args ...interface{}) {
	format = " DEBUG: " + format
	traceID, ok := ctx.Value(ContextTraceKey).(string)
	if ok {
		format = "ContextTraceKey: " + traceID + format
	}
	format = format + "\n"
	fmt.Printf(format, args...)
}

func (impl *LoggerImpl) LogInfoContextf(ctx context.Context, format string, args ...interface{}) {
	format = " INFO: " + format
	traceID, ok := ctx.Value(ContextTraceKey).(string)
	if ok {
		format = "ContextTraceKey: " + traceID + format
	}
	format = format + "\n"
	fmt.Printf(format, args...)
}

func (impl *LoggerImpl) LogWarnContextf(ctx context.Context, format string, args ...interface{}) {
	format = " WARN: " + format
	traceID, ok := ctx.Value(ContextTraceKey).(string)
	if ok {
		format = "ContextTraceKey: " + traceID + format
	}
	format = format + "\n"
	fmt.Printf(format, args...)
}

func (impl *LoggerImpl) LogErrorContextf(ctx context.Context, format string, args ...interface{}) {
	format = " ERROR: " + format
	traceID, ok := ctx.Value(ContextTraceKey).(string)
	if ok {
		format = "ContextTraceKey: " + traceID + format
	}
	format = format + "\n"
	fmt.Printf(format, args...)
}

func (impl *LoggerImpl) LogFatalContextf(ctx context.Context, format string, args ...interface{}) {
	format = " FATAL: " + format
	traceID, ok := ctx.Value(ContextTraceKey).(string)
	if ok {
		format = "ContextTraceKey: " + traceID + format
	}
	format = format + "\n"
	fmt.Printf(format, args...)
}

func LogDebugContextf(ctx context.Context, format string, args ...interface{}) {
	GetDefaultLogger().LogDebugContextf(ctx, format, args...)
}

func LogInfoContextf(ctx context.Context, format string, args ...interface{}) {
	GetDefaultLogger().LogInfoContextf(ctx, format, args...)
}

func LogWarnContextf(ctx context.Context, format string, args ...interface{}) {
	GetDefaultLogger().LogWarnContextf(ctx, format, args...)
}

func LogErrorContextf(ctx context.Context, format string, args ...interface{}) {
	GetDefaultLogger().LogErrorContextf(ctx, format, args...)
}

func LogFatalContextf(ctx context.Context, format string, args ...interface{}) {
	GetDefaultLogger().LogFatalContextf(ctx, format, args...)
}
