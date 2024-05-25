package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PathFinder5/protobuf-list-devices/device"
	"google.golang.org/protobuf/proto"
)

func handleGetDevices(){
	http.HandleFunc("/devices", func(w http.ResponseWriter, r *http.Request) {
        devices, err := GetDevices()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        deviceList := &device.ListAvailableDevices{
            Devices: devices,
        }

        data, err := proto.Marshal(deviceList)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/x-protobuf")
        w.Write(data)
		log.Println("Write response to client...")

    })
}

func handleJsonGetDevices(){
	http.HandleFunc("/jsondevices", func(w http.ResponseWriter, r *http.Request) {
		devices, err := GetDevices()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonDeviceList := &device.ListAvailableDevices{
			Devices: devices,
		}

		jsonData, err := json.Marshal(jsonDeviceList)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	})
}
