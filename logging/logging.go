package logging

import (
	"os"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("Gospirational")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func init() {
	backend1 := logging.NewLogBackend(os.Stdout, "", 0)
	backend1Formatted := logging.NewBackendFormatter(backend1, format)
	logging.SetBackend(backend1Formatted)
}

// Log returns the current logger
func Log() *logging.Logger {
	return log
}
