package main

import (
	"io"
	"log"
	"net/http"

	"github.com/PathFinder5/protobuf-list-devices/device"

	"google.golang.org/protobuf/proto"
)

// GetDevicesFromServer sends an HTTP GET request to the server's /devices endpoint
// and returns the list of devices received in the response.
func GetDevicesFromServer() (*device.ListAvailableDevices, error) {
    resp, err := http.Get("http://localhost:8080/devices")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    // Read the response body
    data, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    // Unmarshal the protobuf-encoded response
    var deviceList device.ListAvailableDevices
    err = proto.Unmarshal(data, &deviceList)
    if err != nil {
        return nil, err
    }

    return &deviceList, nil
}


func main() {
    
    // Get the list of devices from the server
    devices, err := GetDevicesFromServer()
    if err != nil {
        log.Fatalf("Error getting devices from server: %v", err)
    }

    // Print the list of devices
    for _, d := range devices.Devices {
        log.Printf("Type: %s, Path: %s, Vendor ID: %s, Product ID: %s", d.Type, d.Path, d.VendorId, d.ProductId)
    }

}

