package packets

import (
	"bytes"
	"encoding/json"
	"io"
	"slices"

	"github.com/thegogod/fmq/common/protocol"
)

type UnSubscribe struct {
	Header `json:"header"`

	MessageID uint16   `json:"message_id"`
	Topics    []string `json:"topics"`
}

func (self UnSubscribe) ToProtocol() protocol.Packet {
	return &protocol.UnSubscribe{
		ID:     self.MessageID,
		Topics: self.Topics,
	}
}

func (self *UnSubscribe) FromProtocol(packet *protocol.UnSubscribe) *UnSubscribe {
	self.Header = Header{Code: UNSUBSCRIBE}
	self.MessageID = packet.ID
	self.Topics = packet.Topics
	return self
}

func (self *UnSubscribe) Read(reader io.Reader) error {
	var err error

	self.MessageID, err = decodeUint16(reader)

	if err != nil {
		return err
	}

	topic, err := decodeString(reader)

	for err == nil && topic != "" {
		topic, err = decodeString(reader)

		if err != nil {
			return err
		}

		self.Topics = append(self.Topics, topic)
	}

	return err
}

func (self *UnSubscribe) Write(writer io.Writer) error {
	var header bytes.Buffer
	var body bytes.Buffer
	var err error

	body.Write(encodeUint16(self.MessageID))

	for _, topic := range self.Topics {
		body.Write(encodeString(topic))
	}

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

func (self UnSubscribe) String() string {
	b, _ := json.MarshalIndent(self, "", " ")
	return string(b)
}
