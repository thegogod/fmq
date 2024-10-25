package main

import (
	"github.com/thegogod/fmq/common/protocol"
	"github.com/thegogod/fmq/plugins/mqtt"
)

var Plugins = map[string]protocol.Plugin{
	"mqtt": mqtt.New(),
}
