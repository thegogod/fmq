package protocol

type Code byte

const (
	ERROR Code = iota
	CONNECT
	CONNECT_ACK
	DISCONNECT
	PUBLISH
	PUBLISH_ACK
	SUBSCRIBE
	SUBSCRIBE_ACK
	UNSUBSCRIBE
	UNSUBSCRIBE_ACK
	PING
	PING_ACK
)

func (self Code) Valid() bool {
	switch self {
	case ERROR:
		fallthrough
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
	case ERROR:
		return "error"
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
