package main

import (
	"fmt"
	"strings"

	"github.com/Nibelheims/gometrics/pkg/keyboard"
	"github.com/Nibelheims/gometrics/pkg/monitoring"
	"github.com/karalabe/hid"
)

func main() {
	// Enumerate all the HID devices in alphabetical path order
	if !hid.Supported() {
		fmt.Println("HID lib not supported on this platform, exiting")
		return
	}
	hids := hid.Enumerate(keyboard.Lily58.VendorID, keyboard.Lily58.ProductID)
	if hids == nil {
		panic("failed to enumerate HID devices")
	}
	for i := 0; i < len(hids); i++ {
		for j := i + 1; j < len(hids); j++ {
			if hids[i].Path > hids[j].Path {
				hids[i], hids[j] = hids[j], hids[i]
			}
		}
	}
	fmt.Printf("hid.Supported() %v\n", hid.Supported())
	for i, hid := range hids {
		fmt.Println(strings.Repeat("-", 128))
		fmt.Printf("HID #%d\n", i)
		fmt.Printf("  OS Path:      %s\n", hid.Path)
		fmt.Printf("  Vendor ID:    %#04x\n", hid.VendorID)
		fmt.Printf("  Product ID:   %#04x\n", hid.ProductID)
		fmt.Printf("  Release:      %d\n", hid.Release)
		fmt.Printf("  Serial:       %s\n", hid.Serial)
		fmt.Printf("  Manufacturer: %s\n", hid.Manufacturer)
		fmt.Printf("  Product:      %s\n", hid.Product)
		fmt.Printf("  Usage Page:   %#04x\n", hid.UsagePage)
		fmt.Printf("  Usage:        %d\n", hid.Usage)
		fmt.Printf("  Interface:    %d\n", hid.Interface)
	}
	fmt.Println(strings.Repeat("=", 128))

	m := monitoring.NewMonitor(500)
	m.Run()

	i := 0
	for u := range m.C() {
		fmt.Printf("cpu: %f%%\tmem: %f%%\n", u.CpuPercent, u.MemPercent)
		i++
		if i > 10 {
			m.Stop()
		}
	}
}
