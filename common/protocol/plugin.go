package protocol

import (
	"net"
)

type Plugin interface {
	Name() string
	Version() string
}

type Protocol interface {
	Plugin

	Connect(conn net.Conn) (Connection, error)
}

type Connection interface {
	ID() string
	Handshake(username string, password string) error
	Read() (Packet, error)
	Write(packet Packet) error
	Close() error
}
