package main

import (
	"io"
	"log/slog"

	"github.com/thegogod/fmq/common/protocol"
	"github.com/thegogod/fmq/logger"
)

type Event[T protocol.Packet] struct {
	Packet T
	From   protocol.Connection
}

type Client struct {
	log  *slog.Logger
	conn protocol.Connection

	onClose       func()
	onPublish     func(event Event[*protocol.Publish])
	onSubscrube   func(event Event[*protocol.Subscribe])
	onUnSubscribe func(event Event[*protocol.UnSubscribe])
}

func NewClient(
	conn protocol.Connection,
	onClose func(),
	onPublish func(event Event[*protocol.Publish]),
	onSubscrube func(event Event[*protocol.Subscribe]),
	onUnSubscribe func(event Event[*protocol.UnSubscribe]),
) *Client {
	return &Client{
		log:  logger.New(conn.ID()),
		conn: conn,

		onClose:       onClose,
		onPublish:     onPublish,
		onSubscrube:   onSubscrube,
		onUnSubscribe: onUnSubscribe,
	}
}

func (self *Client) ID() string {
	return self.conn.ID()
}

func (self *Client) Listen(username string, password string) error {
	defer func() {
		self.conn.Close()
		self.log.Debug("closed...")
		self.onClose()
	}()

	if err := self.conn.Handshake(username, password); err != nil {
		self.log.Error(err.Error())
		return err
	}

	self.log.Info("connected...")

	for {
		packet, err := self.conn.Read()

		if err != nil {
			// connection closed
			if err == io.EOF {
				return nil
			}

			self.log.Error(err.Error())
			return err
		}

		switch packet := packet.(type) {
		case *protocol.Ping:
			err = self.conn.Write(&protocol.PingAck{})
			break
		case *protocol.Disconnect:
			return nil
		case *protocol.Publish:
			self.onPublish(Event[*protocol.Publish]{
				Packet: packet,
				From:   self.conn,
			})

			if packet.Qos == 1 {
				err = self.conn.Write(&protocol.PublishAck{ID: packet.ID})
			}

			break
		case *protocol.Subscribe:
			self.onSubscrube(Event[*protocol.Subscribe]{
				Packet: packet,
				From:   self.conn,
			})

			err = self.conn.Write(&protocol.SubscribeAck{ID: packet.ID})
			break
		case *protocol.UnSubscribe:
			self.onUnSubscribe(Event[*protocol.UnSubscribe]{
				Packet: packet,
				From:   self.conn,
			})

			err = self.conn.Write(&protocol.UnSubscribeAck{ID: packet.ID})
			break
		}

		if err != nil {
			self.log.Error(err.Error())
			return err
		}
	}
}
