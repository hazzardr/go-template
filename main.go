package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/hazzardr/go-template/cmd/cli"
)

func main() {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
	})
	slog.SetDefault(slog.New(logger))

	if err := cli.Execute(); err != nil {
		slog.Error("command failed", "error", err)
		os.Exit(1)
	}
}
