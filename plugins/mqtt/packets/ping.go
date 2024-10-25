package packets

import (
	"encoding/json"
	"io"

	"github.com/thegogod/fmq/common/protocol"
)

type Ping struct {
	Header `json:"header"`
}

func (self Ping) ToProtocol() protocol.Packet {
	return &protocol.Ping{}
}

func (self *Ping) FromProtocol(packet *protocol.Ping) *Ping {
	self.Header = Header{Code: PING}
	return self
}

func (self *Ping) Read(reader io.Reader) error {
	return nil
}

func (self *Ping) Write(writer io.Writer) error {
	return self.Header.Write(writer)
}

func (self Ping) String() string {
	b, _ := json.MarshalIndent(self, "", " ")
	return string(b)
}
