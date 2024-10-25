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

			log.Debug(packet.String())

			switch packet := packet.(type) {
			case *protocol.Ping:
				err = c.Write(&protocol.PingAck{})
				break
			case *protocol.PingAck:
				break
			case *protocol.Disconnect:
				return
			case *protocol.Publish:
				PublishQueue <- PublishEvent{packet}
				err = c.Write(&protocol.PublishAck{ID: packet.ID})
				break
			case *protocol.Subscribe:
				SubscribeQueue <- SubscribeEvent{packet, c}
				err = c.Write(&protocol.SubscribeAck{ID: packet.ID})
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
