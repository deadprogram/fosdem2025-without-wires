package main

import (
	"time"

	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

func main() {
	must("enable BLE stack", adapter.Enable())
	adv := adapter.DefaultAdvertisement()
	must("config adv", adv.Configure(bluetooth.AdvertisementOptions{
		AdvertisementType: bluetooth.AdvertisingTypeScanInd,
		LocalName:         "Go Bluetooth",
	}))
	must("start adv", adv.Start())

	println("starting advertising...")
	for {
		println("Go Bluetooth is advertising.")
		time.Sleep(time.Second)
	}
}

func must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}
