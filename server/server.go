package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"word-of-wisdom/internal/logger"
	"word-of-wisdom/internal/messages"
	"word-of-wisdom/internal/network"
	"word-of-wisdom/internal/pow"
	"word-of-wisdom/internal/quotes"
)

type Server struct {
	listener net.Listener
	quotes   *quotes.Quotes
	pow      *pow.Manager
	logger   *zap.Logger
}

func New(quotes *quotes.Quotes) (*Server, error) {
	log, err := logger.New()
	if err != nil {
		return nil, err
	}
	listener, err := net.Listen("tcp", network.GetAddress())
	if err != nil {
		return nil, err
	}
	return &Server{
		listener: listener,
		quotes:   quotes,
		pow:      pow.NewManager(),
		logger:   log,
	}, nil
}

func (s *Server) Close() error {
	return s.listener.Close()
}

func (s *Server) Listen() error {
	for {
		connection, err := s.listener.Accept()
		if err != nil {
			return err
		}
		s.logger.Info("got new connection", zap.String("remote address", connection.RemoteAddr().String()))
		go s.HandleConnection(connection)
	}
}

func (s *Server) HandleConnection(connection net.Conn) {
	if err := s.handleConnection(connection); err != nil {
		s.logger.Error("something going wrong", zap.Error(err))
	}
	connection.Close()
}

func (s *Server) handleConnection(connection net.Conn) error {
	var (
		message messages.Client
		log     = s.logger.With(zap.String("remote address", connection.RemoteAddr().String()))
	)
	if err := json.NewDecoder(connection).Decode(&message); err != nil {
		return err
	}
	if message.Type != messages.ClientTypeRequest {
		return errors.New("invalid request")
	}
	log.Info("got start message from client")

	challenge := uuid.New()
	err := network.SendMessage(connection, &messages.Server{
		Type: messages.ServerTypeRequest,
		Data: challenge.String(),
	})
	if err != nil {
		return err
	}
	log.Info("send challenge to client")

	if err := json.NewDecoder(connection).Decode(&message); err != nil {
		return err
	}
	if message.Type != messages.ClientTypeResponse {
		return errors.New("invalid request")
	}
	if err := s.pow.Validate(challenge.String(), message.ChallengeAnswer); err != nil {
		return err
	}
	log.Info("got correct solution from client")

	quote := s.quotes.GetRandomQuote()
	return network.SendMessage(connection, &messages.Server{
		Type: messages.ServerTypeResponse,
		Data: fmt.Sprintf("%s (c) %s", quote.Text, quote.Author),
	})
}
