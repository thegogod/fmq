package protocol

import "encoding/json"

type Disconnect struct{}

func (self Disconnect) Code() Code {
	return DISCONNECT
}

func (self Disconnect) String() string {
	b, _ := json.MarshalIndent(self, "", " ")
	return string(b)
}
