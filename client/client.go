package client

import (
	"encoding/json"
	"errors"
	"net"

	"go.uber.org/zap"

	"word-of-wisdom/internal/logger"
	"word-of-wisdom/internal/messages"
	"word-of-wisdom/internal/network"
	"word-of-wisdom/internal/pow"
)

type Client struct {
	connection net.Conn
	pow        *pow.Manager
	logger     *zap.Logger
}

func New() (*Client, error) {
	log, err := logger.New()
	if err != nil {
		return nil, err
	}
	conn, err := net.Dial("tcp", network.GetAddress())
	if err != nil {
		return nil, err
	}
	return &Client{
		connection: conn,
		pow:        pow.NewManager(),
		logger:     log,
	}, nil
}

func (c *Client) Close() error {
	return c.connection.Close()
}

func (c *Client) Connect() error {
	log := c.logger.With(zap.String("remote address", c.connection.RemoteAddr().String()))
	if err := network.SendMessage(c.connection, &messages.Client{
		Type: messages.ClientTypeRequest,
	}); err != nil {
		return nil
	}
	log.Info("send start message")

	var challenge messages.Server
	if err := json.NewDecoder(c.connection).Decode(&challenge); err != nil {
		return err
	}
	if challenge.Type != messages.ServerTypeRequest {
		return errors.New("invalid request")
	}
	log.Info("get challenge message from server", zap.String("challenge", challenge.Data))

	err := network.SendMessage(c.connection, &messages.Client{
		Type:            messages.ClientTypeResponse,
		ChallengeAnswer: c.pow.Calculate(challenge.Data),
	})
	if err != nil {
		return nil
	}
	log.Info("find result and send to server")

	var result messages.Server
	if err := json.NewDecoder(c.connection).Decode(&result); err != nil {
		return err
	}
	if result.Type != messages.ServerTypeResponse {
		return errors.New("invalid request")
	}
	log.Info("get approve solution from server", zap.String("quote", result.Data))
	return nil
}
