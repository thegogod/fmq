package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/thegogod/fmq/async"
	"github.com/thegogod/fmq/common/env"
	"github.com/thegogod/fmq/common/protocol"
	"github.com/thegogod/fmq/common/slices"
	"github.com/thegogod/fmq/logger"
)

type Message struct {
	ID      uint16 `json:"id"`
	Payload []byte `json:"payload"`
}

var topics = map[string]chan Message{}
var tasks = async.New(50)
var router = NewRouter()

func main() {
	log := logger.New("")
	port, err := strconv.Atoi(env.GetOrDefault("FMQ_PORT", "1883"))

	if err != nil {
		panic(fmt.Errorf("`FMQ_PORT` must be an integer"))
	}

	plugins := slices.Map(slices.Filter(strings.Split(os.Getenv("FMQ_PLUGINS"), ","), func(name string) bool {
		_, exists := Plugins[name]
		return exists
	}), func(name string) protocol.Plugin {
		return Plugins[name]
	})

	protocols := slices.Map(slices.Filter(plugins, func(p protocol.Plugin) bool {
		_, ok := p.(protocol.Protocol)
		return ok
	}), func(p protocol.Plugin) protocol.Protocol {
		return p.(protocol.Protocol)
	})

	if len(protocols) == 0 {
		panic(fmt.Errorf("must enable at least 1 protocol plugin"))
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		panic(err)
	}

	log.Info(fmt.Sprintf("listening on port %d...", port))
	tasks.Start()

	for {
		conn, err := listener.Accept()

		if err != nil {
			panic(err)
		}

		go onConnection(protocols, conn)
	}
}
