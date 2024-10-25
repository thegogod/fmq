package packets

import (
	"bytes"
	"encoding/json"
	"io"
	"slices"

	"github.com/thegogod/fmq/common/protocol"
)

type Subscribe struct {
	Header `json:"header"`

	MessageID uint16   `json:"message_id"`
	Topics    []string `json:"topics"`
	Qoss      []byte   `json:"qoss"`
}

func (self Subscribe) ToProtocol() protocol.Packet {
	return &protocol.Subscribe{
		ID:     self.MessageID,
		Topics: self.Topics,
	}
}

func (self *Subscribe) FromProtocol(packet *protocol.Subscribe) *Subscribe {
	self.Header = Header{Code: SUBSCRIBE}
	self.MessageID = packet.ID
	self.Topics = packet.Topics
	return self
}

func (self *Subscribe) Read(reader io.Reader) error {
	var err error
	self.MessageID, err = decodeUint16(reader)

	if err != nil {
		return err
	}

	payloadLength := self.Header.RemainingLength - 2

	for payloadLength > 0 {
		topic, err := decodeString(reader)

		if err != nil {
			return err
		}

		self.Topics = append(self.Topics, topic)
		qos, err := decodeByte(reader)

		if err != nil {
			return err
		}

		self.Qoss = append(self.Qoss, qos)
		payloadLength -= 2 + len(topic) + 1
	}

	return nil
}

func (self *Subscribe) Write(writer io.Writer) error {
	var header bytes.Buffer
	var body bytes.Buffer
	var err error

	body.Write(encodeUint16(self.MessageID))

	for i, topic := range self.Topics {
		body.Write(encodeString(topic))
		body.WriteByte(self.Qoss[i])
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

func (self Subscribe) String() string {
	b, _ := json.MarshalIndent(self, "", " ")
	return string(b)
}
