package packets

import (
	"bytes"
	"encoding/json"
	"io"
	"slices"

	"github.com/thegogod/fmq/common/protocol"
)

type SubscribeAck struct {
	Header `json:"header"`

	MessageID   uint16 `json:"message_id"`
	ReturnCodes []byte `json:"return_codes"`
}

func (self SubscribeAck) ToProtocol() protocol.Packet {
	return &protocol.SubscribeAck{
		ID:          self.MessageID,
		ReturnCodes: self.ReturnCodes,
	}
}

func (self *SubscribeAck) FromProtocol(packet *protocol.SubscribeAck) *SubscribeAck {
	self.Header = Header{Code: SUBSCRIBE_ACK}
	self.MessageID = packet.ID
	self.ReturnCodes = packet.ReturnCodes
	return self
}

func (self *SubscribeAck) Read(reader io.Reader) error {
	var qosBuffer bytes.Buffer
	var err error

	self.MessageID, err = decodeUint16(reader)

	if err != nil {
		return err
	}

	_, err = qosBuffer.ReadFrom(reader)

	if err != nil {
		return err
	}

	self.ReturnCodes = qosBuffer.Bytes()
	return nil
}

func (self *SubscribeAck) Write(writer io.Writer) error {
	var header bytes.Buffer
	var body bytes.Buffer
	var err error

	body.Write(encodeUint16(self.MessageID))
	body.Write(self.ReturnCodes)

	self.Header.RemainingLength = body.Len()

	if err := self.Header.Write(&header); err != nil {
		return err
	}

	_, err = writer.Write(slices.Concat(
		header.Bytes(),
		body.Bytes(),
	))

	return err
}

func (self SubscribeAck) String() string {
	b, _ := json.MarshalIndent(self, "", " ")
	return string(b)
}
