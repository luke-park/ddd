package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"os"

	"ddds/config"
	"ddds/handler"
	"ddds/templating"
)

var buildVersionString string

func main() {
	fmt.Printf("ddds (%v)\n", buildVersionString)

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() error {
	config, err := config.Load()
	if err != nil {
		return err
	}

	err = templating.Setup()
	if err != nil {
		return err
	}

	if config.SymmetricKey != "" {
		rawKey, err := hex.DecodeString(config.SymmetricKey)
		if err != nil {
			return err
		}

		err = handler.SetSymmetricKey(rawKey)
		if err != nil {
			return err
		}
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", config.Port))
	if err != nil {
		return err
	}

	acceptFailureCount := 0

	for {
		conn, err := l.Accept()
		if err != nil {
			acceptFailureCount++
			if acceptFailureCount == 5 {
				return err
			}
		}

		acceptFailureCount = 0
		go handler.Handle(conn)
	}

	return nil
}
