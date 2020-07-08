package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	logger := log.New()

	service, err := NewService(ctx, logger)
	if err != nil {
		logger.WithError(err).Fatal()
	}
	group := errgroup.Group{}

	group.Go(func() error {
		return service.CryptMsg(ctx)
	})
	go func() {
		signalListener := make(chan os.Signal, 1)
		defer close(signalListener)

		signal.Notify(signalListener,
			syscall.SIGHUP,
			syscall.SIGINT,
			syscall.SIGQUIT)

		stop := <-signalListener
		logger.Info("Received", stop)
		logger.Fatal()
	}()

	err = group.Wait()
	if err != nil {
		logger.Fatal()
	}
}
