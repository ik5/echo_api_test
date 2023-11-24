package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func signalling() {
	chanQuitSig := make(chan os.Signal, 1)
	chanLogSig := make(chan os.Signal, 1)

	signal.Notify(
		chanQuitSig, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT,
		syscall.SIGABRT,
	)

	signal.Notify(chanLogSig, syscall.SIGHUP)

	infinite := true

	for infinite {
		select {
		case sig := <-chanQuitSig:
			if logger != nil {
				logger.Debug("Had quit signal", slog.Any("signal", sig))
			}

			infinite = false
			quit <- true
		case sig := <-chanLogSig:
			if logger != nil {
				logger.Debug("Had log signal", slog.Any("signal", sig))
			}
		}
	}
}
