package network

import (
	"encoding/json"
	"net"
	"os"
)

const (
	defaultHost = "127.0.0.1"
	defaultPort = "8081"
)

func GetAddress() string {
	host := os.Getenv("POW_SERVER_HOST")
	if host == "" {
		host = defaultHost
	}
	port := os.Getenv("POW_SERVER_PORT")
	if port == "" {
		port = defaultPort
	}
	return net.JoinHostPort(host, port)
}

func SendMessage(connection net.Conn, message interface{}) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	_, err = connection.Write(data)
	return err
}
