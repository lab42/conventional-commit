package logger

import (
	"os"

	"golang.org/x/exp/slog"
)

var Log *slog.Logger

func init() {
	Log = slog.New(slog.NewTextHandler(os.Stdout))
}
