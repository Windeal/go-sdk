package wxworkrobot

import (
	"context"
	"reflect"
	"testing"
)

func TestGetDefaultLogger(t *testing.T) {
	tests := []struct {
		name string
		want Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDefaultLogger(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDefaultLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogDebugContextf(t *testing.T) {
	type args struct {
		ctx    context.Context
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "TestLogDebugContextf",
			args: args{
				ctx:    context.Background(),
				format: "TestLogDebugContextf:%d, %s",
				args:   []interface{}{1, "======Hello====="},
			},
		},
		{
			name: "TestLogDebugContextf context",
			args: args{
				ctx:    context.WithValue(context.Background(), ContextTraceKey, "trace-id-001"),
				format: "TestLogDebugContextf:%d, %s",
				args:   []interface{}{1, "======Hello====="},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogDebugContextf(tt.args.ctx, tt.args.format, tt.args.args...)
		})
	}
}

func TestLogErrorContextf(t *testing.T) {
	type args struct {
		ctx    context.Context
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogErrorContextf(tt.args.ctx, tt.args.format, tt.args.args...)
		})
	}
}

func TestLogFatalContextf(t *testing.T) {
	type args struct {
		ctx    context.Context
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogFatalContextf(tt.args.ctx, tt.args.format, tt.args.args...)
		})
	}
}

func TestLogInfoContextf(t *testing.T) {
	type args struct {
		ctx    context.Context
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogInfoContextf(tt.args.ctx, tt.args.format, tt.args.args...)
		})
	}
}

func TestLogWarnContextf(t *testing.T) {
	type args struct {
		ctx    context.Context
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogWarnContextf(tt.args.ctx, tt.args.format, tt.args.args...)
		})
	}
}

func TestLoggerImpl_LogDebugContextf(t *testing.T) {
	type args struct {
		ctx    context.Context
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := &LoggerImpl{}
			impl.LogDebugContextf(tt.args.ctx, tt.args.format, tt.args.args...)
		})
	}
}

func TestLoggerImpl_LogErrorContextf(t *testing.T) {
	type args struct {
		ctx    context.Context
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := &LoggerImpl{}
			impl.LogErrorContextf(tt.args.ctx, tt.args.format, tt.args.args...)
		})
	}
}

func TestLoggerImpl_LogFatalContextf(t *testing.T) {
	type args struct {
		ctx    context.Context
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := &LoggerImpl{}
			impl.LogFatalContextf(tt.args.ctx, tt.args.format, tt.args.args...)
		})
	}
}

func TestLoggerImpl_LogInfoContextf(t *testing.T) {
	type args struct {
		ctx    context.Context
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := &LoggerImpl{}
			impl.LogInfoContextf(tt.args.ctx, tt.args.format, tt.args.args...)
		})
	}
}

func TestLoggerImpl_LogWarnContextf(t *testing.T) {
	type args struct {
		ctx    context.Context
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := &LoggerImpl{}
			impl.LogWarnContextf(tt.args.ctx, tt.args.format, tt.args.args...)
		})
	}
}

func TestSetDefaultLogger(t *testing.T) {
	type args struct {
		logger Logger
	}
	tests := []struct {
		name string
		args args
		want Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetDefaultLogger(tt.args.logger)
			if !reflect.DeepEqual(tt.args.logger, GetDefaultLogger()) {
				t.Errorf("SetDefaultLogger() = %v, want %v", GetDefaultLogger(), tt.want)
			}
		})
	}
}
