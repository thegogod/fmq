package main

import "github.com/thegogod/fmq/common/protocol"

var PublishQueue chan PublishEvent
var SubscribeQueue chan SubscribeEvent

type PublishEvent struct {
	Packet *protocol.Publish `json:"packet"`
}

type SubscribeEvent struct {
	Packet *protocol.Subscribe `json:"packet"`
	Conn   protocol.Connection `json:"conn"`
}
