package main

import (
	"fmt"

	gpsgenpb "github.com/mmadfox/go-gpsgen/proto"
	"github.com/mmadfox/gpsgend/pkg/websocket"
)

func main() {
	addr := "ws://127.0.0.1:32300"
	websocket.Listen(addr,
		func(dev *gpsgenpb.Device) {
			fmt.Printf("device id=%s, lat=%f, lon=%f, alt=%f\n",
				dev.Id,
				dev.Location.Lat,
				dev.Location.Lon,
				dev.Location.Alt)
		},
		func(err error) {
			fmt.Printf("[ERROR] %v\n", err)
		})
}
