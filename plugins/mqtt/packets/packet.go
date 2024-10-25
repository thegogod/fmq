package packets

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/thegogod/fmq/common/protocol"
)

type Code byte

const (
	CONNECT         Code = 1
	CONNECT_ACK     Code = 2
	PUBLISH         Code = 3
	PUBLISH_ACK     Code = 4
	PUBLISH_REC     Code = 5
	PUBLISH_REL     Code = 6
	PUBLISH_COMP    Code = 7
	SUBSCRIBE       Code = 8
	SUBSCRIBE_ACK   Code = 9
	UNSUBSCRIBE     Code = 10
	UNSUBSCRIBE_ACK Code = 11
	PING            Code = 12
	PING_ACK        Code = 13
	DISCONNECT      Code = 14
)

func (self Code) Valid() bool {
	switch self {
	case CONNECT:
		fallthrough
	case CONNECT_ACK:
		fallthrough
	case DISCONNECT:
		fallthrough
	case PUBLISH:
		fallthrough
	case PUBLISH_ACK:
		fallthrough
	case PUBLISH_REC:
		fallthrough
	case PUBLISH_REL:
		fallthrough
	case PUBLISH_COMP:
		fallthrough
	case SUBSCRIBE:
		fallthrough
	case SUBSCRIBE_ACK:
		fallthrough
	case UNSUBSCRIBE:
		fallthrough
	case UNSUBSCRIBE_ACK:
		fallthrough
	case PING:
		fallthrough
	case PING_ACK:
		return true
	}

	return false
}

func (self Code) String() string {
	switch self {
	case CONNECT:
		return "connect"
	case CONNECT_ACK:
		return "connect_ack"
	case DISCONNECT:
		return "disconnect"
	case PUBLISH:
		return "publish"
	case PUBLISH_ACK:
		return "publish_ack"
	case PUBLISH_REC:
		return "publish_rec"
	case PUBLISH_REL:
		return "publish_rel"
	case PUBLISH_COMP:
		return "publish_comp"
	case SUBSCRIBE:
		return "subscribe"
	case SUBSCRIBE_ACK:
		return "subscribe_ack"
	case UNSUBSCRIBE:
		return "unsubscribe"
	case UNSUBSCRIBE_ACK:
		return "unsubscribe_ack"
	case PING:
		return "ping"
	case PING_ACK:
		return "ping_ack"
	}

	return ""
}

type Header struct {
	Code            Code `json:"code"`
	Dup             bool `json:"dup"`
	Qos             byte `json:"qos"`
	Retain          bool `json:"retain"`
	RemainingLength int  `json:"remaining_length"`
}

func (self *Header) Read(typeAndFlags byte, reader io.Reader) error {
	self.Code = Code(typeAndFlags >> 4)

	if !self.Code.Valid() {
		return errors.New("invalid packet code")
	}

	self.Dup = (typeAndFlags>>3)&0x01 > 0
	self.Qos = (typeAndFlags >> 1) & 0x03
	self.Retain = typeAndFlags&0x01 > 0

	length, err := decodeLength(reader)

	if err != nil {
		return err
	}

	self.RemainingLength = length
	return nil
}

func (self *Header) Write(writer io.Writer) error {
	var header bytes.Buffer

	header.WriteByte(byte(self.Code)<<4 | boolToByte(self.Dup)<<3 | self.Qos<<1 | boolToByte(self.Retain))
	header.Write(encodeLength(self.RemainingLength))
	_, err := header.WriteTo(writer)

	return err
}

func (self Header) String() string {
	b, _ := json.MarshalIndent(self, "", " ")
	return string(b)
}

type Packet interface {
	fmt.Stringer

	ToProtocol() protocol.Packet
	Read(reader io.Reader) error
	Write(writer io.Writer) error
}

func New(code Code) Packet {
	packet, _ := FromHeader(Header{Code: code})
	return packet
}

func FromHeader(header Header) (Packet, error) {
	switch header.Code {
	case CONNECT:
		return &Connect{Header: header}, nil
	case CONNECT_ACK:
		return &ConnectAck{Header: header}, nil
	case DISCONNECT:
		return &Disconnect{Header: header}, nil
	case PING:
		return &Ping{Header: header}, nil
	case PING_ACK:
		return &PingAck{Header: header}, nil
	case PUBLISH:
		return &Publish{Header: header}, nil
	case PUBLISH_ACK:
		return &PublishAck{Header: header}, nil
	case SUBSCRIBE:
		return &Subscribe{Header: header}, nil
	case SUBSCRIBE_ACK:
		return &SubscribeAck{Header: header}, nil
	case UNSUBSCRIBE:
		return &UnSubscribe{Header: header}, nil
	case UNSUBSCRIBE_ACK:
		return &UnSubscribeAck{Header: header}, nil
	}

	return nil, fmt.Errorf("unsupported packet type 0x%x", header.Code)
}

