package keyboard

import (
	"errors"

	"github.com/Nibelheims/gometrics/pkg/monitoring"
)

type Keyboard struct {
	VendorID  uint16
	ProductID uint16
	UsagePage uint16
	UsageID   uint16
	Interface int
	Name      string
}

var Lily58 = Keyboard{
	0x04d8,
	0xeb2d,
	0xFF60,
	0x0061,
	1,
	"Lily58",
}

// gometrics
var MAGIC = []byte{0x67, 0x6F, 0x6D, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73}

func UsagesToHIDReport(usages []monitoring.Usage) ([]byte, error) {
	buffLen := len(MAGIC) + monitoring.USAGE_SIZE*len(usages)
	if buffLen >= 32 /* hid report size */ {
		return nil, errors.New("too much data to fit in one report, splitting not supported")
	}
	b := make([]byte, buffLen)

	copy(b[0:len(MAGIC)], MAGIC) // every sent packet will start with this magic id
	start := len(MAGIC)
	for _, u := range usages {
		copy(b[start:start+4], u.Name[:])
		copy(b[start+4:start+4+1], []byte{uint8(u.Percent)})
		start += 4 /*len(Name)*/ + 1 /*1 byte for the percentage*/
	}

	return b, nil
}
