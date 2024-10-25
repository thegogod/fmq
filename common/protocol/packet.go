package protocol

import (
	"fmt"
)

type Packet interface {
	fmt.Stringer

	Code() Code
}
