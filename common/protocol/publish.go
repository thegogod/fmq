package protocol

import "encoding/json"

type Publish struct {
	ID      uint16 `json:"id"`
	Topic   string `json:"topic"`
	Payload []byte `json:"payload"`
}

func (self Publish) Code() Code {
	return PUBLISH
}

func (self Publish) String() string {
	b, _ := json.MarshalIndent(self, "", " ")
	return string(b)
}

type PublishAck struct {
	ID uint16 `json:"id"`
}

func (self PublishAck) Code() Code {
	return PUBLISH_ACK
}

func (self PublishAck) String() string {
	b, _ := json.MarshalIndent(self, "", " ")
	return string(b)
}
