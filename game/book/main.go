package main

import (
	"context"
	"io"
	"log"
	"log/slog"
	"os"
)

type Me struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	ctx := context.WithValue(context.Background(), "request-id", "pien")

	slog.Info("main msg", "me", Me{Name: "john doe", Age: 128})

	// Handler Interface: https://pkg.go.dev/golang.org/x/exp/slog#Handler
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		// AddSource: true,
		// Level:     slog.LevelInfo,
	}))
	slog.SetDefault(logger)
	// 通常のログまで変わるらしい！
	log.Print("normal log msg")

	// log object (as a json type) without defining a Go struct.
	logger.Info("group log", slog.Group("group_key", "name", "ooo"))

	loggerWithGroup := logger.WithGroup("group_key")
	loggerWithGroup.Info("msg22", slog.String("key", "value"))

	myHandler := &myJSONHandler{
		Handler: slog.NewJSONHandler(os.Stdout, nil),
	}
	myLogger := slog.New(myHandler)
	myLogger.InfoContext(ctx, "message yo")

	lg := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	lg.InfoContext(ctx, "secret test", slog.Any("password", Secret("12345678")))

	// words()
}

type myJSONHandler struct {
	slog.Handler
	w io.Writer
}

func (h *myJSONHandler) Handle(ctx context.Context, r slog.Record) error {
	v := ctx.Value("request-id").(string)
	r.AddAttrs(slog.String("ctx", v))
	return h.Handler.Handle(ctx, r)
}

type Secret string

func (s Secret) LogValue() slog.Value {
	return slog.StringValue("********")
}
