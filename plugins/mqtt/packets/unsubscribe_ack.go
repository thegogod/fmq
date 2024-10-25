package packets

import (
	"bytes"
	"encoding/json"
	"io"
	"slices"

	"github.com/thegogod/fmq/common/protocol"
)

type UnSubscribeAck struct {
	Header `json:"header"`

	MessageID uint16 `json:"message_id"`
}

func (self UnSubscribeAck) ToProtocol() protocol.Packet {
	return &protocol.UnSubscribeAck{ID: self.MessageID}
}

func (self *UnSubscribeAck) FromProtocol(packet *protocol.UnSubscribeAck) *UnSubscribeAck {
	self.Header = Header{Code: UNSUBSCRIBE_ACK}
	self.MessageID = packet.ID
	return self
}

func (self *UnSubscribeAck) Read(reader io.Reader) error {
	messageId, err := decodeUint16(reader)

	if err != nil {
		return err
	}

	self.MessageID = messageId
	return nil
}

func (self *UnSubscribeAck) Write(writer io.Writer) error {
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

func (self UnSubscribeAck) String() string {
	b, _ := json.MarshalIndent(self, "", " ")
	return string(b)
}
