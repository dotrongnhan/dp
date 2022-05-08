package main

import "adapter/adapter"

func main() {
	client := &adapter.Client{}
	window := &adapter.Window{}
	client.InsertSquareUSBInComputer(window)

	mac := &adapter.MacAdapter{
		MacMachine: adapter.Mac{},
	}
	client.InsertSquareUSBInComputer(mac)
}
