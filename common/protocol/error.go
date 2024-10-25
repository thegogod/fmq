package protocol

import (
	"encoding/json"
)

type Error struct {
	Message string `json:"message"`
}

func (self Error) Code() Code {
	return ERROR
}

func (self Error) String() string {
	b, _ := json.MarshalIndent(self, "", " ")
	return string(b)
}
