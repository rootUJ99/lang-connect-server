package main

import "log/slog"

func main() {
	slog.Default()
	slog.Info("starting the extractor server")
}
