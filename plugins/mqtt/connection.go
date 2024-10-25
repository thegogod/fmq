package mqtt

import (
	"bufio"
	"fmt"
	"log/slog"
	"net"
	"sync"
	"time"

	"github.com/thegogod/fmq/common/protocol"
	"github.com/thegogod/fmq/logger"
	"github.com/thegogod/fmq/plugins/mqtt/packets"
)

type Connection struct {
	conn      net.Conn
	log       *slog.Logger
	reader    *bufio.Reader
	timeout   *time.Timer
	keepalive time.Duration
	mu        sync.RWMutex
}

func newConnection(conn net.Conn) *Connection {
	return &Connection{
		conn:   conn,
		log:    logger.New("mqtt"),
		reader: bufio.NewReader(conn),
		mu:     sync.RWMutex{},
	}
}

func (self *Connection) ID() string {
	return self.conn.RemoteAddr().String()
}

func (self *Connection) Handshake(username string, password string) error {
	self.mu.Lock()
	defer self.mu.Unlock()

	packet, err := packets.Read(self.conn)

	if err != nil {
		return err
	}

	switch p := packet.(type) {
	case *packets.Connect:
		res := &packets.ConnectAck{
			Header:     packets.Header{Code: packets.CONNECT_ACK},
			ReturnCode: packets.Accepted,
		}

		if p.Username != username || string(p.Password) != password {
			res.ReturnCode = packets.ErrRefusedNotAuthorised
		}

		self.keepalive = time.Duration(int64(p.Keepalive+(p.Keepalive/2)) * int64(time.Second))

		if err := res.Write(self.conn); err != nil {
			return err
		}

		self.timeout = time.AfterFunc(self.keepalive, func() {
			self.log.Warn("closing due to inactivity...")
			self.Close()
		})

		return nil
	}

	return fmt.Errorf("expected Connect packet, received 0x%x", packet.ToProtocol().Code())
}

func (self *Connection) Read() (protocol.Packet, error) {
	self.mu.RLock()
	defer self.mu.RUnlock()

	packet, err := packets.Read(self.conn)

	if err != nil {
		return nil, err
	}

	self.timeout.Reset(self.keepalive)
	return packet.ToProtocol(), nil
}

func (self *Connection) Write(packet protocol.Packet) error {
	self.mu.RLock()
	defer self.mu.RUnlock()

	p, err := packets.FromProtocol(packet)

	if err != nil {
		return err
	}

	return p.Write(self.conn)
}

func (self *Connection) Close() error {
	return self.conn.Close()
}
