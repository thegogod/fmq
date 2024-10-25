package packets

import (
	"bytes"
	"encoding/json"
	"io"
	"slices"

	"github.com/thegogod/fmq/common/protocol"
)

type PublishAck struct {
	Header `json:"header"`

	MessageID uint16 `json:"message_id"`
}

func (self PublishAck) ToProtocol() protocol.Packet {
	return &protocol.PublishAck{ID: self.MessageID}
}

func (self *PublishAck) FromProtocol(packet *protocol.PublishAck) *PublishAck {
	self.Header = Header{Code: PUBLISH_ACK}
	self.MessageID = packet.ID
	return self
}

func (self *PublishAck) Read(reader io.Reader) error {
	messageId, err := decodeUint16(reader)

	if err != nil {
		return err
	}

	self.MessageID = messageId
	return nil
}

func (self *PublishAck) Write(writer io.Writer) error {
	var header bytes.Buffer
	var body bytes.Buffer

	body.Write(encodeUint16(self.MessageID))

	self.Header.RemainingLength = 2

	if err := self.Header.Write(&header); err != nil {
		return err
	}

	_, err := writer.Write(slices.Concat(
		header.Bytes(),
		body.Bytes(),
	))

	return err
}

func (self PublishAck) String() string {
	b, _ := json.MarshalIndent(self, "", " ")
	return string(b)
}
