package main

import (
	"fmt"
	"os"

	"word-of-wisdom/internal/quotes"
	"word-of-wisdom/server"
)

func main() {
	quotes, err := quotes.NewQuotes()
	if err != nil {
		fmt.Fprintf(os.Stderr, "get quotes failed: %s\n", err)
		os.Exit(1)
	}
	serv, err := server.New(quotes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create server failed: %s\n", err)
		os.Exit(1)
	}
	defer serv.Close()
	if err := serv.Listen(); err != nil {
		fmt.Fprintf(os.Stderr, "something going wrong: %s\n", err)
		os.Exit(1)
	}
}
