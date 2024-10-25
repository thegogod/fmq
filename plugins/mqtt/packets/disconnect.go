package packets

import (
	"encoding/json"
	"io"

	"github.com/thegogod/fmq/common/protocol"
)

type Disconnect struct {
	Header `json:"header"`
}

func (self Disconnect) ToProtocol() protocol.Packet {
	return &protocol.Disconnect{}
}

func (self *Disconnect) FromProtocol(packet *protocol.Disconnect) *Disconnect {
	self.Header = Header{Code: DISCONNECT}
	return self
}

func (self *Disconnect) Read(reader io.Reader) error {
	return nil
}

func (self *Disconnect) Write(writer io.Writer) error {
	return self.Header.Write(writer)
}

func (self Disconnect) String() string {
	b, _ := json.MarshalIndent(self, "", " ")
	return string(b)
}
