package main

import (
	"context"
	"fr33d0mz/moneyflowx"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	srv := new(moneyflowx.Server)

	go func() {
		if err := srv.Run("", nil); err != nil {
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		return
	}
}
