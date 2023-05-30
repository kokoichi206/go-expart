package main

func main() {
	// // sloger is a wrapper of slog, which is a structured logger
	// // in official golang.org/x/exp repo (still experimental)
	// // see: https://pkg.go.dev/golang.org/x/exp/slog
	// logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
	// 	Level: slog.LevelInfo,
	// }))
	// logger.Info("hello world", slog.String("foo", "bar"))
	// logger.Debug("hello world", slog.String("foo", "bar"))

	words()
}
