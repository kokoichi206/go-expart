package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strconv"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var marshalMu sync.Mutex
var slogger *slog.Logger
var zapLogger *zap.Logger

type logMsg struct {
	Time    string `json:"time"`
	Level   string `json:"level"`
	Message string `json:"message"`
}

func marshalBenchmark(n int, w io.Writer) {
	for i := 0; i < n; i++ {
		marshalCheck(strconv.Itoa(n), w)
	}
}

func marshalCheck(msg string, w io.Writer) {
	marshalMu.Lock()
	defer marshalMu.Unlock()

	b, _ := json.Marshal(&logMsg{
		Time:    time.Now().Format(time.RFC3339),
		Level:   "DEBUG",
		Message: msg,
	})
	fmt.Fprintln(w, string(b))
}

func slogBenchmark(n int) {
	for i := 0; i < n; i++ {
		slogCheck(strconv.Itoa(i))
	}
}

func slogCheck(msg string) {
	slogger.Debug(msg)
}

func zapBenchmark(n int) {
	for i := 0; i < n; i++ {
		zapCheck(strconv.Itoa(i))
	}
	zapLogger.Sync()
}

func zapCheck(msg string) {
	zapLogger.Info(msg)
}

func main() {
	marshalCheck("ぱおん", os.Stdout)

	// set time format for log package (not log/slog)
	// log.SetFlags(log.Ldate | log.Ltime)
	// log.Println("log test")

	// ------------------ initialize slog logger ------------------
	slogger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	// CANNOT change time format of JSONHandler
	slogCheck("pien; tukareta")

	// ------------------ initialize zap logger ------------------
	kafkaEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	topicDebugging := zapcore.AddSync(os.Stdout)
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})
	core := zapcore.NewTee(
		zapcore.NewCore(kafkaEncoder, topicDebugging, lowPriority),
	)
	// From a zapcore.Core, it's easy to construct a Logger.
	zapLogger = zap.New(core)

	zapCheck("zap test")
}
