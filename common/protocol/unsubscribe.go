package protocol

import "encoding/json"

type UnSubscribe struct {
	ID     uint16   `json:"id"`
	Topics []string `json:"topics"`
}

func (self UnSubscribe) Code() Code {
	return UNSUBSCRIBE
}

func (self UnSubscribe) String() string {
	b, _ := json.MarshalIndent(self, "", " ")
	return string(b)
}

type UnSubscribeAck struct {
	ID uint16 `json:"id"`
}

func (self UnSubscribeAck) Code() Code {
	return UNSUBSCRIBE_ACK
}

func (self UnSubscribeAck) String() string {
	b, _ := json.MarshalIndent(self, "", " ")
	return string(b)
}
