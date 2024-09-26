package util

import (
	"os"
	"os/signal"
	"syscall"
)

func MonitorStopSignal() <-chan os.Signal {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	return signalCh
}
