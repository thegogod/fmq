package packets

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"slices"

	"github.com/thegogod/fmq/common/protocol"
)

type ConnectReturnCode byte

const (
	Accepted                        ConnectReturnCode = 0x00
	ErrRefusedBadProtocolVersion    ConnectReturnCode = 0x01
	ErrRefusedIDRejected            ConnectReturnCode = 0x02
	ErrRefusedServerUnavailable     ConnectReturnCode = 0x03
	ErrRefusedBadUsernameOrPassword ConnectReturnCode = 0x04
	ErrRefusedNotAuthorised         ConnectReturnCode = 0x05
	ErrNetworkError                 ConnectReturnCode = 0xFE
	ErrProtocolViolation            ConnectReturnCode = 0xFF
)

func (self ConnectReturnCode) Valid() bool {
	switch self {
	case Accepted:
		fallthrough
	case ErrRefusedBadProtocolVersion:
		fallthrough
	case ErrRefusedIDRejected:
		fallthrough
	case ErrRefusedServerUnavailable:
		fallthrough
	case ErrRefusedBadUsernameOrPassword:
		fallthrough
	case ErrRefusedNotAuthorised:
		fallthrough
	case ErrNetworkError:
		fallthrough
	case ErrProtocolViolation:
		return true
	}

	return false
}

type ConnectAck struct {
	Header `json:"header"`

	SessionPresent bool              `json:"session_present"`
	ReturnCode     ConnectReturnCode `json:"return_code"`
}

func (self ConnectAck) ToProtocol() protocol.Packet {
	return &protocol.ConnectAck{ReturnCode: byte(self.ReturnCode)}
}

func (self *ConnectAck) FromProtocol(packet *protocol.ConnectAck) *ConnectAck {
	self.Header = Header{Code: CONNECT_ACK}
	self.ReturnCode = ConnectReturnCode(packet.ReturnCode)
	return self
}

func (self *ConnectAck) Read(reader io.Reader) error {
	flags, err := decodeByte(reader)

	if err != nil {
		return err
	}

	self.SessionPresent = 1&flags > 0
	code, err := decodeByte(reader)

	if err != nil {
		return err
	}

	self.ReturnCode = ConnectReturnCode(code)

	if !self.ReturnCode.Valid() {
		return errors.New("invalid connect return code")
	}

	return nil
}

func (self *ConnectAck) Write(writer io.Writer) error {
	var body bytes.Buffer
	var header bytes.Buffer

	body.WriteByte(boolToByte(self.SessionPresent))
	body.WriteByte(byte(self.ReturnCode))

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

func (self ConnectAck) String() string {
	b, _ := json.MarshalIndent(self, "", " ")
	return string(b)
}
