package main

import (
	"fmt"
	"log/slog"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/thegogod/fmq/async"
	"github.com/thegogod/fmq/common/env"
	"github.com/thegogod/fmq/common/protocol"
	"github.com/thegogod/fmq/logger"
)

var publish = make(chan Event[*protocol.Publish], 100000)
var subscribe = make(chan Event[*protocol.Subscribe], 100000)
var unSubscribe = make(chan Event[*protocol.UnSubscribe], 100000)

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
	workers := async.New(500)
	workers.Start()

	for i := 0; i < workers.Count(); i++ {
		workers.Push(listen(log, topics))
	}

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

		client := NewClient(c)
		go client.Listen(
			os.Getenv("FMQ_USERNAME"),
			os.Getenv("FMQ_PASSWORD"),
		)
	}
}

func listen(_ *slog.Logger, topics *Topics) func() error {
	return func() error {
		for {
			select {
			case event := <-subscribe:
				for _, topic := range event.Packet.Topics {
					topics.Subscribe(topic, event.From)
				}
			case event := <-unSubscribe:
				for _, topic := range event.Packet.Topics {
					topics.UnSubscribe(topic, event.From.ID())
				}
			case event := <-publish:
				topics.Publish(event.Packet.Topic, event.Packet)
			default:
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}
