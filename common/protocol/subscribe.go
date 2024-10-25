package protocol

import "encoding/json"

type Subscribe struct {
	ID     uint16   `json:"id"`
	Topics []string `json:"topics"`
}

func (self Subscribe) Code() Code {
	return SUBSCRIBE
}

func (self Subscribe) String() string {
	b, _ := json.MarshalIndent(self, "", " ")
	return string(b)
}

type SubscribeAck struct {
	ID          uint16 `json:"id"`
	ReturnCodes []byte `json:"return_codes"`
}

func (self SubscribeAck) Code() Code {
	return SUBSCRIBE_ACK
}

func (self SubscribeAck) String() string {
	b, _ := json.MarshalIndent(self, "", " ")
	return string(b)
}
