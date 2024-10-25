package packets

import (
	"encoding/json"
	"io"

	"github.com/thegogod/fmq/common/protocol"
)

type PingAck struct {
	Header `json:"header"`
}

func (self PingAck) ToProtocol() protocol.Packet {
	return &protocol.PingAck{}
}

func (self *PingAck) FromProtocol(packet *protocol.PingAck) *PingAck {
	self.Header = Header{Code: PING_ACK}
	return self
}

func (self *PingAck) Read(reader io.Reader) error {
	return nil
}

func (self *PingAck) Write(writer io.Writer) error {
	return self.Header.Write(writer)
}

func (self PingAck) String() string {
	b, _ := json.MarshalIndent(self, "", " ")
	return string(b)
}
