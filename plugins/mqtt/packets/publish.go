package packets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"slices"

	"github.com/thegogod/fmq/common/protocol"
)

type Publish struct {
	Header `json:"header"`

	Topic     string `json:"topic"`
	MessageID uint16 `json:"message_id"`
	Payload   []byte `json:"payload"`
}

func (self Publish) ToProtocol() protocol.Packet {
	return &protocol.Publish{
		ID:      self.MessageID,
		Qos:     self.Header.Qos,
		Topic:   self.Topic,
		Payload: self.Payload,
	}
}

func (self *Publish) FromProtocol(packet *protocol.Publish) *Publish {
	self.Header = Header{Code: PUBLISH}
	self.Topic = packet.Topic
	self.MessageID = packet.ID
	self.Payload = packet.Payload
	return self
}

func (self *Publish) Read(reader io.Reader) error {
	var payloadLength = self.Header.RemainingLength
	var err error

	self.Topic, err = decodeString(reader)

	if err != nil {
		return err
	}

	if self.Qos > 0 {
		self.MessageID, err = decodeUint16(reader)

		if err != nil {
			return err
		}

		payloadLength -= len(self.Topic) + 4
	} else {
		payloadLength -= len(self.Topic) + 2
	}

	if payloadLength < 0 {
		return fmt.Errorf("error unpacking publish, payload length < 0")
	}

	self.Payload = make([]byte, payloadLength)
	_, err = reader.Read(self.Payload)

	return err
}

func (self *Publish) Write(writer io.Writer) error {
	var header bytes.Buffer
	var body bytes.Buffer
	var err error

	body.Write(encodeString(self.Topic))

	if self.Qos > 0 {
		body.Write(encodeUint16(self.MessageID))
	}

	self.Header.RemainingLength = body.Len() + len(self.Payload)

	if err := self.Header.Write(&header); err != nil {
		return err
	}

	_, err = writer.Write(slices.Concat(
		header.Bytes(),
		body.Bytes(),
		self.Payload,
	))

	return err
}

func (self Publish) String() string {
	b, _ := json.MarshalIndent(self, "", " ")
	return string(b)
}
