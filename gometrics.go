package main

import (
	"errors"
	"fmt"

	"github.com/Nibelheims/gometrics/pkg/keyboard"
	"github.com/Nibelheims/gometrics/pkg/monitoring"
	"github.com/karalabe/hid"
)

func getHID(k *keyboard.Keyboard) (*hid.DeviceInfo, error) {
	hids, err := hid.Enumerate(keyboard.Lily58.VendorID, keyboard.Lily58.ProductID)
	if err != nil {
		return nil, errors.New("failed to enumerate HID devices")
	}
	for _, hid := range hids {
		if hid.UsagePage == keyboard.Lily58.UsagePage &&
			hid.Usage == keyboard.Lily58.UsageID {
			return &hid, nil
		}
	}
    // on linux libusb typically do not fetch usage page and id
    // use the interface number instead
	for _, hid := range hids {
        if hid.Interface == keyboard.Lily58.Interface {
			return &hid, nil
        }
    }
	return nil, errors.New("Could not find the device \"" + k.Name + "\"")
}

func main() {
	if !hid.Supported() {
		fmt.Println("HID lib not supported on this platform, exiting")
		return
	}

	keebInfo, err := getHID(&keyboard.Lily58)
	if err != nil {
		fmt.Println(err)
		return
	}

	keeb, err := keebInfo.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer keeb.Close()
	fmt.Println("Successfully opened  \"" + keyboard.Lily58.Name + "\"")

	m := monitoring.NewMonitor(500)
	m.Run()

	i := 0
	for u := range m.C() {
		fmt.Printf("cpu: %f%%\tmem: %f%%\n", u.CpuPercent, u.MemPercent)
		buffer := keyboard.UsageToHIDReport(u)
		//fmt.Println(buffer)
		keeb.Write(buffer)
        //report := append([]byte{0x00}, buffer...)
        //keeb.SendFeatureReport(report)
		i++
		if i > 10 {
			m.Stop()
		}
	}
}
