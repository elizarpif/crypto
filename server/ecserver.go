package main

import (
	"context"
	cryptoapi "github.com/elizarpif/crypto/api"
	log "github.com/sirupsen/logrus"
)

type EcServer struct {
	logger *log.Logger
}

func (s *EcServer) SendMsg(ctx context.Context, request *cryptoapi.MsgRequest) (*cryptoapi.MsgResponse, error) {
	logger := s.logger

	text := request.Text
	logger.WithField("text", text).Info("received message")

	return &cryptoapi.MsgResponse{
		Text: encrypt(text),
	}, nil
}

const key = "secret key"

func encrypt(input string) string {
	output := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = input[i] ^ key[i%len(key)]
	}

	return string(output)
}

