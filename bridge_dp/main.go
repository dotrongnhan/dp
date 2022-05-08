package main

import "bridge/bridge"

func main() {
	hpPrinter := &bridge.HP{}
	epsonPrinter := &bridge.Epson{}

	macComputer := &bridge.Mac{}
	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()

	windowComputer := &bridge.Windows{}

	windowComputer.SetPrinter(hpPrinter)
	windowComputer.Print()

	windowComputer.SetPrinter(epsonPrinter)
	windowComputer.Print()

}
