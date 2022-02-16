package main

import (
	"fmt"
	"os"

	"word-of-wisdom/client"
)

func main() {
	cli, err := client.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "create client failed: %s\n", err)
		os.Exit(1)
	}
	if err := cli.Connect(); err != nil {
		fmt.Fprintf(os.Stderr, "something going wrong: %s\n", err)
		cli.Close()
		os.Exit(1)
	}
	cli.Close()
}
