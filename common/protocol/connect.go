package protocol

import "encoding/json"

type ConnectReturnCode byte

type Connect struct {
	ClientID string `json:"client_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (self Connect) Code() Code {
	return CONNECT
}

func (self Connect) String() string {
	b, _ := json.MarshalIndent(self, "", " ")
	return string(b)
}

type ConnectAck struct {
	ReturnCode byte `json:"return_code"`
}

func (self ConnectAck) Code() Code {
	return CONNECT_ACK
}

func (self ConnectAck) String() string {
	b, _ := json.MarshalIndent(self, "", " ")
	return string(b)
}
