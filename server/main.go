package main

import (
	"context"
	"errors"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	cryptoapi "github.com/elizarpif/crypto/api"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	grpcAddress = "localhost:443"
)

func main() {
	//ctx := context.Background()
	logger := log.New()
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		logger.WithError(err).WithField("address", grpcAddress).Fatal("listen for grpc")
	}

	opts := []grpc.ServerOption{grpc.UnaryInterceptor(validateToken)}
	grpcServer := grpc.NewServer(opts...)
	defer grpcServer.GracefulStop()

	cryptoapi.RegisterCryptoServer(grpcServer, &EcServer{logger: logger})

	group := errgroup.Group{}
	group.Go(func() error {
		logger.WithField("address", grpcAddress).Info("start grpc server")
		return grpcServer.Serve(lis)
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

func validateToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("missing metadata")
	}
	token := strings.TrimPrefix(md["authorization"][0], "Bearer ")
	if token != "access-token" {
		return nil, errors.New("not expected token")
	}
	return handler(ctx, req)
}
