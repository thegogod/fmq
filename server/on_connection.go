package main

import (
	"io"
	"net"
	"os"

	"github.com/thegogod/fmq/common/protocol"
	"github.com/thegogod/fmq/logger"
)

func onConnection(protocols []protocol.Protocol, conn net.Conn) {
	log := logger.New(conn.RemoteAddr().String())

	for _, plugin := range protocols {
		c, err := plugin.Connect(conn)

		defer func() {
			c.Close()
			log.Info("closed...")
		}()

		if err != nil {
			if err == io.EOF {
				return
			}

			log.Warn(err.Error())
			continue
		}

		if err := c.Handshake(os.Getenv("FMQ_USERNAME"), os.Getenv("FMQ_PASSWORD")); err != nil {
			log.Error(err.Error())
			return
		}

		log.Info("connected...")

		for {
			packet, err := c.Read()

			if err != nil {
				log.Error(err.Error())
				return
			}

			switch packet := packet.(type) {
			case *protocol.Ping:
				err = c.Write(&protocol.PingAck{})
				break
			case *protocol.PingAck:
				break
			case *protocol.Disconnect:
				return
			case *protocol.Publish:
				tasks.Push(onPublish(packet, c))
				break
			case *protocol.Subscribe:
				tasks.Push(onSubscribe(packet, c))
				break
			case *protocol.UnSubscribe:
				err = c.Write(&protocol.UnSubscribeAck{ID: packet.ID})
				break
			}

			if err != nil {
				log.Error(err.Error())
				return
			}
		}
	}
}

func onPublish(packet *protocol.Publish, conn protocol.Connection) func() error {
	return func() error {
		next := router.Next(packet.Topic)

		if next != nil {
			next.Write(packet)
			return conn.Write(&protocol.PublishAck{ID: packet.ID})
		}

		return nil
	}
}

func onSubscribe(packet *protocol.Subscribe, conn protocol.Connection) func() error {
	return func() error {
		for _, topic := range packet.Topics {
			router.On(topic, conn)
		}

		return conn.Write(&protocol.SubscribeAck{ID: packet.ID})
	}
}
