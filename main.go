package main

import (
	"os"
	"os/signal"
	"syscall"
	"worker-golang/cmd/jobs"
)

func main() {
	interrupt := make(chan os.Signal, 1)

	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	defer signal.Stop(interrupt)

	jobs := jobs.New()

	select {
	case <-interrupt:
		go jobs.Hello.Stop()
		os.Exit(0)
	}
}