func FromProtocol(in protocol.Packet) (Packet, error) {
	switch packet := in.(type) {
	case *protocol.Connect:
		return (&Connect{}).FromProtocol(packet), nil
	case *protocol.ConnectAck:
		return (&ConnectAck{}).FromProtocol(packet), nil
	case *protocol.Disconnect:
		return (&Disconnect{}).FromProtocol(packet), nil
	case *protocol.Ping:
		return (&Ping{}).FromProtocol(packet), nil
	case *protocol.PingAck:
		return (&PingAck{}).FromProtocol(packet), nil
	case *protocol.Publish:
		return (&Publish{}).FromProtocol(packet), nil
	case *protocol.PublishAck:
		return (&PublishAck{}).FromProtocol(packet), nil
	case *protocol.Subscribe:
		return (&Subscribe{}).FromProtocol(packet), nil
	case *protocol.SubscribeAck:
		return (&SubscribeAck{}).FromProtocol(packet), nil
	case *protocol.UnSubscribe:
		return (&UnSubscribe{}).FromProtocol(packet), nil
	case *protocol.UnSubscribeAck:
		return (&UnSubscribeAck{}).FromProtocol(packet), nil
	}

	return nil, fmt.Errorf("unsupported packet type 0x%x", in.Code())
}

func Read(reader io.Reader) (Packet, error) {
	header := Header{}
	b := make([]byte, 1)

	if _, err := io.ReadFull(reader, b); err != nil {
		return nil, err
	}

	if err := header.Read(b[0], reader); err != nil {
		return nil, err
	}

	packet, err := FromHeader(header)

	if err != nil {
		return nil, err
	}

	packetBytes := make([]byte, header.RemainingLength)
	n, err := io.ReadFull(reader, packetBytes)

	if err != nil {
		return nil, err
	}

	if n != header.RemainingLength {
		return nil, errors.New("failed to read expected data")
	}

	err = packet.Read(bytes.NewBuffer(packetBytes))
	return packet, err
}

//
// CONVERSIONS
//

func boolToByte(b bool) byte {
	switch b {
	case true:
		return 1
	default:
		return 0
	}
}

//
// DECODING
//

func decodeByte(b io.Reader) (byte, error) {
	num := make([]byte, 1)
	_, err := b.Read(num)
	if err != nil {
		return 0, err
	}

	return num[0], nil
}

func decodeUint16(b io.Reader) (uint16, error) {
	num := make([]byte, 2)
	_, err := b.Read(num)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16(num), nil
}

func encodeUint16(num uint16) []byte {
	bytesResult := make([]byte, 2)
	binary.BigEndian.PutUint16(bytesResult, num)
	return bytesResult
}

func encodeString(field string) []byte {
	return encodeBytes([]byte(field))
}

func decodeString(b io.Reader) (string, error) {
	buf, err := decodeBytes(b)
	return string(buf), err
}

func decodeBytes(b io.Reader) ([]byte, error) {
	fieldLength, err := decodeUint16(b)
	if err != nil {
		return nil, err
	}

	field := make([]byte, fieldLength)
	_, err = b.Read(field)
	if err != nil {
		return nil, err
	}

	return field, nil
}

func encodeBytes(field []byte) []byte {
	fieldLength := make([]byte, 2)
	binary.BigEndian.PutUint16(fieldLength, uint16(len(field)))
	return append(fieldLength, field...)
}

func encodeLength(length int) []byte {
	var encLength []byte
	for {
		digit := byte(length % 128)
		length /= 128
		if length > 0 {
			digit |= 0x80
		}
		encLength = append(encLength, digit)
		if length == 0 {
			break
		}
	}
	return encLength
}

func decodeLength(r io.Reader) (int, error) {
	var rLength uint32
	var multiplier uint32
	b := make([]byte, 1)
	for multiplier < 27 { // fix: Infinite '(digit & 128) == 1' will cause the dead loop
		_, err := io.ReadFull(r, b)
		if err != nil {
			return 0, err
		}

		digit := b[0]
		rLength |= uint32(digit&127) << multiplier
		if (digit & 128) == 0 {
			break
		}
		multiplier += 7
	}
	return int(rLength), nil
}
