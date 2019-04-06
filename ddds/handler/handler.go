package handler

import (
	"ddds/commands"
	"ddds/templating"
	"errors"
	"net"
	"time"

	"dddl/bundling"
	"dddl/deployment"
	"dddl/web"

	"github.com/BurntSushi/toml"
)

// ReadTimeout is how long a handler will wait for data before it disconnects a
// socket.
const ReadTimeout = time.Second * 3

var symmetricKey []byte

// ErrInvalidKeyLength is returned if the length of the provided symmetric key
// is not 128, 192 or 256 bits.
var ErrInvalidKeyLength = errors.New("the specified key length is invalid, it must be one of 128, 192, or 256 bits in length")

// SetSymmetricKey sets the symmetric key used for transport security.
func SetSymmetricKey(key []byte) error {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return ErrInvalidKeyLength
	}

	symmetricKey = key
	return nil
}

// Handle handles a newly connected TCP connection.
func Handle(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(ReadTimeout))
	socket := web.NewSocket(conn, getBundler())
	defer socket.Close()

	msg, err := socket.Receive()
	if err != nil {
		return
	}

	payload := deployment.Payload{}
	err = toml.Unmarshal(msg, &payload)
	if err != nil {
		socket.SendError(err)
		return
	}

	err = templating.Perform(payload)
	if err != nil {
		socket.SendError(err)
		return
	}

	err = commands.RestartDocker()
	if err != nil {
		socket.SendError(err)
		return
	}

	socket.Send([]byte{1})
}

// getBundler returns a new bundler for a new socket.
func getBundler() bundling.MessageBundler {
	if len(symmetricKey) == 0 {
		return bundling.NewLengthPrefixedBundler()
	}

	return bundling.NewSymmetricEncryptionBundler(symmetricKey)
}
