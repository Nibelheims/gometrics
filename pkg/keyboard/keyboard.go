package keyboard

type Keyboard struct {
	VendorID  uint16
	ProductID uint16
	UsagePage uint16
	UsageID   uint16
	Name      string
}

var Lily58 = Keyboard{
	0x04d8,
	0xeb2d,
	0xFF60,
	0x0061,
	"Lily58",
}
