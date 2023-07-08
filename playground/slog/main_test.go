package main

import (
	"io"
	"log/slog"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func BenchmarkMarshal(b *testing.B) {
	marshalBenchmark(b.N, io.Discard)
}

func BenchmarkSlog(b *testing.B) {
	slogger = slog.New(slog.NewJSONHandler(
		io.Discard,
		&slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))

	b.ResetTimer()
	slogBenchmark(b.N)
}

func BenchmarkZap(b *testing.B) {
	kafkaEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	topicDebugging := zapcore.AddSync(io.Discard)
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})
	core := zapcore.NewTee(
		zapcore.NewCore(kafkaEncoder, topicDebugging, lowPriority),
	)
	zapLogger = zap.New(core)

	b.ResetTimer()
	zapBenchmark(b.N)
}
