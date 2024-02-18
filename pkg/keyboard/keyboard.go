package keyboard

import (
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

func UsageToHIDReport(u monitoring.Usage) []byte {
	b := make([]byte, len(MAGIC)+2)

	copy(b[0:len(MAGIC)], MAGIC)
	copy(b[len(MAGIC)+0:len(MAGIC)+1], []byte{uint8(u.CpuPercent)})
	copy(b[len(MAGIC)+1:len(MAGIC)+2], []byte{uint8(u.MemPercent)})

	return b
}
