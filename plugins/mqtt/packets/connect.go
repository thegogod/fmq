package packets

import (
	"bytes"
	"encoding/json"
	"io"
	"slices"

	"github.com/thegogod/fmq/common/protocol"
)

type Connect struct {
	Header `json:"header"`

	ProtocolName    string `json:"protocol_name"`
	ProtocolVersion byte   `json:"protocol_version"`
	CleanSession    bool   `json:"clean_session"`
	WillFlag        bool   `json:"will_flag"`
	WillQos         byte   `json:"will_qos"`
	WillRetain      bool   `json:"will_retain"`
	UsernameFlag    bool   `json:"username_flag"`
	PasswordFlag    bool   `json:"password_flag"`
	ReservedBit     byte   `json:"reserved_bit"`
	Keepalive       uint16 `json:"keep_alive"`

	ClientIdentifier string `json:"client_identifier"`
	WillTopic        string `json:"will_topic"`
	WillMessage      []byte `json:"will_message"`
	Username         string `json:"username"`
	Password         []byte `json:"-"`
}

func (self Connect) ToProtocol() protocol.Packet {
	return &protocol.Connect{
		ClientID: self.ClientIdentifier,
		Username: self.Username,
		Password: string(self.Password),
	}
}

func (self *Connect) FromProtocol(packet *protocol.Connect) *Connect {
	self.Header = Header{Code: CONNECT}
	self.ClientIdentifier = packet.ClientID
	self.Username = packet.Username
	self.Password = []byte(packet.Password)
	return self
}

func (self *Connect) Read(reader io.Reader) error {
	var err error
	self.ProtocolName, err = decodeString(reader)

	if err != nil {
		return err
	}

	self.ProtocolVersion, err = decodeByte(reader)

	if err != nil {
		return err
	}

	options, err := decodeByte(reader)

	if err != nil {
		return err
	}

	self.ReservedBit = 1 & options
	self.CleanSession = 1&(options>>1) > 0
	self.WillFlag = 1&(options>>2) > 0
	self.WillQos = 3 & (options >> 3)
	self.WillRetain = 1&(options>>5) > 0
	self.PasswordFlag = 1&(options>>6) > 0
	self.UsernameFlag = 1&(options>>7) > 0
	self.Keepalive, err = decodeUint16(reader)

	if err != nil {
		return err
	}

	self.ClientIdentifier, err = decodeString(reader)

	if err != nil {
		return err
	}

	if self.WillFlag {
		self.WillTopic, err = decodeString(reader)

		if err != nil {
			return err
		}

		self.WillMessage, err = decodeBytes(reader)

		if err != nil {
			return err
		}
	}

	if self.UsernameFlag {
		self.Username, err = decodeString(reader)

		if err != nil {
			return err
		}
	}

	if self.PasswordFlag {
		self.Password, err = decodeBytes(reader)

		if err != nil {
			return err
		}
	}

	return nil
}

func (self *Connect) Write(writer io.Writer) error {
	var header bytes.Buffer
	var body bytes.Buffer
	var err error

	body.Write(encodeString(self.ProtocolName))
	body.WriteByte(self.ProtocolVersion)
	body.WriteByte(boolToByte(self.CleanSession)<<1 | boolToByte(self.WillFlag)<<2 | self.WillQos<<3 | boolToByte(self.WillRetain)<<5 | boolToByte(self.PasswordFlag)<<6 | boolToByte(self.UsernameFlag)<<7)
	body.Write(encodeUint16(self.Keepalive))
	body.Write(encodeString(self.ClientIdentifier))

	if self.WillFlag {
		body.Write(encodeString(self.WillTopic))
		body.Write(encodeBytes(self.WillMessage))
	}

	if self.UsernameFlag {
		body.Write(encodeString(self.Username))
	}

	if self.PasswordFlag {
		body.Write(encodeBytes(self.Password))
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

func (self Connect) ReturnCode() ConnectReturnCode {
	if self.PasswordFlag && !self.UsernameFlag {
		return ErrRefusedBadUsernameOrPassword
	}

	if self.ReservedBit != 0 {
		return ErrProtocolViolation
	}

	if (self.ProtocolName == "MQIsdp" && self.ProtocolVersion != 3) || (self.ProtocolName == "MQTT" && self.ProtocolVersion != 4) {
		return ErrRefusedBadProtocolVersion
	}

	if self.ProtocolName != "MQIsdp" && self.ProtocolName != "MQTT" {
		return ErrProtocolViolation
	}

	if len(self.ClientIdentifier) > 65535 || len(self.Username) > 65535 || len(self.Password) > 65535 {
		return ErrProtocolViolation
	}

	if len(self.ClientIdentifier) == 0 && !self.CleanSession {
		return ErrRefusedIDRejected
	}

	return Accepted
}

func (self Connect) String() string {
	b, _ := json.MarshalIndent(self, "", " ")
	return string(b)
}
