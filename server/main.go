package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/thegogod/fmq/common/env"
	"github.com/thegogod/fmq/common/protocol"
	"github.com/thegogod/fmq/logger"
)

var publish = make(chan Event[*protocol.Publish])
var subscribe = make(chan Event[*protocol.Subscribe])
var unSubscribe = make(chan Event[*protocol.UnSubscribe])

func main() {
	log := logger.New("main")
	port, err := strconv.Atoi(env.GetOrDefault("FMQ_PORT", "1883"))

	if err != nil {
		panic(fmt.Errorf("`FMQ_PORT` must be an integer"))
	}

	protocolName := strings.ToLower(env.GetOrDefault("FMQ_PROTOCOL", "mqtt"))
	plugin, exists := Plugins[protocolName]

	if !exists {
		panic(fmt.Errorf("`FMQ_PROTOCOL` value `%s` is not a valid plugin", protocolName))
	}

	proto, ok := plugin.(protocol.Protocol)

	if !ok {
		panic(fmt.Errorf("`FMQ_PROTOCOL` value `%s` is not a valid protocol", protocolName))
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		panic(err)
	}

	log.Info(fmt.Sprintf("listening on port %d...", port))
	topics := newTopics()
	clients := map[string]*Client{}

	for {
		conn, err := listener.Accept()

		if err != nil {
			panic(err)
		}

		c, err := proto.Connect(conn)

		if err != nil {
			conn.Close()
			continue
		}

		client := NewClient(
			c,
			func() {
				delete(clients, c.ID())
			},
			func(event Event[*protocol.Publish]) {
				topics.Publish(event.Packet.Topic, event.Packet)
			},
			func(event Event[*protocol.Subscribe]) {
				for _, topic := range event.Packet.Topics {
					topics.Subscribe(topic, c)
				}
			},
			func(event Event[*protocol.UnSubscribe]) {
				for _, topic := range event.Packet.Topics {
					topics.UnSubscribe(topic, c.ID())
				}
			},
		)

		clients[client.ID()] = client
		go client.Listen(
			os.Getenv("FMQ_USERNAME"),
			os.Getenv("FMQ_PASSWORD"),
		)
	}
}
