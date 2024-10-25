package mqtt

import (
	"net"

	"github.com/thegogod/fmq/common/protocol"
)

type Plugin struct {
}

func New() *Plugin {
	return &Plugin{}
}

func (self *Plugin) Name() string {
	return "mqtt"
}

func (self *Plugin) Version() string {
	return "0.0.0"
}

func (self *Plugin) Connect(conn net.Conn) (protocol.Connection, error) {
	return newConnection(conn), nil
}
