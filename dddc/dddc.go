package main

import (
	"bytes"
	"dddc/config"
	"dddl/bundling"
	"dddl/web"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"os"
	"strings"

	"dddl/deployment"

	"github.com/BurntSushi/toml"
)

var buildVersionString string

// ErrInvalidKeyLength is returned if the length of the provided symmetric key
// is not 128, 192 or 256 bits.
var ErrInvalidKeyLength = errors.New("the specified key length is invalid, it must be one of 128, 192, or 256 bits in length")

func main() {
	fmt.Printf("dddc (%v)\n", buildVersionString)

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) != 3 {
		return errors.New("usage: dddc <hostname> <deployment-file>")
	}

	config, err := config.Load()
	if err != nil {
		return err
	}

	hostname := os.Args[1]
	deploymentFile := os.Args[2]

	deploymentPayload := deployment.Payload{}
	_, err = toml.DecodeFile(deploymentFile, &deploymentPayload)
	if err != nil {
		return err
	}

	for _, v := range deploymentPayload {
		if strings.HasPrefix(v.Tag, "env:") {
			v.Tag = os.Getenv(v.Tag[4:])
		}
	}

	var bundler bundling.MessageBundler
	if config.SymmetricKey != "" {
		key, err := hex.DecodeString(config.SymmetricKey)
		if err != nil {
			return err
		}

		if len(key) != 16 && len(key) != 24 && len(key) != 32 {
			return ErrInvalidKeyLength
		}

		bundler = bundling.NewSymmetricEncryptionBundler(key)
	} else {
		bundler = bundling.NewLengthPrefixedBundler()
	}

	conn, err := net.Dial("tcp", hostname)
	if err != nil {
		return err
	}

	socket := web.NewSocket(conn, bundler)
	defer socket.Close()

	buf := &bytes.Buffer{}
	encoder := toml.NewEncoder(buf)
	err = encoder.Encode(deploymentPayload)
	if err != nil {
		return err
	}

	rawBuf := buf.Bytes()
	socket.Send(rawBuf)
	fmt.Printf("Sent %v bytes of deployment information to server...\n", len(rawBuf))
	msg, err := socket.Receive()
	if err != nil {
		return err
	}

	if len(msg) != 1 || msg[0] != 1 {
		return fmt.Errorf("server-side deployment error: %v", string(msg))
	}

	fmt.Printf("-- Successful Deployment --\n\n")

	return nil
}
