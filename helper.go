package main

import (
	"log"

	"github.com/PathFinder5/protobuf-list-devices/device"
	"github.com/google/gousb"
	"github.com/google/gousb/usbid"
)


func GetDevices() ([]*device.Device, error) {
	ctx := gousb.NewContext()
	defer ctx.Close()

	// Open all USB devices
	devs, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		// The usbid package can be used to print out human readable information.
		// fmt.Printf("%03d.%03d %s:%s %s\n", desc.Bus, desc.Address, desc.Vendor, desc.Product, usbid.Describe(desc))
		
		// return true
		// force return true to list all devices because otherwise it will not return any devices
		return true
	})

	if err != nil {
		log.Printf("Error listing devices: %v", err)

		return nil, err
	}
	defer func() {
		for _, dev := range devs {
			dev.Close()
		}
	}()

	var deviceList []*device.Device
	for _, dev := range devs {
		desc := dev.Desc
		devType := usbid.Describe(desc)
		devPath := desc.String()
		vendorID := desc.Vendor.String()
		productID := desc.Product.String()

		dev := &device.Device{
			Type:      devType,
			Path:      devPath,
			VendorId:  vendorID,
			ProductId: productID,
		}
		deviceList = append(deviceList, dev)
	}

	log.Printf("All %d device(s) listed successfully!", len(deviceList))

	return deviceList, nil
}