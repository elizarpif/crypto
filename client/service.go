package main

import (
	"context"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/local"
	"google.golang.org/grpc/credentials/oauth"

	cryptoapi "github.com/elizarpif/crypto/api"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	client cryptoapi.CryptoClient
	logger *log.Logger
}

func NewService(ctx context.Context, logger *log.Logger) (*Service, error) {
	creds := oauth.NewOauthAccess(&oauth2.Token{AccessToken: "access-token"})

	ts := local.NewCredentials()
	conn, err := grpc.DialContext(ctx, "localhost:443",
		grpc.WithPerRPCCredentials(creds),
		grpc.WithTransportCredentials(ts)) // grpc.WithStatsHandler(&ocgrpc.ClientHandler{})
	if err != nil {
		return nil, err
	}

	cl := cryptoapi.NewCryptoClient(conn)
	return &Service{client: cl, logger: logger}, nil
}

func (s *Service) CryptMsg(ctx context.Context) error {
	req := &cryptoapi.MsgRequest{Text: "hello server"}

	resp, err := s.client.SendMsg(ctx, req)
	if err != nil {
		s.logger.WithError(err).Error("cannot connect")
		return err
	}
	log.WithField("decrypt_msg", resp.Text).Info("decrypt")

	req.Text = resp.Text

	resp, err = s.client.SendMsg(ctx, req)
	if err != nil {
		s.logger.WithError(err).Error("cannot connect")
		return err
	}

	log.WithField("encrypt_msg", resp.Text).Info("encrypt")
	return nil
}
