package protocol

import "encoding/json"

type Ping struct{}

func (self Ping) Code() Code {
	return PING
}

func (self Ping) String() string {
	b, _ := json.MarshalIndent(self, "", " ")
	return string(b)
}

type PingAck struct{}

func (self PingAck) Code() Code {
	return PING_ACK
}

func (self PingAck) String() string {
	b, _ := json.MarshalIndent(self, "", " ")
	return string(b)
}
